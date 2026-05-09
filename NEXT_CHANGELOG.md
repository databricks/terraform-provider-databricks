# NEXT CHANGELOG

## Release v1.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_library`, `databricks_share`, and `databricks_quality_monitor` failing to decode prior state with `Error decoding ... from prior state: missing expected {` after upgrading from v1.113.0 to v1.114.0 ([#5669](https://github.com/databricks/terraform-provider-databricks/issues/5669)). Reverts [#5582](https://github.com/databricks/terraform-provider-databricks/pull/5582).

  > **Note for users upgrading from v1.114.0**: any state written by v1.114.0 for `databricks_library`, `databricks_share`, or `databricks_quality_monitor` encodes `provider_config` as a single object instead of a list, and `terraform plan` against the upgraded provider will fail to decode it. Mitigate with a one-time edit of each affected resource instance in your state file: change the `provider_config` value from the object form to either the list form or null.
  >
  >   - If you set `provider_config` explicitly in HCL: change `"provider_config": {"workspace_id": "X"}` to `"provider_config": [{"workspace_id": "X"}]` (wrap the existing object in a single-element list).
  >   - If you did NOT set `provider_config` in HCL: change `"provider_config": {"workspace_id": "X"}` to `"provider_config": null`. This avoids a one-time replacement plan on `databricks_library` (where the block-level plan modifier forces replacement on any provider_config diff).
  >
  > Users on v1.113.0 are unaffected — their state already matches the restored schema.

* Fix `databricks_service_principal` data source failing on account-level provider with `cannot populate provider_config for service principal: failed to resolve workspace_id` ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_service_principals` data source failing on account-level provider with the same `cannot populate provider_config for service principals: failed to resolve workspace_id` regression ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_mws_workspaces` and `databricks_mws_credentials` data sources failing on account-level provider with `cannot populate provider_config for mws workspaces: failed to resolve workspace_id` ([#5672](https://github.com/databricks/terraform-provider-databricks/issues/5672)). These account-only data sources are now exempted from the post-Read workspace-tracking hook, and `provider_config` (which had no effect on them) is now deprecated and will be removed in a future major release.

### Documentation

### Exporter

### Internal Changes
