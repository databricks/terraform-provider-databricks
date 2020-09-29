package access

import (
	"fmt"
	"path"
	"sort"
	"strconv"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/compute"
	"github.com/databrickslabs/databricks-terraform/identity"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// ObjectACL is a structure to generically describe access control
type ObjectACL struct {
	ObjectID          string           `json:"object_id,omitempty"`
	ObjectType        string           `json:"object_type,omitempty"`
	AccessControlList []*AccessControl `json:"access_control_list"`
}

// AccessControl is a structure to describe user/group permissions
type AccessControl struct {
	UserName       *string       `json:"user_name,omitempty"`
	GroupName      *string       `json:"group_name,omitempty"`
	AllPermissions []*Permission `json:"all_permissions,omitempty"`
}

func (ac AccessControl) String() string {
	s := ""
	switch {
	case ac.GroupName != nil:
		s += *ac.GroupName
	case ac.UserName != nil:
		s += *ac.UserName
	default:
		s += "something"
	}
	s += " "
	for _, ap := range ac.AllPermissions {
		if ap == nil {
			continue
		}
		s += ap.String()
	}
	return s
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
	AccessControlList []*AccessControlChange `json:"access_control_list"`
}

// AccessControlChange is API wrapper for changing permissions
type AccessControlChange struct {
	UserName             *string `json:"user_name,omitempty"`
	GroupName            *string `json:"group_name,omitempty"`
	ServicePrincipalName *string `json:"service_principal_name,omitempty"`
	PermissionLevel      string  `json:"permission_level"`
}

func (acc AccessControlChange) String() string {
	return fmt.Sprintf("%v%v%v %s",
		acc.UserName, acc.GroupName, acc.ServicePrincipalName,
		acc.PermissionLevel)
}

// ToAccessControlChangeList converts data formats
func (oa *ObjectACL) ToAccessControlChangeList() *AccessControlChangeList {
	acl := new(AccessControlChangeList)
	for _, accessControl := range oa.AccessControlList {
		for _, permission := range accessControl.AllPermissions {
			if permission.Inherited {
				continue
			}
			item := new(AccessControlChange)
			acl.AccessControlList = append(acl.AccessControlList, item)
			item.PermissionLevel = permission.PermissionLevel
			if accessControl.UserName != nil {
				item.UserName = accessControl.UserName
			} else if accessControl.GroupName != nil {
				item.GroupName = accessControl.GroupName
			}
		}
	}
	acl.Sort()
	return acl
}

// Sort AccessControlList for consistent results
func (acl *AccessControlChangeList) Sort() {
	sort.Slice(acl.AccessControlList, func(i, j int) bool {
		return acl.AccessControlList[i].String() > acl.AccessControlList[j].String()
	})
}

// AccessControl exports data for TF
func (acl *AccessControlChangeList) AccessControl(me string) []map[string]string {
	result := []map[string]string{}
	for _, control := range acl.AccessControlList {
		item := map[string]string{}
		if control.UserName != nil && *control.UserName != "" {
			if me == *control.UserName {
				continue
			}
			item["user_name"] = *control.UserName
		} else if control.GroupName != nil && *control.GroupName != "" {
			item["group_name"] = *control.GroupName
		}
		item["permission_level"] = control.PermissionLevel
		result = append(result, item)
	}
	return result
}

// NewPermissionsAPI creates PermissionsAPI instance from provider meta
func NewPermissionsAPI(m interface{}) PermissionsAPI {
	return PermissionsAPI{client: m.(*common.DatabricksClient)}
}

// PermissionsAPI exposes general permission related methods
type PermissionsAPI struct {
	client *common.DatabricksClient
}

// AddOrModify works with permissions change list
func (a PermissionsAPI) AddOrModify(objectID string, objectACL *AccessControlChangeList) error {
	return a.client.Patch("/preview/permissions"+objectID, objectACL)
}

// SetOrDelete updates object permissions
func (a PermissionsAPI) SetOrDelete(objectID string, objectACL *AccessControlChangeList) error {
	return a.client.Put("/preview/permissions"+objectID, objectACL)
}

// Read gets all relevant permissions for the object, including inherited ones
func (a PermissionsAPI) Read(objectID string) (objectACL ObjectACL, err error) {
	err = a.client.Get("/preview/permissions"+objectID, nil, &objectACL)
	return
}

func parsePermissionsFromData(d *schema.ResourceData,
	client *common.DatabricksClient) (*AccessControlChangeList, string, error) {
	var objectId string
	acl := new(AccessControlChangeList)
	for _, mapping := range permissionsResourceIDFields() {
		v, ok := d.GetOk(mapping.field)
		if !ok {
			continue
		}
		id, err := mapping.idRetriever(client, v.(string))
		if err != nil {
			return nil, "", err
		}
		objectId = fmt.Sprintf(
			"/%s/%s",
			mapping.resourceType, id)
		err = d.Set("object_type", mapping.objectType)
		if err != nil {
			return nil, "", err
		}
	}
	if objectId == "" {
		return nil, "", fmt.Errorf("At least one type of resource identifiers must be set")
	}
	changes := 0
	if data, ok := d.GetOk("access_control"); ok {
		for _, element := range data.([]interface{}) {
			rawAccessControl := element.(map[string]interface{})
			change := new(AccessControlChange)
			acl.AccessControlList = append(acl.AccessControlList, change)
			if v, ok := rawAccessControl["group_name"].(string); ok && v != "" {
				change.GroupName = &v
			}
			if v, ok := rawAccessControl["user_name"].(string); ok && v != "" {
				change.UserName = &v
			}
			if v, ok := rawAccessControl["permission_level"].(string); ok {
				change.PermissionLevel = v
			}
			changes++
		}
	}
	if changes < 1 {
		return nil, "", fmt.Errorf("at least one access_control is required")
	}
	acl.Sort()
	return acl, objectId, nil
}

func resourcePermissionsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	acl, objectID, err := parsePermissionsFromData(d, client)
	if err != nil {
		return err
	}
	err = NewPermissionsAPI(m).AddOrModify(objectID, acl)
	if err != nil {
		return err
	}
	d.SetId(objectID)
	return resourcePermissionsRead(d, m)
}

func resourcePermissionsRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	objectACL, err := NewPermissionsAPI(m).Read(id)
	if e, ok := err.(common.APIError); ok && e.IsMissing() {
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}
	for _, mapping := range permissionsResourceIDFields() {
		if mapping.objectType != objectACL.ObjectType {
			continue
		}
		err = d.Set("object_type", mapping.objectType)
		if err != nil {
			return fmt.Errorf("Cannot set object type: %v", mapping.objectType)
		}
		pathVariant := d.Get(mapping.objectType + "_path")
		if pathVariant != "" {
			// we're not importing and it's a path... it's set, so let's not re-set it
			break
		}
		identifier := path.Base(id)
		err := d.Set(mapping.field, identifier)
		if err != nil {
			return errors.Wrapf(err,
				"Cannot set mapping field %s to %s",
				mapping.field, id)
		}
		break
	}
	acl := objectACL.ToAccessControlChangeList()
	me, err := identity.NewUsersAPI(m).Me()
	if err != nil {
		return errors.Wrapf(err, "Cannot self-identify")
	}
	accessControl := acl.AccessControl(me.UserName)
	err = d.Set("access_control", accessControl)
	if err != nil {
		return err
	}
	return nil
}

func resourcePermissionsDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	objectType := d.Get("object_type").(string)

	if objectType == "job" {
		job, err := compute.NewJobsAPI(m).Read(id)

		if err != nil {
			return err
		}
		userName := job.CreatorUserName

		userACL := AccessControlChange{
			UserName:        &userName,
			PermissionLevel: "IS_OWNER",
		}

		accessControlChange := []*AccessControlChange{&userACL}

		resourceACL := AccessControlChangeList{
			accessControlChange,
		}
		return NewPermissionsAPI(m).SetOrDelete(id, &resourceACL)
	}
	return NewPermissionsAPI(m).SetOrDelete(id, new(AccessControlChangeList))

}

// permissionsIDFieldMapping holds mapping
type permissionsIDFieldMapping struct {
	field        string
	objectType   string
	resourceType string
	idRetriever  func(client *common.DatabricksClient, id string) (string, error)
}

// PermissionsResourceIDFields shows mapping of id columns to resource types
func permissionsResourceIDFields() []permissionsIDFieldMapping {
	SIMPLE := func(client *common.DatabricksClient, id string) (string, error) {
		return id, nil
	}
	PATH := func(client *common.DatabricksClient, path string) (string, error) {
		info, err := workspace.NewNotebooksAPI(client).Read(path)
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
	}
}

func conflictingFields(field string) []string {
	conflicting := []string{}
	for _, mapping := range permissionsResourceIDFields() {
		if mapping.field == field {
			continue
		}
		conflicting = append(conflicting, mapping.field)
	}
	return conflicting
}

// ResourcePermissions definition
func ResourcePermissions() *schema.Resource {
	fields := map[string]*schema.Schema{
		"object_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"access_control": {
			ForceNew:   true,
			Type:       schema.TypeList,
			MinItems:   1,
			Required:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"user_name": {
						ForceNew: true,
						Type:     schema.TypeString,
						Optional: true,
					},
					"group_name": {
						ForceNew: true,
						Type:     schema.TypeString,
						Optional: true,
					},
					"permission_level": {
						ForceNew: true,
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
	for _, mapping := range permissionsResourceIDFields() {
		fields[mapping.field] = &schema.Schema{
			ForceNew:      true,
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: conflictingFields(mapping.field),
		}
	}
	return &schema.Resource{
		Create: resourcePermissionsCreate,
		Read:   resourcePermissionsRead,
		Delete: resourcePermissionsDelete,
		Schema: fields,
	}
}
