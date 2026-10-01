# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes

 * Fixed `databricks_app` updates. The resource now uses the asynchronous update API with a field mask of the changed attributes, so changing a description, scopes, resources or sizing no longer fails with "Compute size updates are not supported in this update API" ([#6034](https://github.com/databricks/terraform-provider-databricks/pull/6034)).

### Documentation

### Exporter

* Fixed a nil pointer panic in `emitRfaAccessRequestDestinations` when exporting workspace-level UC securables with an account-level provider ([#6028](https://github.com/databricks/terraform-provider-databricks/issues/6028)).

### Internal Changes
