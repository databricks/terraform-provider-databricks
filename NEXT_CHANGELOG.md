# NEXT CHANGELOG

## Release v1.106.0

### Breaking Changes

### New Features and Improvements

* Add `role_arn` field to `databricks_mws_storage_configurations` resource to support sharing S3 buckets between root storage and Unity Catalog ([#5222](https://github.com/databricks/terraform-provider-databricks/issues/5222))

### Bug Fixes

* Fixed `databricks_users` data source `extra_attributes` parameter issues ([#5308](https://github.com/databricks/terraform-provider-databricks/issues/5308)): (1) Single-attribute inputs (e.g., `extra_attributes = "active"`) were silently ignored at account level due to incorrect value quoting. (2) Complex multi-valued attributes like `emails` and `roles` returned null at account level even when explicitly requested in `extra_attributes`. (3) `extra_attributes` were not forwarded to the SCIM API at workspace level.

### Documentation

### Exporter

### Internal Changes
