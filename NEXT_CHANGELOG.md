# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_custom_app_integration` not applying `enable_single_use_refresh_tokens = false` and removal of the `token_access_policy` block, and returning a plan-time error instead of a no-op update when `user_authorized_scopes` is removed ([#6262](https://github.com/databricks/terraform-provider-databricks/pull/6262)).

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
