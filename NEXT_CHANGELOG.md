# NEXT CHANGELOG

## Release v1.100.0

### Breaking Changes

### New Features and Improvements

* Added `databricks_permission` resource for managing permissions on Databricks objects for individual principals ([#5161](https://github.com/databricks/terraform-provider-databricks/pull/5161)).
* Recreate `databricks_credential` on the `name` change ([#5248](https://github.com/databricks/terraform-provider-databricks/pull/5248)).
* Allow to specify default catalog and schema for `databricks_dashboard` ([#5259](https://github.com/databricks/terraform-provider-databricks/pull/5259)).

### Bug Fixes

* Fix retrieving of latest DBR versions in `databricks_spark_version` ([#5255](https://github.com/databricks/terraform-provider-databricks/pull/5255))

### Documentation

### Exporter

* Added `-targetCloud` and `-nodeTypeMappingFile` flags for cross-cloud attribute and node-type conversion ([#5236](https://github.com/databricks/terraform-provider-databricks/issues/5236)).

### Internal Changes
