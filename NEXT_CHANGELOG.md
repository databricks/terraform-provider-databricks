# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

* Deprecate the SDKv2 fallback implementations of `databricks_library`, `databricks_quality_monitor`, and `databricks_share` resources, and the `databricks_share`, `databricks_shares`, and `databricks_volumes` data sources. These resources have been served by the Plugin Framework by default since their migration; the SDKv2 implementations remain only as opt-in fallbacks via the `USE_SDK_V2_RESOURCES` / `USE_SDK_V2_DATA_SOURCES` environment variables. Setting either environment variable for any of these names now emits a runtime warning (visible with `TF_LOG=WARN` or higher), and the SDKv2 implementations will be removed in the next major release of the provider.
* Reintroduce provider-level `workspace_id` support on `databricks_library`, `databricks_quality_monitor`, and `databricks_share`. `provider_config` is now a single nested attribute (`provider_config = { workspace_id = "..." }`) matching the shape used by `databricks_app`, `databricks_tag_policy`, and other Plugin Framework resources, and the inner `workspace_id` is Optional+Computed so the provider populates it from the provider-level `workspace_id` / host metadata when omitted. A schema v0→v1 state upgrader migrates state files written by earlier releases (where `provider_config` was a list-shaped block) so existing users upgrade without "Error decoding ... missing expected {" — the failure mode that prompted the earlier revert in #5685.

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).
* Allow setting `user_api_scopes = []` on `databricks_app` to disable OBO (On-Behalf-Of) user authorization ([#5834](https://github.com/databricks/terraform-provider-databricks/pull/5834)).

  The Apps API omits `user_api_scopes` from its response when OBO is inactive, so a configured empty list previously failed with `Provider produced inconsistent result after apply`. The provider now preserves a configured empty list in state, mirroring the reconciliation used by `databricks_app_space`.

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
