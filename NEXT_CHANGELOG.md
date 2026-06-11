# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements

* Deprecate the SDKv2 fallback implementations of `databricks_library`, `databricks_quality_monitor`, and `databricks_share` resources, and the `databricks_share`, `databricks_shares`, and `databricks_volumes` data sources. These resources have been served by the Plugin Framework by default since their migration; the SDKv2 implementations remain only as opt-in fallbacks via the `USE_SDK_V2_RESOURCES` / `USE_SDK_V2_DATA_SOURCES` environment variables. Setting either environment variable for any of these names now emits a runtime warning (visible with `TF_LOG=WARN` or higher), and the SDKv2 implementations will be removed in the next major release of the provider.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
