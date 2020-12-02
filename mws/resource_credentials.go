package mws

import (
	"context"
	"fmt"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewCredentialsAPI creates MWSCredentialsAPI instance from provider meta
func NewCredentialsAPI(ctx context.Context, m interface{}) CredentialsAPI {
	return CredentialsAPI{m.(*common.DatabricksClient), ctx}
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

// ResourceCredentials ...
func ResourceCredentials() *schema.Resource {
	p := util.NewPairSeparatedID("account_id", "credentials_id", "/")
	return util.CommonResource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID := d.Get("account_id").(string)
			roleArn := d.Get("role_arn").(string)
			credentialsName := d.Get("credentials_name").(string)
			credentials, err := NewCredentialsAPI(ctx, c).Create(accountID, credentialsName, roleArn)
			if err != nil {
				return err
			}
			d.Set("credentials_id", credentials.CredentialsID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, credsID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			credentials, err := NewCredentialsAPI(ctx, c).Read(accountID, credsID)
			if err != nil {
				return err
			}
			d.Set("credentials_name", credentials.CredentialsName)
			d.Set("role_arn", credentials.AwsCredentials.StsRole.RoleArn)
			d.Set("creation_time", credentials.CreationTime)
			return d.Set("external_id", credentials.AwsCredentials.StsRole.ExternalID)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			accountID, credsID, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return NewCredentialsAPI(ctx, c).Delete(accountID, credsID)
		},
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
	}.ToResource()
}