# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add `databricks_app_slack` and `databricks_app_teams` configuration blocks to `databricks_notification_destination`, and accept `DATABRICKS_APP_SLACK` and `DATABRICKS_APP_TEAMS` as `type` in the `databricks_notification_destinations` data source ([#6131](https://github.com/databricks/terraform-provider-databricks/pull/6131)).

### Bug Fixes

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

### Documentation

* Document that Databricks plans to deprecate classic AWS workspace creation with a Databricks-managed VPC in `databricks_mws_workspaces`, `databricks_mws_credentials`, and `databricks_aws_crossaccount_policy`, and use `policy_type = "customer"` in examples.

### Exporter

### Internal Changes
