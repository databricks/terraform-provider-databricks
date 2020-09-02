package identity

import (
	"log"
	"reflect"
	"sort"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type scimUser struct {
	UserName       string   `json:"user_name"`
	DisplayName    string   `json:"display_name,omitempty"`
	Roles          []string `json:"roles,omitempty" tf:"slice_set"`
	Entitlements   []string `json:"entitlements,omitempty" tf:"slice_set"`
	DefaultRoles   []string `json:"default_roles" tf:"slice_set"`
	InheritedRoles []string `json:"inherited_roles,omitempty" tf:"slice_set,computed"`
	SetAdmin       bool     `json:"set_admin,omitempty"`
}

// ResourceScimUser ..
func ResourceScimUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceScimUserCreate,
		Update: resourceScimUserUpdate,
		Read:   resourceScimUserRead,
		Delete: resourceScimUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: internal.StructToSchema(scimUser{}, func(
			s map[string]*schema.Schema) map[string]*schema.Schema {
			s["user_name"].ForceNew = true
			s["set_admin"].Default = false
			return s
		}),
	}
}

func convertInterfaceSliceToStringSlice(input []interface{}) []string {
	resp := []string{}
	for _, item := range input {
		resp = append(resp, item.(string))
	}
	sort.Strings(resp)
	return resp
}

func resourceScimUserCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	userName := d.Get("user_name").(string)
	setAdmin := d.Get("set_admin").(bool)
	var displayName string
	var roles []string
	var entitlements []string
	if rDisplayName, ok := d.GetOk("display_name"); ok {
		displayName = rDisplayName.(string)
	}
	if rRoles, ok := d.GetOk("roles"); ok {
		roles = convertInterfaceSliceToStringSlice(rRoles.(*schema.Set).List())
		log.Println(roles)
	}
	if rEntitlements, ok := d.GetOk("entitlements"); ok {
		entitlements = convertInterfaceSliceToStringSlice(rEntitlements.(*schema.Set).List())
		log.Println(entitlements)
	}
	user, err := NewUsersAPI(client).Create(userName, displayName, entitlements, roles)
	if err != nil {
		return err
	}

	// Hack to fix user, for some reason if entitlements is empty it will auto create user with
	// allow-cluster-create permissions so we will apply a put operation to overwrite the user
	err = NewUsersAPI(client).Update(user.ID, userName, displayName, entitlements, roles)
	if err != nil {
		return err
	}

	if setAdmin {
		adminGroup, err := NewGroupsAPI(client).GetAdminGroup()
		if err != nil {
			return err
		}
		err = NewUsersAPI(client).SetUserAsAdmin(user.ID, adminGroup.ID)
		if err != nil {
			return err
		}
	}
	d.SetId(user.ID)

	return resourceScimUserRead(d, m)
}

func getListOfRoles(roleList []RoleListItem) []string {
	resp := []string{}
	for _, role := range roleList {
		resp = append(resp, role.Value)
	}
	return resp
}

func resourceScimUserRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	user, err := NewUsersAPI(client).Read(id)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}

	adminGroup, err := NewGroupsAPI(client).GetAdminGroup()
	if err != nil {
		return err
	}
	isAdmin, err := NewUsersAPI(client).VerifyUserAsAdmin(user.ID, adminGroup.ID)
	if err != nil {
		return err
	}
	err = d.Set("set_admin", isAdmin)
	if err != nil {
		return err
	}

	//entitlements := getListOfEntitlements(user.Entitlements)
	var entitlements []string
	for _, entitlement := range user.Entitlements {
		entitlements = append(entitlements, string(entitlement.Value))
	}
	err = d.Set("entitlements", entitlements)
	if err != nil {
		return err
	}

	err = d.Set("display_name", user.DisplayName)
	if err != nil {
		return err
	}

	roles := getListOfRoles(user.Roles)
	inheritedRoles := getListOfRoles(user.InheritedRoles)
	err = d.Set("inherited_roles", inheritedRoles)
	if err != nil {
		return err
	}

	myActualRolesThatAreNotInheritedFromGroups := diff(roles, inheritedRoles)

	defaultRolesInterfaceList := d.Get("default_roles").(*schema.Set).List()
	allUserInheritedRoles := internal.ConvertListInterfaceToString(defaultRolesInterfaceList)
	//allUserInheritedRoles := getListOfRoles(metaUser.Roles)
	myActualRolesThatAreNotInherited := diff(myActualRolesThatAreNotInheritedFromGroups, allUserInheritedRoles)
	setRoles, ok := d.GetOk("roles")
	if ok {
		allRules := append(inheritedRoles, allUserInheritedRoles...)
		for _, role := range setRoles.(*schema.Set).List() {
			if sliceContains(role.(string), allRules) {
				myActualRolesThatAreNotInherited = append(myActualRolesThatAreNotInherited, role.(string))
			}
		}
	}

	log.Println(myActualRolesThatAreNotInherited)
	err = d.Set("roles", myActualRolesThatAreNotInherited)
	if err != nil {
		return err
	}

	return err
}

func resourceScimUserUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	userName := d.Get("user_name").(string)
	var displayName string
	var roles []string
	var entitlements []string
	if rDisplayName, ok := d.GetOk("display_name"); ok {
		displayName = rDisplayName.(string)
	}
	if rRoles, ok := d.GetOk("roles"); ok {
		roles = convertInterfaceSliceToStringSlice(rRoles.(*schema.Set).List())
	}
	if rEntitlements, ok := d.GetOk("entitlements"); ok {
		entitlements = convertInterfaceSliceToStringSlice(rEntitlements.(*schema.Set).List())
	}
	err := NewUsersAPI(client).Update(id, userName, displayName, entitlements, roles)
	if err != nil {
		return err
	}
	if d.HasChange("set_admin") {
		setAdmin := d.Get("set_admin").(bool)
		if setAdmin {
			adminGroup, err := NewGroupsAPI(client).GetAdminGroup()
			if err != nil {
				return err
			}
			err = NewUsersAPI(client).SetUserAsAdmin(id, adminGroup.ID)
			if err != nil {
				return err
			}
		} else {
			adminGroup, err := NewGroupsAPI(client).GetAdminGroup()
			if err != nil {
				return err
			}
			err = NewUsersAPI(client).RemoveUserAsAdmin(id, adminGroup.ID)
			if err != nil {
				return err
			}
		}
	}
	return resourceScimUserRead(d, m)
}

func resourceScimUserDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	err := NewUsersAPI(client).Delete(id)
	return err
}

func sliceContains(value string, list []string) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}
