# NEXT CHANGELOG

## Release v1.138.0

### Important Changes

### Breaking Changes

* Remove the `databricks_mason_managed_memory_entry`, `databricks_mason_managed_memory_store`, `databricks_mason_session`, and `databricks_mason_session_store` resources and their data sources ([#6087](https://github.com/databricks/terraform-provider-databricks/pull/6087)).

  The Databricks Go SDK renamed the Mason API to AgentKit.

* Remove the `databricks_feature_engineering_kafka_config` and `databricks_feature_engineering_materialized_feature` resources and their data sources ([#6087](https://github.com/databricks/terraform-provider-databricks/pull/6087)).

  The API no longer exposes these resources to Terraform.

### New Features and Improvements

* Add `scopes` and `autoscope_enabled` arguments to `databricks_obo_token` and `databricks_token` resources ([#6046](https://github.com/databricks/terraform-provider-databricks/issues/6046)).

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
