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

* Keep resources with `provider_config` in state when SCIM `/Me` answers 404 while a workspace-level provider checks `provider_config.workspace_id` ([#6226](https://github.com/databricks/terraform-provider-databricks/pull/6226)).

  Refresh and destroy now fail with that error. Before, refresh dropped the resource from state without reading it, so the next apply created it again, and destroy removed it from state without deleting it.

### Documentation

### Exporter

### Internal Changes
