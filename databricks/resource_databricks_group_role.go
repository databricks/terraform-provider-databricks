package databricks

import (
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceGroupRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupRoleCreate,
		Read:   resourceGroupRoleRead,
		Delete: resourceGroupRoleDelete,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"instance_profile_id": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: ValidateInstanceProfileARN,
			},
		},
	}
}



func resourceGroupRoleCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	groupID := d.Get("group_id").(string)
	instanceProfileID := d.Get("instance_profile_id").(string)
	groupRoleID := &GroupRoleID{
		GroupID: groupID,
		InstanceProfileID: instanceProfileID,
	}

	roleAddList := []string{groupRoleID.InstanceProfileID}
	err := client.Groups().Patch(groupRoleID.GroupID, roleAddList, nil, model.GroupRolesPath)
	if err != nil {
		return err
	}

	d.SetId(groupRoleID.String())
	return resourceGroupRoleRead(d, m)
}

func resourceGroupRoleRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	groupRoleID := parseGroupRoleID(id)
	group, err := client.Groups().Read(groupRoleID.GroupID)

	// First verify if the group exists
	if err != nil {
		if isScimGroupMissing(err.Error(), groupRoleID.GroupID) {
			log.Printf("Missing  group with id: %s.", groupRoleID.GroupID)
			d.SetId("")
			return nil
		}
		return err
	}

	// Set Id to null if instance profile is not in group
	if !iRoleInGroup(groupRoleID.InstanceProfileID, &group) {
		log.Printf("Missing role %s in group with id: %s.", groupRoleID.InstanceProfileID, groupRoleID.GroupID)
		d.SetId("")
		return nil
	}

	err = d.Set("group_id", groupRoleID.GroupID)
	if err != nil {
		return err
	}

	err = d.Set("instance_profile_id", groupRoleID.InstanceProfileID)
	return err
}

func resourceGroupRoleDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	groupRoleID := parseGroupRoleID(id)

	roleRemoveList := []string{groupRoleID.InstanceProfileID}
	// Patch op to remove role from group
	err := client.Groups().Patch(groupRoleID.GroupID, nil, roleRemoveList, model.GroupRolesPath)
	return err
}


type GroupRoleID struct {
	GroupID string
	InstanceProfileID string
}

func (g GroupRoleID) String() string {
	return fmt.Sprintf("%s|%s", g.GroupID, g.InstanceProfileID)
}

func parseGroupRoleID(id string) *GroupRoleID {
	parts := strings.Split(id, "|")
	return &GroupRoleID{
		GroupID:parts[0],
		InstanceProfileID:parts[1],
	}
}

func iRoleInGroup(role string, group *model.Group) bool {
	for _, groupRole := range(group.Roles) {
		if groupRole.Value == role {
			return true
		}
	}
	return false
}