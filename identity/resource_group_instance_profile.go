package identity

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGroupInstanceProfile() *schema.Resource {
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

// ValidateInstanceProfileARN is a ValidateFunc that ensures the role id is a valid aws iam instance profile arn
func ValidateInstanceProfileARN(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)

	if v == "" {
		return nil, []error{fmt.Errorf("%s is empty got: %s, must be an aws instance profile arn", key, v)}
	}

	// Parse and verify instance profiles
	instanceProfileArn, err := arn.Parse(v)
	if err != nil {
		return nil, []error{fmt.Errorf("%s is invalid got: %s received error: %w", key, v, err)}
	}
	// Verify instance profile resource type, Resource gets parsed as instance-profile/<profile-name>
	if !strings.HasPrefix(instanceProfileArn.Resource, "instance-profile") {
		return nil, []error{fmt.Errorf("%s must be an instance profile resource, got: %s in %s",
			key, instanceProfileArn.Resource, v)}
	}
	return nil, nil
}

func resourceGroupInstanceProfileCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	groupID := d.Get("group_id").(string)
	instanceProfileID := d.Get("instance_profile_id").(string)
	groupInstanceProfileID := &groupInstanceProfileID{
		GroupID:           groupID,
		InstanceProfileID: instanceProfileID,
	}

	roleAddList := []string{groupInstanceProfileID.InstanceProfileID}
	err := NewGroupsAPI(client).Patch(groupInstanceProfileID.GroupID, roleAddList, nil, GroupRolesPath)
	if err != nil {
		return err
	}

	d.SetId(groupInstanceProfileID.String())
	return resourceGroupInstanceProfileRead(d, m)
}

func resourceGroupInstanceProfileRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*common.DatabricksClient)
	groupInstanceProfileID := parsegroupInstanceProfileID(id)
	group, err := NewGroupsAPI(client).Read(groupInstanceProfileID.GroupID)

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
	if !InstanceProfileInGroup(groupInstanceProfileID.InstanceProfileID, &group) {
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
	client := m.(*common.DatabricksClient)
	groupInstanceProfileID := parsegroupInstanceProfileID(id)

	roleRemoveList := []string{groupInstanceProfileID.InstanceProfileID}
	// Patch op to remove role from group
	err := NewGroupsAPI(client).Patch(groupInstanceProfileID.GroupID, nil, roleRemoveList, GroupRolesPath)
	return err
}

type groupInstanceProfileID struct {
	GroupID           string
	InstanceProfileID string
}

func (g groupInstanceProfileID) String() string {
	return fmt.Sprintf("%s|%s", g.GroupID, g.InstanceProfileID)
}

func parsegroupInstanceProfileID(id string) *groupInstanceProfileID {
	parts := strings.Split(id, "|")
	return &groupInstanceProfileID{
		GroupID:           parts[0],
		InstanceProfileID: parts[1],
	}
}

func InstanceProfileInGroup(role string, group *Group) bool {
	for _, groupRole := range group.Roles {
		if groupRole.Value == role {
			return true
		}
	}
	return false
}
