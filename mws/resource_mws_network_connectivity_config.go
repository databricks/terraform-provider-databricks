package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceMwsNetworkConnectivityConfig() common.Resource {
	s := common.StructToSchema(settings.NetworkConnectivityConfiguration{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.MustSchemaPath(m, "name").ValidateFunc = validation.StringLenBetween(3, 30)
		for _, p := range []string{"name", "region"} {
			common.MustSchemaPath(m, p).Optional = false
			common.MustSchemaPath(m, p).Required = true
		}
		for _, p := range []string{"account_id", "network_connectivity_config_id", "creation_time", "updated_time"} {
			common.MustSchemaPath(m, p).Optional = false
			common.MustSchemaPath(m, p).Required = false
			common.MustSchemaPath(m, p).Computed = true
		}
		return m
	})
	p := common.NewPairSeparatedID("account_id", "network_connectivity_config_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create settings.CreateNetworkConnectivityConfigRequest
			common.DataToStructPointer(d, s, &create)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			ncc, err := acc.NetworkConnectivity.CreateNetworkConnectivityConfiguration(ctx, create)
			if err != nil {
				return err
			}
			common.StructToData(ncc, s, d)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, nccId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			ncc, err := acc.NetworkConnectivity.GetNetworkConnectivityConfigurationByNetworkConnectivityConfigId(ctx, nccId)
			if err != nil {
				return err
			}
			return common.StructToData(ncc, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, nccId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			return acc.NetworkConnectivity.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(ctx, nccId)
		},
	}
}
