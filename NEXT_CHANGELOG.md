# NEXT CHANGELOG

## Release v1.108.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix the `databricks_mws_ncc_private_endpoint_rule` resource to support updates for non-S3 endpoint rules ([#5326](https://github.com/databricks/terraform-provider-databricks/pull/5326))
* Prevent users from locking themselves out of managing a secret scope in `databricks_secret_acl` ([#5373](https://github.com/databricks/terraform-provider-databricks/pull/5373)).
* [Fix] `databricks_app` resource fail to read app when deleted outside terraform ([#5365](https://github.com/databricks/terraform-provider-databricks/pull/5365))
* Fixed `databricks_users` data source `extra_attributes` parameter issues ([#5308](https://github.com/databricks/terraform-provider-databricks/issues/5308)): (1) Single-attribute inputs (e.g., `extra_attributes = "active"`) were silently ignored at account level due to incorrect value quoting. (2) Complex multi-valued attributes like `emails` and `roles` returned null at account level even when explicitly requested in `extra_attributes`. (3) `extra_attributes` were not forwarded to the SCIM API at workspace level.
* Fix `databricks_app` resource fail to read app when deleted outside terraform ([#5365](https://github.com/databricks/terraform-provider-databricks/pull/5365))
* Fix num_workers Default With Policy Defaults ([#5380](https://github.com/databricks/terraform-provider-databricks/pull/5380))

### Documentation

* Added note to `databricks_mws_ncc_binding` that a workspace can only have one NCC binding at a time.

### Exporter

### Internal Changes
