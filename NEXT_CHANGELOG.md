# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_sql_table` timing out after 50 seconds when creating tables with a custom `warehouse_id` by polling for statement completion instead of cancelling ([#5340](https://github.com/databricks/terraform-provider-databricks/issues/5340))

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
* Use workspace `GetStatus` API to check parent folder existence before creating Lakeview dashboards, replacing fragile error string matching.
* Use workspace `GetStatus` API to check parent folder existence before uploading workspace files, replacing fragile error string matching.
