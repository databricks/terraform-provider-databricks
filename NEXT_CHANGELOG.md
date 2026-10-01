# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes

### Documentation

### Exporter

* Add support for exporting dependencies of the `ai_runtime_task`, `clean_rooms_notebook_task`, `dbt_cloud_task`, `dbt_platform_task`, `gen_ai_compute_task`, and `python_operator_task` job task types. ([#6027](https://github.com/databricks/terraform-provider-databricks/pull/6027)).
* Fixed a nil pointer panic in `emitRfaAccessRequestDestinations` when exporting workspace-level UC securables with an account-level provider ([#6028](https://github.com/databricks/terraform-provider-databricks/issues/6028)).

### Internal Changes
