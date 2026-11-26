// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go"

	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/billing"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/disasterrecovery"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iamv2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/networking"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/oauth2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/provisioning"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settings"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settingsv2"
	"github.com/stretchr/testify/mock"
)

type MockAccountClient struct {
	AccountClient *databricks.AccountClient
}

// NewMockAccountClient creates new mocked version of Databricks SDK client for Accounts
// which can be used for testing.
func NewMockAccountClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountClient {
	cli := &MockAccountClient{
		AccountClient: &databricks.AccountClient{
			Config: nil,

			AccessControl:                    iam.NewMockAccountAccessControlInterface(t),
			BillableUsage:                    billing.NewMockBillableUsageInterface(t),
			BudgetPolicy:                     billing.NewMockBudgetPolicyInterface(t),
			Budgets:                          billing.NewMockBudgetsInterface(t),
			Credentials:                      provisioning.NewMockCredentialsInterface(t),
			CustomAppIntegration:             oauth2.NewMockCustomAppIntegrationInterface(t),
			DisasterRecovery:                 disasterrecovery.NewMockDisasterRecoveryInterface(t),
			EncryptionKeys:                   provisioning.NewMockEncryptionKeysInterface(t),
			Endpoints:                        networking.NewMockEndpointsInterface(t),
			FederationPolicy:                 oauth2.NewMockAccountFederationPolicyInterface(t),
			GroupsV2:                         iam.NewMockAccountGroupsV2Interface(t),
			IamV2:                            iamv2.NewMockAccountIamV2Interface(t),
			IpAccessLists:                    settings.NewMockAccountIpAccessListsInterface(t),
			LogDelivery:                      billing.NewMockLogDeliveryInterface(t),
			MetastoreAssignments:             catalog.NewMockAccountMetastoreAssignmentsInterface(t),
			Metastores:                       catalog.NewMockAccountMetastoresInterface(t),
			NetworkConnectivity:              settings.NewMockNetworkConnectivityInterface(t),
			NetworkPolicies:                  settings.NewMockNetworkPoliciesInterface(t),
			Networks:                         provisioning.NewMockNetworksInterface(t),
			OAuthPublishedApps:               oauth2.NewMockOAuthPublishedAppsInterface(t),
			PrivateAccess:                    provisioning.NewMockPrivateAccessInterface(t),
			PrivateNetworkGateways:           networking.NewMockPrivateNetworkGatewaysInterface(t),
			PublishedAppIntegration:          oauth2.NewMockPublishedAppIntegrationInterface(t),
			ServicePrincipalFederationPolicy: oauth2.NewMockServicePrincipalFederationPolicyInterface(t),
			ServicePrincipalSecrets:          oauth2.NewMockServicePrincipalSecretsInterface(t),
			ServicePrincipalsV2:              iam.NewMockAccountServicePrincipalsV2Interface(t),
			Settings:                         settings.NewMockAccountSettingsInterface(t),
			SettingsV2:                       settingsv2.NewMockAccountSettingsV2Interface(t),
			Storage:                          provisioning.NewMockStorageInterface(t),
			StorageCredentials:               catalog.NewMockAccountStorageCredentialsInterface(t),
			UsageDashboards:                  billing.NewMockUsageDashboardsInterface(t),
			UsersV2:                          iam.NewMockAccountUsersV2Interface(t),
			VpcEndpoints:                     provisioning.NewMockVpcEndpointsInterface(t),
			WorkspaceAssignment:              iam.NewMockWorkspaceAssignmentInterface(t),
			WorkspaceNetworkConfiguration:    settings.NewMockWorkspaceNetworkConfigurationInterface(t),
			Workspaces:                       provisioning.NewMockWorkspacesInterface(t),
			Groups:                           iam.NewMockAccountGroupsInterface(t),
			ServicePrincipals:                iam.NewMockAccountServicePrincipalsInterface(t),
			Users:                            iam.NewMockAccountUsersInterface(t),
		},
	}

	mockAccountSettingsAPI := cli.GetMockAccountSettingsAPI()

	mockCspEnablementAccount := settings.NewMockCspEnablementAccountInterface(t)
	mockAccountSettingsAPI.On("CspEnablementAccount").Return(mockCspEnablementAccount).Maybe()

	mockDisableLegacyFeatures := settings.NewMockDisableLegacyFeaturesInterface(t)
	mockAccountSettingsAPI.On("DisableLegacyFeatures").Return(mockDisableLegacyFeatures).Maybe()

	mockEnableIpAccessLists := settings.NewMockEnableIpAccessListsInterface(t)
	mockAccountSettingsAPI.On("EnableIpAccessLists").Return(mockEnableIpAccessLists).Maybe()

	mockEsmEnablementAccount := settings.NewMockEsmEnablementAccountInterface(t)
	mockAccountSettingsAPI.On("EsmEnablementAccount").Return(mockEsmEnablementAccount).Maybe()

	mockLlmProxyPartnerPoweredAccount := settings.NewMockLlmProxyPartnerPoweredAccountInterface(t)
	mockAccountSettingsAPI.On("LlmProxyPartnerPoweredAccount").Return(mockLlmProxyPartnerPoweredAccount).Maybe()

	mockLlmProxyPartnerPoweredEnforce := settings.NewMockLlmProxyPartnerPoweredEnforceInterface(t)
	mockAccountSettingsAPI.On("LlmProxyPartnerPoweredEnforce").Return(mockLlmProxyPartnerPoweredEnforce).Maybe()

	mockPersonalCompute := settings.NewMockPersonalComputeInterface(t)
	mockAccountSettingsAPI.On("PersonalCompute").Return(mockPersonalCompute).Maybe()

	return cli
}

func (m *MockAccountClient) GetMockCspEnablementAccountAPI() *settings.MockCspEnablementAccountInterface {
	api, ok := m.GetMockAccountSettingsAPI().CspEnablementAccount().(*settings.MockCspEnablementAccountInterface)
	if !ok {
		panic(fmt.Sprintf("expected CspEnablementAccount to be *settings.MockCspEnablementAccountInterface, actual was %T", m.GetMockAccountSettingsAPI().CspEnablementAccount()))
	}
	return api
}

func (m *MockAccountClient) GetMockDisableLegacyFeaturesAPI() *settings.MockDisableLegacyFeaturesInterface {
	api, ok := m.GetMockAccountSettingsAPI().DisableLegacyFeatures().(*settings.MockDisableLegacyFeaturesInterface)
	if !ok {
		panic(fmt.Sprintf("expected DisableLegacyFeatures to be *settings.MockDisableLegacyFeaturesInterface, actual was %T", m.GetMockAccountSettingsAPI().DisableLegacyFeatures()))
	}
	return api
}

func (m *MockAccountClient) GetMockEnableIpAccessListsAPI() *settings.MockEnableIpAccessListsInterface {
	api, ok := m.GetMockAccountSettingsAPI().EnableIpAccessLists().(*settings.MockEnableIpAccessListsInterface)
	if !ok {
		panic(fmt.Sprintf("expected EnableIpAccessLists to be *settings.MockEnableIpAccessListsInterface, actual was %T", m.GetMockAccountSettingsAPI().EnableIpAccessLists()))
	}
	return api
}

func (m *MockAccountClient) GetMockEsmEnablementAccountAPI() *settings.MockEsmEnablementAccountInterface {
	api, ok := m.GetMockAccountSettingsAPI().EsmEnablementAccount().(*settings.MockEsmEnablementAccountInterface)
	if !ok {
		panic(fmt.Sprintf("expected EsmEnablementAccount to be *settings.MockEsmEnablementAccountInterface, actual was %T", m.GetMockAccountSettingsAPI().EsmEnablementAccount()))
	}
	return api
}

func (m *MockAccountClient) GetMockLlmProxyPartnerPoweredAccountAPI() *settings.MockLlmProxyPartnerPoweredAccountInterface {
	api, ok := m.GetMockAccountSettingsAPI().LlmProxyPartnerPoweredAccount().(*settings.MockLlmProxyPartnerPoweredAccountInterface)
	if !ok {
		panic(fmt.Sprintf("expected LlmProxyPartnerPoweredAccount to be *settings.MockLlmProxyPartnerPoweredAccountInterface, actual was %T", m.GetMockAccountSettingsAPI().LlmProxyPartnerPoweredAccount()))
	}
	return api
}

func (m *MockAccountClient) GetMockLlmProxyPartnerPoweredEnforceAPI() *settings.MockLlmProxyPartnerPoweredEnforceInterface {
	api, ok := m.GetMockAccountSettingsAPI().LlmProxyPartnerPoweredEnforce().(*settings.MockLlmProxyPartnerPoweredEnforceInterface)
	if !ok {
		panic(fmt.Sprintf("expected LlmProxyPartnerPoweredEnforce to be *settings.MockLlmProxyPartnerPoweredEnforceInterface, actual was %T", m.GetMockAccountSettingsAPI().LlmProxyPartnerPoweredEnforce()))
	}
	return api
}

func (m *MockAccountClient) GetMockPersonalComputeAPI() *settings.MockPersonalComputeInterface {
	api, ok := m.GetMockAccountSettingsAPI().PersonalCompute().(*settings.MockPersonalComputeInterface)
	if !ok {
		panic(fmt.Sprintf("expected PersonalCompute to be *settings.MockPersonalComputeInterface, actual was %T", m.GetMockAccountSettingsAPI().PersonalCompute()))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountAccessControlAPI() *iam.MockAccountAccessControlInterface {
	api, ok := m.AccountClient.AccessControl.(*iam.MockAccountAccessControlInterface)
	if !ok {
		panic(fmt.Sprintf("expected AccessControl to be *iam.MockAccountAccessControlInterface, actual was %T", m.AccountClient.AccessControl))
	}
	return api
}

func (m *MockAccountClient) GetMockBillableUsageAPI() *billing.MockBillableUsageInterface {
	api, ok := m.AccountClient.BillableUsage.(*billing.MockBillableUsageInterface)
	if !ok {
		panic(fmt.Sprintf("expected BillableUsage to be *billing.MockBillableUsageInterface, actual was %T", m.AccountClient.BillableUsage))
	}
	return api
}

func (m *MockAccountClient) GetMockBudgetPolicyAPI() *billing.MockBudgetPolicyInterface {
	api, ok := m.AccountClient.BudgetPolicy.(*billing.MockBudgetPolicyInterface)
	if !ok {
		panic(fmt.Sprintf("expected BudgetPolicy to be *billing.MockBudgetPolicyInterface, actual was %T", m.AccountClient.BudgetPolicy))
	}
	return api
}

func (m *MockAccountClient) GetMockBudgetsAPI() *billing.MockBudgetsInterface {
	api, ok := m.AccountClient.Budgets.(*billing.MockBudgetsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Budgets to be *billing.MockBudgetsInterface, actual was %T", m.AccountClient.Budgets))
	}
	return api
}

func (m *MockAccountClient) GetMockCredentialsAPI() *provisioning.MockCredentialsInterface {
	api, ok := m.AccountClient.Credentials.(*provisioning.MockCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Credentials to be *provisioning.MockCredentialsInterface, actual was %T", m.AccountClient.Credentials))
	}
	return api
}

func (m *MockAccountClient) GetMockCustomAppIntegrationAPI() *oauth2.MockCustomAppIntegrationInterface {
	api, ok := m.AccountClient.CustomAppIntegration.(*oauth2.MockCustomAppIntegrationInterface)
	if !ok {
		panic(fmt.Sprintf("expected CustomAppIntegration to be *oauth2.MockCustomAppIntegrationInterface, actual was %T", m.AccountClient.CustomAppIntegration))
	}
	return api
}

func (m *MockAccountClient) GetMockDisasterRecoveryAPI() *disasterrecovery.MockDisasterRecoveryInterface {
	api, ok := m.AccountClient.DisasterRecovery.(*disasterrecovery.MockDisasterRecoveryInterface)
	if !ok {
		panic(fmt.Sprintf("expected DisasterRecovery to be *disasterrecovery.MockDisasterRecoveryInterface, actual was %T", m.AccountClient.DisasterRecovery))
	}
	return api
}

func (m *MockAccountClient) GetMockEncryptionKeysAPI() *provisioning.MockEncryptionKeysInterface {
	api, ok := m.AccountClient.EncryptionKeys.(*provisioning.MockEncryptionKeysInterface)
	if !ok {
		panic(fmt.Sprintf("expected EncryptionKeys to be *provisioning.MockEncryptionKeysInterface, actual was %T", m.AccountClient.EncryptionKeys))
	}
	return api
}

func (m *MockAccountClient) GetMockEndpointsAPI() *networking.MockEndpointsInterface {
	api, ok := m.AccountClient.Endpoints.(*networking.MockEndpointsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Endpoints to be *networking.MockEndpointsInterface, actual was %T", m.AccountClient.Endpoints))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountFederationPolicyAPI() *oauth2.MockAccountFederationPolicyInterface {
	api, ok := m.AccountClient.FederationPolicy.(*oauth2.MockAccountFederationPolicyInterface)
	if !ok {
		panic(fmt.Sprintf("expected FederationPolicy to be *oauth2.MockAccountFederationPolicyInterface, actual was %T", m.AccountClient.FederationPolicy))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountGroupsV2API() *iam.MockAccountGroupsV2Interface {
	api, ok := m.AccountClient.GroupsV2.(*iam.MockAccountGroupsV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected GroupsV2 to be *iam.MockAccountGroupsV2Interface, actual was %T", m.AccountClient.GroupsV2))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountIamV2API() *iamv2.MockAccountIamV2Interface {
	api, ok := m.AccountClient.IamV2.(*iamv2.MockAccountIamV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected IamV2 to be *iamv2.MockAccountIamV2Interface, actual was %T", m.AccountClient.IamV2))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountIpAccessListsAPI() *settings.MockAccountIpAccessListsInterface {
	api, ok := m.AccountClient.IpAccessLists.(*settings.MockAccountIpAccessListsInterface)
	if !ok {
		panic(fmt.Sprintf("expected IpAccessLists to be *settings.MockAccountIpAccessListsInterface, actual was %T", m.AccountClient.IpAccessLists))
	}
	return api
}

func (m *MockAccountClient) GetMockLogDeliveryAPI() *billing.MockLogDeliveryInterface {
	api, ok := m.AccountClient.LogDelivery.(*billing.MockLogDeliveryInterface)
	if !ok {
		panic(fmt.Sprintf("expected LogDelivery to be *billing.MockLogDeliveryInterface, actual was %T", m.AccountClient.LogDelivery))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountMetastoreAssignmentsAPI() *catalog.MockAccountMetastoreAssignmentsInterface {
	api, ok := m.AccountClient.MetastoreAssignments.(*catalog.MockAccountMetastoreAssignmentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected MetastoreAssignments to be *catalog.MockAccountMetastoreAssignmentsInterface, actual was %T", m.AccountClient.MetastoreAssignments))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountMetastoresAPI() *catalog.MockAccountMetastoresInterface {
	api, ok := m.AccountClient.Metastores.(*catalog.MockAccountMetastoresInterface)
	if !ok {
		panic(fmt.Sprintf("expected Metastores to be *catalog.MockAccountMetastoresInterface, actual was %T", m.AccountClient.Metastores))
	}
	return api
}

func (m *MockAccountClient) GetMockNetworkConnectivityAPI() *settings.MockNetworkConnectivityInterface {
	api, ok := m.AccountClient.NetworkConnectivity.(*settings.MockNetworkConnectivityInterface)
	if !ok {
		panic(fmt.Sprintf("expected NetworkConnectivity to be *settings.MockNetworkConnectivityInterface, actual was %T", m.AccountClient.NetworkConnectivity))
	}
	return api
}

func (m *MockAccountClient) GetMockNetworkPoliciesAPI() *settings.MockNetworkPoliciesInterface {
	api, ok := m.AccountClient.NetworkPolicies.(*settings.MockNetworkPoliciesInterface)
	if !ok {
		panic(fmt.Sprintf("expected NetworkPolicies to be *settings.MockNetworkPoliciesInterface, actual was %T", m.AccountClient.NetworkPolicies))
	}
	return api
}

func (m *MockAccountClient) GetMockNetworksAPI() *provisioning.MockNetworksInterface {
	api, ok := m.AccountClient.Networks.(*provisioning.MockNetworksInterface)
	if !ok {
		panic(fmt.Sprintf("expected Networks to be *provisioning.MockNetworksInterface, actual was %T", m.AccountClient.Networks))
	}
	return api
}

func (m *MockAccountClient) GetMockOAuthPublishedAppsAPI() *oauth2.MockOAuthPublishedAppsInterface {
	api, ok := m.AccountClient.OAuthPublishedApps.(*oauth2.MockOAuthPublishedAppsInterface)
	if !ok {
		panic(fmt.Sprintf("expected OAuthPublishedApps to be *oauth2.MockOAuthPublishedAppsInterface, actual was %T", m.AccountClient.OAuthPublishedApps))
	}
	return api
}

func (m *MockAccountClient) GetMockPrivateAccessAPI() *provisioning.MockPrivateAccessInterface {
	api, ok := m.AccountClient.PrivateAccess.(*provisioning.MockPrivateAccessInterface)
	if !ok {
		panic(fmt.Sprintf("expected PrivateAccess to be *provisioning.MockPrivateAccessInterface, actual was %T", m.AccountClient.PrivateAccess))
	}
	return api
}

func (m *MockAccountClient) GetMockPrivateNetworkGatewaysAPI() *networking.MockPrivateNetworkGatewaysInterface {
	api, ok := m.AccountClient.PrivateNetworkGateways.(*networking.MockPrivateNetworkGatewaysInterface)
	if !ok {
		panic(fmt.Sprintf("expected PrivateNetworkGateways to be *networking.MockPrivateNetworkGatewaysInterface, actual was %T", m.AccountClient.PrivateNetworkGateways))
	}
	return api
}

func (m *MockAccountClient) GetMockPublishedAppIntegrationAPI() *oauth2.MockPublishedAppIntegrationInterface {
	api, ok := m.AccountClient.PublishedAppIntegration.(*oauth2.MockPublishedAppIntegrationInterface)
	if !ok {
		panic(fmt.Sprintf("expected PublishedAppIntegration to be *oauth2.MockPublishedAppIntegrationInterface, actual was %T", m.AccountClient.PublishedAppIntegration))
	}
	return api
}

func (m *MockAccountClient) GetMockServicePrincipalFederationPolicyAPI() *oauth2.MockServicePrincipalFederationPolicyInterface {
	api, ok := m.AccountClient.ServicePrincipalFederationPolicy.(*oauth2.MockServicePrincipalFederationPolicyInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipalFederationPolicy to be *oauth2.MockServicePrincipalFederationPolicyInterface, actual was %T", m.AccountClient.ServicePrincipalFederationPolicy))
	}
	return api
}

func (m *MockAccountClient) GetMockServicePrincipalSecretsAPI() *oauth2.MockServicePrincipalSecretsInterface {
	api, ok := m.AccountClient.ServicePrincipalSecrets.(*oauth2.MockServicePrincipalSecretsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipalSecrets to be *oauth2.MockServicePrincipalSecretsInterface, actual was %T", m.AccountClient.ServicePrincipalSecrets))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountServicePrincipalsV2API() *iam.MockAccountServicePrincipalsV2Interface {
	api, ok := m.AccountClient.ServicePrincipalsV2.(*iam.MockAccountServicePrincipalsV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipalsV2 to be *iam.MockAccountServicePrincipalsV2Interface, actual was %T", m.AccountClient.ServicePrincipalsV2))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountSettingsAPI() *settings.MockAccountSettingsInterface {
	api, ok := m.AccountClient.Settings.(*settings.MockAccountSettingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Settings to be *settings.MockAccountSettingsInterface, actual was %T", m.AccountClient.Settings))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountSettingsV2API() *settingsv2.MockAccountSettingsV2Interface {
	api, ok := m.AccountClient.SettingsV2.(*settingsv2.MockAccountSettingsV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected SettingsV2 to be *settingsv2.MockAccountSettingsV2Interface, actual was %T", m.AccountClient.SettingsV2))
	}
	return api
}

func (m *MockAccountClient) GetMockStorageAPI() *provisioning.MockStorageInterface {
	api, ok := m.AccountClient.Storage.(*provisioning.MockStorageInterface)
	if !ok {
		panic(fmt.Sprintf("expected Storage to be *provisioning.MockStorageInterface, actual was %T", m.AccountClient.Storage))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountStorageCredentialsAPI() *catalog.MockAccountStorageCredentialsInterface {
	api, ok := m.AccountClient.StorageCredentials.(*catalog.MockAccountStorageCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected StorageCredentials to be *catalog.MockAccountStorageCredentialsInterface, actual was %T", m.AccountClient.StorageCredentials))
	}
	return api
}

func (m *MockAccountClient) GetMockUsageDashboardsAPI() *billing.MockUsageDashboardsInterface {
	api, ok := m.AccountClient.UsageDashboards.(*billing.MockUsageDashboardsInterface)
	if !ok {
		panic(fmt.Sprintf("expected UsageDashboards to be *billing.MockUsageDashboardsInterface, actual was %T", m.AccountClient.UsageDashboards))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountUsersV2API() *iam.MockAccountUsersV2Interface {
	api, ok := m.AccountClient.UsersV2.(*iam.MockAccountUsersV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected UsersV2 to be *iam.MockAccountUsersV2Interface, actual was %T", m.AccountClient.UsersV2))
	}
	return api
}

func (m *MockAccountClient) GetMockVpcEndpointsAPI() *provisioning.MockVpcEndpointsInterface {
	api, ok := m.AccountClient.VpcEndpoints.(*provisioning.MockVpcEndpointsInterface)
	if !ok {
		panic(fmt.Sprintf("expected VpcEndpoints to be *provisioning.MockVpcEndpointsInterface, actual was %T", m.AccountClient.VpcEndpoints))
	}
	return api
}

func (m *MockAccountClient) GetMockWorkspaceAssignmentAPI() *iam.MockWorkspaceAssignmentInterface {
	api, ok := m.AccountClient.WorkspaceAssignment.(*iam.MockWorkspaceAssignmentInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceAssignment to be *iam.MockWorkspaceAssignmentInterface, actual was %T", m.AccountClient.WorkspaceAssignment))
	}
	return api
}

func (m *MockAccountClient) GetMockWorkspaceNetworkConfigurationAPI() *settings.MockWorkspaceNetworkConfigurationInterface {
	api, ok := m.AccountClient.WorkspaceNetworkConfiguration.(*settings.MockWorkspaceNetworkConfigurationInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceNetworkConfiguration to be *settings.MockWorkspaceNetworkConfigurationInterface, actual was %T", m.AccountClient.WorkspaceNetworkConfiguration))
	}
	return api
}

func (m *MockAccountClient) GetMockWorkspacesAPI() *provisioning.MockWorkspacesInterface {
	api, ok := m.AccountClient.Workspaces.(*provisioning.MockWorkspacesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Workspaces to be *provisioning.MockWorkspacesInterface, actual was %T", m.AccountClient.Workspaces))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountGroupsAPI() *iam.MockAccountGroupsInterface {
	api, ok := m.AccountClient.Groups.(*iam.MockAccountGroupsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Groups to be *iam.MockAccountGroupsInterface, actual was %T", m.AccountClient.Groups))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountServicePrincipalsAPI() *iam.MockAccountServicePrincipalsInterface {
	api, ok := m.AccountClient.ServicePrincipals.(*iam.MockAccountServicePrincipalsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipals to be *iam.MockAccountServicePrincipalsInterface, actual was %T", m.AccountClient.ServicePrincipals))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountUsersAPI() *iam.MockAccountUsersInterface {
	api, ok := m.AccountClient.Users.(*iam.MockAccountUsersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Users to be *iam.MockAccountUsersInterface, actual was %T", m.AccountClient.Users))
	}
	return api
}
