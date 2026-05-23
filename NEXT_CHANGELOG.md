# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Migrate `databricks_mws_ncc_private_endpoint_rule` to the Plugin Framework ([#XXXX](https://github.com/databricks/terraform-provider-databricks/pull/XXXX)). No HCL or state changes; the legacy SDKv2 implementation remains in the repo and can be re-selected per-resource via `USE_SDK_V2_RESOURCES=databricks_mws_ncc_private_endpoint_rule`.
