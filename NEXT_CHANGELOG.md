# NEXT CHANGELOG

## Release v1.98.0

### Breaking Changes

### New Features and Improvements

* Add `databricks_users` data source ([#4028](https://github.com/databricks/terraform-provider-databricks/pull/4028))
* Improve `databricks_service_principals` data source ([#5164](https://github.com/databricks/terraform-provider-databricks/pull/5164))

### Bug Fixes

* Fix spurious plan diffs in `databricks_model_serving` and `databricks_model_serving_provisioned_throughput` resources due to tag reordering ([#5120](https://github.com/databricks/terraform-provider-databricks/pull/5120))
* Move Spark Version selector defaults to Terraform ([#5219](https://github.com/databricks/terraform-provider-databricks/pull/5219)).

### Documentation

* Document tag policies in `databricks_access_control_rule_set` ([#5209](https://github.com/databricks/terraform-provider-databricks/pull/5209)).
* Document missing `aws_attributes.ebs_*` properties in `databricks_cluster` ([#5196](https://github.com/databricks/terraform-provider-databricks/pull/5196)).
* Document support for serverless workspaces on GCP ([#5124](https://github.com/databricks/terraform-provider-databricks/pull/5124))

### Exporter

* Added support for `databricks_account_network_policy` and `databricks_workspace_network_option` resources ([#5238](https://github.com/databricks/terraform-provider-databricks/pull/5238)).
* Added support for `databricks_data_quality_monitor` resource ([#5193](https://github.com/databricks/terraform-provider-databricks/pull/5193)).
* Fix typo in the name of environment variable ([#5158](https://github.com/databricks/terraform-provider-databricks/pull/5158)).
* Export permission assignments on workspace level ([#5169](https://github.com/databricks/terraform-provider-databricks/pull/5169)).
* Added support for Databricks Apps resources ([#5208](https://github.com/databricks/terraform-provider-databricks/pull/5208)).
* Added support for Database Instance resource (aka Lakebase) ([#5212](https://github.com/databricks/terraform-provider-databricks/pull/5212)).

### Internal Changes
