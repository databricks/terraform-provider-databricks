package mws

import (
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func getTestBudget() *billing.BudgetConfiguration {
	return &billing.BudgetConfiguration{
		AccountId: "account_id",
		AlertConfigurations: []billing.AlertConfiguration{
			{
				ActionConfigurations: []billing.ActionConfiguration{
					{
						ActionType: billing.ActionConfigurationTypeEmailNotification,
						Target:     "me@databricks.com",
					},
				},
				AlertConfigurationId: "alert_configuration_id",
				QuantityThreshold:    "840",
				QuantityType:         billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				TimePeriod:           billing.AlertConfigurationTimePeriodMonth,
				TriggerType:          billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
			},
		},
		Filter: &billing.BudgetConfigurationFilter{
			Tags: []billing.BudgetConfigurationFilterTagClause{
				{
					Key: "Environment",
					Value: &billing.BudgetConfigurationFilterClause{
						Operator: billing.BudgetConfigurationFilterOperatorIn,
						Values:   []string{"Testing"},
					},
				},
			},
			WorkspaceId: &billing.BudgetConfigurationFilterWorkspaceIdClause{
				Operator: billing.BudgetConfigurationFilterOperatorIn,
				Values: []int64{
					1234567890098765,
				},
			},
		},
		BudgetConfigurationId: "budget_configuration_id",
		DisplayName:           "budget_name",
	}
}

func TestResourceBudgetCreate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockbudgetsAPI().EXPECT()
			api.Create(mock.Anything, billing.CreateBudgetConfigurationRequest{
				Budget: billing.CreateBudgetConfigurationBudget{
					AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{
						{
							ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{
								{
									ActionType: billing.ActionConfigurationTypeEmailNotification,
									Target:     "me@databricks.com",
								},
							},
							QuantityThreshold: "840",
							QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
							TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
							TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
						},
					},
					DisplayName: getTestBudget().DisplayName,
					Filter:      getTestBudget().Filter,
				},
			}).Return(&billing.CreateBudgetConfigurationResponse{Budget: getTestBudget()}, nil)
			api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
				&billing.GetBudgetConfigurationResponse{Budget: getTestBudget()}, nil,
			)
		},
		Create:    true,
		AccountID: "account_id",
		HCL: `
		display_name = "budget_name"
	
		alert_configurations {
			time_period         = "MONTH"
			trigger_type        = "CUMULATIVE_SPENDING_EXCEEDED" 
			quantity_type       = "LIST_PRICE_DOLLARS_USD"
			quantity_threshold  = "840"
	
			action_configurations {
				action_type = "EMAIL_NOTIFICATION"
				target      = "me@databricks.com"
			}
		}

		filter {
			tags {
				key   = "Environment"
				value {
					operator = "IN"
					values = ["Testing"]
				}
			}

			workspace_id {
				operator = "IN"
				values   = [
					1234567890098765
				]                              
			}
		}
		`,
		Resource: ResourceMwsBudget(),
	}.ApplyAndExpectData(t, nil) // ???
}

func TestResourceMwsBudgetRead(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockbudgetsAPI().EXPECT().
				GetByBudgetId(mock.Anything, "budget_configuration_id").
				Return(&billing.GetBudgetConfigurationResponse{Budget: getTestBudget()}, nil)
		},
		Resource:  ResourceMwsBudget(),
		Read:      true,
		New:       true,
		AccountID: "account_id",
		ID:        "account_id|budget_configuration_id",
	}.ApplyAndExpectData(t, nil) // ???
}

func TestResourceMwsBudgetDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockbudgetsAPI().EXPECT().DeleteByBudgetId(mock.Anything, "budget_configuration_id").Return(nil)
		},
		Resource:  ResourceMwsBudget(),
		AccountID: "account_id",
		Delete:    true,
		ID:        "account_id|budget_configuration_id",
	}.ApplyAndExpectData(t, nil)
}
