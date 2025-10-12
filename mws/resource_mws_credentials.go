package mws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// NewCredentialsAPI creates MWSCredentialsAPI instance from provider meta
func NewCredentialsAPI(ctx context.Context, m any) CredentialsAPI {
	return CredentialsAPI{m.(*common.DatabricksClient), ctx}
}

// CredentialsAPI exposes the mws credentials API
type CredentialsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// List lists all the available credentials object in the mws account
func (a CredentialsAPI) List(mwsAcctID string) ([]Credentials, error) {
	var mwsCredsList []Credentials
	credentialsAPIPath := fmt.Sprintf("/accounts/%s/credentials", mwsAcctID)
	err := a.client.Get(a.context, credentialsAPIPath, nil, &mwsCredsList)
	return mwsCredsList, err
}

type CredentialInfo struct {
	common.Namespace
	// The account id - this is for backwards compatiblity
	AccountId string `json:"account_id,omitempty" tf:"force_new,suppress_diff"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name" tf:"force_new"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn string `json:"role_arn" tf:"force_new"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime int64 `json:"creation_time,omitempty" tf:"computed"`
	// Databricks credential configuration ID.
	CredentialsId string `json:"credentials_id,omitempty" tf:"computed"`
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId string `json:"external_id,omitempty" tf:"computed"`
}

func ResourceMwsCredentials() common.Resource {
	p := common.NewPairSeparatedID("account_id", "credentials_id", "/")
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClientWithAccountIdFromConfig(d)
			if err != nil {
				return err
			}
			roleArn := d.Get("role_arn").(string)
			credentialsName := d.Get("credentials_name").(string)

			credentials, err := acc.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
				CredentialsName: credentialsName,
				AwsCredentials: provisioning.CreateCredentialAwsCredentials{
					StsRole: &provisioning.CreateCredentialStsRole{
						RoleArn: roleArn,
					},
				},
			})
			if err != nil {
				return err
			}
			d.Set("credentials_id", credentials.CredentialsId)
			d.Set("account_id", c.Config.AccountID)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			_, credsId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			credentials, err := acc.Credentials.GetByCredentialsId(ctx, credsId)
			if err != nil {
				return err
			}
			d.Set("credentials_name", credentials.CredentialsName)
			d.Set("role_arn", credentials.AwsCredentials.StsRole.RoleArn)
			d.Set("creation_time", credentials.CreationTime)
			return d.Set("external_id", credentials.AwsCredentials.StsRole.ExternalId)
		},
		// this resource cannot be updated, add this to prevent "doesn't support update" error from TF
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			_, credsId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			return acc.Credentials.DeleteByCredentialsId(ctx, credsId)
		},
		Schema: common.StructToSchema(CredentialInfo{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
			// nolint
			s["account_id"].Deprecated = "`account_id` should be set as part of the Databricks Config, not in the resource."
			return s
		}),
	}
}
