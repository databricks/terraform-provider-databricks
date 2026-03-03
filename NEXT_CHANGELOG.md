# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

* Added support for updating `git_repository` on `databricks_app` resource and switched updates for `description`, `resources`, `budget_policy_id`, `compute_size`, `usage_policy_id`, and `user_api_scopes` to use the async `CreateUpdate` API, as the previous synchronous `Update` API does not support `git_repository` changes and does not handle fields that require async processing ([#5378](https://github.com/databricks/terraform-provider-databricks/pull/5378))

### Bug Fixes

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
