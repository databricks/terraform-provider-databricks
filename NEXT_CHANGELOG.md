# NEXT CHANGELOG

## Release v1.85.0

### Breaking Changes

### New Features and Improvements

* Add `bearer_token` to the list of sensitive options in `databricks_connection` ([#4812](https://github.com/databricks/terraform-provider-databricks/pull/4812)).
* Use single-node cluster for `databricks_sql_permissions` ([#4813](https://github.com/databricks/terraform-provider-databricks/pull/4813)).
* Add support for Lakebase `databricks_database_instance` in  `databricks_permissions` ([#4824](https://github.com/databricks/terraform-provider-databricks/pull/4824)).
* Added support for Alert V2 in `databricks_permissions` ([#4831](https://github.com/databricks/terraform-provider-databricks/pull/4831)).
* Replace instead of dropping Delta `databricks_sql_table` ([#2424](https://github.com/databricks/terraform-provider-databricks/pull/2424)).

### Bug Fixes

* Fix updating of of `fallback_config` in `databricks_model_serving` ([#4830](https://github.com/databricks/terraform-provider-databricks/pull/4830)).
* Fixed `databricks_secret_acl` resource creation to properly verify ACL application with retry logic when API returns success but doesn't actually create the ACL.

### Documentation

* Update documentation for single-node clusters in `databricks_cluster` resource ([#4817](https://github.com/databricks/terraform-provider-databricks/pull/4817)).
* Update GCP example for `databricks_external_location` resource ([#4826](https://github.com/databricks/terraform-provider-databricks/pull/4826))
* Fix formatting for HTTP connection example in `databricks_connection` resource ([#4826](https://github.com/databricks/terraform-provider-databricks/pull/4826))
* Improve documentation for `databricks_git_credential` resource ([#4837](https://github.com/databricks/terraform-provider-databricks/pull/4837))

### Exporter

* Fix generation of columns in `databricks_sql_table` resource ([#4819](https://github.com/databricks/terraform-provider-databricks/pull/4819)).

### Internal Changes

 * Updated the contributing guide with instructions on how to modify the changelog ([#4404](https://github.com/databricks/terraform-provider-databricks/pull/4404)).
