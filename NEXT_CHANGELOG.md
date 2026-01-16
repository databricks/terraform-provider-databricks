# NEXT CHANGELOG

## Release v1.105.0

### Breaking Changes

* Return empty string for `data_source_id` in `databricks_sql_warehouse` and `databricks_sql_endpoint` if data source API call fails ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### New Features and Improvements

### Bug Fixes
* Fix the `databricks_mws_ncc_private_endpoint_rule` resource to support updates for non-S3 endpoint rules ([#5326](https://github.com/databricks/terraform-provider-databricks/pull/5326))

### Documentation

* Mark `data_source_id` as deprecated in `databricks_sql_warehouse` and `databricks_sql_endpoint` ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### Exporter

### Internal Changes
