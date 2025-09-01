# NEXT CHANGELOG

## Release v1.88.0

### Breaking Changes

### New Features and Improvements

* Document and handle additional Slack options in `databricks_notification_destination` ([#4933](https://github.com/databricks/terraform-provider-databricks/pull/4933))
* Lower case `warehouse_name_contains` when comparing in `databricks_warehouses` ([#4966](https://github.com/databricks/terraform-provider-databricks/pull/4966))

### Bug Fixes

* Suppress diff for `databricks_catalog` attributes ([#4975](https://github.com/databricks/terraform-provider-databricks/pull/4975))
* Correct which file event fields should be reset in `databricks_external_location` ([#4945](https://github.com/databricks/terraform-provider-databricks/pull/4945))
* Fix `ExactlyOneOf` in `databricks_app` ([#4946](https://github.com/databricks/terraform-provider-databricks/pull/4946))
* Enable update of `databricks_mws_ncc_private_endpoint_rule` resource ([#4957](https://github.com/databricks/terraform-provider-databricks/pull/4957))
* Fix metastore ID interpolation for account-level storage credential imports ([#4980](https://github.com/databricks/terraform-provider-databricks/pull/4980)). Storage credentials can now be imported using two formats: `<storage_credential_name>` when using a workspace-level provider, and `<metastore_id>|<storage_credential_name>` when using an account-level provider. Previously, importing storage credentials with an account-level provider would fail due to missing metastore ID in the API call.

### Documentation

* Improve documentation for grant resource ([#4906](https://github.com/databricks/terraform-provider-databricks/pull/4935))
* Document `gcp_attributes.first_on_demand` attribute in `databricks_cluster` ([#4934](https://github.com/databricks/terraform-provider-databricks/pull/4934))
* Improve `databricks_mws_permission_assignment` documentation ([#4943](https://github.com/databricks/terraform-provider-databricks/pull/4943))
* Update `databricks_job` example to use `environment_version` ([#4942](https://github.com/databricks/terraform-provider-databricks/pull/4942))
* Fix broken link in docs/index.md ([#4949](https://github.com/databricks/terraform-provider-databricks/pull/4949))
* Update existing docs to match Go SDK 0.81.0 ([#4960](https://github.com/databricks/terraform-provider-databricks/pull/4960))

### Exporter

* Add match by name to more exported resources ([#4939](https://github.com/databricks/terraform-provider-databricks/pull/4939))
* Improve handling of new dependencies in jobs, pipelines, model serving ([#4914](https://github.com/databricks/terraform-provider-databricks/pull/4914))
* Add support for `databricks_budget` ([#4957](https://github.com/databricks/terraform-provider-databricks/pull/4957))
* Correct support for `library.glob` in `databricks_pipeline` ([#4937](https://github.com/databricks/terraform-provider-databricks/pull/4937))
* Resolve references also for map values ([#4944](https://github.com/databricks/terraform-provider-databricks/pull/4944))

### Internal Changes

* Caching group membership in `databricks_group_member` to improve performance ([#4581](https://github.com/databricks/terraform-provider-databricks/pull/4581)).
* Replaced `common.APIErrorBody` with corresponding structs in Go SDK ([#4936](https://github.com/databricks/terraform-provider-databricks/pull/4936))
* Reimplement `databricks_group` data source to use combination of List + Get API ([#4947](https://github.com/databricks/terraform-provider-databricks/pull/4947))
* Use databricks_sql_table Instead of databricks_table in Sharing Tests ([#4981](https://github.com/databricks/terraform-provider-databricks/pull/4981)
