package db

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
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"entitlements": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
	return resourceScimUserRead(d, m)
}

func resourceScimUserDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
	err := client.Users().Delete(id)
	return err
}

func isScimUserMissing(errorMsg, resourceId string) bool {
	return strings.Contains(errorMsg, "urn:ietf:params:scim:api:messages:2.0:Error") &&
		strings.Contains(errorMsg, fmt.Sprintf("User with id %s not found.", resourceId)) &&
		strings.Contains(errorMsg, "404")
}
