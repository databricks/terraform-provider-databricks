# Acceptance-to-replay inventory

Initial AST inventory: **545 acceptance entry points in 176 files**. Detection
uses acceptance test prefixes or direct calls to `acceptance.*Level`. This is a
migration baseline, not an exhaustive list of every indirect wrapper or hermetic
Terraform test. Some inventoried cases already use local mocks.

**396 copied; 149 pending.** Counts refer to original Go test functions, not
individual steps or cloud permutations. A copy must preserve meaningful
assertions; unsupported assertions remain pending rather than being omitted.
Original files are unchanged.

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
| [aws/instance_profile_test.go](../../aws/instance_profile_test.go) | 0 | 1 |
| [aws/service_principal_role_provider_config_test.go](../../aws/service_principal_role_provider_config_test.go) | 1 | 0 |
| [aws/user_role_provider_config_test.go](../../aws/user_role_provider_config_test.go) | 1 | 0 |
| [aws/user_role_test.go](../../aws/user_role_test.go) | 2 | 0 |
| [catalog/artifact_allowlist_test.go](../../catalog/artifact_allowlist_test.go) | 1 | 0 |
| [catalog/catalog_test.go](../../catalog/catalog_test.go) | 7 | 3 |
| [catalog/connection_test.go](../../catalog/connection_test.go) | 2 | 1 |
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
| [catalog/grant_test.go](../../catalog/grant_test.go) | 4 | 1 |
| [catalog/grants_test.go](../../catalog/grants_test.go) | 5 | 1 |
| [catalog/metastore_assignment_provider_config_test.go](../../catalog/metastore_assignment_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_assignment_test.go](../../catalog/metastore_assignment_test.go) | 2 | 0 |
| [catalog/metastore_data_access_provider_config_test.go](../../catalog/metastore_data_access_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_data_access_test.go](../../catalog/metastore_data_access_test.go) | 2 | 0 |
| [catalog/metastore_provider_config_test.go](../../catalog/metastore_provider_config_test.go) | 2 | 0 |
| [catalog/metastore_test.go](../../catalog/metastore_test.go) | 3 | 1 |
| [catalog/quality_monitor_test.go](../../catalog/quality_monitor_test.go) | 2 | 0 |
| [catalog/registered_model_test.go](../../catalog/registered_model_test.go) | 1 | 0 |
| [catalog/resource_sql_table_acc_test.go](../../catalog/resource_sql_table_acc_test.go) | 2 | 0 |
| [catalog/schema_test.go](../../catalog/schema_test.go) | 1 | 1 |
| [catalog/storage_credential_provider_config_test.go](../../catalog/storage_credential_provider_config_test.go) | 2 | 0 |
| [catalog/storage_credential_test.go](../../catalog/storage_credential_test.go) | 1 | 1 |
| [catalog/system_schema_test.go](../../catalog/system_schema_test.go) | 1 | 0 |
| [catalog/table_provider_config_test.go](../../catalog/table_provider_config_test.go) | 1 | 0 |
| [catalog/volume_test.go](../../catalog/volume_test.go) | 2 | 0 |
| [catalog/workspace_binding_test.go](../../catalog/workspace_binding_test.go) | 0 | 1 |
| [clusters/cluster_policy_test.go](../../clusters/cluster_policy_test.go) | 3 | 0 |
| [clusters/data_cluster_acc_test.go](../../clusters/data_cluster_acc_test.go) | 1 | 0 |
| [clusters/data_clusters_acc_test.go](../../clusters/data_clusters_acc_test.go) | 2 | 1 |
| [clusters/data_spark_version_acc_test.go](../../clusters/data_spark_version_acc_test.go) | 1 | 0 |
| [clusters/node_type_test.go](../../clusters/node_type_test.go) | 0 | 6 |
| [dashboards/dashboard_test.go](../../dashboards/dashboard_test.go) | 5 | 3 |
| [finops/budget_test.go](../../finops/budget_test.go) | 2 | 0 |
| [internal/acceptance/aibi_embedding_settings_test.go](../../internal/acceptance/aibi_embedding_settings_test.go) | 1 | 0 |
| [internal/acceptance/tf_provider_test.go](../../internal/acceptance/tf_provider_test.go) | 0 | 1 |
| [internal/acceptance/uc_resources_workspace_id_acc_test.go](../../internal/acceptance/uc_resources_workspace_id_acc_test.go) | 6 | 0 |
| [internal/acceptance/unified_host_acc_test.go](../../internal/acceptance/unified_host_acc_test.go) | 6 | 0 |
| [internal/acceptance/workspace_id_pf_app_acc_test.go](../../internal/acceptance/workspace_id_pf_app_acc_test.go) | 21 | 0 |
| [internal/acceptance/workspace_id_pf_tag_policy_acc_test.go](../../internal/acceptance/workspace_id_pf_tag_policy_acc_test.go) | 21 | 0 |
| [internal/acceptance/workspace_id_sdkv2_gosdk_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_gosdk_acc_test.go) | 21 | 0 |
| [internal/acceptance/workspace_id_sdkv2_http_acc_test.go](../../internal/acceptance/workspace_id_sdkv2_http_acc_test.go) | 19 | 2 |
| [internal/providers/pluginfw/products/app/data_app_acc_test.go](../../internal/providers/pluginfw/products/app/data_app_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/app/data_apps_acc_test.go](../../internal/providers/pluginfw/products/app/data_apps_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/app/resource_app_acc_test.go](../../internal/providers/pluginfw/products/app/resource_app_acc_test.go) | 6 | 2 |
| [internal/providers/pluginfw/products/cluster/data_cluster_acc_test.go](../../internal/providers/pluginfw/products/cluster/data_cluster_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/dashboards/data_dashboards_acc_test.go](../../internal/providers/pluginfw/products/dashboards/data_dashboards_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go](../../internal/providers/pluginfw/products/networks/privateendpointrule/resource_acc_test.go) | 0 | 3 |
| [internal/providers/pluginfw/products/notificationdestinations/data_notification_destinations_acc_test.go](../../internal/providers/pluginfw/products/notificationdestinations/data_notification_destinations_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go](../../internal/providers/pluginfw/products/qualitymonitor/resource_quality_monitor_acc_test.go) | 3 | 3 |
| [internal/providers/pluginfw/products/sharing/data_recipients_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_recipients_acc_test.go) | 1 | 0 |
| [internal/providers/pluginfw/products/sharing/data_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_share_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/sharing/data_shares_acc_test.go](../../internal/providers/pluginfw/products/sharing/data_shares_acc_test.go) | 2 | 0 |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | 20 | 5 |
| [internal/providers/pluginfw/products/user/data_users_acc_test.go](../../internal/providers/pluginfw/products/user/data_users_acc_test.go) | 4 | 4 |
| [internal/providers/pluginfw/products/volume/data_volumes_acc_test.go](../../internal/providers/pluginfw/products/volume/data_volumes_acc_test.go) | 1 | 0 |
| [jobs/data_job_acc_test.go](../../jobs/data_job_acc_test.go) | 5 | 0 |
| [jobs/data_jobs_acc_test.go](../../jobs/data_jobs_acc_test.go) | 4 | 0 |
| [jobs/job_test.go](../../jobs/job_test.go) | 13 | 3 |
| [mlflow/data_mlflow_experiment_acc_test.go](../../mlflow/data_mlflow_experiment_acc_test.go) | 1 | 0 |
| [mlflow/data_mlflow_model_acc_test.go](../../mlflow/data_mlflow_model_acc_test.go) | 1 | 0 |
| [mlflow/data_mlflow_models_acc_test.go](../../mlflow/data_mlflow_models_acc_test.go) | 1 | 0 |
| [mlflow/mlflow_experiment_test.go](../../mlflow/mlflow_experiment_test.go) | 2 | 0 |
| [mlflow/mlflow_model_test.go](../../mlflow/mlflow_model_test.go) | 1 | 0 |
| [mws/data_current_config_acc_test.go](../../mws/data_current_config_acc_test.go) | 8 | 0 |
| [mws/data_mws_credentials_acc_test.go](../../mws/data_mws_credentials_acc_test.go) | 1 | 0 |
| [mws/data_mws_network_connectivity_config_acc_test.go](../../mws/data_mws_network_connectivity_config_acc_test.go) | 1 | 0 |
| [mws/data_mws_network_connectivity_configs_acc_test.go](../../mws/data_mws_network_connectivity_configs_acc_test.go) | 0 | 1 |
| [mws/data_mws_workspaces_acc_test.go](../../mws/data_mws_workspaces_acc_test.go) | 1 | 0 |
| [mws/mws_credentials_test.go](../../mws/mws_credentials_test.go) | 1 | 0 |
| [mws/mws_customer_managed_keys_test.go](../../mws/mws_customer_managed_keys_test.go) | 2 | 0 |
| [mws/mws_log_delivery_test.go](../../mws/mws_log_delivery_test.go) | 0 | 1 |
| [mws/mws_network_connectivity_config_test.go](../../mws/mws_network_connectivity_config_test.go) | 1 | 0 |
| [mws/mws_networks_test.go](../../mws/mws_networks_test.go) | 2 | 0 |
| [mws/mws_permissionassignments_test.go](../../mws/mws_permissionassignments_test.go) | 2 | 0 |
| [mws/mws_private_access_settings_test.go](../../mws/mws_private_access_settings_test.go) | 2 | 0 |
| [mws/mws_storage_configurations_test.go](../../mws/mws_storage_configurations_test.go) | 1 | 0 |
| [mws/mws_vpc_endpoint_test.go](../../mws/mws_vpc_endpoint_test.go) | 2 | 0 |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | 4 | 5 |
| [mws/resource_mws_ncc_private_endpoint_rule_acc_test.go](../../mws/resource_mws_ncc_private_endpoint_rule_acc_test.go) | 0 | 1 |
| [permissions/access_control_rule_set_provider_config_test.go](../../permissions/access_control_rule_set_provider_config_test.go) | 1 | 0 |
| [permissions/account_rule_set_test.go](../../permissions/account_rule_set_test.go) | 0 | 2 |
| [permissions/permissions_provider_config_test.go](../../permissions/permissions_provider_config_test.go) | 1 | 0 |
| [permissions/permissions_test.go](../../permissions/permissions_test.go) | 0 | 27 |
| [permissions/resource_permissions_test.go](../../permissions/resource_permissions_test.go) | 0 | 1 |
| [pipelines/data_pipelines_acc_test.go](../../pipelines/data_pipelines_acc_test.go) | 1 | 0 |
| [pipelines/pipeline_provider_config_test.go](../../pipelines/pipeline_provider_config_test.go) | 1 | 0 |
| [pipelines/pipeline_test.go](../../pipelines/pipeline_test.go) | 5 | 2 |
| [policies/cluster_policy_provider_config_test.go](../../policies/cluster_policy_provider_config_test.go) | 1 | 0 |
| [pools/instance_pool_provider_config_test.go](../../pools/instance_pool_provider_config_test.go) | 1 | 0 |
| [pools/resource_instance_pool_acc_test.go](../../pools/resource_instance_pool_acc_test.go) | 1 | 0 |
| [repos/git_credential_test.go](../../repos/git_credential_test.go) | 1 | 0 |
| [repos/repo_provider_config_test.go](../../repos/repo_provider_config_test.go) | 1 | 0 |
| [scim/data_group_acc_test.go](../../scim/data_group_acc_test.go) | 0 | 2 |
| [scim/data_service_principal_acc_test.go](../../scim/data_service_principal_acc_test.go) | 5 | 0 |
| [scim/data_service_principals_acc_test.go](../../scim/data_service_principals_acc_test.go) | 5 | 0 |
| [scim/data_user_acc_test.go](../../scim/data_user_acc_test.go) | 4 | 0 |
| [scim/entitlement_provider_config_test.go](../../scim/entitlement_provider_config_test.go) | 2 | 0 |
| [scim/group_member_provider_config_test.go](../../scim/group_member_provider_config_test.go) | 1 | 0 |
| [scim/group_member_test.go](../../scim/group_member_test.go) | 0 | 3 |
| [scim/group_provider_config_test.go](../../scim/group_provider_config_test.go) | 1 | 0 |
| [scim/group_role_provider_config_test.go](../../scim/group_role_provider_config_test.go) | 1 | 0 |
| [scim/group_role_test.go](../../scim/group_role_test.go) | 2 | 0 |
| [scim/group_test.go](../../scim/group_test.go) | 6 | 0 |
| [scim/service_principal_provider_config_test.go](../../scim/service_principal_provider_config_test.go) | 1 | 0 |
| [scim/service_principal_test.go](../../scim/service_principal_test.go) | 6 | 2 |
| [scim/user_provider_config_test.go](../../scim/user_provider_config_test.go) | 1 | 0 |
| [scim/user_test.go](../../scim/user_test.go) | 5 | 3 |
| [secrets/secret_acl_test.go](../../secrets/secret_acl_test.go) | 1 | 1 |
| [secrets/secret_scope_provider_config_test.go](../../secrets/secret_scope_provider_config_test.go) | 1 | 0 |
| [secrets/secret_scope_test.go](../../secrets/secret_scope_test.go) | 1 | 1 |
| [secrets/secret_test.go](../../secrets/secret_test.go) | 1 | 0 |
| [serving/model_serving_test.go](../../serving/model_serving_test.go) | 4 | 0 |
| [settings/default_namespace_test.go](../../settings/default_namespace_test.go) | 3 | 1 |
| [settings/disable_legacy_access_setting_test.go](../../settings/disable_legacy_access_setting_test.go) | 0 | 1 |
| [settings/disable_legacy_dbfs_setting_test.go](../../settings/disable_legacy_dbfs_setting_test.go) | 0 | 1 |
| [settings/disable_legacy_features_test.go](../../settings/disable_legacy_features_test.go) | 0 | 1 |
| [settings/notification_destination_provider_config_test.go](../../settings/notification_destination_provider_config_test.go) | 1 | 0 |
| [settings/notification_destination_test.go](../../settings/notification_destination_test.go) | 5 | 0 |
| [settings/restrict_workspace_admins_test.go](../../settings/restrict_workspace_admins_test.go) | 0 | 1 |
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
| [sql/sql_permissions_test.go](../../sql/sql_permissions_test.go) | 0 | 1 |
| [sql/sql_query_provider_config_test.go](../../sql/sql_query_provider_config_test.go) | 1 | 0 |
| [sql/sql_query_test.go](../../sql/sql_query_test.go) | 1 | 0 |
| [sql/sql_table_test.go](../../sql/sql_table_test.go) | 2 | 16 |
| [sql/sql_visualization_provider_config_test.go](../../sql/sql_visualization_provider_config_test.go) | 1 | 0 |
| [sql/sql_widget_provider_config_test.go](../../sql/sql_widget_provider_config_test.go) | 1 | 0 |
| [storage/dbfs_file_provider_config_test.go](../../storage/dbfs_file_provider_config_test.go) | 1 | 0 |
| [storage/dbfs_file_test.go](../../storage/dbfs_file_test.go) | 1 | 1 |
| [storage/file_test.go](../../storage/file_test.go) | 0 | 7 |
| [storage/mount_provider_config_test.go](../../storage/mount_provider_config_test.go) | 1 | 0 |
| [tokens/obo_token_provider_config_test.go](../../tokens/obo_token_provider_config_test.go) | 1 | 0 |
| [tokens/obo_token_test.go](../../tokens/obo_token_test.go) | 0 | 1 |
| [tokens/service_principal_secret_provider_config_test.go](../../tokens/service_principal_secret_provider_config_test.go) | 1 | 0 |
| [tokens/token_provider_config_test.go](../../tokens/token_provider_config_test.go) | 1 | 0 |
| [tokens/token_test.go](../../tokens/token_test.go) | 1 | 0 |
| [vectorsearch/vector_search_test.go](../../vectorsearch/vector_search_test.go) | 1 | 0 |
| [workspace/data_notebook_acc_test.go](../../workspace/data_notebook_acc_test.go) | 0 | 3 |
| [workspace/directory_test.go](../../workspace/directory_test.go) | 1 | 0 |
| [workspace/global_init_script_test.go](../../workspace/global_init_script_test.go) | 1 | 0 |
| [workspace/notebook_test.go](../../workspace/notebook_test.go) | 1 | 8 |
| [workspace/workspace_conf_provider_config_test.go](../../workspace/workspace_conf_provider_config_test.go) | 1 | 0 |
| [workspace/workspace_conf_test.go](../../workspace/workspace_conf_test.go) | 1 | 5 |
| [workspace/workspace_file_test.go](../../workspace/workspace_file_test.go) | 0 | 5 |

## Known adaptations and gaps

See [FOLLOW_UPS.md](FOLLOW_UPS.md) for deferred problems and the scenarios they
block. Keep new problems in that list as more source tests are reviewed.

- Workspace routing copies directory, app, tag-policy, and notebook cases into
  `features/workspace_routing/`, split by carrier. Preserve workspace-ID state
  checks, no-op/replacement plan actions, and diagnostic patterns. Routing headers
  distinguish requests for the old and new workspaces. Step-local dummy metadata
  models account/workspace clients and host transitions on the local server.
  YAML aliases repeat consumable fixtures; they do not make calls reusable.
  The provider-upgrade cases retain the originals' current-provider approximation.
  Removed defaults/overrides retain diagnostic checks. The two notebook removal
  cases remain pending because the provider attempts creation rather than
  returning the expected planning diagnostic.
- Current-config cloud branches use separate YAML files for AWS, Azure, and
  GCP but still count as one source function. Explicit API cases extend
  `features/api_selection`; inference and cloud overrides stay in the
  `current_config` resource suite.
- Compute cases copy cluster-policy rename/family overrides, cluster lookup
  and name filtering, Spark-version selection, instance-profile listing, and
  global-init-script creation. With fixed Spark-version responses, the selected
  ID assertion includes the original Scala 2.13 suffix requirement.
  Built-in policy cleanup resets its override instead of deleting the policy.
- `TestAccDataSourceClusters_FilterBy` also checks SDK-only cluster fields;
  copying only its Terraform data-source reads would lose those assertions.
- Workspace-file uploads use multipart requests, while the current fixture
  handler matches JSON bodies. Source-file scenarios also need a stable way
  to provide local files to Terraform. Those cases remain pending.
- Identity cases add user-name normalization, service-principal lifecycles,
  user/service-principal data sources, and group roles. Explicit `api` cases
  belong to `features/api_selection`; other cases stay with their resource.
  User data-source/resource ID equality is checked against the same fixture ID.
- Azure service-principal templates generate a new application ID for each
  step, so their second steps replace the resource. AWS display-name changes
  update it in place. Account destroy preserves the default disable operation.
- The three independent users in `TestAccUserResource` expose an existing race
  in `common/unified_provider.go`'s shared client cache under `-race`. That case
  remains pending; no artificial resource dependencies or production fixes are
  added to these ports. Home-directory cleanup and arbitrary SDK checks also
  remain pending.
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
- `secrets/secret_scope_test.go` includes SDK-only ACL and current-user checks.
  Copying just its resource lifecycle would lose those assertions.
- `internal/acceptance/tf_provider_test.go` needs an additional noop provider
  for deferred configuration. It is already hermetic and remains pending.
- `internal/acceptance/dual_resource_api_validation_test.go` contains additional
  already-hermetic tests without acceptance prefixes; it is outside these counts.
- Import and arbitrary SDK/Go hooks need explicit engine support or a faithful
  equivalent before those scenarios can be counted as copied.
- Account provisioning copies credentials, AWS/GCP networks and customer-managed
  keys, storage configurations, the GCP VPC endpoint, and credential/workspace
  listing data sources. The skipped AWS VPC-endpoint case stays pending.
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
  Zero-worker and last-modified checks remain pending under P4.
- Apps retain no-compute state, Git input-only fields and reapply, external
  deletion/recreation, invalid resource configuration, and app/list lookups.
  The known-empty-list regression and import lifecycle remain pending under P4/P5.
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

The remaining straightforward batch copies 16 source functions across 13 suites:

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
