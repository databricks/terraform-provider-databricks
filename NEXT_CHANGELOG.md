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
* Preserve the `refresh_token` option when refreshing `databricks_connection` to prevent permanent configuration drift ([#5987](https://github.com/databricks/terraform-provider-databricks/pull/5987)).

### Documentation

### Exporter

### Internal Changes
