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
* Fixed a panic when a resource holds an enum value that the Go SDK does not know about yet, for example `scan_status` on `databricks_app` deployments, which crashed the exporter ([#6205](https://github.com/databricks/terraform-provider-databricks/issues/6205)).

### Documentation

### Exporter

### Internal Changes
