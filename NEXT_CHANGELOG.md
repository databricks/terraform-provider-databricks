# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_parser`.
* Add resource and data sources for `databricks_lakewatch_parser_version`.
* Add data sources for `databricks_lakewatch_builtin_parser`.
* Add data sources for `databricks_lakewatch_builtin_parser_version`.

### Bug Fixes
* Fix renaming a `databricks_catalog`: the update now sends `new_name`, so changing `name` renames the catalog instead of reporting success without a rename ([#6274](https://github.com/databricks/terraform-provider-databricks/pull/6274)).

### Documentation

### Exporter

### Internal Changes
