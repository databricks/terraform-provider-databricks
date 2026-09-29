# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

 * Fixed `databricks_app` updates. The resource now uses the asynchronous update API with a field mask of the changed attributes, so changing a description, scopes, resources or sizing no longer fails with "Compute size updates are not supported in this update API" ([#6034](https://github.com/databricks/terraform-provider-databricks/pull/6034)).

### Documentation

### Exporter

### Internal Changes
