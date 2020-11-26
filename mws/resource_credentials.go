package mws

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"log"
)

// NewCredentialsAPI creates MWSCredentialsAPI instance from provider meta
func NewCredentialsAPI(m interface{}) CredentialsAPI {
	return CredentialsAPI{m.(*common.DatabricksClient), context.TODO()}
}

// CredentialsAPI exposes the mws credentials API
type CredentialsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// TODO: move mwsAcctID into provider configuration...

// Create creates a set of MWS Credentials for the cross account role
func (a CredentialsAPI) Create(mwsAcctID, credentialsName string, roleArn string) (Credentials, error) {
	var mwsCreds Credentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctID)
	err := a.client.Post(a.context, credentialsAPIPath, Credentials{
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
func (a CredentialsAPI) Read(mwsAcctID, credentialsID string) (Credentials, error) {
	var mwsCreds Credentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctID, credentialsID)
	err := a.client.Get(a.context, credentialsAPIPath, nil, &mwsCreds)
	return mwsCreds, err
}

// Delete deletes the credentials object given a credentials id
func (a CredentialsAPI) Delete(mwsAcctID, credentialsID string) error {
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials/%s", mwsAcctID, credentialsID)
	return a.client.Delete(a.context, credentialsAPIPath, nil)
}

// List lists all the available credentials object in the mws account
func (a CredentialsAPI) List(mwsAcctID string) ([]Credentials, error) {
	var mwsCredsList []Credentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctID)
	err := a.client.Get(a.context, credentialsAPIPath, nil, &mwsCredsList)
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
	credentials, err := NewCredentialsAPI(m).Create(mwsAcctID, credentialsName, roleArn)
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
	credentials, err := NewCredentialsAPI(m).Read(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
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
	err = NewCredentialsAPI(m).Delete(packagedMwsID.MwsAcctID, packagedMwsID.ResourceID)
	return err
}
