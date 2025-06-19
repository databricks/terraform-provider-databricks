package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMwsNccPrivateEndpointRule() common.Resource {
	s := common.StructToSchema(settings.NccPrivateEndpointRule{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		for _, p := range []string{"network_connectivity_config_id", "group_id", "resource_id"} {
			common.CustomizeSchemaPath(m, p).SetRequired().SetForceNew()
		}
		for _, p := range []string{"rule_id", "endpoint_name", "connection_state", "creation_time", "updated_time"} {
			common.CustomizeSchemaPath(m, p).SetComputed()
		}
		conflicts := []string{"group_id", "resource_names", "domain_names"}
		for _, p := range conflicts {
			common.CustomizeSchemaPath(m, p).SetOptional().SetConflictsWith(conflicts)
		}
		return m
	})
	p := common.NewPairSeparatedID("network_connectivity_config_id", "rule_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create settings.CreatePrivateEndpointRuleRequest
			common.DataToStructPointer(d, s, &create.PrivateEndpointRule)
			create.NetworkConnectivityConfigId = d.Get("network_connectivity_config_id").(string)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			rule, err := acc.NetworkConnectivity.CreatePrivateEndpointRule(ctx, create)
			if err != nil {
				return err
			}
			common.StructToData(rule, s, d)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			nccId, ruleId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			rule, err := acc.NetworkConnectivity.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
			if err != nil {
				return err
			}
			return common.StructToData(rule, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			nccId, ruleId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			_, err = acc.NetworkConnectivity.DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
			return err
		},
	}
}
