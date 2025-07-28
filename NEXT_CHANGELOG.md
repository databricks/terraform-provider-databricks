# NEXT CHANGELOG

## Release v1.86.0

### Breaking Changes

### New Features and Improvements

* Added output attribute `endpoint_url` in `databricks_model_serving`([#4877](https://github.com/databricks/terraform-provider-databricks/pull/4877)).

### Bug Fixes

* Corrected accidentally removed `SpID` field from `databricks_service_principal` ([#4868](https://github.com/databricks/terraform-provider-databricks/pull/4868)).
* Corrected optional fields in `databricks_mws_ncc_private_endpoint_rule`([#4856](https://github.com/databricks/terraform-provider-databricks/pull/4856)).
* Restricted create or replace statement to managed tables in `databricks_sql_table`([#4874](https://github.com/databricks/terraform-provider-databricks/pull/4874)).
* Mitigate issue due to internal caching in `databricks_secret_acl` by retrying until ACL are applied with the right permission ([]())

### Documentation

* Refreshed `databricks_job` documentation ([#4861](https://github.com/databricks/terraform-provider-databricks/pull/4861)).
* Document `environment` block in `databricks_pipeline` ([#4878](https://github.com/databricks/terraform-provider-databricks/pull/4878)).
* Updated documentation for `databricks_disable_legacy_dbfs_setting` resource ([#4870](https://github.com/databricks/terraform-provider-databricks/pull/4870)).
* Add deprecation notice to `databricks_dbfs_file` and `databricks_mount` ([#4876](https://github.com/databricks/terraform-provider-databricks/pull/4876))

### Exporter

### Internal Changes
