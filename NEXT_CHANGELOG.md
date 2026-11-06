# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_datasource`.
* Add resource and data source for `databricks_lakewatch_datasource_config`.
* Add resource and data source for `databricks_lakewatch_rule_config`.
* Add resource and data source for `databricks_lakewatch_alert_config`.
* Add resource and data source for `databricks_lakewatch_common_config`.
* Add resource and data sources for `databricks_lakewatch_rule`.

* Add `databricks_app_slack` and `databricks_app_teams` configuration blocks to `databricks_notification_destination`, and accept `DATABRICKS_APP_SLACK` and `DATABRICKS_APP_TEAMS` as `type` in the `databricks_notification_destinations` data source ([#6131](https://github.com/databricks/terraform-provider-databricks/pull/6131)).

* Added `principal_id` argument to `databricks_git_credential` resource, enabling management of Git credentials on behalf of service principals.
* Add resource and data source for `databricks_postgres_catalog`.
* Add resource and data source for `databricks_postgres_synced_table`.
* Add resource and data sources for `databricks_environments_workspace_base_environment`.
* Add resource and data source for `databricks_environments_default_workspace_base_environment`.

* Added optional `cloud` argument to `databricks_current_config` data source to explicitly set the cloud type (`aws`, `azure`, `gcp`) instead of relying on host-based detection.

* Added `api` field to dual account/workspace resources (`databricks_user`, `databricks_service_principal`, `databricks_group`, `databricks_group_role`, `databricks_group_member`, `databricks_user_role`, `databricks_service_principal_role`, `databricks_user_instance_profile`, `databricks_group_instance_profile`, `databricks_metastore`, `databricks_metastore_assignment`, `databricks_metastore_data_access`, `databricks_storage_credential`, `databricks_service_principal_secret`, `databricks_access_control_rule_set`) to explicitly control whether account-level or workspace-level APIs are used. This enables support for unified hosts like `api.databricks.com` where the API level cannot be inferred from the host ([#5483](https://github.com/databricks/terraform-provider-databricks/pull/5483)).

### Bug Fixes

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

* Preserve `databricks_pipeline.expected_last_modified` during refresh to avoid post-apply drift ([#6133](https://github.com/databricks/terraform-provider-databricks/pull/6133)).

### Documentation

* Documented `principal_id` argument for `databricks_git_credential` resource.

### Exporter

### Internal Changes
