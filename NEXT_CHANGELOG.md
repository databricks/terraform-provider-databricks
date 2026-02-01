# NEXT CHANGELOG

## Release v1.106.0

### Breaking Changes

### New Features and Improvements

* Add `role_arn` field to `databricks_mws_storage_configurations` resource to support sharing S3 buckets between root storage and Unity Catalog ([#5222](https://github.com/databricks/terraform-provider-databricks/issues/5222))

### Bug Fixes

* Fixed `databricks_users` data source incorrectly quoting `extra_attributes` value, causing single-attribute inputs to be silently ignored at account level. Also fixed `extra_attributes` not being forwarded to the SCIM API at workspace level.

### Documentation

### Exporter

### Internal Changes
