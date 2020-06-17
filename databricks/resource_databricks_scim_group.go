package databricks

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceScimGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceScimGroupCreate,
		Update: resourceScimGroupUpdate,
		Read:   resourceScimGroupRead,
		Delete: resourceScimGroupDelete,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:     schema.TypeSet,
				Optional: true,
				//Computed: true,
				//ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{Type: schema.TypeString},
				Set:  schema.HashString,
			},
			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
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
			"inherited_roles": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"entitlements": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

//func convertInterfaceSliceToStringSlice(input []interface{}) []string {
//	resp := []string{}
//	for _, item := range input {
//		resp = append(resp, item.(string))
//	}
//	return resp
//}

func resourceScimGroupCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	groupName := d.Get("display_name").(string)
	var members []string

	if rMembers, ok := d.GetOk("members"); ok {
		members = convertInterfaceSliceToStringSlice(rMembers.(*schema.Set).List())
	}
	var roles []string

	if rRoles, ok := d.GetOk("roles"); ok {
		roles = convertInterfaceSliceToStringSlice(rRoles.(*schema.Set).List())
	}

	var entitlements []string

	if rEntitlements, ok := d.GetOk("entitlements"); ok {
		entitlements = convertInterfaceSliceToStringSlice(rEntitlements.(*schema.Set).List())
	}

	group, err := client.Groups().Create(groupName, members, roles, entitlements)
	if err != nil {
		return err
	}
	d.SetId(group.ID)
	return resourceScimGroupRead(d, m)
}

func getListOfMemberRefs(memberList []model.GroupMember) []string {
	resp := []string{}
	for _, member := range memberList {
		resp = append(resp, member.Value)
	}
	log.Println("Members list =")
	log.Println(resp)
	return resp
}

func getListOfEntitlements(entitlementList []model.EntitlementsListItem) []string {
	resp := []string{}
	for _, entitlement := range entitlementList {
		resp = append(resp, string(entitlement.Value))
	}
	return resp
}

func resourceScimGroupRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	group, err := client.Groups().Read(id)
	if err != nil {
		if isScimGroupMissing(err.Error(), id) {
			log.Printf("Missing scim group with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}

	err = d.Set("display_name", group.DisplayName)
	if err != nil {
		return err
	}

	err = d.Set("members", getListOfMemberRefs(group.Members))
	if err != nil {
		return err
	}

	err = d.Set("roles", getListOfRoles(group.Roles))
	if err != nil {
		return err
	}
	err = d.Set("inherited_roles", getListOfRoles(group.InheritedRoles))
	if err != nil {
		return err
	}

	err = d.Set("entitlements", getListOfEntitlements(group.Entitlements))
	if err != nil {
		return err
	}

	return err
}

func diff(sliceA []string, sliceB []string) []string {
	var output []string
	m := make(map[string]int)
	for _, y := range sliceB {
		m[y]++
	}
	for _, x := range sliceA {
		if m[x] > 0 {
			m[x]--
			continue
		}
		output = append(output, x)
	}
	return output
}

func resourceScimGroupUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)

	group, err := client.Groups().Read(id)
	if err != nil {
		return err
	}

	// In the future we may want these to be in one patch statement as opposed to 3 separate patches to make a
	// better all or none scenario
	if d.HasChange("members") {
		oldMembersInterface, newMembersInterface := d.GetChange("members")
		oldMembers := convertInterfaceSliceToStringSlice(oldMembersInterface.(*schema.Set).List())
		newMembers := convertInterfaceSliceToStringSlice(newMembersInterface.(*schema.Set).List())
		addMembers := diff(newMembers, oldMembers)
		removeMembers := diff(oldMembers, newMembers)
		log.Println("add members")
		log.Println(addMembers)
		log.Println("remove members")
		log.Println(removeMembers)
		// We need to check if the group has any members especially if addMembers is 0 and removeMembers is not 0
		// A user might be deleted and a patch operation on an empty group will fail
		if len(group.Members) > 0 || len(addMembers) > 0 {
			err = client.Groups().Patch(group.ID, addMembers, removeMembers, model.GroupMembersPath)
			if err != nil {
				return err
			}
		}
	}

	if d.HasChange("roles") {
		oldRolesInterface, newRolesInterface := d.GetChange("roles")
		oldRoles := convertInterfaceSliceToStringSlice(oldRolesInterface.(*schema.Set).List())
		newRoles := convertInterfaceSliceToStringSlice(newRolesInterface.(*schema.Set).List())
		addRoles := diff(newRoles, oldRoles)
		removeRoles := diff(oldRoles, newRoles)
		log.Println("add roles")
		log.Println(addRoles)
		log.Println("remove roles")
		log.Println(removeRoles)
		err = client.Groups().Patch(group.ID, addRoles, removeRoles, model.GroupRolesPath)
		if err != nil {
			return err
		}
	}

	if d.HasChange("entitlements") {
		oldEntitlementsInterface, newEntitlementsInterface := d.GetChange("entitlements")
		oldEntitlements := convertInterfaceSliceToStringSlice(oldEntitlementsInterface.(*schema.Set).List())
		newEntitlements := convertInterfaceSliceToStringSlice(newEntitlementsInterface.(*schema.Set).List())
		addEntitlements := diff(newEntitlements, oldEntitlements)
		removeEntitlements := diff(oldEntitlements, newEntitlements)
		log.Println("add entitlements")
		log.Println(addEntitlements)
		log.Println("remove entitlements")
		log.Println(removeEntitlements)
		err = client.Groups().Patch(group.ID, addEntitlements, removeEntitlements, model.GroupEntitlementsPath)
		if err != nil {
			return err
		}
	}

	return resourceScimGroupRead(d, m)
}

func resourceScimGroupDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	err := client.Groups().Delete(id)
	return err
}

func isScimGroupMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, "urn:ietf:params:scim:api:messages:2.0:Error") &&
		strings.Contains(errorMsg, fmt.Sprintf("Group with id %s not found.", resourceID)) &&
		strings.Contains(errorMsg, "404")
}
