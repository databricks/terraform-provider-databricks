# NEXT CHANGELOG

## Release v1.75.0

### New Features and Improvements

 * Add support for `power_bi_task` in jobs ([#4647](https://github.com/databricks/terraform-provider-databricks/pull/4647))
 * Make `spark_version` optional in the context of jobs such that a cluster policy can provide a default value ([#4643](https://github.com/databricks/terraform-provider-databricks/pull/4643))

### Bug Fixes

### Documentation

 * Document `performance_target` in `databricks_job` ([#4651](https://github.com/databricks/terraform-provider-databricks/pull/4651))
 * Add more examples for `databricks_model_serving` ([#4658](https://github.com/databricks/terraform-provider-databricks/pull/4658))

### Exporter

 * Correctly handle account-level identities when generating the code ([#4650](https://github.com/databricks/terraform-provider-databricks/pull/4650))

### Internal Changes
