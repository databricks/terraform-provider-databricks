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
* Keep a valid `databricks_volume` ID when the create or update response has no `full_name`, so the read after the write no longer fails ([#6208](https://github.com/databricks/terraform-provider-databricks/pull/6208)).

### Documentation

### Exporter

### Internal Changes
