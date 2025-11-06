# NEXT CHANGELOG

## Release v1.97.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed spurious diffs in `databricks_model_serving` and `databricks_model_serving_provisioned_throughput` resources when the API returns `served_models` or `served_entities` in a different order than specified in the configuration ([#5188](https://github.com/databricks/terraform-provider-databricks/pull/5188)).

### Documentation

* Explain that special characters in URLs should be percent-encoded ([#5178](https://github.com/databricks/terraform-provider-databricks/pull/5178)).
* Improve documentation about `preloaded_docker_images.basic_auth` in `databricks_cluster` and `databricks_instance_pool` ([#5154](https://github.com/databricks/terraform-provider-databricks/pull/5154)).

### Exporter

* Initial support for resources implemented with plugin framework ([#5176](https://github.com/databricks/terraform-provider-databricks/pull/5176)).

### Internal Changes
