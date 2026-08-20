# NEXT CHANGELOG

## Release v1.129.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

* Skip system-managed jobs during export and add missing file references for job task parameters ([#5956](https://github.com/databricks/terraform-provider-databricks/issues/5956)).
* Added support for exporting `databricks_endpoint` resource ([#5951](https://github.com/databricks/terraform-provider-databricks/pull/5951)).
* Add an `exporter` dimension to the user agent ([#5954](https://github.com/databricks/terraform-provider-databricks/pull/5954)).
* Allow to generate named variables from references; introduce `databricks_account_id` variable for account-level exports; bug fixes ([#5952](https://github.com/databricks/terraform-provider-databricks/pull/5952)).
* Preserve zero `value` fields when exporting `databricks_workspace_setting_v2` and `databricks_account_setting_v2` ([#5955](https://github.com/databricks/terraform-provider-databricks/issues/5955)).
* Resolve references embedded in `databricks_cluster_policy` definitions instead of emitting hardcoded values ([#5953](https://github.com/databricks/terraform-provider-databricks/issues/5953)).

### Internal Changes

* Use account host check instead of account ID check in `databricks_access_control_rule_set` to determine client type ([#5484](https://github.com/databricks/terraform-provider-databricks/pull/5484)).

* Significantly reduced the number of SCIM and IAM API calls during `terraform plan`/`apply` for large deployments by introducing shared in-memory caches with `sync.RWMutex` and `singleflight` deduplication. Resources `databricks_group`, `databricks_user`, `databricks_group_member`, `databricks_permission_assignment`, and `databricks_mws_permission_assignment` now each issue a single list API call per plan cycle instead of one call per resource instance, eliminating redundant requests and rate-limit (429) errors.
