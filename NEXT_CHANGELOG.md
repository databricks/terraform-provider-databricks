# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix perpetual diffs and `cannot be updated` errors in `databricks_budget` when the API returns alerts or actions in a different order than configured ([#6100](https://github.com/databricks/terraform-provider-databricks/pull/6100)).

  Alerts are now matched by content, so their order no longer matters, and changing an alert replaces it. Plans also ignore the empty filter that the API returns for budgets without one, and the decimal formatting of `override_threshold`.

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
