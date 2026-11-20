# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_parser`.
* Add resource and data sources for `databricks_lakewatch_parser_version`.
* Add data sources for `databricks_lakewatch_builtin_parser`.
* Add data sources for `databricks_lakewatch_builtin_parser_version`.

* Allow `MINUTES` as a `unit` for the `databricks_job` `trigger.periodic` block ([#5972](https://github.com/databricks/terraform-provider-databricks/pull/5972)).

  The Jobs API accepts and schedules minute-based periodic triggers, but the provider rejected them during validation.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
