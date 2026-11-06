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
* Fix `databricks_cluster` dropping a configured `autoscale` block when resizing a running cluster ([#6018](https://github.com/databricks/terraform-provider-databricks/pull/6018)).

  When `num_workers` changed in the plan while the `autoscale` block stayed the same, the provider called `POST /api/2.1/clusters/resize` with `num_workers` and `ForceSendFields: ["NumWorkers"]`, so the API discarded autoscaling and pinned the running cluster to the fixed size - `0` workers when the configuration declares only `autoscale`. The apply reported success. The four predicates that selected the resize path decided from what had changed in the diff rather than from the desired configuration: `isNumWorkersResizeForNonAutoscalingCluster` never checked `cluster.Autoscale == nil` despite its name. They are replaced by a single condition that builds the resize request from the desired configuration, so an `autoscale` block in the configuration always takes precedence over `num_workers`. This also fixes removing `autoscale` while `num_workers` is unchanged, which previously issued a resize request carrying neither a size nor a range.

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

* Preserve `databricks_pipeline.expected_last_modified` during refresh to avoid post-apply drift ([#6133](https://github.com/databricks/terraform-provider-databricks/pull/6133)).

### Documentation

### Exporter

### Internal Changes
