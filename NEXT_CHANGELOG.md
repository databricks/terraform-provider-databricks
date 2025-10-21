# NEXT CHANGELOG

## Release v1.94.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Mark `workload_size` as not computed in `databricks_model_serving` ([#5111](https://github.com/databricks/terraform-provider-databricks/pull/5111))
* Fix crash when error happens during reading `databricks_job` ([#5110](https://github.com/databricks/terraform-provider-databricks/pull/5110))

### Documentation

* Document `table_update` trigger in `databricks_job` resource ([#5107](https://github.com/databricks/terraform-provider-databricks/pull/5107))
* Document new attributes in `databricks_app` resource and data sources ([#5108](https://github.com/databricks/terraform-provider-databricks/pull/5108))
* Document `git_email` in `databricks_git_credential` resource ([#5099](https://github.com/databricks/terraform-provider-databricks/pull/5099))

### Exporter

### Internal Changes

* Refactor `catalog_test.go` to use internal plan checks ([#5112](https://github.com/databricks/terraform-provider-databricks/pull/5112)).
