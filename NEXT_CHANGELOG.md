# NEXT CHANGELOG

## Release v1.89.0

### Breaking Changes

### New Features and Improvements

* Document and handle additional Microsoft Teams options in `databricks_notification_destination` ([#4990](https://github.com/databricks/terraform-provider-databricks/pull/4990))
* Optimize `databricks_grant` and `databricks_grants` to not call the `Update` API if the requested permissions are already granted ([#5095](https://github.com/databricks/terraform-provider-databricks/pull/5095))

### Bug Fixes

* Fix regression with `databricks_group` data source introduced by a recent change ([#4995](https://github.com/databricks/terraform-provider-databricks/pull/4995))

### Documentation

* Document `continuous.task_retry_mode` in `databricks_job` ([#4993](https://github.com/databricks/terraform-provider-databricks/pull/4993))

### Exporter

* Improve handling of dependencies for vector search index ([#4989](https://github.com/databricks/terraform-provider-databricks/pull/4989)).

### Internal Changes
