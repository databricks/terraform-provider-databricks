# NEXT CHANGELOG

## Release v1.89.0

### Breaking Changes

### New Features and Improvements

* Document and handle additional Microsoft Teams options in `databricks_notification_destination` ([#4990](https://github.com/databricks/terraform-provider-databricks/pull/4990))
* Added the following resources and data sources: `account_settings_v2_setting`, `rfa_access_request_destinations` and `workspace_settings_v2_setting` ([#5016](https://github.com/databricks/terraform-provider-databricks/pull/5016)).

### Bug Fixes

* Fix regression with `databricks_group` data source introduced by a recent change ([#4995](https://github.com/databricks/terraform-provider-databricks/pull/4995))

### Documentation

* Document `continuous.task_retry_mode` in `databricks_job` ([#4993](https://github.com/databricks/terraform-provider-databricks/pull/4993))

### Exporter

* Improve handling of dependencies for vector search index ([#4989](https://github.com/databricks/terraform-provider-databricks/pull/4989)).

### Internal Changes

* Update Go SDK to v0.83.0 ([#5016](https://github.com/databricks/terraform-provider-databricks/pull/5016)).
