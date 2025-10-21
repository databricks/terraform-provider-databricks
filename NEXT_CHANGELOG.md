# NEXT CHANGELOG

## Release v1.94.0

### Breaking Changes

### New Features and Improvements

* Added `expected_workspace_status` to `databricks_mws_workspaces` to support creating workspaces in provisioning status ([#5019](https://github.com/databricks/terraform-provider-databricks/pull/5019))

### Bug Fixes

* Fix crash when error happens during reading `databricks_job` ([#5110](https://github.com/databricks/terraform-provider-databricks/pull/5110))

### Documentation

* Document `table_update` trigger in `databricks_job` resource ([#5107](https://github.com/databricks/terraform-provider-databricks/pull/5107))
* Document new attributes in `databricks_app` resource and data sources ([#5108](https://github.com/databricks/terraform-provider-databricks/pull/5108))
* Document `git_email` in `databricks_git_credential` resource ([#5099](https://github.com/databricks/terraform-provider-databricks/pull/5099))

### Exporter

### Internal Changes
* Refactor `catalog_test.go` to use internal plan checks ([#5112](https://github.com/databricks/terraform-provider-databricks/pull/5112)).
