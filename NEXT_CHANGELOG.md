# NEXT CHANGELOG

## Release v1.95.0

### New Features and Improvements

* Add `arm` option to `databricks_node_type` instead of `graviton` ([#5028](https://github.com/databricks/terraform-provider-databricks/pull/5028))
* Optimize `databricks_grant` and `databricks_grants` to not call the `Update` API if the requested permissions are already granted ([#5095](https://github.com/databricks/terraform-provider-databricks/pull/5095))

### New Features and Improvements

* Added `expected_workspace_status` to `databricks_mws_workspaces` to support creating workspaces in provisioning status ([#5019](https://github.com/databricks/terraform-provider-databricks/pull/5019))

### Bug Fixes

* Mark `workload_size` as not computed in `databricks_model_serving` ([#5111](https://github.com/databricks/terraform-provider-databricks/pull/5111))
* Fix Inconsistent Plan Errors in Permissions Resource ([#5091](https://github.com/databricks/terraform-provider-databricks/pull/5091))

### Documentation

* Add documentation on the `service_principal_client_id` attribute of `databricks_app` and related [#5134](https://github.com/databricks/terraform-provider-databricks/pull/5134)
* Add documentation for Unified Terraform Provider ([#5122](https://github.com/databricks/terraform-provider-databricks/pull/5122))

### Exporter
* Remove `METRIC_VIEW` from `sql_table` resource ([#5135](https://github.com/databricks/terraform-provider-databricks/pull/5135))

### Internal Changes

* Fix custom_app_integration_test ([#5129](https://github.com/databricks/terraform-provider-databricks/pull/5129))
