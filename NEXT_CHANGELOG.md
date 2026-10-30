# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Support fixed-size REALTIME warehouses with `max_dbu`, `statement_timeout`, and `customer_version_label` in `databricks_sql_endpoint` and `databricks_sql_warehouse` ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Bug Fixes

* Fix perpetual diffs and `cannot be updated` errors in `databricks_budget` when the API returns alerts or actions in a different order than configured ([#6100](https://github.com/databricks/terraform-provider-databricks/pull/6100)).

  Alerts are now matched by content, so their order no longer matters, and changing an alert replaces it. Plans also ignore the empty filter that the API returns for budgets without one, and the decimal formatting of `override_threshold`.

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

* Preserve `auto_stop_mins`, REALTIME `statement_timeout`, and `max_dbu` in warehouse exports ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Internal Changes
