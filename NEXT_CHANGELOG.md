# NEXT CHANGELOG

## Release v1.82.0

### Breaking Changes

### New Features and Improvements

 * Add `databricks_users` data source ([#4028](https://github.com/databricks/terraform-provider-databricks/pull/4028))
 * Support configuration of file events in `databricks_external_location` [#4749](https://github.com/databricks/terraform-provider-databricks/pull/4749).
 * Improve support for new fields in `databricks_pipeline` [#4744](https://github.com/databricks/terraform-provider-databricks/pull/4744).

### Bug Fixes

 * Populate `partitions` when reading `databricks_sql_table` ([#4674](https://github.com/databricks/terraform-provider-databricks/pull/4674)).
 * Fail when creating `databricks_query` and `databricks_alert` with already existing names [#4697](https://github.com/databricks/terraform-provider-databricks/pull/4697).

### Documentation

### Exporter

### Internal Changes
