package db

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stikkireddy/databricks-tf-provider/client/model"
	"github.com/stikkireddy/databricks-tf-provider/client/service"
	"log"
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
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"entitlements": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				//Computed: true,
				Elem: &schema.Schema{Type: schema.TypeString},
				Set:  schema.HashString,
			},
		},
	}
}

func convertInterfaceSliceToStringSlice(input []interface{}) []string {
	resp := []string{}
	for _, item := range input {
		resp = append(resp, item.(string))
	}
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
	user, err := client.Users().Create(userName, displayName, roles, entitlements)
	if err != nil {
		return err
	}
	d.SetId(user.ID)
	err = d.Set("entitlements", entitlements)
	if err != nil {
		return err
	}
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
		return err
	}

	roles := getListOfRoles(user.Roles)
	//entitlements := getListOfEntitlements(user.Entitlements)

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
	err := client.Users().Update(id, userName, displayName, roles, entitlements)
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
