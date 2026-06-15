package mws

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/databricks/terraform-provider-databricks/common"
)

func ResourceMwsLogDelivery() common.Resource {
	p := common.NewPairID("account_id", "config_id")
	s := common.StructToSchema(billing.CreateLogDeliveryConfigurationParams{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			s["account_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			}
			s["config_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			}

			for _, field := range []string{
				"config_name",
				"credentials_id",
				"delivery_path_prefix",
				"delivery_start_time",
				"log_type",
				"output_format",
				"storage_configuration_id",
				"workspace_ids_filter",
			} {
				common.CustomizeSchemaPath(s, field).SetForceNew()
			}
			common.CustomizeSchemaPath(s, "config_name").SetValidateFunc(validation.StringLenBetween(0, 255))
			common.CustomizeSchemaPath(s, "delivery_start_time").SetComputed().SetCustomSuppressDiff(func(
				k, old, new string, d *schema.ResourceData) bool {
				return false
			})
			common.CustomizeSchemaPath(s, "status").SetComputed()
			return s
		})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClientWithAccountIdFromConfig(d)
			if err != nil {
				return err
			}
			var create billing.CreateLogDeliveryConfigurationParams
			common.DataToStructPointer(d, s, &create)
			response, err := acc.LogDelivery.Create(ctx, billing.WrappedCreateLogDeliveryConfiguration{
				LogDeliveryConfiguration: create,
			})
			if err != nil {
				return err
			}
			if response.LogDeliveryConfiguration == nil {
				return fmt.Errorf("empty log delivery configuration response")
			}
			if err = d.Set("config_id", response.LogDeliveryConfiguration.ConfigId); err != nil {
				return err
			}
			p.Pack(d)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, configID, err := c.AccountClientWithAccountIdFromPair(d, p)
			if err != nil {
				return err
			}
			return acc.LogDelivery.PatchStatus(ctx, billing.UpdateLogDeliveryConfigurationStatusRequest{
				LogDeliveryConfigurationId: configID,
				Status:                     billing.LogDeliveryConfigStatus(d.Get("status").(string)),
			})
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, configID, err := c.AccountClientWithAccountIdFromPair(d, p)
			if err != nil {
				return err
			}
			response, err := acc.LogDelivery.GetByLogDeliveryConfigurationId(ctx, configID)
			if err != nil {
				return err
			}
			if response.LogDeliveryConfiguration == nil {
				return fmt.Errorf("log delivery configuration not found: %w", apierr.ErrNotFound)
			}
			if response.LogDeliveryConfiguration.ConfigId == "" &&
				response.LogDeliveryConfiguration.Status == billing.LogDeliveryConfigStatusDisabled {
				d.SetId("")
				return nil
			}
			return common.StructToData(response.LogDeliveryConfiguration, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, configID, err := c.AccountClientWithAccountIdFromPair(d, p)
			if err != nil {
				return err
			}
			return acc.LogDelivery.PatchStatus(ctx, billing.UpdateLogDeliveryConfigurationStatusRequest{
				LogDeliveryConfigurationId: configID,
				Status:                     billing.LogDeliveryConfigStatusDisabled,
			})
		},
	}
}
