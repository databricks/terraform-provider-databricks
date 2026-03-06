# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

* Added `databricks_postgres_database` resource and data source.
* Renamed `databricks_apps_space` to `databricks_app_space`.

### Bug Fixes

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
* Use workspace `GetStatus` API to check parent folder existence before creating Lakeview dashboards, replacing fragile error string matching.
* Use workspace `GetStatus` API to check parent folder existence before uploading workspace files, replacing fragile error string matching.
* Skip `TestMwsAccNetworkConnectivityConfig` test to unblock merging PRs in the repository
* * Use workspace `GetStatus` API to check parent folder existence before importing notebooks, replacing fragile error string matching.
