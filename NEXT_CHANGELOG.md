# NEXT CHANGELOG

## Release v1.129.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix column handling on a `VIEW` in `databricks_sql_table` ([#5958](https://github.com/databricks/terraform-provider-databricks/pull/5958)). A view's columns are derived from its query, but the provider treated them like table columns and emitted DDL that Databricks rejects with a `PARSE_SYNTAX_ERROR`: `NOT NULL` in `CREATE VIEW`, and `ALTER VIEW ... ADD COLUMN` / `DROP COLUMN` / `RENAME COLUMN` / `ALTER COLUMN ... SET|DROP NOT NULL` on update. Editing the query of a view that has column comments also dropped those comments, because `ALTER VIEW ... AS` clears them and the provider did not re-apply them in the same run. For views, the provider now applies only column comments (via `COMMENT ON COLUMN`), matches columns by name rather than by position, re-applies the configured comments whenever the view definition changes, and suppresses diffs on `column.nullable`, whose server-derived value cannot be set from the configuration.

### Documentation

### Exporter

* Skip system-managed jobs during export and add missing file references for job task parameters ([#5956](https://github.com/databricks/terraform-provider-databricks/issues/5956)).
* Added support for exporting `databricks_endpoint` resource ([#5951](https://github.com/databricks/terraform-provider-databricks/pull/5951)).
* Add an `exporter` dimension to the user agent ([#5954](https://github.com/databricks/terraform-provider-databricks/pull/5954)).
* Allow to generate named variables from references; introduce `databricks_account_id` variable for account-level exports; bug fixes ([#5952](https://github.com/databricks/terraform-provider-databricks/pull/5952)).
* Preserve zero `value` fields when exporting `databricks_workspace_setting_v2` and `databricks_account_setting_v2` ([#5955](https://github.com/databricks/terraform-provider-databricks/issues/5955)).
* Resolve references embedded in `databricks_cluster_policy` definitions instead of emitting hardcoded values ([#5953](https://github.com/databricks/terraform-provider-databricks/issues/5953)).

### Internal Changes
