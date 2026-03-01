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
