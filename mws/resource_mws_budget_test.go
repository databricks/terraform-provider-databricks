package mws

import (
	"fmt"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/stretchr/testify/assert"
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
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockbudgetsAPI().EXPECT()
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
							QuantityThreshold: getTestBudget().AlertConfigurations[0].QuantityThreshold,
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
		Resource: ResourceMwsBudget(),
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "account_id|budget_configuration_id", d.Id())
	assert.Equal(t, "budget_name", d.Get("display_name"))
	assert.Len(t, d.Get("alert_configurations"), 1)
	assert.Len(t, d.Get("filter"), 1)
}

func TestResourceBudgetRead(t *testing.T) {
	d, err := qa.ResourceFixture{
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
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "account_id|budget_configuration_id", d.Id())
	assert.Equal(t, "budget_name", d.Get("display_name"))
	assert.Len(t, d.Get("alert_configurations"), 1)
	assert.Len(t, d.Get("filter"), 1)
}

func TestResourceBudgetUpdate(t *testing.T) {
	d, err := qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			api := a.GetMockbudgetsAPI().EXPECT()
			api.Update(mock.Anything, billing.UpdateBudgetConfigurationRequest{
				Budget: billing.UpdateBudgetConfigurationBudget{
					AccountId:             getTestBudget().AccountId,
					AlertConfigurations:   getTestBudget().AlertConfigurations,
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
		Resource: ResourceMwsBudget(),
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
		ID:        "account_id|budget_configuration_id",
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "account_id|budget_configuration_id", d.Id())
	assert.Equal(t, "budget_name_update", d.Get("display_name"))
}

func TestResourceBudgetDelete(t *testing.T) {
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

func TestResourceMwsBudget_BigDecimalToString(t *testing.T) {
	v, err := BigDecimalToString("840")
	assert.NoError(t, err)
	assert.Equal(t, "840", v)

	v, err = BigDecimalToString("840.8400000000")
	assert.NoError(t, err)
	assert.Equal(t, "840.84", v)

	v, err = BigDecimalToString("840.8400123000")
	assert.NoError(t, err)
	assert.Equal(t, "840.8400123", v)

	v, err = BigDecimalToString("840.840000000000000001")
	assert.NoError(t, err)
	assert.Equal(t, "840.840000000000000001", v)

	_, err = BigDecimalToString("")
	assert.Error(t, err)
}

func TestResourceMwsBudget_StringToBigDecimal(t *testing.T) {
	v, err := StringToBigDecimal("840")
	assert.NoError(t, err)
	assert.Equal(t, "840.000000000000000000", v)

	v, err = StringToBigDecimal("840.84")
	assert.NoError(t, err)
	assert.Equal(t, "840.840000000000000000", v)

	v, err = StringToBigDecimal("840.8400123")
	assert.NoError(t, err)
	assert.Equal(t, "840.840012300000000000", v)

	v, err = StringToBigDecimal("")
	assert.Error(t, err)
}
