# NEXT CHANGELOG

## Release v1.124.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_gateway_model_service`.
* Add resource and data sources for `databricks_ai_gateway_model_provider_service`.
* Add resource and data sources for `databricks_ai_gateway_mcp_service`.

### Bug Fixes
* Use a 600 second HTTP timeout when creating a `databricks_repo`, so cloning a larger repository no longer fails with `request timed out after 1m5s of inactivity`. A higher `http_timeout_seconds` is still respected.

### Documentation

### Exporter

### Internal Changes
