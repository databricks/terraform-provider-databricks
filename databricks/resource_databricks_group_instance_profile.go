package databricks

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGroupInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupInstanceProfileCreate,
		Read:   resourceGroupInstanceProfileRead,
		Delete: resourceGroupInstanceProfileDelete,

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

func resourceGroupInstanceProfileCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	groupID := d.Get("group_id").(string)
	instanceProfileID := d.Get("instance_profile_id").(string)
	groupInstanceProfileID := &GroupInstanceProfileID{
		GroupID:           groupID,
		InstanceProfileID: instanceProfileID,
	}

	roleAddList := []string{groupInstanceProfileID.InstanceProfileID}
	err := client.Groups().Patch(groupInstanceProfileID.GroupID, roleAddList, nil, model.GroupRolesPath)
	if err != nil {
		return err
	}

	d.SetId(groupInstanceProfileID.String())
	return resourceGroupInstanceProfileRead(d, m)
}

func resourceGroupInstanceProfileRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	groupInstanceProfileID := parseGroupInstanceProfileID(id)
	group, err := client.Groups().Read(groupInstanceProfileID.GroupID)

	// First verify if the group exists
	if err != nil {
		if isScimGroupMissing(err.Error(), groupInstanceProfileID.GroupID) {
			log.Printf("Missing  group with id: %s.", groupInstanceProfileID.GroupID)
			d.SetId("")
			return nil
		}
		return err
	}

	// Set Id to null if instance profile is not in group
	if !iInstanceProfileInGroup(groupInstanceProfileID.InstanceProfileID, &group) {
		log.Printf("Missing role %s in group with id: %s.", groupInstanceProfileID.InstanceProfileID, groupInstanceProfileID.GroupID)
		d.SetId("")
		return nil
	}

	err = d.Set("group_id", groupInstanceProfileID.GroupID)
	if err != nil {
		return err
	}

	err = d.Set("instance_profile_id", groupInstanceProfileID.InstanceProfileID)
	return err
}

func resourceGroupInstanceProfileDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	groupInstanceProfileID := parseGroupInstanceProfileID(id)

	roleRemoveList := []string{groupInstanceProfileID.InstanceProfileID}
	// Patch op to remove role from group
	err := client.Groups().Patch(groupInstanceProfileID.GroupID, nil, roleRemoveList, model.GroupRolesPath)
	return err
}

type GroupInstanceProfileID struct {
	GroupID           string
	InstanceProfileID string
}

func (g GroupInstanceProfileID) String() string {
	return fmt.Sprintf("%s|%s", g.GroupID, g.InstanceProfileID)
}

func parseGroupInstanceProfileID(id string) *GroupInstanceProfileID {
	parts := strings.Split(id, "|")
	return &GroupInstanceProfileID{
		GroupID:           parts[0],
		InstanceProfileID: parts[1],
	}
}

func iInstanceProfileInGroup(role string, group *model.Group) bool {
	for _, groupRole := range group.Roles {
		if groupRole.Value == role {
			return true
		}
	}
	return false
}
