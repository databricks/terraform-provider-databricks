package databricks

import (
	"fmt"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceProfileCreate,
		Read:   resourceInstanceProfileRead,
		Delete: resourceInstanceProfileDelete,

		Schema: map[string]*schema.Schema{
			"instance_profile_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"skip_validation": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceInstanceProfileCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DBApiClient)
	instanceProfileArn := d.Get("instance_profile_arn").(string)
	skipValidation := d.Get("skip_validation").(bool)
	err := client.InstanceProfiles().Create(instanceProfileArn, skipValidation)
	if err != nil {
		return err
	}
	d.SetId(instanceProfileArn)
	err = d.Set("skip_validation", skipValidation)
	if err != nil {
		return err
	}

	return resourceInstanceProfileRead(d, m)
}

func resourceInstanceProfileRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	profile, err := client.InstanceProfiles().Read(id)
	if err != nil {
		if isInstanceProfileMissing(err.Error(), id) {
			log.Printf("Missing instance profile with id: %s.", id)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("instance_profile_arn", profile)
	return err
}

func resourceInstanceProfileDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	err := client.InstanceProfiles().Delete(id)
	return err
}

func isInstanceProfileMissing(errorMsg, resourceID string) bool {
	return strings.Contains(errorMsg, fmt.Sprintf("Instance profile with name: %s not found in "+
		"list of instance profiles in the workspace!", resourceID))
}
