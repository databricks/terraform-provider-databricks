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

* Fix perpetual diffs and `cannot be updated` errors in `databricks_budget` when the API returns alerts or actions in a different order than configured ([#6100](https://github.com/databricks/terraform-provider-databricks/pull/6100)).

  Alerts are now matched by content, so their order no longer matters, and changing an alert replaces it. Plans also ignore the empty filter that the API returns for budgets without one, and the decimal formatting of `override_threshold`.

### Documentation

### Exporter

### Internal Changes
