# Pending acceptance scenarios

**38 original functions remain pending** in the 545-function baseline. See
[MIGRATION.md](MIGRATION.md) for counts and [FOLLOW_UPS.md](FOLLOW_UPS.md) for
blocker details. Original tests remain unchanged.

A row can describe a reproduced failure, a limitation of the current runner, or
a source-inspection risk. Inspection risks are not claims that every listed
scenario was run and failed. P17 is an inventory exception, not missing coverage.
Provider races are tracked independently of conversion; a P1 warning alone
does not require different replay fixtures or prevent writing the test.
P4 now tracks 16 cases ready for conversion after source review, not a missing
engine feature. Its SQL-warehouse import case is tracked under P5 instead.
The P4 cases still count as pending until their replay copies pass.
All five P5 import cases now have replay copies.

| Source | Original function | Blockers | Reason |
| --- | --- | --- | --- |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | `TestAccCatalog_ProviderConfig_Recreate` | P9 | Expected replacement conflicts with current workspace validation. |
| [catalog/credential_test.go](../../catalog/credential_test.go) | `TestAccCredentialOwner` | P13 | Workspace credential configured without workspace routing. |
| [internal/acceptance/tf_provider_test.go](../../internal/acceptance/tf_provider_test.go) | `TestAccProviderPlanShouldSucceedWithIncompleteConfiguration` | P6 | Additional noop provider; already hermetic. |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | `TestMwsAccWorkspaceIDHttp_RemoveDefault` | P8 | Current diagnostic differs from the original expectation. |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | `TestMwsAccWorkspaceIDHttp_RemoveOverrideNoFallback` | P8 | Current diagnostic differs from the original expectation. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRulePluginFrameworkOptIn` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRuleMigrationFromSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | `TestMwsAccPrivateEndpointRuleMigrationToSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccUpdateQualityMonitorTransitionFromSdkV2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | `TestUcAccUpdateQualityMonitorTransitionFromPluginFw` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareMigrationFromSDKv2` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareMigrationFromPluginFramework` | P16 | Requires the original per-step provider factories. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Invalid` | P8 | Current diagnostic differs from the original expectation. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Recreate` | P9 | Expected replacement conflicts with current workspace validation. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestUcAccShareVolume` | P16 | Requires the original per-step provider factories. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccWorkspaces` | P1/P20 | Concurrent client initialization plus derived workspace token host (inspection). |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccWorkspacesTokenUpdate` | P1/P20 | Concurrent client initialization plus derived workspace token host (inspection). |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesProvisioningToRunning` | P14 | Provisioning Read loses the input-only expected status. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesUnsetExpectedState` | P14 | Provisioning Read loses the input-only expected status. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccAwsChangeToServicePrincipal` | P16 | Requires the original per-step provider factories. |
| [permissions/resource_permissions_test.go](../../permissions/resource_permissions_test.go) | `TestAccessControlHashFunction` | P17 | Already hermetic; inventory classification. |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | `TestAccPipelineResourcLastModified` | P19 | Read clears the configured concurrency precondition. |
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
