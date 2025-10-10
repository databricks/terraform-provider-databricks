# NEXT CHANGELOG

## Release v1.92.0

### Breaking Changes

### New Features and Improvements

* Add `arm` option to `databricks_node_type` instead of `graviton` ([#5028](https://github.com/databricks/terraform-provider-databricks/pull/5028))

### Bug Fixes
 * Fix `databricks_cluster` drift when `is_single_node` is `true` ([#4360](https://github.com/databricks/terraform-provider-databricks/issues/4360)).

* Mark `storage_location` as read-only in `databricks_catalog` ([#5075](https://github.com/databricks/terraform-provider-databricks/pull/5075))

### Documentation

* Add instructions for migration from deprecated `databricks_catalog_workspace_binding` to `databricks_workspace_binding` ([#5054](https://github.com/databricks/terraform-provider-databricks/pull/5054))

### Exporter

### Internal Changes
