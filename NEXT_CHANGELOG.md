# NEXT CHANGELOG

## Release v1.105.0

### Breaking Changes

* Return empty string for `data_source_id` in `databricks_sql_warehouse` and `databricks_sql_endpoint` if data source API call fails ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### New Features and Improvements

### Bug Fixes

### Documentation

* Mark `data_source_id` as deprecated in `databricks_sql_warehouse` and `databricks_sql_endpoint` ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### Exporter

### Internal Changes

* Switch to use Go SDK struct in `databricks_metastore` resource ([#5088](https://github.com/databricks/terraform-provider-databricks/pull/5088))
* Change default value for `delta_sharing_recipient_token_lifetime_in_seconds` from 0 to 31536000 (1 year) ([#5296](https://github.com/databricks/terraform-provider-databricks/pull/5296))
