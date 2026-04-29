# NEXT CHANGELOG

## Release v1.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_service_principal` data source failing on account-level provider with `cannot populate provider_config for service principal: failed to resolve workspace_id` ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_service_principals` data source failing on account-level provider with the same `cannot populate provider_config for service principals: failed to resolve workspace_id` regression ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Remove invalid `provider_config` attribute from account-only data sources `databricks_mws_workspaces` and `databricks_mws_credentials` ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)).

### Documentation

### Exporter

### Internal Changes
