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

* Added `cascade_on_destroy` attribute to `databricks_pipeline` to control whether destroying a pipeline also deletes its datasets (materialized views, streaming tables, and views). Defaults to `true`; set to `false` to preserve the datasets on destroy ([#5860](https://github.com/databricks/terraform-provider-databricks/pull/5860)).

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
