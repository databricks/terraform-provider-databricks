# NEXT CHANGELOG

## Release v1.110.0

### Breaking Changes

### New Features and Improvements

* Changed default AWS availability for auto-created utility clusters from `SPOT` to `SPOT_WITH_FALLBACK` (API default). `SPOT_WITH_FALLBACK` improves reliability by falling back to on-demand instances when spot capacity is unavailable. Affects internal clusters created by `databricks_aws_s3_mount`, `databricks_mount`, `databricks_sql_permissions`, `databricks_sql_table`, and the exporter.

### Bug Fixes

* Fixed `databricks_sql_table` timing out after 50 seconds when creating tables with a custom `warehouse_id` by polling for statement completion instead of cancelling ([#5340](https://github.com/databricks/terraform-provider-databricks/issues/5340))

### Documentation

### Exporter

### Internal Changes
