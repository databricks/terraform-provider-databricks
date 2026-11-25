// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package billing_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/billing"
)

func ExampleBudgetsAPI_Create_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.Budgets.Create(ctx, billing.CreateBudgetConfigurationRequest{
		Budget: billing.CreateBudgetConfigurationBudget{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter: &billing.BudgetConfigurationFilter{
				Tags: []billing.BudgetConfigurationFilterTagClause{{
					Key: "tagName",
					Value: &billing.BudgetConfigurationFilterClause{
						Operator: billing.BudgetConfigurationFilterOperatorIn,
						Values:   []string{"all"},
					},
				}},
			},
			AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{{
				TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
				QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
				QuantityThreshold: "100",
				ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{{
					ActionType: billing.ActionConfigurationTypeEmailNotification,
					Target:     "admin@example.com",
				}},
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetConfigurationId)
	if err != nil {
		panic(err)
	}

}

func ExampleBudgetsAPI_Get_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.Budgets.Create(ctx, billing.CreateBudgetConfigurationRequest{
		Budget: billing.CreateBudgetConfigurationBudget{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter: &billing.BudgetConfigurationFilter{
				Tags: []billing.BudgetConfigurationFilterTagClause{{
					Key: "tagName",
					Value: &billing.BudgetConfigurationFilterClause{
						Operator: billing.BudgetConfigurationFilterOperatorIn,
						Values:   []string{"all"},
					},
				}},
			},
			AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{{
				TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
				QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
				QuantityThreshold: "100",
				ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{{
					ActionType: billing.ActionConfigurationTypeEmailNotification,
					Target:     "admin@example.com",
				}},
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.Budgets.GetByBudgetId(ctx, created.Budget.BudgetConfigurationId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetConfigurationId)
	if err != nil {
		panic(err)
	}

}

func ExampleBudgetsAPI_List_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.Budgets.ListAll(ctx, billing.ListBudgetConfigurationsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleBudgetsAPI_Update_budgets() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.Budgets.Create(ctx, billing.CreateBudgetConfigurationRequest{
		Budget: billing.CreateBudgetConfigurationBudget{
			DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter: &billing.BudgetConfigurationFilter{
				Tags: []billing.BudgetConfigurationFilterTagClause{{
					Key: "tagName",
					Value: &billing.BudgetConfigurationFilterClause{
						Operator: billing.BudgetConfigurationFilterOperatorIn,
						Values:   []string{"all"},
					},
				}},
			},
			AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{{
				TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
				QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
				QuantityThreshold: "100",
				ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{{
					ActionType: billing.ActionConfigurationTypeEmailNotification,
					Target:     "admin@example.com",
				}},
			}},
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = a.Budgets.Update(ctx, billing.UpdateBudgetConfigurationRequest{
		BudgetId: created.Budget.BudgetConfigurationId,
		Budget: billing.UpdateBudgetConfigurationBudget{
			BudgetConfigurationId: created.Budget.BudgetConfigurationId,
			DisplayName:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
			Filter: &billing.BudgetConfigurationFilter{
				Tags: []billing.BudgetConfigurationFilterTagClause{{
					Key: "tagName",
					Value: &billing.BudgetConfigurationFilterClause{
						Operator: billing.BudgetConfigurationFilterOperatorIn,
						Values:   []string{"all"},
					},
				}},
			},
			AlertConfigurations: []billing.AlertConfiguration{{
				AlertConfigurationId: created.Budget.AlertConfigurations[0].AlertConfigurationId,
				TimePeriod:           billing.AlertConfigurationTimePeriodMonth,
				QuantityType:         billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				TriggerType:          billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
				QuantityThreshold:    "50",
				ActionConfigurations: created.Budget.AlertConfigurations[0].ActionConfigurations,
			}},
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetConfigurationId)
	if err != nil {
		panic(err)
	}

}
