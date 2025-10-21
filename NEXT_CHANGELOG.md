# NEXT CHANGELOG

## Release v1.94.0

### New Features and Improvements

* Add `arm` option to `databricks_node_type` instead of `graviton` ([#5028](https://github.com/databricks/terraform-provider-databricks/pull/5028))
* Optimize `databricks_grant` and `databricks_grants` to not call the `Update` API if the requested permissions are already granted ([#5095](https://github.com/databricks/terraform-provider-databricks/pull/5095))

### New Features and Improvements

### Bug Fixes

* Fix crash when error happens during reading `databricks_job` ([#5110](https://github.com/databricks/terraform-provider-databricks/pull/5110))

### Documentation

* Document `table_update` trigger in `databricks_job` resource ([#5107](https://github.com/databricks/terraform-provider-databricks/pull/5107))
* Document new attributes in `databricks_app` resource and data sources ([#5108](https://github.com/databricks/terraform-provider-databricks/pull/5108))
* Document `git_email` in `databricks_git_credential` resource ([#5099](https://github.com/databricks/terraform-provider-databricks/pull/5099))

### Exporter

### Internal Changes
* Refactor `catalog_test.go` to use internal plan checks ([#5112](https://github.com/databricks/terraform-provider-databricks/pull/5112)).
