# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Added `statement_timeout` support to the `databricks_sql_endpoint` resource and `databricks_sql_warehouse` data source ([#6089](https://github.com/databricks/terraform-provider-databricks/pull/6089)).

  The resource defaults to `0` (workspace inheritance). Removing the argument clears any existing warehouse override on the next apply.

### Bug Fixes

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

### Internal Changes
