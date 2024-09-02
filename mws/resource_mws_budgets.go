package mws

import (
	"context"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceMwsBudget() common.Resource {
	s := common.StructToSchema(billing.BudgetConfiguration{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(m, "display_name").SetValidateFunc(validation.StringLenBetween(1, 128))
		for _, p := range []string{"account_id", "budget_configuration_id", "create_time", "update_time"} {
			common.CustomizeSchemaPath(m, p).SetComputed()
		}
		return m
	})
	p := common.NewPairSeparatedID("account_id", "budget_configuration_id", "/")
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create billing.CreateBudgetConfigurationBudget
			common.DataToStructPointer(d, s, &create)
			acc, err := c.AccountClientWithAccountIdFromConfig(d)
			if err != nil {
				return err
			}
			budget, err := acc.Budgets.Create(ctx, billing.CreateBudgetConfigurationRequest{Budget: create})
			if err != nil {
				return err
			}
			d.SetId(budget.Budget.BudgetConfigurationId)
			d.Set("budget_configuration_id", budget.Budget.BudgetConfigurationId)
			d.Set("account_id", c.Config.AccountID)
			common.StructToData(budget, s, d)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, id, err := p.Unpack(d)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			res, err := acc.Budgets.GetByBudgetId(ctx, id)
			if err != nil {
				return err
			}
			return common.StructToData(res.Budget, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, id, err := p.Unpack(d)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			return acc.Budgets.DeleteByBudgetId(ctx, id)
		},
		Schema: s,
	}
}
