package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListAccountNetworkPolicies(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockNetworkPoliciesAPI().EXPECT().ListNetworkPoliciesRpcAll(mock.Anything, settings.ListNetworkPoliciesRequest{}).Return([]settings.AccountNetworkPolicy{
			{
				NetworkPolicyId: "policy-1",
			},
			{
				NetworkPolicyId: "policy-2",
			},
			{
				NetworkPolicyId: "default-policy", // Should be skipped
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")

		err := resourcesMap["databricks_account_network_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_account_network_policy[<unknown>] (id: policy-1)"])
		assert.True(t, ic.testEmits["databricks_account_network_policy[<unknown>] (id: policy-2)"])
		assert.False(t, ic.testEmits["databricks_account_network_policy[<unknown>] (id: default-policy)"])
	})
}

func TestListAccountNetworkPoliciesWithMatch(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockNetworkPoliciesAPI().EXPECT().ListNetworkPoliciesRpcAll(mock.Anything, settings.ListNetworkPoliciesRequest{}).Return([]settings.AccountNetworkPolicy{
			{
				NetworkPolicyId: "policy-1",
			},
			{
				NetworkPolicyId: "policy-2",
			},
			{
				NetworkPolicyId: "test-policy",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")
		ic.match = "test"

		err := resourcesMap["databricks_account_network_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_account_network_policy[<unknown>] (id: test-policy)"])
		assert.False(t, ic.testEmits["databricks_account_network_policy[<unknown>] (id: policy-1)"])
		assert.False(t, ic.testEmits["databricks_account_network_policy[<unknown>] (id: policy-2)"])
	})
}

func TestListWorkspaceNetworkOptions(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return([]provisioning.Workspace{
			{
				WorkspaceId:     123,
				WorkspaceName:   "workspace-1",
				WorkspaceStatus: "RUNNING",
			},
			{
				WorkspaceId:     456,
				WorkspaceName:   "workspace-2",
				WorkspaceStatus: "RUNNING",
			},
			{
				WorkspaceId:     789,
				WorkspaceName:   "workspace-3",
				WorkspaceStatus: "STOPPED", // Should be skipped
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")

		err := resourcesMap["databricks_workspace_network_option"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_workspace_network_option[<unknown>] (id: 123)"])
		assert.True(t, ic.testEmits["databricks_workspace_network_option[<unknown>] (id: 456)"])
		assert.False(t, ic.testEmits["databricks_workspace_network_option[<unknown>] (id: 789)"])
	})
}

func TestListWorkspaceNetworkOptionsWithMatch(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return([]provisioning.Workspace{
			{
				WorkspaceId:     123,
				WorkspaceName:   "test-workspace-1",
				WorkspaceStatus: "RUNNING",
			},
			{
				WorkspaceId:     456,
				WorkspaceName:   "workspace-2",
				WorkspaceStatus: "RUNNING",
			},
			{
				WorkspaceId:     789,
				WorkspaceName:   "test-workspace-3",
				WorkspaceStatus: "RUNNING",
			},
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")
		ic.match = "test"

		err := resourcesMap["databricks_workspace_network_option"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_workspace_network_option[<unknown>] (id: 123)"])
		assert.True(t, ic.testEmits["databricks_workspace_network_option[<unknown>] (id: 789)"])
		assert.False(t, ic.testEmits["databricks_workspace_network_option[<unknown>] (id: 456)"])
	})
}

func TestImportWorkspaceNetworkOption(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockWorkspaceNetworkConfigurationAPI().EXPECT().GetWorkspaceNetworkOptionRpc(mock.Anything, settings.GetWorkspaceNetworkOptionRequest{
			WorkspaceId: 123,
		}).Return(&settings.WorkspaceNetworkOption{
			WorkspaceId:     123,
			NetworkPolicyId: "policy-1",
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")

		r := &resource{
			ID: "123",
		}

		err := resourcesMap["databricks_workspace_network_option"].Import(ic, r)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_account_network_policy[<unknown>] (id: policy-1)"])
	})
}

func TestImportWorkspaceNetworkOptionDefaultPolicy(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockWorkspaceNetworkConfigurationAPI().EXPECT().GetWorkspaceNetworkOptionRpc(mock.Anything, settings.GetWorkspaceNetworkOptionRequest{
			WorkspaceId: 123,
		}).Return(&settings.WorkspaceNetworkOption{
			WorkspaceId:     123,
			NetworkPolicyId: "default-policy", // Should not be emitted
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")

		r := &resource{
			ID: "123",
		}

		err := resourcesMap["databricks_workspace_network_option"].Import(ic, r)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}

func TestImportWorkspaceNetworkOptionEmptyPolicy(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockWorkspaceNetworkConfigurationAPI().EXPECT().GetWorkspaceNetworkOptionRpc(mock.Anything, settings.GetWorkspaceNetworkOptionRequest{
			WorkspaceId: 123,
		}).Return(&settings.WorkspaceNetworkOption{
			WorkspaceId:     123,
			NetworkPolicyId: "", // Empty policy should not be emitted
		}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")

		r := &resource{
			ID: "123",
		}

		err := resourcesMap["databricks_workspace_network_option"].Import(ic, r)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}

func TestImportWorkspaceNetworkOptionInvalidId(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		// GetWorkspaceNetworkOptionRpc should not be called with invalid ID
		// No expectation set, so if it's called the test will fail
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")

		r := &resource{
			ID: "invalid-id", // Invalid workspace ID
		}

		err := resourcesMap["databricks_workspace_network_option"].Import(ic, r)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid syntax")
	})
}

func TestImportWorkspaceNetworkOptionGetError(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		ma.GetMockWorkspaceNetworkConfigurationAPI().EXPECT().GetWorkspaceNetworkOptionRpc(mock.Anything, settings.GetWorkspaceNetworkOptionRequest{
			WorkspaceId: 123,
		}).Return(nil, assert.AnError)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "seg")

		r := &resource{
			ID: "123",
		}

		// The function returns nil on error, not the error itself (it just logs a warning)
		err := resourcesMap["databricks_workspace_network_option"].Import(ic, r)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}
