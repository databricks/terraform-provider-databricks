# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

* Added `databricks_gcp_crossaccount_policy`, `databricks_gcp_vpc_policy`, and `databricks_gcp_unity_catalog_policy` data sources to simplify creation of GCP custom IAM roles for Databricks workspaces and Unity Catalog ([#5425](https://github.com/databricks/terraform-provider-databricks/pull/5425)).

### Bug Fixes

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Host-agnostic cloud detection via node type patterns, replacing host-URL-based `IsAws()`/`IsAzure()`/`IsGcp()` checks.
