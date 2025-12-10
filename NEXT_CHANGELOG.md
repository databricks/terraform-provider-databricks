# NEXT CHANGELOG

## Release v1.100.0

### Breaking Changes

### New Features and Improvements

* Handle new fields in `databricks_pipeline` resource ([#5249](https://github.com/databricks/terraform-provider-databricks/pull/5249))
* Recreate `databricks_credential` on the `name` change ([#5248](https://github.com/databricks/terraform-provider-databricks/pull/5248)).
* Allow to specify default catalog and schema for `databricks_dashboard` ([#5259](https://github.com/databricks/terraform-provider-databricks/pull/5259)).
* Added `databricks_pipeline` data source for single pipeline lookup ([#5271](https://github.com/databricks/terraform-provider-databricks/pull/5271))


### Bug Fixes

* Fix retrieving of latest DBR versions in `databricks_spark_version` ([#5255](https://github.com/databricks/terraform-provider-databricks/pull/5255))
* Reset PO flag for non-managed UC Catalogs ([#5260](https://github.com/databricks/terraform-provider-databricks/pull/5260)).

### Documentation

* Add missing GCP IAM permission `topics.detachSubscription` for Databricks file events ([#5269](https://github.com/databricks/terraform-provider-databricks/pull/5269))

### Exporter

* Added `-targetCloud` and `-nodeTypeMappingFile` flags for cross-cloud attribute and node-type conversion ([#5236](https://github.com/databricks/terraform-provider-databricks/issues/5236)).

### Internal Changes
