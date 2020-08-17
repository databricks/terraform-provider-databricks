package identity

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Use this to manage individual members of a particular group
func ResourceGroupMember() *schema.Resource {
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
			// TODO: add group_name
			// TODO: add user_name
			"member_id": {
				// is a user or a group - we do not know
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceGroupMemberCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	groupID := d.Get("group_id").(string)
	memberID := d.Get("member_id").(string)

	groupMemberID := &groupMemberID{
		GroupID:  groupID,
		MemberID: memberID,
	}

	memberAddList := []string{groupMemberID.MemberID}
	err := NewGroupsAPI(client).Patch(groupMemberID.GroupID, memberAddList, nil, GroupMembersPath)
	if err != nil {
		return err
	}

	d.SetId(groupMemberID.String())
	return resourceGroupMemberRead(d, m)
}

func resourceGroupMemberRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	groupMemberID := parsegroupMemberID(id)
	group, err := NewGroupsAPI(client).Read(groupMemberID.GroupID)

	// First verify if the group exists
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
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
	client := m.(*common.DatabricksClient)
	groupMemberID := parsegroupMemberID(id)

	memberRemoveList := []string{groupMemberID.MemberID}
	// Patch op to remove member from group
	err := NewGroupsAPI(client).Patch(groupMemberID.GroupID, nil, memberRemoveList, GroupMembersPath)
	return err
}

type groupMemberID struct {
	GroupID  string
	MemberID string
}

func (g groupMemberID) String() string {
	return fmt.Sprintf("%s|%s", g.GroupID, g.MemberID)
}

func parsegroupMemberID(id string) *groupMemberID {
	parts := strings.Split(id, "|")
	return &groupMemberID{
		GroupID:  parts[0],
		MemberID: parts[1],
	}
}

func iMemberInGroup(member string, group *Group) bool {
	for _, groupMember := range group.Members {
		if groupMember.Value == member {
			return true
		}
	}
	return false
}
