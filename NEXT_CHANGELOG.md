# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Support fixed-size REALTIME warehouses with `max_dbu`, `statement_timeout`, and `customer_version_label` in `databricks_sql_endpoint` and `databricks_sql_warehouse` ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Bug Fixes

* Fixed perpetual `databricks_share` drift when CDF and history sharing are enabled ([#5949](https://github.com/databricks/terraform-provider-databricks/pull/5949)).

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

* Preserve `auto_stop_mins`, REALTIME `statement_timeout`, and `max_dbu` in warehouse exports ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Internal Changes
