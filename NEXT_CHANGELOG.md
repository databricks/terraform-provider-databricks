# NEXT CHANGELOG

## Release v1.106.0

### Breaking Changes

### New Features and Improvements

* Support for payload bigger than 10Mb in `databricks_workspace_file` ([#5293](https://github.com/databricks/terraform-provider-databricks/pull/5293))
* Add `role_arn` field to `databricks_mws_storage_configurations` resource to support sharing S3 buckets between root storage and Unity Catalog ([#5222](https://github.com/databricks/terraform-provider-databricks/issues/5222))

### Bug Fixes

* Fix the `databricks_mws_ncc_private_endpoint_rule` resource to support updates for non-S3 endpoint rules ([#5326](https://github.com/databricks/terraform-provider-databricks/pull/5326))
* Prevent users from locking themselves out of managing a secret scope in `databricks_secret_acl` ([#5373](https://github.com/databricks/terraform-provider-databricks/pull/5373)).
* [Fix] `databricks_app` resource fail to read app when deleted outside terraform ([#5365](https://github.com/databricks/terraform-provider-databricks/pull/5365))
* Fixed `databricks_users` data source `extra_attributes` parameter issues ([#5308](https://github.com/databricks/terraform-provider-databricks/issues/5308)): (1) Single-attribute inputs (e.g., `extra_attributes = "active"`) were silently ignored at account level due to incorrect value quoting. (2) Complex multi-valued attributes like `emails` and `roles` returned null at account level even when explicitly requested in `extra_attributes`. (3) `extra_attributes` were not forwarded to the SCIM API at workspace level.
* Fix `databricks_app` resource fail to read app when deleted outside terraform ([#5365](https://github.com/databricks/terraform-provider-databricks/pull/5365))

### Documentation

* Add `run_as` configuration block documentation to `databricks_pipeline` resource ([#5342](https://github.com/databricks/terraform-provider-databricks/pull/5342))
* Document `driver_node_type_flexibility` and `worker_node_type_flexibility` in `databricks_cluster` ([#5379](https://github.com/databricks/terraform-provider-databricks/pull/5379))

### Exporter

### Internal Changes

* Change default value for `delta_sharing_recipient_token_lifetime_in_seconds` from 0 to 31536000 (1 year) ([#5296](https://github.com/databricks/terraform-provider-databricks/pull/5296))
* Update Go SDK to v0.106.0.
