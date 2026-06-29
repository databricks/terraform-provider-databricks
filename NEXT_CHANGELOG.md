# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

* Remove the `USE_SDK_V2_RESOURCES` and `USE_SDK_V2_DATA_SOURCES` environment variables. Setting either no longer has any effect; the Plugin Framework is now the only path for `databricks_library`, `databricks_quality_monitor`, and `databricks_share` resources and the `databricks_share`, `databricks_shares`, and `databricks_volumes` data sources. Users who previously relied on these env vars as a fallback should remove them; any existing Terraform state written against the SDKv2 implementations continues to be readable by the Plugin Framework versions.

### New Features and Improvements

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).
* Allow setting `user_api_scopes = []` on `databricks_app` to disable OBO (On-Behalf-Of) user authorization ([#5834](https://github.com/databricks/terraform-provider-databricks/pull/5834)).

  The Apps API omits `user_api_scopes` from its response when OBO is inactive, so a configured empty list previously failed with `Provider produced inconsistent result after apply`. The provider now preserves a configured empty list in state, mirroring the reconciliation used by `databricks_app_space`.

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
