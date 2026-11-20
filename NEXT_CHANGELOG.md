# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_parser`.
* Add resource and data sources for `databricks_lakewatch_parser_version`.
* Add data sources for `databricks_lakewatch_builtin_parser`.
* Add data sources for `databricks_lakewatch_builtin_parser_version`.

* Support GCP targets in the opt-in Plugin Framework implementation of `databricks_mws_ncc_private_endpoint_rule`. Set `DATABRICKS_TF_ENABLED_PF_RESOURCES=databricks_mws_ncc_private_endpoint_rule` to use this support; the default SDKv2 implementation is unchanged ([#5973](https://github.com/databricks/terraform-provider-databricks/pull/5973)).

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
