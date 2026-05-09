# NEXT CHANGELOG

## Release v1.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_service_principal` data source failing on account-level provider with `cannot populate provider_config for service principal: failed to resolve workspace_id` ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_service_principals` data source failing on account-level provider with the same `cannot populate provider_config for service principals: failed to resolve workspace_id` regression ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_mws_workspaces` and `databricks_mws_credentials` data sources failing on account-level provider with `cannot populate provider_config for mws workspaces: failed to resolve workspace_id` ([#5672](https://github.com/databricks/terraform-provider-databricks/issues/5672)). These account-only data sources are now exempted from the post-Read workspace-tracking hook, and `provider_config` (which had no effect on them) is now deprecated and will be removed in a future major release.
* Fix workspace-level SCIM data sources and resources (`databricks_user`, `databricks_group`, `databricks_service_principal(s)`) failing with `HTTP 400 Unable to load OAuth Config` when `api = "workspace"` is set on a unified host.

  On a unified host, workspace-level SCIM requests need an `X-Databricks-Org-Id` header so the host can route the call to the right workspace. The Go SDK does not auto-inject this header from `Config.WorkspaceID` for general API calls, so the provider now adds it from `Config.WorkspaceID` for non-account SCIM requests on a unified host.

### Documentation

### Exporter

### Internal Changes
