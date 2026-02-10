# NEXT CHANGELOG

## Release v1.106.0

### Breaking Changes

### New Features and Improvements

* Add `role_arn` field to `databricks_mws_storage_configurations` resource to support sharing S3 buckets between root storage and Unity Catalog ([#5222](https://github.com/databricks/terraform-provider-databricks/issues/5222))

### Bug Fixes

* Prevent users from locking themselves out of managing a secret scope in `databricks_secret_acl` ([#5373](https://github.com/databricks/terraform-provider-databricks/pull/5373)).

### Documentation

### Exporter

### Internal Changes
