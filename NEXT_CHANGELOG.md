# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

* Add new `databricks_app_deployment` resource for deploying code to Databricks Apps from workspace filesystem paths or Git repositories ([#5452](https://github.com/databricks/terraform-provider-databricks/pull/5452)).

### Bug Fixes

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
