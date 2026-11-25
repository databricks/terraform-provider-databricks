# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fixed a panic when a resource holds an enum value that the Go SDK does not know about yet, for example `scan_status` on `databricks_app` deployments, which crashed the exporter ([#6205](https://github.com/databricks/terraform-provider-databricks/issues/6205)).

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
