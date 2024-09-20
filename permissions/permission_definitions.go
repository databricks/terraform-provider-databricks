package permissions

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

// resourcePermissions captures all the information needed to manage permissions for a given object type.
type resourcePermissions struct {
	// Mandatory Fields

	// The attribute name that users configure with the ID of the object to manage
	// e.g. "cluster_id" for a cluster
	field string
	// The object type in the API response for this resource type, e.g. "cluster" for a cluster.
	objectType string
	// The name of the object in the ID of the TF resource, e.g. "clusters" for a cluster,
	// where the ID would be /clusters/<cluster-id>. This should also match the prefix of the
	// object ID in the API response, unless idMatcher is set.
	resourceType string
	// The allowed permission levels for this object type.
	allowedPermissionLevels map[string]permissionLevelOptions

	// ID Remapping Options

	// Returns the object ID for the given user-specified ID. This is necessary because permissions for
	// some objects are done by path, whereas others are by ID. Those by path need to be converted to the
	// internal object ID before being stored in the state. If not specified, the default ID is "/<resource_type>/<id>".
	idRetriever func(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error)
	// A custom matcher to check whether an object ID in the API response matches this resource type.
	// This is necessary because some objects have a different ID in the API response than the ID in the
	// Terraform state. If unset, the default is to check whether the object ID starts with "/<resource_type>".
	idMatcher func(objectID string) bool
	// A custom matcher to check whether a given ID matches this resource type.
	// Most resources can be determined by looking at the attribute name used to configure the permission, but
	// tokens & passwords are special cases where the resource type is determined by the value of this attribute.
	stateMatcher func(id string) bool

	// Behavior Options and Customizations

	// The alternative name of the "path" attribute for this resource. E.g. "workspace_file_path" for a file.
	// If not set, default is "<object_type>_path".
	pathVariant string
	// Customizers when handling create & update requests
	updateAclCustomizers []aclUpdateCustomizer
	// Customizers when handling delete requests
	deleteAclCustomizers []aclUpdateCustomizer
	// Customizers when handling read requests
	readAclCustomizers []aclReadCustomizer

	// Returns the creator of the object. Used when deleting databricks_permissions resources, when the
	// creator of the object is restored as the owner.
	fetchObjectCreator func(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error)

	// Request Options

	// Returns the path for the object ID.
	// If not set, the default is to use "/permissions/<object_id>"
	makeRequestPath func(objectId string) string
	// If true, use POST instead of PUT for permissions changes.
	// By default, PUT is used for permissions changes.
	usePost bool
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

// getRequestPath returns the request URI path for the object ID.
func (p resourcePermissions) getRequestPath(objectID string) string {
	if p.makeRequestPath != nil {
		return p.makeRequestPath(objectID)
	}
	return "/permissions" + objectID
}

// isTypeOf returns true if the object ID is the ID of a permission for this resourcePermission's resource type.
func (p resourcePermissions) isTypeOf(objectID string) bool {
	if p.idMatcher != nil {
		return p.idMatcher(objectID)
	}
	if objectID != "" && objectID[0] == '/' {
		return strings.HasPrefix(objectID[1:], p.resourceType)
	}
	return false
}

// validate checks that the user is not trying to set permissions for the admin group or remove their own management permissions.
func (p resourcePermissions) validate(changes []AccessControlChangeApiRequest, currentUsername string) error {
	for _, change := range changes {
		// Check if the user is trying to set permissions for the admin group outside of the tokens/passwords permissions.
		if change.GroupName == "admins" && p.field != "authorization" {
			return fmt.Errorf("it is not possible to modify admin permissions for %s resources", p.objectType)
		}
		// Check that the user is preventing themselves from managing the object
		if (change.UserName == currentUsername || change.ServicePrincipalName == currentUsername) && !p.allowedPermissionLevels[change.PermissionLevel].isManagementPermission {
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
	return fmt.Sprintf("/%s/%s", p.resourceType, id), nil
}

// prepareForUpdate prepares the access control list for an update request by calling all update customizers.
func (p resourcePermissions) prepareForUpdate(objectID string, objectACL []AccessControlChangeApiRequest, currentUser string) ([]AccessControlChangeApiRequest, error) {
	cachedCurrentUser := func() (string, error) { return currentUser, nil }
	ctx := aclUpdateCustomizerContext{
		getCurrentUser: cachedCurrentUser,
		getId:          func() string { return objectID },
	}
	var err error
	for _, customizer := range p.updateAclCustomizers {
		objectACL, err = customizer(ctx, objectACL)
		if err != nil {
			return nil, err
		}
	}
	return objectACL, nil
}

// prepareForDelete prepares the access control list for a delete request by calling all delete customizers.
func (p resourcePermissions) prepareForDelete(objectACL ObjectAclApiResponse, getCurrentUser func() (string, error)) ([]AccessControlChangeApiRequest, error) {
	accl := make([]AccessControlChangeApiRequest, 0, len(objectACL.AccessControlList))
	for _, acl := range objectACL.AccessControlList {
		// When deleting permissions for a resource with explicit admin permissions, delete should remove
		// admin permissions as well. Otherwise, admin permissions should be left as-is.
		if acl.GroupName == "admins" {
			if change, direct := acl.toAccessControlChange(); direct {
				// keep everything direct for admin group
				accl = append(accl, change)
			}
		}
	}
	ctx := aclUpdateCustomizerContext{
		getCurrentUser: getCurrentUser,
		getId:          func() string { return objectACL.ObjectID },
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
func (p resourcePermissions) prepareResponse(objectID string, objectACL ObjectAclApiResponse) (ObjectAclApiResponse, error) {
	ctx := aclReadCustomizerContext{
		getId: func() string { return objectID },
	}
	for _, customizer := range p.readAclCustomizers {
		var err error
		objectACL, err = customizer(ctx, objectACL)
		if err != nil {
			return ObjectAclApiResponse{}, err
		}
	}
	return objectACL, nil
}

// addOwnerPermissionIfNeeded adds the owner permission to the object ACL if the owner permission is allowed and not already set.
func (p resourcePermissions) addOwnerPermissionIfNeeded(objectACL []AccessControlChangeApiRequest, ownerOpt string) []AccessControlChangeApiRequest {
	_, ok := p.allowedPermissionLevels["IS_OWNER"]
	if !ok {
		return objectACL
	}

	for _, acl := range objectACL {
		if acl.PermissionLevel == "IS_OWNER" {
			return objectACL
		}
	}

	return append(objectACL, AccessControlChangeApiRequest{
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

// getResourcePermissions returns the resourcePermissions for the given object ID.
// This ID must be the ID of the object in the Terraform state.
func getResourcePermissions(objectId string) (resourcePermissions, error) {
	objectParts := strings.Split(objectId, "/")
	objectType := strings.Join(objectParts[1:len(objectParts)-1], "/")
	for _, p := range allResourcePermissions() {
		if p.resourceType == objectType {
			id := objectParts[len(objectParts)-1]
			if p.stateMatcher != nil && !p.stateMatcher(id) {
				continue
			}
			return p, nil
		}
	}
	return resourcePermissions{}, fmt.Errorf("no permissions resource found for object type %s", objectType)
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
func getResourcePermissionsForObjectAcl(objectACL ObjectAclApiResponse) (resourcePermissions, string, error) {
	for _, mapping := range allResourcePermissions() {
		if mapping.objectType == objectACL.ObjectType && mapping.isTypeOf(objectACL.ObjectID) {
			return mapping, objectACL.ObjectID, nil
		}
	}
	return resourcePermissions{}, "", fmt.Errorf("unknown object type %s", objectACL.ObjectType)
}

// allResourcePermissions is the list of all resource types that can be managed by the databricks_permissions resource.
func allResourcePermissions() []resourcePermissions {
	PATH := func(ctx context.Context, w *databricks.WorkspaceClient, path string) (string, error) {
		info, err := w.Workspace.GetStatusByPath(ctx, path)
		if err != nil {
			return "", fmt.Errorf("cannot load path %s: %s", path, err)
		}
		return strconv.FormatInt(info.ObjectId, 10), nil
	}
	SQL_REQUEST_PATH := func(objectId string) string {
		return "/preview/sql/permissions" + objectId[4:]
	}
	return []resourcePermissions{
		{
			field:        "cluster_policy_id",
			objectType:   "cluster-policy",
			resourceType: "cluster-policies",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {isManagementPermission: true},
			},
		},
		{
			field:        "instance_pool_id",
			objectType:   "instance-pool",
			resourceType: "instance-pools",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_ATTACH_TO": {isManagementPermission: false},
				"CAN_MANAGE":    {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
		},
		{
			field:        "cluster_id",
			objectType:   "cluster",
			resourceType: "clusters",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_ATTACH_TO": {isManagementPermission: false},
				"CAN_RESTART":   {isManagementPermission: false},
				"CAN_MANAGE":    {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
		},
		{
			field:        "pipeline_id",
			objectType:   "pipelines",
			resourceType: "pipelines",
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
			field:        "job_id",
			objectType:   "job",
			resourceType: "jobs",
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
			field:        "notebook_id",
			objectType:   "notebook",
			resourceType: "notebooks",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
		},
		{
			field:        "notebook_path",
			objectType:   "notebook",
			resourceType: "notebooks",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			idRetriever: PATH,
		},
		{
			field:        "directory_id",
			objectType:   "directory",
			resourceType: "directories",
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
		},
		{
			field:        "directory_path",
			objectType:   "directory",
			resourceType: "directories",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			idRetriever: PATH,
		},
		{
			field:        "workspace_file_id",
			objectType:   "file",
			resourceType: "files",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			pathVariant: "workspace_file_path",
		},
		{
			field:        "workspace_file_path",
			objectType:   "file",
			resourceType: "files",
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
			field:        "repo_id",
			objectType:   "repo",
			resourceType: "repos",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
		},
		{
			field:        "repo_path",
			objectType:   "repo",
			resourceType: "repos",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			idRetriever: PATH,
		},
		{
			field:        "authorization",
			objectType:   "tokens",
			resourceType: "authorization",
			stateMatcher: func(id string) bool {
				return id == "tokens"
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
			field:        "authorization",
			objectType:   "passwords",
			resourceType: "authorization",
			stateMatcher: func(id string) bool {
				return id == "passwords"
			},
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {isManagementPermission: true},
			},
		},
		{
			field:        "sql_endpoint_id",
			objectType:   "warehouses",
			resourceType: "sql/warehouses",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":     {isManagementPermission: false},
				"CAN_MANAGE":  {isManagementPermission: true},
				"CAN_MONITOR": {isManagementPermission: false},
				"IS_OWNER":    {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			fetchObjectCreator: func(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error) {
				warehouse, err := w.Warehouses.GetById(ctx, strings.ReplaceAll(objectID, "/sql/warehouses/", ""))
				if err != nil {
					return "", common.IgnoreNotFoundError(err)
				}
				return warehouse.CreatorName, nil
			},
		},
		{
			field:        "sql_dashboard_id",
			objectType:   "dashboard",
			resourceType: "sql/dashboards",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				"CAN_VIEW":   {isManagementPermission: false},
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				copyServicePrincipalToUserCustomizer,
			},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			readAclCustomizers: []aclReadCustomizer{
				copyUserToServicePrincipalCustomizer,
			},
			makeRequestPath: SQL_REQUEST_PATH,
			usePost:         true,
			idMatcher: func(objectID string) bool {
				return strings.HasPrefix(objectID, "dashboards/")
			},
		},
		{
			field:        "sql_alert_id",
			objectType:   "alert",
			resourceType: "sql/alerts",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				"CAN_VIEW":   {isManagementPermission: false},
			},
			idMatcher: func(objectID string) bool {
				return strings.HasPrefix(objectID, "alerts/")
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				copyServicePrincipalToUserCustomizer,
			},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			readAclCustomizers: []aclReadCustomizer{
				copyUserToServicePrincipalCustomizer,
			},
			makeRequestPath: SQL_REQUEST_PATH,
			usePost:         true,
		},
		{
			field:        "sql_query_id",
			objectType:   "query",
			resourceType: "sql/queries",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				"CAN_VIEW":   {isManagementPermission: false},
			},
			updateAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
				copyServicePrincipalToUserCustomizer,
			},
			deleteAclCustomizers: []aclUpdateCustomizer{
				addCurrentUserAsManageCustomizer,
			},
			readAclCustomizers: []aclReadCustomizer{
				copyUserToServicePrincipalCustomizer,
			},
			idMatcher: func(objectID string) bool {
				return strings.HasPrefix(objectID, "queries/")
			},
			makeRequestPath: SQL_REQUEST_PATH,
			usePost:         true,
		},
		{
			field:        "dashboard_id",
			objectType:   "dashboard",
			resourceType: "dashboards",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_RUN":    {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
				"CAN_READ":   {isManagementPermission: false},
			},
			readAclCustomizers: []aclReadCustomizer{
				func(ctx aclReadCustomizerContext, objectAcls ObjectAclApiResponse) (ObjectAclApiResponse, error) {
					if strings.HasPrefix(objectAcls.ObjectID, "/dashboards/") {
						// workaround for inconsistent API response returning object ID of file in the workspace
						objectAcls.ObjectID = ctx.getId()
					}
					return objectAcls, nil
				},
			},
		},
		{
			field:        "experiment_id",
			objectType:   "mlflowExperiment",
			resourceType: "experiments",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {isManagementPermission: false},
				"CAN_EDIT":   {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
		},
		{
			field:        "registered_model_id",
			objectType:   "registered-model",
			resourceType: "registered-models",
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
		},
		{
			field:        "serving_endpoint_id",
			objectType:   "serving-endpoint",
			resourceType: "serving-endpoints",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_VIEW":   {isManagementPermission: false},
				"CAN_QUERY":  {isManagementPermission: false},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			updateAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
			deleteAclCustomizers: []aclUpdateCustomizer{addCurrentUserAsManageCustomizer},
		},
	}
}
