package mws

import (
	"context"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMwsNccPrivateEndpointRule() common.Resource {
	s := common.StructToSchema(settings.NccPrivateEndpointRule{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		for _, p := range []string{"endpoint_service", "group_id", "resource_id", "domain_names", "resource_names"} {
			common.CustomizeSchemaPath(m, p).SetOptional()
		}
		for _, p := range []string{"endpoint_service", "group_id", "resource_id"} {
			common.CustomizeSchemaPath(m, p).SetForceNew()
		}

		for _, p := range []string{"rule_id", "endpoint_name", "connection_state", "creation_time", "updated_time", "vpc_endpoint_id"} {
			common.CustomizeSchemaPath(m, p).SetComputed()
		}

		common.CustomizeSchemaPath(m, "network_connectivity_config_id").SetRequired().SetForceNew()
		common.CustomizeSchemaPath(m, "enabled").SetOptional().SetComputed()

		// only one of `domain_names`, `resource_names`, `group_id` can be specified, as they are applicable
		// to different type of endpoints
		supportedFields := []string{"group_id", "resource_names", "domain_names"}
		for _, key := range supportedFields {
			conflicts := make([]string, 0, len(supportedFields)-1)
			for _, otherKey := range supportedFields {
				if key != otherKey {
					conflicts = append(conflicts, otherKey)
				}
			}
			common.CustomizeSchemaPath(m, key).SetConflictsWith(conflicts)
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
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			nccId, ruleId, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}

			// only enabled, domain names & resource names are updatable
			// they do require update_mask to be set
			// resource_names are not applicable to Azure, so we exclude them from the update
			updateMask := []string{"enabled"}
			updatePrivateEndpointRule := settings.UpdatePrivateEndpointRule{
				Enabled: d.Get("enabled").(bool),
			}

			if d.HasChange("domain_names") {
				updateMask = append(updateMask, "domain_names")
				newDomainNames := []string{}
				for _, v := range d.Get("domain_names").([]any) {
					newDomainNames = append(newDomainNames, v.(string))
				}
				updatePrivateEndpointRule.DomainNames = newDomainNames
			} else if d.HasChange("resource_names") {
				updateMask = append(updateMask, "resource_names")
				newResourceNames := []string{}
				for _, v := range d.Get("resource_names").([]any) {
					newResourceNames = append(newResourceNames, v.(string))
				}
				updatePrivateEndpointRule.ResourceNames = newResourceNames
			}
			rule, err := acc.NetworkConnectivity.UpdatePrivateEndpointRule(ctx, settings.UpdateNccPrivateEndpointRuleRequest{
				NetworkConnectivityConfigId: nccId,
				PrivateEndpointRuleId:       ruleId,
				PrivateEndpointRule:         updatePrivateEndpointRule,
				UpdateMask:                  strings.Join(updateMask, ","),
			})
			if err != nil {
				return err
			}
			common.StructToData(rule, s, d)
			p.Pack(d)
			return nil
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
