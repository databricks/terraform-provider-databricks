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

* Require replacement of `databricks_budget` when its filter changes ([#6091](https://github.com/databricks/terraform-provider-databricks/pull/6091)).

  Replacement creates a new budget; accumulated usage and alert behavior should be reviewed before applying.

### Documentation

### Exporter

### Internal Changes
