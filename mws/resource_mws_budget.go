package mws

import (
	"context"
	"fmt"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strconv"
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
				create.AlertConfigurations[i].QuantityThreshold, err = StringToBigDecimal(create.AlertConfigurations[i].QuantityThreshold)
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
				budget.Budget.AlertConfigurations[i].QuantityThreshold, err = BigDecimalToString(budget.Budget.AlertConfigurations[i].QuantityThreshold)
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
				update.AlertConfigurations[i].QuantityThreshold, err = StringToBigDecimal(update.AlertConfigurations[i].QuantityThreshold)
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
func BigDecimalToString(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	if strings.Contains(input, ".") {
		input = strings.TrimRight(input, "0")
		input = strings.TrimSuffix(input, ".")
	}
	return input, nil
}

func StringToBigDecimal(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	if _, err := strconv.ParseFloat(input, 64); err != nil {
		return "", fmt.Errorf("invalid input string")
	}

	if !strings.Contains(input, ".") {
		return input + ".000000000000000000", nil
	}

	parts := strings.SplitN(input, ".", 2)
	decimalPart := parts[1]

	if len(decimalPart) < 18 {
		decimalPart += strings.Repeat("0", 18-len(decimalPart))
	} else {
		decimalPart = decimalPart[:18]
	}
	return parts[0] + "." + decimalPart, nil
}
