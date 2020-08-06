package databricks

import (
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceMWSCredentials() *schema.Resource {
	return &schema.Resource{
		Create: resourceMWSCredentialsCreate,
		Read:   resourceMWSCredentialsRead,
		Delete: resourceMWSCredentialsDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"credentials_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMWSCredentialsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*service.DatabricksClient)
	credentialsName := d.Get("credentials_name").(string)
	roleArn := d.Get("role_arn").(string)
	mwsAcctID := d.Get("account_id").(string)
	credentials, err := client.MWSCredentials().Create(mwsAcctID, credentialsName, roleArn)
	if err != nil {
		return err
	}
	credentialsResourceID := PackagedMWSIds{
		MwsAcctID:  mwsAcctID,
		ResourceID: credentials.CredentialsID,
	}
	d.SetId(packMWSAccountID(credentialsResourceID))
	return resourceMWSCredentialsRead(d, m)
}

func resourceMWSCredentialsRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	credentials, err := client.MWSCredentials().Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		if e, ok := err.(service.APIError); ok && e.IsMissing() {
			log.Printf("missing resource due to error: %v\n", e)
			d.SetId("")
			return nil
		}
		return err
	}
	err = d.Set("credentials_name", credentials.CredentialsName)
	if err != nil {
		return err
	}
	err = d.Set("role_arn", credentials.AwsCredentials.StsRole.RoleArn)
	if err != nil {
		return err
	}
	err = d.Set("creation_time", credentials.CreationTime)
	if err != nil {
		return err
	}
	err = d.Set("external_id", credentials.AwsCredentials.StsRole.ExternalID)
	if err != nil {
		return err
	}
	err = d.Set("credentials_id", credentials.CredentialsID)
	if err != nil {
		return err
	}
	return nil
}

func resourceMWSCredentialsDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DatabricksClient)
	packagedMwsID, err := unpackMWSAccountID(id)
	if err != nil {
		return err
	}
	err = client.MWSCredentials().Delete(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	return err
}
