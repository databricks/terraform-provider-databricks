# NEXT CHANGELOG

## Release v1.132.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

* Fix column handling on a `VIEW` in `databricks_sql_table` ([#5958](https://github.com/databricks/terraform-provider-databricks/pull/5958)). A view's columns are derived from its query, but the provider treated them like table columns and emitted DDL that Databricks rejects with a `PARSE_SYNTAX_ERROR`: `NOT NULL` in `CREATE VIEW`, and `ALTER VIEW ... ADD COLUMN` / `DROP COLUMN` / `RENAME COLUMN` / `ALTER COLUMN ... SET|DROP NOT NULL` on update. Editing the query of a view that has column comments also dropped those comments, because `ALTER VIEW ... AS` clears them and the provider did not re-apply them in the same run. For views, the provider now applies only column comments (via `COMMENT ON COLUMN`), matches columns by name rather than by position, re-applies the configured comments whenever the view definition changes, and suppresses diffs on `column.nullable`, whose server-derived value cannot be set from the configuration.

### Documentation

* Document the `auto_deploy` and `caller_credential_id` fields of the `databricks_app` `git_repository` block, the `git_source` block, and the top-level `source_code_path` argument. Clarify that `auto_deploy` requires `git_source` to specify a `branch`. These fields become manageable with [#5977](https://github.com/databricks/terraform-provider-databricks/pull/5977).

### Exporter

### Internal Changes
