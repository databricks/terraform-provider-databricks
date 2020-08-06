package databricks

import (
	"log"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceProfileCreate,
		Read:   resourceInstanceProfileRead,
		Delete: resourceInstanceProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
	client := m.(*service.DatabricksClient)
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
	client := m.(*service.DatabricksClient)
	profile, err := client.InstanceProfiles().Read(id)
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
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
	client := m.(*service.DatabricksClient)
	err := client.InstanceProfiles().Delete(id)
	return err
}
