# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

* Added `principal_id` argument to `databricks_git_credential` resource, enabling management of Git credentials on behalf of service principals.
* Add resource and data source for `databricks_postgres_catalog`.
* Add resource and data source for `databricks_postgres_synced_table`.
* Add resource and data sources for `databricks_environments_workspace_base_environment`.
* Add resource and data source for `databricks_environments_default_workspace_base_environment`.

* Added optional `cloud` argument to `databricks_current_config` data source to explicitly set the cloud type (`aws`, `azure`, `gcp`) instead of relying on host-based detection.

* Added `api` field to dual account/workspace resources (`databricks_user`, `databricks_service_principal`, `databricks_group`, `databricks_group_role`, `databricks_group_member`, `databricks_user_role`, `databricks_service_principal_role`, `databricks_user_instance_profile`, `databricks_group_instance_profile`, `databricks_metastore`, `databricks_metastore_assignment`, `databricks_metastore_data_access`, `databricks_storage_credential`, `databricks_service_principal_secret`, `databricks_access_control_rule_set`) to explicitly control whether account-level or workspace-level APIs are used. This enables support for unified hosts like `api.databricks.com` where the API level cannot be inferred from the host ([#5483](https://github.com/databricks/terraform-provider-databricks/pull/5483)).

### Bug Fixes

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.

### Documentation

### Exporter

### Internal Changes

* Add `internal/retrier` package for unified retry and backoff handling ([#5746](https://github.com/databricks/terraform-provider-databricks/pull/5746)).
* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).
