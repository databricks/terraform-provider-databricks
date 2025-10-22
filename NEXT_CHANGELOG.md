# NEXT CHANGELOG

## Release v1.95.0

### Breaking Changes

### New Features and Improvements

* Add `arm` option to `databricks_node_type` instead of `graviton` ([#5028](https://github.com/databricks/terraform-provider-databricks/pull/5028))
* Add `provider_config` support for manual SDKv2 resources and data sources on Go SDK([#5115](https://github.com/databricks/terraform-provider-databricks/pull/5115))
* Added `expected_workspace_status` to `databricks_mws_workspaces` to support creating workspaces in provisioning status ([#5019](https://github.com/databricks/terraform-provider-databricks/pull/5019))

### Bug Fixes

* Fix Inconsistent Plan Errors in Permissions Resource ([#5091](https://github.com/databricks/terraform-provider-databricks/pull/5091))

### Documentation

* Add documentation on the `service_principal_client_id` attribute of `databricks_app` and related [#5134](https://github.com/databricks/terraform-provider-databricks/pull/5134)
### Exporter

### Internal Changes

* Fix custom_app_integration_test ([#5129](https://github.com/databricks/terraform-provider-databricks/pull/5129))
