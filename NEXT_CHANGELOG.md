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

* Roll back the `databricks_volume` owner change when an apply that also changes other fields is cancelled while their update is being retried, instead of leaving the new owner on the volume ([#6240](https://github.com/databricks/terraform-provider-databricks/pull/6240)).

### Documentation

### Exporter

### Internal Changes
