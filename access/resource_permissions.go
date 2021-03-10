package access

import (
	"context"
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/databrickslabs/terraform-provider-databricks/identity"

	"github.com/databrickslabs/terraform-provider-databricks/workspace"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
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
func NewPermissionsAPI(ctx context.Context, m interface{}) PermissionsAPI {
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

func urlPathForObjectID(objectID string) string {
	if strings.HasPrefix(objectID, "/sql/") {
		// Permissions for SQLA entities are routed differently from the others.
		return "/preview/sql/permissions" + objectID[4:]
	}
	return "/preview/permissions" + objectID
}

// Helper function to select the correct HTTP method depending on the object types.
func (a PermissionsAPI) put(objectID string, objectACL AccessControlChangeList) error {
	if strings.HasPrefix(objectID, "/sql/") {
		// SQLA entities always have `CAN_MANAGE` permission for the calling user.
		me, err := identity.NewUsersAPI(a.context, a.client).Me()
		if err != nil {
			return err
		}
		objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
			UserName:        me.UserName,
			PermissionLevel: "CAN_MANAGE",
		})

		// The SQLA entities use HTTP POST for permission updates.
		return a.client.Post(a.context, urlPathForObjectID(objectID), objectACL, nil)
	}

	return a.client.Put(a.context, urlPathForObjectID(objectID), objectACL)
}

// Update updates object permissions. Technically, it's using method named SetOrDelete, but here we do more
func (a PermissionsAPI) Update(objectID string, objectACL AccessControlChangeList) error {
	if "/authorization/tokens" == objectID {
		// Cannot remove admins's CAN_MANAGE permission on tokens
		objectACL.AccessControlList = append(objectACL.AccessControlList, AccessControlChange{
			GroupName:       "admins",
			PermissionLevel: "CAN_MANAGE",
		})
	}
	if strings.HasPrefix(objectID, "/jobs") {
		owners := 0
		for _, acl := range objectACL.AccessControlList {
			if acl.PermissionLevel == "IS_OWNER" {
				owners++
			}
		}
		if owners == 0 {
			me, err := identity.NewUsersAPI(a.context, a.client).Me()
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
	return a.put(objectID, objectACL)
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
	if strings.HasPrefix(objectID, "/jobs") {
		job, err := compute.NewJobsAPI(a.context, a.client).Read(strings.ReplaceAll(objectID, "/jobs/", ""))
		if err != nil {
			return err
		}
		accl.AccessControlList = append(accl.AccessControlList, AccessControlChange{
			UserName:        job.CreatorUserName,
			PermissionLevel: "IS_OWNER",
		})
	}
	return a.put(objectID, accl)
}

// Read gets all relevant permissions for the object, including inherited ones
func (a PermissionsAPI) Read(objectID string) (objectACL ObjectACL, err error) {
	err = a.client.Get(a.context, urlPathForObjectID(objectID), nil, &objectACL)
	return
}

// permissionsIDFieldMapping holds mapping
type permissionsIDFieldMapping struct {
	field, objectType, resourceType string

	idRetriever func(client *common.DatabricksClient, id string) (string, error)
}

// PermissionsResourceIDFields shows mapping of id columns to resource types
func permissionsResourceIDFields(ctx context.Context) []permissionsIDFieldMapping {
	SIMPLE := func(client *common.DatabricksClient, id string) (string, error) {
		return id, nil
	}
	PATH := func(client *common.DatabricksClient, path string) (string, error) {
		info, err := workspace.NewNotebooksAPI(ctx, client).Read(path)
		if err != nil {
			return "", errors.Wrapf(err, "Cannot load path %s", path)
		}
		return strconv.FormatInt(info.ObjectID, 10), nil
	}
	return []permissionsIDFieldMapping{
		{"cluster_policy_id", "cluster-policy", "cluster-policies", SIMPLE},
		{"instance_pool_id", "instance-pool", "instance-pools", SIMPLE},
		{"cluster_id", "cluster", "clusters", SIMPLE},
		{"job_id", "job", "jobs", SIMPLE},
		{"notebook_id", "notebook", "notebooks", SIMPLE},
		{"notebook_path", "notebook", "notebooks", PATH},
		{"directory_id", "directory", "directories", SIMPLE},
		{"directory_path", "directory", "directories", PATH},
		{"authorization", "tokens", "authorization", SIMPLE},
		{"authorization", "passwords", "authorization", SIMPLE},
		{"sql_endpoint_id", "endpoint", "sql/endpoints", SIMPLE},
		{"sql_dashboard_id", "dashboard", "sql/dashboards", SIMPLE},
		{"sql_alert_id", "alert", "sql/alerts", SIMPLE},
		{"sql_query_id", "query", "sql/queries", SIMPLE},
	}
}

// PermissionsEntity is the one used for resource metadata
type PermissionsEntity struct {
	ObjectType        string                `json:"object_type,omitempty" tf:"computed"`
	AccessControlList []AccessControlChange `json:"access_control" tf:"slice_set"`
}

// ToPermissionsEntity ..
func (oa *ObjectACL) ToPermissionsEntity(ctx context.Context, d *schema.ResourceData, me string) (PermissionsEntity, error) {
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
	for _, mapping := range permissionsResourceIDFields(ctx) {
		if mapping.objectType != oa.ObjectType {
			continue
		}
		entity.ObjectType = mapping.objectType
		pathVariant := d.Get(mapping.objectType + "_path")
		if pathVariant != nil && pathVariant.(string) != "" {
			// we're not importing and it's a path... it's set, so let's not re-set it
			return entity, nil
		}
		identifier := path.Base(oa.ObjectID)
		err := d.Set(mapping.field, identifier)
		if err != nil {
			return entity, err
		}
		return entity, nil
	}
	return entity, fmt.Errorf("Unknown object type %s", oa.ObjectType)
}

// ResourcePermissions definition
func ResourcePermissions() *schema.Resource {
	s := common.StructToSchema(PermissionsEntity{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		ctx := context.Background()
		for _, mapping := range permissionsResourceIDFields(ctx) {
			s[mapping.field] = &schema.Schema{
				ForceNew: true,
				Type:     schema.TypeString,
				Optional: true,
			}
			for _, m := range permissionsResourceIDFields(ctx) {
				if m.field == mapping.field {
					continue
				}
				s[mapping.field].ConflictsWith = append(s[mapping.field].ConflictsWith, m.field)
			}
		}
		s["access_control"].MinItems = 1
		if groupNameSchema, err := common.SchemaPath(s,
			"access_control", "group_name"); err == nil {
			groupNameSchema.ValidateDiagFunc = func(i interface{}, p cty.Path) diag.Diagnostics {
				if v, ok := i.(string); ok {
					if "admins" == strings.ToLower(v) {
						return diag.Diagnostics{
							{
								Summary:       "It is not possible to restrict any permissions from `admins`.",
								Severity:      diag.Error,
								AttributePath: p,
							},
						}
					}
				}
				return nil
			}
		}
		return s
	})
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		id := d.Id()
		objectACL, err := NewPermissionsAPI(ctx, m).Read(id)
		if aerr, ok := err.(common.APIError); ok && aerr.IsMissing() {
			d.SetId("")
			return nil
		}
		if err != nil {
			return diag.FromErr(err)
		}
		me, err := identity.NewUsersAPI(ctx, m).Me()
		if err != nil {
			return diag.FromErr(err)
		}
		entity, err := objectACL.ToPermissionsEntity(ctx, d, me.UserName)
		if err != nil {
			return diag.FromErr(err)
		}
		if len(entity.AccessControlList) == 0 {
			// empty "modifiable" access control list is the same as resource absence
			d.SetId("")
			return nil
		}
		err = common.StructToData(entity, s, d)
		if err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	return &schema.Resource{
		Schema:      s,
		ReadContext: readContext,
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var entity PermissionsEntity
			err := common.DataToStructPointer(d, s, &entity)
			if err != nil {
				return diag.FromErr(err)
			}
			for _, mapping := range permissionsResourceIDFields(ctx) {
				if v, ok := d.GetOk(mapping.field); ok {
					id, err := mapping.idRetriever(m.(*common.DatabricksClient), v.(string))
					if err != nil {
						return diag.FromErr(err)
					}
					objectID := fmt.Sprintf("/%s/%s", mapping.resourceType, id)
					err = NewPermissionsAPI(ctx, m).Update(objectID, AccessControlChangeList{
						AccessControlList: entity.AccessControlList,
					})
					if err != nil {
						return diag.FromErr(err)
					}
					d.SetId(objectID)
					return readContext(ctx, d, m)
				}
			}
			return diag.Errorf("At least one type of resource identifiers must be set")
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var entity PermissionsEntity
			err := common.DataToStructPointer(d, s, &entity)
			if err != nil {
				return diag.FromErr(err)
			}
			err = NewPermissionsAPI(ctx, m).Update(d.Id(), AccessControlChangeList{
				AccessControlList: entity.AccessControlList,
			})
			if err != nil {
				return diag.FromErr(err)
			}
			return readContext(ctx, d, m)
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			err := NewPermissionsAPI(ctx, m).Delete(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
