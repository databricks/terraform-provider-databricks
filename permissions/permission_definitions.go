package permissions

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
)

// resourcePermissions captures all the information needed to manage permissions for a given object type.
type resourcePermissions struct {
	// Mandatory Fields

	// The attribute name that users configure with the ID of the object to manage
	// e.g. "cluster_id" for a cluster
	field string
	// The object type to use in the Permissions API, e.g. "cluster" for a cluster.
	objectType string
	// The name of the object in the ID of the TF resource, e.g. "clusters" for a cluster,
	// where the ID would be /clusters/<cluster-id>. This should also match the prefix of the
	// object ID in the API response, unless idMatcher is set.
	requestObjectType string
	// The allowed permission levels for this object type and its options.
	allowedPermissionLevels map[string]permissionLevelOptions

	// ID Remapping Options

	// Returns the object ID for the given user-specified ID. This is necessary because permissions for
	// some objects are done by path, whereas others are by ID. Those by path need to be converted to the
	// internal object ID before being stored in the state. If not specified, the default ID is "/<resource_type>/<id>".
	idRetriever func(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error)
	// By default, a resourcePermissions can be retrieved based on the structure of the ID, as described above.
	// If this function is set, it will be used to determine whether the ID matches this resource type.
	idMatcher func(id string) bool
	// A custom matcher to check whether a given ID matches this resource type.
	// Most resources can be determined by looking at the attribute name used to configure the permission, but
	// tokens & passwords are special cases where the resource type is determined by the value of this attribute.
	stateMatcher func(id string) bool

	// Behavior Options and Customizations

	// The alternative name of the "path" attribute for this resource. E.g. "workspace_file_path" for a file.
	// If not set, default is "<object_type>_path".
	pathVariant string
	// Customizers when handling permission resource creation and update.
	//
	// Most resources that have a CAN_MANAGE permission level should add addCurrentUserAsManageCustomizer to this list
	// to ensure that the user applying the template always has management permissions on the underlying resource.
	updateAclCustomizers []aclUpdateCustomizer
	// Customizers when handling permission resource deletion.
	//
	// Most resources that have a CAN_MANAGE permission level should add addCurrentUserAsManageCustomizer to this list
	// to ensure that the user applying the template always has management permissions on the underlying resource.
	deleteAclCustomizers []aclUpdateCustomizer
	// Customizers when handling permission resource read.
	//
	// Resources for which admins inherit permissions should add removeAdminPermissionsCustomizer to this list. This
	// prevents the admin group from being included in the permissions when reading the state.
	readAclCustomizers []aclReadCustomizer

	// Returns the creator of the object. Used when deleting databricks_permissions resources, when the
	// creator of the object is restored as the owner.
	fetchObjectCreator func(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error)
}

// getAllowedPermissionLevels returns the list of permission levels that are allowed for this resource type.
func (p resourcePermissions) getAllowedPermissionLevels(includeNonManagementPermissions bool) []string {
	levels := make([]string, 0, len(p.allowedPermissionLevels))
	for level := range p.allowedPermissionLevels {
		if includeNonManagementPermissions || p.allowedPermissionLevels[level].isManagementPermission {
			levels = append(levels, level)
		}
	}
	sort.Strings(levels)
	return levels
}

// resourceStatus captures the status of a resource with permissions. If the resource doesn't exist,
// the provider will not try to update its permissions. Otherwise, the creator will be returned if
// it can be determined for the given resource type.
type resourceStatus struct {
	exists  bool
	creator string
}

// getObjectStatus returns the creator of the object and whether the object exists. If the object creator cannot be determined for this
// resource type, an empty string is returned. Resources without fetchObjectCreator are assumed to exist and have an unknown creator.
func (p resourcePermissions) getObjectStatus(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (resourceStatus, error) {
	if p.fetchObjectCreator != nil {
		creator, err := p.fetchObjectCreator(ctx, w, objectID)
		if err != nil {
			return resourceStatus{}, err
		}
		return resourceStatus{exists: creator != "", creator: creator}, nil
	}
	return resourceStatus{exists: true, creator: ""}, nil
}

// getPathVariant returns the name of the path attribute for this resource type.
func (p resourcePermissions) getPathVariant() string {
	if p.pathVariant != "" {
		return p.pathVariant
	}
	return p.objectType + "_path"
}

// validate checks that the user is not trying to set permissions for the admin group or remove their own management permissions.
func (p resourcePermissions) validate(entity PermissionsEntity, currentUsername string) error {
	for _, change := range entity.AccessControlList {
		// Check if the user is trying to set permissions for the admin group outside of the tokens/passwords permissions.
		if change.GroupName == "admins" && p.field != "authorization" {
			return fmt.Errorf("it is not possible to modify admin permissions for %s resources", p.objectType)
		}
		// Check that the user is preventing themselves from managing the object
		if (change.UserName == currentUsername || change.ServicePrincipalName == currentUsername) && !p.allowedPermissionLevels[string(change.PermissionLevel)].isManagementPermission {
			allowedLevelsForCurrentUser := p.getAllowedPermissionLevels(false)
			return fmt.Errorf("cannot remove management permissions for the current user for %s, allowed levels: %s", p.objectType, strings.Join(allowedLevelsForCurrentUser, ", "))
		}
	}
	return nil
}

// getID returns the object ID for the given user-specified ID.
func (p resourcePermissions) getID(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error) {
	var err error
	if p.idRetriever != nil {
		id, err = p.idRetriever(ctx, w, id)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("/%s/%s", p.requestObjectType, id), nil
}

// prepareForUpdate prepares the access control list for an update request by calling all update customizers.
func (p resourcePermissions) prepareForUpdate(objectID string, entity PermissionsEntity, currentUser string) (PermissionsEntity, error) {
	cachedCurrentUser := func() (string, error) { return currentUser, nil }
	ctx := aclUpdateCustomizerContext{
		getCurrentUser: cachedCurrentUser,
		getId:          func() string { return objectID },
	}
	var err error
	for _, customizer := range p.updateAclCustomizers {
		entity.AccessControlList, err = customizer(ctx, entity.AccessControlList)
		if err != nil {
			return PermissionsEntity{}, err
		}
	}
	return entity, nil
}

// prepareForDelete prepares the access control list for a delete request by calling all delete customizers.
func (p resourcePermissions) prepareForDelete(objectACL *iam.ObjectPermissions, getCurrentUser func() (string, error)) ([]iam.AccessControlRequest, error) {
	accl := make([]iam.AccessControlRequest, 0, len(objectACL.AccessControlList))
	// By default, only admins have access to a resource when databricks_permissions for that resource are deleted.
	for _, acl := range objectACL.AccessControlList {
		if acl.GroupName != "admins" {
			continue
		}
		for _, permission := range acl.AllPermissions {
			if !permission.Inherited {
				// keep everything direct for admin group
				accl = append(accl, iam.AccessControlRequest{
					GroupName:       acl.GroupName,
					PermissionLevel: permission.PermissionLevel,
				})
				break
			}
		}
	}
	ctx := aclUpdateCustomizerContext{
		getCurrentUser: getCurrentUser,
		getId:          func() string { return objectACL.ObjectId },
	}
	var err error
	for _, customizer := range p.deleteAclCustomizers {
		accl, err = customizer(ctx, accl)
		if err != nil {
			return nil, err
		}
	}
	return accl, nil
}

// prepareResponse prepares the access control list for a read response by calling all read customizers.
//
// If the user does not include an access_control block for themselves, it will not be included in the state. This
// prevents diffs when the applying user is not included in the access_control block for the resource but is
// added by addCurrentUserAsManageCustomizer.
//
// Read customizers are able to access the current state of the object in order to customize the response accordingly.
// For example, the SQL API previously used CAN_VIEW for read-only permission, but the GA API uses CAN_READ. Users may
// have CAN_VIEW in their resource configuration, so the read customizer will rewrite the response from CAN_READ to
// CAN_VIEW to match the user's configuration.
func (p resourcePermissions) prepareResponse(objectID string, objectACL *iam.ObjectPermissions, existing PermissionsEntity, me string) (PermissionsEntity, error) {
	ctx := aclReadCustomizerContext{
		getId:                        func() string { return objectID },
		getExistingPermissionsEntity: func() PermissionsEntity { return existing },
	}
	if objectACL.ObjectType != p.objectType {
		return PermissionsEntity{}, fmt.Errorf("expected object type %s, got %s", p.objectType, objectACL.ObjectType)
	}
	for _, customizer := range p.readAclCustomizers {
		var err error
		objectACL, err = customizer(ctx, objectACL)
		if err != nil {
			return PermissionsEntity{}, err
		}
	}
	entity := PermissionsEntity{}
	for _, accessControl := range objectACL.AccessControlList {
		// If the user doesn't include an access_control block for themselves, do not include it in the state.
		// On create/update, the provider will automatically include the current user in the access_control block
		// for appropriate resources. Otherwise, it must be included in state to prevent configuration drift.
		if me == accessControl.UserName || me == accessControl.ServicePrincipalName {
			if !existing.containsUserOrServicePrincipal(me) {
				continue
			}
		}
		entity.AccessControlList = append(entity.AccessControlList, iam.AccessControlRequest{
			GroupName:            accessControl.GroupName,
			UserName:             accessControl.UserName,
			ServicePrincipalName: accessControl.ServicePrincipalName,
			PermissionLevel:      accessControl.AllPermissions[0].PermissionLevel,
		})
	}
	return entity, nil
}

// addOwnerPermissionIfNeeded adds the owner permission to the object ACL if the owner permission is allowed and not already set.
func (p resourcePermissions) addOwnerPermissionIfNeeded(objectACL []iam.AccessControlRequest, ownerOpt string) []iam.AccessControlRequest {
	_, ok := p.allowedPermissionLevels["IS_OWNER"]
	if !ok {
		return objectACL
	}

	for _, acl := range objectACL {
		if acl.PermissionLevel == "IS_OWNER" {
			return objectACL
		}
	}

	return append(objectACL, iam.AccessControlRequest{
		UserName:        ownerOpt,
		PermissionLevel: "IS_OWNER",
	})
}

// permissionLevelOptions indicates the properties of a permissions level. Today, the only property
// is whether the current user can set the permission level for themselves.
type permissionLevelOptions struct {
	// Whether users with this permission level are allowed to manage the resource.
	// For some resources where ACLs don't define who can manage the resource, this might be unintuitive,
	// e.g. all cluster policies permissions are considered management permissions because cluster policy
	// ACLs don't define who can manage the cluster policy.
	isManagementPermission bool
}

func getResourcePermissionsFromId(id string) (resourcePermissions, error) {
	idParts := strings.Split(id, "/")
	objectType := strings.Join(idParts[1:len(idParts)-1], "/")
	for _, mapping := range allResourcePermissions() {
		if mapping.idMatcher != nil {
			if mapping.idMatcher(id) {
				return mapping, nil
			}
			continue
		}
		if mapping.requestObjectType == objectType {
			return mapping, nil
		}
	}
	return resourcePermissions{}, fmt.Errorf("resource type for %s not found", id)
}

// getResourcePermissionsFromState returns the resourcePermissions for the given state.
func getResourcePermissionsFromState(d interface{ GetOk(string) (any, bool) }) (resourcePermissions, string, error) {
	allPermissions := allResourcePermissions()
	for _, mapping := range allPermissions {
		if v, ok := d.GetOk(mapping.field); ok {
			id := v.(string)
			if mapping.stateMatcher != nil && !mapping.stateMatcher(id) {
				continue
			}
			return mapping, id, nil
		}
	}
	allFields := make([]string, 0, len(allPermissions))
	seen := make(map[string]struct{})
	for _, mapping := range allPermissions {
		if _, ok := seen[mapping.field]; ok {
			continue
		}
		seen[mapping.field] = struct{}{}
		allFields = append(allFields, mapping.field)
	}
	sort.Strings(allFields)
	return resourcePermissions{}, "", fmt.Errorf("at least one type of resource identifier must be set; allowed fields: %s", strings.Join(allFields, ", "))
}

// getResourcePermissionsForObjectAcl returns the resourcePermissions for the given ObjectAclApiResponse.
// allResourcePermissions is the list of all resource types that can be managed by the databricks_permissions resource.
func allResourcePermissions() []resourcePermissions {
	PATH := func(ctx context.Context, w *databricks.WorkspaceClient, path string) (string, error) {
		info, err := w.Workspace.GetStatusByPath(ctx, path)
		if err != nil {
			return "", fmt.Errorf("cannot load path %s: %s", path, err)
		}
		return strconv.FormatInt(info.ObjectId, 10), nil
	}
	return []resourcePermissions{
		{
			field:             "cluster_policy_id",
			objectType:        "cluster-policy",
			requestObjectType: "cluster-policies",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {isManagementPermission: true},
			},
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "instance_pool_id",
			objectType:        "instance-pool",
			requestObjectType: "instance-pools",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_ATTACH_TO": {isManagementPermission: false},
				"CAN_MANAGE":    {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			readAclCustomizers:   []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "cluster_id",
			objectType:        "cluster",
			requestObjectType: "clusters",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_ATTACH_TO": {isManagementPermission: false},
				"CAN_RESTART":   {isManagementPermission: false},
				"CAN_MANAGE":    {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			readAclCustomizers:   []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "pipeline_id",
			objectType:        "pipelines",
			requestObjectType: "pipelines",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_VIEW":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				"IS_OWNER":   {isManagementPermission: true},
			},
			fetchObjectCreator: func(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error) {
				pipeline, err := w.Pipelines.GetByPipelineId(ctx, strings.ReplaceAll(objectID, "/pipelines/", ""))
				if err != nil {
					return "", common.IgnoreNotFoundError(err)
				}
				return pipeline.CreatorUserName, nil
			},
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "job_id",
			objectType:        "job",
			requestObjectType: "jobs",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_VIEW":       {isManagementPermission: false},
				"CAN_MANAGE_RUN": {isManagementPermission: false},
				"IS_OWNER":       {isManagementPermission: true},
				"CAN_MANAGE":     {isManagementPermission: true},
			},
			fetchObjectCreator: func(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error) {
				jobId, err := strconv.ParseInt(strings.ReplaceAll(objectID, "/jobs/", ""), 10, 64)
				if err != nil {
					return "", err
				}
				job, err := w.Jobs.GetByJobId(ctx, jobId)
				if err != nil {
					return "", common.IgnoreNotFoundError(err)
				}
				return job.CreatorUserName, nil
			},
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "notebook_id",
			objectType:        "notebook",
			requestObjectType: "notebooks",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "notebook_path",
			objectType:        "notebook",
			requestObjectType: "notebooks",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			idRetriever:        PATH,
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "directory_id",
			objectType:        "directory",
			requestObjectType: "directories",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addAdminAclCustomizer(func(objectId string) bool {
					return objectId == "/directories/0"
				}),
			},
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "directory_path",
			objectType:        "directory",
			requestObjectType: "directories",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			idRetriever:        PATH,
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "workspace_file_id",
			objectType:        "file",
			requestObjectType: "files",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			pathVariant:        "workspace_file_path",
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "workspace_file_path",
			objectType:        "file",
			requestObjectType: "files",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			idRetriever:        PATH,
			pathVariant:        "workspace_file_path",
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "repo_id",
			objectType:        "repo",
			requestObjectType: "repos",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "repo_path",
			objectType:        "repo",
			requestObjectType: "repos",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			idRetriever:        PATH,
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "authorization",
			objectType:        "tokens",
			requestObjectType: "authorization",
			stateMatcher: func(id string) bool {
				return id == "tokens"
			},
			idMatcher: func(id string) bool {
				return id == "/authorization/tokens"
			},
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":    {isManagementPermission: true},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addAdminAclCustomizer(func(objectId string) bool {
					return objectId == "/authorization/tokens"
				}),
			},
		},
		{
			field:             "authorization",
			objectType:        "passwords",
			requestObjectType: "authorization",
			stateMatcher: func(id string) bool {
				return id == "passwords"
			},
			idMatcher: func(id string) bool {
				return id == "/authorization/passwords"
			},
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {isManagementPermission: true},
			},
		},
		{
			field:             "sql_endpoint_id",
			objectType:        "warehouses",
			requestObjectType: "sql/warehouses",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":     {isManagementPermission: false},
				"CAN_MANAGE":  {isManagementPermission: true},
				"CAN_MONITOR": {isManagementPermission: false},
				"IS_OWNER":    {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			readAclCustomizers:   []aclReadCustomizer{removeAdminPermissionsCustomizer},
			fetchObjectCreator: func(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error) {
				warehouse, err := w.Warehouses.GetById(ctx, strings.ReplaceAll(objectID, "/sql/warehouses/", ""))
				if err != nil {
					return "", common.IgnoreNotFoundError(err)
				}
				return warehouse.CreatorName, nil
			},
		},
		{
			field:             "sql_dashboard_id",
			objectType:        "dashboard",
			requestObjectType: "dbsql-dashboards",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				// This was originally called CAN_VIEW in the preview API, but was renamed to CAN_READ in the GA API.
				"CAN_VIEW": {isManagementPermission: false},
			},
			idMatcher: func(id string) bool {
				return strings.HasPrefix(id, "/dbsql-dashboards/") || strings.HasPrefix(id, "/sql/dashboards/")
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				rewritePermissionForUpdate("CAN_VIEW", "CAN_READ"),
			},
			deleteAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				rewritePermissionForUpdate("CAN_VIEW", "CAN_READ"),
			},
			readAclCustomizers: []aclReadCustomizer{
				removeAdminPermissionsCustomizer,
				rewritePermissionForRead("CAN_READ", "CAN_VIEW"),
			},
		},
		{
			field:             "sql_alert_id",
			objectType:        "alert",
			requestObjectType: "sql/alerts",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				// This was originally called CAN_VIEW in the preview API, but was renamed to CAN_READ in the GA API.
				"CAN_VIEW": {isManagementPermission: false},
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				rewritePermissionForUpdate("CAN_VIEW", "CAN_READ"),
			},
			deleteAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				rewritePermissionForUpdate("CAN_VIEW", "CAN_READ"),
			},
			readAclCustomizers: []aclReadCustomizer{
				removeAdminPermissionsCustomizer,
				rewritePermissionForRead("CAN_READ", "CAN_VIEW"),
			},
		},
		{
			field:             "sql_query_id",
			objectType:        "query",
			requestObjectType: "sql/queries",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				// This was originally called CAN_VIEW in the preview API, but was renamed to CAN_READ in the GA API.
				"CAN_VIEW": {isManagementPermission: false},
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				rewritePermissionForUpdate("CAN_VIEW", "CAN_READ"),
			},
			deleteAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				rewritePermissionForUpdate("CAN_VIEW", "CAN_READ"),
			},
			readAclCustomizers: []aclReadCustomizer{
				removeAdminPermissionsCustomizer,
				rewritePermissionForRead("CAN_READ", "CAN_VIEW"),
			},
		},
		{
			field:             "dashboard_id",
			objectType:        "dashboard",
			requestObjectType: "dashboards",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				"CAN_READ":   {isManagementPermission: false},
			},
			readAclCustomizers: []aclReadCustomizer{
				removeAdminPermissionsCustomizer,
				func(ctx aclReadCustomizerContext, objectAcls *iam.ObjectPermissions) (*iam.ObjectPermissions, error) {
					if strings.HasPrefix(objectAcls.ObjectId, "/dashboards/") {
						// workaround for inconsistent API response returning object ID of file in the workspace
						objectAcls.ObjectId = ctx.getId()
					}
					return objectAcls, nil
				},
			},
		},
		{
			field:             "experiment_id",
			objectType:        "mlflowExperiment",
			requestObjectType: "experiments",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			readAclCustomizers: []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "registered_model_id",
			objectType:        "registered-model",
			requestObjectType: "registered-models",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":                       {isManagementPermission: false},
				"CAN_EDIT":                       {isManagementPermission: false},
				"CAN_MANAGE_STAGING_VERSIONS":    {isManagementPermission: false},
				"CAN_MANAGE_PRODUCTION_VERSIONS": {isManagementPermission: false},
				"CAN_MANAGE":                     {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				addAdminAclCustomizer(func(objectId string) bool {
					return objectId == "/registered-models/root"
				}),
			},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			readAclCustomizers:   []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
		{
			field:             "serving_endpoint_id",
			objectType:        "serving-endpoint",
			requestObjectType: "serving-endpoints",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_VIEW":   {isManagementPermission: false},
				"CAN_QUERY":  {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			readAclCustomizers:   []aclReadCustomizer{removeAdminPermissionsCustomizer},
		},
	}
}
