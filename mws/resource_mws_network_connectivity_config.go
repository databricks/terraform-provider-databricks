package mws

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// isNccStillInUseError reports whether the given error is the Databricks
// Accounts API's "NCC is still in use" response — which is expected during
// teardown when rule/binding deletes haven't yet propagated. These errors
// are safe to retry; the backend eventually observes the dependent deletes
// and allows the NCC removal.
func isNccStillInUseError(err error) bool {
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) {
		return false
	}
	msg := strings.ToLower(apiErr.Message)
	if strings.Contains(msg, "has one or more private endpoint rules") {
		return true
	}
	if strings.Contains(msg, "attached to one or more workspaces") {
		return true
	}
	return false
}

func ResourceMwsNetworkConnectivityConfig() common.Resource {
	s := common.StructToSchema(settings.NetworkConnectivityConfiguration{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(m, "name").SetValidateFunc(validation.StringLenBetween(3, 30))
		for _, p := range []string{"name", "region"} {
			common.CustomizeSchemaPath(m, p).SetRequired().SetForceNew()
		}
		for _, p := range []string{"account_id", "network_connectivity_config_id", "creation_time", "updated_time", "egress_config"} {
			common.CustomizeSchemaPath(m, p).SetComputed()
		}
		return m
	})
	p := common.NewPairSeparatedID("account_id", "network_connectivity_config_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create settings.CreateNetworkConnectivityConfigRequest
			common.DataToStructPointer(d, s, &create.NetworkConnectivityConfig)
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
			// Retry while the backend reports the NCC is still in use
			// (rules still attached, or workspaces still bound). This absorbs
			// eventual-consistency delays between dependent deletes and the
			// NCC's own deletion check.
			return resource.RetryContext(ctx, 15*time.Minute, func() *resource.RetryError {
				err := acc.NetworkConnectivity.DeleteNetworkConnectivityConfigurationByNetworkConnectivityConfigId(ctx, nccId)
				if err == nil {
					return nil
				}
				if isNccStillInUseError(err) {
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			})
		},
	}
}
