# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

* Remove the `USE_SDK_V2_RESOURCES` / `USE_SDK_V2_DATA_SOURCES` environment variables and the corresponding SDKv2 implementations of `databricks_library`, `databricks_share` resources and `databricks_share`, `databricks_shares`, `databricks_volumes` data sources. The Plugin Framework is now the only path for these surfaces; the env vars no longer have any effect. Existing Terraform state written against the SDKv2 implementations continues to be readable by the Plugin Framework versions through the `ConfigureAsSdkV2Compatible()` schema shim, which is now covered by the `TestResource*_SchemaPreserved` unit tests in place of the previous transition acceptance tests. The SDKv2 implementation of `databricks_quality_monitor` is retained as it backs the still-supported (but deprecated) `databricks_lakehouse_monitor` alias.

### New Features and Improvements

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).
* Allow setting `user_api_scopes = []` on `databricks_app` to disable OBO (On-Behalf-Of) user authorization ([#5834](https://github.com/databricks/terraform-provider-databricks/pull/5834)).

  The Apps API omits `user_api_scopes` from its response when OBO is inactive, so a configured empty list previously failed with `Provider produced inconsistent result after apply`. The provider now preserves a configured empty list in state, mirroring the reconciliation used by `databricks_app_space`.

### Documentation

### Exporter

### Internal Changes

* Make notification destination acceptance tests robust to the eventual consistency of the notification destinations list API.
