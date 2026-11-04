# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

* `databricks_app` now updates apps through the asynchronous Apps update API with a wildcard
  update mask ([#6151](https://github.com/databricks/terraform-provider-databricks/pull/6151)).
  Previously the synchronous update API ignored `compute_size`, `compute_min_instances`, and
  `compute_max_instances`, so changes to them never took effect. Every updatable field is now
  applied from the configuration, and fields omitted from it are cleared or reset to their
  backend default on update, matching create: an omitted `compute_size` resets to `MEDIUM`,
  omitted `compute_min_instances`/`compute_max_instances` disable horizontal scaling, and an
  omitted `git_source` clears the app's default git source. Set these fields explicitly to keep
  values configured outside Terraform. Updates now wait for the app update to reach `SUCCEEDED`.

### Breaking Changes

### New Features and Improvements

* Add `databricks_app_slack` and `databricks_app_teams` configuration blocks to `databricks_notification_destination`, and accept `DATABRICKS_APP_SLACK` and `DATABRICKS_APP_TEAMS` as `type` in the `databricks_notification_destinations` data source ([#6131](https://github.com/databricks/terraform-provider-databricks/pull/6131)).

### Bug Fixes

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

### Documentation

### Exporter

### Internal Changes
