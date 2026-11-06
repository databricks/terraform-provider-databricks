# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

### Breaking Changes

* Require exactly one nonempty GCP target in `databricks_mws_ncc_private_endpoint_rule`. Omit unused targets instead of setting `all_vpc_sc_services = false` or `service_attachment = ""`; `google_api_endpoints.endpoints` must be nonempty ([#5973](https://github.com/databricks/terraform-provider-databricks/pull/5973)).

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_datasource`.
* Add resource and data source for `databricks_lakewatch_datasource_config`.
* Add resource and data source for `databricks_lakewatch_rule_config`.
* Add resource and data source for `databricks_lakewatch_alert_config`.
* Add resource and data source for `databricks_lakewatch_common_config`.
* Add resource and data sources for `databricks_lakewatch_rule`.

* Add `databricks_app_slack` and `databricks_app_teams` configuration blocks to `databricks_notification_destination`, and accept `DATABRICKS_APP_SLACK` and `DATABRICKS_APP_TEAMS` as `type` in the `databricks_notification_destinations` data source ([#6131](https://github.com/databricks/terraform-provider-databricks/pull/6131)).

* Add GCP targets to `databricks_mws_ncc_private_endpoint_rule` ([#5973](https://github.com/databricks/terraform-provider-databricks/pull/5973)).

### Bug Fixes

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

* Preserve `databricks_pipeline.expected_last_modified` during refresh to avoid post-apply drift ([#6133](https://github.com/databricks/terraform-provider-databricks/pull/6133)).

### Documentation

### Exporter

### Internal Changes
