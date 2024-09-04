package mws

import (
	"context"
	"fmt"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"math/big"
	"strings"
)

func ResourceMwsBudget() common.Resource {
	s := common.StructToSchema(billing.BudgetConfiguration{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(m, "display_name").SetValidateFunc(validation.StringLenBetween(1, 128))
		for _, p := range []string{"account_id", "budget_configuration_id", "create_time", "update_time"} {
			common.CustomizeSchemaPath(m, p).SetComputed()
		}
		common.CustomizeSchemaPath(m, "alert_configurations", "alert_configuration_id").SetComputed()
		common.CustomizeSchemaPath(m, "alert_configurations", "action_configurations", "action_configuration_id").SetComputed()
		return m
	})
	p := common.NewPairID("account_id", "budget_configuration_id")
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var create billing.CreateBudgetConfigurationBudget
			common.DataToStructPointer(d, s, &create)
			for i := range create.AlertConfigurations {
				var err error
				create.AlertConfigurations[i].QuantityThreshold, err = stringToBigDecimal(create.AlertConfigurations[i].QuantityThreshold)
				if err != nil {
					return err
				}
			}
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
			for i := range budget.Budget.AlertConfigurations {
				budget.Budget.AlertConfigurations[i].QuantityThreshold, err = bigDecimalToString(budget.Budget.AlertConfigurations[i].QuantityThreshold)
				if err != nil {
					return err
				}
			}
			return common.StructToData(budget.Budget, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var update billing.UpdateBudgetConfigurationBudget
			_, id, err := p.Unpack(d)
			if err != nil {
				return err
			}
			common.DataToStructPointer(d, s, &update)
			for i := range update.AlertConfigurations {
				update.AlertConfigurations[i].QuantityThreshold, err = stringToBigDecimal(update.AlertConfigurations[i].QuantityThreshold)
				if err != nil {
					return err
				}
			}
			acc, err := c.AccountClient()
			if err != nil {
				return err
			}
			budget, err := acc.Budgets.Update(ctx, billing.UpdateBudgetConfigurationRequest{
				Budget:   update,
				BudgetId: id,
			})
			if err != nil {
				return err
			}
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

// API always return a BigDecimal for QuantityThreshold in string representation.
// To avoid drift, we convert the provided number to BigDecimal in string representation
// when making requests to the API.
func bigDecimalToString(input string) (string, error) {
	v := new(big.Rat)
	if _, ok := v.SetString(input); !ok {
		return "", fmt.Errorf("invalid input string")
	}
	s := v.FloatString(18)
	s = strings.TrimRight(s, "0")
	if strings.HasSuffix(s, ".") {
		s = s[:len(s)-1]
	}
	return s, nil
}

func stringToBigDecimal(input string) (string, error) {
	v := new(big.Rat)
	if _, ok := v.SetString(input); !ok {
		return "", fmt.Errorf("invalid input string")
	}
	return v.FloatString(18), nil
}
