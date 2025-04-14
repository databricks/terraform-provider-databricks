# NEXT CHANGELOG

## Release v1.73.0

### New Features and Improvements

 * Mark `databricks_sql_dashaboard` as deprecated ([#4641](https://github.com/databricks/terraform-provider-databricks/pull/4641))
 * Add a new settings resource `databricks_disable_legacy_dbfs_setting` ([#4605](https://github.com/databricks/terraform-provider-databricks/pull/4605))

### Bug Fixes

 * Increase the default HTTP timeout to 65 seconds.

### Documentation

 * Document `user_api_scopes` in `databricks_app` resource and data sources ([#4614](https://github.com/databricks/terraform-provider-databricks/pull/4614))
 * Document new fields in `databricks_model_serving` resource ([#4615](https://github.com/databricks/terraform-provider-databricks/pull/4615))

### Exporter

 * Add export of `databricks_mws_network_connectivity_config` and `databricks_mws_ncc_private_endpoint_rule` ([#4613](https://github.com/databricks/terraform-provider-databricks/pull/4613))

### Internal Changes

* Refactor `databricks_mws_workspaces` resource and datasource, as well as `databricks_mws_ncc_binding`, to use the Go SDK ([#4633](https://github.com/databricks/terraform-provider-databricks/pull/4633)).
* Add `TestMwsAccGcpWorkspaces` and `TestMwsAccGcpByovpcWorkspaces` to flaky test ([#4624](https://github.com/databricks/terraform-provider-databricks/pull/4624)).
* Bump Go SDK to latest release ([#4645](https://github.com/databricks/terraform-provider-databricks/pull/4645))
