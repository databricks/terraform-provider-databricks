# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.


### Bug Fixes

* Fix `databricks_pipeline` so that `photon`, `serverless`, `continuous` and `development` set to `false` are sent in the create/update request if specified ([#5806](https://github.com/databricks/terraform-provider-databricks/pull/5806)).

### Documentation
* Added `disabled` field to `task` block in `databricks_job` resource, allowing individual tasks to be disabled ([#5767](https://github.com/databricks/terraform-provider-databricks/pull/5767)).

### Exporter

### Internal Changes
