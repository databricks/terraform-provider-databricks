# NEXT CHANGELOG

## Release v1.105.0

### Breaking Changes

* Return empty string for `data_source_id` in `databricks_sql_warehouse` and `databricks_sql_endpoint` if data source API call fails ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### New Features and Improvements

### Bug Fixes

### Documentation

* Mark `data_source_id` as deprecated in `databricks_sql_warehouse` and `databricks_sql_endpoint` ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### Exporter

* Add exporting for RFA destinations and Delta Sharing providers ([#5337](https://github.com/databricks/terraform-provider-databricks/pull/5337))

### Internal Changes

* Update Go SDK to v0.104.0.
