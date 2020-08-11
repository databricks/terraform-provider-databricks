package mws

import (
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

// NewMWSCredentialsAPI creates MWSCredentialsAPI instance from provider meta
func NewMWSCredentialsAPI(m interface{}) MWSCredentialsAPI {
	return MWSCredentialsAPI{client: m.(*common.DatabricksClient)}
}

// MWSCredentialsAPI exposes the mws credentials API
type MWSCredentialsAPI struct {
	client *common.DatabricksClient
}

// TODO: move mwsAcctID into provider configuration...

// Create creates a set of MWS Credentials for the cross account role
func (a MWSCredentialsAPI) Create(mwsAcctID, credentialsName string, roleArn string) (MWSCredentials, error) {
	var mwsCreds MWSCredentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctID)
	err := a.client.Post(credentialsAPIPath, MWSCredentials{
		CredentialsName: credentialsName,
		AwsCredentials: &AwsCredentials{
			StsRole: &StsRole{
				RoleArn: roleArn,
			},
		},
	}, &mwsCreds)
	return mwsCreds, err
}

// Read returns the credentials object along with metadata
func (a MWSCredentialsAPI) Read(mwsAcctID, credentialsID string) (MWSCredentials, error) {
	var mwsCreds MWSCredentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctID, credentialsID)
	err := a.client.Get(credentialsAPIPath, nil, &mwsCreds)
	return mwsCreds, err
}

// Delete deletes the credentials object given a credentials id
func (a MWSCredentialsAPI) Delete(mwsAcctID, credentialsID string) error {
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctID, credentialsID)
	return a.client.Delete(credentialsAPIPath, nil)
}

// List lists all the available credentials object in the mws account
func (a MWSCredentialsAPI) List(mwsAcctID string) ([]MWSCredentials, error) {
	var mwsCredsList []MWSCredentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctID)
	err := a.client.Get(credentialsAPIPath, nil, &mwsCredsList)
	return mwsCredsList, err
}

func ResourceCredentials() *schema.Resource {
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
	credentialsName := d.Get("credentials_name").(string)
	roleArn := d.Get("role_arn").(string)
	mwsAcctID := d.Get("account_id").(string)
	credentials, err := NewMWSCredentialsAPI(m).Create(mwsAcctID, credentialsName, roleArn)
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
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	credentials, err := NewMWSCredentialsAPI(m).Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	if err != nil {
		if e, ok := err.(common.APIError); ok && e.IsMissing() {
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
	packagedMwsID, err := UnpackMWSAccountID(id)
	if err != nil {
		return err
	}
	err = NewMWSCredentialsAPI(m).Delete(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	return err
}
