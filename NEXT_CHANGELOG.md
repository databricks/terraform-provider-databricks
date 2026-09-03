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

* Allow `MINUTES` as a `unit` for the `databricks_job` `trigger.periodic` block ([#5972](https://github.com/databricks/terraform-provider-databricks/pull/5972)).

  The Jobs API accepts and schedules minute-based periodic triggers, but the provider rejected them during validation.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
