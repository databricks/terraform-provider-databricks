# NEXT CHANGELOG

## Release v1.113.0

### Breaking Changes

### New Features and Improvements
* Add resource and data source for `databricks_postgres_catalog`.
* Add resource and data source for `databricks_postgres_synced_table`.
* Add resource and data sources for `databricks_environments_workspace_base_environment`.
* Add resource and data source for `databricks_environments_default_workspace_base_environment`.

* Added optional `cloud` argument to `databricks_current_config` data source to explicitly set the cloud type (`aws`, `azure`, `gcp`) instead of relying on host-based detection.

* Added `api` field to dual account/workspace resources (`databricks_user`, `databricks_service_principal`, `databricks_group`, `databricks_group_role`, `databricks_group_member`, `databricks_user_role`, `databricks_service_principal_role`, `databricks_user_instance_profile`, `databricks_group_instance_profile`, `databricks_metastore`, `databricks_metastore_assignment`, `databricks_metastore_data_access`, `databricks_storage_credential`, `databricks_service_principal_secret`, `databricks_access_control_rule_set`) to explicitly control whether account-level or workspace-level APIs are used. This enables support for unified hosts like `api.databricks.com` where the API level cannot be inferred from the host ([#5483](https://github.com/databricks/terraform-provider-databricks/pull/5483)).

### Bug Fixes

* Fixed import inconsistency for `force_destroy` and other schema-only fields causing "Provider produced inconsistent final plan" errors ([#5487](https://github.com/databricks/terraform-provider-databricks/pull/5487)).

### Documentation

### Exporter

### Internal Changes

* Update Go SDK to v0.127.0.
* Use account host check instead of account ID check in `databricks_access_control_rule_set` to determine client type ([#5484](https://github.com/databricks/terraform-provider-databricks/pull/5484)).
