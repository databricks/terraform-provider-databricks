# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_registered_model` aliases not being applied to Unity Catalog by using dedicated `SetAlias`/`DeleteAlias` API calls ([#5448](https://github.com/databricks/terraform-provider-databricks/pull/5448)).

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
