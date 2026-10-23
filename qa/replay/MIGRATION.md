# Acceptance-to-replay inventory

Initial AST inventory: **545 acceptance entry points in 176 files**. Detection
uses acceptance test prefixes or direct calls to `acceptance.*Level`. This is a
migration baseline, not an exhaustive list of every indirect wrapper or hermetic
Terraform test. Some inventoried cases already use local mocks.

**43 copied; 502 pending.** Counts refer to original Go test functions, not
individual steps or cloud permutations. A copy must preserve meaningful
assertions; unsupported assertions remain pending rather than being omitted.
Original files are unchanged.

Suite assignment requires reading each scenario. In particular, do not infer
the subject from the first HCL resource: it may be a prerequisite.

| Original source file | Copied | Pending |
| --- | ---: | ---: |
| [access/ip_access_list_test.go](../../access/ip_access_list_test.go) | 0 | 1 |
| [access/sql_permissions_provider_config_test.go](../../access/sql_permissions_provider_config_test.go) | 2 | 0 |
| [apps/custom_app_integration_test.go](../../apps/custom_app_integration_test.go) | 0 | 2 |
| [aws/data_aws_crossaccount_policy_acc_test.go](../../aws/data_aws_crossaccount_policy_acc_test.go) | 0 | 1 |
| [aws/data_instance_profiles_acc_test.go](../../aws/data_instance_profiles_acc_test.go) | 0 | 1 |
| [aws/instance_profile_provider_config_test.go](../../aws/instance_profile_provider_config_test.go) | 1 | 0 |
| [aws/instance_profile_test.go](../../aws/instance_profile_test.go) | 0 | 1 |
| [aws/service_principal_role_provider_config_test.go](../../aws/service_principal_role_provider_config_test.go) | 1 | 0 |
| [aws/user_role_provider_config_test.go](../../aws/user_role_provider_config_test.go) | 1 | 0 |
| [aws/user_role_test.go](../../aws/user_role_test.go) | 0 | 2 |
| [catalog/artifact_allowlist_test.go](../../catalog/artifact_allowlist_test.go) | 0 | 1 |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | 0 | 10 |
| [catalog/connection_test.go](../../catalog/connection_test.go) | 0 | 3 |
| [catalog/credential_provider_config_test.go](../../catalog/credential_provider_config_test.go) | 2 | 0 |
| [catalog/credential_test.go](../../catalog/credential_test.go) | 0 | 2 |
| [catalog/data_catalog_acc_test.go](../../catalog/data_catalog_acc_test.go) | 0 | 1 |
| [catalog/data_current_metastore_acc_test.go](../../catalog/data_current_metastore_acc_test.go) | 0 | 1 |
| [catalog/data_external_locations_acc_test.go](../../catalog/data_external_locations_acc_test.go) | 0 | 1 |
| [catalog/data_metastore_acc_test.go](../../catalog/data_metastore_acc_test.go) | 0 | 1 |
| [catalog/data_metastores_acc_test.go](../../catalog/data_metastores_acc_test.go) | 0 | 1 |
| [catalog/data_schema_acc_test.go](../../catalog/data_schema_acc_test.go) | 0 | 1 |
| [catalog/data_schemas_acc_test.go](../../catalog/data_schemas_acc_test.go) | 0 | 1 |
| [catalog/data_storage_credential_acc_test.go](../../catalog/data_storage_credential_acc_test.go) | 0 | 1 |
| [catalog/data_storage_credentials_acc_test.go](../../catalog/data_storage_credentials_acc_test.go) | 0 | 1 |
| [catalog/data_table_acc_test.go](../../catalog/data_table_acc_test.go) | 0 | 1 |
| [catalog/data_tables_acc_test.go](../../catalog/data_tables_acc_test.go) | 0 | 1 |
| [catalog/data_volume_acc_test.go](../../catalog/data_volume_acc_test.go) | 0 | 1 |
| [catalog/data_volumes_acc_test.go](../../catalog/data_volumes_acc_test.go) | 0 | 1 |
| [catalog/external_location_test.go](../../catalog/external_location_test.go) | 0 | 3 |
| [catalog/grant_test.go](../../catalog/grant_test.go) | 0 | 5 |
| [catalog/grants_test.go](../../catalog/grants_test.go) | 0 | 6 |
| [catalog/metastore_assignment_provider_config_test.go](../../catalog/metastore_assignment_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_assignment_test.go](../../catalog/metastore_assignment_test.go) | 0 | 2 |
| [catalog/metastore_data_access_provider_config_test.go](../../catalog/metastore_data_access_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_data_access_test.go](../../catalog/metastore_data_access_test.go) | 0 | 2 |
| [catalog/metastore_provider_config_test.go](../../catalog/metastore_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_test.go](../../catalog/metastore_test.go) | 0 | 4 |
| [catalog/quality_monitor_test.go](../../catalog/quality_monitor_test.go) | 0 | 2 |
| [catalog/registered_model_test.go](../../catalog/registered_model_test.go) | 0 | 1 |
| [catalog/resource_sql_table_acc_test.go](../../catalog/resource_sql_table_acc_test.go) | 0 | 2 |
| [catalog/schema_test.go](../../catalog/schema_test.go) | 0 | 2 |
| [catalog/storage_credential_provider_config_test.go](../../catalog/storage_credential_provider_config_test.go) | 2 | 0 |
| [catalog/storage_credential_test.go](../../catalog/storage_credential_test.go) | 0 | 2 |
| [catalog/system_schema_test.go](../../catalog/system_schema_test.go) | 0 | 1 |
| [catalog/table_provider_config_test.go](../../catalog/table_provider_config_test.go) | 1 | 0 |
| [catalog/volume_test.go](../../catalog/volume_test.go) | 0 | 2 |
| [catalog/workspace_binding_test.go](../../catalog/workspace_binding_test.go) | 0 | 1 |
| [clusters/cluster_policy_test.go](../../clusters/cluster_policy_test.go) | 0 | 3 |
| [clusters/data_cluster_acc_test.go](../../clusters/data_cluster_acc_test.go) | 0 | 1 |
| [clusters/data_clusters_acc_test.go](../../clusters/data_clusters_acc_test.go) | 0 | 3 |
| [clusters/data_spark_version_acc_test.go](../../clusters/data_spark_version_acc_test.go) | 0 | 1 |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | 0 | 6 |
| [dashboards/dashboard_test.go](../../dashboards/dashboard_test.go) | 0 | 8 |
| [finops/budget_test.go](../../finops/budget_test.go) | 0 | 2 |
| [internal/acceptance/aibi_embedding_settings_test.go](../../internal/acceptance/aibi_embedding_settings_test.go) | 0 | 1 |
| [internal/acceptance/tf_provider_test.go](../../internal/acceptance/tf_provider_test.go) | 0 | 1 |
| [internal/acceptance/uc_resources_workspace_id_acc_test.go](../../internal/acceptance/uc_resources_workspace_id_acc_test.go) | 0 | 6 |
| [internal/acceptance/unified_host_acc_test.go](../../internal/acceptance/unified_host_acc_test.go) | 0 | 6 |
| [internal/acceptance/workspace_id_pf_app_acc_test.go](../../internal/acceptance/workspace_id_pf_app_acc_test.go) | 0 | 21 |
| [internal/acceptance/workspace_id_pf_tag_policy_acc_test.go](../../internal/acceptance/workspace_id_pf_tag_policy_acc_test.go) | 0 | 21 |
| [internal/acceptance/workspace_id_sdkv2_gosdk_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_gosdk_acc_test.go) | 0 | 21 |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | 0 | 21 |
| [internal/providers/pluginfw/products/app/data_app_acc_test.go](../../internal/providers/pluginfw/products/app/data_app_acc_test.go) | 0 | 1 |
| [internal/providers/pluginfw/products/app/data_apps_acc_test.go](../../internal/providers/pluginfw/products/app/data_apps_acc_test.go) | 0 | 1 |
| [internal/providers/pluginfw/products/app/resource_app_acc_test.go](../../internal/providers/pluginfw/products/app/resource_app_acc_test.go) | 0 | 8 |
| [internal/providers/pluginfw/products/cluster/data_cluster_acc_test.go](../../internal/providers/pluginfw/products/cluster/data_cluster_acc_test.go) | 0 | 2 |
| [internal/providers/pluginfw/products/dashboards/data_dashboards_acc_test.go](../../internal/providers/pluginfw/products/dashboards/data_dashboards_acc_test.go) | 0 | 1 |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | 0 | 3 |
| [internal/providers/pluginfw/products/notificationdestinations/data_notification_destinations_acc_test.go](../../internal/providers/pluginfw/products/notificationdestinations/data_notification_destinations_acc_test.go) | 0 | 2 |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | 0 | 6 |
| [internal/providers/pluginfw/products/sharing/data_recipients_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_recipients_acc_test.go) | 0 | 1 |
| [internal/providers/pluginfw/products/sharing/data_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_share_acc_test.go) | 0 | 2 |
| [internal/providers/pluginfw/products/sharing/data_shares_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_shares_acc_test.go) | 0 | 2 |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | 0 | 25 |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | 0 | 8 |
| [internal/providers/pluginfw/products/volume/data_volumes_acc_test.go](../../internal/providers/pluginfw/products/volume/data_volumes_acc_test.go) | 0 | 1 |
| [jobs/data_job_acc_test.go](../../jobs/data_job_acc_test.go) | 0 | 5 |
| [jobs/data_jobs_acc_test.go](../../jobs/data_jobs_acc_test.go) | 0 | 4 |
| [jobs/job_test.go](../../jobs/job_test.go) | 0 | 16 |
| [mlflow/data_mlflow_experiment_acc_test.go](../../mlflow/data_mlflow_experiment_acc_test.go) | 0 | 1 |
| [mlflow/data_mlflow_model_acc_test.go](../../mlflow/data_mlflow_model_acc_test.go) | 0 | 1 |
| [mlflow/data_mlflow_models_acc_test.go](../../mlflow/data_mlflow_models_acc_test.go) | 0 | 1 |
| [mlflow/mlflow_experiment_test.go](../../mlflow/mlflow_experiment_test.go) | 0 | 2 |
| [mlflow/mlflow_model_test.go](../../mlflow/mlflow_model_test.go) | 0 | 1 |
| [mws/data_current_config_acc_test.go](../../mws/data_current_config_acc_test.go) | 0 | 8 |
| [mws/data_mws_credentials_acc_test.go](../../mws/data_mws_credentials_acc_test.go) | 0 | 1 |
| [mws/data_mws_network_connectivity_config_acc_test.go](../../mws/data_mws_network_connectivity_config_acc_test.go) | 0 | 1 |
| [mws/data_mws_network_connectivity_configs_acc_test.go](../../mws/data_mws_network_connectivity_configs_acc_test.go) | 0 | 1 |
| [mws/data_mws_workspaces_acc_test.go](../../mws/data_mws_workspaces_acc_test.go) | 0 | 1 |
| [mws/mws_credentials_test.go](../../mws/mws_credentials_test.go) | 0 | 1 |
| [mws/mws_customer_managed_keys_test.go](../../mws/mws_customer_managed_keys_test.go) | 0 | 2 |
| [mws/mws_log_delivery_test.go](../../mws/mws_log_delivery_test.go) | 0 | 1 |
| [mws/mws_network_connectivity_config_test.go](../../mws/mws_network_connectivity_config_test.go) | 0 | 1 |
| [mws/mws_networks_test.go](../../mws/mws_networks_test.go) | 0 | 2 |
| [mws/mws_permissionassignments_test.go](../../mws/mws_permissionassignments_test.go) | 0 | 2 |
| [mws/mws_private_access_settings_test.go](../../mws/mws_private_access_settings_test.go) | 0 | 2 |
| [mws/mws_storage_configurations_test.go](../../mws/mws_storage_configurations_test.go) | 0 | 1 |
| [mws/mws_vpc_endpoint_test.go](../../mws/mws_vpc_endpoint_test.go) | 0 | 2 |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | 0 | 9 |
| [mws/resource_mws_ncc_private_endpoint_rule_acc_test.go](../../mws/resource_mws_ncc_private_endpoint_rule_acc_test.go) | 0 | 1 |
| [permissions/access_control_rule_set_provider_config_test.go](../../permissions/access_control_rule_set_provider_config_test.go) | 1 | 0 |
| [permissions/account_rule_set_test.go](../../permissions/account_rule_set_test.go) | 0 | 2 |
| [permissions/permissions_provider_config_test.go](../../permissions/permissions_provider_config_test.go) | 1 | 0 |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | 0 | 27 |
| [permissions/resource_permissions_test.go](../../permissions/resource_permissions_test.go) | 0 | 1 |
| [pipelines/data_pipelines_acc_test.go](../../pipelines/data_pipelines_acc_test.go) | 0 | 1 |
| [pipelines/pipeline_provider_config_test.go](../../pipelines/pipeline_provider_config_test.go) | 1 | 0 |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | 0 | 7 |
| [policies/cluster_policy_provider_config_test.go](../../policies/cluster_policy_provider_config_test.go) | 1 | 0 |
| [pools/instance_pool_provider_config_test.go](../../pools/instance_pool_provider_config_test.go) | 1 | 0 |
| [pools/resource_instance_pool_acc_test.go](../../pools/resource_instance_pool_acc_test.go) | 0 | 1 |
| [repos/git_credential_test.go](../../repos/git_credential_test.go) | 0 | 1 |
| [repos/repo_provider_config_test.go](../../repos/repo_provider_config_test.go) | 1 | 0 |
| [scim/data_group_acc_test.go](../../scim/data_group_acc_test.go) | 0 | 2 |
| [scim/data_service_principal_acc_test.go](../../scim/data_service_principal_acc_test.go) | 0 | 5 |
| [scim/data_service_principals_acc_test.go](../../scim/data_service_principals_acc_test.go) | 0 | 5 |
| [scim/data_user_acc_test.go](../../scim/data_user_acc_test.go) | 0 | 4 |
| [scim/entitlement_provider_config_test.go](../../scim/entitlement_provider_config_test.go) | 2 | 0 |
| [scim/group_member_provider_config_test.go](../../scim/group_member_provider_config_test.go) | 1 | 0 |
| [scim/group_member_test.go](../../scim/group_member_test.go) | 0 | 3 |
| [scim/group_provider_config_test.go](../../scim/group_provider_config_test.go) | 1 | 0 |
| [scim/group_role_provider_config_test.go](../../scim/group_role_provider_config_test.go) | 1 | 0 |
| [scim/group_role_test.go](../../scim/group_role_test.go) | 0 | 2 |
| [scim/group_test.go](../../scim/group_test.go) | 0 | 6 |
| [scim/service_principal_provider_config_test.go](../../scim/service_principal_provider_config_test.go) | 1 | 0 |
| [scim/service_principal_test.go](../../scim/service_principal_test.go) | 0 | 8 |
| [scim/user_provider_config_test.go](../../scim/user_provider_config_test.go) | 1 | 0 |
| [scim/user_test.go](../../scim/user_test.go) | 0 | 8 |
| [secrets/secret_acl_test.go](../../secrets/secret_acl_test.go) | 0 | 2 |
| [secrets/secret_scope_provider_config_test.go](../../secrets/secret_scope_provider_config_test.go) | 1 | 0 |
| [secrets/secret_scope_test.go](../../secrets/secret_scope_test.go) | 0 | 2 |
| [secrets/secret_test.go](../../secrets/secret_test.go) | 0 | 1 |
| [serving/model_serving_test.go](../../serving/model_serving_test.go) | 0 | 4 |
| [settings/default_namespace_test.go](../../settings/default_namespace_test.go) | 0 | 4 |
| [settings/disable_legacy_access_setting_test.go](../../settings/disable_legacy_access_setting_test.go) | 0 | 1 |
| [settings/disable_legacy_dbfs_setting_test.go](../../settings/disable_legacy_dbfs_setting_test.go) | 0 | 1 |
| [settings/disable_legacy_features_test.go](../../settings/disable_legacy_features_test.go) | 0 | 1 |
| [settings/notification_destination_provider_config_test.go](../../settings/notification_destination_provider_config_test.go) | 1 | 0 |
| [settings/notification_destination_test.go](../../settings/notification_destination_test.go) | 0 | 5 |
| [settings/restrict_workspace_admins_test.go](../../settings/restrict_workspace_admins_test.go) | 0 | 1 |
| [sharing/data_shares_acc_test.go](../../sharing/data_shares_acc_test.go) | 0 | 1 |
| [sharing/provider_provider_config_test.go](../../sharing/provider_provider_config_test.go) | 1 | 0 |
| [sharing/provider_test.go](../../sharing/provider_test.go) | 0 | 1 |
| [sharing/recipient_test.go](../../sharing/recipient_test.go) | 0 | 3 |
| [sharing/share_test.go](../../sharing/share_test.go) | 0 | 2 |
| [sql/alert_test.go](../../sql/alert_test.go) | 0 | 1 |
| [sql/data_sql_warehouse_acc_test.go](../../sql/data_sql_warehouse_acc_test.go) | 0 | 1 |
| [sql/query_test.go](../../sql/query_test.go) | 0 | 1 |
| [sql/sql_alert_test.go](../../sql/sql_alert_test.go) | 0 | 1 |
| [sql/sql_dashboard_provider_config_test.go](../../sql/sql_dashboard_provider_config_test.go) | 1 | 0 |
| [sql/sql_endpoint_test.go](../../sql/sql_endpoint_test.go) | 0 | 1 |
| [sql/sql_global_config_provider_config_test.go](../../sql/sql_global_config_provider_config_test.go) | 1 | 0 |
| [sql/sql_global_config_test.go](../../sql/sql_global_config_test.go) | 0 | 4 |
| [sql/sql_permissions_test.go](../../sql/sql_permissions_test.go) | 0 | 1 |
| [sql/sql_query_provider_config_test.go](../../sql/sql_query_provider_config_test.go) | 1 | 0 |
| [sql/sql_query_test.go](../../sql/sql_query_test.go) | 0 | 1 |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | 0 | 18 |
| [sql/sql_visualization_provider_config_test.go](../../sql/sql_visualization_provider_config_test.go) | 1 | 0 |
| [sql/sql_widget_provider_config_test.go](../../sql/sql_widget_provider_config_test.go) | 1 | 0 |
| [storage/dbfs_file_provider_config_test.go](../../storage/dbfs_file_provider_config_test.go) | 1 | 0 |
| [storage/dbfs_file_test.go](../../storage/dbfs_file_test.go) | 0 | 2 |
| [storage/file_test.go](../../storage/file_test.go) | 0 | 7 |
| [storage/mount_provider_config_test.go](../../storage/mount_provider_config_test.go) | 1 | 0 |
| [tokens/obo_token_provider_config_test.go](../../tokens/obo_token_provider_config_test.go) | 1 | 0 |
| [tokens/obo_token_test.go](../../tokens/obo_token_test.go) | 0 | 1 |
| [tokens/service_principal_secret_provider_config_test.go](../../tokens/service_principal_secret_provider_config_test.go) | 1 | 0 |
| [tokens/token_provider_config_test.go](../../tokens/token_provider_config_test.go) | 1 | 0 |
| [tokens/token_test.go](../../tokens/token_test.go) | 0 | 1 |
| [vectorsearch/vector_search_test.go](../../vectorsearch/vector_search_test.go) | 0 | 1 |
| [workspace/data_notebook_acc_test.go](../../workspace/data_notebook_acc_test.go) | 0 | 3 |
| [workspace/directory_test.go](../../workspace/directory_test.go) | 0 | 1 |
| [workspace/global_init_script_test.go](../../workspace/global_init_script_test.go) | 0 | 1 |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | 0 | 9 |
| [workspace/workspace_conf_provider_config_test.go](../../workspace/workspace_conf_provider_config_test.go) | 1 | 0 |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | 0 | 6 |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | 0 | 5 |

## Known adaptations and gaps

- Negative provider-config plan-only cases fail before apply or any API call.
  They use `expect_error` without `plan_only` because the replay runner rejects
  that combination. The original HCL and error patterns are preserved.
- `secrets/secret_scope_test.go` includes SDK-only ACL and current-user checks.
  Copying just its resource lifecycle would lose those assertions.
- `internal/acceptance/tf_provider_test.go` needs an additional noop provider
  for deferred configuration. It is already hermetic and remains pending.
- `internal/acceptance/dual_resource_api_validation_test.go` contains additional
  already-hermetic tests without acceptance prefixes; it is outside these counts.
- Import and arbitrary SDK/Go hooks need explicit engine support or a faithful
  equivalent before those scenarios can be counted as copied.
