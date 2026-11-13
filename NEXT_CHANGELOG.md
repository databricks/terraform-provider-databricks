# NEXT CHANGELOG

## Release v1.143.0

### Important Changes

* `databricks_app` now updates apps through the asynchronous Apps update API ([#6151](https://github.com/databricks/terraform-provider-databricks/pull/6151)).
  The synchronous update API ignored `compute_size`, `compute_min_instances`, and
  `compute_max_instances`, so changes to them never took effect; they are now applied. As
  before, an update replaces the app's updatable fields with the configuration, so fields
  omitted from it are cleared. That now also covers the fields the synchronous API skipped: an
  omitted `compute_size` resets to `MEDIUM`, omitted `compute_min_instances` and
  `compute_max_instances` disable horizontal scaling, and an omitted `git_source` clears the
  app's default git source. Updates now wait for the app update to reach `SUCCEEDED`.

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_assume_group_restricted_sharing_group`.
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

* Added `principal_id` argument to `databricks_git_credential` resource, enabling management of Git credentials on behalf of service principals.
* Add resource and data source for `databricks_postgres_catalog`.
* Add resource and data source for `databricks_postgres_synced_table`.
* Add resource and data sources for `databricks_environments_workspace_base_environment`.
* Add resource and data source for `databricks_environments_default_workspace_base_environment`.

* Added optional `cloud` argument to `databricks_current_config` data source to explicitly set the cloud type (`aws`, `azure`, `gcp`) instead of relying on host-based detection.

* Added `api` field to dual account/workspace resources (`databricks_user`, `databricks_service_principal`, `databricks_group`, `databricks_group_role`, `databricks_group_member`, `databricks_user_role`, `databricks_service_principal_role`, `databricks_user_instance_profile`, `databricks_group_instance_profile`, `databricks_metastore`, `databricks_metastore_assignment`, `databricks_metastore_data_access`, `databricks_storage_credential`, `databricks_service_principal_secret`, `databricks_access_control_rule_set`) to explicitly control whether account-level or workspace-level APIs are used. This enables support for unified hosts like `api.databricks.com` where the API level cannot be inferred from the host ([#5483](https://github.com/databricks/terraform-provider-databricks/pull/5483)).

### Bug Fixes
* Keep the configured values of the `pem_key_passphrase`, `ssl_keystore_key`, `ssl_key_password` and `schema_registry_api_secret` options of `databricks_connection`, which the API doesn't return, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Ignore whitespace-only differences in PEM-encoded options of `databricks_connection`, such as `userProvidedServerCertificate`, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Documentation

* Documented `principal_id` argument for `databricks_git_credential` resource.

### Exporter

### Internal Changes
