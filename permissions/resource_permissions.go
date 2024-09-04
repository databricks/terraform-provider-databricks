package permissions

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path"
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

// As described in https://github.com/databricks/terraform-provider-databricks/issues/1504,
// certain object types require that we explicitly grant the calling user CAN_MANAGE
// permissions when POSTing permissions changes through the REST API, to avoid accidentally
// revoking the calling user's ability to manage the current object.
func (a PermissionsAPI) shouldExplicitlyGrantCallingUserManagePermissions(objectID string) bool {
	for _, prefix := range [...]string{"/registered-models/", "/clusters/", "/instance-pools/", "/serving-endpoints/", "/queries/", "/sql/warehouses"} {
		if strings.HasPrefix(objectID, prefix) {
			return true
		}
	}
	return isDbsqlPermissionsWorkaroundNecessary(objectID)
}

func isOwnershipWorkaroundNecessary(objectID string) bool {
	return strings.HasPrefix(objectID, "/jobs") || strings.HasPrefix(objectID, "/pipelines") || strings.HasPrefix(objectID, "/sql/warehouses")
}

func (a PermissionsAPI) getObjectCreator(objectID string) (string, error) {
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(objectID, "/jobs") {
		jobId, err := strconv.ParseInt(strings.ReplaceAll(objectID, "/jobs/", ""), 10, 64)
		if err != nil {
			return "", err
		}
		job, err := w.Jobs.GetByJobId(a.context, jobId)
		if err != nil {
			return "", common.IgnoreNotFoundError(err)
		}
		return job.CreatorUserName, nil
	} else if strings.HasPrefix(objectID, "/pipelines") {
		pipeline, err := w.Pipelines.GetByPipelineId(a.context, strings.ReplaceAll(objectID, "/pipelines/", ""))
		if err != nil {
			return "", common.IgnoreNotFoundError(err)
		}
		return pipeline.CreatorUserName, nil
	} else if strings.HasPrefix(objectID, "/sql/warehouses") {
		warehouse, err := w.Warehouses.GetById(a.context, strings.ReplaceAll(objectID, "/sql/warehouses/", ""))
		if err != nil {
			return "", common.IgnoreNotFoundError(err)
		}
		return warehouse.CreatorName, nil
	}
	return "", nil
}

func (a PermissionsAPI) ensureCurrentUserCanManageObject(objectID string, objectACL AccessControlChangeList) (AccessControlChangeList, error) {
	if !a.shouldExplicitlyGrantCallingUserManagePermissions(objectID) {
		return objectACL, nil
	}
	w, err := a.client.WorkspaceClient()
	if err != nil {
		return objectACL, err
	}
	me, err := w.CurrentUser.Me(a.context)
	if err != nil {
		return objectACL, err
	}
	objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
		UserName:        me.UserName,
		PermissionLevel: "CAN_MANAGE",
	})
	return objectACL, nil
}

// Helper function for applying permissions changes. Ensures that
// we select the correct HTTP method based on the object type and preserve the calling
// user's ability to manage the specified object when applying permissions changes.
func (a PermissionsAPI) put(objectID string, objectACL AccessControlChangeList) error {
	objectACL, err := a.ensureCurrentUserCanManageObject(objectID, objectACL)
	if err != nil {
		return err
	}
	if isDbsqlPermissionsWorkaroundNecessary(objectID) {
		// SQLA entities use POST for permission updates.
		return a.client.Post(a.context, urlPathForObjectID(objectID), objectACL, nil)
	}
	log.Printf("[DEBUG] PUT %s %v", objectID, objectACL)
	return a.client.Put(a.context, urlPathForObjectID(objectID), objectACL)
}

// safePutWithOwner is a workaround for the limitation where warehouse without owners cannot have IS_OWNER set
func (a PermissionsAPI) safePutWithOwner(objectID string, objectACL AccessControlChangeList, originalAcl []AccessControlChange) error {
	err := a.put(objectID, objectACL)
	if err != nil {
		if strings.Contains(err.Error(), "with no existing owner must provide a new owner") {
			objectACL.AccessControlList = originalAcl
			return a.put(objectID, objectACL)
		}
		return err
	}
	return nil
}

// Update updates object permissions. Technically, it's using method named SetOrDelete, but here we do more
func (a PermissionsAPI) Update(objectID string, objectACL AccessControlChangeList) error {
	if objectID == "/authorization/tokens" || objectID == "/registered-models/root" || objectID == "/directories/0" {
		// Prevent "Cannot change permissions for group 'admins' to None."
		objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
			GroupName:       "admins",
			PermissionLevel: "CAN_MANAGE",
		})
	}
	originalAcl := make([]AccessControlChange, len(objectACL.AccessControlList))
	_ = copy(originalAcl, objectACL.AccessControlList)
	if isOwnershipWorkaroundNecessary(objectID) {
		owners := 0
		for _, acl := range objectACL.AccessControlList {
			if acl.PermissionLevel == "IS_OWNER" {
				owners++
			}
		}
		if owners == 0 {
			w, err := a.client.WorkspaceClient()
			if err != nil {
				return err
			}
			me, err := w.CurrentUser.Me(a.context)
			if err != nil {
				return err
			}
			// add owner if it's missing, otherwise automated planning might be difficult
			objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
				UserName:        me.UserName,
				PermissionLevel: "IS_OWNER",
			})
		}
	}
	return a.safePutWithOwner(objectID, objectACL, originalAcl)
}

// Delete gracefully removes permissions. Technically, it's using method named SetOrDelete, but here we do more
func (a PermissionsAPI) Delete(objectID string) error {
	objectACL, err := a.Read(objectID)
	if err != nil {
		return err
	}
	accl := AccessControlChangeList{}
	for _, acl := range objectACL.AccessControlList {
		if acl.GroupName == "admins" && objectID != "/authorization/passwords" {
			if change, direct := acl.toAccessControlChange(); direct {
				// keep everything direct for admin group
				accl.AccessControlList = append(accl.AccessControlList, change)
			}
		}
	}
	originalAcl := make([]AccessControlChange, len(accl.AccessControlList))
	_ = copy(originalAcl, accl.AccessControlList)
	if isOwnershipWorkaroundNecessary(objectID) {
		creator, err := a.getObjectCreator(objectID)
		if err != nil {
			return err
		}
		if creator == "" {
			return nil
		}
		accl.AccessControlList = append(accl.AccessControlList, AccessControlChange{
			UserName:        creator,
			PermissionLevel: "IS_OWNER",
		})
	}
	return a.safePutWithOwner(objectID, accl, originalAcl)
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

// permissionsIDFieldMapping holds mapping
type permissionsIDFieldMapping struct {
	field, objectType, resourceType string

	allowedPermissionLevels []string

	idRetriever func(ctx context.Context, w *databricks.WorkspaceClient, id string) (string, error)
}

// PermissionsResourceIDFields shows mapping of id columns to resource types
func permissionsResourceIDFields() []permissionsIDFieldMapping {
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
	return []permissionsIDFieldMapping{
		{"cluster_policy_id", "cluster-policy", "cluster-policies", []string{"CAN_USE"}, SIMPLE},
		{"instance_pool_id", "instance-pool", "instance-pools", []string{"CAN_ATTACH_TO", "CAN_MANAGE"}, SIMPLE},
		{"cluster_id", "cluster", "clusters", []string{"CAN_ATTACH_TO", "CAN_RESTART", "CAN_MANAGE"}, SIMPLE},
		{"pipeline_id", "pipelines", "pipelines", []string{"CAN_VIEW", "CAN_RUN", "CAN_MANAGE", "IS_OWNER"}, SIMPLE},
		{"job_id", "job", "jobs", []string{"CAN_VIEW", "CAN_MANAGE_RUN", "IS_OWNER", "CAN_MANAGE"}, SIMPLE},
		{"notebook_id", "notebook", "notebooks", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, SIMPLE},
		{"notebook_path", "notebook", "notebooks", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, PATH},
		{"directory_id", "directory", "directories", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, SIMPLE},
		{"directory_path", "directory", "directories", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, PATH},
		{"workspace_file_id", "file", "files", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, SIMPLE},
		{"workspace_file_path", "file", "files", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, PATH},
		{"repo_id", "repo", "repos", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, SIMPLE},
		{"repo_path", "repo", "repos", []string{"CAN_READ", "CAN_RUN", "CAN_EDIT", "CAN_MANAGE"}, PATH},
		{"authorization", "tokens", "authorization", []string{"CAN_USE"}, SIMPLE},
		{"authorization", "passwords", "authorization", []string{"CAN_USE"}, SIMPLE},
		{"sql_endpoint_id", "warehouses", "sql/warehouses", []string{"CAN_USE", "CAN_MANAGE", "CAN_MONITOR", "IS_OWNER"}, SIMPLE},
		{"sql_dashboard_id", "dashboard", "sql/dashboards", []string{"CAN_EDIT", "CAN_RUN", "CAN_MANAGE", "CAN_VIEW"}, SIMPLE},
		{"sql_alert_id", "alert", "sql/alerts", []string{"CAN_EDIT", "CAN_RUN", "CAN_MANAGE", "CAN_VIEW"}, SIMPLE},
		{"sql_query_id", "query", "sql/queries", []string{"CAN_EDIT", "CAN_RUN", "CAN_MANAGE", "CAN_VIEW"}, SIMPLE},
		{"dashboard_id", "dashboard", "dashboards", []string{"CAN_EDIT", "CAN_RUN", "CAN_MANAGE", "CAN_READ"}, SIMPLE},
		{"experiment_id", "mlflowExperiment", "experiments", []string{"CAN_READ", "CAN_EDIT", "CAN_MANAGE"}, SIMPLE},
		{"registered_model_id", "registered-model", "registered-models", []string{
			"CAN_READ", "CAN_EDIT", "CAN_MANAGE_STAGING_VERSIONS", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE"}, SIMPLE},
		{"serving_endpoint_id", "serving-endpoint", "serving-endpoints", []string{"CAN_VIEW", "CAN_QUERY", "CAN_MANAGE"}, SIMPLE},
	}
}

// PermissionsEntity is the one used for resource metadata
type PermissionsEntity struct {
	ObjectType        string                `json:"object_type,omitempty" tf:"computed"`
	AccessControlList []AccessControlChange `json:"access_control" tf:"slice_set"`
}

func (oa *ObjectACL) isMatchingMapping(mapping permissionsIDFieldMapping) bool {
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

func (oa *ObjectACL) ToPermissionsEntity(d *schema.ResourceData, me string) (PermissionsEntity, error) {
	entity := PermissionsEntity{}
	for _, accessControl := range oa.AccessControlList {
		if accessControl.GroupName == "admins" && d.Id() != "/authorization/passwords" {
			// not possible to lower admins permissions anywhere from CAN_MANAGE
			continue
		}
		if me == accessControl.UserName || me == accessControl.ServicePrincipalName {
			// not possible to lower one's permissions anywhere from CAN_MANAGE
			continue
		}
		if change, direct := accessControl.toAccessControlChange(); direct {
			entity.AccessControlList = append(entity.AccessControlList, change)
		}
	}
	for _, mapping := range permissionsResourceIDFields() {
		if !oa.isMatchingMapping(mapping) {
			continue
		}
		entity.ObjectType = mapping.objectType
		var pathVariant any
		if mapping.objectType == "file" {
			pathVariant = d.Get("workspace_file_path")
		} else {
			pathVariant = d.Get(mapping.objectType + "_path")
		}
		if pathVariant != nil && pathVariant.(string) != "" {
			// we're not importing and it's a path... it's set, so let's not re-set it
			return entity, nil
		}
		identifier := path.Base(oa.ObjectID)
		return entity, d.Set(mapping.field, identifier)
	}
	return entity, fmt.Errorf("unknown object type %s", oa.ObjectType)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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
			// Plan time validation for object permission levels
			for _, mapping := range permissionsResourceIDFields() {
				if _, ok := diff.GetOk(mapping.field); !ok {
					continue
				}
				access_control_list := diff.Get("access_control").(*schema.Set).List()
				for _, access_control := range access_control_list {
					m := access_control.(map[string]any)
					permission_level := m["permission_level"].(string)
					if !stringInSlice(permission_level, mapping.allowedPermissionLevels) {
						return fmt.Errorf(`permission_level %s is not supported with %s objects`,
							permission_level, mapping.field)
					}
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
			entity, err := objectACL.ToPermissionsEntity(d, me.UserName)
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
			me, err := w.CurrentUser.Me(ctx)
			if err != nil {
				return err
			}
			for _, mapping := range permissionsResourceIDFields() {
				if v, ok := d.GetOk(mapping.field); ok {
					id, err := mapping.idRetriever(ctx, w, v.(string))
					if err != nil {
						return err
					}
					objectID := fmt.Sprintf("/%s/%s", mapping.resourceType, id)
					// this logic was moved from CustomizeDiff because of undeterministic auth behavior
					// in the corner-case scenarios.
					// see https://github.com/databricks/terraform-provider-databricks/issues/2052
					for _, v := range entity.AccessControlList {
						if v.UserName == me.UserName {
							format := "it is not possible to decrease administrative permissions for the current user: %s"
							return fmt.Errorf(format, me.UserName)
						}

						if v.GroupName == "admins" && mapping.resourceType != "authorization" {
							// should allow setting admins permissions for passwords and tokens usage
							return fmt.Errorf("it is not possible to restrict any permissions from `admins`")
						}
					}
					err = NewPermissionsAPI(ctx, c).Update(objectID, AccessControlChangeList{
						AccessControlList: entity.AccessControlList,
					})
					if err != nil {
						return err
					}
					d.SetId(objectID)
					return nil
				}
			}
			return errors.New("at least one type of resource identifiers must be set")
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var entity PermissionsEntity
			common.DataToStructPointer(d, s, &entity)
			return NewPermissionsAPI(ctx, c).Update(d.Id(), AccessControlChangeList{
				AccessControlList: entity.AccessControlList,
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewPermissionsAPI(ctx, c).Delete(d.Id())
		},
	}
}
