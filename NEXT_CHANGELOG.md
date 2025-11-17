# NEXT CHANGELOG

## Release v1.98.0

### Breaking Changes

### New Features and Improvements

* Add `databricks_users` data source ([#4028](https://github.com/databricks/terraform-provider-databricks/pull/4028))
* Improve `databricks_service_principals` data source ([#5164](https://github.com/databricks/terraform-provider-databricks/pull/5164))
* Add `role_arn` field to `databricks_mws_storage_configurations` resource to support sharing S3 buckets between root storage and Unity Catalog ([#5222](https://github.com/databricks/terraform-provider-databricks/issues/5222))

### Bug Fixes

* Fix spurious plan diffs in `databricks_model_serving` and `databricks_model_serving_provisioned_throughput` resources due to tag reordering ([#5120](https://github.com/databricks/terraform-provider-databricks/pull/5120))

### Documentation

### Exporter

* Added support for `databricks_data_quality_monitor` resource ([#5193](https://github.com/databricks/terraform-provider-databricks/pull/5193)).
* Fix typo in the name of environment variable ([#5158](https://github.com/databricks/terraform-provider-databricks/pull/5158)).
* Export permission assignments on workspace level ([#5169](https://github.com/databricks/terraform-provider-databricks/pull/5169)).
* Added support for Databricks Apps resources ([#5208](https://github.com/databricks/terraform-provider-databricks/pull/5208)).

### Internal Changes
