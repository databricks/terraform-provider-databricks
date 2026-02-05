package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Tests for Custom App Integrations

func TestListCustomAppIntegrations(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		integrations := []oauth2.GetCustomAppIntegrationOutput{
			{
				IntegrationId: "integration-1",
				Name:          "Test Integration 1",
			},
			{
				IntegrationId: "integration-2",
				Name:          "Test Integration 2",
			},
		}
		ma.GetMockCustomAppIntegrationAPI().EXPECT().
			List(mock.Anything, oauth2.ListCustomAppIntegrationsRequest{}).
			Return(createIteratorFromSlice(integrations))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth")

		err := resourcesMap["databricks_custom_app_integration"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_custom_app_integration[<unknown>] (id: integration-1)"])
		assert.True(t, ic.testEmits["databricks_custom_app_integration[<unknown>] (id: integration-2)"])
	})
}

func TestListCustomAppIntegrationsWithMatch(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		integrations := []oauth2.GetCustomAppIntegrationOutput{
			{
				IntegrationId: "integration-1",
				Name:          "Test Integration",
			},
			{
				IntegrationId: "integration-2",
				Name:          "Production Integration",
			},
		}
		ma.GetMockCustomAppIntegrationAPI().EXPECT().
			List(mock.Anything, oauth2.ListCustomAppIntegrationsRequest{}).
			Return(createIteratorFromSlice(integrations))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth")
		ic.match = "Test"

		err := resourcesMap["databricks_custom_app_integration"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_custom_app_integration[<unknown>] (id: integration-1)"])
	})
}

func TestListCustomAppIntegrationsEmpty(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockCustomAppIntegrationAPI().EXPECT().
			List(mock.Anything, oauth2.ListCustomAppIntegrationsRequest{}).
			Return(createIteratorFromSlice([]oauth2.GetCustomAppIntegrationOutput{}))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth")

		err := resourcesMap["databricks_custom_app_integration"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}

// Tests for Account Federation Policies

func TestListAccountFederationPolicies(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		policies := []oauth2.FederationPolicy{
			{
				PolicyId:           "policy-1",
				Name:               "accounts/test-account/federationPolicies/policy-1",
				ServicePrincipalId: 0, // Not a service principal policy
			},
			{
				PolicyId:           "policy-2",
				Name:               "accounts/test-account/federationPolicies/policy-2",
				ServicePrincipalId: 0,
			},
		}
		ma.GetMockAccountFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListAccountFederationPoliciesRequest{}).
			Return(createIteratorFromSlice(policies))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth")

		err := resourcesMap["databricks_account_federation_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_account_federation_policy[acc_fed_policy_policy-1] (id: policy-1)"])
		assert.True(t, ic.testEmits["databricks_account_federation_policy[acc_fed_policy_policy-2] (id: policy-2)"])
	})
}

func TestListAccountFederationPoliciesWithMatch(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		policies := []oauth2.FederationPolicy{
			{
				PolicyId:           "test-policy",
				Name:               "accounts/test-account/federationPolicies/test-policy",
				ServicePrincipalId: 0,
			},
			{
				PolicyId:           "prod-policy",
				Name:               "accounts/test-account/federationPolicies/prod-policy",
				ServicePrincipalId: 0,
			},
		}
		ma.GetMockAccountFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListAccountFederationPoliciesRequest{}).
			Return(createIteratorFromSlice(policies))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth")
		ic.match = "test-policy"

		err := resourcesMap["databricks_account_federation_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_account_federation_policy[acc_fed_policy_test-policy] (id: test-policy)"])
	})
}

func TestListAccountFederationPoliciesSkipServicePrincipal(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		policies := []oauth2.FederationPolicy{
			{
				PolicyId:           "account-policy",
				Name:               "accounts/test-account/federationPolicies/account-policy",
				ServicePrincipalId: 0, // Account-level policy
			},
			{
				PolicyId:           "sp-policy",
				Name:               "accounts/test-account/servicePrincipals/123/federationPolicies/sp-policy",
				ServicePrincipalId: 123, // Service principal policy - should be skipped
			},
		}
		ma.GetMockAccountFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListAccountFederationPoliciesRequest{}).
			Return(createIteratorFromSlice(policies))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth")

		err := resourcesMap["databricks_account_federation_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_account_federation_policy[acc_fed_policy_account-policy] (id: account-policy)"])
		assert.False(t, ic.testEmits["databricks_account_federation_policy[acc_fed_policy_sp-policy] (id: sp-policy)"])
	})
}

func TestListAccountFederationPoliciesEmpty(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockAccountFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListAccountFederationPoliciesRequest{}).
			Return(createIteratorFromSlice([]oauth2.FederationPolicy{}))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth")

		err := resourcesMap["databricks_account_federation_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}

// Tests for Service Principal Federation Policies

func TestListServicePrincipalFederationPolicies(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		servicePrincipals := []iam.ServicePrincipal{
			{
				Id: "123",
			},
			{
				Id: "456",
			},
		}
		ma.GetMockAccountServicePrincipalsAPI().EXPECT().
			ListAll(mock.Anything, iam.ListAccountServicePrincipalsRequest{Attributes: "id"}).
			Return(servicePrincipals, nil)

		// Mock federation policies for SP 123
		policies123 := []oauth2.FederationPolicy{
			{
				PolicyId: "policy-1",
				Name:     "accounts/test-account/servicePrincipals/123/federationPolicies/policy-1",
			},
		}
		ma.GetMockServicePrincipalFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListServicePrincipalFederationPoliciesRequest{ServicePrincipalId: 123}).
			Return(createIteratorFromSlice(policies123))

		// Mock federation policies for SP 456
		policies456 := []oauth2.FederationPolicy{
			{
				PolicyId: "policy-2",
				Name:     "accounts/test-account/servicePrincipals/456/federationPolicies/policy-2",
			},
		}
		ma.GetMockServicePrincipalFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListServicePrincipalFederationPoliciesRequest{ServicePrincipalId: 456}).
			Return(createIteratorFromSlice(policies456))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth,users")

		err := resourcesMap["databricks_service_principal_federation_policy"].List(ic)
		assert.NoError(t, err)
		// 2 service principals + 2 federation policies
		require.Equal(t, 4, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_service_principal[<unknown>] (id: 123)"])
		assert.True(t, ic.testEmits["databricks_service_principal[<unknown>] (id: 456)"])
		assert.True(t, ic.testEmits["databricks_service_principal_federation_policy[sp_fed_policy_123_policy-1] (id: 123,policy-1)"])
		assert.True(t, ic.testEmits["databricks_service_principal_federation_policy[sp_fed_policy_456_policy-2] (id: 456,policy-2)"])
	})
}

func TestListServicePrincipalFederationPoliciesEmpty(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		ma.GetMockAccountServicePrincipalsAPI().EXPECT().
			ListAll(mock.Anything, iam.ListAccountServicePrincipalsRequest{Attributes: "id"}).
			Return([]iam.ServicePrincipal{}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth,users")

		err := resourcesMap["databricks_service_principal_federation_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 0, len(ic.testEmits))
	})
}

func TestListServicePrincipalFederationPoliciesNoFederationPolicies(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		servicePrincipals := []iam.ServicePrincipal{
			{
				Id: "123",
			},
		}
		ma.GetMockAccountServicePrincipalsAPI().EXPECT().
			ListAll(mock.Anything, iam.ListAccountServicePrincipalsRequest{Attributes: "id"}).
			Return(servicePrincipals, nil)

		// No federation policies for this SP
		ma.GetMockServicePrincipalFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListServicePrincipalFederationPoliciesRequest{ServicePrincipalId: 123}).
			Return(createIteratorFromSlice([]oauth2.FederationPolicy{}))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth,users")

		err := resourcesMap["databricks_service_principal_federation_policy"].List(ic)
		assert.NoError(t, err)
		// Only the service principal should be emitted
		require.Equal(t, 1, len(ic.testEmits))
		assert.True(t, ic.testEmits["databricks_service_principal[<unknown>] (id: 123)"])
	})
}

func TestListServicePrincipalFederationPoliciesNameExtraction(t *testing.T) {
	qa.MockAccountsApply(t, func(ma *mocks.MockAccountClient) {
		setupAwsAccountConfig(ma)
		servicePrincipals := []iam.ServicePrincipal{
			{
				Id: "123",
			},
		}
		ma.GetMockAccountServicePrincipalsAPI().EXPECT().
			ListAll(mock.Anything, iam.ListAccountServicePrincipalsRequest{Attributes: "id"}).
			Return(servicePrincipals, nil)

		// Policy with full path
		policies := []oauth2.FederationPolicy{
			{
				PolicyId: "complex-policy-name",
				Name:     "accounts/abc-123/servicePrincipals/123/federationPolicies/my-policy-name",
			},
		}
		ma.GetMockServicePrincipalFederationPolicyAPI().EXPECT().
			List(mock.Anything, oauth2.ListServicePrincipalFederationPoliciesRequest{ServicePrincipalId: 123}).
			Return(createIteratorFromSlice(policies))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForAccountTestWithClient(ctx, client, "oauth,users")

		err := resourcesMap["databricks_service_principal_federation_policy"].List(ic)
		assert.NoError(t, err)
		require.Equal(t, 2, len(ic.testEmits))
		// Check that the name is extracted from the path correctly
		assert.True(t, ic.testEmits["databricks_service_principal_federation_policy[sp_fed_policy_123_my-policy-name] (id: 123,complex-policy-name)"])
	})
}
