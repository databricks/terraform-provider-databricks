# NEXT CHANGELOG

## Release v1.92.0

### Breaking Changes

* Remove stale resources/datasources/documentation related to Clean Room services.
* databricks\_permissions resource no longer updates permissions on delete. This is to mitigate an issue with incorrect IS\_OWNER being set.

### New Features and Improvements

* Add `arm` option to `databricks_node_type` instead of `graviton` ([#5028](https://github.com/databricks/terraform-provider-databricks/pull/5028))
* Add `data_quality_monitor` resource and data sources ([#5092](https://github.com/databricks/terraform-provider-databricks/pull/5092)).
* Add `data_quality_refresh` resource and data sources ([#5092](https://github.com/databricks/terraform-provider-databricks/pull/5092)).
* Perform workspace-level permission assignment by `user_name`, `group_name`, or `service_principal_name` ([#5068](https://github.com/databricks/terraform-provider-databricks/pull/5068)).

### Bug Fixes

* Fixed syncing of effective fields in plugin framework implementation of share resource ([#4969](https://github.com/databricks/terraform-provider-databricks/pull/4969))
* Mark `storage_location` as read-only in `databricks_catalog` ([#5075](https://github.com/databricks/terraform-provider-databricks/pull/5075))

### Documentation

* Add instructions for migration from deprecated `databricks_catalog_workspace_binding` to `databricks_workspace_binding` ([#5054](https://github.com/databricks/terraform-provider-databricks/pull/5054))
* Document output attributes in `databricks_storage_credential` ([#5093](https://github.com/databricks/terraform-provider-databricks/pull/5093))

### Exporter

### Internal Changes

* Bump the Go SDK to v0.86.0 ([#5092](https://github.com/databricks/terraform-provider-databricks/pull/5092)).
