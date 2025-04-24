# NEXT CHANGELOG

## Release v1.75.0

### New Features and Improvements

 * Add support for `power_bi_task` in jobs ([#4647](https://github.com/databricks/terraform-provider-databricks/pull/4647))
 * Add support for `dashboard_task` in jobs ([#4646](https://github.com/databricks/terraform-provider-databricks/pull/4646))
 * Add `compute_mode` to `databricks_mws_workspaces` to support creating serverless workspaces ([#4670](https://github.com/databricks/terraform-provider-databricks/pull/4670)).

### Bug Fixes

### Documentation

 * Document `performance_target` in `databricks_job` ([#4651](https://github.com/databricks/terraform-provider-databricks/pull/4651))
 * Add more examples for `databricks_model_serving` ([#4658](https://github.com/databricks/terraform-provider-databricks/pull/4658))
 * Document `on_streaming_backlog_exceeded` in email/webhook notifications in `databricks_job` ([#4660](https://github.com/databricks/terraform-provider-databricks/pull/4660))
 * Refresh `spark_python_task` option in `databricks_job` ([#4666](https://github.com/databricks/terraform-provider-databricks/pull/4666))

### Exporter

 * Correctly handle account-level identities when generating the code ([#4650](https://github.com/databricks/terraform-provider-databricks/pull/4650))
 * Add export of dashboard tasks in `datarbicks_job` ([#4665](https://github.com/databricks/terraform-provider-databricks/pull/4665))
 * Add export of PowerBI tasks in `datarbicks_job` ([#4668](https://github.com/databricks/terraform-provider-databricks/pull/4668))
 * Add `Ignore` implementation for `databricks_grants` to fix issue with wrongly generated dependencies ([#4650](https://github.com/databricks/terraform-provider-databricks/pull/4650))
 * Improve handling of `owner` for UC resources ([#4669](https://github.com/databricks/terraform-provider-databricks/pull/4669))

### Internal Changes
