# NEXT CHANGELOG

## Release v1.96.0

### Breaking Changes

### New Features and Improvements

* Add `provider_config` support for manual plugin framework resources and data sources([#5127](https://github.com/databricks/terraform-provider-databricks/pull/5127))
* Improve `databricks_service_principals` data source ([#5164](https://github.com/databricks/terraform-provider-databricks/pull/5164))

* Added support for custom instance profiles on instance pools on AWS ([#5144](https://github.com/databricks/terraform-provider-databricks/pull/5144))


### Bug Fixes

* Remove unnecessary `SetSuppressDiff()` for `workload_size` in `databricks_model_serving` ([#5152](https://github.com/databricks/terraform-provider-databricks/pull/5152)).

### Documentation

* Fix missing GCP IAM permissions for workspace creation in GCP guides ([#5123](https://github.com/databricks/terraform-provider-databricks/pull/5123)).

### Exporter

### Internal Changes

* Caching group membership in `databricks_group_member` to improve performance ([#4581](https://github.com/databricks/terraform-provider-databricks/pull/4581)).
