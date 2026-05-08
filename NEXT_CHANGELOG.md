# NEXT CHANGELOG

## Release v1.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_service_principal` data source failing on account-level provider with `cannot populate provider_config for service principal: failed to resolve workspace_id` ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_service_principals` data source failing on account-level provider with the same `cannot populate provider_config for service principals: failed to resolve workspace_id` regression ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_mws_workspaces` and `databricks_mws_credentials` data sources failing on account-level provider with `cannot populate provider_config for mws workspaces: failed to resolve workspace_id` ([#5672](https://github.com/databricks/terraform-provider-databricks/issues/5672)). These account-only data sources are now exempted from the post-Read workspace-tracking hook, and `provider_config` (which had no effect on them) is now deprecated and will be removed in a future major release.
* Fix `databricks_disable_legacy_features_setting` failing on account-level provider with `cannot populate provider_config for disable legacy features setting: failed to resolve workspace_id: ... Unable to load OAuth Config`. This account-only setting is now exempted from the post-Read workspace-tracking hook, and the auto-injected `provider_config` block is deprecated. The fix is applied at the generic-setting builder level (`makeSettingResource` in `settings/generic_setting.go`), so any future `accountSetting`-based resource inherits the opt-out automatically.

### Documentation

### Exporter

### Internal Changes
