package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/provisioning"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceMwsPrivateAccessSettings() common.Resource {
	s := common.StructToSchema(provisioning.PrivateAccessSettings{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// nolint
		s["private_access_settings_name"].ValidateFunc = validation.StringLenBetween(4, 256)
		common.SetRequired(s["private_access_settings_name"])
		common.SetRequired(s["region"])

		s["private_access_level"].ValidateFunc = validation.StringInSlice([]string{"ACCOUNT", "ENDPOINT"}, true)
		common.SetDefault(s["private_access_level"], "ACCOUNT")

		s["private_access_settings_id"].Computed = true

		common.AddAccountIdField(s)
		return s
	})
	p := common.NewPairSeparatedID("account_id", "private_access_settings_id", "/")
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, err := c.AccountClientWithAccountIdFromConfig(d)
			if err != nil {
				return err
			}
			var pas provisioning.UpsertPrivateAccessSettingsRequest
			common.DataToStructPointer(d, s, &pas)
			res, err := a.PrivateAccess.Create(ctx, pas)
			if err != nil {
				return err
			}
			d.Set("private_access_settings_id", res.PrivateAccessSettingsId)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, pasID, err := c.AccountClientWithAccountIdFromPair(d, p)
			if err != nil {
				return err
			}
			pas, err := a.PrivateAccess.GetByPrivateAccessSettingsId(ctx, pasID)
			if err != nil {
				return err
			}
			return common.StructToData(pas, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, pasID, err := c.AccountClientWithAccountIdFromPair(d, p)
			if err != nil {
				return err
			}
			var pas provisioning.UpsertPrivateAccessSettingsRequest
			common.DataToStructPointer(d, s, &pas)
			pas.PrivateAccessSettingsId = pasID
			return a.PrivateAccess.Replace(ctx, pas)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, pasID, err := c.AccountClientWithAccountIdFromPair(d, p)
			if err != nil {
				return err
			}
			return a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, pasID)
		},
	}
}
