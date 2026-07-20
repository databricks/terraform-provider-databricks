# NEXT CHANGELOG

## Release v1.123.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add `databricks_genie_space` resource and `databricks_genie_spaces` data source, plus Genie Space support in `databricks_permissions` via the new `genie_space_id` attribute ([#5770](https://github.com/databricks/terraform-provider-databricks/pull/5770)).

  The resource normalizes `serialized_space` to suppress whitespace and key-order diffs and auto-creates a missing `parent_path` on first apply. Delete is trash-aware (treats an already-trashed space as a successful delete). Tags can be attached using the existing `databricks_workspace_entity_tag_assignment` resource with `entity_type = "geniespaces"`.

### Bug Fixes

* Fix updating a column comment on a `VIEW` in `databricks_sql_table` ([#5855](https://github.com/databricks/terraform-provider-databricks/pull/5855)). The provider emitted `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a `PARSE_SYNTAX_ERROR`, leaving the change stuck as a perpetual, un-appliable diff. Column comment changes on views are now applied in place via `COMMENT ON COLUMN`, matching how column comments on tables are updated.
* Fix `databricks_access_control_rule_set` drift detection when all `grant_rules` are removed outside Terraform ([#5589](https://github.com/databricks/terraform-provider-databricks/issues/5589)).

### Documentation

### Exporter

### Internal Changes
