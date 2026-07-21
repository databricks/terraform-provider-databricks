# NEXT CHANGELOG

## Release v1.123.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix updating a column comment on a `VIEW` in `databricks_sql_table` ([#5855](https://github.com/databricks/terraform-provider-databricks/pull/5855)). The provider emitted `ALTER VIEW ... ALTER COLUMN ... COMMENT`, which Databricks rejects with a `PARSE_SYNTAX_ERROR`, leaving the change stuck as a perpetual, un-appliable diff. Column comment changes on views are now applied in place via `COMMENT ON COLUMN`, matching how column comments on tables are updated.
* Fix `databricks_access_control_rule_set` drift detection when all `grant_rules` are removed outside Terraform ([#5589](https://github.com/databricks/terraform-provider-databricks/issues/5589)).

### Documentation

### Exporter

### Internal Changes

* Use account host check instead of account ID check in `databricks_access_control_rule_set` to determine client type ([#5484](https://github.com/databricks/terraform-provider-databricks/pull/5484)).

* Significantly reduced the number of SCIM and IAM API calls during `terraform plan`/`apply` for large deployments by introducing shared in-memory caches with `sync.RWMutex` and `singleflight` deduplication. Resources `databricks_group`, `databricks_user`, `databricks_group_member`, `databricks_permission_assignment`, and `databricks_mws_permission_assignment` now each issue a single list API call per plan cycle instead of one call per resource instance, eliminating redundant requests and rate-limit (429) errors.
