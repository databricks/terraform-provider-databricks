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

### Bug Fixes
* Default to a 600 second HTTP timeout for `databricks_repo` operations that run git commands inline (the clone on create and the branch/tag checkout on update), so larger repositories no longer fail with `request timed out after 1m5s of inactivity`. An explicitly configured `http_timeout_seconds` still takes precedence.
* Prevent repeated diffs for `databricks_mws_private_access_settings.allowed_vpc_endpoint_ids` when the Accounts API returns the same IDs in a different order ([#5970](https://github.com/databricks/terraform-provider-databricks/pull/5970)).
* `databricks_app` can now manage `git_source`, `source_code_path`, and `git_repository.caller_credential_id`. These are input_only (the Apps API accepts them on write but does not echo them on read), so the resource now calls the generated `SyncFields` reconciliation to preserve the configured value across reads, and treats the nested `git_source` descendants (`git_repository`, `resolved_commit`) as non-Computed. This prevents the "Provider produced inconsistent result after apply" error and the perpetual diff that previously occurred when any of these fields was set.

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

* Preserve `databricks_pipeline.expected_last_modified` during refresh to avoid post-apply drift ([#6133](https://github.com/databricks/terraform-provider-databricks/pull/6133)).

### Documentation

### Exporter

### Internal Changes
