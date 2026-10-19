# NEXT CHANGELOG

## Release v1.139.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* [Breaking] Remove resource and data sources for `databricks_mason_managed_memory_store`.
* [Breaking] Remove resource and data sources for `databricks_mason_managed_memory_entry`.
* [Breaking] Remove resource and data sources for `databricks_mason_session_store`.
* [Breaking] Remove resource and data sources for `databricks_mason_session`.
* Add resource and data sources for `databricks_agent_kit_managed_memory_store`.
* Add resource and data sources for `databricks_agent_kit_managed_memory_entry`.
* Add resource and data sources for `databricks_agent_kit_session_store`.
* Add resource and data sources for `databricks_agent_kit_session`.
* [Breaking] Remove resource and data sources for `databricks_feature_engineering_materialized_feature`.
* [Breaking] Remove resource and data sources for `databricks_feature_engineering_kafka_config`.

* Support fixed-size REALTIME warehouses with `max_dbu`, `statement_timeout`, and `customer_version_label` in `databricks_sql_endpoint` and `databricks_sql_warehouse` ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Bug Fixes

### Documentation

### Exporter

* Preserve `auto_stop_mins`, REALTIME `statement_timeout`, and `max_dbu` in warehouse exports ([#6092](https://github.com/databricks/terraform-provider-databricks/pull/6092)).

### Internal Changes
