# NEXT CHANGELOG

## Release v1.92.0

### Breaking Changes

### New Features and Improvements

* Add `arm` option to `databricks_node_type` instead of `graviton` ([#5028](https://github.com/databricks/terraform-provider-databricks/pull/5028))

### Bug Fixes
 * Add support for share resource in plugin framework implementation to be SDKv2 compatible ([#4965](https://github.com/databricks/terraform-provider-databricks/pull/4965))

* Mark `storage_location` as read-only in `databricks_catalog` ([#5075](https://github.com/databricks/terraform-provider-databricks/pull/5075))

### Documentation

* Add instructions for migration from deprecated `databricks_catalog_workspace_binding` to `databricks_workspace_binding` ([#5054](https://github.com/databricks/terraform-provider-databricks/pull/5054))

### Exporter

### Internal Changes

* Make plugin framework implementation of share resource as default ([#4967](https://github.com/databricks/terraform-provider-databricks/pull/4967))
