# NEXT CHANGELOG

## Release v1.124.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_gateway_model_service`.
* Add resource and data sources for `databricks_ai_gateway_model_provider_service`.
* Add resource and data sources for `databricks_ai_gateway_mcp_service`.

### Bug Fixes
* Fixed `databricks_share` failing with `produced an unexpected new value: .comment` when a share's comment is set or removed outside of Terraform (e.g. in the UI). The `comment` attribute is now `Optional`+`Computed`, so an out-of-band value is adopted into state instead of erroring, while `comment = ""` still explicitly clears the description.

### Documentation

### Exporter

### Internal Changes
