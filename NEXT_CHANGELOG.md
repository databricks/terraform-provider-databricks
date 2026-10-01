# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes

* Fix column handling on a `VIEW` in `databricks_sql_table` ([#5958](https://github.com/databricks/terraform-provider-databricks/pull/5958)). A view's columns are derived from its query, but the provider treated them like table columns and emitted DDL that Databricks rejects with a `PARSE_SYNTAX_ERROR`: `NOT NULL` in `CREATE VIEW`, and `ALTER VIEW ... ADD COLUMN` / `DROP COLUMN` / `RENAME COLUMN` / `ALTER COLUMN ... SET|DROP NOT NULL` on update. Editing the query of a view that has column comments also dropped those comments, because `ALTER VIEW ... AS` clears them and the provider did not re-apply them in the same run. For views, the provider now applies only column comments (via `COMMENT ON COLUMN`), matches columns by name rather than by position, re-applies the configured comments whenever the view definition changes, and suppresses diffs on `column.nullable`, whose server-derived value cannot be set from the configuration.

### Documentation

### Exporter

### Internal Changes
