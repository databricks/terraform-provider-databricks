# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Preserve configured privileges when revoking `ALL_PRIVILEGES` in `databricks_grants` and `databricks_grant` ([#6198](https://github.com/databricks/terraform-provider-databricks/pull/6198)).

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
