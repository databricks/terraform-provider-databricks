# NEXT CHANGELOG

## Release v1.67.0

### New Features and Improvements

 * Add support for cluster logs delivery to UC Volumes ([#4492](https://github.com/databricks/terraform-provider-databricks/pull/4492)).
 * Expose more attributes for `databricks_connection` resource ([#4502](https://github.com/databricks/terraform-provider-databricks/pull/4502)).

### Bug Fixes

 * Delete `databricks_sql_endpoint` that failed to start ([#4520](https://github.com/databricks/terraform-provider-databricks/pull/4520))

### Documentation

 * Update `databricks_cluster` and `databricks_clusters` data source documentation ([#4506](https://github.com/databricks/terraform-provider-databricks/pull/4506)).

### Exporter

 * Explicitly abort execution via `panic` if list of users can't be fetched ([#4500](https://github.com/databricks/terraform-provider-databricks/pull/4500)).
 * Fix matching on `workspace_path` and refactoring ([#4504](https://github.com/databricks/terraform-provider-databricks/pull/4504)).

### Internal Changes

 * Bump golang version to 1.24.0 ([#4508](https://github.com/databricks/terraform-provider-databricks/pull/4508)).