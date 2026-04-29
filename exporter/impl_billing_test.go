package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/finops"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListBudgetPolicies(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockBudgetPolicyAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetPoliciesRequest{}).Return([]billing.BudgetPolicy{
			{
				PolicyId:   "policy-1",
				PolicyName: "Test Policy 1",
			},
			{
				PolicyId:   "policy-2",
				PolicyName: "Test Policy 2",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		err := resourcesMap["databricks_budget_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_budget_policy[<unknown>] (id: policy-1)"])
		assert.True(t, ic.testEmits["databricks_budget_policy[<unknown>] (id: policy-2)"])
	})
}

func TestListBudgetPoliciesWithMatch(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockBudgetPolicyAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetPoliciesRequest{}).Return([]billing.BudgetPolicy{
			{
				PolicyId:   "policy-1",
				PolicyName: "Test Policy 1",
			},
			{
				PolicyId:   "policy-2",
				PolicyName: "Production Policy",
			},
			{
				PolicyId:   "policy-3",
				PolicyName: "Test Policy 3",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")
		ic.match = "Test"

		err := resourcesMap["databricks_budget_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_budget_policy[<unknown>] (id: policy-1)"])
		assert.True(t, ic.testEmits["databricks_budget_policy[<unknown>] (id: policy-3)"])
		assert.False(t, ic.testEmits["databricks_budget_policy[<unknown>] (id: policy-2)"])
	})
}

func TestListBudgetPoliciesEmpty(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockBudgetPolicyAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetPoliciesRequest{}).Return([]billing.BudgetPolicy{}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		err := resourcesMap["databricks_budget_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}

func TestListBudgetPoliciesSkipEmptyPolicyId(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockBudgetPolicyAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetPoliciesRequest{}).Return([]billing.BudgetPolicy{
			{
				PolicyId:   "policy-1",
				PolicyName: "Test Policy 1",
			},
			{
				PolicyId:   "", // Empty policy ID should be skipped
				PolicyName: "Invalid Policy",
			},
			{
				PolicyId:   "policy-2",
				PolicyName: "Test Policy 2",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		err := resourcesMap["databricks_budget_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_budget_policy[<unknown>] (id: policy-1)"])
		assert.True(t, ic.testEmits["databricks_budget_policy[<unknown>] (id: policy-2)"])
	})
}

func TestListBudgetPoliciesError(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockBudgetPolicyAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetPoliciesRequest{}).Return(nil, assert.AnError)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		err := resourcesMap["databricks_budget_policy"].List(ic)
		assert.Error(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestImportBudgetPolicyWithWorkspaceBindings(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// No additional API calls needed for import
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,access,mws")
		ic.Client.Config.Host = "https://accounts.cloud.databricks.com"

		// Create a mock resource with workspace bindings
		r := &resource{
			ID:       "policy-1",
			Resource: "databricks_budget_policy",
		}

		// Create a mock DataWrapper with binding_workspace_ids
		wrapper := &mockResourceDataWrapper{
			data: map[string]interface{}{
				"binding_workspace_ids": []int64{123, 456},
			},
		}
		r.DataWrapper = wrapper

		err := resourcesMap["databricks_budget_policy"].Import(ic, r)
		assert.NoError(t, err)

		// Should emit access control rule set and workspace resources
		require.Equal(t, 3, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_access_control_rule_set[<unknown>] (id: accounts/"+testAccountID+"/budgetPolicies/policy-1/ruleSets/default)"])
		assert.True(t, ic.testEmits["databricks_mws_workspaces[<unknown>] (id: "+testAccountID+"/123)"])
		assert.True(t, ic.testEmits["databricks_mws_workspaces[<unknown>] (id: "+testAccountID+"/456)"])
	})
}

func TestImportBudgetPolicyNoWorkspaceBindings(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// No additional API calls needed for import
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,access")

		// Create a mock resource without workspace bindings
		r := &resource{
			ID:       "policy-1",
			Resource: "databricks_budget_policy",
		}

		// Create a mock DataWrapper without binding_workspace_ids
		wrapper := &mockResourceDataWrapper{
			data: map[string]interface{}{},
		}
		r.DataWrapper = wrapper

		err := resourcesMap["databricks_budget_policy"].Import(ic, r)
		assert.NoError(t, err)

		// Should only emit access control rule set
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_access_control_rule_set[<unknown>] (id: accounts/"+testAccountID+"/budgetPolicies/policy-1/ruleSets/default)"])
	})
}

func TestImportBudgetPolicyAzure(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// No additional API calls needed for import
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,access")
		ic.Client.Config.Host = "https://accounts.azuredatabricks.net"

		// Create a mock resource with workspace bindings
		r := &resource{
			ID:       "policy-1",
			Resource: "databricks_budget_policy",
		}

		// Create a mock DataWrapper with binding_workspace_ids
		wrapper := &mockResourceDataWrapper{
			data: map[string]interface{}{
				"binding_workspace_ids": []int64{123, 456},
			},
		}
		r.DataWrapper = wrapper

		err := resourcesMap["databricks_budget_policy"].Import(ic, r)
		assert.NoError(t, err)

		// On Azure, should only emit access control rule set, not workspace resources
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_access_control_rule_set[<unknown>] (id: accounts/"+testAccountID+"/budgetPolicies/policy-1/ruleSets/default)"])
	})
}

func TestImportBudgetPolicyNilDataWrapper(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// No additional API calls needed for import
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		// Create a mock resource with nil DataWrapper
		r := &resource{
			ID:          "policy-1",
			Resource:    "databricks_budget_policy",
			DataWrapper: nil,
		}

		err := resourcesMap["databricks_budget_policy"].Import(ic, r)
		assert.NoError(t, err)

		// Should return early with no emits
		require.Equal(t, 0, len(ic.testEmits))
	})
}

// mockResourceDataWrapper is a simple mock implementation of ResourceDataWrapper for testing
type mockResourceDataWrapper struct {
	data map[string]interface{}
}

func (m *mockResourceDataWrapper) Get(key string) interface{} {
	return m.data[key]
}

func (m *mockResourceDataWrapper) GetOk(key string) (interface{}, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *mockResourceDataWrapper) Id() string {
	if id, ok := m.data["id"].(string); ok {
		return id
	}
	return ""
}

func (m *mockResourceDataWrapper) SetId(id string) {
	m.data["id"] = id
}

func (m *mockResourceDataWrapper) Set(key string, value interface{}) error {
	m.data[key] = value
	return nil
}

func (m *mockResourceDataWrapper) GetSchema() SchemaWrapper {
	return nil
}

func (m *mockResourceDataWrapper) IsPluginFramework() bool {
	return true
}

func (m *mockResourceDataWrapper) GetTypedStruct(ctx context.Context, target interface{}) error {
	return nil
}

// Tests for Budget resource

func TestListBudgets(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// Set Config on the mock AccountClient
		ma.AccountClient.Config = &config.Config{
			AccountID: testAccountID,
		}
		ma.GetMockBudgetsAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetConfigurationsRequest{}).Return([]billing.BudgetConfiguration{
			{
				BudgetConfigurationId: "budget-1",
				DisplayName:           "Test Budget 1",
				CreateTime:            1700000000000,
				UpdateTime:            1700000000000,
			},
			{
				BudgetConfigurationId: "budget-2",
				DisplayName:           "Test Budget 2",
				CreateTime:            1700000000000,
				UpdateTime:            1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		err := resourcesMap["databricks_budget"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_budget[Test Budget 1] (id: "+testAccountID+"|budget-1)"])
		assert.True(t, ic.testEmits["databricks_budget[Test Budget 2] (id: "+testAccountID+"|budget-2)"])
	})
}

func TestListBudgetsIncremental(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.AccountClient.Config = &config.Config{
			AccountID: testAccountID,
		}
		ma.GetMockBudgetsAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetConfigurationsRequest{}).Return([]billing.BudgetConfiguration{
			{
				BudgetConfigurationId: "budget-1",
				DisplayName:           "Old Budget",
				CreateTime:            1600000000000, // Old budget
				UpdateTime:            1600000000000,
			},
			{
				BudgetConfigurationId: "budget-2",
				DisplayName:           "New Budget",
				CreateTime:            1700000000000, // New budget
				UpdateTime:            1700000000000,
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")
		ic.incremental = true
		ic.updatedSinceMs = 1650000000000 // Between the two create times

		err := resourcesMap["databricks_budget"].List(ic)
		assert.NoError(t, err)
		// Only the new budget should be emitted
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_budget[New Budget] (id: "+testAccountID+"|budget-2)"])
		assert.False(t, ic.testEmits["databricks_budget[Old Budget] (id: "+testAccountID+"|budget-1)"])
	})
}

func TestListBudgetsEmpty(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockBudgetsAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetConfigurationsRequest{}).Return([]billing.BudgetConfiguration{}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		err := resourcesMap["databricks_budget"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}

func TestListBudgetsError(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockBudgetsAPI().EXPECT().ListAll(mock.Anything, billing.ListBudgetConfigurationsRequest{}).Return(nil, assert.AnError)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		err := resourcesMap["databricks_budget"].List(ic)
		assert.Error(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}

func TestImportBudgetWithWorkspaces(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.AccountClient.Config = &config.Config{
			AccountID: testAccountID,
			Host:      "https://accounts.cloud.databricks.com",
		}
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,mws")
		ic.Client.Config.Host = "https://accounts.cloud.databricks.com"

		// Create test resource data
		d := finops.ResourceBudget().ToResource().TestResourceData()
		d.SetId(testAccountID + "|budget-1")

		// Set filter with workspace IDs
		filter := []interface{}{
			map[string]interface{}{
				"workspace_id": []interface{}{
					map[string]interface{}{
						"operator": "IN",
						"values":   []interface{}{int64(123), int64(456)},
					},
				},
			},
		}
		d.Set("filter", filter)

		r := &resource{
			ID:       testAccountID + "|budget-1",
			Resource: "databricks_budget",
			Data:     d,
		}

		err := resourcesMap["databricks_budget"].Import(ic, r)
		assert.NoError(t, err)

		// Should emit workspace resources
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_workspaces[<unknown>] (id: "+testAccountID+"/123)"])
		assert.True(t, ic.testEmits["databricks_mws_workspaces[<unknown>] (id: "+testAccountID+"/456)"])
	})
}

func TestImportBudgetWithAlertEmails(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.AccountClient.Config = &config.Config{
			AccountID: testAccountID,
		}
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,users")

		// Create test resource data
		d := finops.ResourceBudget().ToResource().TestResourceData()
		d.SetId(testAccountID + "|budget-1")

		// Set alert configurations with email notifications
		alerts := []interface{}{
			map[string]interface{}{
				"alert_configuration_id": "alert-1",
				"action_configurations": []interface{}{
					map[string]interface{}{
						"action_type": "EMAIL_NOTIFICATION",
						"target":      "user1@example.com",
					},
					map[string]interface{}{
						"action_type": "EMAIL_NOTIFICATION",
						"target":      "user2@example.com",
					},
				},
			},
		}
		d.Set("alert_configurations", alerts)

		r := &resource{
			ID:       testAccountID + "|budget-1",
			Resource: "databricks_budget",
			Data:     d,
		}

		err := resourcesMap["databricks_budget"].Import(ic, r)
		assert.NoError(t, err)

		// Should emit user resources for email notifications
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: user1@example.com)"])
		assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: user2@example.com)"])
	})
}

func TestImportBudgetWithWorkspacesAndAlerts(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.AccountClient.Config = &config.Config{
			AccountID: testAccountID,
			Host:      "https://accounts.cloud.databricks.com",
		}
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,mws,users")
		ic.Client.Config.Host = "https://accounts.cloud.databricks.com"

		// Create test resource data
		d := finops.ResourceBudget().ToResource().TestResourceData()
		d.SetId(testAccountID + "|budget-1")

		// Set filter with workspace IDs
		filter := []interface{}{
			map[string]interface{}{
				"workspace_id": []interface{}{
					map[string]interface{}{
						"operator": "IN",
						"values":   []interface{}{int64(789)},
					},
				},
			},
		}
		d.Set("filter", filter)

		// Set alert configurations with email notifications
		alerts := []interface{}{
			map[string]interface{}{
				"alert_configuration_id": "alert-1",
				"action_configurations": []interface{}{
					map[string]interface{}{
						"action_type": "EMAIL_NOTIFICATION",
						"target":      "admin@example.com",
					},
				},
			},
		}
		d.Set("alert_configurations", alerts)

		r := &resource{
			ID:       testAccountID + "|budget-1",
			Resource: "databricks_budget",
			Data:     d,
		}

		err := resourcesMap["databricks_budget"].Import(ic, r)
		assert.NoError(t, err)

		// Should emit both workspace and user resources
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_mws_workspaces[<unknown>] (id: "+testAccountID+"/789)"])
		assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: admin@example.com)"])
	})
}

func TestImportBudgetAzureNoWorkspaces(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.AccountClient.Config = &config.Config{
			AccountID: testAccountID,
			Host:      "https://accounts.azuredatabricks.net",
		}
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,mws,users")
		ic.Client.Config.Host = "https://accounts.azuredatabricks.net"

		// Create test resource data
		d := finops.ResourceBudget().ToResource().TestResourceData()
		d.SetId(testAccountID + "|budget-1")

		// Set filter with workspace IDs
		filter := []interface{}{
			map[string]interface{}{
				"workspace_id": []interface{}{
					map[string]interface{}{
						"operator": "IN",
						"values":   []interface{}{int64(123), int64(456)},
					},
				},
			},
		}
		d.Set("filter", filter)

		// Set alert configurations
		alerts := []interface{}{
			map[string]interface{}{
				"alert_configuration_id": "alert-1",
				"action_configurations": []interface{}{
					map[string]interface{}{
						"action_type": "EMAIL_NOTIFICATION",
						"target":      "user@example.com",
					},
				},
			},
		}
		d.Set("alert_configurations", alerts)

		r := &resource{
			ID:       testAccountID + "|budget-1",
			Resource: "databricks_budget",
			Data:     d,
		}

		err := resourcesMap["databricks_budget"].Import(ic, r)
		assert.NoError(t, err)

		// On Azure, should only emit user, not workspaces
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_user[<unknown>] (user_name: user@example.com)"])
		assert.False(t, ic.testEmits["databricks_mws_workspaces[<unknown>] (id: "+testAccountID+"/123)"])
	})
}

func TestImportBudgetNoWorkspacesNoAlerts(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// No additional API calls needed for import
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing")

		// Create test resource data with no filter or alerts
		d := finops.ResourceBudget().ToResource().TestResourceData()
		d.SetId(testAccountID + "|budget-1")

		r := &resource{
			ID:       testAccountID + "|budget-1",
			Resource: "databricks_budget",
			Data:     d,
		}

		err := resourcesMap["databricks_budget"].Import(ic, r)
		assert.NoError(t, err)

		// Should not emit any resources
		require.Equal(t, 0, len(ic.testEmits))
	})
}

func TestImportBudgetNonEmailAction(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// No additional API calls needed for import
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "billing,users")

		// Create test resource data
		d := finops.ResourceBudget().ToResource().TestResourceData()
		d.SetId(testAccountID + "|budget-1")

		// Set alert configurations with non-email action type
		alerts := []interface{}{
			map[string]interface{}{
				"alert_configuration_id": "alert-1",
				"action_configurations": []interface{}{
					map[string]interface{}{
						"action_type": "WEBHOOK",
						"target":      "https://example.com/webhook",
					},
				},
			},
		}
		d.Set("alert_configurations", alerts)

		r := &resource{
			ID:       testAccountID + "|budget-1",
			Resource: "databricks_budget",
			Data:     d,
		}

		err := resourcesMap["databricks_budget"].Import(ic, r)
		assert.NoError(t, err)

		// Should not emit any user resources for non-email actions
		require.Equal(t, 0, len(ic.testEmits))
	})
}
