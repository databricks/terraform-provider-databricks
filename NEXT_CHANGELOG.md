# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements
* Add resource and data source for `databricks_postgres_data_api`.

* Deprecate the SDKv2 fallback implementations of `databricks_library`, `databricks_quality_monitor`, and `databricks_share` resources, and the `databricks_share`, `databricks_shares`, and `databricks_volumes` data sources. These resources have been served by the Plugin Framework by default since their migration; the SDKv2 implementations remain only as opt-in fallbacks via the `USE_SDK_V2_RESOURCES` / `USE_SDK_V2_DATA_SOURCES` environment variables. Setting either environment variable for any of these names now emits a runtime warning (visible with `TF_LOG=WARN` or higher), and the SDKv2 implementations will be removed in the next major release of the provider.

### Bug Fixes

* Fix intermittent inconsistent final plan errors in `databricks_permissions` for SQL warehouse access control reads after permission updates ([#5837](https://github.com/databricks/terraform-provider-databricks/issues/5837)).
* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).
* Allow setting `user_api_scopes = []` on `databricks_app` to disable OBO (On-Behalf-Of) user authorization ([#5834](https://github.com/databricks/terraform-provider-databricks/pull/5834)).

  The Apps API omits `user_api_scopes` from its response when OBO is inactive, so a configured empty list previously failed with `Provider produced inconsistent result after apply`. The provider now preserves a configured empty list in state, mirroring the reconciliation used by `databricks_app_space`.

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
