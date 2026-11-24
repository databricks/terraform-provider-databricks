// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go"

	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/agentbricks"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/agentkit"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/aifunctions"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/aisearch"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/apps"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/bundledeployments"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/cleanrooms"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/compute"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/computemigrations"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/dashboards"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/database"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/dataclassification"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/dataquality"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/domains"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/environments"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/files"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/genieagents"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iamv2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/jobs"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/knowledgeassistants"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/lakewatch"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/marketplace"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/ml"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/oauth2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/pipelines"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/postgres"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/qualitymonitorv2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/sandbox"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/serving"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settings"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settingsv2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/sharing"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/softwareartifacts"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/sql"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/statementexecutionv2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/supervisoragents"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/tags"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/vectorsearch"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/workspace"
	"github.com/stretchr/testify/mock"
)

type MockWorkspaceClient struct {
	WorkspaceClient *databricks.WorkspaceClient
}

// NewMockWorkspaceClient creates new mocked version of Databricks SDK client for Workspaces
// which can be used for testing.
func NewMockWorkspaceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkspaceClient {
	cli := &MockWorkspaceClient{
		WorkspaceClient: &databricks.WorkspaceClient{
			Config: nil,

			AccessControl:                       iam.NewMockAccessControlInterface(t),
			AccountAccessControlProxy:           iam.NewMockAccountAccessControlProxyInterface(t),
			AgentBricks:                         agentbricks.NewMockAgentBricksInterface(t),
			AgentKit:                            agentkit.NewMockAgentKitInterface(t),
			AiFunctions:                         aifunctions.NewMockAiFunctionsInterface(t),
			AiGateway:                           catalog.NewMockAiGatewayInterface(t),
			AiSearch:                            aisearch.NewMockAiSearchInterface(t),
			Alerts:                              sql.NewMockAlertsInterface(t),
			AlertsLegacy:                        sql.NewMockAlertsLegacyInterface(t),
			AlertsV2:                            sql.NewMockAlertsV2Interface(t),
			Apps:                                apps.NewMockAppsInterface(t),
			AppsSettings:                        apps.NewMockAppsSettingsInterface(t),
			ArtifactAllowlists:                  catalog.NewMockArtifactAllowlistsInterface(t),
			AssumeGroup:                         iam.NewMockAssumeGroupInterface(t),
			BundleDeployments:                   bundledeployments.NewMockBundleDeploymentsInterface(t),
			Catalogs:                            catalog.NewMockCatalogsInterface(t),
			CleanRoomAssetRevisions:             cleanrooms.NewMockCleanRoomAssetRevisionsInterface(t),
			CleanRoomAssets:                     cleanrooms.NewMockCleanRoomAssetsInterface(t),
			CleanRoomAutoApprovalRules:          cleanrooms.NewMockCleanRoomAutoApprovalRulesInterface(t),
			CleanRoomTaskRuns:                   cleanrooms.NewMockCleanRoomTaskRunsInterface(t),
			CleanRooms:                          cleanrooms.NewMockCleanRoomsInterface(t),
			ClusterPolicies:                     compute.NewMockClusterPoliciesInterface(t),
			Clusters:                            compute.NewMockClustersInterface(t),
			CommandExecution:                    compute.NewMockCommandExecutionInterface(t),
			ComputeMigrations:                   computemigrations.NewMockComputeMigrationsInterface(t),
			Connections:                         catalog.NewMockConnectionsInterface(t),
			ConsumerFulfillments:                marketplace.NewMockConsumerFulfillmentsInterface(t),
			ConsumerInstallations:               marketplace.NewMockConsumerInstallationsInterface(t),
			ConsumerListings:                    marketplace.NewMockConsumerListingsInterface(t),
			ConsumerPersonalizationRequests:     marketplace.NewMockConsumerPersonalizationRequestsInterface(t),
			ConsumerProviders:                   marketplace.NewMockConsumerProvidersInterface(t),
			Credentials:                         catalog.NewMockCredentialsInterface(t),
			CredentialsManager:                  settings.NewMockCredentialsManagerInterface(t),
			CurrentUser:                         iam.NewMockCurrentUserInterface(t),
			DashboardWidgets:                    sql.NewMockDashboardWidgetsInterface(t),
			Dashboards:                          sql.NewMockDashboardsInterface(t),
			DataClassification:                  dataclassification.NewMockDataClassificationInterface(t),
			DataQuality:                         dataquality.NewMockDataQualityInterface(t),
			DataSources:                         sql.NewMockDataSourcesInterface(t),
			Database:                            database.NewMockDatabaseInterface(t),
			Dbfs:                                files.NewMockDbfsInterface(t),
			DbsqlPermissions:                    sql.NewMockDbsqlPermissionsInterface(t),
			Domains:                             domains.NewMockDomainsInterface(t),
			EntityTagAssignments:                catalog.NewMockEntityTagAssignmentsInterface(t),
			Environments:                        environments.NewMockEnvironmentsInterface(t),
			Experiments:                         ml.NewMockExperimentsInterface(t),
			ExternalLineage:                     catalog.NewMockExternalLineageInterface(t),
			ExternalLocations:                   catalog.NewMockExternalLocationsInterface(t),
			ExternalMetadata:                    catalog.NewMockExternalMetadataInterface(t),
			FeatureEngineering:                  ml.NewMockFeatureEngineeringInterface(t),
			FeatureStore:                        ml.NewMockFeatureStoreInterface(t),
			Files:                               files.NewMockFilesInterface(t),
			Forecasting:                         ml.NewMockForecastingInterface(t),
			Functions:                           catalog.NewMockFunctionsInterface(t),
			Genie:                               dashboards.NewMockGenieInterface(t),
			GenieAgents:                         genieagents.NewMockGenieAgentsInterface(t),
			GitCredentials:                      workspace.NewMockGitCredentialsInterface(t),
			GlobalInitScripts:                   compute.NewMockGlobalInitScriptsInterface(t),
			Grants:                              catalog.NewMockGrantsInterface(t),
			GroupsV2:                            iam.NewMockGroupsV2Interface(t),
			InstancePools:                       compute.NewMockInstancePoolsInterface(t),
			InstanceProfiles:                    compute.NewMockInstanceProfilesInterface(t),
			IpAccessLists:                       settings.NewMockIpAccessListsInterface(t),
			Jobs:                                jobs.NewMockJobsInterface(t),
			KnowledgeAssistants:                 knowledgeassistants.NewMockKnowledgeAssistantsInterface(t),
			Lakeview:                            dashboards.NewMockLakeviewInterface(t),
			LakeviewEmbedded:                    dashboards.NewMockLakeviewEmbeddedInterface(t),
			Lakewatch:                           lakewatch.NewMockLakewatchInterface(t),
			Libraries:                           compute.NewMockLibrariesInterface(t),
			MaterializedFeatures:                ml.NewMockMaterializedFeaturesInterface(t),
			Metastores:                          catalog.NewMockMetastoresInterface(t),
			ModelRegistry:                       ml.NewMockModelRegistryInterface(t),
			ModelVersions:                       catalog.NewMockModelVersionsInterface(t),
			NotificationDestinations:            settings.NewMockNotificationDestinationsInterface(t),
			OnlineTables:                        catalog.NewMockOnlineTablesInterface(t),
			PermissionMigration:                 iam.NewMockPermissionMigrationInterface(t),
			Permissions:                         iam.NewMockPermissionsInterface(t),
			Pipelines:                           pipelines.NewMockPipelinesInterface(t),
			Policies:                            catalog.NewMockPoliciesInterface(t),
			PolicyComplianceForClusters:         compute.NewMockPolicyComplianceForClustersInterface(t),
			PolicyComplianceForJobs:             jobs.NewMockPolicyComplianceForJobsInterface(t),
			PolicyFamilies:                      compute.NewMockPolicyFamiliesInterface(t),
			Postgres:                            postgres.NewMockPostgresInterface(t),
			ProviderExchangeFilters:             marketplace.NewMockProviderExchangeFiltersInterface(t),
			ProviderExchanges:                   marketplace.NewMockProviderExchangesInterface(t),
			ProviderFiles:                       marketplace.NewMockProviderFilesInterface(t),
			ProviderListings:                    marketplace.NewMockProviderListingsInterface(t),
			ProviderPersonalizationRequests:     marketplace.NewMockProviderPersonalizationRequestsInterface(t),
			ProviderProviderAnalyticsDashboards: marketplace.NewMockProviderProviderAnalyticsDashboardsInterface(t),
			ProviderProviders:                   marketplace.NewMockProviderProvidersInterface(t),
			Providers:                           sharing.NewMockProvidersInterface(t),
			QualityMonitorV2:                    qualitymonitorv2.NewMockQualityMonitorV2Interface(t),
			QualityMonitors:                     catalog.NewMockQualityMonitorsInterface(t),
			Queries:                             sql.NewMockQueriesInterface(t),
			QueriesLegacy:                       sql.NewMockQueriesLegacyInterface(t),
			QueryHistory:                        sql.NewMockQueryHistoryInterface(t),
			QueryVisualizations:                 sql.NewMockQueryVisualizationsInterface(t),
			QueryVisualizationsLegacy:           sql.NewMockQueryVisualizationsLegacyInterface(t),
			RecipientActivation:                 sharing.NewMockRecipientActivationInterface(t),
			RecipientFederationPolicies:         sharing.NewMockRecipientFederationPoliciesInterface(t),
			Recipients:                          sharing.NewMockRecipientsInterface(t),
			RedashConfig:                        sql.NewMockRedashConfigInterface(t),
			RegisteredModels:                    catalog.NewMockRegisteredModelsInterface(t),
			Repos:                               workspace.NewMockReposInterface(t),
			ResourceQuotas:                      catalog.NewMockResourceQuotasInterface(t),
			Rfa:                                 catalog.NewMockRfaInterface(t),
			Sandbox:                             sandbox.NewMockSandboxInterface(t),
			Schemas:                             catalog.NewMockSchemasInterface(t),
			Secrets:                             workspace.NewMockSecretsInterface(t),
			SecretsUc:                           catalog.NewMockSecretsUcInterface(t),
			ServerlessComputeService:            compute.NewMockServerlessComputeServiceInterface(t),
			ServicePrincipalSecretsProxy:        oauth2.NewMockServicePrincipalSecretsProxyInterface(t),
			ServicePrincipalsV2:                 iam.NewMockServicePrincipalsV2Interface(t),
			ServingEndpoints:                    serving.NewMockServingEndpointsInterface(t),
			ServingEndpointsDataPlane:           serving.NewMockServingEndpointsDataPlaneInterface(t),
			Settings:                            settings.NewMockSettingsInterface(t),
			Shares:                              sharing.NewMockSharesInterface(t),
			SoftwareArtifacts:                   softwareartifacts.NewMockSoftwareArtifactsInterface(t),
			StatementExecution:                  sql.NewMockStatementExecutionInterface(t),
			StatementExecutionV2:                statementexecutionv2.NewMockStatementExecutionV2Interface(t),
			StorageCredentials:                  catalog.NewMockStorageCredentialsInterface(t),
			SupervisorAgents:                    supervisoragents.NewMockSupervisorAgentsInterface(t),
			SystemSchemas:                       catalog.NewMockSystemSchemasInterface(t),
			TableConstraints:                    catalog.NewMockTableConstraintsInterface(t),
			Tables:                              catalog.NewMockTablesInterface(t),
			TagPolicies:                         tags.NewMockTagPoliciesInterface(t),
			TemporaryPathCredentials:            catalog.NewMockTemporaryPathCredentialsInterface(t),
			TemporaryTableCredentials:           catalog.NewMockTemporaryTableCredentialsInterface(t),
			TemporaryVolumeCredentials:          catalog.NewMockTemporaryVolumeCredentialsInterface(t),
			TokenManagement:                     settings.NewMockTokenManagementInterface(t),
			Tokens:                              settings.NewMockTokensInterface(t),
			UsersV2:                             iam.NewMockUsersV2Interface(t),
			VectorSearchEndpoints:               vectorsearch.NewMockVectorSearchEndpointsInterface(t),
			VectorSearchIndexes:                 vectorsearch.NewMockVectorSearchIndexesInterface(t),
			Volumes:                             catalog.NewMockVolumesInterface(t),
			Warehouses:                          sql.NewMockWarehousesInterface(t),
			Workspace:                           workspace.NewMockWorkspaceInterface(t),
			WorkspaceBindings:                   catalog.NewMockWorkspaceBindingsInterface(t),
			WorkspaceConf:                       settings.NewMockWorkspaceConfInterface(t),
			WorkspaceEntityTagAssignments:       tags.NewMockWorkspaceEntityTagAssignmentsInterface(t),
			WorkspaceIamV2:                      iamv2.NewMockWorkspaceIamV2Interface(t),
			WorkspaceSettingsV2:                 settingsv2.NewMockWorkspaceSettingsV2Interface(t),
			Groups:                              iam.NewMockGroupsInterface(t),
			ServicePrincipals:                   iam.NewMockServicePrincipalsInterface(t),
			Users:                               iam.NewMockUsersInterface(t),
		},
	}

	mocksettingsAPI := cli.GetMockSettingsAPI()

	mockaibiDashboardEmbeddingAccessPolicy := settings.NewMockAibiDashboardEmbeddingAccessPolicyInterface(t)
	mocksettingsAPI.On("AibiDashboardEmbeddingAccessPolicy").Return(mockaibiDashboardEmbeddingAccessPolicy).Maybe()

	mockaibiDashboardEmbeddingApprovedDomains := settings.NewMockAibiDashboardEmbeddingApprovedDomainsInterface(t)
	mocksettingsAPI.On("AibiDashboardEmbeddingApprovedDomains").Return(mockaibiDashboardEmbeddingApprovedDomains).Maybe()

	mockautomaticClusterUpdate := settings.NewMockAutomaticClusterUpdateInterface(t)
	mocksettingsAPI.On("AutomaticClusterUpdate").Return(mockautomaticClusterUpdate).Maybe()

	mockcomplianceSecurityProfile := settings.NewMockComplianceSecurityProfileInterface(t)
	mocksettingsAPI.On("ComplianceSecurityProfile").Return(mockcomplianceSecurityProfile).Maybe()

	mockdashboardEmailSubscriptions := settings.NewMockDashboardEmailSubscriptionsInterface(t)
	mocksettingsAPI.On("DashboardEmailSubscriptions").Return(mockdashboardEmailSubscriptions).Maybe()

	mockdefaultNamespace := settings.NewMockDefaultNamespaceInterface(t)
	mocksettingsAPI.On("DefaultNamespace").Return(mockdefaultNamespace).Maybe()

	mockdefaultWarehouseId := settings.NewMockDefaultWarehouseIdInterface(t)
	mocksettingsAPI.On("DefaultWarehouseId").Return(mockdefaultWarehouseId).Maybe()

	mockdisableLegacyAccess := settings.NewMockDisableLegacyAccessInterface(t)
	mocksettingsAPI.On("DisableLegacyAccess").Return(mockdisableLegacyAccess).Maybe()

	mockdisableLegacyDbfs := settings.NewMockDisableLegacyDbfsInterface(t)
	mocksettingsAPI.On("DisableLegacyDbfs").Return(mockdisableLegacyDbfs).Maybe()

	mockenableExportNotebook := settings.NewMockEnableExportNotebookInterface(t)
	mocksettingsAPI.On("EnableExportNotebook").Return(mockenableExportNotebook).Maybe()

	mockenableNotebookTableClipboard := settings.NewMockEnableNotebookTableClipboardInterface(t)
	mocksettingsAPI.On("EnableNotebookTableClipboard").Return(mockenableNotebookTableClipboard).Maybe()

	mockenableResultsDownloading := settings.NewMockEnableResultsDownloadingInterface(t)
	mocksettingsAPI.On("EnableResultsDownloading").Return(mockenableResultsDownloading).Maybe()

	mockenhancedSecurityMonitoring := settings.NewMockEnhancedSecurityMonitoringInterface(t)
	mocksettingsAPI.On("EnhancedSecurityMonitoring").Return(mockenhancedSecurityMonitoring).Maybe()

	mockllmProxyPartnerPoweredWorkspace := settings.NewMockLlmProxyPartnerPoweredWorkspaceInterface(t)
	mocksettingsAPI.On("LlmProxyPartnerPoweredWorkspace").Return(mockllmProxyPartnerPoweredWorkspace).Maybe()

	mockrestrictWorkspaceAdmins := settings.NewMockRestrictWorkspaceAdminsInterface(t)
	mocksettingsAPI.On("RestrictWorkspaceAdmins").Return(mockrestrictWorkspaceAdmins).Maybe()

	mocksqlResultsDownload := settings.NewMockSqlResultsDownloadInterface(t)
	mocksettingsAPI.On("SqlResultsDownload").Return(mocksqlResultsDownload).Maybe()

	return cli
}

func (m *MockWorkspaceClient) GetMockAibiDashboardEmbeddingAccessPolicyAPI() *settings.MockAibiDashboardEmbeddingAccessPolicyInterface {
	api, ok := m.GetMockSettingsAPI().AibiDashboardEmbeddingAccessPolicy().(*settings.MockAibiDashboardEmbeddingAccessPolicyInterface)
	if !ok {
		panic(fmt.Sprintf("expected AibiDashboardEmbeddingAccessPolicy to be *settings.MockAibiDashboardEmbeddingAccessPolicyInterface, actual was %T", m.GetMockSettingsAPI().AibiDashboardEmbeddingAccessPolicy()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAibiDashboardEmbeddingApprovedDomainsAPI() *settings.MockAibiDashboardEmbeddingApprovedDomainsInterface {
	api, ok := m.GetMockSettingsAPI().AibiDashboardEmbeddingApprovedDomains().(*settings.MockAibiDashboardEmbeddingApprovedDomainsInterface)
	if !ok {
		panic(fmt.Sprintf("expected AibiDashboardEmbeddingApprovedDomains to be *settings.MockAibiDashboardEmbeddingApprovedDomainsInterface, actual was %T", m.GetMockSettingsAPI().AibiDashboardEmbeddingApprovedDomains()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAutomaticClusterUpdateAPI() *settings.MockAutomaticClusterUpdateInterface {
	api, ok := m.GetMockSettingsAPI().AutomaticClusterUpdate().(*settings.MockAutomaticClusterUpdateInterface)
	if !ok {
		panic(fmt.Sprintf("expected AutomaticClusterUpdate to be *settings.MockAutomaticClusterUpdateInterface, actual was %T", m.GetMockSettingsAPI().AutomaticClusterUpdate()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockComplianceSecurityProfileAPI() *settings.MockComplianceSecurityProfileInterface {
	api, ok := m.GetMockSettingsAPI().ComplianceSecurityProfile().(*settings.MockComplianceSecurityProfileInterface)
	if !ok {
		panic(fmt.Sprintf("expected ComplianceSecurityProfile to be *settings.MockComplianceSecurityProfileInterface, actual was %T", m.GetMockSettingsAPI().ComplianceSecurityProfile()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDashboardEmailSubscriptionsAPI() *settings.MockDashboardEmailSubscriptionsInterface {
	api, ok := m.GetMockSettingsAPI().DashboardEmailSubscriptions().(*settings.MockDashboardEmailSubscriptionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected DashboardEmailSubscriptions to be *settings.MockDashboardEmailSubscriptionsInterface, actual was %T", m.GetMockSettingsAPI().DashboardEmailSubscriptions()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDefaultNamespaceAPI() *settings.MockDefaultNamespaceInterface {
	api, ok := m.GetMockSettingsAPI().DefaultNamespace().(*settings.MockDefaultNamespaceInterface)
	if !ok {
		panic(fmt.Sprintf("expected DefaultNamespace to be *settings.MockDefaultNamespaceInterface, actual was %T", m.GetMockSettingsAPI().DefaultNamespace()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDefaultWarehouseIdAPI() *settings.MockDefaultWarehouseIdInterface {
	api, ok := m.GetMockSettingsAPI().DefaultWarehouseId().(*settings.MockDefaultWarehouseIdInterface)
	if !ok {
		panic(fmt.Sprintf("expected DefaultWarehouseId to be *settings.MockDefaultWarehouseIdInterface, actual was %T", m.GetMockSettingsAPI().DefaultWarehouseId()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDisableLegacyAccessAPI() *settings.MockDisableLegacyAccessInterface {
	api, ok := m.GetMockSettingsAPI().DisableLegacyAccess().(*settings.MockDisableLegacyAccessInterface)
	if !ok {
		panic(fmt.Sprintf("expected DisableLegacyAccess to be *settings.MockDisableLegacyAccessInterface, actual was %T", m.GetMockSettingsAPI().DisableLegacyAccess()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDisableLegacyDbfsAPI() *settings.MockDisableLegacyDbfsInterface {
	api, ok := m.GetMockSettingsAPI().DisableLegacyDbfs().(*settings.MockDisableLegacyDbfsInterface)
	if !ok {
		panic(fmt.Sprintf("expected DisableLegacyDbfs to be *settings.MockDisableLegacyDbfsInterface, actual was %T", m.GetMockSettingsAPI().DisableLegacyDbfs()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockEnableExportNotebookAPI() *settings.MockEnableExportNotebookInterface {
	api, ok := m.GetMockSettingsAPI().EnableExportNotebook().(*settings.MockEnableExportNotebookInterface)
	if !ok {
		panic(fmt.Sprintf("expected EnableExportNotebook to be *settings.MockEnableExportNotebookInterface, actual was %T", m.GetMockSettingsAPI().EnableExportNotebook()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockEnableNotebookTableClipboardAPI() *settings.MockEnableNotebookTableClipboardInterface {
	api, ok := m.GetMockSettingsAPI().EnableNotebookTableClipboard().(*settings.MockEnableNotebookTableClipboardInterface)
	if !ok {
		panic(fmt.Sprintf("expected EnableNotebookTableClipboard to be *settings.MockEnableNotebookTableClipboardInterface, actual was %T", m.GetMockSettingsAPI().EnableNotebookTableClipboard()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockEnableResultsDownloadingAPI() *settings.MockEnableResultsDownloadingInterface {
	api, ok := m.GetMockSettingsAPI().EnableResultsDownloading().(*settings.MockEnableResultsDownloadingInterface)
	if !ok {
		panic(fmt.Sprintf("expected EnableResultsDownloading to be *settings.MockEnableResultsDownloadingInterface, actual was %T", m.GetMockSettingsAPI().EnableResultsDownloading()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockEnhancedSecurityMonitoringAPI() *settings.MockEnhancedSecurityMonitoringInterface {
	api, ok := m.GetMockSettingsAPI().EnhancedSecurityMonitoring().(*settings.MockEnhancedSecurityMonitoringInterface)
	if !ok {
		panic(fmt.Sprintf("expected EnhancedSecurityMonitoring to be *settings.MockEnhancedSecurityMonitoringInterface, actual was %T", m.GetMockSettingsAPI().EnhancedSecurityMonitoring()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLlmProxyPartnerPoweredWorkspaceAPI() *settings.MockLlmProxyPartnerPoweredWorkspaceInterface {
	api, ok := m.GetMockSettingsAPI().LlmProxyPartnerPoweredWorkspace().(*settings.MockLlmProxyPartnerPoweredWorkspaceInterface)
	if !ok {
		panic(fmt.Sprintf("expected LlmProxyPartnerPoweredWorkspace to be *settings.MockLlmProxyPartnerPoweredWorkspaceInterface, actual was %T", m.GetMockSettingsAPI().LlmProxyPartnerPoweredWorkspace()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRestrictWorkspaceAdminsAPI() *settings.MockRestrictWorkspaceAdminsInterface {
	api, ok := m.GetMockSettingsAPI().RestrictWorkspaceAdmins().(*settings.MockRestrictWorkspaceAdminsInterface)
	if !ok {
		panic(fmt.Sprintf("expected RestrictWorkspaceAdmins to be *settings.MockRestrictWorkspaceAdminsInterface, actual was %T", m.GetMockSettingsAPI().RestrictWorkspaceAdmins()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSqlResultsDownloadAPI() *settings.MockSqlResultsDownloadInterface {
	api, ok := m.GetMockSettingsAPI().SqlResultsDownload().(*settings.MockSqlResultsDownloadInterface)
	if !ok {
		panic(fmt.Sprintf("expected SqlResultsDownload to be *settings.MockSqlResultsDownloadInterface, actual was %T", m.GetMockSettingsAPI().SqlResultsDownload()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAccessControlAPI() *iam.MockAccessControlInterface {
	api, ok := m.WorkspaceClient.AccessControl.(*iam.MockAccessControlInterface)
	if !ok {
		panic(fmt.Sprintf("expected AccessControl to be *iam.MockAccessControlInterface, actual was %T", m.WorkspaceClient.AccessControl))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAccountAccessControlProxyAPI() *iam.MockAccountAccessControlProxyInterface {
	api, ok := m.WorkspaceClient.AccountAccessControlProxy.(*iam.MockAccountAccessControlProxyInterface)
	if !ok {
		panic(fmt.Sprintf("expected AccountAccessControlProxy to be *iam.MockAccountAccessControlProxyInterface, actual was %T", m.WorkspaceClient.AccountAccessControlProxy))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAgentBricksAPI() *agentbricks.MockAgentBricksInterface {
	api, ok := m.WorkspaceClient.AgentBricks.(*agentbricks.MockAgentBricksInterface)
	if !ok {
		panic(fmt.Sprintf("expected AgentBricks to be *agentbricks.MockAgentBricksInterface, actual was %T", m.WorkspaceClient.AgentBricks))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAgentKitAPI() *agentkit.MockAgentKitInterface {
	api, ok := m.WorkspaceClient.AgentKit.(*agentkit.MockAgentKitInterface)
	if !ok {
		panic(fmt.Sprintf("expected AgentKit to be *agentkit.MockAgentKitInterface, actual was %T", m.WorkspaceClient.AgentKit))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAiFunctionsAPI() *aifunctions.MockAiFunctionsInterface {
	api, ok := m.WorkspaceClient.AiFunctions.(*aifunctions.MockAiFunctionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected AiFunctions to be *aifunctions.MockAiFunctionsInterface, actual was %T", m.WorkspaceClient.AiFunctions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAiGatewayAPI() *catalog.MockAiGatewayInterface {
	api, ok := m.WorkspaceClient.AiGateway.(*catalog.MockAiGatewayInterface)
	if !ok {
		panic(fmt.Sprintf("expected AiGateway to be *catalog.MockAiGatewayInterface, actual was %T", m.WorkspaceClient.AiGateway))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAiSearchAPI() *aisearch.MockAiSearchInterface {
	api, ok := m.WorkspaceClient.AiSearch.(*aisearch.MockAiSearchInterface)
	if !ok {
		panic(fmt.Sprintf("expected AiSearch to be *aisearch.MockAiSearchInterface, actual was %T", m.WorkspaceClient.AiSearch))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAlertsAPI() *sql.MockAlertsInterface {
	api, ok := m.WorkspaceClient.Alerts.(*sql.MockAlertsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Alerts to be *sql.MockAlertsInterface, actual was %T", m.WorkspaceClient.Alerts))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAlertsLegacyAPI() *sql.MockAlertsLegacyInterface {
	api, ok := m.WorkspaceClient.AlertsLegacy.(*sql.MockAlertsLegacyInterface)
	if !ok {
		panic(fmt.Sprintf("expected AlertsLegacy to be *sql.MockAlertsLegacyInterface, actual was %T", m.WorkspaceClient.AlertsLegacy))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAlertsV2API() *sql.MockAlertsV2Interface {
	api, ok := m.WorkspaceClient.AlertsV2.(*sql.MockAlertsV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected AlertsV2 to be *sql.MockAlertsV2Interface, actual was %T", m.WorkspaceClient.AlertsV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAppsAPI() *apps.MockAppsInterface {
	api, ok := m.WorkspaceClient.Apps.(*apps.MockAppsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Apps to be *apps.MockAppsInterface, actual was %T", m.WorkspaceClient.Apps))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAppsSettingsAPI() *apps.MockAppsSettingsInterface {
	api, ok := m.WorkspaceClient.AppsSettings.(*apps.MockAppsSettingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected AppsSettings to be *apps.MockAppsSettingsInterface, actual was %T", m.WorkspaceClient.AppsSettings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockArtifactAllowlistsAPI() *catalog.MockArtifactAllowlistsInterface {
	api, ok := m.WorkspaceClient.ArtifactAllowlists.(*catalog.MockArtifactAllowlistsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ArtifactAllowlists to be *catalog.MockArtifactAllowlistsInterface, actual was %T", m.WorkspaceClient.ArtifactAllowlists))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAssumeGroupAPI() *iam.MockAssumeGroupInterface {
	api, ok := m.WorkspaceClient.AssumeGroup.(*iam.MockAssumeGroupInterface)
	if !ok {
		panic(fmt.Sprintf("expected AssumeGroup to be *iam.MockAssumeGroupInterface, actual was %T", m.WorkspaceClient.AssumeGroup))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockBundleDeploymentsAPI() *bundledeployments.MockBundleDeploymentsInterface {
	api, ok := m.WorkspaceClient.BundleDeployments.(*bundledeployments.MockBundleDeploymentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected BundleDeployments to be *bundledeployments.MockBundleDeploymentsInterface, actual was %T", m.WorkspaceClient.BundleDeployments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCatalogsAPI() *catalog.MockCatalogsInterface {
	api, ok := m.WorkspaceClient.Catalogs.(*catalog.MockCatalogsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Catalogs to be *catalog.MockCatalogsInterface, actual was %T", m.WorkspaceClient.Catalogs))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCleanRoomAssetRevisionsAPI() *cleanrooms.MockCleanRoomAssetRevisionsInterface {
	api, ok := m.WorkspaceClient.CleanRoomAssetRevisions.(*cleanrooms.MockCleanRoomAssetRevisionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected CleanRoomAssetRevisions to be *cleanrooms.MockCleanRoomAssetRevisionsInterface, actual was %T", m.WorkspaceClient.CleanRoomAssetRevisions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCleanRoomAssetsAPI() *cleanrooms.MockCleanRoomAssetsInterface {
	api, ok := m.WorkspaceClient.CleanRoomAssets.(*cleanrooms.MockCleanRoomAssetsInterface)
	if !ok {
		panic(fmt.Sprintf("expected CleanRoomAssets to be *cleanrooms.MockCleanRoomAssetsInterface, actual was %T", m.WorkspaceClient.CleanRoomAssets))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCleanRoomAutoApprovalRulesAPI() *cleanrooms.MockCleanRoomAutoApprovalRulesInterface {
	api, ok := m.WorkspaceClient.CleanRoomAutoApprovalRules.(*cleanrooms.MockCleanRoomAutoApprovalRulesInterface)
	if !ok {
		panic(fmt.Sprintf("expected CleanRoomAutoApprovalRules to be *cleanrooms.MockCleanRoomAutoApprovalRulesInterface, actual was %T", m.WorkspaceClient.CleanRoomAutoApprovalRules))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCleanRoomTaskRunsAPI() *cleanrooms.MockCleanRoomTaskRunsInterface {
	api, ok := m.WorkspaceClient.CleanRoomTaskRuns.(*cleanrooms.MockCleanRoomTaskRunsInterface)
	if !ok {
		panic(fmt.Sprintf("expected CleanRoomTaskRuns to be *cleanrooms.MockCleanRoomTaskRunsInterface, actual was %T", m.WorkspaceClient.CleanRoomTaskRuns))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCleanRoomsAPI() *cleanrooms.MockCleanRoomsInterface {
	api, ok := m.WorkspaceClient.CleanRooms.(*cleanrooms.MockCleanRoomsInterface)
	if !ok {
		panic(fmt.Sprintf("expected CleanRooms to be *cleanrooms.MockCleanRoomsInterface, actual was %T", m.WorkspaceClient.CleanRooms))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockClusterPoliciesAPI() *compute.MockClusterPoliciesInterface {
	api, ok := m.WorkspaceClient.ClusterPolicies.(*compute.MockClusterPoliciesInterface)
	if !ok {
		panic(fmt.Sprintf("expected ClusterPolicies to be *compute.MockClusterPoliciesInterface, actual was %T", m.WorkspaceClient.ClusterPolicies))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockClustersAPI() *compute.MockClustersInterface {
	api, ok := m.WorkspaceClient.Clusters.(*compute.MockClustersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Clusters to be *compute.MockClustersInterface, actual was %T", m.WorkspaceClient.Clusters))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCommandExecutionAPI() *compute.MockCommandExecutionInterface {
	api, ok := m.WorkspaceClient.CommandExecution.(*compute.MockCommandExecutionInterface)
	if !ok {
		panic(fmt.Sprintf("expected CommandExecution to be *compute.MockCommandExecutionInterface, actual was %T", m.WorkspaceClient.CommandExecution))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockComputeMigrationsAPI() *computemigrations.MockComputeMigrationsInterface {
	api, ok := m.WorkspaceClient.ComputeMigrations.(*computemigrations.MockComputeMigrationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ComputeMigrations to be *computemigrations.MockComputeMigrationsInterface, actual was %T", m.WorkspaceClient.ComputeMigrations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConnectionsAPI() *catalog.MockConnectionsInterface {
	api, ok := m.WorkspaceClient.Connections.(*catalog.MockConnectionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Connections to be *catalog.MockConnectionsInterface, actual was %T", m.WorkspaceClient.Connections))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerFulfillmentsAPI() *marketplace.MockConsumerFulfillmentsInterface {
	api, ok := m.WorkspaceClient.ConsumerFulfillments.(*marketplace.MockConsumerFulfillmentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerFulfillments to be *marketplace.MockConsumerFulfillmentsInterface, actual was %T", m.WorkspaceClient.ConsumerFulfillments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerInstallationsAPI() *marketplace.MockConsumerInstallationsInterface {
	api, ok := m.WorkspaceClient.ConsumerInstallations.(*marketplace.MockConsumerInstallationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerInstallations to be *marketplace.MockConsumerInstallationsInterface, actual was %T", m.WorkspaceClient.ConsumerInstallations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerListingsAPI() *marketplace.MockConsumerListingsInterface {
	api, ok := m.WorkspaceClient.ConsumerListings.(*marketplace.MockConsumerListingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerListings to be *marketplace.MockConsumerListingsInterface, actual was %T", m.WorkspaceClient.ConsumerListings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerPersonalizationRequestsAPI() *marketplace.MockConsumerPersonalizationRequestsInterface {
	api, ok := m.WorkspaceClient.ConsumerPersonalizationRequests.(*marketplace.MockConsumerPersonalizationRequestsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerPersonalizationRequests to be *marketplace.MockConsumerPersonalizationRequestsInterface, actual was %T", m.WorkspaceClient.ConsumerPersonalizationRequests))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerProvidersAPI() *marketplace.MockConsumerProvidersInterface {
	api, ok := m.WorkspaceClient.ConsumerProviders.(*marketplace.MockConsumerProvidersInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerProviders to be *marketplace.MockConsumerProvidersInterface, actual was %T", m.WorkspaceClient.ConsumerProviders))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCredentialsAPI() *catalog.MockCredentialsInterface {
	api, ok := m.WorkspaceClient.Credentials.(*catalog.MockCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Credentials to be *catalog.MockCredentialsInterface, actual was %T", m.WorkspaceClient.Credentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCredentialsManagerAPI() *settings.MockCredentialsManagerInterface {
	api, ok := m.WorkspaceClient.CredentialsManager.(*settings.MockCredentialsManagerInterface)
	if !ok {
		panic(fmt.Sprintf("expected CredentialsManager to be *settings.MockCredentialsManagerInterface, actual was %T", m.WorkspaceClient.CredentialsManager))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCurrentUserAPI() *iam.MockCurrentUserInterface {
	api, ok := m.WorkspaceClient.CurrentUser.(*iam.MockCurrentUserInterface)
	if !ok {
		panic(fmt.Sprintf("expected CurrentUser to be *iam.MockCurrentUserInterface, actual was %T", m.WorkspaceClient.CurrentUser))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDashboardWidgetsAPI() *sql.MockDashboardWidgetsInterface {
	api, ok := m.WorkspaceClient.DashboardWidgets.(*sql.MockDashboardWidgetsInterface)
	if !ok {
		panic(fmt.Sprintf("expected DashboardWidgets to be *sql.MockDashboardWidgetsInterface, actual was %T", m.WorkspaceClient.DashboardWidgets))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDashboardsAPI() *sql.MockDashboardsInterface {
	api, ok := m.WorkspaceClient.Dashboards.(*sql.MockDashboardsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Dashboards to be *sql.MockDashboardsInterface, actual was %T", m.WorkspaceClient.Dashboards))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDataClassificationAPI() *dataclassification.MockDataClassificationInterface {
	api, ok := m.WorkspaceClient.DataClassification.(*dataclassification.MockDataClassificationInterface)
	if !ok {
		panic(fmt.Sprintf("expected DataClassification to be *dataclassification.MockDataClassificationInterface, actual was %T", m.WorkspaceClient.DataClassification))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDataQualityAPI() *dataquality.MockDataQualityInterface {
	api, ok := m.WorkspaceClient.DataQuality.(*dataquality.MockDataQualityInterface)
	if !ok {
		panic(fmt.Sprintf("expected DataQuality to be *dataquality.MockDataQualityInterface, actual was %T", m.WorkspaceClient.DataQuality))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDataSourcesAPI() *sql.MockDataSourcesInterface {
	api, ok := m.WorkspaceClient.DataSources.(*sql.MockDataSourcesInterface)
	if !ok {
		panic(fmt.Sprintf("expected DataSources to be *sql.MockDataSourcesInterface, actual was %T", m.WorkspaceClient.DataSources))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDatabaseAPI() *database.MockDatabaseInterface {
	api, ok := m.WorkspaceClient.Database.(*database.MockDatabaseInterface)
	if !ok {
		panic(fmt.Sprintf("expected Database to be *database.MockDatabaseInterface, actual was %T", m.WorkspaceClient.Database))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDbfsAPI() *files.MockDbfsInterface {
	api, ok := m.WorkspaceClient.Dbfs.(*files.MockDbfsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Dbfs to be *files.MockDbfsInterface, actual was %T", m.WorkspaceClient.Dbfs))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDbsqlPermissionsAPI() *sql.MockDbsqlPermissionsInterface {
	api, ok := m.WorkspaceClient.DbsqlPermissions.(*sql.MockDbsqlPermissionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected DbsqlPermissions to be *sql.MockDbsqlPermissionsInterface, actual was %T", m.WorkspaceClient.DbsqlPermissions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDomainsAPI() *domains.MockDomainsInterface {
	api, ok := m.WorkspaceClient.Domains.(*domains.MockDomainsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Domains to be *domains.MockDomainsInterface, actual was %T", m.WorkspaceClient.Domains))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockEntityTagAssignmentsAPI() *catalog.MockEntityTagAssignmentsInterface {
	api, ok := m.WorkspaceClient.EntityTagAssignments.(*catalog.MockEntityTagAssignmentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected EntityTagAssignments to be *catalog.MockEntityTagAssignmentsInterface, actual was %T", m.WorkspaceClient.EntityTagAssignments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockEnvironmentsAPI() *environments.MockEnvironmentsInterface {
	api, ok := m.WorkspaceClient.Environments.(*environments.MockEnvironmentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Environments to be *environments.MockEnvironmentsInterface, actual was %T", m.WorkspaceClient.Environments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExperimentsAPI() *ml.MockExperimentsInterface {
	api, ok := m.WorkspaceClient.Experiments.(*ml.MockExperimentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Experiments to be *ml.MockExperimentsInterface, actual was %T", m.WorkspaceClient.Experiments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExternalLineageAPI() *catalog.MockExternalLineageInterface {
	api, ok := m.WorkspaceClient.ExternalLineage.(*catalog.MockExternalLineageInterface)
	if !ok {
		panic(fmt.Sprintf("expected ExternalLineage to be *catalog.MockExternalLineageInterface, actual was %T", m.WorkspaceClient.ExternalLineage))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExternalLocationsAPI() *catalog.MockExternalLocationsInterface {
	api, ok := m.WorkspaceClient.ExternalLocations.(*catalog.MockExternalLocationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ExternalLocations to be *catalog.MockExternalLocationsInterface, actual was %T", m.WorkspaceClient.ExternalLocations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExternalMetadataAPI() *catalog.MockExternalMetadataInterface {
	api, ok := m.WorkspaceClient.ExternalMetadata.(*catalog.MockExternalMetadataInterface)
	if !ok {
		panic(fmt.Sprintf("expected ExternalMetadata to be *catalog.MockExternalMetadataInterface, actual was %T", m.WorkspaceClient.ExternalMetadata))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFeatureEngineeringAPI() *ml.MockFeatureEngineeringInterface {
	api, ok := m.WorkspaceClient.FeatureEngineering.(*ml.MockFeatureEngineeringInterface)
	if !ok {
		panic(fmt.Sprintf("expected FeatureEngineering to be *ml.MockFeatureEngineeringInterface, actual was %T", m.WorkspaceClient.FeatureEngineering))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFeatureStoreAPI() *ml.MockFeatureStoreInterface {
	api, ok := m.WorkspaceClient.FeatureStore.(*ml.MockFeatureStoreInterface)
	if !ok {
		panic(fmt.Sprintf("expected FeatureStore to be *ml.MockFeatureStoreInterface, actual was %T", m.WorkspaceClient.FeatureStore))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFilesAPI() *files.MockFilesInterface {
	api, ok := m.WorkspaceClient.Files.(*files.MockFilesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Files to be *files.MockFilesInterface, actual was %T", m.WorkspaceClient.Files))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockForecastingAPI() *ml.MockForecastingInterface {
	api, ok := m.WorkspaceClient.Forecasting.(*ml.MockForecastingInterface)
	if !ok {
		panic(fmt.Sprintf("expected Forecasting to be *ml.MockForecastingInterface, actual was %T", m.WorkspaceClient.Forecasting))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFunctionsAPI() *catalog.MockFunctionsInterface {
	api, ok := m.WorkspaceClient.Functions.(*catalog.MockFunctionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Functions to be *catalog.MockFunctionsInterface, actual was %T", m.WorkspaceClient.Functions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGenieAPI() *dashboards.MockGenieInterface {
	api, ok := m.WorkspaceClient.Genie.(*dashboards.MockGenieInterface)
	if !ok {
		panic(fmt.Sprintf("expected Genie to be *dashboards.MockGenieInterface, actual was %T", m.WorkspaceClient.Genie))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGenieAgentsAPI() *genieagents.MockGenieAgentsInterface {
	api, ok := m.WorkspaceClient.GenieAgents.(*genieagents.MockGenieAgentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected GenieAgents to be *genieagents.MockGenieAgentsInterface, actual was %T", m.WorkspaceClient.GenieAgents))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGitCredentialsAPI() *workspace.MockGitCredentialsInterface {
	api, ok := m.WorkspaceClient.GitCredentials.(*workspace.MockGitCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected GitCredentials to be *workspace.MockGitCredentialsInterface, actual was %T", m.WorkspaceClient.GitCredentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGlobalInitScriptsAPI() *compute.MockGlobalInitScriptsInterface {
	api, ok := m.WorkspaceClient.GlobalInitScripts.(*compute.MockGlobalInitScriptsInterface)
	if !ok {
		panic(fmt.Sprintf("expected GlobalInitScripts to be *compute.MockGlobalInitScriptsInterface, actual was %T", m.WorkspaceClient.GlobalInitScripts))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGrantsAPI() *catalog.MockGrantsInterface {
	api, ok := m.WorkspaceClient.Grants.(*catalog.MockGrantsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Grants to be *catalog.MockGrantsInterface, actual was %T", m.WorkspaceClient.Grants))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGroupsV2API() *iam.MockGroupsV2Interface {
	api, ok := m.WorkspaceClient.GroupsV2.(*iam.MockGroupsV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected GroupsV2 to be *iam.MockGroupsV2Interface, actual was %T", m.WorkspaceClient.GroupsV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockInstancePoolsAPI() *compute.MockInstancePoolsInterface {
	api, ok := m.WorkspaceClient.InstancePools.(*compute.MockInstancePoolsInterface)
	if !ok {
		panic(fmt.Sprintf("expected InstancePools to be *compute.MockInstancePoolsInterface, actual was %T", m.WorkspaceClient.InstancePools))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockInstanceProfilesAPI() *compute.MockInstanceProfilesInterface {
	api, ok := m.WorkspaceClient.InstanceProfiles.(*compute.MockInstanceProfilesInterface)
	if !ok {
		panic(fmt.Sprintf("expected InstanceProfiles to be *compute.MockInstanceProfilesInterface, actual was %T", m.WorkspaceClient.InstanceProfiles))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockIpAccessListsAPI() *settings.MockIpAccessListsInterface {
	api, ok := m.WorkspaceClient.IpAccessLists.(*settings.MockIpAccessListsInterface)
	if !ok {
		panic(fmt.Sprintf("expected IpAccessLists to be *settings.MockIpAccessListsInterface, actual was %T", m.WorkspaceClient.IpAccessLists))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockJobsAPI() *jobs.MockJobsInterface {
	api, ok := m.WorkspaceClient.Jobs.(*jobs.MockJobsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Jobs to be *jobs.MockJobsInterface, actual was %T", m.WorkspaceClient.Jobs))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockKnowledgeAssistantsAPI() *knowledgeassistants.MockKnowledgeAssistantsInterface {
	api, ok := m.WorkspaceClient.KnowledgeAssistants.(*knowledgeassistants.MockKnowledgeAssistantsInterface)
	if !ok {
		panic(fmt.Sprintf("expected KnowledgeAssistants to be *knowledgeassistants.MockKnowledgeAssistantsInterface, actual was %T", m.WorkspaceClient.KnowledgeAssistants))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLakeviewAPI() *dashboards.MockLakeviewInterface {
	api, ok := m.WorkspaceClient.Lakeview.(*dashboards.MockLakeviewInterface)
	if !ok {
		panic(fmt.Sprintf("expected Lakeview to be *dashboards.MockLakeviewInterface, actual was %T", m.WorkspaceClient.Lakeview))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLakeviewEmbeddedAPI() *dashboards.MockLakeviewEmbeddedInterface {
	api, ok := m.WorkspaceClient.LakeviewEmbedded.(*dashboards.MockLakeviewEmbeddedInterface)
	if !ok {
		panic(fmt.Sprintf("expected LakeviewEmbedded to be *dashboards.MockLakeviewEmbeddedInterface, actual was %T", m.WorkspaceClient.LakeviewEmbedded))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLakewatchAPI() *lakewatch.MockLakewatchInterface {
	api, ok := m.WorkspaceClient.Lakewatch.(*lakewatch.MockLakewatchInterface)
	if !ok {
		panic(fmt.Sprintf("expected Lakewatch to be *lakewatch.MockLakewatchInterface, actual was %T", m.WorkspaceClient.Lakewatch))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLibrariesAPI() *compute.MockLibrariesInterface {
	api, ok := m.WorkspaceClient.Libraries.(*compute.MockLibrariesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Libraries to be *compute.MockLibrariesInterface, actual was %T", m.WorkspaceClient.Libraries))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockMaterializedFeaturesAPI() *ml.MockMaterializedFeaturesInterface {
	api, ok := m.WorkspaceClient.MaterializedFeatures.(*ml.MockMaterializedFeaturesInterface)
	if !ok {
		panic(fmt.Sprintf("expected MaterializedFeatures to be *ml.MockMaterializedFeaturesInterface, actual was %T", m.WorkspaceClient.MaterializedFeatures))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockMetastoresAPI() *catalog.MockMetastoresInterface {
	api, ok := m.WorkspaceClient.Metastores.(*catalog.MockMetastoresInterface)
	if !ok {
		panic(fmt.Sprintf("expected Metastores to be *catalog.MockMetastoresInterface, actual was %T", m.WorkspaceClient.Metastores))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockModelRegistryAPI() *ml.MockModelRegistryInterface {
	api, ok := m.WorkspaceClient.ModelRegistry.(*ml.MockModelRegistryInterface)
	if !ok {
		panic(fmt.Sprintf("expected ModelRegistry to be *ml.MockModelRegistryInterface, actual was %T", m.WorkspaceClient.ModelRegistry))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockModelVersionsAPI() *catalog.MockModelVersionsInterface {
	api, ok := m.WorkspaceClient.ModelVersions.(*catalog.MockModelVersionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ModelVersions to be *catalog.MockModelVersionsInterface, actual was %T", m.WorkspaceClient.ModelVersions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockNotificationDestinationsAPI() *settings.MockNotificationDestinationsInterface {
	api, ok := m.WorkspaceClient.NotificationDestinations.(*settings.MockNotificationDestinationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected NotificationDestinations to be *settings.MockNotificationDestinationsInterface, actual was %T", m.WorkspaceClient.NotificationDestinations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockOnlineTablesAPI() *catalog.MockOnlineTablesInterface {
	api, ok := m.WorkspaceClient.OnlineTables.(*catalog.MockOnlineTablesInterface)
	if !ok {
		panic(fmt.Sprintf("expected OnlineTables to be *catalog.MockOnlineTablesInterface, actual was %T", m.WorkspaceClient.OnlineTables))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPermissionMigrationAPI() *iam.MockPermissionMigrationInterface {
	api, ok := m.WorkspaceClient.PermissionMigration.(*iam.MockPermissionMigrationInterface)
	if !ok {
		panic(fmt.Sprintf("expected PermissionMigration to be *iam.MockPermissionMigrationInterface, actual was %T", m.WorkspaceClient.PermissionMigration))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPermissionsAPI() *iam.MockPermissionsInterface {
	api, ok := m.WorkspaceClient.Permissions.(*iam.MockPermissionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Permissions to be *iam.MockPermissionsInterface, actual was %T", m.WorkspaceClient.Permissions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPipelinesAPI() *pipelines.MockPipelinesInterface {
	api, ok := m.WorkspaceClient.Pipelines.(*pipelines.MockPipelinesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Pipelines to be *pipelines.MockPipelinesInterface, actual was %T", m.WorkspaceClient.Pipelines))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPoliciesAPI() *catalog.MockPoliciesInterface {
	api, ok := m.WorkspaceClient.Policies.(*catalog.MockPoliciesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Policies to be *catalog.MockPoliciesInterface, actual was %T", m.WorkspaceClient.Policies))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPolicyComplianceForClustersAPI() *compute.MockPolicyComplianceForClustersInterface {
	api, ok := m.WorkspaceClient.PolicyComplianceForClusters.(*compute.MockPolicyComplianceForClustersInterface)
	if !ok {
		panic(fmt.Sprintf("expected PolicyComplianceForClusters to be *compute.MockPolicyComplianceForClustersInterface, actual was %T", m.WorkspaceClient.PolicyComplianceForClusters))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPolicyComplianceForJobsAPI() *jobs.MockPolicyComplianceForJobsInterface {
	api, ok := m.WorkspaceClient.PolicyComplianceForJobs.(*jobs.MockPolicyComplianceForJobsInterface)
	if !ok {
		panic(fmt.Sprintf("expected PolicyComplianceForJobs to be *jobs.MockPolicyComplianceForJobsInterface, actual was %T", m.WorkspaceClient.PolicyComplianceForJobs))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPolicyFamiliesAPI() *compute.MockPolicyFamiliesInterface {
	api, ok := m.WorkspaceClient.PolicyFamilies.(*compute.MockPolicyFamiliesInterface)
	if !ok {
		panic(fmt.Sprintf("expected PolicyFamilies to be *compute.MockPolicyFamiliesInterface, actual was %T", m.WorkspaceClient.PolicyFamilies))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPostgresAPI() *postgres.MockPostgresInterface {
	api, ok := m.WorkspaceClient.Postgres.(*postgres.MockPostgresInterface)
	if !ok {
		panic(fmt.Sprintf("expected Postgres to be *postgres.MockPostgresInterface, actual was %T", m.WorkspaceClient.Postgres))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderExchangeFiltersAPI() *marketplace.MockProviderExchangeFiltersInterface {
	api, ok := m.WorkspaceClient.ProviderExchangeFilters.(*marketplace.MockProviderExchangeFiltersInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderExchangeFilters to be *marketplace.MockProviderExchangeFiltersInterface, actual was %T", m.WorkspaceClient.ProviderExchangeFilters))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderExchangesAPI() *marketplace.MockProviderExchangesInterface {
	api, ok := m.WorkspaceClient.ProviderExchanges.(*marketplace.MockProviderExchangesInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderExchanges to be *marketplace.MockProviderExchangesInterface, actual was %T", m.WorkspaceClient.ProviderExchanges))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderFilesAPI() *marketplace.MockProviderFilesInterface {
	api, ok := m.WorkspaceClient.ProviderFiles.(*marketplace.MockProviderFilesInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderFiles to be *marketplace.MockProviderFilesInterface, actual was %T", m.WorkspaceClient.ProviderFiles))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderListingsAPI() *marketplace.MockProviderListingsInterface {
	api, ok := m.WorkspaceClient.ProviderListings.(*marketplace.MockProviderListingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderListings to be *marketplace.MockProviderListingsInterface, actual was %T", m.WorkspaceClient.ProviderListings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderPersonalizationRequestsAPI() *marketplace.MockProviderPersonalizationRequestsInterface {
	api, ok := m.WorkspaceClient.ProviderPersonalizationRequests.(*marketplace.MockProviderPersonalizationRequestsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderPersonalizationRequests to be *marketplace.MockProviderPersonalizationRequestsInterface, actual was %T", m.WorkspaceClient.ProviderPersonalizationRequests))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderProviderAnalyticsDashboardsAPI() *marketplace.MockProviderProviderAnalyticsDashboardsInterface {
	api, ok := m.WorkspaceClient.ProviderProviderAnalyticsDashboards.(*marketplace.MockProviderProviderAnalyticsDashboardsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderProviderAnalyticsDashboards to be *marketplace.MockProviderProviderAnalyticsDashboardsInterface, actual was %T", m.WorkspaceClient.ProviderProviderAnalyticsDashboards))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderProvidersAPI() *marketplace.MockProviderProvidersInterface {
	api, ok := m.WorkspaceClient.ProviderProviders.(*marketplace.MockProviderProvidersInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderProviders to be *marketplace.MockProviderProvidersInterface, actual was %T", m.WorkspaceClient.ProviderProviders))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProvidersAPI() *sharing.MockProvidersInterface {
	api, ok := m.WorkspaceClient.Providers.(*sharing.MockProvidersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Providers to be *sharing.MockProvidersInterface, actual was %T", m.WorkspaceClient.Providers))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQualityMonitorV2API() *qualitymonitorv2.MockQualityMonitorV2Interface {
	api, ok := m.WorkspaceClient.QualityMonitorV2.(*qualitymonitorv2.MockQualityMonitorV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected QualityMonitorV2 to be *qualitymonitorv2.MockQualityMonitorV2Interface, actual was %T", m.WorkspaceClient.QualityMonitorV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQualityMonitorsAPI() *catalog.MockQualityMonitorsInterface {
	api, ok := m.WorkspaceClient.QualityMonitors.(*catalog.MockQualityMonitorsInterface)
	if !ok {
		panic(fmt.Sprintf("expected QualityMonitors to be *catalog.MockQualityMonitorsInterface, actual was %T", m.WorkspaceClient.QualityMonitors))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueriesAPI() *sql.MockQueriesInterface {
	api, ok := m.WorkspaceClient.Queries.(*sql.MockQueriesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Queries to be *sql.MockQueriesInterface, actual was %T", m.WorkspaceClient.Queries))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueriesLegacyAPI() *sql.MockQueriesLegacyInterface {
	api, ok := m.WorkspaceClient.QueriesLegacy.(*sql.MockQueriesLegacyInterface)
	if !ok {
		panic(fmt.Sprintf("expected QueriesLegacy to be *sql.MockQueriesLegacyInterface, actual was %T", m.WorkspaceClient.QueriesLegacy))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueryHistoryAPI() *sql.MockQueryHistoryInterface {
	api, ok := m.WorkspaceClient.QueryHistory.(*sql.MockQueryHistoryInterface)
	if !ok {
		panic(fmt.Sprintf("expected QueryHistory to be *sql.MockQueryHistoryInterface, actual was %T", m.WorkspaceClient.QueryHistory))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueryVisualizationsAPI() *sql.MockQueryVisualizationsInterface {
	api, ok := m.WorkspaceClient.QueryVisualizations.(*sql.MockQueryVisualizationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected QueryVisualizations to be *sql.MockQueryVisualizationsInterface, actual was %T", m.WorkspaceClient.QueryVisualizations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueryVisualizationsLegacyAPI() *sql.MockQueryVisualizationsLegacyInterface {
	api, ok := m.WorkspaceClient.QueryVisualizationsLegacy.(*sql.MockQueryVisualizationsLegacyInterface)
	if !ok {
		panic(fmt.Sprintf("expected QueryVisualizationsLegacy to be *sql.MockQueryVisualizationsLegacyInterface, actual was %T", m.WorkspaceClient.QueryVisualizationsLegacy))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRecipientActivationAPI() *sharing.MockRecipientActivationInterface {
	api, ok := m.WorkspaceClient.RecipientActivation.(*sharing.MockRecipientActivationInterface)
	if !ok {
		panic(fmt.Sprintf("expected RecipientActivation to be *sharing.MockRecipientActivationInterface, actual was %T", m.WorkspaceClient.RecipientActivation))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRecipientFederationPoliciesAPI() *sharing.MockRecipientFederationPoliciesInterface {
	api, ok := m.WorkspaceClient.RecipientFederationPolicies.(*sharing.MockRecipientFederationPoliciesInterface)
	if !ok {
		panic(fmt.Sprintf("expected RecipientFederationPolicies to be *sharing.MockRecipientFederationPoliciesInterface, actual was %T", m.WorkspaceClient.RecipientFederationPolicies))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRecipientsAPI() *sharing.MockRecipientsInterface {
	api, ok := m.WorkspaceClient.Recipients.(*sharing.MockRecipientsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Recipients to be *sharing.MockRecipientsInterface, actual was %T", m.WorkspaceClient.Recipients))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRedashConfigAPI() *sql.MockRedashConfigInterface {
	api, ok := m.WorkspaceClient.RedashConfig.(*sql.MockRedashConfigInterface)
	if !ok {
		panic(fmt.Sprintf("expected RedashConfig to be *sql.MockRedashConfigInterface, actual was %T", m.WorkspaceClient.RedashConfig))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRegisteredModelsAPI() *catalog.MockRegisteredModelsInterface {
	api, ok := m.WorkspaceClient.RegisteredModels.(*catalog.MockRegisteredModelsInterface)
	if !ok {
		panic(fmt.Sprintf("expected RegisteredModels to be *catalog.MockRegisteredModelsInterface, actual was %T", m.WorkspaceClient.RegisteredModels))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockReposAPI() *workspace.MockReposInterface {
	api, ok := m.WorkspaceClient.Repos.(*workspace.MockReposInterface)
	if !ok {
		panic(fmt.Sprintf("expected Repos to be *workspace.MockReposInterface, actual was %T", m.WorkspaceClient.Repos))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockResourceQuotasAPI() *catalog.MockResourceQuotasInterface {
	api, ok := m.WorkspaceClient.ResourceQuotas.(*catalog.MockResourceQuotasInterface)
	if !ok {
		panic(fmt.Sprintf("expected ResourceQuotas to be *catalog.MockResourceQuotasInterface, actual was %T", m.WorkspaceClient.ResourceQuotas))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRfaAPI() *catalog.MockRfaInterface {
	api, ok := m.WorkspaceClient.Rfa.(*catalog.MockRfaInterface)
	if !ok {
		panic(fmt.Sprintf("expected Rfa to be *catalog.MockRfaInterface, actual was %T", m.WorkspaceClient.Rfa))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSandboxAPI() *sandbox.MockSandboxInterface {
	api, ok := m.WorkspaceClient.Sandbox.(*sandbox.MockSandboxInterface)
	if !ok {
		panic(fmt.Sprintf("expected Sandbox to be *sandbox.MockSandboxInterface, actual was %T", m.WorkspaceClient.Sandbox))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSchemasAPI() *catalog.MockSchemasInterface {
	api, ok := m.WorkspaceClient.Schemas.(*catalog.MockSchemasInterface)
	if !ok {
		panic(fmt.Sprintf("expected Schemas to be *catalog.MockSchemasInterface, actual was %T", m.WorkspaceClient.Schemas))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSecretsAPI() *workspace.MockSecretsInterface {
	api, ok := m.WorkspaceClient.Secrets.(*workspace.MockSecretsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Secrets to be *workspace.MockSecretsInterface, actual was %T", m.WorkspaceClient.Secrets))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSecretsUcAPI() *catalog.MockSecretsUcInterface {
	api, ok := m.WorkspaceClient.SecretsUc.(*catalog.MockSecretsUcInterface)
	if !ok {
		panic(fmt.Sprintf("expected SecretsUc to be *catalog.MockSecretsUcInterface, actual was %T", m.WorkspaceClient.SecretsUc))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServerlessComputeServiceAPI() *compute.MockServerlessComputeServiceInterface {
	api, ok := m.WorkspaceClient.ServerlessComputeService.(*compute.MockServerlessComputeServiceInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServerlessComputeService to be *compute.MockServerlessComputeServiceInterface, actual was %T", m.WorkspaceClient.ServerlessComputeService))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServicePrincipalSecretsProxyAPI() *oauth2.MockServicePrincipalSecretsProxyInterface {
	api, ok := m.WorkspaceClient.ServicePrincipalSecretsProxy.(*oauth2.MockServicePrincipalSecretsProxyInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipalSecretsProxy to be *oauth2.MockServicePrincipalSecretsProxyInterface, actual was %T", m.WorkspaceClient.ServicePrincipalSecretsProxy))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServicePrincipalsV2API() *iam.MockServicePrincipalsV2Interface {
	api, ok := m.WorkspaceClient.ServicePrincipalsV2.(*iam.MockServicePrincipalsV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipalsV2 to be *iam.MockServicePrincipalsV2Interface, actual was %T", m.WorkspaceClient.ServicePrincipalsV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServingEndpointsAPI() *serving.MockServingEndpointsInterface {
	api, ok := m.WorkspaceClient.ServingEndpoints.(*serving.MockServingEndpointsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServingEndpoints to be *serving.MockServingEndpointsInterface, actual was %T", m.WorkspaceClient.ServingEndpoints))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServingEndpointsDataPlaneAPI() *serving.MockServingEndpointsDataPlaneInterface {
	api, ok := m.WorkspaceClient.ServingEndpointsDataPlane.(*serving.MockServingEndpointsDataPlaneInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServingEndpointsDataPlane to be *serving.MockServingEndpointsDataPlaneInterface, actual was %T", m.WorkspaceClient.ServingEndpointsDataPlane))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSettingsAPI() *settings.MockSettingsInterface {
	api, ok := m.WorkspaceClient.Settings.(*settings.MockSettingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Settings to be *settings.MockSettingsInterface, actual was %T", m.WorkspaceClient.Settings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSharesAPI() *sharing.MockSharesInterface {
	api, ok := m.WorkspaceClient.Shares.(*sharing.MockSharesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Shares to be *sharing.MockSharesInterface, actual was %T", m.WorkspaceClient.Shares))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSoftwareArtifactsAPI() *softwareartifacts.MockSoftwareArtifactsInterface {
	api, ok := m.WorkspaceClient.SoftwareArtifacts.(*softwareartifacts.MockSoftwareArtifactsInterface)
	if !ok {
		panic(fmt.Sprintf("expected SoftwareArtifacts to be *softwareartifacts.MockSoftwareArtifactsInterface, actual was %T", m.WorkspaceClient.SoftwareArtifacts))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockStatementExecutionAPI() *sql.MockStatementExecutionInterface {
	api, ok := m.WorkspaceClient.StatementExecution.(*sql.MockStatementExecutionInterface)
	if !ok {
		panic(fmt.Sprintf("expected StatementExecution to be *sql.MockStatementExecutionInterface, actual was %T", m.WorkspaceClient.StatementExecution))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockStatementExecutionV2API() *statementexecutionv2.MockStatementExecutionV2Interface {
	api, ok := m.WorkspaceClient.StatementExecutionV2.(*statementexecutionv2.MockStatementExecutionV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected StatementExecutionV2 to be *statementexecutionv2.MockStatementExecutionV2Interface, actual was %T", m.WorkspaceClient.StatementExecutionV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockStorageCredentialsAPI() *catalog.MockStorageCredentialsInterface {
	api, ok := m.WorkspaceClient.StorageCredentials.(*catalog.MockStorageCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected StorageCredentials to be *catalog.MockStorageCredentialsInterface, actual was %T", m.WorkspaceClient.StorageCredentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSupervisorAgentsAPI() *supervisoragents.MockSupervisorAgentsInterface {
	api, ok := m.WorkspaceClient.SupervisorAgents.(*supervisoragents.MockSupervisorAgentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected SupervisorAgents to be *supervisoragents.MockSupervisorAgentsInterface, actual was %T", m.WorkspaceClient.SupervisorAgents))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSystemSchemasAPI() *catalog.MockSystemSchemasInterface {
	api, ok := m.WorkspaceClient.SystemSchemas.(*catalog.MockSystemSchemasInterface)
	if !ok {
		panic(fmt.Sprintf("expected SystemSchemas to be *catalog.MockSystemSchemasInterface, actual was %T", m.WorkspaceClient.SystemSchemas))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTableConstraintsAPI() *catalog.MockTableConstraintsInterface {
	api, ok := m.WorkspaceClient.TableConstraints.(*catalog.MockTableConstraintsInterface)
	if !ok {
		panic(fmt.Sprintf("expected TableConstraints to be *catalog.MockTableConstraintsInterface, actual was %T", m.WorkspaceClient.TableConstraints))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTablesAPI() *catalog.MockTablesInterface {
	api, ok := m.WorkspaceClient.Tables.(*catalog.MockTablesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Tables to be *catalog.MockTablesInterface, actual was %T", m.WorkspaceClient.Tables))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTagPoliciesAPI() *tags.MockTagPoliciesInterface {
	api, ok := m.WorkspaceClient.TagPolicies.(*tags.MockTagPoliciesInterface)
	if !ok {
		panic(fmt.Sprintf("expected TagPolicies to be *tags.MockTagPoliciesInterface, actual was %T", m.WorkspaceClient.TagPolicies))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTemporaryPathCredentialsAPI() *catalog.MockTemporaryPathCredentialsInterface {
	api, ok := m.WorkspaceClient.TemporaryPathCredentials.(*catalog.MockTemporaryPathCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected TemporaryPathCredentials to be *catalog.MockTemporaryPathCredentialsInterface, actual was %T", m.WorkspaceClient.TemporaryPathCredentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTemporaryTableCredentialsAPI() *catalog.MockTemporaryTableCredentialsInterface {
	api, ok := m.WorkspaceClient.TemporaryTableCredentials.(*catalog.MockTemporaryTableCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected TemporaryTableCredentials to be *catalog.MockTemporaryTableCredentialsInterface, actual was %T", m.WorkspaceClient.TemporaryTableCredentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTemporaryVolumeCredentialsAPI() *catalog.MockTemporaryVolumeCredentialsInterface {
	api, ok := m.WorkspaceClient.TemporaryVolumeCredentials.(*catalog.MockTemporaryVolumeCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected TemporaryVolumeCredentials to be *catalog.MockTemporaryVolumeCredentialsInterface, actual was %T", m.WorkspaceClient.TemporaryVolumeCredentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTokenManagementAPI() *settings.MockTokenManagementInterface {
	api, ok := m.WorkspaceClient.TokenManagement.(*settings.MockTokenManagementInterface)
	if !ok {
		panic(fmt.Sprintf("expected TokenManagement to be *settings.MockTokenManagementInterface, actual was %T", m.WorkspaceClient.TokenManagement))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTokensAPI() *settings.MockTokensInterface {
	api, ok := m.WorkspaceClient.Tokens.(*settings.MockTokensInterface)
	if !ok {
		panic(fmt.Sprintf("expected Tokens to be *settings.MockTokensInterface, actual was %T", m.WorkspaceClient.Tokens))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockUsersV2API() *iam.MockUsersV2Interface {
	api, ok := m.WorkspaceClient.UsersV2.(*iam.MockUsersV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected UsersV2 to be *iam.MockUsersV2Interface, actual was %T", m.WorkspaceClient.UsersV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockVectorSearchEndpointsAPI() *vectorsearch.MockVectorSearchEndpointsInterface {
	api, ok := m.WorkspaceClient.VectorSearchEndpoints.(*vectorsearch.MockVectorSearchEndpointsInterface)
	if !ok {
		panic(fmt.Sprintf("expected VectorSearchEndpoints to be *vectorsearch.MockVectorSearchEndpointsInterface, actual was %T", m.WorkspaceClient.VectorSearchEndpoints))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockVectorSearchIndexesAPI() *vectorsearch.MockVectorSearchIndexesInterface {
	api, ok := m.WorkspaceClient.VectorSearchIndexes.(*vectorsearch.MockVectorSearchIndexesInterface)
	if !ok {
		panic(fmt.Sprintf("expected VectorSearchIndexes to be *vectorsearch.MockVectorSearchIndexesInterface, actual was %T", m.WorkspaceClient.VectorSearchIndexes))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockVolumesAPI() *catalog.MockVolumesInterface {
	api, ok := m.WorkspaceClient.Volumes.(*catalog.MockVolumesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Volumes to be *catalog.MockVolumesInterface, actual was %T", m.WorkspaceClient.Volumes))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWarehousesAPI() *sql.MockWarehousesInterface {
	api, ok := m.WorkspaceClient.Warehouses.(*sql.MockWarehousesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Warehouses to be *sql.MockWarehousesInterface, actual was %T", m.WorkspaceClient.Warehouses))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceAPI() *workspace.MockWorkspaceInterface {
	api, ok := m.WorkspaceClient.Workspace.(*workspace.MockWorkspaceInterface)
	if !ok {
		panic(fmt.Sprintf("expected Workspace to be *workspace.MockWorkspaceInterface, actual was %T", m.WorkspaceClient.Workspace))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceBindingsAPI() *catalog.MockWorkspaceBindingsInterface {
	api, ok := m.WorkspaceClient.WorkspaceBindings.(*catalog.MockWorkspaceBindingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceBindings to be *catalog.MockWorkspaceBindingsInterface, actual was %T", m.WorkspaceClient.WorkspaceBindings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceConfAPI() *settings.MockWorkspaceConfInterface {
	api, ok := m.WorkspaceClient.WorkspaceConf.(*settings.MockWorkspaceConfInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceConf to be *settings.MockWorkspaceConfInterface, actual was %T", m.WorkspaceClient.WorkspaceConf))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceEntityTagAssignmentsAPI() *tags.MockWorkspaceEntityTagAssignmentsInterface {
	api, ok := m.WorkspaceClient.WorkspaceEntityTagAssignments.(*tags.MockWorkspaceEntityTagAssignmentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceEntityTagAssignments to be *tags.MockWorkspaceEntityTagAssignmentsInterface, actual was %T", m.WorkspaceClient.WorkspaceEntityTagAssignments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceIamV2API() *iamv2.MockWorkspaceIamV2Interface {
	api, ok := m.WorkspaceClient.WorkspaceIamV2.(*iamv2.MockWorkspaceIamV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceIamV2 to be *iamv2.MockWorkspaceIamV2Interface, actual was %T", m.WorkspaceClient.WorkspaceIamV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceSettingsV2API() *settingsv2.MockWorkspaceSettingsV2Interface {
	api, ok := m.WorkspaceClient.WorkspaceSettingsV2.(*settingsv2.MockWorkspaceSettingsV2Interface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceSettingsV2 to be *settingsv2.MockWorkspaceSettingsV2Interface, actual was %T", m.WorkspaceClient.WorkspaceSettingsV2))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGroupsAPI() *iam.MockGroupsInterface {
	api, ok := m.WorkspaceClient.Groups.(*iam.MockGroupsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Groups to be *iam.MockGroupsInterface, actual was %T", m.WorkspaceClient.Groups))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServicePrincipalsAPI() *iam.MockServicePrincipalsInterface {
	api, ok := m.WorkspaceClient.ServicePrincipals.(*iam.MockServicePrincipalsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipals to be *iam.MockServicePrincipalsInterface, actual was %T", m.WorkspaceClient.ServicePrincipals))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockUsersAPI() *iam.MockUsersInterface {
	api, ok := m.WorkspaceClient.Users.(*iam.MockUsersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Users to be *iam.MockUsersInterface, actual was %T", m.WorkspaceClient.Users))
	}
	return api
}
