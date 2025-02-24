# NEXT CHANGELOG

## Release v1.67.0

### New Features and Improvements

 * Expose more attributes for `databricks_connection` resource ([#4502](https://github.com/databricks/terraform-provider-databricks/pull/4502)).

### Bug Fixes

 * Mark `default_catalog_name` attribute in `databricks_metastore_assignment` as deprecated ([#4522](https://github.com/databricks/terraform-provider-databricks/pull/4522))

### Documentation

 * Update `databricks_cluster` and `databricks_clusters` data source documentation ([#4506](https://github.com/databricks/terraform-provider-databricks/pull/4506)).

### Exporter

 * Explicitly abort execution via `panic` if list of users can't be fetched ([#4500](https://github.com/databricks/terraform-provider-databricks/pull/4500)).
 * Fix matching on `workspace_path` and refactoring ([#4504](https://github.com/databricks/terraform-provider-databricks/pull/4504)).

### Internal Changes

 * Bump golang version to 1.24.0 ([#4508](https://github.com/databricks/terraform-provider-databricks/pull/4508)).