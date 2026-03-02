# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed AI Gateway rate limits not being sent when `calls` or `tokens` is explicitly set to `0` in `databricks_model_serving` resource ([#5333](https://github.com/databricks/terraform-provider-databricks/issues/5333)).

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
