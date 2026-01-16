# NEXT CHANGELOG

## Release v1.103.0

### Breaking Changes

### New Features and Improvements
* Add polling on `databricks_resource_mws_ncc_private_endpoint_rule` so that Terraform also behaves synchronously, even with an async API

### Bug Fixes
* Fixed `databricks_dashboard` resource to detect content changes when using the `file_path` attribute. Previously, only changes to the path string itself triggered updates, not changes to the file content.

### Documentation

### Exporter

### Internal Changes
