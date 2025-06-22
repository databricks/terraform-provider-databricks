# NEXT CHANGELOG

## Release v1.84.0

### Breaking Changes

### New Features and Improvements

* Added `time_rotating` argument to `databricks_service_principal_secret` to enable secret rotation ([#4789](https://github.com/databricks/terraform-provider-databricks/pull/4789)).
* Add AWS support for `databricks_mws_ncc_private_endpoint_rule` ([#4804](https://github.com/databricks/terraform-provider-databricks/pull/4804)).
* Added `key` argument to `databricks_jobs` data source to enable mapping by job ID and allow duplicate job names ([#4796](https://github.com/databricks/terraform-provider-databricks/pull/4796)).

### Bug Fixes

### Documentation

* Added HCL provider example for WIF authentication ([#4799](https://github.com/databricks/terraform-provider-databricks/pull/4799))
* Added link to Workload Identity Federation page ([#4786](https://github.com/databricks/terraform-provider-databricks/pull/4786)).
* auto `zone_id` can only be used for fleet node types in `databricks_instance_pool` resource ([#4782](https://github.com/databricks/terraform-provider-databricks/pull/4782)).
* Document `tags` attribute in `databricks_pipeline` resource ([#4783](https://github.com/databricks/terraform-provider-databricks/pull/4783)).
* Recommend OAuth instead of PAT in guides ([#4787](https://github.com/databricks/terraform-provider-databricks/pull/4787))
* Document new options in `databricks_model_serving` resource ([#4789](https://github.com/databricks/terraform-provider-databricks/pull/4789))
* Added example of granting Account Admin role with `databricks_service_principal_role` resource ([#4807](https://github.com/databricks/terraform-provider-databricks/pull/4807))

### Exporter

* Fix crash when importing model serving endpoint ([#4806](https://github.com/databricks/terraform-provider-databricks/pull/4806))

### Internal Changes
