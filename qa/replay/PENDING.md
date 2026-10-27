# Pending acceptance scenarios

**65 original functions remain pending** in the 545-function baseline. See
[MIGRATION.md](MIGRATION.md) for counts and [FOLLOW_UPS.md](FOLLOW_UPS.md) for
blocker details. Original tests remain unchanged.

A row can describe a reproduced failure, a limitation of the current runner, or
a source-inspection risk. Inspection risks are not claims that every listed
scenario was run and failed. P17 is an inventory exception, not missing coverage.
Provider races are tracked independently of conversion; a P1 warning alone
does not require different replay fixtures or prevent writing the test.
P4 now tracks 16 cases ready for conversion after source review, not a missing
engine feature. Its SQL-warehouse import case is tracked under P5 instead.
These cases still count as pending until their replay copies pass. P5 import
support is implemented; four remaining P5 cases are also conversion backlog.

| Source | Original function | Blockers | Reason |
| --- | --- | --- | --- |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | `TestAccCatalog_ProviderConfig_Recreate` | P9 | Expected replacement conflicts with current workspace validation. |
| [catalog/connection_test.go](../../catalog/connection_test.go) | `TestUcAccConnectionsSchemaLevelResourceFullLifecycle` | P5 | Ready to convert: import support is implemented; retain the original import ID and verification settings. |
| [catalog/credential_test.go](../../catalog/credential_test.go) | `TestAccCredentialOwner` | P13 | Workspace credential configured without workspace routing. |
| [catalog/metastore_test.go](../../catalog/metastore_test.go) | `TestUcAccMetastoreWithOwnerUpdates` | P7 | Metastore name is not sent by the account update model. |
| [catalog/schema_test.go](../../catalog/schema_test.go) | `TestUcAccSchema` | P10 | Reproduced shared SDKv2 schema-slice mutation under -race. |
| [internal/acceptance/tf_provider_test.go](../../internal/acceptance/tf_provider_test.go) | `TestAccProviderPlanShouldSucceedWithIncompleteConfiguration` | P6 | Additional noop provider; already hermetic. |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | `TestMwsAccWorkspaceIDHttp_RemoveDefault` | P8 | Current diagnostic differs from the original expectation. |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | `TestMwsAccWorkspaceIDHttp_RemoveOverrideNoFallback` | P8 | Current diagnostic differs from the original expectation. |
| [internal/providers/pluginfw/products/app/resource_app_acc_test.go](../../internal/providers/pluginfw/products/app/resource_app_acc_test.go) | `TestAccAppResource` | P5 | Ready to convert: import support is implemented; retain the original import ID and verification settings. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRulePluginFrameworkOptIn` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRuleMigrationFromSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRuleMigrationToSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccUpdateQualityMonitorTransitionFromSdkV2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccUpdateQualityMonitorTransitionFromPluginFw` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccQualityMonitorImportPluginFramework` | P5 | Ready to convert: import support is implemented; retain the original import ID and verification settings. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareMigrationFromSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareMigrationFromPluginFramework` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Invalid` | P8 | Current diagnostic differs from the original expectation. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Recreate` | P9 | Expected replacement conflicts with current workspace validation. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareVolume` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | `TestAccDataSourceUsers_WithGroups` | P18 | Source configuration contradicts the checked resource/filter. |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | `TestWorkspaceDataSourceUsers_WithGroups` | P18 | Source configuration contradicts the checked resource/filter. |
| [jobs/job_test.go](../../jobs/job_test.go) | `TestAccJobControlRunState` | P4 | Ready to convert: retain all four HCL steps and exact run-now/cancel requests; omit backend scheduling polls. |
| [jobs/job_test.go](../../jobs/job_test.go) | `TestAccJobClusterPolicySparkVersion` | P12 | Original permits a nonempty post-apply plan. |
| [mws/data_mws_network_connectivity_configs_acc_test.go](../../mws/data_mws_network_connectivity_configs_acc_test.go) | `TestAccDataSourceMwsNetworkConnectivityConfigsTest` | P18 | Source configuration contradicts the checked resource/filter. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccWorkspaces` | P1/P20 | Concurrent client initialization plus derived workspace token host (inspection). |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccWorkspacesTokenUpdate` | P1/P20 | Concurrent client initialization plus derived workspace token host (inspection). |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesProvisioningToRunning` | P14 | Provisioning Read loses the input-only expected status. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesUnsetExpectedState` | P14 | Provisioning Read loses the input-only expected status. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccAwsChangeToServicePrincipal` | P16 | Requires the original per-step provider factories. |
| [mws/resource_mws_ncc_private_endpoint_rule_acc_test.go](../../mws/resource_mws_ncc_private_endpoint_rule_acc_test.go) | `TestMwsAccNccPrivateEndpointRule_NoSpuriousAccountIdDrift` | P15 | Original account_id precondition rejected by current schema. |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Job` | P1/P4 | Ready to convert: verify permission transitions, rejection, and creator-owner reset requests; omit SDK readback. P1 remains tracked. |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Pipeline` | P1/P4 | Ready to convert: retain permission transitions and rejection; current pipeline permission Delete is a no-op, unlike the old owner-reset assertion. P1 remains tracked. |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Directory_RootDirectoryCorrectlyHandlesAdminUsers` | P1/P4 | Ready to convert: verify root-ID/root-path handling and admin-only reset requests; omit SDK readback. P1 remains tracked. |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_Authorization_Tokens` | P1/P4 | Ready to convert: verify group/user transitions and admin-only reset requests; omit SDK readback. P1 remains tracked. |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_SqlWarehouses` | P1/P5 | Ready to convert: import support accepts the /warehouses/<ID> alias. Omit backend-only owner readback; P1 remains tracked. |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | `TestAccPermissions_RegisteredModel_Root` | P1/P4 | Ready to convert: verify root-model transitions, rejection, and admin-only reset requests; omit SDK readback. P1 remains tracked. |
| [permissions/resource_permissions_test.go](../../permissions/resource_permissions_test.go) | `TestAccessControlHashFunction` | P17 | Already hermetic; inventory classification. |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | `TestAccPipelineResource_CreatePipelineWithoutWorkers` | P4 | Ready to convert: require explicit num_workers: 0 and no autoscale in the create payload; omit backend SDK ForceSendFields assertion. |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | `TestAccPipelineResourcLastModified` | P19 | Read clears the configured concurrency precondition. |
| [scim/service_principal_test.go](../../scim/service_principal_test.go) | `TestAccServicePrincipalHomeDeleteSuccess` | P4 | Ready to convert: require identity deletion and recursive /Users/<application_id> cleanup in destroy fixtures; omit SDK readback. |
| [scim/user_test.go](../../scim/user_test.go) | `TestAccUserHomeDelete` | P4 | Ready to convert: require user deletion and recursive /Users/<username> cleanup in destroy fixtures; omit SDK readback. |
| [secrets/secret_acl_test.go](../../secrets/secret_acl_test.go) | `TestAccSecretAclResource` | P4 | Ready to convert: verify principal/READ ACL requests and provider state; omit backend default-owner inventory checks. |
| [secrets/secret_scope_test.go](../../secrets/secret_scope_test.go) | `TestAccSecretScopeResource` | P4 | Ready to convert: retain name/backend state checks and model external deletion in next-step refresh, requiring recreation; omit SDK ACL readback. |
| [settings/default_namespace_test.go](../../settings/default_namespace_test.go) | `TestAccDefaultNamespaceSetting` | P4 | Ready to convert: verify setting update/delete requests, etag propagation, and provider state; omit backend reset readback and SDK etag probes. |
| [settings/disable_legacy_access_setting_test.go](../../settings/disable_legacy_access_setting_test.go) | `TestAccDisableLegacyAccessSetting` | P4 | Ready to convert: verify setting update/delete requests, etag propagation, and provider state; omit backend reset readback and SDK etag probes. |
| [settings/disable_legacy_dbfs_setting_test.go](../../settings/disable_legacy_dbfs_setting_test.go) | `TestAccDisableLegacyDbfsSetting` | P4 | Ready to convert: verify setting update/delete requests, etag propagation, and provider state; omit backend reset readback and SDK etag probes. |
| [settings/disable_legacy_features_test.go](../../settings/disable_legacy_features_test.go) | `TestMwsAccDisableLegacyFeaturesSetting` | P4 | Ready to convert: verify setting update/delete requests, etag propagation, and provider state; omit backend reset readback and SDK etag probes. |
| [settings/restrict_workspace_admins_test.go](../../settings/restrict_workspace_admins_test.go) | `TestAccRestrictWorkspaceAdminsSetting` | P4 | Ready to convert: verify setting update/delete requests, etag propagation, and provider state; omit backend reset readback and SDK etag probes. |
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
