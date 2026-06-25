# Fakebricks acceptance status

_Updated 2026-06-25._ Open failures: **156** · Fixed: **0** · Passing: **131** · Skipped (gated): **261**.

Retest a set after a fakebricks fix, e.g. `python3 fakebricks-acc/retest.py run --match "/scim/v2/Users"`.

## Open failures by cause

### 501_NOT_IMPLEMENTED — `POST /api/2.0/preview/scim/v2/Users` (17)
- [ ] aws/TestAccUserRole
- [ ] aws/TestAccUserRoleWithApiField
- [ ] permissions/TestAccPermissions_Directory_Id
- [ ] permissions/TestAccPermissions_Notebook_Id
- [ ] permissions/TestAccPermissions_Pipeline
- [ ] permissions/TestMwsAccAccountGroupRuleSetsFullLifeCycle
- [ ] scim/TestAccDataUserWithApiField
- [ ] scim/TestAccEntitlementsAddToEmpty
- [ ] scim/TestAccEntitlementsSetExplicitlyToFalse
- [ ] scim/TestAccForceUserImport
- [ ] scim/TestAccUserData
- [ ] scim/TestAccUserHomeDelete
- [ ] scim/TestAccUserHomeDeleteNotDeleted
- [ ] scim/TestAccUserResource
- [ ] scim/TestAccUserResourceCaseInsensitive
- [ ] scim/TestAccUserWithApiField
- [ ] scim/TestMwsAccUserData

### NOT_SEEDED (15)
- [ ] catalog/TestUcAccRegisteredModel — cannot create registered model: Catalog 'main' does not exist. Error: cannot create registered model: Catalog 'main' does not exist. Error running apply: exit status 1
- [ ] clusters/TestAccClusterPolicyResourceOverrideBuiltIn — cannot create cluster policy: Policy named 'Personal Compute' does not exist Error: cannot create cluster policy: Policy named 'Personal Compute' does not exist Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_Authorization_Tokens — cannot create permissions: Node ID tokens does not exist. Error: cannot create permissions: Node ID tokens does not exist. Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_ClusterPolicy — cannot create permissions: Node ID 0000000000000001 does not exist. Error: cannot create permissions: Node ID 0000000000000001 does not exist. Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_Directory_RootDirectoryCorrectlyHandlesAdminUsers — cannot create permissions: Node ID 0 does not exist. Error: cannot create permissions: Node ID 0 does not exist. Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_RegisteredModel_Root — cannot create permissions: Node ID root does not exist. Error: cannot create permissions: Node ID root does not exist. Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_Repo_Id — cannot create permissions: Node ID 100000006 does not exist. Error: cannot create permissions: Node ID 100000006 does not exist. Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_Repo_Path — cannot create permissions: Node ID 100000006 does not exist. Error: cannot create permissions: Node ID 100000006 does not exist. Error running apply: exit status 1
- [ ] storage/TestUcAccFileBase64FullLifeCycle — cannot create schema: Catalog 'main' does not exist. Error: cannot create schema: Catalog 'main' does not exist. Error running apply: exit status 1
- [ ] storage/TestUcAccFileDontUpdateIfNoChange — cannot create schema: Catalog 'main' does not exist. Error: cannot create schema: Catalog 'main' does not exist. Error running apply: exit status 1
- [ ] storage/TestUcAccFileFullLifeCycle — cannot create schema: Catalog 'main' does not exist. Error: cannot create schema: Catalog 'main' does not exist. Error running apply: exit status 1
- [ ] storage/TestUcAccFileNoUpdateIfFileDoesNotChange — cannot create schema: Catalog 'main' does not exist. Error: cannot create schema: Catalog 'main' does not exist. Error running apply: exit status 1
- [ ] storage/TestUcAccFileUpdateOnLocalContentChange — cannot create schema: Catalog 'main' does not exist. Error: cannot create schema: Catalog 'main' does not exist. Error running apply: exit status 1
- [ ] storage/TestUcAccFileUpdateOnLocalFileChange — cannot create schema: Catalog 'main' does not exist. Error: cannot create schema: Catalog 'main' does not exist. Error running apply: exit status 1
- [ ] storage/TestUcAccFileUpdateServerChange — cannot create schema: Catalog 'main' does not exist. Error: cannot create schema: Catalog 'main' does not exist. 2026-06-25T05:31:58.723Z [ERROR] sdk.helper_resource: Expected an error with pattern ( the refresh plan was not empty.): test_name=TestUcAccFileUpdateServerChange test_terraform_path=/usr

### WORKSPACE_ID (14)
- [ ] catalog/TestAccCatalog_ProviderConfig_Match — failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1500000000000000 but got 6051921418418893 in provider_config. please check the workspace_id provided in provider_config Error: failed to validate workspace_id: workspace_id mismatch: provider is configured f
- [ ] catalog/TestAccCatalog_ProviderConfig_Recreate — failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1500000000000000 but got 6051921418418893 in provider_config. please check the workspace_id provided in provider_config Error: failed to validate workspace_id: workspace_id mismatch: provider is configured f
- [ ] catalog/TestAccCatalog_ProviderConfig_Remove — failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1500000000000000 but got 6051921418418893 in provider_config. please check the workspace_id provided in provider_config Error: failed to validate workspace_id: workspace_id mismatch: provider is configured f
- [ ] catalog/TestUcAccDataSourceExternalLocations — Check failed: names are empty: map[%:3 id:_ provider_config.#:1 provider_config.0.%:1 provider_config.0.workspace_id:1500000000000000]
- [ ] catalog/TestUcAccDataSourceStorageCredentials — Check failed: names are empty: map[%:3 id:_ provider_config.#:1 provider_config.0.%:1 provider_config.0.workspace_id:1500000000000000]
- [ ] internal/acceptance/TestMwsAccWorkspaceIDHttp_NoDefaultNoOverride — 2026-06-25T05:34:18.140Z [ERROR] sdk.helper_resource: Expected an error with pattern (managing workspace-level resources requires a workspace_id, but none was found in the resource's provider_config block or the provider's workspace_id attribute): test_step_number=1 test_terraform_path=/usr/local/bi
- [ ] internal/acceptance/TestMwsAccWorkspaceID_NoDefaultNoOverride — 2026-06-25T05:34:22.944Z [ERROR] sdk.helper_resource: Expected an error with pattern (managing workspace-level resources requires a workspace_id, but none was found in the resource's provider_config block or the provider's workspace_id attribute): test_terraform_path=/usr/local/bin/terraform test_wo
- [ ] internal/providers/pluginfw/products/sharing/TestAccShareData_ProviderConfig_Mismatched — failed to create share Error: failed to create share 2026-06-25T05:23:03.667Z [ERROR] sdk.helper_resource: Expected an error with pattern ((?s)failed to get workspace client.*workspace_id mismatch.*please check the workspace_id provided in provider_config): test_step_number=1 test_name=TestAccShareD
- [ ] internal/providers/pluginfw/products/sharing/TestAccShare_ProviderConfig_Invalid — failed to get workspace client Error: failed to get workspace client 2026-06-25T05:22:52.560Z [ERROR] sdk.helper_resource: Expected an error with pattern ((?s)Attribute provider_config\[0\]\.workspace_id workspace_id must be a valid.*integer, got: invalid): test_name=TestAccShare_ProviderConfig_Inva
- [ ] internal/providers/pluginfw/products/sharing/TestAccSharesData_ProviderConfig_Mismatched — failed to create share Error: failed to create share 2026-06-25T05:23:32.766Z [ERROR] sdk.helper_resource: Expected an error with pattern ((?s)failed to get workspace client.*workspace_id mismatch.*please check the workspace_id provided in provider_config): test_name=TestAccSharesData_ProviderConfig
- [ ] jobs/TestAccDataSourcesJob — cannot read jobs: cannot get client jobs: failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1500000000000000 but got 6051921418418893 in provider_config. please check the workspace_id provided in provider_config Error: cannot read jobs: cannot get client jo
- [ ] workspace/TestAccNotebook_ProviderConfig_Match — failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1500000000000000 but got 6051921418418893 in provider_config. please check the workspace_id provided in provider_config Error: failed to validate workspace_id: workspace_id mismatch: provider is configured f
- [ ] workspace/TestAccNotebook_ProviderConfig_Recreate — failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1500000000000000 but got 6051921418418893 in provider_config. please check the workspace_id provided in provider_config Error: failed to validate workspace_id: workspace_id mismatch: provider is configured f
- [ ] workspace/TestAccNotebook_ProviderConfig_Remove — failed to validate workspace_id: workspace_id mismatch: provider is configured for workspace 1500000000000000 but got 6051921418418893 in provider_config. please check the workspace_id provided in provider_config Error: failed to validate workspace_id: workspace_id mismatch: provider is configured f

### OTHER (13)
- [ ] internal/acceptance/TestAccWorkspaceIDApp_ProviderUpgrade — Check failed: env var THIS_WORKSPACE_ID is not set
- [ ] internal/acceptance/TestAccWorkspaceIDHttp_DefaultOnWorkspaceProvider_Same — Check failed: env var THIS_WORKSPACE_ID is not set
- [ ] internal/acceptance/TestAccWorkspaceIDHttp_ProviderUpgrade — Check failed: env var THIS_WORKSPACE_ID is not set
- [ ] internal/acceptance/TestAccWorkspaceIDTagPolicy_ProviderUpgrade — failed to create tag_policy Error: failed to create tag_policy Error running apply: exit status 1
- [ ] internal/acceptance/TestAccWorkspaceID_DefaultOnWorkspaceProvider_Same — Check failed: env var THIS_WORKSPACE_ID is not set
- [ ] internal/acceptance/TestAccWorkspaceID_ProviderUpgrade — Check failed: env var THIS_WORKSPACE_ID is not set
- [ ] internal/providers/pluginfw/products/app/TestAccApp_ProviderConfig_Apply — failed to get workspace client Error: failed to get workspace client Error running pre-apply plan: exit status 1
- [ ] permissions/TestAccPermissions_ServingEndpoint — Argument is deprecated cannot create model serving: fakebricks serving_endpoints models only external-model endpoints Error: cannot create model serving: fakebricks serving_endpoints models only external-model endpoints Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_SqlWarehouses — cannot read permissions: expected object type warehouses, got sql/warehouses Error: cannot read permissions: expected object type warehouses, got sql/warehouses Error running apply: exit status 1
- [ ] serving/TestAccModelServing — Argument is deprecated failed to list endpoints Error: failed to list endpoints Error running pre-apply plan: exit status 1
- [ ] sql/TestAccSQLGlobalConfig
- [ ] sql/TestAccSQLGlobalConfigServerless
- [ ] sql/TestAccSQLGlobalConfig_GoogleServiceAccount_OnAWS

### 501_NOT_IMPLEMENTED — `POST /api/2.1/unity-catalog/shares` (12)
- [ ] internal/providers/pluginfw/products/sharing/TestAccShareData_ProviderConfig_Apply
- [ ] internal/providers/pluginfw/products/sharing/TestAccShare_ProviderConfig_Match
- [ ] internal/providers/pluginfw/products/sharing/TestAccShare_ProviderConfig_NotProvided
- [ ] internal/providers/pluginfw/products/sharing/TestAccShare_ProviderConfig_Recreate
- [ ] internal/providers/pluginfw/products/sharing/TestAccShare_ProviderConfig_Remove
- [ ] internal/providers/pluginfw/products/sharing/TestUcAccShareMigrationFromPluginFramework
- [ ] internal/providers/pluginfw/products/sharing/TestUcAccShareMigrationFromSDKv2
- [ ] internal/providers/pluginfw/products/sharing/TestUcAccShareReorderObject
- [ ] internal/providers/pluginfw/products/sharing/TestUcAccUpdateShareComplexObjectChanges
- [ ] internal/providers/pluginfw/products/sharing/TestUcAccUpdateShareNoChanges
- [ ] internal/providers/pluginfw/products/sharing/TestUcAccUpdateShareOutsideTerraform
- [ ] internal/providers/pluginfw/products/sharing/TestUcAccUpdateShareRemoveAllObjects

### 501_NOT_IMPLEMENTED — `GET /api/2.1/unity-catalog/current-metastore-assignment` (9)
- [ ] catalog/TestUcAccCatalog
- [ ] catalog/TestUcAccCatalogForceDestroyConsistentAfterImport
- [ ] catalog/TestUcAccCatalogIsolated
- [ ] catalog/TestUcAccDataSourceCatalog
- [ ] catalog/TestUcAccDataSourceSchema
- [ ] catalog/TestUcAccDataSourceSchemas
- [ ] catalog/TestUcAccDataSourceVolume
- [ ] catalog/TestUcAccDataSourceVolumes
- [ ] internal/providers/pluginfw/products/volume/TestUcAccDataSourceVolumes

### 501_NOT_IMPLEMENTED — `GET /api/2.1/clusters/spark-versions` (9)
- [ ] clusters/TestAccClusterResource_CreateAndUpdateAwsAttributes
- [ ] clusters/TestAccSparkVersion
- [ ] jobs/TestAccDataSourceJob
- [ ] jobs/TestAccDataSourceJobApply
- [ ] jobs/TestAccJobCluster_ProviderConfig_Match
- [ ] jobs/TestAccJobCluster_ProviderConfig_Recreate
- [ ] jobs/TestAccJobControlRunState
- [ ] jobs/TestAccJobDisabledTask
- [ ] jobs/TestAccJobRunAsUser

### REQUEST_PARSE (7)
- [ ] permissions/TestAccPermissions_WorkspaceFile_Id — cannot create workspace file: invalid character '-' in numeric literal Error: cannot create workspace file: invalid character '-' in numeric literal Error running apply: exit status 1
- [ ] permissions/TestAccPermissions_WorkspaceFile_Path — cannot create workspace file: invalid character '-' in numeric literal Error: cannot create workspace file: invalid character '-' in numeric literal Error running apply: exit status 1
- [ ] workspace/TestAccWorkspaceFile — cannot create workspace file: invalid character '-' in numeric literal Error: cannot create workspace file: invalid character '-' in numeric literal Error running apply: exit status 1
- [ ] workspace/TestAccWorkspaceFileBase64 — cannot create workspace file: invalid character '-' in numeric literal Error: cannot create workspace file: invalid character '-' in numeric literal Error running apply: exit status 1
- [ ] workspace/TestAccWorkspaceFileCreate_NonExistentParent — cannot create workspace file: invalid character '-' in numeric literal Error: cannot create workspace file: invalid character '-' in numeric literal Error running apply: exit status 1
- [ ] workspace/TestAccWorkspaceFileEmptyFile — cannot create workspace file: invalid character '-' in numeric literal Error: cannot create workspace file: invalid character '-' in numeric literal Error running apply: exit status 1
- [ ] workspace/TestAccWorkspaceFileZipFile — cannot create workspace file: invalid character '-' in numeric literal Error: cannot create workspace file: invalid character '-' in numeric literal Error running apply: exit status 1

### 501_NOT_IMPLEMENTED — `GET /api/2.1/clusters/list-node-types` (6)
- [ ] jobs/TestAccJobClusterPolicySparkVersion
- [ ] jobs/TestAccJobCluster_ProviderConfig_Mismatched
- [ ] jobs/TestAccJobCluster_ProviderConfig_Remove
- [ ] jobs/TestAccJobTasks
- [ ] permissions/TestAccPermissions_InstancePool
- [ ] pools/TestAccInstancePool_EnableElasticDiskFalse

### 501_NOT_IMPLEMENTED — `POST /api/2.0/preview/scim/v2/ServicePrincipals` (6)
- [ ] permissions/TestAccPermissions_Directory_Path
- [ ] permissions/TestAccPermissions_Notebook_Path
- [ ] permissions/TestMwsAccAccountServicePrincipalRuleSetsFullLifeCycle
- [ ] scim/TestAccEntitlementsRemoveExisting
- [ ] scim/TestAccEntitlementsSomeTrueSomeFalse
- [ ] scim/TestAccServicePrincipalWithApiField

### 404_NO_HANDLER — `no fakebricks handler for GET /.well-known/databricks-config` (5)
- [ ] internal/providers/pluginfw/products/app/TestAccAppResource
- [ ] jobs/TestAccDataSourceJob_EmptyBlock
- [ ] jobs/TestAccPeriodicTrigger
- [ ] permissions/TestAccPermissions_Job
- [ ] sharing/TestUcAccCreateProviderDb2Open

### 501_NOT_IMPLEMENTED — `POST /api/2.0/accounts/{acct}/scim/v2/ServicePrincipals` (5)
- [ ] scim/TestMwsAccDataSourceSPNByDisplayNameOnAccount
- [ ] scim/TestMwsAccDataSourceSPNByDisplayNameOnAccountNoApi
- [ ] scim/TestMwsAccDataSourceSPNsByDisplayNameOnAccount
- [ ] scim/TestMwsAccDataSourceSPNsByDisplayNameOnAccountNoApi
- [ ] scim/TestMwsAccServicePrincipalWithApiField

### 501_NOT_IMPLEMENTED — `PATCH /api/2.0/settings/types/default_namespace_ws/names/default` (4)
- [ ] settings/TestAccDefaultNamespaceSetting
- [ ] settings/TestAccDefaultNamespaceSetting_ProviderConfig_Match
- [ ] settings/TestAccDefaultNamespaceSetting_ProviderConfig_Recreate
- [ ] settings/TestAccDefaultNamespaceSetting_ProviderConfig_Remove

### CHECK_ASSERTION (4)
- [ ] internal/acceptance/TestMwsAccWorkspaceIDApp_NoDefaultNoOverride — 2026-06-25T05:34:02.541Z [ERROR] sdk.helper_resource: Expected an error with pattern ((?s)failed to get workspace client): test_working_directory=/tmp/plugintest2329402484 test_step_number=1 test_name=TestMwsAccWorkspaceIDApp_NoDefaultNoOverride test_terraform_path=/usr/local/bin/terraform
- [ ] internal/acceptance/TestMwsAccWorkspaceIDTagPolicy_NoDefaultNoOverride — 2026-06-25T05:34:16.299Z [ERROR] sdk.helper_resource: Expected an error with pattern ((?s)failed to get workspace client): test_terraform_path=/usr/local/bin/terraform test_step_number=1
- [ ] mws/TestMwsAccDataCurrentConfig — Error Trace:	/home/pieter.noordhuis/terraform-provider-databricks/mws/data_current_config_acc_test.go:21 Error:      	Not equal: --- Expected
- [ ] mws/TestMwsAccDataCurrentConfigCloudOverride — Error Trace:	/home/pieter.noordhuis/terraform-provider-databricks/mws/data_current_config_acc_test.go:34 Error:      	Not equal: --- Expected

### NOTEBOOK_DBC (3)
- [ ] workspace/TestAccNotebookData_InState — cannot create notebook: fakebricks: valid DBC extraction is not modeled Error: cannot create notebook: fakebricks: valid DBC extraction is not modeled Error running apply: exit status 1
- [ ] workspace/TestAccNotebookData_ProviderConfig_Mismatched — cannot create notebook: fakebricks: valid DBC extraction is not modeled Error: cannot create notebook: fakebricks: valid DBC extraction is not modeled 2026-06-25T05:32:40.579Z [ERROR] sdk.helper_resource: Expected an error with pattern (workspace_id mismatch.*please check the workspace_id provided i
- [ ] workspace/TestAccNotebookResourceDbcUpdate — cannot create notebook: fakebricks: valid DBC extraction is not modeled Error: cannot create notebook: fakebricks: valid DBC extraction is not modeled Error running apply: exit status 1

### POST_APPLY_DRIFT (3)
- [ ] serving/TestAccModelServingExternalModel — After applying this test step, the non-refresh plan was not empty.
- [ ] serving/TestUcAccModelServingProvisionedThroughput — After applying this test step, the non-refresh plan was not empty.
- [ ] sql/TestAccSQLEndpoint — After applying this test step, the non-refresh plan was not empty.

### 501_NOT_IMPLEMENTED — `POST /api/2.0/accounts/{acct}/oauth2/custom-app-integrations` (2)
- [ ] apps/TestMwsAccCustomAppIntegrationCreate
- [ ] apps/TestMwsAccCustomAppIntegrationUpdate

### 501_NOT_IMPLEMENTED — `GET /api/2.1/unity-catalog/metastore_summary` (2)
- [ ] catalog/TestUcAccDataSourceCurrentMetastore
- [ ] catalog/TestUcAccResourceSystemSchema

### 501_NOT_IMPLEMENTED — `POST /api/2.1/accounts/{acct}/budgets` (2)
- [ ] finops/TestMwsAccBudgetCreate
- [ ] finops/TestMwsAccBudgetUpdate

### 501_NOT_IMPLEMENTED — `POST /api/2.0/accounts/{acct}/scim/v2/Users` (2)
- [ ] scim/TestMwsAccDataUserWithApiField
- [ ] scim/TestMwsAccUserWithApiField

### 501_NOT_IMPLEMENTED — `POST /api/2.0/accounts/{acct}/scim/v2/Groups` (2)
- [ ] scim/TestMwsAccGroupMemberWithApiField
- [ ] scim/TestMwsAccGroupWithApiField

### GROUP_ROLES (2)
- [ ] scim/TestAccGroupRole — cannot read group role: Group has no role Error: cannot read group role: Group has no role Error running apply: exit status 1
- [ ] scim/TestAccGroupRoleWithApiField — cannot read group role: Group has no role Error: cannot read group role: Group has no role Error running apply: exit status 1

### 501_NOT_IMPLEMENTED — `POST /api/2.0/ip-access-lists` (1)
- [ ] access/TestAccIPACLListsResourceFullLifecycle

### 501_NOT_IMPLEMENTED — `PUT /api/2.1/unity-catalog/artifact-allowlists/LIBRARY_MAVEN` (1)
- [ ] catalog/TestUcAccArtifactAllowlistResourceFullLifecycle

### 501_NOT_IMPLEMENTED — `PATCH /api/2.0/settings/types/aibi_dash_embed_ws_acc_policy/names/default` (1)
- [ ] internal/acceptance/TestAccAiBiEmbeddings

### 501_NOT_IMPLEMENTED — `GET /api/2.0/mlflow/registered-models/list` (1)
- [ ] mlflow/TestAccDataMlflowModels

### 501_NOT_IMPLEMENTED — `POST /api/2.0/serving-endpoints/pt` (1)
- [ ] serving/TestUcAccModelServingProvisionedThroughputResource

### 501_NOT_IMPLEMENTED — `PATCH /api/2.0/settings/types/disable_legacy_access/names/default` (1)
- [ ] settings/TestAccDisableLegacyAccessSetting

### 501_NOT_IMPLEMENTED — `PATCH /api/2.0/settings/types/disable_legacy_dbfs/names/default` (1)
- [ ] settings/TestAccDisableLegacyDbfsSetting

### 501_NOT_IMPLEMENTED — `PATCH /api/2.0/settings/types/restrict_workspace_admins/names/default` (1)
- [ ] settings/TestAccRestrictWorkspaceAdminsSetting

### 501_NOT_IMPLEMENTED — `PATCH /api/2.0/accounts/{acct}/settings/types/disable_legacy_features/names/default` (1)
- [ ] settings/TestMwsAccDisableLegacyFeaturesSetting

### 501_NOT_IMPLEMENTED — `POST /api/2.1/unity-catalog/recipients` (1)
- [ ] sharing/TestUcAccCreateRecipientDb2Open

### 501_NOT_IMPLEMENTED — `POST /api/2.0/global-init-scripts` (1)
- [ ] workspace/TestAccGlobalInitScriptResource_Create

### 501_NOT_IMPLEMENTED — `PATCH /api/2.0/workspace-conf` (1)
- [ ] workspace/TestAccWorkspaceConfFullLifecycle

## Skipped — gated, not a fakebricks failure (grouped)

- **Missing env TEST_DEFAULT_WAREHOUSE_ID variable.** (55)
- **Missing env TEST_WORKSPACE_ID variable.** (49)
- **Missing env TEST_INSTANCE_POOL_ID variable.** (18)
- **Skipping UC_ACCOUNT test in UC_WORKSPACE environment** (16)
- **Missing env TEST_METASTORE_DATA_ACCESS_ARN variable.** (14)
- **Missing env THIS_WORKSPACE_ID variable.** (14)
- **Skipping ACCOUNT/UC_ACCOUNT test in UC_WORKSPACE environment** (10)
- **Missing env TEST_WORKSPACE_URL variable.** (8)
- **Environment variable ARM_CLIENT_ID is missing** (6)
- **Environment variable TEST_EC2_INSTANCE_PROFILE is missing** (4)
- **Environment variable TEST_ROOT_BUCKET is missing** (4)
- **UNIFIED_HOST environment variable is missing** (4)
- **Environment variable ACCOUNT_LEVEL_SERVICE_PRINCIPAL_ID is missing** (4)
- **Missing env TEST_CROSSACCOUNT_ARN variable.** (4)
- **Missing env TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID variable.** (4)
- **Missing env TEST_DATA_ENG_GROUP variable.** (3)
- **Missing env TEST_METASTORE_ID variable.** (3)
- **Missing env TEST_DEFAULT_CLUSTER_ID variable.** (3)
- **Missing env GOOGLE_PROJECT variable.** (3)
- **Missing env AWS_REGION variable.** (3)
- **Skipping test because it requires Azure** (2)
- **Skipping test because it requires GCP** (2)
- **Environment variable TEST_WORKSPACE_ID is missing** (2)
- **Environment variable GOOGLE_CREDENTIALS is missing** (2)
- **Test only valid for Azure** (2)
- **Missing env DUMMY_EC2_INSTANCE_PROFILE variable.** (1)
- **Missing env TEST_METASTORE_ADMIN_GROUP_NAME variable.** (1)
- **Missing env TEST_DATA_SCI_GROUP variable.** (1)
- **Skipping this test because feature not enabled in Prod** (1)
- **There is no API to create notification destinations. Once available, add here and enable this test.** (1)
- **Missing env TEST_MANAGED_KMS_KEY_ARN variable.** (1)
- **Missing env TEST_GCP_KMS_KEY_ID variable.** (1)
- **** (1)
- **Missing env TEST_PREFIX variable.** (1)
- **TestMwsAccGcpWorkspacesProvisioningToRunning is currently only supported on GCP** (1)
- **TestMwsAccGcpWorkspacesUnsetExpectedState is currently only supported on GCP** (1)
- **Missing env TEST_LOGDELIVERY_ARN variable.** (1)
- **=== RUN   TestMwsAccPrivateAccessSettings** (1)
- **=== RUN   TestMwsAccVpcEndpoint** (1)
- **ACLs for passwords are disabled on testing workspaces** (1)
- **Only GCP service principals are treated as users** (1)
- **service principal isn't defined** (1)
- **Missing env TEST_BUCKET variable.** (1)
- **Test only runs on GCP** (1)
- **Skipping this test since users will be able to disable public DBFS root which is required for this test** (1)
- **Environment variable THIS_WORKSPACE_ID is missing** (1)
- **Missing env TEST_GLOBAL_METASTORE_ID variable.** (1)
