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

* Fail with a clear error when SCIM `/Me` returns workspace ID 0 for a workspace-level provider, instead of writing `"0"` to `provider_config.workspace_id` of resources with `provider_config`, which made every later plan fail ([#6225](https://github.com/databricks/terraform-provider-databricks/pull/6225)).

### Documentation

### Exporter

### Internal Changes
