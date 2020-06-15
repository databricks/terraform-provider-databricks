package databricks

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceScimUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceScimUserCreate,
		Update: resourceScimUserUpdate,
		Read:   resourceScimUserRead,
		Delete: resourceScimUserDelete,

		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"roles": {
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem:       &schema.Schema{Type: schema.TypeString},
				Set:        schema.HashString,
			},
			"entitlements": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"inherited_roles": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"default_roles": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"set_admin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Set:      schema.HashString,
			},
		},
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
	client := m.(*service.DBApiClient)
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
	user, err := client.Users().Create(userName, displayName, entitlements, roles)
	if err != nil {
		return err
	}

	// Hack to fix user, for some reason if entitlements is empty it will auto create user with
	// allow-cluster-create permissions so we will apply a put operation to overwrite the user
	err = client.Users().Update(user.ID, userName, displayName, entitlements, roles)
	if err != nil {
		return err
	}

	if setAdmin {
		adminGroup, err := client.Groups().GetAdminGroup()
		if err != nil {
			return err
		}
		err = client.Users().SetUserAsAdmin(user.ID, adminGroup.ID)
		if err != nil {
			return err
		}
	}
	d.SetId(user.ID)

	return resourceScimUserRead(d, m)
}

func getListOfRoles(roleList []model.RoleListItem) []string {
	resp := []string{}
	for _, role := range roleList {
		resp = append(resp, role.Value)
	}
	return resp
}

func resourceScimUserRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	user, err := client.Users().Read(id)
	if err != nil {
		if isScimUserMissing(err.Error(), id) {
			log.Printf("Missing scim user with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}

	adminGroup, err := client.Groups().GetAdminGroup()
	if err != nil {
		return err
	}
	isAdmin, err := client.Users().VerifyUserAsAdmin(user.ID, adminGroup.ID)
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
	allUserInheritedRoles := convertListInterfaceToString(defaultRolesInterfaceList)
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
	client := m.(*service.DBApiClient)
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
	err := client.Users().Update(id, userName, displayName, entitlements, roles)
	if err != nil {
		return err
	}
	if d.HasChange("set_admin") {
		setAdmin := d.Get("set_admin").(bool)
		if setAdmin {
			adminGroup, err := client.Groups().GetAdminGroup()
			if err != nil {
				return err
			}
			err = client.Users().SetUserAsAdmin(id, adminGroup.ID)
			if err != nil {
				return err
			}
		} else {
			adminGroup, err := client.Groups().GetAdminGroup()
			if err != nil {
				return err
			}
			err = client.Users().RemoveUserAsAdmin(id, adminGroup.ID)
			if err != nil {
				return err
			}
		}
	}
	return resourceScimUserRead(d, m)
}

func resourceScimUserDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	err := client.Users().Delete(id)
	return err
}

func isScimUserMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "urn:ietf:params:scim:api:messages:2.0:Error") &&
		strings.Contains(errorMsg, fmt.Sprintf("User with id %s not found.", resourceID)) &&
		strings.Contains(errorMsg, "404")
}

func sliceContains(value string, list []string) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}
