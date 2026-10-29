# Acceptance-to-replay inventory

Initial AST inventory: **545 acceptance entry points in 176 files**. Detection
uses acceptance test prefixes or direct calls to `acceptance.*Level`. This is a
migration baseline, not an exhaustive list of every indirect wrapper or hermetic
Terraform test. Some inventoried cases already use local mocks.

**520 copied; 25 pending.** Counts refer to original Go test functions, not
individual steps or cloud permutations. Copies preserve provider-owned behavior
and meaningful assertions. Backend-only persistence checks are omitted; fixture
requests verify what the provider sends. Unsupported provider behavior remains
pending. Original files are unchanged.

Suite assignment requires reading each scenario. In particular, do not infer
the subject from the first HCL resource: it may be a prerequisite.

| Original source file | Copied | Pending |
| --- | ---: | ---: |
| [access/ip_access_list_test.go](../../access/ip_access_list_test.go) | 1 | 0 |
| [access/sql_permissions_provider_config_test.go](../../access/sql_permissions_provider_config_test.go) | 2 | 0 |
| [apps/custom_app_integration_test.go](../../apps/custom_app_integration_test.go) | 2 | 0 |
| [aws/data_aws_crossaccount_policy_acc_test.go](../../aws/data_aws_crossaccount_policy_acc_test.go) | 1 | 0 |
| [aws/data_instance_profiles_acc_test.go](../../aws/data_instance_profiles_acc_test.go) | 1 | 0 |
| [aws/instance_profile_provider_config_test.go](../../aws/instance_profile_provider_config_test.go) | 1 | 0 |
| [aws/instance_profile_test.go](../../aws/instance_profile_test.go) | 1 | 0 |
| [aws/service_principal_role_provider_config_test.go](../../aws/service_principal_role_provider_config_test.go) | 1 | 0 |
| [aws/user_role_provider_config_test.go](../../aws/user_role_provider_config_test.go) | 1 | 0 |
| [aws/user_role_test.go](../../aws/user_role_test.go) | 2 | 0 |
| [catalog/artifact_allowlist_test.go](../../catalog/artifact_allowlist_test.go) | 1 | 0 |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | 9 | 1 |
| [catalog/connection_test.go](../../catalog/connection_test.go) | 3 | 0 |
| [catalog/credential_provider_config_test.go](../../catalog/credential_provider_config_test.go) | 2 | 0 |
| [catalog/credential_test.go](../../catalog/credential_test.go) | 1 | 1 |
| [catalog/data_catalog_acc_test.go](../../catalog/data_catalog_acc_test.go) | 1 | 0 |
| [catalog/data_current_metastore_acc_test.go](../../catalog/data_current_metastore_acc_test.go) | 1 | 0 |
| [catalog/data_external_locations_acc_test.go](../../catalog/data_external_locations_acc_test.go) | 1 | 0 |
| [catalog/data_metastore_acc_test.go](../../catalog/data_metastore_acc_test.go) | 1 | 0 |
| [catalog/data_metastores_acc_test.go](../../catalog/data_metastores_acc_test.go) | 1 | 0 |
| [catalog/data_schema_acc_test.go](../../catalog/data_schema_acc_test.go) | 1 | 0 |
| [catalog/data_schemas_acc_test.go](../../catalog/data_schemas_acc_test.go) | 1 | 0 |
| [catalog/data_storage_credential_acc_test.go](../../catalog/data_storage_credential_acc_test.go) | 1 | 0 |
| [catalog/data_storage_credentials_acc_test.go](../../catalog/data_storage_credentials_acc_test.go) | 1 | 0 |
| [catalog/data_table_acc_test.go](../../catalog/data_table_acc_test.go) | 1 | 0 |
| [catalog/data_tables_acc_test.go](../../catalog/data_tables_acc_test.go) | 1 | 0 |
| [catalog/data_volume_acc_test.go](../../catalog/data_volume_acc_test.go) | 1 | 0 |
| [catalog/data_volumes_acc_test.go](../../catalog/data_volumes_acc_test.go) | 1 | 0 |
| [catalog/external_location_test.go](../../catalog/external_location_test.go) | 3 | 0 |
| [catalog/grant_test.go](../../catalog/grant_test.go) | 5 | 0 |
| [catalog/grants_test.go](../../catalog/grants_test.go) | 6 | 0 |
| [catalog/metastore_assignment_provider_config_test.go](../../catalog/metastore_assignment_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_assignment_test.go](../../catalog/metastore_assignment_test.go) | 2 | 0 |
| [catalog/metastore_data_access_provider_config_test.go](../../catalog/metastore_data_access_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_data_access_test.go](../../catalog/metastore_data_access_test.go) | 2 | 0 |
| [catalog/metastore_provider_config_test.go](../../catalog/metastore_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_test.go](../../catalog/metastore_test.go) | 4 | 0 |
| [catalog/quality_monitor_test.go](../../catalog/quality_monitor_test.go) | 2 | 0 |
| [catalog/registered_model_test.go](../../catalog/registered_model_test.go) | 1 | 0 |
| [catalog/resource_sql_table_acc_test.go](../../catalog/resource_sql_table_acc_test.go) | 2 | 0 |
| [catalog/schema_test.go](../../catalog/schema_test.go) | 2 | 0 |
| [catalog/storage_credential_provider_config_test.go](../../catalog/storage_credential_provider_config_test.go) | 2 | 0 |
| [catalog/storage_credential_test.go](../../catalog/storage_credential_test.go) | 2 | 0 |
| [catalog/system_schema_test.go](../../catalog/system_schema_test.go) | 1 | 0 |
| [catalog/table_provider_config_test.go](../../catalog/table_provider_config_test.go) | 1 | 0 |
| [catalog/volume_test.go](../../catalog/volume_test.go) | 2 | 0 |
| [catalog/workspace_binding_test.go](../../catalog/workspace_binding_test.go) | 1 | 0 |
| [clusters/cluster_policy_test.go](../../clusters/cluster_policy_test.go) | 3 | 0 |
| [clusters/data_cluster_acc_test.go](../../clusters/data_cluster_acc_test.go) | 1 | 0 |
| [clusters/data_clusters_acc_test.go](../../clusters/data_clusters_acc_test.go) | 3 | 0 |
| [clusters/data_spark_version_acc_test.go](../../clusters/data_spark_version_acc_test.go) | 1 | 0 |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | 6 | 0 |
| [dashboards/dashboard_test.go](../../dashboards/dashboard_test.go) | 8 | 0 |
| [finops/budget_test.go](../../finops/budget_test.go) | 2 | 0 |
| [internal/acceptance/aibi_embedding_settings_test.go](../../internal/acceptance/aibi_embedding_settings_test.go) | 1 | 0 |
| [internal/acceptance/tf_provider_test.go](../../internal/acceptance/tf_provider_test.go) | 0 | 1 |
| [internal/acceptance/uc_resources_workspace_id_acc_test.go](../../internal/acceptance/uc_resources_workspace_id_acc_test.go) | 6 | 0 |
| [internal/acceptance/unified_host_acc_test.go](../../internal/acceptance/unified_host_acc_test.go) | 6 | 0 |
| [internal/acceptance/workspace_id_pf_app_acc_test.go](../../internal/acceptance/workspace_id_pf_app_acc_test.go) | 21 | 0 |
| [internal/acceptance/workspace_id_pf_tag_policy_acc_test.go](../../internal/acceptance/workspace_id_pf_tag_policy_acc_test.go) | 21 | 0 |
| [internal/acceptance/workspace_id_sdkv2_gosdk_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_gosdk_acc_test.go) | 21 | 0 |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | 21 | 0 |
| [internal/providers/pluginfw/products/app/data_app_acc_test.go](../../internal/providers/pluginfw/products/app/data_app_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/app/data_apps_acc_test.go](../../internal/providers/pluginfw/products/app/data_apps_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/app/resource_app_acc_test.go](../../internal/providers/pluginfw/products/app/resource_app_acc_test.go) | 8 | 0 |
| [internal/providers/pluginfw/products/cluster/data_cluster_acc_test.go](../../internal/providers/pluginfw/products/cluster/data_cluster_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/dashboards/data_dashboards_acc_test.go](../../internal/providers/pluginfw/products/dashboards/data_dashboards_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | 3 | 0 |
| [internal/providers/pluginfw/products/notificationdestinations/data_notification_destinations_acc_test.go](../../internal/providers/pluginfw/products/notificationdestinations/data_notification_destinations_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | 6 | 0 |
| [internal/providers/pluginfw/products/sharing/data_recipients_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_recipients_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/sharing/data_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_share_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/sharing/data_shares_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_shares_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | 23 | 2 |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | 8 | 0 |
| [internal/providers/pluginfw/products/volume/data_volumes_acc_test.go](../../internal/providers/pluginfw/products/volume/data_volumes_acc_test.go) | 1 | 0 |
| [jobs/data_job_acc_test.go](../../jobs/data_job_acc_test.go) | 5 | 0 |
| [jobs/data_jobs_acc_test.go](../../jobs/data_jobs_acc_test.go) | 4 | 0 |
| [jobs/job_test.go](../../jobs/job_test.go) | 16 | 0 |
| [mlflow/data_mlflow_experiment_acc_test.go](../../mlflow/data_mlflow_experiment_acc_test.go) | 1 | 0 |
| [mlflow/data_mlflow_model_acc_test.go](../../mlflow/data_mlflow_model_acc_test.go) | 1 | 0 |
| [mlflow/data_mlflow_models_acc_test.go](../../mlflow/data_mlflow_models_acc_test.go) | 1 | 0 |
| [mlflow/mlflow_experiment_test.go](../../mlflow/mlflow_experiment_test.go) | 2 | 0 |
| [mlflow/mlflow_model_test.go](../../mlflow/mlflow_model_test.go) | 1 | 0 |
| [mws/data_current_config_acc_test.go](../../mws/data_current_config_acc_test.go) | 8 | 0 |
| [mws/data_mws_credentials_acc_test.go](../../mws/data_mws_credentials_acc_test.go) | 1 | 0 |
| [mws/data_mws_network_connectivity_config_acc_test.go](../../mws/data_mws_network_connectivity_config_acc_test.go) | 1 | 0 |
| [mws/data_mws_network_connectivity_configs_acc_test.go](../../mws/data_mws_network_connectivity_configs_acc_test.go) | 1 | 0 |
| [mws/data_mws_workspaces_acc_test.go](../../mws/data_mws_workspaces_acc_test.go) | 1 | 0 |
| [mws/mws_credentials_test.go](../../mws/mws_credentials_test.go) | 1 | 0 |
| [mws/mws_customer_managed_keys_test.go](../../mws/mws_customer_managed_keys_test.go) | 2 | 0 |
| [mws/mws_log_delivery_test.go](../../mws/mws_log_delivery_test.go) | 1 | 0 |
| [mws/mws_network_connectivity_config_test.go](../../mws/mws_network_connectivity_config_test.go) | 1 | 0 |
| [mws/mws_networks_test.go](../../mws/mws_networks_test.go) | 2 | 0 |
| [mws/mws_permissionassignments_test.go](../../mws/mws_permissionassignments_test.go) | 2 | 0 |
| [mws/mws_private_access_settings_test.go](../../mws/mws_private_access_settings_test.go) | 2 | 0 |
| [mws/mws_storage_configurations_test.go](../../mws/mws_storage_configurations_test.go) | 1 | 0 |
| [mws/mws_vpc_endpoint_test.go](../../mws/mws_vpc_endpoint_test.go) | 2 | 0 |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | 7 | 2 |
| [mws/resource_mws_ncc_private_endpoint_rule_acc_test.go](../../mws/resource_mws_ncc_private_endpoint_rule_acc_test.go) | 1 | 0 |
| [permissions/access_control_rule_set_provider_config_test.go](../../permissions/access_control_rule_set_provider_config_test.go) | 1 | 0 |
| [permissions/account_rule_set_test.go](../../permissions/account_rule_set_test.go) | 2 | 0 |
| [permissions/permissions_provider_config_test.go](../../permissions/permissions_provider_config_test.go) | 1 | 0 |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | 27 | 0 |
| [permissions/resource_permissions_test.go](../../permissions/resource_permissions_test.go) | 0 | 1 |
| [pipelines/data_pipelines_acc_test.go](../../pipelines/data_pipelines_acc_test.go) | 1 | 0 |
| [pipelines/pipeline_provider_config_test.go](../../pipelines/pipeline_provider_config_test.go) | 1 | 0 |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | 6 | 1 |
| [policies/cluster_policy_provider_config_test.go](../../policies/cluster_policy_provider_config_test.go) | 1 | 0 |
| [pools/instance_pool_provider_config_test.go](../../pools/instance_pool_provider_config_test.go) | 1 | 0 |
| [pools/resource_instance_pool_acc_test.go](../../pools/resource_instance_pool_acc_test.go) | 1 | 0 |
| [repos/git_credential_test.go](../../repos/git_credential_test.go) | 1 | 0 |
| [repos/repo_provider_config_test.go](../../repos/repo_provider_config_test.go) | 1 | 0 |
| [scim/data_group_acc_test.go](../../scim/data_group_acc_test.go) | 2 | 0 |
| [scim/data_service_principal_acc_test.go](../../scim/data_service_principal_acc_test.go) | 5 | 0 |
| [scim/data_service_principals_acc_test.go](../../scim/data_service_principals_acc_test.go) | 5 | 0 |
| [scim/data_user_acc_test.go](../../scim/data_user_acc_test.go) | 4 | 0 |
| [scim/entitlement_provider_config_test.go](../../scim/entitlement_provider_config_test.go) | 2 | 0 |
| [scim/group_member_provider_config_test.go](../../scim/group_member_provider_config_test.go) | 1 | 0 |
| [scim/group_member_test.go](../../scim/group_member_test.go) | 3 | 0 |
| [scim/group_provider_config_test.go](../../scim/group_provider_config_test.go) | 1 | 0 |
| [scim/group_role_provider_config_test.go](../../scim/group_role_provider_config_test.go) | 1 | 0 |
| [scim/group_role_test.go](../../scim/group_role_test.go) | 2 | 0 |
| [scim/group_test.go](../../scim/group_test.go) | 6 | 0 |
| [scim/service_principal_provider_config_test.go](../../scim/service_principal_provider_config_test.go) | 1 | 0 |
| [scim/service_principal_test.go](../../scim/service_principal_test.go) | 8 | 0 |
| [scim/user_provider_config_test.go](../../scim/user_provider_config_test.go) | 1 | 0 |
| [scim/user_test.go](../../scim/user_test.go) | 8 | 0 |
| [secrets/secret_acl_test.go](../../secrets/secret_acl_test.go) | 2 | 0 |
| [secrets/secret_scope_provider_config_test.go](../../secrets/secret_scope_provider_config_test.go) | 1 | 0 |
| [secrets/secret_scope_test.go](../../secrets/secret_scope_test.go) | 2 | 0 |
| [secrets/secret_test.go](../../secrets/secret_test.go) | 1 | 0 |
| [serving/model_serving_test.go](../../serving/model_serving_test.go) | 4 | 0 |
| [settings/default_namespace_test.go](../../settings/default_namespace_test.go) | 4 | 0 |
| [settings/disable_legacy_access_setting_test.go](../../settings/disable_legacy_access_setting_test.go) | 1 | 0 |
| [settings/disable_legacy_dbfs_setting_test.go](../../settings/disable_legacy_dbfs_setting_test.go) | 1 | 0 |
| [settings/disable_legacy_features_test.go](../../settings/disable_legacy_features_test.go) | 1 | 0 |
| [settings/notification_destination_provider_config_test.go](../../settings/notification_destination_provider_config_test.go) | 1 | 0 |
| [settings/notification_destination_test.go](../../settings/notification_destination_test.go) | 5 | 0 |
| [settings/restrict_workspace_admins_test.go](../../settings/restrict_workspace_admins_test.go) | 1 | 0 |
| [sharing/data_shares_acc_test.go](../../sharing/data_shares_acc_test.go) | 1 | 0 |
| [sharing/provider_provider_config_test.go](../../sharing/provider_provider_config_test.go) | 1 | 0 |
| [sharing/provider_test.go](../../sharing/provider_test.go) | 1 | 0 |
| [sharing/recipient_test.go](../../sharing/recipient_test.go) | 3 | 0 |
| [sharing/share_test.go](../../sharing/share_test.go) | 2 | 0 |
| [sql/alert_test.go](../../sql/alert_test.go) | 1 | 0 |
| [sql/data_sql_warehouse_acc_test.go](../../sql/data_sql_warehouse_acc_test.go) | 1 | 0 |
| [sql/query_test.go](../../sql/query_test.go) | 1 | 0 |
| [sql/sql_alert_test.go](../../sql/sql_alert_test.go) | 1 | 0 |
| [sql/sql_dashboard_provider_config_test.go](../../sql/sql_dashboard_provider_config_test.go) | 1 | 0 |
| [sql/sql_endpoint_test.go](../../sql/sql_endpoint_test.go) | 1 | 0 |
| [sql/sql_global_config_provider_config_test.go](../../sql/sql_global_config_provider_config_test.go) | 1 | 0 |
| [sql/sql_global_config_test.go](../../sql/sql_global_config_test.go) | 4 | 0 |
| [sql/sql_permissions_test.go](../../sql/sql_permissions_test.go) | 1 | 0 |
| [sql/sql_query_provider_config_test.go](../../sql/sql_query_provider_config_test.go) | 1 | 0 |
| [sql/sql_query_test.go](../../sql/sql_query_test.go) | 1 | 0 |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | 2 | 16 |
| [sql/sql_visualization_provider_config_test.go](../../sql/sql_visualization_provider_config_test.go) | 1 | 0 |
| [sql/sql_widget_provider_config_test.go](../../sql/sql_widget_provider_config_test.go) | 1 | 0 |
| [storage/dbfs_file_provider_config_test.go](../../storage/dbfs_file_provider_config_test.go) | 1 | 0 |
| [storage/dbfs_file_test.go](../../storage/dbfs_file_test.go) | 2 | 0 |
| [storage/file_test.go](../../storage/file_test.go) | 7 | 0 |
| [storage/mount_provider_config_test.go](../../storage/mount_provider_config_test.go) | 1 | 0 |
| [tokens/obo_token_provider_config_test.go](../../tokens/obo_token_provider_config_test.go) | 1 | 0 |
| [tokens/obo_token_test.go](../../tokens/obo_token_test.go) | 1 | 0 |
| [tokens/service_principal_secret_provider_config_test.go](../../tokens/service_principal_secret_provider_config_test.go) | 1 | 0 |
| [tokens/token_provider_config_test.go](../../tokens/token_provider_config_test.go) | 1 | 0 |
| [tokens/token_test.go](../../tokens/token_test.go) | 1 | 0 |
| [vectorsearch/vector_search_test.go](../../vectorsearch/vector_search_test.go) | 1 | 0 |
| [workspace/data_notebook_acc_test.go](../../workspace/data_notebook_acc_test.go) | 3 | 0 |
| [workspace/directory_test.go](../../workspace/directory_test.go) | 1 | 0 |
| [workspace/global_init_script_test.go](../../workspace/global_init_script_test.go) | 1 | 0 |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | 9 | 0 |
| [workspace/workspace_conf_provider_config_test.go](../../workspace/workspace_conf_provider_config_test.go) | 1 | 0 |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | 6 | 0 |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | 5 | 0 |

## Known adaptations and gaps

See [PENDING.md](PENDING.md) for the remaining per-test blockers and
[FOLLOW_UPS.md](FOLLOW_UPS.md) for details and reproduction status.

- Workspace routing copies directory, app, tag-policy, and notebook cases into
  `features/workspace_routing/`, split by carrier. Preserve workspace-ID state
  checks, no-op/replacement plan actions, and diagnostic patterns. Routing headers
  distinguish requests for the old and new workspaces. Step-local dummy metadata
  models account/workspace clients and host transitions on the local server.
  YAML aliases repeat consumable fixtures; they do not make calls reusable.
  The provider-upgrade cases retain the originals' current-provider approximation.
  Removed defaults/overrides retain diagnostic checks. Notebook removal refreshes
  the workspace saved in state before rejecting the missing workspace fallback.
- Current-config cloud branches use separate YAML files for AWS, Azure, and
  GCP but still count as one source function. Explicit API cases extend
  `features/api_selection`; inference and cloud overrides stay in the
  `current_config` resource suite.
- Compute cases copy cluster-policy rename/family overrides, cluster lookup
  and name filtering, Spark-version selection, instance-profile listing, and
  global-init-script creation. With fixed Spark-version responses, the selected
  ID assertion includes the original Scala 2.13 suffix requirement.
  Built-in policy cleanup resets its override instead of deleting the policy.
- `TestAccDataSourceClusters_FilterBy` adds a read-only cluster data source to
  check the same source/state/policy fields as the SDK reads. The pinned-filter
  step and exact filter requests remain intact.
- Workspace-file and volume-file suites match multipart and raw upload bytes.
  Source files are staged inside Terraform's isolated working directory; later
  steps can replace their contents without changing the configured path.
- Identity cases add user-name normalization, service-principal lifecycles,
  user/service-principal data sources, and group roles. Explicit `api` cases
  belong to `features/api_selection`; other cases stay with their resource.
  User data-source/resource ID equality is checked against the same fixture ID.
- Azure service-principal templates generate a new application ID for each
  step, so their second steps replace the resource. AWS display-name changes
  update it in place. Account destroy preserves the default disable operation.
- The three independent users in `TestAccUserResource` expose an existing race
  in `common/unified_provider.go`'s shared client cache under `-race`. That case
  is copied independently of the P1 fix; no artificial resource dependencies
  are added. Home preservation is copied through an explicit
  identity-removal step and directory data-source read; forced home cleanup
  remains pending but needs no new engine feature after the P4 review.
- Token timestamps use a fixed future window with the original lifetime, so
  wall-clock expiry does not invalidate its otherwise static lifecycle transcript.
- Notification-destination SDK checks of display name and type become
  provider-read state checks. Inferred responses return `*_set` flags for
  input-only fields, exercising preservation of the original configured values.
- The group suite copies workspace/account display-name updates and external SCIM
  changes from `scim/group_test.go`. SDK name reads become provider-read state
  checks; external mutations become next-step refresh responses. The two explicit
  API-selection cases are copied into the separate `features/api_selection` suite.
- Negative provider-config plan-only cases fail before apply or any API call.
  They use `expect_error` without `plan_only` because the replay runner rejects
  that combination. The original HCL and error patterns are preserved.
- `secrets/secret_scope_test.go` can retain name/backend checks and recreation
  after external deletion using next-step refresh fixtures. Its backend-only
  ACL/current-user checks can be omitted; the conversion is pending under P4.
- `internal/acceptance/tf_provider_test.go` needs an additional noop provider
  for deferred configuration. It is already hermetic and remains pending.
- `internal/acceptance/dual_resource_api_validation_test.go` contains additional
  already-hermetic tests without acceptance prefixes; it is outside these counts.
- Import and arbitrary SDK/Go hooks need explicit engine support or a faithful
  equivalent before those scenarios can be counted as copied.
- Account provisioning copies credentials, AWS/GCP networks and customer-managed
  keys, storage configurations, the GCP VPC endpoint, and credential/workspace
  listing data sources. The skipped AWS VPC-endpoint case is now copied too.
- Catalog and schema suites retain owner/comment changes, predictive optimization,
  isolation transitions, external-location grants, and catalog replacement when
  options change. The multi-schema/grants case stays pending because of P10.
  Service-credential scenarios preserve AWS/GCP branches. Storage credentials
  include the GCP branch; the complete source function stays pending because
  its two independent AWS credentials reproduce P1. Metastore create scenarios
  cover AWS, Azure, and GCP storage-root/region formats.
- Sharing retains schema-object additions, updates, removal, reordering plans,
  recipient ownership/comment changes, and external sharing profiles. Matching
  or removing share provider configuration retains the original update-plan
  checks, even when it produces no HTTP mutation.
- MLflow cases preserve experiment path normalization, model tags, name/ID
  lookups, and resource replacement where the original uses a new random name
  on each step. Nonempty collection checks use a selected fixture-backed member.
- The identical volume-list scenarios in `catalog/data_volumes_acc_test.go` and
  `internal/providers/pluginfw/products/volume/data_volumes_acc_test.go` share
  one replay file. Its selected volume ID proves the original minimum-one-volume
  check; the original output expression is retained. Both source functions count
  as copied; cloud permutations count only once.
- Metastore renaming, changed notebook/share diagnostics, and share workspace
  replacement remain pending. See P7-P9 in the follow-up list; their assertions
  have not been weakened to make the copies pass.

- Jobs retain multi-task dependencies, shared clusters, disabled-task and periodic
  trigger values, run-as changes, and job/job-list lookups. Provider configuration
  cases retain no-op plans and the rejected cross-workspace replacement.
  Notebook import content uses `any_string` only where the original HCL embeds
  Terraform's temporary working directory; other request fields remain exact.
- Pipelines retain cluster settings, AWS cluster updates, run-as transitions and
  rejected removal, plus the nonempty data-source result check. Repeated whole
  calls and configurations use YAML anchors; calls remain consumable fixtures.
  The zero-worker case is ready for conversion under P4 with an exact request
  requiring explicit `num_workers: 0`; last-modified replay
  reproduces input-state drift (P19).
- Apps retain no-compute state, Git input-only fields and reapply, external
  deletion/recreation, invalid resource configuration, and app/list lookups.
  The known-empty-list regression uses a scalar output to distinguish null
  from empty; the full import lifecycle remains pending under P5.
  Out-of-band deletion is represented by a missing-resource refresh response.
- SQL suites cover current queries/alerts, legacy queries/alerts and visualization,
  warehouse creation/replacement and lookup, global configuration, and table/view
  column-comment updates. Comment regressions retain in-place-update/no-op plans
  and state checks. Multi-property/option SQL cases remain pending; see P11.
- SDK checks of disabled tasks, periodic triggers, and warehouse serverless mode
  become provider-read state checks of the same fields, backed by exact outgoing
  requests. Fixture responses do not independently validate server behavior.
- Additional suites copy account app integrations, budgets, workspace permission
  assignments, user roles, inline DBFS files, model-serving updates, and existing
  cluster/notification-destination lookups. Network-connectivity lookup covers AWS
  and Azure; the original function counts once. The AWS private-access-settings
  scenario was unconditionally skipped in the original and is now exercised by
  replay; its separate GCP scenario also has a replay copy.
- Table and schema lookups retain their resource-existence and ID checks. Schema
  IDs are checked on both the created resources and the listing data source.
- External-volume cases retain their storage-credential and location setup,
  owner changes, comments, and renames. Account storage-credential ownership
  uses the created service principal's application ID.
- Registered models retain grant removal, comment updates, replacement, and
  model/version data-source lookups. MLflow and vector-search endpoints retain
  their associated permissions resources.
- Share lifecycle cases retain table-object updates, shared aliases, replacement
  ID checks, and the empty-plan convergence regression. External object additions
  appear in the next refresh and must be removed by apply; external comments
  must survive refresh with an empty plan. The SDK comment read is represented
  by provider-read state checks, not an independent server-side assertion.
- Embedding settings retain approved-domain updates and dependent policy setup.
  Account cases add serverless workspace creation, GCP private-access settings,
  and metastore data access.
- The foreach-job and webhook-removal scenarios were skipped in their live
  originals. Replay exercises their existing configurations without requiring
  those live environment capabilities. Original skip statements are unchanged.
- Unity Catalog routing retains update/no-op plans for catalogs, schemas, and
  storage credentials, and the replacement plan for metastore data access.
  Unified-host suites retain job/repo creation and their name/path checks,
  using dummy account or unified metadata and explicit routing headers.
- Share suites retain create/recipient grants, four-step owner/comment updates,
  list population, and dedicated data-source provider configuration. Diagnostic
  patterns tolerate Terraform line wrapping without dropping diagnostic words.
  Legacy and Plugin Framework source scenarios run through the production mux;
  their filename suffixes distinguish source cases, not forced implementations.
- Quality monitors retain inference, time-series and snapshot configurations,
  schedule changes and granularity updates. Dashboard cases retain permission
  setup, JSON updates, unchanged timestamps, remote edits and path normalization.
  Fixed response timestamps let state checks retain the original equality and
  inequality checks without comparing dynamically generated timestamps.
- Account network cases add AWS/Azure private-endpoint replacement, AWS VPC
  endpoint creation, and GCP workspace creation with BYOVPC/private connectivity.
  The two GCP provisioning-status scenarios remain pending under P14.
- Grant and grants suites retain UC secret READ_SECRET and credential replacement
  followed by an invalid-privilege error. Model serving retains its two-model
  traffic split, model removal, tag addition/change/removal, and permissions.

## Additional lifecycles and concrete-value checks

The initial part of this batch copies 16 source functions across 13 suites:

- Provider configuration: quality-monitor workspace mismatch, rejected before
  any prerequisite is applied.
- Workspace routing: workspace-binding replacement when the provider's workspace
  changes; cleanup still uses the last applied workspace.
- Force import: create conflicts with a server-side user, then `force` adopts it.
  The out-of-band SDK creation is represented by the conflict/list responses.
- Users: account/workspace extra-attribute lookups retain the `active` check and
  check the selected ID in place of a generic count. Account deletion retains
  the home-cleanup flag and verifies the absence of workspace cleanup requests.
- Dashboard and SQL table lists: fixed member IDs replace generic nonempty/count
  checks. The SQL table IDs are checked on both resources and the data source;
  the original table-count output is retained.
- Git credentials and the default secret ACL: provider-read state checks retain
  the username, principal, and permission checked by SDK reads. Collection counts
  are replaced with these concrete values; they are not exact-cardinality checks.
- Workspace configuration: state checks retain enable/disable values and the
  rejected-key diagnostic. Reads before and after updates have separate responses.
- Notification destinations and Azure Key Vault secret scopes: retain the
  original create/read/destroy configurations and destination ID/type checks.
- External SQL tables: retain storage credential/location setup, columns, owner,
  and storage location. SQL statements are matched exactly.
- AWS cross-account policy: retain the default-policy data source and decode its
  JSON in an added output. A plan-output assertion checks the policy version in
  place of the generic nonempty-string check, without pinning the full policy.
- AI Gateway grants: retain prefixed service-name handling and convergence.
  Response fixtures omit the input-only secret value, and cleanup removes the
  grant before deleting its service and schema.

All API responses are inferred inputs, not recordings or independent evidence of
server behavior. Original acceptance tests remain unchanged. Collection checks
use concrete members where practical, as agreed for replay; backend-only checks,
imports, provider-factory switches, local files, and known races remain pending.

The account log-delivery attempt remains pending: concurrent credential/storage
reads reproduce the SDK configuration race in P1. Both AWS/Azure attempts of the
NCC account-ID drift regression are blocked by the current schema (P15). Notebook
configuration attempts still require their local source files (P3).

## Additional supported cases

The YAML conversion batch adds **13 source functions** after its initial 16.
The stacked client-helper change adds **11 helper-level source functions**,
bringing the combined batch to **40** and the cumulative inventory to
**420 copied / 125 pending**. The new YAML cases
include an additional Azure rule-set permutation, counted once by source.

- Account group membership checks use a read-only group data source and scalar
  outputs for exact child/member counts and member-type IDs. A no-op plan step
  checks those outputs once Terraform has computed them. Explicit account API
  selection remains in the generic API-selection suite.
- Account rule sets retain manager-role configuration and exact rule counts;
  service principals cover AWS and Azure application-ID behavior.
- The account user list preserves both created IDs and its original filter.
  The workspace equivalent reproduces P1; the two with-groups variants have a
  contradictory equality filter and remain pending under P18.
- Notebook/dashboard missing-parent cases retain the first failed create,
  mkdirs and retry, then check parent existence through the directory data source.
- Home-preservation cases remove the identity in an explicit step, then read the
  prepared home directory through a data source. They retain automatic final
  cleanup; SDK environment preparation is represented by inferred responses.
- The app empty-scope regression keeps its two steps. An output maps null to -1
  and an empty list to 0, preserving the distinction without a new YAML matcher.
- Table ACLs retain the SQL-permissions configuration against a prepared table
  and cluster, including exact grant/revoke commands. The live source is skipped;
  its replay copy runs offline.
- Six cloud-detection cases and five workspace-configuration helper cases were
  SDK/helper calls without Terraform steps. Go tests now replay their HTTP calls
  with the same strict handler and retain their Go assertions. They use their own
  local client helper, not an invented Terraform lifecycle or new runner hooks.

All fixtures remain manually inferred and clearly marked. The client-helper
change adds test infrastructure for direct SDK/helper replay. It does not change
the Terraform runner or production code, add race workarounds, or make live calls.

## Concurrent resource lifecycles

This batch copies **22 original functions** into eight existing or new resource
suites, bringing the inventory to **442 copied / 103 pending**. The original
acceptance tests remain unchanged. Provider race fixes are reviewed separately;
known race-detector failures do not require a different fixture format.

- Users retain the three-user entitlement checks, reapply, and workspace listing
  of two independently created users.
- Groups retain mixed user/service-principal/child-group checks; group membership
  retains the two-member count. Their existing account counterparts remain.
- Storage credentials complete the AWS IAM-role and Cloudflare branch, alongside
  the existing GCP lifecycle. Log delivery retains credential/storage setup and
  both audit and billable-usage configurations.
- Job run-as retains its user, notebook and cluster lookups. OBO tokens retain
  service-principal creation and admin-group membership before token creation.
- Permissions cover cluster policies, instance pools, directories by ID/path,
  password authorization, legacy/current SQL queries and alerts, dashboards,
  experiments, registered models, serving endpoints and apps. They retain
  principal changes and rejected removal of management permissions.

All HTTP responses are inferred and marked. The engine is unchanged: independent
requests can interleave, and identical signatures consume responses FIFO. No
artificial resource dependencies, forced ordering or race-detector skips were
added. Remaining file, import and SDK-check limitations stay in the pending list.

The same-process repeat check also reproduces P21 in the workspace membership,
mixed-member data-source and OBO-token scenarios. Their first runs pass,
including with the separate cache-locking fix under `-race`. The YAML does not
reset provider globals, relax fixture consumption or use dynamic IDs to hide it.

## Local-file and upload lifecycles

This batch copies **29 source functions**, bringing the inventory to
**479 copied / 66 pending**. Original acceptance tests remain unchanged.

- Notebooks retain path renames and DBC/IPYNB content updates. Dedicated
  provider-configuration cases retain matching, removed and rejected workspace
  IDs, no-op plans, and the notebook data-source workspace-ID state check.
- Workspace files cover source/base64 uploads, empty and ZIP files, renaming,
  and automatic parent creation after a failed upload. The parent check uses a
  dependent directory data source. DBFS retains source-file reading and its
  upload/close/read sequence.
- Volume files retain source/base64 lifecycles and changed/unchanged inputs.
  Fixed metadata timestamps preserve the original equality/inequality checks
  through provider-read state. Exact upload requests also reject unintended or
  missing updates. External file edits appear in the next refresh's metadata;
  unchanged HCL must cause the configured bytes to be uploaded again.
- File-backed dashboards retain content updates, explicit false/true publish
  flags, remote edits, and parent-path replacement. Exact create/update/publish
  requests validate provider behavior; MD5 checks validate local file hashing.
  SDK assertions that the backend persisted dashboard JSON or publish settings
  are omitted.
- Notebook and workspace-file permissions retain both path/ID variants, all
  principal transitions, and rejected management-permission removal. Their
  independent resources retain ordinary Terraform concurrency.

Inputs are deliberately small: comment-only Python, an empty file, a tiny ZIP,
minimal DBC/IPYNB notebooks, and one-page dashboard JSON. Updated inputs still
change the uploaded bytes.
YAML anchors repeat consumed fixtures; they do not make calls optional. All API
responses are inferred and marked; no live server or recorder was used.

All 29 cases previously grouped under file/payload blockers are copied. The
volume-file SDK mutation is modeled by the next refresh response; its artificial
check-hook refresh error and SDK download assertion are omitted. The replay
verifies the provider reconciles the edit by uploading the configured bytes.

## Import lifecycles

All five P5 source functions now have replay copies. The inventory is
**484 copied / 61 pending**. Every import explicitly chooses whether to persist
its state; the copied acceptance scenarios use `persist: false`, matching their
native separate-state import checks. Original integration tests are unchanged.

- Catalog force-destroy consistency retains strict native import verification
  and the two ignored audit fields.
- Schema-level connections retain catalog/schema setup, parent/full-name
  stability, a host update and composite-ID import. The inferred API response
  omits backend auth_scheme drift: identical HCL must plan no changes, while a
  changed host must plan an update.
- Apps retain all prerequisite resources, description updates and update-time
  checks, plus import by name. AWS includes model serving; GCP omits it, as in
  the source. Verification stays disabled, matching the original scenario.
- Quality monitors retain inference-table setup, import by table name and strict
  native verification with `table_name` as the matching identifier.
- Warehouse permissions retain all principal types/levels, rejected management
  removal, owner-reset requests, recreation, and the legacy `/warehouses/<ID>`
  import alias. Backend-only SDK readback is omitted. Independent resources keep
  their ordinary concurrency; P1 remains tracked separately.

All HTTP responses are inferred inputs and are marked as such. Engine examples
also exercise `persist: true`, including import without creation, subsequent
updates, preservation of existing resources/checkpoints, and automatic destroy.

## Remaining supported YAML lifecycles

This batch copies **23 original functions** as 27 YAML scenarios, including AWS,
Azure and GCP permutations. The inventory is **507 copied / 38 pending**.

- Permission suites retain principal transitions and rejected loss of management
  rights. Pipeline permission deletion is a provider no-op; its backend owner
  assertion is omitted. Other permission resources retain their reset requests.
- Settings retain configured values, computed etags and etag-bearing deletion.
  Identity cleanup verifies recursive home deletion; secrets retain group ACLs
  and recreation after an external scope deletion appears during refresh.
- Continuous jobs retain all four steps and the provider's run/cancel requests.
  Cluster-policy defaults use a convergent API response, retaining the policy
  inputs without requiring permanent backend drift.
- Pipelines retain explicit zero workers. The optimistic-concurrency timestamp
  case remains pending under P19 because it requires a separate provider fix.
- Schema grants remain independent resources. Metastore owner/name scenarios
  retain their changing storage roots and replacement plans; they do not prove
  an in-place metastore rename. The separate rename/API concern stays tracked.
- User-list filters and the NCC-list resource typo are corrected to match their
  original intent. Concrete user/group/name checks replace generic nonempty tests.
- NCC account-ID stability starts with a real Terraform checkpoint containing
  the server-populated account ID, then applies HCL without configuring that
  computed-only field and requires an empty plan.

Responses are marked placeholders, inferred from API models rather than captured
from a live service. No integration test is changed.

## Provider migration and native workspace-token replays

This batch copies **13 original functions** in 16 replay scenarios: 13 YAML and
three native Go, including both NCC cloud forms. The inventory is
**520 copied / 25 pending**.

- Notebook removal cases use ordinary expected-error steps: planning refreshes
  the old workspace successfully, then rejects the missing workspace fallback
  before apply. Workspace and account SCIM /Me fixtures distinguish requests by
  their workspace-routing header.
- Endpoint rules explicitly select Plugin Framework for opt-in/import and switch
  SDKv2/PF in both directions. Settled connection state, strict import verification,
  and no-op migration plans remain checked.
- Quality monitors switch implementations while updating inference granularities.
  Shares retain both schema objects/comments, reverse adoption and volume sharing,
  including the original independent prerequisite resources.
- Workspace/token cases retain AWS credentials, keys, storage and network setup,
  token replacement and foo/bar list checks. Their explicit Go transport maps one
  fictional workspace URL to the local fixture server; all other destinations
  still pass through the runner's rejecting transport.
- The authentication-switch case retains all five steps: create, tolerate an
  inaccessible token, remove it, reject recreation and restore the original
  identity. Fixed local credentials and exact Authorization-header fixtures
  distinguish identities. OAuth exchange and backend propagation polling are
  omitted; these are not provider behavior.

`provider_implementations` selects SDKv2 or Plugin Framework per resource type
and step. The migration suite uses ordinary YAML discovery; cleanup retains the
selection associated with its configuration. `RunNative` remains available for
Go checks and custom transports/authentication, with `ProviderFactories` sharing
the runner's mux construction. Fixtures remain stateless and strict; responses
are marked placeholders. Original integration tests remain unchanged.
