# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

* Deprecate the SDKv2 fallback implementations of `databricks_library`, `databricks_quality_monitor`, and `databricks_share` resources, and the `databricks_share`, `databricks_shares`, and `databricks_volumes` data sources. These resources have been served by the Plugin Framework by default since their migration; the SDKv2 implementations remain only as opt-in fallbacks via the `USE_SDK_V2_RESOURCES` / `USE_SDK_V2_DATA_SOURCES` environment variables. Setting either environment variable for any of these names now emits a runtime warning (visible with `TF_LOG=WARN` or higher), and the SDKv2 implementations will be removed in the next major release of the provider.

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).
* Fix `databricks_ip_access_list` ignoring `enabled = false` on create and update ([#5829](https://github.com/databricks/terraform-provider-databricks/pull/5829))

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
