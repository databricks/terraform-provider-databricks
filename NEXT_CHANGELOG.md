# NEXT CHANGELOG

## Release v1.78.0

### New Features and Improvements

 * Faster and more reliable schema deletion. It now uses schemas/delete call with force=true flag instead of manually listing and deleting all resources.  [#4705](https://github.com/databricks/terraform-provider-databricks/pull/4705)

### Bug Fixes

 * Fix validation of S3 bucket name in `databricks_aws_unity_catalog_policy` and `databricks_aws_bucket_policy` [#4691](https://github.com/databricks/terraform-provider-databricks/pull/4691)

### Documentation

* Fix import documentation for all resources ([#4699](https://github.com/databricks/terraform-provider-databricks/pull/4699/files)).

### Exporter

### Internal Changes
