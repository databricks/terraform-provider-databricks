# NEXT CHANGELOG

## Release v1.97.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed spurious diffs in `databricks_model_serving` and `databricks_model_serving_provisioned_throughput` resources when the API returns `served_models` or `served_entities` in a different order than specified in the configuration ([#5188](https://github.com/databricks/terraform-provider-databricks/pull/5188)).

### Documentation

* Improve documentation about `preloaded_docker_images.basic_auth` in `databricks_cluster` and `databricks_instance_pool` ([#5154](https://github.com/databricks/terraform-provider-databricks/pull/5154)).

### Exporter

* Support exporting of `databricks_mws_ncc_binding` ([#5184](https://github.com/databricks/terraform-provider-databricks/pull/5184)).
* Initial support for resources implemented with plugin framework ([#5176](https://github.com/databricks/terraform-provider-databricks/pull/5176)).

### Internal Changes
