# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

* For workspace-level hosts, resolve and validate the provider's `workspace_id` from the host's `/.well-known/databricks-config` discovery metadata instead of a SCIM `/Me` call ([#5922](https://github.com/databricks/terraform-provider-databricks/pull/5922)).

  This removes the authenticated `/Me` request on workspace hosts (avoiding false failures for service principals that can manage resources but cannot call `/Me`), and makes a `workspace_id` that disagrees with the host fail at plan time instead of at apply. Hosts whose metadata does not advertise a `workspace_id` fall back to the previous `/Me` behavior.

* Added `principal_id` argument to `databricks_git_credential` resource, enabling management of Git credentials on behalf of service principals.
* Add resource and data source for `databricks_postgres_catalog`.
* Add resource and data source for `databricks_postgres_synced_table`.
* Add resource and data sources for `databricks_environments_workspace_base_environment`.
* Add resource and data source for `databricks_environments_default_workspace_base_environment`.

* Added optional `cloud` argument to `databricks_current_config` data source to explicitly set the cloud type (`aws`, `azure`, `gcp`) instead of relying on host-based detection.

* Added `api` field to dual account/workspace resources (`databricks_user`, `databricks_service_principal`, `databricks_group`, `databricks_group_role`, `databricks_group_member`, `databricks_user_role`, `databricks_service_principal_role`, `databricks_user_instance_profile`, `databricks_group_instance_profile`, `databricks_metastore`, `databricks_metastore_assignment`, `databricks_metastore_data_access`, `databricks_storage_credential`, `databricks_service_principal_secret`, `databricks_access_control_rule_set`) to explicitly control whether account-level or workspace-level APIs are used. This enables support for unified hosts like `api.databricks.com` where the API level cannot be inferred from the host ([#5483](https://github.com/databricks/terraform-provider-databricks/pull/5483)).

### Bug Fixes

### Documentation

* Documented `principal_id` argument for `databricks_git_credential` resource.

### Exporter

### Internal Changes
