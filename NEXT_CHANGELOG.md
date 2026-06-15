# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.


### Bug Fixes

### Documentation
* Added `disabled` field to `task` block in `databricks_job` resource, allowing individual tasks to be disabled ([#5767](https://github.com/databricks/terraform-provider-databricks/pull/5767)).

### Exporter

* Rewrote Exporter logging so it works with Databricks Go SDK logging ([#5805](https://github.com/databricks/terraform-provider-databricks/pull/5805)).

### Internal Changes

* Use Go SDK for `databricks_mws_log_delivery` ([#5807](https://github.com/databricks/terraform-provider-databricks/pull/5807)).
