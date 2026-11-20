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
* Fix `databricks_workspace_file` and `databricks_notebook` creation failing with `The parent folder (...) does not exist.` when other resources create the same parent folder concurrently ([#6145](https://github.com/databricks/terraform-provider-databricks/issues/6145)).

### Documentation

### Exporter

### Internal Changes
