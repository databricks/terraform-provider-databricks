# NEXT CHANGELOG

## Release v1.104.0

### Breaking Changes

### New Features and Improvements
* Added resources and data sources for `databricks_account_user_setting_v2` and `databricks_default_warehouse_override` ([#5329](https://github.com/databricks/terraform-provider-databricks/pull/5329)).


* Add `provider_config` support for manual plugin framework resources and data sources([#5127](https://github.com/databricks/terraform-provider-databricks/pull/5127))

* Added support for custom instance profiles on instance pools on AWS ([#5144](https://github.com/databricks/terraform-provider-databricks/pull/5144))

* Added `deployment_names` attribute to `databricks_mws_workspaces` data block ([#5100](https://github.com/databricks/terraform-provider-databricks/pull/5100))

### Bug Fixes

* Fix importing of the `databricks_share` ([#5311](https://github.com/databricks/terraform-provider-databricks/pull/5311))
* Fix `databricks_dashboard` resource to include `dataset_catalog` and `dataset_schema` when retrying creation after parent folder creation ([#5327](https://github.com/databricks/terraform-provider-databricks/pull/5327))

### Documentation

### Exporter

* Rewrite cloud-specific attributes and node types in cluster policy definitions when using `-targetCloud` flag ([#5297](https://github.com/databricks/terraform-provider-databricks/issues/5297)).
* Added support for `databricks_account_network_policy` and `databricks_workspace_network_option` resources ([#5238](https://github.com/databricks/terraform-provider-databricks/pull/5238)).
* Rewrite exporting of `databricks_share` to plugin framework ([#5328](https://github.com/databricks/terraform-provider-databricks/pull/5328))

### Internal Changes
* Add support for unified host, migrate IsAccountClient to HostType
