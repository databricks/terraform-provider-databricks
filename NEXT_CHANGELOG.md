# NEXT CHANGELOG

## Release v1.86.0

### Breaking Changes
* Do not set `run_as from run_as_user_name in DLT pipelines. This fixes an issue where the value for run_as was unintentionally cached in the terraform state. More details and the workaround are specified in the PR: ([#4886](https://github.com/databricks/terraform-provider-databricks/pull/4886))

### New Features and Improvements

* Added ability to add `comment` when defining a new `databricks_share` ([#4802](https://github.com/databricks/terraform-provider-databricks/pull/4802))
* Added output attribute `endpoint_url` in `databricks_model_serving` ([#4877](https://github.com/databricks/terraform-provider-databricks/pull/4877)).
* Deprecate `egg` library type in `databricks_cluster`, `databricks_job`, and `databricks_library` ([#4881](https://github.com/databricks/terraform-provider-databricks/pull/4881)).

### Bug Fixes

* Corrected accidentally removed `SpID` field from `databricks_service_principal` ([#4868](https://github.com/databricks/terraform-provider-databricks/pull/4868)).
* Corrected optional fields in `databricks_mws_ncc_private_endpoint_rule`([#4856](https://github.com/databricks/terraform-provider-databricks/pull/4856)).
* Fix handling of `force` option in `databricks_git_credential` ([#4873](https://github.com/databricks/terraform-provider-databricks/pull/4873)).
* Restricted create or replace statement to managed tables in `databricks_sql_table`([#4874](https://github.com/databricks/terraform-provider-databricks/pull/4874)).
* Mitigate issue due to internal caching in `databricks_secret_acl` by retrying until ACL are applied with the right permission ([#4885](https://github.com/databricks/terraform-provider-databricks/pull/4885)).
* Fix schema mismatch bug in `databricks_functions` data source ([#4902](https://github.com/databricks/terraform-provider-databricks/pull/4902)).

### Documentation

* Updated `share` documentation to be more in line with Terraform styling ([#4802](https://github.com/databricks/terraform-provider-databricks/pull/4802))
* Refreshed `databricks_job` documentation ([#4861](https://github.com/databricks/terraform-provider-databricks/pull/4861)).
* Document `environment` block in `databricks_pipeline` ([#4878](https://github.com/databricks/terraform-provider-databricks/pull/4878)).
* Updated documentation for `databricks_disable_legacy_dbfs_setting` resource ([#4870](https://github.com/databricks/terraform-provider-databricks/pull/4870)).
* Add deprecation notice to `databricks_dbfs_file` and `databricks_mount` ([#4876](https://github.com/databricks/terraform-provider-databricks/pull/4876))
* Updated documentation for `databricks_disable_legacy_features_setting` resource ([#4884](https://github.com/databricks/terraform-provider-databricks/pull/4884)).
* Improve docs for `databricks_compliance_security_profile_setting` ([#4880](https://github.com/databricks/terraform-provider-databricks/pull/4880)).
* Improve instructions for the Terraform Exporter ([#4892](https://github.com/databricks/terraform-provider-databricks/pull/4892)).
* Improve documentation for service principal data sources ([#4900](https://github.com/databricks/terraform-provider-databricks/pull/4900)).

### Exporter

* Added support for exporting of workspaces and related resources ([#4899](https://github.com/databricks/terraform-provider-databricks/pull/4899)).

### Internal Changes

* Promote Plugin Framework Share Resource to Production ([#4846](https://github.com/databricks/terraform-provider-databricks/pull/4846)).
