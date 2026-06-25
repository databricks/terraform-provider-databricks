# NEXT CHANGELOG

## Release v1.120.0

### Breaking Changes

### New Features and Improvements

* Added `databricks_gcp_crossaccount_policy`, `databricks_gcp_vpc_policy`, and `databricks_gcp_unity_catalog_policy` data sources to simplify creation of GCP custom IAM roles for Databricks workspaces and Unity Catalog ([#5425](https://github.com/databricks/terraform-provider-databricks/pull/5425)).

### Bug Fixes

* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5757](hattps://github.com/databricks/terraform-provider-databricks/issues/5757)).

### Documentation

### Exporter

### Internal Changes
