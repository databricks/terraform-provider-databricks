# NEXT CHANGELOG

## Release v1.110.0

### Breaking Changes

### New Features and Improvements

* Changed default AWS availability for auto-created utility clusters from `SPOT` to `SPOT_WITH_FALLBACK` (API default). `SPOT_WITH_FALLBACK` improves reliability by falling back to on-demand instances when spot capacity is unavailable. Affects internal clusters created by `databricks_aws_s3_mount`, `databricks_mount`, `databricks_sql_permissions`, `databricks_sql_table`, and the exporter.

### Bug Fixes

* Mark plaintext credential fields in `databricks_model_serving` as sensitive to prevent them from being displayed in plan/apply output ([#5409](https://github.com/databricks/terraform-provider-databricks/pull/5409)).

### Documentation

### Exporter

### Internal Changes

* Add support for host agnostic SQL global config resource.
