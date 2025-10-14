# NEXT CHANGELOG

## Release v1.92.0

### Breaking Changes

* Remove stale resources/datasources/documentation related to Clean Room services.
  
### New Features and Improvements

* Added `expected_workspace_status` to `databricks_mws_workspaces` to support creating workspaces in provisioning status ([#5019](https://github.com/databricks/terraform-provider-databricks/pull/5019))
* Add `arm` option to `databricks_node_type` instead of `graviton` ([#5028](https://github.com/databricks/terraform-provider-databricks/pull/5028))

### Bug Fixes

* Fixed syncing of effective fields in plugin framework implementation of share resource ([#4969](https://github.com/databricks/terraform-provider-databricks/pull/4969))
* Mark `storage_location` as read-only in `databricks_catalog` ([#5075](https://github.com/databricks/terraform-provider-databricks/pull/5075))

### Documentation

* Add instructions for migration from deprecated `databricks_catalog_workspace_binding` to `databricks_workspace_binding` ([#5054](https://github.com/databricks/terraform-provider-databricks/pull/5054))

### Exporter

### Internal Changes

* Use `Jobs.Get` instead of `JobsGetByJobId` ([#5029](https://github.com/databricks/terraform-provider-databricks/pull/5029))
