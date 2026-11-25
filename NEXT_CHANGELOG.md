# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fail with a clear error when SCIM `/Me` returns workspace ID 0 for a workspace-level provider, instead of writing `"0"` to `provider_config.workspace_id` of resources with `provider_config`, which made every later plan fail ([#6225](https://github.com/databricks/terraform-provider-databricks/pull/6225)).

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
