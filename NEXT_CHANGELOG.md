# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.


### Bug Fixes

* Fix `databricks_global_init_script` silently ignoring `enabled = false` and `position = 0` on update ([#5718](https://github.com/databricks/terraform-provider-databricks/issues/5718)).
* Fixed `databricks_mws_workspaces` failing to update `private_access_settings_id` and other fields on GCP workspaces ([#5430](https://github.com/databricks/terraform-provider-databricks/issues/5430)).

### Documentation
* Added `disabled` field to `task` block in `databricks_job` resource, allowing individual tasks to be disabled ([#5767](https://github.com/databricks/terraform-provider-databricks/pull/5767)).

### Exporter

* Rewrote Exporter logging so it works with Databricks Go SDK logging ([#5805](https://github.com/databricks/terraform-provider-databricks/pull/5805)).

### Internal Changes
