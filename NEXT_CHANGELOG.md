# NEXT CHANGELOG

## Release v1.108.0

### Breaking Changes

### New Features and Improvements

* Changed default AWS availability for auto-created utility clusters from `SPOT` to `SPOT_WITH_FALLBACK` (API default). `SPOT_WITH_FALLBACK` improves reliability by falling back to on-demand instances when spot capacity is unavailable. Affects internal clusters created by `databricks_aws_s3_mount`, `databricks_mount`, `databricks_sql_permissions`, `databricks_sql_table`, and the exporter.
* Added `node_type_flexibility` block to `databricks_instance_pool` resource ([#5381](https://github.com/databricks/terraform-provider-databricks/pull/5381)).

### Bug Fixes

* Handle error during WorkspaceClient() creation in databricks_grant and databricks_grants resources ([#5403](https://github.com/databricks/terraform-provider-databricks/pull/5403))

### Documentation

* Added note to `databricks_mws_ncc_binding` that a workspace can only have one NCC binding at a time.

### Exporter

### Internal Changes
