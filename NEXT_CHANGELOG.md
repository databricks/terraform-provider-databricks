# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fixed `enable_file_events = false` having no effect on `databricks_external_location` when Databricks had enabled file events automatically ([#6200](https://github.com/databricks/terraform-provider-databricks/pull/6200)).

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
