# Pending acceptance scenarios

**136 original functions remain pending** in the 545-function baseline. See
[MIGRATION.md](MIGRATION.md) for counts and [FOLLOW_UPS.md](FOLLOW_UPS.md) for
blocker details. Original tests remain unchanged.

A row can describe a reproduced failure, a limitation of the current runner, or
a source-inspection risk. Inspection risks are not claims that every listed
scenario was run and failed. P17 is an inventory exception, not missing coverage.
The GCP branch of `TestUcAccStorageCredential` is copied, but its full source
function stays pending because the AWS branch reproduces P1.

| Source | Original function | Blockers | Reason |
| --- | --- | --- | --- |
| [aws/instance_profile_test.go](../../aws/instance_profile_test.go) | `TestAccInstanceProfileIntegrationSuite` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | `TestUcAccCatalogForceDestroyConsistentAfterImport` | P5 | Import or import-state verification is part of the scenario. |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | `TestUcAccCatalogHmsConnectionUpdate` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | `TestAccCatalog_ProviderConfig_Recreate` | P9 | Expected replacement conflicts with current workspace validation. |
| [catalog/connection_test.go](../../catalog/connection_test.go) | `TestUcAccConnectionsSchemaLevelResourceFullLifecycle` | P5 | Import or import-state verification is part of the scenario. |
| [catalog/credential_test.go](../../catalog/credential_test.go) | `TestAccCredentialOwner` | P13 | Workspace credential configured without workspace routing. |
| [catalog/grant_test.go](../../catalog/grant_test.go) | `TestUcAccGrant` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [catalog/grants_test.go](../../catalog/grants_test.go) | `TestUcAccGrants` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [catalog/metastore_test.go](../../catalog/metastore_test.go) | `TestUcAccMetastoreWithOwnerUpdates` | P7 | Metastore name is not sent by the account update model. |
| [catalog/schema_test.go](../../catalog/schema_test.go) | `TestUcAccSchema` | P10 | Reproduced shared SDKv2 schema-slice mutation under -race. |
| [catalog/storage_credential_test.go](../../catalog/storage_credential_test.go) | `TestUcAccStorageCredential` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [catalog/workspace_binding_test.go](../../catalog/workspace_binding_test.go) | `TestUcAccWorkspaceBindingToOtherWorkspace` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | `TestAccClusterAPI_ListNodeTypesAzureUcws` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | `TestAccClusterAPI_ListNodeTypesAzure` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | `TestAccClusterAPI_ListNodeTypesGcpUcws` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | `TestAccClusterAPI_ListNodeTypesGcp` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | `TestAccCluster_ListNodeTypesAWSUcws` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | `TestAccCluster_ListNodeTypesAWS` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [dashboards/dashboard_test.go](../../dashboards/dashboard_test.go) | `TestAccDashboardWithFilePath` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [dashboards/dashboard_test.go](../../dashboards/dashboard_test.go) | `TestAccDashboardTestAll` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [internal/acceptance/tf_provider_test.go](../../internal/acceptance/tf_provider_test.go) | `TestAccProviderPlanShouldSucceedWithIncompleteConfiguration` | P6 | Additional noop provider; already hermetic. |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | `TestMwsAccWorkspaceIDHttp_RemoveDefault` | P8 | Current diagnostic differs from the original expectation. |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | `TestMwsAccWorkspaceIDHttp_RemoveOverrideNoFallback` | P8 | Current diagnostic differs from the original expectation. |
| [internal/providers/pluginfw/products/app/resource_app_acc_test.go](../../internal/providers/pluginfw/products/app/resource_app_acc_test.go) | `TestAccAppResource` | P5 | Import or import-state verification is part of the scenario. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRulePluginFrameworkOptIn` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRuleMigrationFromSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRuleMigrationToSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccUpdateQualityMonitorTransitionFromSdkV2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccUpdateQualityMonitorTransitionFromPluginFw` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccQualityMonitorImportPluginFramework` | P5 | Import or import-state verification is part of the scenario. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareMigrationFromSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareMigrationFromPluginFramework` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Invalid` | P8 | Current diagnostic differs from the original expectation. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Recreate` | P9 | Expected replacement conflicts with current workspace validation. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareVolume` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | `TestWorkspaceDataSourceDataUsers` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | `TestAccDataSourceUsers_WithGroups` | P18 | Source configuration contradicts the checked resource/filter. |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | `TestWorkspaceDataSourceUsers_WithGroups` | P18 | Source configuration contradicts the checked resource/filter. |
| [jobs/job_test.go](../../jobs/job_test.go) | `TestAccJobControlRunState` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [jobs/job_test.go](../../jobs/job_test.go) | `TestAccJobRunAsUser` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [jobs/job_test.go](../../jobs/job_test.go) | `TestAccJobClusterPolicySparkVersion` | P12 | Original permits a nonempty post-apply plan. |
| [mws/data_mws_network_connectivity_configs_acc_test.go](../../mws/data_mws_network_connectivity_configs_acc_test.go) | `TestAccDataSourceMwsNetworkConnectivityConfigsTest` | P18 | Source configuration contradicts the checked resource/filter. |
| [mws/mws_log_delivery_test.go](../../mws/mws_log_delivery_test.go) | `TestMwsAccLogDelivery` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccWorkspaces` | P1/P20 | Concurrent client initialization plus derived workspace token host (inspection). |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccWorkspacesTokenUpdate` | P1/P20 | Concurrent client initialization plus derived workspace token host (inspection). |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesProvisioningToRunning` | P14 | Provisioning Read loses the input-only expected status. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesUnsetExpectedState` | P14 | Provisioning Read loses the input-only expected status. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccAwsChangeToServicePrincipal` | P16 | Requires the original per-step provider factories. |
| [mws/resource_mws_ncc_private_endpoint_rule_acc_test.go](../../mws/resource_mws_ncc_private_endpoint_rule_acc_test.go) | `TestMwsAccNccPrivateEndpointRule_NoSpuriousAccountIdDrift` | P15 | Original account_id precondition rejected by current schema. |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_ClusterPolicy` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_InstancePool` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Job` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Pipeline` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Notebook_Path` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Notebook_Id` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Directory_Path` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Directory_Id` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Directory_RootDirectoryCorrectlyHandlesAdminUsers` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_WorkspaceFile_Path` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_WorkspaceFile_Id` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Repo_Id` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Repo_Path` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Authorization_Passwords` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Authorization_Tokens` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_SqlWarehouses` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_SqlAlert` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_SqlQuery` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Dashboard` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Experiment` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_RegisteredModel` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_RegisteredModel_Root` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_ServingEndpoint` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Alert` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Query` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_App` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestUcAccPermissions_DatabaseProject` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [permissions/resource_permissions_test.go](../../permissions/resource_permissions_test.go) | `TestAccessControlHashFunction` | P17 | Already hermetic; inventory classification. |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | `TestAccPipelineResource_CreatePipelineWithoutWorkers` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | `TestAccPipelineResourcLastModified` | P19 | Read clears the configured concurrency precondition. |
| [scim/data_group_acc_test.go](../../scim/data_group_acc_test.go) | `TestAccGroupDataSplitMembers` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [scim/group_member_test.go](../../scim/group_member_test.go) | `TestAccGroupMemberResource` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [scim/service_principal_test.go](../../scim/service_principal_test.go) | `TestAccServicePrincipalHomeDeleteSuccess` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [scim/user_test.go](../../scim/user_test.go) | `TestAccUserHomeDelete` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [scim/user_test.go](../../scim/user_test.go) | `TestAccUserResource` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [secrets/secret_acl_test.go](../../secrets/secret_acl_test.go) | `TestAccSecretAclResource` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [secrets/secret_scope_test.go](../../secrets/secret_scope_test.go) | `TestAccSecretScopeResource` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [settings/default_namespace_test.go](../../settings/default_namespace_test.go) | `TestAccDefaultNamespaceSetting` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [settings/disable_legacy_access_setting_test.go](../../settings/disable_legacy_access_setting_test.go) | `TestAccDisableLegacyAccessSetting` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [settings/disable_legacy_dbfs_setting_test.go](../../settings/disable_legacy_dbfs_setting_test.go) | `TestAccDisableLegacyDbfsSetting` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [settings/disable_legacy_features_test.go](../../settings/disable_legacy_features_test.go) | `TestMwsAccDisableLegacyFeaturesSetting` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [settings/restrict_workspace_admins_test.go](../../settings/restrict_workspace_admins_test.go) | `TestAccRestrictWorkspaceAdminsSetting` | P4 | Retain SDK-only assertions/mutations unavailable to current YAML checks. |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_Managed` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTableWithIdentityColumn_Managed` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_WarehousePartition` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_Liquid` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_RenameColumn` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_ColumnTypeSuppressDiff` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_AddColumnComment` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_DropColumnNullable` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_MultipleColumnUpdates` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_ChangeColumnTypeThrows` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_DropColumn` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_DropMultipleColumns` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_AddColumn` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_AddMultipleColumns` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_AddColumnAndUpdateThrows` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | `TestUcAccResourceSqlTable_DropColumnAndUpdateThrows` | P11 | Multi-property SQL, including the shared template, has map-order risk (inspection). |
| [storage/dbfs_file_test.go](../../storage/dbfs_file_test.go) | `TestAccDatabricksDBFSFile_CreateViaSource` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [storage/file_test.go](../../storage/file_test.go) | `TestUcAccFileDontUpdateIfNoChange` | P2/P3 | Raw file upload/download and local input files. |
| [storage/file_test.go](../../storage/file_test.go) | `TestUcAccFileUpdateOnLocalContentChange` | P2/P3 | Raw file upload/download and local input files. |
| [storage/file_test.go](../../storage/file_test.go) | `TestUcAccFileUpdateOnLocalFileChange` | P2/P3 | Raw file upload/download and local input files. |
| [storage/file_test.go](../../storage/file_test.go) | `TestUcAccFileNoUpdateIfFileDoesNotChange` | P2/P3 | Raw file upload/download and local input files. |
| [storage/file_test.go](../../storage/file_test.go) | `TestUcAccFileUpdateServerChange` | P2/P3 | Raw file upload/download and local input files. |
| [storage/file_test.go](../../storage/file_test.go) | `TestUcAccFileFullLifeCycle` | P2/P3 | Raw file upload/download and local input files. |
| [storage/file_test.go](../../storage/file_test.go) | `TestUcAccFileBase64FullLifeCycle` | P2/P3 | Raw file upload/download and local input files. |
| [tokens/obo_token_test.go](../../tokens/obo_token_test.go) | `TestAccAwsOboTokenResource` | P1 | Concurrent workspace client initialization; reproduced or shares the same unsafe cache (see follow-ups). |
| [workspace/data_notebook_acc_test.go](../../workspace/data_notebook_acc_test.go) | `TestAccNotebookData_ProviderConfig_Mismatched` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/data_notebook_acc_test.go](../../workspace/data_notebook_acc_test.go) | `TestAccNotebookData_ProviderConfig_EmptyID` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/data_notebook_acc_test.go](../../workspace/data_notebook_acc_test.go) | `TestAccNotebookData_InState` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | `TestAccNotebookResourceScalability` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | `TestAccNotebookResourceDbcUpdate` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | `TestAccNotebookResourceJupiterUpdate` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | `TestAccNotebook_ProviderConfig_Mismatched` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | `TestAccNotebook_ProviderConfig_Match` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | `TestAccNotebook_ProviderConfig_Recreate` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | `TestAccNotebook_ProviderConfig_Remove` | P3 | Original source/file_path needs a local file or PreConfig hook. |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | `TestAccWorkspaceConf_DeleteInvalidKey` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | `TestAccWorkspaceConf_GetInvalidKey` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | `TestAccWorkspaceConf_GetOnlyInvalidKeys` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | `TestAccWorkspaceConf_GetValidKey` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | `TestAccWorkspaceConf_SetInvalidKey` | P21 | Go helper replay is reviewed in the stacked client-helper PR. |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | `TestAccWorkspaceFile` | P2/P3 | Multipart upload; source variants also need local files. |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | `TestAccWorkspaceFileEmptyFile` | P2/P3 | Multipart upload; source variants also need local files. |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | `TestAccWorkspaceFileZipFile` | P2/P3 | Multipart upload; source variants also need local files. |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | `TestAccWorkspaceFileCreate_NonExistentParent` | P2/P3 | Multipart upload; source variants also need local files. |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | `TestAccWorkspaceFileBase64` | P2/P3 | Multipart upload; source variants also need local files. |
