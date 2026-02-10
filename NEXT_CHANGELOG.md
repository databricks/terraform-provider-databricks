# NEXT CHANGELOG

## Release v1.106.0

### Breaking Changes

### New Features and Improvements

* Add `role_arn` field to `databricks_mws_storage_configurations` resource to support sharing S3 buckets between root storage and Unity Catalog ([#5222](https://github.com/databricks/terraform-provider-databricks/issues/5222))

### Bug Fixes
* Fix the `databricks_mws_ncc_private_endpoint_rule` resource to support updates for non-S3 endpoint rules ([#5326](https://github.com/databricks/terraform-provider-databricks/pull/5326))

* Fixed `databricks_users` data source `extra_attributes` parameter issues ([#5308](https://github.com/databricks/terraform-provider-databricks/issues/5308)): (1) Single-attribute inputs (e.g., `extra_attributes = "active"`) were silently ignored at account level due to incorrect value quoting. (2) Complex multi-valued attributes like `emails` and `roles` returned null at account level even when explicitly requested in `extra_attributes`. (3) `extra_attributes` were not forwarded to the SCIM API at workspace level.
* Fix `databricks_app` resource fail to read app when deleted outside terraform ([#5365](https://github.com/databricks/terraform-provider-databricks/pull/5365))

### Documentation

* Add `run_as` configuration block documentation to `databricks_pipeline` resource ([#5342](https://github.com/databricks/terraform-provider-databricks/pull/5342))

### Exporter

### Internal Changes
