# NEXT CHANGELOG

## Release v1.73.0

### New Features and Improvements

 * Add `databricks_disable_legacy_access_setting` resource to disable legacy access methods ([#4578](https://github.com/databricks/terraform-provider-databricks/pull/4578)).
 * Customize and document `event_log` block in `databricks_pipeline` ([#4612](https://github.com/databricks/terraform-provider-databricks/pull/4612))
 * Add automatic clustering support for `databricks_sql_table` ([#4607](https://github.com/databricks/terraform-provider-databricks/pull/4607))
 * Add a new settings resource `databricks_disable_legacy_dbfs_setting` ([#4605](https://github.com/databricks/terraform-provider-databricks/pull/4605))

### Bug Fixes

### Documentation

 * Document `user_api_scopes` in `databricks_app` resource and data sources ([#4614](https://github.com/databricks/terraform-provider-databricks/pull/4614))
 * Document new fields in `databricks_model_serving` resource ([#4615](https://github.com/databricks/terraform-provider-databricks/pull/4615))

### Exporter

 * Add export of `databricks_mws_network_connectivity_config` and `databricks_mws_ncc_private_endpoint_rule` ([#4613](https://github.com/databricks/terraform-provider-databricks/pull/4613))

### Internal Changes
