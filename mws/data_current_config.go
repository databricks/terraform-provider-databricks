package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type currentConfig struct {
	Cloud     string `json:"cloud,omitempty" tf:"optional"`
	IsAccount bool   `json:"is_account,omitempty" tf:"computed"`
	AccountId string `json:"account_id,omitempty" tf:"computed"`
	Host      string `json:"host,omitempty" tf:"computed"`
	CloudType string `json:"cloud_type,omitempty" tf:"computed"`
	AuthType  string `json:"auth_type,omitempty" tf:"computed"`
}

var validCloudValues = []string{"aws", "azure", "gcp"}

func cloudTypeFromConfig(cfg *config.Config) string {
	if cfg.IsAws() {
		return "aws"
	} else if cfg.IsAzure() {
		return "azure"
	} else if cfg.IsGcp() {
		return "gcp"
	}
	return "unknown"
}

func DataSourceCurrentConfiguration() common.Resource {
	r := common.DataResource(currentConfig{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*currentConfig)
		data.IsAccount = false
		if c.Config.HostType() == config.AccountHost {
			data.AccountId = c.Config.AccountID
			data.IsAccount = true
		}
		data.Host = c.Config.Host
		if data.Cloud != "" {
			data.CloudType = data.Cloud
		} else {
			data.CloudType = cloudTypeFromConfig(c.Config)
		}
		data.AuthType = c.Config.AuthType
		return nil
	})
	r.Schema["cloud"].ValidateFunc = validation.StringInSlice(validCloudValues, false)
	return r
}
