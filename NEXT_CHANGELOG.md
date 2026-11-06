# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_datasource`.
* Add resource and data source for `databricks_lakewatch_datasource_config`.
* Add resource and data source for `databricks_lakewatch_rule_config`.
* Add resource and data source for `databricks_lakewatch_alert_config`.
* Add resource and data source for `databricks_lakewatch_common_config`.
* Add resource and data sources for `databricks_lakewatch_rule`.

* Add `databricks_app_slack` and `databricks_app_teams` configuration blocks to `databricks_notification_destination`, and accept `DATABRICKS_APP_SLACK` and `DATABRICKS_APP_TEAMS` as `type` in the `databricks_notification_destinations` data source ([#6131](https://github.com/databricks/terraform-provider-databricks/pull/6131)).

* Add `databricks_genie_space` resource and `databricks_genie_spaces` data source, plus Genie Space support in `databricks_permissions` via the new `genie_space_id` attribute ([#5770](https://github.com/databricks/terraform-provider-databricks/pull/5770)).

  The resource normalizes `serialized_space` to suppress whitespace and key-order diffs and auto-creates a missing `parent_path` on first apply. Delete is trash-aware (treats an already-trashed space as a successful delete). Tags can be attached using the existing `databricks_workspace_entity_tag_assignment` resource with `entity_type = "geniespaces"`.

### Bug Fixes

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

* Preserve `databricks_pipeline.expected_last_modified` during refresh to avoid post-apply drift ([#6133](https://github.com/databricks/terraform-provider-databricks/pull/6133)).

### Documentation

### Exporter

### Internal Changes
