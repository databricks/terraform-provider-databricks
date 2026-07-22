# NEXT CHANGELOG

## Release v1.123.0

### Important Changes

### Breaking Changes

### New Features and Improvements

 * Added `user_agent_extra` provider configuration attribute to append products to the `User-Agent` header, equivalent to the `DATABRICKS_USER_AGENT_EXTRA` environment variable ([#5863](https://github.com/databricks/terraform-provider-databricks/pull/5863)).

   This lets Terraform modules built on top of the provider configure usage attribution in their `provider` block without requiring users to set environment variables.

### Bug Fixes

* Fix updating a column comment on a `VIEW` in `databricks_sql_table` ([#5855](https://github.com/databricks/terraform-provider-databricks/pull/5855)). The provider emitted `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a `PARSE_SYNTAX_ERROR`, leaving the change stuck as a perpetual, un-appliable diff. Column comment changes on views are now applied in place via `COMMENT ON COLUMN`, matching how column comments on tables are updated.
* Fix `databricks_access_control_rule_set` drift detection when all `grant_rules` are removed outside Terraform ([#5589](https://github.com/databricks/terraform-provider-databricks/issues/5589)).

### Documentation

### Exporter

### Internal Changes
