package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	s := common.StructToSchema(currentConfig{}, func(
		s map[string]*schema.Schema) map[string]*schema.Schema {
		s["cloud"].ValidateFunc = validation.StringInSlice(validCloudValues, false)
		common.AddApiField(s)
		return s
	})
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		IsDual: true,
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			if err := common.ValidateApiLevelForUnifiedHostFromData(d, m); err != nil {
				return err
			}
			newClient, err := m.DatabricksClientForDualResource(ctx, d)
			if err != nil {
				return err
			}
			var data currentConfig
			common.DataToStructPointer(d, s, &data)
			data.IsAccount = false
			if common.IsAccountLevel(d, newClient) {
				data.AccountId = newClient.Config.AccountID
				data.IsAccount = true
			}
			data.Host = newClient.Config.Host
			if data.Cloud != "" {
				data.CloudType = data.Cloud
			} else {
				data.CloudType = cloudTypeFromConfig(newClient.Config)
			}
			data.AuthType = newClient.Config.AuthType
			common.StructToData(&data, s, d)
			d.SetId("_")
			return nil
		},
	}
}
