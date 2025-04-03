# NEXT CHANGELOG

## Release v1.72.0

### New Features and Improvements

 * Customize and document `event_log` block in `databricks_pipeline` ([#4612](https://github.com/databricks/terraform-provider-databricks/pull/4612))
 * Add automatic clustering support for `databricks_sql_table` ([#4607](https://github.com/databricks/terraform-provider-databricks/pull/4607))

### Bug Fixes

 * Fix `databricks_workspace_binding` for existing resources using `external-location` or `storage-credential` securable types ([#4611](https://github.com/databricks/terraform-provider-databricks/pull/4611)).
 * Suppress diff in `databricks_mlflow_experiment` name ([#4606](https://github.com/databricks/terraform-provider-databricks/pull/4606))

### Documentation

 * Add import instructions for `databricks_share` and `databricks_recipient` ([#4608](https://github.com/databricks/terraform-provider-databricks/pull/4608))

### Exporter

### Internal Changes
