# NEXT CHANGELOG

## Release v1.138.0

### Important Changes

### Breaking Changes

* Remove the `databricks_mason_managed_memory_entry`, `databricks_mason_managed_memory_store`, `databricks_mason_session`, and `databricks_mason_session_store` resources and their data sources ([#6087](https://github.com/databricks/terraform-provider-databricks/pull/6087)).

  The Databricks Go SDK renamed the Mason API to AgentKit.

* Remove the `databricks_feature_engineering_kafka_config` and `databricks_feature_engineering_materialized_feature` resources and their data sources ([#6087](https://github.com/databricks/terraform-provider-databricks/pull/6087)).

  The API no longer exposes these resources to Terraform.

### New Features and Improvements

### Bug Fixes

* Require replacement of `databricks_budget` when its filter changes.

  Replacement creates a new budget; accumulated usage and alert behavior should be reviewed before applying.

### Documentation

### Exporter

### Internal Changes
