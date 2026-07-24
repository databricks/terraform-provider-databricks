# NEXT CHANGELOG

## Release v1.123.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add `trace_location` to `databricks_mlflow_experiment` for storing experiment traces in a Unity Catalog schema ([#5869](https://github.com/databricks/terraform-provider-databricks/pull/5869)). The block is immutable (`ForceNew`); `table_prefix` is optional and, when omitted, server-defaulted, with the resolved value exposed on the read-only `effective_table_prefix`.

### Bug Fixes

* Fix updating a column comment on a `VIEW` in `databricks_sql_table` ([#5855](https://github.com/databricks/terraform-provider-databricks/pull/5855)). The provider emitted `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a `PARSE_SYNTAX_ERROR`, leaving the change stuck as a perpetual, un-appliable diff. Column comment changes on views are now applied in place via `COMMENT ON COLUMN`, matching how column comments on tables are updated.
* Fix `databricks_access_control_rule_set` drift detection when all `grant_rules` are removed outside Terraform ([#5589](https://github.com/databricks/terraform-provider-databricks/issues/5589)).

### Documentation

### Exporter

### Internal Changes
