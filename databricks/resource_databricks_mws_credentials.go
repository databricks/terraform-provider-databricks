package databricks

import (
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
	"strings"
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
	client := m.(*service.DBApiClient)
	credentialsName := d.Get("credentials_name").(string)
	roleArn := d.Get("role_arn").(string)
	mwsAcctId := d.Get("account_id").(string)
	credentials, err := client.MWSCredentials().Create(mwsAcctId, credentialsName, roleArn)
	if err != nil {
		return err
	}
	credentialsResourceId := PackagedMWSIds{
		MwsAcctId:  mwsAcctId,
		ResourceId: credentials.CredentialsID,
	}
	d.SetId(packMWSAccountId(credentialsResourceId))
	return resourceMWSCredentialsRead(d, m)
}

func resourceMWSCredentialsRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	client := m.(*service.DBApiClient)
	packagedMwsId, err := unpackMWSAccountId(id)
	if err != nil {
		return err
	}
	credentials, err := client.MWSCredentials().Read(packagedMwsId.MwsAcctId, packagedMwsId.ResourceId)
	if err != nil {
		if isMWSCredentialsMissing(err.Error()) {
			log.Printf("Missing e2 credentials with id: %s.", packagedMwsId.ResourceId)
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
	client := m.(*service.DBApiClient)
	packagedMwsId, err := unpackMWSAccountId(id)
	if err != nil {
		return err
	}
	err = client.MWSCredentials().Delete(packagedMwsId.MwsAcctId, packagedMwsId.ResourceId)
	return err
}

func isMWSCredentialsMissing(errorMsg string) bool {
	return strings.Contains(errorMsg, "RESOURCE_DOES_NOT_EXIST")
}
