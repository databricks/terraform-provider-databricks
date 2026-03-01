# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fixed `databricks_sql_endpoint` resource failing when `min_num_clusters` was changed outside of Terraform by adding `Default: 1` to match `max_num_clusters` behavior ([#5294](https://github.com/databricks/terraform-provider-databricks/issues/5294)).

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes
