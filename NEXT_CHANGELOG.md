# NEXT CHANGELOG

## Release v1.88.0

### Breaking Changes

### New Features and Improvements

* Document and handle additional Slack options in `databricks_notification_destination` ([#4933](https://github.com/databricks/terraform-provider-databricks/pull/4933))

### Bug Fixes

### Documentation
* Improve documentation for grant resource ([#4906](https://github.com/databricks/terraform-provider-databricks/pull/4935))

* Document `gcp_attributes.first_on_demand` attribute in `databricks_cluster` ([#4934](https://github.com/databricks/terraform-provider-databricks/pull/4934))

* Update `databricks_job` example to use `environment_version` ([#4942](https://github.com/databricks/terraform-provider-databricks/pull/4942))

### Exporter

* Improve handling of new dependencies in jobs, pipelines, model serving ([#4914](https://github.com/databricks/terraform-provider-databricks/pull/4914))

### Internal Changes

* Replaced `common.APIErrorBody` with corresponding structs in Go SDK ([#4936](https://github.com/databricks/terraform-provider-databricks/pull/4936))
