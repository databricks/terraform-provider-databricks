# NEXT CHANGELOG

## Release v1.119.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

* Document that read-only workspace bindings aren't applicable for non-catalog objects ([#5611](https://github.com/databricks/terraform-provider-databricks/pull/5611))

### Exporter

### Internal Changes

* Use account host check instead of account ID check in `databricks_access_control_rule_set` to determine client type ([#5484](https://github.com/databricks/terraform-provider-databricks/pull/5484)).

* Significantly reduced the number of SCIM and IAM API calls during `terraform plan`/`apply` for large deployments by introducing shared in-memory caches with `sync.RWMutex` and `singleflight` deduplication. Resources `databricks_group`, `databricks_user`, `databricks_group_member`, `databricks_permission_assignment`, and `databricks_mws_permission_assignment` now each issue a single list API call per plan cycle instead of one call per resource instance, eliminating redundant requests and rate-limit (429) errors.
