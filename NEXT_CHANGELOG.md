# NEXT CHANGELOG

## Release v1.90.0

### Breaking Changes

### New Features and Improvements

* Added `no_wait` option for `databricks_sql_endpoint` to skip waiting to start on warehouse creation ([#5014](https://github.com/databricks/terraform-provider-databricks/pull/5014))

### Bug Fixes

* Remove incorrect customization for `databricks_catalog` ([#5021](https://github.com/databricks/terraform-provider-databricks/pull/5021))
* Fix filling of `active` attribute in `databricks_user` data source ([#5026](https://github.com/databricks/terraform-provider-databricks/pull/5026))
* Add support for share resource in plugin framework implementation to be SDKv2 compatible ([#4965](https://github.com/databricks/terraform-provider-databricks/pull/4965))

### Documentation

### Exporter

### Internal Changes
* Use `Jobs.Get` instead of `JobsGetByJobId` ([#5029](https://github.com/databricks/terraform-provider-databricks/pull/5029))
