# NEXT CHANGELOG

## Release v1.85.0

### Breaking Changes

### New Features and Improvements

* Add `bearer_token` to the list of sensitive options in `databricks_connection` ([#4812](https://github.com/databricks/terraform-provider-databricks/pull/4812)).
* Use single-node cluster for `databricks_sql_permissions` ([#4813](https://github.com/databricks/terraform-provider-databricks/pull/4813)).
* Allow to retrieve service principal data by SCIM ID ([#3142](https://github.com/databricks/terraform-provider-databricks/pull/3142)).
* Add support for Lakebase `databricks_database_instance` in  `databricks_permissions` ([#4824](https://github.com/databricks/terraform-provider-databricks/pull/4824)).
* Document new fields in `databricks_model_serving` and deprecate `invalid_keywords` and `valid_topics` in AI Gateway configuration ([#4851](https://github.com/databricks/terraform-provider-databricks/pull/4851)).
* Added support for Alert V2 in `databricks_permissions` ([#4831](https://github.com/databricks/terraform-provider-databricks/pull/4831)).
* Replace instead of dropping Delta `databricks_sql_table` ([#2424](https://github.com/databricks/terraform-provider-databricks/pull/2424)).

### Bug Fixes

* Don't redeploy `databricks_sql_table` when view definition contains newlines or tabs ([#4003](https://github.com/databricks/terraform-provider-databricks/pull/4003)).
* Preserve `queue` setting for `databricks_job` resource when upgrading from provider version <1.71.0.
* Fix updating of of `fallback_config` in `databricks_model_serving` ([#4830](https://github.com/databricks/terraform-provider-databricks/pull/4830)).

### Documentation

* Update documentation for single-node clusters in `databricks_cluster` resource ([#4817](https://github.com/databricks/terraform-provider-databricks/pull/4817)).
* Update GCP example for `databricks_external_location` resource ([#4826](https://github.com/databricks/terraform-provider-databricks/pull/4826))
* Fix formatting for HTTP connection example in `databricks_connection` resource ([#4826](https://github.com/databricks/terraform-provider-databricks/pull/4826))
* Update Databricks SQL objects documentation ([#4840](https://github.com/databricks/terraform-provider-databricks/pull/4840))
* Improve documentation for `databricks_git_credential` resource ([#4837](https://github.com/databricks/terraform-provider-databricks/pull/4837))
* Rename DLT references to Lakeflow Declarative pipelines ([#4842](https://github.com/databricks/terraform-provider-databricks/pull/4842))
* Clarify and add more examples to `databricks_mws_network_connectivity_config` and `databricks_mws_ncc_private_endpoint_rule` documentation ([#4847](https://github.com/databricks/terraform-provider-databricks/pull/4847))

### Exporter

* Fix generation of columns in `databricks_sql_table` resource ([#4819](https://github.com/databricks/terraform-provider-databricks/pull/4819)).

### Internal Changes

* Updated the contributing guide with instructions on how to modify the changelog ([#4404](https://github.com/databricks/terraform-provider-databricks/pull/4404)).
* Add synthetic field to jobs resource to control behavior of `apply_policy_default_values` ([#4834](https://github.com/databricks/terraform-provider-databricks/pull/4834)).
