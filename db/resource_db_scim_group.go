package db

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"log"
)

func resourceScimGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceScimGroupCreate,
		Update: resourceScimGroupUpdate,
		Read:   resourceScimGroupRead,
		Delete: resourceScimGroupDelete,

		Schema: map[string]*schema.Schema{
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"members": &schema.Schema{
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
	client := m.(service.DBApiClient)
	groupName := d.Get("display_name").(string)
	var members []string

	if rMembers, ok := d.GetOk("members"); ok {
		members = convertInterfaceSliceToStringSlice(rMembers.(*schema.Set).List())
		log.Println(members)
	}
	group, err := client.Groups().Create(groupName, members)
	if err != nil {
		return err
	}
	d.SetId(group.ID)
	return resourceScimGroupRead(d, m)
}

//
func getListOfMemberRefs(memberList []model.GroupMember) []string {
	resp := []string{}
	for _, member := range memberList {
		resp = append(resp, member.Value)
	}
	return resp
}

//func getListOfEntitlements(entitlementList []model.EntitlementsListItem) []string {
//	resp := []string{}
//	for _, entitlement := range entitlementList {
//		resp = append(resp, string(entitlement.Value))
//	}
//	return resp
//}

func resourceScimGroupRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
	group, err := client.Groups().Read(id)
	if err != nil {
		return err
	}

	err = d.Set("members", getListOfMemberRefs(group.Members))
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
	client := m.(service.DBApiClient)
	var currentMembers []string
	group, err := client.Groups().Read(id)
	if err != nil {
		return err
	}
	remoteMembers := getListOfMemberRefs(group.Members)
	if members, ok := d.GetOk("members"); ok {
		currentMembers = convertInterfaceSliceToStringSlice(members.(*schema.Set).List())
	}
	addMembers := diff(currentMembers, remoteMembers)
	removeMembers := diff(remoteMembers, currentMembers)
	log.Println("add members")
	log.Println(addMembers)
	log.Println("remove members")
	log.Println(removeMembers)
	err = client.Groups().Update(group.ID, addMembers, removeMembers)
	if err != nil {
		return err
	}
	return resourceScimGroupRead(d, m)
}

func resourceScimGroupDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(service.DBApiClient)
	err := client.Groups().Delete(id)
	return err
}
