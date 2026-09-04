# NEXT CHANGELOG

## Release v1.131.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add data source for `databricks_account_iam_external_user_v2`.
* Add data source for `databricks_account_iam_external_service_principal_v2`.
* Add data source for `databricks_account_iam_external_group_v2`.
* Add data source for `databricks_workspace_iam_external_user_v2`.
* Add data source for `databricks_workspace_iam_external_service_principal_v2`.
* Add data source for `databricks_workspace_iam_external_group_v2`.
* Add resource and data source for `databricks_postgres_snapshot_schedule`.

### Bug Fixes

* Fixed `databricks_registered_model` aliases not being applied to Unity Catalog by using dedicated `SetAlias`/`DeleteAlias` API calls ([#5448](https://github.com/databricks/terraform-provider-databricks/pull/5448)).

### Documentation

### Exporter

### Internal Changes
