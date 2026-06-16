# NEXT CHANGELOG

## Release v1.118.0

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.
* Add `clear_cloud_attributes_on_remove` to `databricks_cluster` ([#5812](https://github.com/databricks/terraform-provider-databricks/pull/5812)). When set to `true`, removing a cloud attributes block (`aws_attributes`, `azure_attributes`, `gcp_attributes`) from the configuration clears it instead of the removal being silently suppressed. The flag defaults to `false`, preserving the existing diff-suppression behavior that prevents perpetual drift from platform-returned cloud attribute defaults. Keeping a block, even partially specified, is unaffected; only removing the whole block clears.


### Bug Fixes

* Fixed `databricks_mws_workspaces` failing to update `private_access_settings_id` and other fields on GCP workspaces ([#5430](https://github.com/databricks/terraform-provider-databricks/issues/5430)).

### Documentation
* Added `disabled` field to `task` block in `databricks_job` resource, allowing individual tasks to be disabled ([#5767](https://github.com/databricks/terraform-provider-databricks/pull/5767)).

### Exporter

* Rewrote Exporter logging so it works with Databricks Go SDK logging ([#5805](https://github.com/databricks/terraform-provider-databricks/pull/5805)).

### Internal Changes
