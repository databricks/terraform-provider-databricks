package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/provisioning"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceMwsPrivateAccessSettings() common.Resource {
	pasSchema := common.StructToSchema(provisioning.PrivateAccessSettings{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		emptyCtx := common.SchemaPathContext{}
		common.CustomizeSchemaPath(emptyCtx, m, "private_access_settings_name").SetValidateFunc(validation.StringLenBetween(4, 256))
		common.CustomizeSchemaPath(emptyCtx, m, "private_access_settings_name").SetRequired()

		common.CustomizeSchemaPath(emptyCtx, m, "region").SetRequired()

		common.CustomizeSchemaPath(emptyCtx, m, "private_access_settings_id").SetComputed()

		common.CustomizeSchemaPath(emptyCtx, m, "private_access_level").SetValidateFunc(validation.StringInSlice([]string{"ACCOUNT", "ENDPOINT"}, true))
		common.CustomizeSchemaPath(emptyCtx, m, "private_access_level").SetDefault("ACCOUNT")

		common.AddAccountIdField(m)
		return m
	})
	p := common.NewPairSeparatedID("account_id", "private_access_settings_id", "/")
	return common.Resource{
		Schema: pasSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, err := c.AccountClientWithAccountIdFromConfig(d)
			if err != nil {
				return err
			}
			var pas provisioning.UpsertPrivateAccessSettingsRequest
			common.DataToStructPointer(d, pasSchema, &pas)
			common.SetForceSendFields(&pas, d, []string{"public_access_enabled"})
			res, err := a.PrivateAccess.Create(ctx, pas)
			if err != nil {
				return err
			}
			d.Set("private_access_settings_id", res.PrivateAccessSettingsId)
			d.Set("account_id", res.AccountId)
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
			return common.StructToData(pas, pasSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			a, pasID, err := c.AccountClientWithAccountIdFromPair(d, p)
			if err != nil {
				return err
			}
			var pas provisioning.UpsertPrivateAccessSettingsRequest
			common.DataToStructPointer(d, pasSchema, &pas)
			common.SetForceSendFields(&pas, d, []string{"public_access_enabled"})
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
