# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Support `statement_timeout` for all SQL warehouse types; omitting or removing the argument preserves the existing warehouse timeout, while explicit `0` clears the override ([#TBD](https://github.com/databricks/terraform-provider-databricks/pull/TBD)).

* Support fixed-size REALTIME warehouses with `max_dbu`, `statement_timeout`, and `customer_version_label` in `databricks_sql_endpoint` and `databricks_sql_warehouse` ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Bug Fixes

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

* Preserve `statement_timeout`, including `0`, in exports of all SQL warehouse types ([#TBD](https://github.com/databricks/terraform-provider-databricks/pull/TBD)).

* Preserve `auto_stop_mins`, REALTIME `statement_timeout`, and `max_dbu` in warehouse exports ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Internal Changes
