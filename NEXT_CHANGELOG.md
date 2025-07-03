# NEXT CHANGELOG

## Release v1.85.0

### Breaking Changes

### New Features and Improvements

* Add `bearer_token` to the list of sensitive options in `databricks_connection` ([#4812](https://github.com/databricks/terraform-provider-databricks/pull/4812)).
* Use single-node cluster for `databricks_sql_permissions` ([#4813](https://github.com/databricks/terraform-provider-databricks/pull/4813)).

### Bug Fixes

* Fix updating of of `fallback_config` in `databricks_model_serving` ([#4830](https://github.com/databricks/terraform-provider-databricks/pull/4830)).

### Documentation

* Update documentation for single-node clusters in `databricks_cluster` resource ([#4817](https://github.com/databricks/terraform-provider-databricks/pull/4817)).

### Exporter

* Fix generation of columns in `databricks_sql_table` resource ([#4819](https://github.com/databricks/terraform-provider-databricks/pull/4819)).

### Internal Changes

 * Updated the contributing guide with instructions on how to modify the changelog ([#4404](https://github.com/databricks/terraform-provider-databricks/pull/4404)).
