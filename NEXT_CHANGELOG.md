# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

* For workspace-level hosts, resolve and validate the provider's `workspace_id` from the host's `/.well-known/databricks-config` discovery metadata instead of a SCIM `/Me` call ([#5922](https://github.com/databricks/terraform-provider-databricks/pull/5922)).

  This removes the authenticated `/Me` request on workspace hosts (avoiding false failures for service principals that can manage resources but cannot call `/Me`), and makes a `workspace_id` that disagrees with the host fail at plan time instead of at apply. Hosts whose metadata does not advertise a `workspace_id` fall back to the previous `/Me` behavior.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Use account host check instead of account ID check in `databricks_access_control_rule_set` to determine client type ([#5484](https://github.com/databricks/terraform-provider-databricks/pull/5484)).

* Significantly reduced the number of SCIM and IAM API calls during `terraform plan`/`apply` for large deployments by introducing shared in-memory caches with `sync.RWMutex` and `singleflight` deduplication. Resources `databricks_group`, `databricks_user`, `databricks_group_member`, `databricks_permission_assignment`, and `databricks_mws_permission_assignment` now each issue a single list API call per plan cycle instead of one call per resource instance, eliminating redundant requests and rate-limit (429) errors.
