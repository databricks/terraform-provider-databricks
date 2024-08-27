package mws

import (
	"context"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMwsBudget() common.Resource {
	type Tag struct {
		Key   string                                   `json:"key"`
		Value *billing.BudgetConfigurationFilterClause `json:"value"`
	}

	type WorkspaceFilter struct {
		Operator string  `json:"operator,omitempty"`
		Values   []int64 `json:"values,omitempty"`
	}

	type Filter struct {
		WorkspaceId WorkspaceFilter `json:"workspace_id,omitempty"`
		Tags        []Tag           `json:"tags,omitempty"`
	}

	type ActionConfiguration struct {
		ActionType billing.ActionConfigurationType `json:"action_type,omitempty"`
		Target     string                          `json:"target,omitempty"`
	}

	type AlertConfiguration struct {
		ActionConfigurations []ActionConfiguration                  `json:"action_configurations,omitempty"`
		QuantityThreshold    string                                 `json:"quantity_threshold,omitempty"`
		QuantityType         billing.AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
		TimePeriod           billing.AlertConfigurationTimePeriod   `json:"time_period,omitempty"`
		TriggerType          billing.AlertConfigurationTriggerType  `json:"trigger_type,omitempty"`
	}

	type Budget struct {
		BudgetConfigurationId string               `json:"budget_configuration_id,omitempty" tf:"computed"`
		DisplayName           string               `json:"display_name,omitempty" tf:"computed"`
		AlertConfigurations   []AlertConfiguration `json:"alert_configurations,omitempty"`
		Filter                *Filter              `json:"filter,omitempty"`
	}

	s := common.StructToSchema(
		Budget{},
		common.NoCustomize,
	)

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var budget billing.CreateBudgetConfigurationBudget
			common.DataToStructPointer(d, s, &budget)
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			req, err := acc.Budgets.Create(ctx, billing.CreateBudgetConfigurationRequest{
				Budget: budget,
			})
			if err != nil {
				return err
			}
			err = common.StructToData(req, s, d)
			if err != nil {
				return err
			}
			d.SetId(req.Budget.BudgetConfigurationId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			res, err := acc.Budgets.GetByBudgetId(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*res.Budget, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			err = acc.Budgets.DeleteByBudgetId(ctx, d.Id())
			return err
		},
		Schema: s,
	}
}
