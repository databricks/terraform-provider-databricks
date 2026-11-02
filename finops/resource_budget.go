package finops

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceBudget() common.Resource {
	s := common.StructToSchema(billing.BudgetConfiguration{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(m, "display_name").SetValidateFunc(validation.StringLenBetween(1, 128))
		for _, p := range []string{"account_id", "budget_configuration_id", "create_time", "update_time"} {
			common.CustomizeSchemaPath(m, p).SetComputed()
		}
		common.CustomizeSchemaPath(m, "resource_type").SetOptional().SetComputed()
		common.CustomizeSchemaPath(m, "filter").
			SetCustomSuppressDiff(suppressEmptyFilter)
		common.CustomizeSchemaPath(m, "alert_configurations").
			SetCustomSuppressDiff(suppressReorderedAlerts)
		common.CustomizeSchemaPath(m, "alert_configurations", "alert_configuration_id").SetComputed()
		common.CustomizeSchemaPath(m, "alert_configurations", "scope_type").SetOptional().SetComputed()
		common.CustomizeSchemaPath(m, "alert_configurations", "action_configurations", "action_configuration_id").SetComputed()
		// We need SuppressDiff because API returns a string representation of BigDecimal with a lot
		// of trailing 0s, etc.
		common.CustomizeSchemaPath(m, "alert_configurations", "quantity_threshold").
			SetCustomSuppressDiff(suppressDecimalDiff)
		common.CustomizeSchemaPath(m, "alert_configurations", "principal_overrides",
			"override_threshold").SetCustomSuppressDiff(suppressDecimalDiff)
		return m
	})
	p := common.NewPairID("account_id", "budget_configuration_id")
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create billing.CreateBudgetConfigurationBudget
			common.DataToStructPointer(d, s, &create)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			budget, err := acc.Budgets.Create(ctx, billing.CreateBudgetConfigurationRequest{Budget: create})
			if err != nil {
				return err
			}
			d.Set("budget_configuration_id", budget.Budget.BudgetConfigurationId)
			d.Set("account_id", c.Config.AccountID)
			orderAlerts(budget.Budget, alertsOf(d, s))
			common.StructToData(budget.Budget, s, d)
			p.Pack(d)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, id, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			budget, err := acc.Budgets.GetByBudgetId(ctx, id)
			if err != nil {
				return err
			}
			orderAlerts(budget.Budget, alertsOf(d, s))
			return common.StructToData(budget.Budget, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update billing.UpdateBudgetConfigurationBudget
			_, id, err := p.Unpack(d)
			if err != nil {
				return err
			}
			common.DataToStructPointer(d, s, &update)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			current, err := acc.Budgets.GetByBudgetId(ctx, id)
			if err != nil {
				return err
			}
			clearUnsetScopeTypes(d, update.AlertConfigurations)
			assignAlertIDs(update.AlertConfigurations, current.Budget.AlertConfigurations)
			budget, err := acc.Budgets.Update(ctx, billing.UpdateBudgetConfigurationRequest{
				Budget:   update,
				BudgetId: id,
			})
			if err != nil {
				return err
			}
			orderAlerts(budget.Budget, update.AlertConfigurations)
			return common.StructToData(budget.Budget, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			_, id, err := p.Unpack(d)
			if err != nil {
				return err
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			return acc.Budgets.DeleteByBudgetId(ctx, id)
		},
		Schema: s,
	}
}

// suppressEmptyFilter suppresses adding or removing an empty filter, as the API returns
// an empty filter for budgets without one.
func suppressEmptyFilter(k, _, _ string, d *schema.ResourceData) bool {
	before, after := d.GetChange("filter")
	return k == "filter.#" && isEmptyFilter(before) && isEmptyFilter(after)
}

func isEmptyFilter(v any) bool {
	filters, _ := v.([]any)
	for _, f := range filters {
		filter, _ := f.(map[string]any)
		tags, _ := filter["tags"].([]any)
		workspaces, _ := filter["workspace_id"].([]any)
		if len(tags) > 0 || len(workspaces) > 0 {
			return false
		}
	}
	return true
}
