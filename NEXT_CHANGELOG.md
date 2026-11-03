# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add `databricks_app_slack` and `databricks_app_teams` configuration blocks to `databricks_notification_destination`, and accept `DATABRICKS_APP_SLACK` and `DATABRICKS_APP_TEAMS` as `type` in the `databricks_notification_destinations` data source ([#6131](https://github.com/databricks/terraform-provider-databricks/pull/6131)).

### Bug Fixes

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

### Documentation
* Fix the VPC endpoint security group rules in the AWS Private Link guide: add the missing inbound ports (5432 and 8443-8451), source the rules from the workspace security group, and drop the unnecessary egress rules ([#5930](https://github.com/databricks/terraform-provider-databricks/pull/5930)).

### Exporter

### Internal Changes
