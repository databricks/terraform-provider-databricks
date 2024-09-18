package permissions

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ObjectACL is a structure to generically describe access control
type ObjectACL struct {
	ObjectID          string          `json:"object_id,omitempty"`
	ObjectType        string          `json:"object_type,omitempty"`
	AccessControlList []AccessControl `json:"access_control_list"`
}

// AccessControl is a structure to describe user/group permissions
type AccessControl struct {
	UserName             string       `json:"user_name,omitempty"`
	GroupName            string       `json:"group_name,omitempty"`
	ServicePrincipalName string       `json:"service_principal_name,omitempty"`
	AllPermissions       []Permission `json:"all_permissions,omitempty"`

	// SQLA entities don't use the `all_permissions` nesting, but rather a simple
	// top level string with the permission level when retrieving permissions.
	PermissionLevel string `json:"permission_level,omitempty"`
}

func (ac AccessControl) toAccessControlChange() (AccessControlChange, bool) {
	for _, permission := range ac.AllPermissions {
		if permission.Inherited {
			continue
		}
		return AccessControlChange{
			PermissionLevel:      permission.PermissionLevel,
			UserName:             ac.UserName,
			GroupName:            ac.GroupName,
			ServicePrincipalName: ac.ServicePrincipalName,
		}, true
	}
	if ac.PermissionLevel != "" {
		return AccessControlChange{
			PermissionLevel:      ac.PermissionLevel,
			UserName:             ac.UserName,
			GroupName:            ac.GroupName,
			ServicePrincipalName: ac.ServicePrincipalName,
		}, true
	}
	return AccessControlChange{}, false
}

func (ac AccessControl) String() string {
	return fmt.Sprintf("%s%s%s%v", ac.GroupName, ac.UserName, ac.ServicePrincipalName, ac.AllPermissions)
}

// Permission is a structure to describe permission level
type Permission struct {
	PermissionLevel     string   `json:"permission_level"`
	Inherited           bool     `json:"inherited,omitempty"`
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
}

func (p Permission) String() string {
	if len(p.InheritedFromObject) > 0 {
		return fmt.Sprintf("%s (from %s)", p.PermissionLevel, p.InheritedFromObject)
	}
	return p.PermissionLevel
}

// AccessControlChangeList is wrapper around ACL changes for REST API
type AccessControlChangeList struct {
	AccessControlList []AccessControlChange `json:"access_control_list"`
}

// AccessControlChange is API wrapper for changing permissions
type AccessControlChange struct {
	UserName             string `json:"user_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	PermissionLevel      string `json:"permission_level"`
}

func (acc AccessControlChange) String() string {
	return fmt.Sprintf("%v%v%v %s", acc.UserName, acc.GroupName, acc.ServicePrincipalName,
		acc.PermissionLevel)
}

// NewPermissionsAPI creates PermissionsAPI instance from provider meta
func NewPermissionsAPI(ctx context.Context, m any) PermissionsAPI {
	return PermissionsAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// PermissionsAPI exposes general permission related methods
type PermissionsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func isDbsqlPermissionsWorkaroundNecessary(objectID string) bool {
	return strings.HasPrefix(objectID, "/sql/") && !strings.HasPrefix(objectID, "/sql/warehouses")
}

func urlPathForObjectID(objectID string) string {
	if isDbsqlPermissionsWorkaroundNecessary(objectID) {
		// Permissions for SQLA entities are routed differently from the others.
		return "/preview/sql/permissions" + objectID[4:]
	}
	return "/permissions" + objectID
}

// Helper function for applying permissions changes. Ensures that
// we select the correct HTTP method based on the object type and preserve the calling
// user's ability to manage the specified object when applying permissions changes.
func (a PermissionsAPI) put(objectID string, objectACL AccessControlChangeList) error {
	if isDbsqlPermissionsWorkaroundNecessary(objectID) {
		// SQLA entities use POST for permission updates.
		return a.client.Post(a.context, urlPathForObjectID(objectID), objectACL, nil)
	}
	log.Printf("[DEBUG] PUT %s %v", objectID, objectACL)
	return a.client.Put(a.context, urlPathForObjectID(objectID), objectACL)
}

// safePutWithOwner is a workaround for the limitation where warehouse without owners cannot have IS_OWNER set
func (a PermissionsAPI) safePutWithOwner(objectID string, objectACL AccessControlChangeList, getCurrentUser, getOwner func() (string, error)) error {
	mapping, err := getResourcePermissions(objectID)
	if mapping.shouldExplicitlyGrantCallingUserManagePermissions {
		currentUser, err := getCurrentUser()
		if err != nil {
			return err
		}
		objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
			UserName:        currentUser,
			PermissionLevel: "CAN_MANAGE",
		})
	}
	originalAcl := make([]AccessControlChange, len(objectACL.AccessControlList))
	copy(originalAcl, objectACL.AccessControlList)
	if err != nil {
		return err
	}
	if mapping.hasOwnerPermissionLevel() {
		owners := 0
		for _, acl := range objectACL.AccessControlList {
			if acl.PermissionLevel == "IS_OWNER" {
				owners++
			}
		}
		if owners == 0 {
			// add owner if it's missing, otherwise automated planning might be difficult
			owner, err := getOwner()
			if err != nil {
				return err
			}
			if owner == "" {
				return nil
			}
			objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
				UserName:        owner,
				PermissionLevel: "IS_OWNER",
			})
		}
	}
	err = a.put(objectID, objectACL)
	if err != nil {
		if strings.Contains(err.Error(), "with no existing owner must provide a new owner") {
			objectACL.AccessControlList = originalAcl
			return a.put(objectID, objectACL)
		}
		return err
	}
	return nil
}

func (a PermissionsAPI) getCurrentUser() (string, error) {
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return "", err
	}
	me, err := w.CurrentUser.Me(a.context)
	if err != nil {
		return "", err
	}
	return me.UserName, nil
}

// Update updates object permissions. Technically, it's using method named SetOrDelete, but here we do more
func (a PermissionsAPI) Update(objectID string, objectACL AccessControlChangeList, mapping resourcePermissions) error {
	currentUser, err := a.getCurrentUser()
	if err != nil {
		return err
	}
	// this logic was moved from CustomizeDiff because of undeterministic auth behavior
	// in the corner-case scenarios.
	// see https://github.com/databricks/terraform-provider-databricks/issues/2052
	err = mapping.validate(objectACL.AccessControlList, currentUser)
	if err != nil {
		return err
	}
	if mapping.requiresExplicitAdminPermissions(objectID) {
		// Prevent "Cannot change permissions for group 'admins' to None."
		objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
			GroupName:       "admins",
			PermissionLevel: "CAN_MANAGE",
		})
	}
	getCurrentUser := func() (string, error) { return currentUser, nil }
	return a.safePutWithOwner(objectID, objectACL, getCurrentUser, getCurrentUser)
}

// Delete gracefully removes permissions of non-admin users. After this operation, the object is managed
// by the current user and admin group.
func (a PermissionsAPI) Delete(objectID string, mapping resourcePermissions) error {
	objectACL, err := a.Read(objectID)
	if err != nil {
		return err
	}
	accl := AccessControlChangeList{}
	for _, acl := range objectACL.AccessControlList {
		// When deleting permissions for a resource with explicit admin permissions, delete should remove
		// admin permissions as well. Otherwise, admin permissions should be left as-is.
		if acl.GroupName == "admins" && !mapping.allowAdminGroup {
			if change, direct := acl.toAccessControlChange(); direct {
				// keep everything direct for admin group
				accl.AccessControlList = append(accl.AccessControlList, change)
			}
		}
	}
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return err
	}
	return a.safePutWithOwner(objectID, accl, a.getCurrentUser, func() (string, error) { return mapping.getObjectCreator(a.context, w, objectID) })
}

// Read gets all relevant permissions for the object, including inherited ones
func (a PermissionsAPI) Read(objectID string) (objectACL ObjectACL, err error) {
	err = a.client.Get(a.context, urlPathForObjectID(objectID), nil, &objectACL)
	var apiErr *apierr.APIError
	// https://github.com/databricks/terraform-provider-databricks/issues/1227
	// platform propagates INVALID_STATE error for auto-purged clusters in
	// the permissions api. this adds "a logical fix" also here, not to introduce
	// cross-package dependency on "clusters".
	if errors.As(err, &apiErr) && strings.Contains(apiErr.Message, "Cannot access cluster") && apiErr.StatusCode == 400 {
		apiErr.StatusCode = 404
		err = apiErr
		return
	}
	if strings.HasPrefix(objectID, "/dashboards/") {
		// workaround for inconsistent API response returning object ID of file in the workspace
		objectACL.ObjectID = objectID
	}
	return
}

// resourcePermissions holds mapping
type resourcePermissions struct {
	// The top-level attribute name in the schema that holds the ID of the object
	// e.g. "cluster_id" for a cluster
	field string
	// The value of the computed `object_type` field, set by the provider based on the
	// resource type, e.g. "cluster" for a cluster
	objectType string
	// The name of the object in the ID of the TF resource, e.g. "clusters" for a cluster,
	// where the ID would be /clusters/<cluster-id>
	resourceType string
	// The alternative name of the "path" attribute for this resource. E.g. "workspace_file_path" for a file.
	// If not set, default is "<object_type>_path".
	pathVariant string
	// The allowed permission levels for this object type
	allowedPermissionLevels map[string]permissionLevelOptions
	// Allow users to set permissions for the admin group. This is only permitted for tokens and passwords
	// permissions today
	allowAdminGroup bool
	// Whether the object requires explicit manage permissions for the calling user if not set.
	// As described in https://github.com/databricks/terraform-provider-databricks/issues/1504,
	// certain object types require that we explicitly grant the calling user CAN_MANAGE
	// permissions when POSTing permissions changes through the REST API, to avoid accidentally
	// revoking the calling user's ability to manage the current object.
	shouldExplicitlyGrantCallingUserManagePermissions bool
	// Returns the object ID for the given user-specified ID. This is necessary because permissions for
	// some objects are done by path, whereas others are by ID. Those by path need to be converted to the
	// internal object ID before being stored in the state.
	idRetriever func(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error)
	// Returns the creator of the object.
	fetchObjectCreator func(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error)
	// Returns whether the object requires explicit admin permissions. This is only true for tokens and passwords
	explicitAdminPermissionCheck func(objectId string) bool
}

func (p resourcePermissions) getAllowedPermissionLevels(includeNonSettable bool) []string {
	levels := make([]string, 0, len(p.allowedPermissionLevels))
	for level := range p.allowedPermissionLevels {
		if includeNonSettable || p.allowedPermissionLevels[level].currentUserSettable {
			levels = append(levels, level)
		}
	}
	sort.Strings(levels)
	return levels
}

func (p resourcePermissions) hasOwnerPermissionLevel() bool {
	_, ok := p.allowedPermissionLevels["IS_OWNER"]
	return ok
}

func (p resourcePermissions) requiresExplicitAdminPermissions(id string) bool {
	if p.explicitAdminPermissionCheck != nil {
		return p.explicitAdminPermissionCheck(id)
	}
	return false
}

func (p resourcePermissions) getObjectCreator(ctx context.Context, w *databricks.WorkspaceClient, objectID string) (string, error) {
	if p.fetchObjectCreator != nil {
		return p.fetchObjectCreator(ctx, w, objectID)
	}
	return "", nil
}

func (p resourcePermissions) getPathVariant() string {
	if p.pathVariant != "" {
		return p.pathVariant
	}
	return p.objectType + "_path"
}

func (p resourcePermissions) validate(changes []AccessControlChange, currentUsername string) error {
	for _, change := range changes {
		// Check if the user is trying to set permissions for the admin group
		if change.GroupName == "admins" && !p.allowAdminGroup {
			return fmt.Errorf("it is not possible to modify admin permissions for %s resources", p.objectType)
		}
		// Check that the user is not preventing themselves from managing the object
		if change.UserName == currentUsername {
			return fmt.Errorf("it is not possible to decrease administrative permissions for the current user: %s", currentUsername)
		}
	}
	return nil
}

func (p resourcePermissions) getID(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error) {
	id, err := p.idRetriever(ctx, w, id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/%s/%s", p.resourceType, id), nil
}

// permissionLevelOptions indicates the properties of a permissions level. Today, the only property
// is whether the current user can set the permission level for themselves.
type permissionLevelOptions struct {
	// Whether the current user can set the permission level for themselves.
	currentUserSettable bool
}

func getResourcePermissions(objectId string) (resourcePermissions, error) {
	objectParts := strings.Split(objectId, "/")
	objectType := strings.Join(objectParts[1:len(objectParts)-1], "/")
	for _, p := range permissionsResourceIDFields() {
		if p.resourceType == objectType {
			return p, nil
		}
	}
	return resourcePermissions{}, fmt.Errorf("no permissions resource found for object type %s", objectType)
}

func getResourcePermissionsFromState(d interface{ GetOk(string) (any, bool) }) (resourcePermissions, string, error) {
	allPermissions := permissionsResourceIDFields()
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

func getResourcePermissionsForObjectAcl(objectACL *ObjectACL) (resourcePermissions, string, error) {
	for _, p := range permissionsResourceIDFields() {
		if objectACL.isMatchingMapping(p) {
			return p, objectACL.ObjectID, nil
		}
	}
	return resourcePermissions{}, "", fmt.Errorf("unknown object type %s", objectACL.ObjectType)
}

// PermissionsResourceIDFields shows mapping of id columns to resource types
func permissionsResourceIDFields() []resourcePermissions {
	SIMPLE := func(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error) {
		return id, nil
	}
	PATH := func(ctx context.Context, w *databricks.WorkspaceClient, path string) (string, error) {
		info, err := w.Workspace.GetStatusByPath(ctx, path)
		if err != nil {
			return "", fmt.Errorf("cannot load path %s: %s", path, err)
		}
		return strconv.FormatInt(info.ObjectId, 10), nil
	}
	return []resourcePermissions{
		{
			field:        "cluster_policy_id",
			objectType:   "cluster-policy",
			resourceType: "cluster-policies",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
		},
		{
			field:        "instance_pool_id",
			objectType:   "instance-pool",
			resourceType: "instance-pools",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_ATTACH_TO": {currentUserSettable: false},
				"CAN_MANAGE":    {currentUserSettable: true},
			},
			shouldExplicitlyGrantCallingUserManagePermissions: true,
			idRetriever: SIMPLE,
		},
		{
			field:        "cluster_id",
			objectType:   "cluster",
			resourceType: "clusters",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_ATTACH_TO": {currentUserSettable: false},
				"CAN_RESTART":   {currentUserSettable: false},
				"CAN_MANAGE":    {currentUserSettable: true},
			},
			shouldExplicitlyGrantCallingUserManagePermissions: true,
			idRetriever: SIMPLE,
		},
		{
			field:        "pipeline_id",
			objectType:   "pipelines",
			resourceType: "pipelines",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_VIEW":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
				"IS_OWNER":   {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
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
				"CAN_VIEW":       {currentUserSettable: false},
				"CAN_MANAGE_RUN": {currentUserSettable: false},
				"IS_OWNER":       {currentUserSettable: true},
				"CAN_MANAGE":     {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
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
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
		},
		{
			field:        "notebook_path",
			objectType:   "notebook",
			resourceType: "notebooks",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: PATH,
		},
		{
			field:        "directory_id",
			objectType:   "directory",
			resourceType: "directories",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
			explicitAdminPermissionCheck: func(objectId string) bool {
				return objectId == "/directories/0"
			},
		},
		{
			field:        "directory_path",
			objectType:   "directory",
			resourceType: "directories",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: PATH,
		},
		{
			field:        "workspace_file_id",
			objectType:   "file",
			resourceType: "files",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
			pathVariant: "workspace_file_path",
		},
		{
			field:        "workspace_file_path",
			objectType:   "file",
			resourceType: "files",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: PATH,
			pathVariant: "workspace_file_path",
		},
		{
			field:        "repo_id",
			objectType:   "repo",
			resourceType: "repos",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
		},
		{
			field:        "repo_path",
			objectType:   "repo",
			resourceType: "repos",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: PATH,
		},
		{
			field:        "authorization",
			objectType:   "tokens",
			resourceType: "authorization",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":    {currentUserSettable: true},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			allowAdminGroup: true,
			idRetriever:     SIMPLE,
			explicitAdminPermissionCheck: func(objectId string) bool {
				return objectId == "/authorization/tokens"
			},
		},
		{
			field:        "authorization",
			objectType:   "passwords",
			resourceType: "authorization",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE": {currentUserSettable: true},
			},
			allowAdminGroup: true,
			idRetriever:     SIMPLE,
		},
		{
			field:        "sql_endpoint_id",
			objectType:   "warehouses",
			resourceType: "sql/warehouses",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_USE":     {currentUserSettable: false},
				"CAN_MANAGE":  {currentUserSettable: true},
				"CAN_MONITOR": {currentUserSettable: false},
				"IS_OWNER":    {currentUserSettable: true},
			},
			shouldExplicitlyGrantCallingUserManagePermissions: true,
			idRetriever: SIMPLE,
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
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
				"CAN_VIEW":   {currentUserSettable: false},
			},
			idRetriever: SIMPLE,
			shouldExplicitlyGrantCallingUserManagePermissions: true,
		},
		{
			field:        "sql_alert_id",
			objectType:   "alert",
			resourceType: "sql/alerts",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
				"CAN_VIEW":   {currentUserSettable: false},
			},
			idRetriever: SIMPLE,
			shouldExplicitlyGrantCallingUserManagePermissions: true,
		},
		{
			field:        "sql_query_id",
			objectType:   "query",
			resourceType: "sql/queries",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
				"CAN_VIEW":   {currentUserSettable: false},
			},
			shouldExplicitlyGrantCallingUserManagePermissions: true,
			idRetriever: SIMPLE,
		},
		{
			field:        "dashboard_id",
			objectType:   "dashboard",
			resourceType: "dashboards",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_RUN":    {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
				"CAN_READ":   {currentUserSettable: false},
			},
			idRetriever: SIMPLE,
		},
		{
			field:        "experiment_id",
			objectType:   "mlflowExperiment",
			resourceType: "experiments",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":   {currentUserSettable: false},
				"CAN_EDIT":   {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			idRetriever: SIMPLE,
		},
		{
			field:        "registered_model_id",
			objectType:   "registered-model",
			resourceType: "registered-models",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_READ":                       {currentUserSettable: false},
				"CAN_EDIT":                       {currentUserSettable: false},
				"CAN_MANAGE_STAGING_VERSIONS":    {currentUserSettable: false},
				"CAN_MANAGE_PRODUCTION_VERSIONS": {currentUserSettable: false},
				"CAN_MANAGE":                     {currentUserSettable: true},
			},
			shouldExplicitlyGrantCallingUserManagePermissions: true,
			idRetriever: SIMPLE,
			explicitAdminPermissionCheck: func(objectId string) bool {
				return objectId == "/registered-models/root"
			},
		},
		{
			field:        "serving_endpoint_id",
			objectType:   "serving-endpoint",
			resourceType: "serving-endpoints",
			allowedPermissionLevels: map[string]permissionLevelOptions{
				"CAN_VIEW":   {currentUserSettable: false},
				"CAN_QUERY":  {currentUserSettable: false},
				"CAN_MANAGE": {currentUserSettable: true},
			},
			shouldExplicitlyGrantCallingUserManagePermissions: true,
			idRetriever: SIMPLE,
		},
	}
}

// PermissionsEntity is the one used for resource metadata
type PermissionsEntity struct {
	ObjectType        string                `json:"object_type,omitempty" tf:"computed"`
	AccessControlList []AccessControlChange `json:"access_control" tf:"slice_set"`
}

func (p PermissionsEntity) containsUserOrServicePrincipal(name string) bool {
	for _, ac := range p.AccessControlList {
		if ac.UserName == name || ac.ServicePrincipalName == name {
			return true
		}
	}
	return false
}

func (oa *ObjectACL) isMatchingMapping(mapping resourcePermissions) bool {
	if mapping.objectType != oa.ObjectType {
		return false
	}
	if oa.ObjectID != "" && oa.ObjectID[0] == '/' {
		return strings.HasPrefix(oa.ObjectID[1:], mapping.resourceType)
	}
	if strings.HasPrefix(oa.ObjectID, "dashboards/") || strings.HasPrefix(oa.ObjectID, "alerts/") || strings.HasPrefix(oa.ObjectID, "queries/") {
		idx := strings.Index(oa.ObjectID, "/")
		if idx != -1 {
			return mapping.resourceType == "sql/"+oa.ObjectID[:idx]
		}
	}

	return false
}

func (oa *ObjectACL) ToPermissionsEntity(d *schema.ResourceData, existing PermissionsEntity, me string) (PermissionsEntity, error) {
	entity := PermissionsEntity{}
	mapping, _, err := getResourcePermissionsForObjectAcl(oa)
	if err != nil {
		return entity, err
	}
	for _, accessControl := range oa.AccessControlList {
		if accessControl.GroupName == "admins" && !mapping.allowAdminGroup {
			// admin permission is always returned but can only be explicitly set for certain resources
			// For other resources, admin permissions are not included in the resource state.
			continue
		}
		if me == accessControl.UserName || me == accessControl.ServicePrincipalName {
			// If the user doesn't include an access_control block for themselves, do not include it in the state.
			// On create/update, the provider will automatically include the current user in the access_control block
			// for appropriate resources. Otherwise, it must be included in state to prevent configuration drift.
			if !existing.containsUserOrServicePrincipal(me) {
				continue
			}
		}
		if change, direct := accessControl.toAccessControlChange(); direct {
			entity.AccessControlList = append(entity.AccessControlList, change)
		}
	}
	entity.ObjectType = mapping.objectType
	pathVariant := d.Get(mapping.getPathVariant())
	if pathVariant != nil && pathVariant.(string) != "" {
		// we're not importing and it's a path... it's set, so let's not re-set it
		return entity, nil
	}
	identifier := path.Base(oa.ObjectID)
	return entity, d.Set(mapping.field, identifier)
}

// ResourcePermissions definition
func ResourcePermissions() common.Resource {
	s := common.StructToSchema(PermissionsEntity{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		for _, mapping := range permissionsResourceIDFields() {
			s[mapping.field] = &schema.Schema{
				ForceNew: true,
				Type:     schema.TypeString,
				Optional: true,
			}
			for _, m := range permissionsResourceIDFields() {
				if m.field == mapping.field {
					continue
				}
				s[mapping.field].ConflictsWith = append(s[mapping.field].ConflictsWith, m.field)
			}
		}
		s["access_control"].MinItems = 1
		return s
	})
	return common.Resource{
		Schema: s,
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff) error {
			mapping, _, err := getResourcePermissionsFromState(diff)
			if err != nil {
				// TODO: return error here
				return nil
			}
			planned := PermissionsEntity{}
			common.DiffToStructPointer(diff, s, &planned)
			// Plan time validation for object permission levels
			for _, accessControl := range planned.AccessControlList {
				permissionLevel := accessControl.PermissionLevel
				if _, ok := mapping.allowedPermissionLevels[permissionLevel]; !ok {
					return fmt.Errorf(`permission_level %s is not supported with %s objects; allowed levels: %s`,
						permissionLevel, mapping.field, strings.Join(mapping.getAllowedPermissionLevels(false), ", "))
				}
			}
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			id := d.Id()
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			objectACL, err := NewPermissionsAPI(ctx, c).Read(id)
			if err != nil {
				return err
			}
			me, err := w.CurrentUser.Me(ctx)
			if err != nil {
				return err
			}
			var existing PermissionsEntity
			common.DataToStructPointer(d, s, &existing)
			entity, err := objectACL.ToPermissionsEntity(d, existing, me.UserName)
			if err != nil {
				return err
			}
			if len(entity.AccessControlList) == 0 {
				// empty "modifiable" access control list is the same as resource absence
				d.SetId("")
				return nil
			}
			return common.StructToData(entity, s, d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var entity PermissionsEntity
			common.DataToStructPointer(d, s, &entity)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			mapping, v, err := getResourcePermissionsFromState(d)
			if err != nil {
				return err
			}
			objectID, err := mapping.getID(ctx, w, v)
			if err != nil {
				return err
			}
			err = NewPermissionsAPI(ctx, c).Update(objectID, AccessControlChangeList{
				AccessControlList: entity.AccessControlList,
			}, mapping)
			if err != nil {
				return err
			}
			d.SetId(objectID)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var entity PermissionsEntity
			common.DataToStructPointer(d, s, &entity)
			mapping, err := getResourcePermissions(d.Id())
			if err != nil {
				return err
			}
			return NewPermissionsAPI(ctx, c).Update(d.Id(), AccessControlChangeList{
				AccessControlList: entity.AccessControlList,
			}, mapping)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			mapping, err := getResourcePermissions(d.Id())
			if err != nil {
				return err
			}
			return NewPermissionsAPI(ctx, c).Delete(d.Id(), mapping)
		},
	}
}
