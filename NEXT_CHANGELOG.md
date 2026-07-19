# NEXT CHANGELOG

## Release v1.122.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_postgres_cdf_config`.
* Add data sources for `databricks_postgres_cdf_status`.

 * Added `warehouse_id` support to `databricks_sql_permissions` resource, allowing `GRANT`/`REVOKE`/`SHOW GRANT` to execute via a SQL warehouse instead of a cluster.

### Bug Fixes

* Fix updating a column comment on a `VIEW` in `databricks_sql_table` ([#5855](https://github.com/databricks/terraform-provider-databricks/pull/5855)). The provider emitted `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a `PARSE_SYNTAX_ERROR`, leaving the change stuck as a perpetual, un-appliable diff. Column comment changes on views are now applied in place via `COMMENT ON COLUMN`, matching how column comments on tables are updated.

### Documentation

### Exporter

### Internal Changes
