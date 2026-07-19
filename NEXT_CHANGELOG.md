# NEXT CHANGELOG

## Release v1.122.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_postgres_cdf_config`.
* Add data sources for `databricks_postgres_cdf_status`.

### Bug Fixes

* Fix updating a column comment on a `VIEW` in `databricks_sql_table` ([#5855](https://github.com/databricks/terraform-provider-databricks/pull/5855)). The provider emitted `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a `PARSE_SYNTAX_ERROR`, leaving the change stuck as a perpetual, un-appliable diff. Column comment changes on views are now applied in place via `COMMENT ON COLUMN`, matching how column comments on tables are updated.

### Documentation
* Fixed the formatting of databricks mws_network_connectivity_config terraform resource

### Exporter

### Internal Changes
