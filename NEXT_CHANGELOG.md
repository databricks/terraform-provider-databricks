# NEXT CHANGELOG

## Release v1.90.0

### Breaking Changes

### New Features and Improvements
* Added the following resources and data sources: `account_settings_v2_setting`, `rfa_access_request_destinations` and `workspace_settings_v2_setting` ([#5016](https://github.com/databricks/terraform-provider-databricks/pull/5016)).

* Added `no_wait` option for `databricks_sql_endpoint` to skip waiting to start on warehouse creation

### Bug Fixes

* Remove incorrect customization for `databricks_catalog`

### Documentation

### Exporter

### Internal Changes
* Update Go SDK to v0.83.0 ([#5016](https://github.com/databricks/terraform-provider-databricks/pull/5016)).
