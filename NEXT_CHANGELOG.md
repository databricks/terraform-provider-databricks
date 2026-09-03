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

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
