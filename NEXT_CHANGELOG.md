# NEXT CHANGELOG

## Release v1.122.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_postgres_cdf_config`.
* Add data sources for `databricks_postgres_cdf_status`.

* Add `provider_config` support for manual plugin framework resources and data sources([#5127](https://github.com/databricks/terraform-provider-databricks/pull/5127))

* Added support for custom instance profiles on instance pools on AWS ([#5144](https://github.com/databricks/terraform-provider-databricks/pull/5144))

* Added `deployment_names` attribute to `databricks_mws_workspaces` data block ([#5100](https://github.com/databricks/terraform-provider-databricks/pull/5100))

### Bug Fixes

* Fix updating a column comment on a `VIEW` in `databricks_sql_table` ([#5855](https://github.com/databricks/terraform-provider-databricks/pull/5855)). The provider emitted `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a `PARSE_SYNTAX_ERROR`, leaving the change stuck as a perpetual, un-appliable diff. Column comment changes on views are now applied in place via `COMMENT ON COLUMN`, matching how column comments on tables are updated.

### Documentation

### Exporter

### Internal Changes
