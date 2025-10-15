package mws

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
)

type currentConfig struct {
	common.Namespace
	IsAccount bool   `json:"is_account,omitempty" tf:"computed"`
	AccountId string `json:"account_id,omitempty" tf:"computed"`
	Host      string `json:"host,omitempty" tf:"computed"`
	CloudType string `json:"cloud_type,omitempty" tf:"computed"`
	AuthType  string `json:"auth_type,omitempty" tf:"computed"`
}

func DataSourceCurrentConfiguration() common.Resource {
	return common.DataResource(currentConfig{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*currentConfig)
		data.IsAccount = false
		if c.Config.IsAccountClient() {
			data.AccountId = c.Config.AccountID
			data.IsAccount = true
		}
		data.Host = c.Config.Host
		if c.Config.IsAws() {
			data.CloudType = "aws"
		} else if c.Config.IsAzure() {
			data.CloudType = "azure"
		} else if c.Config.IsGcp() {
			data.CloudType = "gcp"
		} else {
			data.CloudType = "unknown"
		}
		data.AuthType = c.Config.AuthType
		return nil
	})
}
