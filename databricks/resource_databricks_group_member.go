package databricks

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Use this to manage individual members of a particular group
func resourceGroupMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupMemberCreate,
		Read:   resourceGroupMemberRead,
		Delete: resourceGroupMemberDelete,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"member_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceGroupMemberCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	groupID := d.Get("group_id").(string)
	memberID := d.Get("member_id").(string)

	groupMemberID := &GroupMemberID{
		GroupID:  groupID,
		MemberID: memberID,
	}

	memberAddList := []string{groupMemberID.MemberID}
	err := client.Groups().Patch(groupMemberID.GroupID, memberAddList, nil, model.GroupMembersPath)
	if err != nil {
		return err
	}

	d.SetId(groupMemberID.String())
	return resourceGroupMemberRead(d, m)
}

func resourceGroupMemberRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	groupMemberID := parseGroupMemberID(id)
	group, err := client.Groups().Read(groupMemberID.GroupID)

	// First verify if the group exists
	if err != nil {
		if isScimGroupMissing(err.Error(), id) {
			log.Printf("Missing group with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}

	// Set Id to null if instance profile is not in group
	if !iMemberInGroup(groupMemberID.MemberID, &group) {
		log.Printf("Missing member %s in group with id: %s.", groupMemberID.MemberID, groupMemberID.GroupID)
		d.SetId("")
		return nil
	}

	err = d.Set("group_id", groupMemberID.GroupID)
	if err != nil {
		return err
	}

	err = d.Set("member_id", groupMemberID.MemberID)
	return err
}

func resourceGroupMemberDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	groupMemberID := parseGroupMemberID(id)

	memberRemoveList := []string{groupMemberID.MemberID}
	// Patch op to remove member from group
	err := client.Groups().Patch(groupMemberID.GroupID, nil, memberRemoveList, model.GroupMembersPath)
	return err
}

type GroupMemberID struct {
	GroupID  string
	MemberID string
}

func (g GroupMemberID) String() string {
	return fmt.Sprintf("%s|%s", g.GroupID, g.MemberID)
}

func parseGroupMemberID(id string) *GroupMemberID {
	parts := strings.Split(id, "|")
	return &GroupMemberID{
		GroupID:  parts[0],
		MemberID: parts[1],
	}
}

func iMemberInGroup(member string, group *model.Group) bool {
	for _, groupMember := range group.Members {
		if groupMember.Value == member {
			return true
		}
	}
	return false
}
