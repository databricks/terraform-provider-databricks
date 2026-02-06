# NEXT CHANGELOG

## Release v1.106.0

### Breaking Changes

### New Features and Improvements
* Add polling on `databricks_resource_mws_ncc_private_endpoint_rule` so that Terraform also behaves synchronously, even with an async API

* Add `role_arn` field to `databricks_mws_storage_configurations` resource to support sharing S3 buckets between root storage and Unity Catalog ([#5222](https://github.com/databricks/terraform-provider-databricks/issues/5222))

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
