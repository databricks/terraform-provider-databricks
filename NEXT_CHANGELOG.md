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
* Fixed `databricks_sql_endpoint` resource failing when `min_num_clusters` was changed outside of Terraform by adding `Default: 1` to match `max_num_clusters` behavior ([#5294](https://github.com/databricks/terraform-provider-databricks/issues/5294)).

### Documentation

### Exporter

### Internal Changes
