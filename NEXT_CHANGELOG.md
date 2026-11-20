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

* Fixed `databricks_custom_app_integration` not applying `enable_single_use_refresh_tokens = false` and removal of the `token_access_policy` block, and returning a plan-time error instead of a no-op update when `user_authorized_scopes` is removed ([#6262](https://github.com/databricks/terraform-provider-databricks/pull/6262)).

### Documentation

### Exporter

### Internal Changes
