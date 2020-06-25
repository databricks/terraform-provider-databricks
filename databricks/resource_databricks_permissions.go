package databricks

import (
	"fmt"
	"path"
	"strconv"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func parsePermissionsFromData(d *schema.ResourceData,
	client *service.DBApiClient) (*model.AccessControlChangeList, string, error) {
	var objectID string
	acl := new(model.AccessControlChangeList)
	for _, mapping := range permissionsResourceIDFields() {
		v, ok := d.GetOk(mapping.field)
		if !ok {
			continue
		}
		id, err := mapping.idRetriever(client, v.(string))
		if err != nil {
			return nil, "", err
		}
		objectID = fmt.Sprintf("/%s/%s", mapping.resourceType, id)
		err = d.Set("object_type", mapping.objectType)
		if err != nil {
			return nil, "", err
		}
	}
	if objectID == "" {
		return nil, "", fmt.Errorf("At least one type of resource identifiers must be set")
	}
	changes := 0
	if data, ok := d.GetOk("access_control"); ok {
		for _, element := range data.([]interface{}) {
			rawAccessControl := element.(map[string]interface{})
			change := new(model.AccessControlChange)
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
		return nil, "", fmt.Errorf("At least one access_control is required")
	}
	return acl, objectID, nil
}

func resourcePermissionsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	acl, objectID, err := parsePermissionsFromData(d, client)
	if err != nil {
		return err
	}
	err = client.Permissions().AddOrModify(objectID, acl)
	if err != nil {
		return err
	}
	d.SetId(objectID)
	return resourcePermissionsRead(d, m)
}

func resourcePermissionsRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	objectACL, err := client.Permissions().Read(id)
	if e, ok := err.(service.APIError); ok && e.IsMissing() {
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}
	for _, mapping := range permissionsResourceIDFields() {
		if mapping.objectType != d.Get("object_type").(string) {
			continue
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
	me, err := client.Users().Me()
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
	client := m.(*service.DBApiClient)
	err := client.Permissions().SetOrDelete(id, new(model.AccessControlChangeList))
	if err != nil {
		return err
	}
	return nil
}

// permissionsIDFieldMapping holds mapping
type permissionsIDFieldMapping struct {
	field        string
	objectType   string
	resourceType string
	idRetriever  func(client *service.DBApiClient, id string) (string, error)
}

// PermissionsResourceIDFields shows mapping of id columns to resource types
func permissionsResourceIDFields() []permissionsIDFieldMapping {
	SIMPLE := func(client *service.DBApiClient, id string) (string, error) {
		return id, nil
	}
	PATH := func(client *service.DBApiClient, path string) (string, error) {
		info, err := client.Notebooks().Read(path)
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

func resourcePermissions() *schema.Resource {
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
						//ConflictsWith: []string{"group_name"},
					},
					"group_name": {
						ForceNew: true,
						Type:     schema.TypeString,
						Optional: true,
						//ConflictsWith: []string{"user_name"},
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
