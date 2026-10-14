package finops

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/stretchr/testify/mock"

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
				QuantityThreshold: "840.840000000000000000",
				QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
				TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
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
			api := a.GetMockBudgetsAPI().EXPECT()
			api.Create(mock.Anything, billing.CreateBudgetConfigurationRequest{
				Budget: billing.CreateBudgetConfigurationBudget{
					AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{
						{
							ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{
								{
									ActionType: getTestBudget().AlertConfigurations[0].ActionConfigurations[0].ActionType,
									Target:     getTestBudget().AlertConfigurations[0].ActionConfigurations[0].Target,
								},
							},
							QuantityThreshold: "840.84",
							QuantityType:      getTestBudget().AlertConfigurations[0].QuantityType,
							TimePeriod:        getTestBudget().AlertConfigurations[0].TimePeriod,
							TriggerType:       getTestBudget().AlertConfigurations[0].TriggerType,
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
		Host:      "https://accounts.cloud.databricks.com",
		HCL: `
		display_name = "budget_name"

		alert_configurations {
			time_period         = "MONTH"
			trigger_type        = "CUMULATIVE_SPENDING_EXCEEDED"
			quantity_type       = "LIST_PRICE_DOLLARS_USD"
			quantity_threshold  = "840.84"

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
		Resource: ResourceBudget(),
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":           "budget_name",
		"id":                     "account_id|budget_configuration_id",
		"alert_configurations.#": 1,
		"filter.#":               1,
	})
}

func TestResourceBudgetRead(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockBudgetsAPI().EXPECT().
				GetByBudgetId(mock.Anything, "budget_configuration_id").
				Return(&billing.GetBudgetConfigurationResponse{Budget: getTestBudget()}, nil)
		},
		Resource:  ResourceBudget(),
		Read:      true,
		New:       true,
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "account_id|budget_configuration_id",
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":           "budget_name",
		"id":                     "account_id|budget_configuration_id",
		"alert_configurations.#": 1,
		"filter.#":               1,
	})
}

func TestResourceBudgetRead_UnpackError(t *testing.T) {
	qa.ResourceFixture{
		Resource:  ResourceBudget(),
		Read:      true,
		New:       true,
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "budget_configuration_id",
	}.ExpectError(t, "invalid ID: budget_configuration_id")
}

func TestResourceBudgetUpdate(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockBudgetsAPI().EXPECT()
			api.Update(mock.Anything, billing.UpdateBudgetConfigurationRequest{
				Budget: billing.UpdateBudgetConfigurationBudget{
					AccountId: getTestBudget().AccountId,
					AlertConfigurations: []billing.AlertConfiguration{
						{
							ActionConfigurations: []billing.ActionConfiguration{
								{
									ActionType: getTestBudget().AlertConfigurations[0].ActionConfigurations[0].ActionType,
									Target:     getTestBudget().AlertConfigurations[0].ActionConfigurations[0].Target,
								},
							},
							QuantityThreshold: "840.84",
							QuantityType:      getTestBudget().AlertConfigurations[0].QuantityType,
							TimePeriod:        getTestBudget().AlertConfigurations[0].TimePeriod,
							TriggerType:       getTestBudget().AlertConfigurations[0].TriggerType,
						},
					},
					BudgetConfigurationId: getTestBudget().BudgetConfigurationId,
					DisplayName:           fmt.Sprintf("%s_update", getTestBudget().DisplayName),
					Filter:                getTestBudget().Filter,
				},
				BudgetId: "budget_configuration_id",
			}).Return(&billing.UpdateBudgetConfigurationResponse{Budget: getTestBudget()}, nil)
			api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
				&billing.GetBudgetConfigurationResponse{Budget: &billing.BudgetConfiguration{
					AccountId:             getTestBudget().AccountId,
					AlertConfigurations:   getTestBudget().AlertConfigurations,
					BudgetConfigurationId: getTestBudget().BudgetConfigurationId,
					DisplayName:           fmt.Sprintf("%s_update", getTestBudget().DisplayName),
					Filter:                getTestBudget().Filter,
				}}, nil,
			)
		},
		Resource: ResourceBudget(),
		Update:   true,
		HCL: `
		display_name = "budget_name_update"

		alert_configurations {
			time_period         = "MONTH"
			trigger_type        = "CUMULATIVE_SPENDING_EXCEEDED"
			quantity_type       = "LIST_PRICE_DOLLARS_USD"
			quantity_threshold  = "840.84"

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
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		ID:        "account_id|budget_configuration_id",
	}.ApplyAndExpectData(t, map[string]any{
		"display_name": "budget_name_update",
		"id":           "account_id|budget_configuration_id",
	})
}

func TestResourceBudgetDelete(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockBudgetsAPI().EXPECT().DeleteByBudgetId(mock.Anything, "budget_configuration_id").Return(nil)
		},
		Resource:  ResourceBudget(),
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		Delete:    true,
		ID:        "account_id|budget_configuration_id",
	}.ApplyAndExpectData(t, nil)
}

func genieFilter() *billing.BudgetConfigurationFilter {
	return &billing.BudgetConfigurationFilter{
		Tags: []billing.BudgetConfigurationFilterTagClause{
			{
				Key: "databricks-product",
				Value: &billing.BudgetConfigurationFilterClause{
					Operator: billing.BudgetConfigurationFilterOperatorIn,
					Values:   []string{"genie"},
				},
			},
		},
	}
}

func getTestGenieSharedBudget() *billing.BudgetConfiguration {
	return &billing.BudgetConfiguration{
		AccountId: "account_id",
		AlertConfigurations: []billing.AlertConfiguration{
			{
				ActionConfigurations: []billing.ActionConfiguration{
					{
						ActionType: billing.ActionConfigurationTypeEmailNotification,
						Target:     "abc@gmail.com",
					},
				},
				QuantityThreshold: "2000.000000000000000000",
				QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				ScopeType:         billing.AlertConfigurationScopeTypeAlertConfigurationScopeTypeShared,
				TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
				TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
			},
		},
		BudgetConfigurationId: "budget_configuration_id",
		DisplayName:           "genie-shared-budget",
		Filter:                genieFilter(),
		ResourceType:          billing.BudgetResourceTypeBudgetResourceTypeUnityAiGateway,
	}
}

func getTestGeniePerUserBudget() *billing.BudgetConfiguration {
	return &billing.BudgetConfiguration{
		AccountId: "account_id",
		AlertConfigurations: []billing.AlertConfiguration{
			{
				ActionConfigurations: []billing.ActionConfiguration{
					{
						ActionType: billing.ActionConfigurationTypeBlockUsage,
					},
				},
				PrincipalOverrides: []billing.PrincipalOverride{
					{
						OverrideThreshold: "300",
						PrincipalId:       1234567890123456,
					},
				},
				QuantityThreshold: "100.000000000000000000",
				QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
				ScopeType:         billing.AlertConfigurationScopeTypeAlertConfigurationScopeTypePerUser,
				TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
				TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
			},
		},
		BudgetConfigurationId: "budget_configuration_id",
		DisplayName:           "genie-per-user-budget",
		Filter:                genieFilter(),
		ResourceType:          billing.BudgetResourceTypeBudgetResourceTypeUnityAiGateway,
	}
}

func TestResourceBudgetCreate_GenieShared(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockBudgetsAPI().EXPECT()
			api.Create(mock.Anything, billing.CreateBudgetConfigurationRequest{
				Budget: billing.CreateBudgetConfigurationBudget{
					AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{
						{
							ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{
								{
									ActionType: billing.ActionConfigurationTypeEmailNotification,
									Target:     "abc@gmail.com",
								},
							},
							QuantityThreshold: "2000",
							QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
							ScopeType:         billing.AlertConfigurationScopeTypeAlertConfigurationScopeTypeShared,
							TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
							TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
						},
					},
					DisplayName:  "genie-shared-budget",
					Filter:       genieFilter(),
					ResourceType: billing.BudgetResourceTypeBudgetResourceTypeUnityAiGateway,
				},
			}).Return(&billing.CreateBudgetConfigurationResponse{Budget: getTestGenieSharedBudget()}, nil)
			api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
				&billing.GetBudgetConfigurationResponse{Budget: getTestGenieSharedBudget()}, nil,
			)
		},
		Create:    true,
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		HCL: `
		display_name  = "genie-shared-budget"
		resource_type = "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY"

		filter {
			tags {
				key = "databricks-product"
				value {
					operator = "IN"
					values   = ["genie"]
				}
			}
		}

		alert_configurations {
			time_period        = "MONTH"
			trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
			quantity_type      = "LIST_PRICE_DOLLARS_USD"
			quantity_threshold = "2000"
			scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_SHARED"

			action_configurations {
				action_type = "EMAIL_NOTIFICATION"
				target      = "abc@gmail.com"
			}
		}
		`,
		Resource: ResourceBudget(),
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":                      "genie-shared-budget",
		"id":                                "account_id|budget_configuration_id",
		"resource_type":                     "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY",
		"alert_configurations.#":            1,
		"alert_configurations.0.scope_type": "ALERT_CONFIGURATION_SCOPE_TYPE_SHARED",
		"filter.#":                          1,
	})
}

func TestResourceBudgetCreate_GeniePerUserBlockUsage(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockBudgetsAPI().EXPECT()
			api.Create(mock.Anything, billing.CreateBudgetConfigurationRequest{
				Budget: billing.CreateBudgetConfigurationBudget{
					AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{
						{
							ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{
								{
									ActionType: billing.ActionConfigurationTypeBlockUsage,
								},
							},
							PrincipalOverrides: []billing.PrincipalOverride{
								{
									OverrideThreshold: "300",
									PrincipalId:       1234567890123456,
								},
							},
							QuantityThreshold: "100",
							QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
							ScopeType:         billing.AlertConfigurationScopeTypeAlertConfigurationScopeTypePerUser,
							TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
							TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
						},
					},
					DisplayName:  "genie-per-user-budget",
					Filter:       genieFilter(),
					ResourceType: billing.BudgetResourceTypeBudgetResourceTypeUnityAiGateway,
				},
			}).Return(&billing.CreateBudgetConfigurationResponse{Budget: getTestGeniePerUserBudget()}, nil)
			api.GetByBudgetId(mock.Anything, "budget_configuration_id").Return(
				&billing.GetBudgetConfigurationResponse{Budget: getTestGeniePerUserBudget()}, nil,
			)
		},
		Create:    true,
		AccountID: "account_id",
		Host:      "https://accounts.cloud.databricks.com",
		HCL: `
		display_name  = "genie-per-user-budget"
		resource_type = "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY"

		filter {
			tags {
				key = "databricks-product"
				value {
					operator = "IN"
					values   = ["genie"]
				}
			}
		}

		alert_configurations {
			time_period        = "MONTH"
			trigger_type       = "CUMULATIVE_SPENDING_EXCEEDED"
			quantity_type      = "LIST_PRICE_DOLLARS_USD"
			quantity_threshold = "100"
			scope_type         = "ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER"

			principal_overrides {
				principal_id       = 1234567890123456
				override_threshold = "300"
			}

			action_configurations {
				action_type = "BLOCK_USAGE"
			}
		}
		`,
		Resource: ResourceBudget(),
	}.ApplyAndExpectData(t, map[string]any{
		"display_name":                      "genie-per-user-budget",
		"id":                                "account_id|budget_configuration_id",
		"resource_type":                     "BUDGET_RESOURCE_TYPE_UNITY_AI_GATEWAY",
		"alert_configurations.#":            1,
		"alert_configurations.0.scope_type": "ALERT_CONFIGURATION_SCOPE_TYPE_PER_USER",
		"alert_configurations.0.principal_overrides.#":                    1,
		"alert_configurations.0.principal_overrides.0.principal_id":       1234567890123456,
		"alert_configurations.0.principal_overrides.0.override_threshold": "300",
		"alert_configurations.0.action_configurations.#":                  1,
		"alert_configurations.0.action_configurations.0.action_type":      "BLOCK_USAGE",
		"filter.#": 1,
	})
}
