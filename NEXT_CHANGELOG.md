# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Added `cascade_on_destroy` attribute to `databricks_pipeline` to control whether destroying a pipeline also deletes its datasets (materialized views, streaming tables, and views). Defaults to `true`; set to `false` to preserve the datasets on destroy ([#5860](https://github.com/databricks/terraform-provider-databricks/pull/5860)).

### Bug Fixes

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
