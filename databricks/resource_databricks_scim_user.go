package databricks

import (
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"sort"
	"strings"
)

func resourceScimUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceScimUserCreate,
		Update: resourceScimUserUpdate,
		Read:   resourceScimUserRead,
		Delete: resourceScimUserDelete,

		Schema: map[string]*schema.Schema{
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"roles": &schema.Schema{
				Type:       schema.TypeSet,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem:       &schema.Schema{Type: schema.TypeString},
				Set:        schema.HashString,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					inheritedRoles := convertListInterfaceToString(d.Get("inherited_roles").(*schema.Set).List())
					if new == "0" {
						return true
					}
					if new == "" {
						for _, role := range inheritedRoles {
							if old == role {
								return true
							}
						}
					}
					return false
				},
			},
			"entitlements": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"inherited_roles": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"set_admin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Set:      schema.HashString,
			},
		},
		//CustomizeDiff: customdiff.IfValueChange("roles", func(old, new, meta interface{}) bool {
		//	return true
		//}, func(resourceDiff *schema.ResourceDiff, i interface{}) error {
		//	return fmt.Errorf("Failed")
		//}),
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
	client := m.(service.DBApiClient)
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
func getListOfEntitlements(entitlementList []model.EntitlementsListItem) []string {
	resp := []string{}
	for _, entitlement := range entitlementList {
		resp = append(resp, string(entitlement.Value))
	}
	return resp
}

func resourceScimUserRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
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

	roles := getListOfRoles(user.Roles)
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

	inheritedRoles := getListOfRoles(user.InheritedRoles)
	err = d.Set("inherited_roles", inheritedRoles)
	if err != nil {
		return err
	}

	err = d.Set("roles", roles)
	if err != nil {
		return err
	}

	return err
}

func resourceScimUserUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
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
	client := m.(service.DBApiClient)
	err := client.Users().Delete(id)
	return err
}

func isScimUserMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "urn:ietf:params:scim:api:messages:2.0:Error") &&
		strings.Contains(errorMsg, fmt.Sprintf("User with id %s not found.", resourceID)) &&
		strings.Contains(errorMsg, "404")
}
