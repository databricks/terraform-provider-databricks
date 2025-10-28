# NEXT CHANGELOG

## Release v1.96.0

### Breaking Changes

### New Features and Improvements

* Add `provider_config` support for manual plugin framework resources and data sources([#5127](https://github.com/databricks/terraform-provider-databricks/pull/5127))

* Added support for custom instance profiles on instance pools on AWS ([#5144](https://github.com/databricks/terraform-provider-databricks/pull/5144))


### Bug Fixes

* Fix permanent drift in `databricks_model_serving` when using `*_plaintext` credential fields for external models ([#5125](https://github.com/databricks/terraform-provider-databricks/pull/5125))
* Remove unnecessary `SetSuppressDiff()` for `workload_size` in `databricks_model_serving` ([#5152](https://github.com/databricks/terraform-provider-databricks/pull/5152)).

### Documentation

### Exporter

### Internal Changes

* Caching group membership in `databricks_group_member` to improve performance ([#4581](https://github.com/databricks/terraform-provider-databricks/pull/4581)).
