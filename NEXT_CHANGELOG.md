# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.


### Bug Fixes

* Fix `databricks_global_init_script` silently ignoring `enabled = false` and `position = 0` on update ([#5718](https://github.com/databricks/terraform-provider-databricks/issues/5718)).

### Documentation
* Added `disabled` field to `task` block in `databricks_job` resource, allowing individual tasks to be disabled ([#5767](https://github.com/databricks/terraform-provider-databricks/pull/5767)).

### Exporter

### Internal Changes
