# NEXT CHANGELOG

## Release v1.121.0

### Breaking Changes

### New Features and Improvements

* Add resource and data sources for `databricks_ai_search_endpoint`.
* Add resource and data sources for `databricks_ai_search_index`.
* Add `clear_cloud_attributes_on_remove` to `databricks_cluster` ([#5812](https://github.com/databricks/terraform-provider-databricks/pull/5812)). When set to `true`, removing a cloud attributes block (`aws_attributes`, `azure_attributes`, `gcp_attributes`) from the configuration clears it instead of the removal being silently suppressed. The flag defaults to `false`, preserving the existing diff-suppression behavior that prevents perpetual drift from platform-returned cloud attribute defaults. Keeping a block, even partially specified, is unaffected; only removing the whole block clears.

### Bug Fixes

* Fix import for jobs with >100 tasks ([#5417](https://github.com/databricks/terraform-provider-databricks/pull/5417)).

### Documentation

* Added an example to `databricks_budget` for creating budgets to control Genie usage costs ([#5817](https://github.com/databricks/terraform-provider-databricks/pull/5817)).

### Exporter

* Generate code in `import.sh` more safely ([#5848](hattps://github.com/databricks/terraform-provider-databricks/issues/5848)).

### Internal Changes

* Deprecate `instance_profiles` attribute, and replace it with `roles` in `databricks_group` data source ([#5086](https://github.com/databricks/terraform-provider-databricks/pull/5086))
