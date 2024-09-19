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

	// Behavior Options and Customizations

	// The alternative name of the "path" attribute for this resource. E.g. "workspace_file_path" for a file.
	// If not set, default is "<object_type>_path".
	pathVariant string
	// Allow users to set permissions for the admin group. This is only permitted for tokens and passwords
	// permissions today.
	allowAdminGroup bool
	// Returns whether the object requires explicit admin permissions. This is only true for tokens and passwords
	explicitAdminPermissionCheck func(objectId string) bool
	// Whether the object requires explicit manage permissions for the calling user if not set.
	// As described in https://github.com/databricks/terraform-provider-databricks/issues/1504,
	// certain object types require that we explicitly grant the calling user CAN_MANAGE
	// permissions when POSTing permissions changes through the REST API, to avoid accidentally
	// revoking the calling user's ability to manage the current object.
	shouldExplicitlyGrantCallingUserManagePermissions bool
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

// hasOwnerPermissionLevel returns whether the owner permission level is allowed for this resource type.
func (p resourcePermissions) hasOwnerPermissionLevel() bool {
	_, ok := p.allowedPermissionLevels["IS_OWNER"]
	return ok
}

// requiresExplicitAdminPermissions returns whether the object requires explicit admin permissions.
func (p resourcePermissions) requiresExplicitAdminPermissions(id string) bool {
	if p.explicitAdminPermissionCheck != nil {
		return p.explicitAdminPermissionCheck(id)
	}
	return false
}

type resourceStatus struct {
	exists  bool
	creator string
}

// getObjectCreator returns the creator of the object. If the object creator cannot be determined for this
// resource type, an empty string is returned.
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
		// Check if the user is trying to set permissions for the admin group
		if change.GroupName == "admins" && !p.allowAdminGroup {
			return fmt.Errorf("it is not possible to modify admin permissions for %s resources", p.objectType)
		}
		// Check that the user is not preventing themselves from managing the object
		if change.UserName == currentUsername && !p.allowedPermissionLevels[change.PermissionLevel].isManagementPermission {
			allowedLevels := p.getAllowedPermissionLevels(false)
			return fmt.Errorf("cannot remove management permissions for the current user for %s, allowed levels: %s", p.objectType, strings.Join(allowedLevels, ", "))
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

func (p resourcePermissions) prepareForUpdate(objectID string, objectACL []AccessControlChangeApiRequest, currentUser string) ([]AccessControlChangeApiRequest, error) {
	// this logic was moved from CustomizeDiff because of undeterministic auth behavior
	// in the corner-case scenarios.
	// see https://github.com/databricks/terraform-provider-databricks/issues/2052
	err := p.validate(objectACL, currentUser)
	if err != nil {
		return nil, err
	}
	if p.requiresExplicitAdminPermissions(objectID) {
		// Prevent "Cannot change permissions for group 'admins' to None."
		objectACL = append(objectACL, AccessControlChangeApiRequest{
			GroupName:       "admins",
			PermissionLevel: "CAN_MANAGE",
		})
	}
	cachedCurrentUser := func() (string, error) { return currentUser, nil }
	return p.prepareCommon(objectACL, cachedCurrentUser)
}

func (p resourcePermissions) prepareForDelete(objectACL ObjectAclApiResponse, getCurrentUser func() (string, error)) ([]AccessControlChangeApiRequest, error) {
	accl := make([]AccessControlChangeApiRequest, 0, len(objectACL.AccessControlList))
	for _, acl := range objectACL.AccessControlList {
		// When deleting permissions for a resource with explicit admin permissions, delete should remove
		// admin permissions as well. Otherwise, admin permissions should be left as-is.
		if acl.GroupName == "admins" && !p.allowAdminGroup {
			if change, direct := acl.toAccessControlChange(); direct {
				// keep everything direct for admin group
				accl = append(accl, change)
			}
		}
	}
	return p.prepareCommon(accl, getCurrentUser)
}

func (p resourcePermissions) prepareCommon(changes []AccessControlChangeApiRequest, getCurrentUser func() (string, error)) ([]AccessControlChangeApiRequest, error) {
	if p.shouldExplicitlyGrantCallingUserManagePermissions {
		currentUser, err := getCurrentUser()
		if err != nil {
			return nil, err
		}
		// The validate() method called in Update() ensures that the current user's permissions are either CAN_MANAGE
		// or IS_OWNER if they are specified. If the current user is not specified in the access control list, we add
		// them with CAN_MANAGE permissions.
		found := false
		for _, acl := range changes {
			if acl.UserName == currentUser || acl.ServicePrincipalName == currentUser {
				found = true
				break
			}
		}
		if !found {
			changes = append(changes, AccessControlChangeApiRequest{
				UserName:        currentUser,
				PermissionLevel: "CAN_MANAGE",
			})
		}
	}
	return changes, nil
}

func (p resourcePermissions) addOwnerPermissionIfNeeded(objectACL []AccessControlChangeApiRequest, ownerOpt string) []AccessControlChangeApiRequest {
	if !p.hasOwnerPermissionLevel() {
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

func getResourcePermissions(objectId string) (resourcePermissions, error) {
	objectParts := strings.Split(objectId, "/")
	objectType := strings.Join(objectParts[1:len(objectParts)-1], "/")
	for _, p := range allResourcePermissions() {
		if p.resourceType == objectType {
			return p, nil
		}
	}
	return resourcePermissions{}, fmt.Errorf("no permissions resource found for object type %s", objectType)
}

func getResourcePermissionsFromState(d interface{ GetOk(string) (any, bool) }) (resourcePermissions, string, error) {
	allPermissions := allResourcePermissions()
	for _, mapping := range allPermissions {
		if v, ok := d.GetOk(mapping.field); ok {
			return mapping, v.(string), nil
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
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
			explicitAdminPermissionCheck: func(objectId string) bool {
				return objectId == "/directories/0"
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
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":    {isManagementPermission: true},
				"CAN_MANAGE": {isManagementPermission: true},
			},
			allowAdminGroup: true,
			explicitAdminPermissionCheck: func(objectId string) bool {
				return objectId == "/authorization/tokens"
			},
		},
		{
			field:        "authorization",
			objectType:   "passwords",
			resourceType: "authorization",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {isManagementPermission: true},
			},
			allowAdminGroup: true,
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
			explicitAdminPermissionCheck: func(objectId string) bool {
				return objectId == "/registered-models/root"
			},
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
			shouldExplicitlyGrantCallingUserManagePermissions: true,
		},
	}
}
