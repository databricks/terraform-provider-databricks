# NEXT CHANGELOG

## Release v1.105.0

### Breaking Changes

* Return empty string for `data_source_id` in `databricks_sql_warehouse` and `databricks_sql_endpoint` if data source API call fails ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_sql_table` timing out after 50 seconds when creating tables with a custom `warehouse_id` by polling for statement completion instead of cancelling ([#5340](https://github.com/databricks/terraform-provider-databricks/issues/5340))

### Documentation

* Mark `data_source_id` as deprecated in `databricks_sql_warehouse` and `databricks_sql_endpoint` ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### Exporter

### Internal Changes
