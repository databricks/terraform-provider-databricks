package mws

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type currentConfig struct {
	IsAccount bool   `json:"is_account,omitempty" tf:"computed"`
	AccountId string `json:"account_id,omitempty" tf:"computed"`
	Host      string `json:"host,omitempty" tf:"computed"`
	CloudType string `json:"cloud_type,omitempty" tf:"computed"`
	AuthType  string `json:"auth_type,omitempty" tf:"computed"`
}

func DataSourceCurrentConfiguration() common.Resource {
	s := common.StructToSchema(currentConfig{}, nil)
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var data currentConfig
			common.DataToStructPointer(d, s, &data)
			data.IsAccount = false
			if newClient.Config.IsAccountClient() {
				data.AccountId = newClient.Config.AccountID
				data.IsAccount = true
			}
			data.Host = newClient.Config.Host
			if newClient.Config.IsAws() {
				data.CloudType = "aws"
			} else if newClient.Config.IsAzure() {
				data.CloudType = "azure"
			} else if newClient.Config.IsGcp() {
				data.CloudType = "gcp"
			} else {
				data.CloudType = "unknown"
			}
			data.AuthType = newClient.Config.AuthType
			err = common.StructToData(data, s, d)
			if err != nil {
				return err
			}
			d.SetId("_")
			return nil
		},
	}
}
