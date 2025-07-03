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
	"github.com/databricks/terraform-provider-databricks/permissions/entity"
	"github.com/databricks/terraform-provider-databricks/permissions/read"
	"github.com/databricks/terraform-provider-databricks/permissions/update"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	// If true, the provider will allow the user to configure the "admins" group for this resource type. Otherwise,
	// validation will fail if the user tries to configure the "admins" group, and admin configurations in API
	// responses will be ignored. This should only be set to true for the "authorization = passwords" resource.
	allowConfiguringAdmins bool
	// Customizers when handling permission resource creation and update.
	//
	// Most resources that have a CAN_MANAGE permission level should add update.AddCurrentUserAsManage to this list
	// to ensure that the user applying the template always has management permissions on the underlying resource.
	updateAclCustomizers []update.ACLCustomizer
	// Customizers when handling permission resource deletion.
	//
	// Most resources that have a CAN_MANAGE permission level should add update.AddCurrentUserAsManage to this list
	// to ensure that the user applying the template always has management permissions on the underlying resource.
	deleteAclCustomizers []update.ACLCustomizer
	// Customizers when handling permission resource read.
	//
	// Resources for which admins inherit permissions should add removeAdminPermissionsCustomizer to this list. This
	// prevents the admin group from being included in the permissions when reading the state.
	readAclCustomizers []read.ACLCustomizer

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
func (p resourcePermissions) validate(ctx context.Context, entity entity.PermissionsEntity, currentUsername string) error {
	for _, change := range entity.AccessControlList {
		// Prevent users from setting permissions for admins.
		if change.GroupName == "admins" && !p.allowConfiguringAdmins {
			return fmt.Errorf("it is not possible to modify admin permissions for %s resources", p.objectType)
		}
		// Check that the user is preventing themselves from managing the object
		level := p.allowedPermissionLevels[string(change.PermissionLevel)]
		if (change.UserName == currentUsername || change.ServicePrincipalName == currentUsername) && !level.isManagementPermission {
			allowedLevelsForCurrentUser := p.getAllowedPermissionLevels(false)
			return fmt.Errorf("cannot remove management permissions for the current user for %s, allowed levels: %s", p.objectType, strings.Join(allowedLevelsForCurrentUser, ", "))
		}

		if level.deprecated != "" {
			tflog.Debug(ctx, fmt.Sprintf("the permission level %s for %s is deprecated: %s", change.PermissionLevel, p.objectType, level.deprecated))
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
func (p resourcePermissions) prepareForUpdate(objectID string, e entity.PermissionsEntity, currentUser string) (entity.PermissionsEntity, error) {
	cachedCurrentUser := func() (string, error) { return currentUser, nil }
	ctx := update.ACLCustomizerContext{
		GetCurrentUser: cachedCurrentUser,
		GetId:          func() string { return objectID },
	}
	var err error
	for _, customizer := range p.updateAclCustomizers {
		e.AccessControlList, err = customizer(ctx, e.AccessControlList)
		if err != nil {
			return entity.PermissionsEntity{}, err
		}
	}
	return e, nil
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
	ctx := update.ACLCustomizerContext{
		GetCurrentUser: getCurrentUser,
		GetId:          func() string { return objectACL.ObjectId },
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
// added by update.AddCurrentUserAsManage.
//
// Read customizers are able to access the current state of the object in order to customize the response accordingly.
// For example, the SQL API previously used CAN_VIEW for read-only permission, but the GA API uses CAN_READ. Users may
// have CAN_VIEW in their resource configuration, so the read customizer will rewrite the response from CAN_READ to
// CAN_VIEW to match the user's configuration.
func (p resourcePermissions) prepareResponse(objectID string, objectACL *iam.ObjectPermissions, existing entity.PermissionsEntity, me string) (entity.PermissionsEntity, error) {
	ctx := read.ACLCustomizerContext{
		GetId:                        func() string { return objectID },
		GetExistingPermissionsEntity: func() entity.PermissionsEntity { return existing },
	}
	acl := *objectACL
	for _, customizer := range p.readAclCustomizers {
		acl = customizer(ctx, acl)
	}
	if acl.ObjectType != p.objectType {
		return entity.PermissionsEntity{}, fmt.Errorf("expected object type %s, got %s", p.objectType, objectACL.ObjectType)
	}
	entity := entity.PermissionsEntity{}
	for _, accessControl := range acl.AccessControlList {
		// If the user doesn't include an access_control block for themselves, do not include it in the state.
		// On create/update, the provider will automatically include the current user in the access_control block
		// for appropriate resources. Otherwise, it must be included in state to prevent configuration drift.
		if me == accessControl.UserName || me == accessControl.ServicePrincipalName {
			if !existing.ContainsUserOrServicePrincipal(me) {
				continue
			}
		}
		// Skip admin permissions for resources where users are not allowed to explicitly configure them.
		if accessControl.GroupName == "admins" && !p.allowConfiguringAdmins {
			continue
		}
		for _, permission := range accessControl.AllPermissions {
			// Inherited permissions can be ignored, as they are not set by the user.
			if permission.Inherited {
				continue
			}
			entity.AccessControlList = append(entity.AccessControlList, iam.AccessControlRequest{
				GroupName:            accessControl.GroupName,
				UserName:             accessControl.UserName,
				ServicePrincipalName: accessControl.ServicePrincipalName,
				PermissionLevel:      permission.PermissionLevel,
			})
		}
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

	// If non-empty, the permission level is deprecated. The string is a message to display to the user when
	// this permission level is used.
	deprecated string
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
	rewriteCanViewToCanRead := update.RewritePermissions(map[iam.PermissionLevel]iam.PermissionLevel{
		iam.PermissionLevelCanView: iam.PermissionLevelCanRead,
	})
	rewriteCanReadToCanView := read.RewritePermissions(map[iam.PermissionLevel]iam.PermissionLevel{
		iam.PermissionLevelCanRead: iam.PermissionLevelCanView,
	})
	return []resourcePermissions{
		{
			field:             "cluster_policy_id",
			objectType:        "cluster-policy",
			requestObjectType: "cluster-policies",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {isManagementPermission: true},
			},
		},
		{
			field:             "instance_pool_id",
			objectType:        "instance-pool",
			requestObjectType: "instance-pools",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_ATTACH_TO": {isManagementPermission: false},
				"CAN_MANAGE":    {isManagementPermission: true},
			},
			updateAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
			deleteAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
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
			updateAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
			deleteAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
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
			idRetriever: PATH,
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
			updateAclCustomizers: []update.ACLCustomizer{
				update.If(update.ObjectIdMatches("/directories/0"), update.AddAdmin),
			},
			deleteAclCustomizers: []update.ACLCustomizer{
				update.If(update.ObjectIdMatches("/directories/0"), update.AddAdmin),
			},
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
			idRetriever: PATH,
			updateAclCustomizers: []update.ACLCustomizer{
				update.If(update.ObjectIdMatches("/directories/0"), update.AddAdmin),
			},
			deleteAclCustomizers: []update.ACLCustomizer{
				update.If(update.ObjectIdMatches("/directories/0"), update.AddAdmin),
			},
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
			pathVariant: "workspace_file_path",
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
			idRetriever: PATH,
			pathVariant: "workspace_file_path",
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
			idRetriever: PATH,
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
			updateAclCustomizers: []update.ACLCustomizer{
				update.If(update.ObjectIdMatches("/authorization/tokens"), update.AddAdmin),
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
			allowConfiguringAdmins: true,
		},
		{
			field:             "sql_endpoint_id",
			objectType:        "warehouses",
			requestObjectType: "sql/warehouses",
			// ISSUE-4143: some older warehouse permissions have an ID that starts with "/warehouses" instead of "/sql/warehouses"
			// Because no idRetriever is defined, any warehouse permissions resources will be created with the "/sql/warehouses" prefix.
			idMatcher: func(id string) bool {
				return strings.HasPrefix(id, "/sql/warehouses/") || strings.HasPrefix(id, "/warehouses/")
			},
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":     {isManagementPermission: false},
				"CAN_MANAGE":  {isManagementPermission: true},
				"CAN_MONITOR": {isManagementPermission: false},
				"CAN_VIEW":    {isManagementPermission: false},
				"IS_OWNER":    {isManagementPermission: true},
			},
			updateAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
			deleteAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
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
				"CAN_READ":   {isManagementPermission: false},
				// This was part of the original SQL permissions API but was replaced by CAN_READ in the GA API.
				"CAN_VIEW": {
					isManagementPermission: false,
					deprecated:             "use CAN_READ instead",
				},
			},
			idMatcher: func(id string) bool {
				return strings.HasPrefix(id, "/dbsql-dashboards/") || strings.HasPrefix(id, "/sql/dashboards/")
			},
			updateAclCustomizers: []update.ACLCustomizer{
				update.AddCurrentUserAsManage,
				rewriteCanViewToCanRead,
			},
			deleteAclCustomizers: []update.ACLCustomizer{
				update.AddCurrentUserAsManage,
				rewriteCanViewToCanRead,
			},
			readAclCustomizers: []read.ACLCustomizer{
				rewriteCanReadToCanView,
				func(ctx read.ACLCustomizerContext, objectAcls iam.ObjectPermissions) iam.ObjectPermissions {
					// The object type in the new API is "dbsql-dashboard", but for compatibility this should
					// be "dashboard" in the state.
					objectAcls.ObjectType = "dashboard"
					return objectAcls
				},
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
				"CAN_READ":   {isManagementPermission: false},
				// This was part of the original SQL permissions API but was replaced by CAN_READ in the GA API.
				// It should eventually be deprecated.
				"CAN_VIEW": {
					isManagementPermission: false,
					deprecated:             "use CAN_READ instead",
				},
			},
			updateAclCustomizers: []update.ACLCustomizer{
				update.AddCurrentUserAsManage,
				rewriteCanViewToCanRead,
			},
			deleteAclCustomizers: []update.ACLCustomizer{
				update.AddCurrentUserAsManage,
				rewriteCanViewToCanRead,
			},
			readAclCustomizers: []read.ACLCustomizer{
				rewriteCanReadToCanView,
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
				"CAN_READ":   {isManagementPermission: false},
				// This was part of the original SQL permissions API but was replaced by CAN_READ in the GA API.
				// It should eventually be deprecated.
				"CAN_VIEW": {
					isManagementPermission: false,
					deprecated:             "use CAN_READ instead",
				},
			},
			updateAclCustomizers: []update.ACLCustomizer{
				update.AddCurrentUserAsManage,
				rewriteCanViewToCanRead,
			},
			deleteAclCustomizers: []update.ACLCustomizer{
				update.AddCurrentUserAsManage,
				rewriteCanViewToCanRead,
			},
			readAclCustomizers: []read.ACLCustomizer{
				rewriteCanReadToCanView,
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
			readAclCustomizers: []read.ACLCustomizer{
				func(ctx read.ACLCustomizerContext, objectAcls iam.ObjectPermissions) iam.ObjectPermissions {
					if strings.HasPrefix(objectAcls.ObjectId, "/dashboards/") {
						// workaround for inconsistent API response returning object ID of file in the workspace
						objectAcls.ObjectId = ctx.GetId()
					}
					return objectAcls
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
			updateAclCustomizers: []update.ACLCustomizer{
				update.AddCurrentUserAsManage,
				update.If(update.ObjectIdMatches("/registered-models/root"), update.AddAdmin),
			},
			deleteAclCustomizers: []update.ACLCustomizer{
				update.If(update.Not(update.ObjectIdMatches("/registered-models/root")), update.AddCurrentUserAsManage),
			},
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
			updateAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
			deleteAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
		},
		{
			field:             "vector_search_endpoint_id",
			objectType:        "vector-search-endpoints",
			requestObjectType: "vector-search-endpoints",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			updateAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
			deleteAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
		},
		{
			field:             "app_name",
			objectType:        "apps",
			requestObjectType: "apps",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			updateAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
			deleteAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
		},
		{
			field:             "alert_v2_id",
			objectType:        "alert_v2",
			requestObjectType: "alert_v2",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			updateAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
			deleteAclCustomizers: []update.ACLCustomizer{update.AddCurrentUserAsManage},
		},
	}
}
