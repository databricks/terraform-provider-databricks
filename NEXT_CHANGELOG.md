# NEXT CHANGELOG

## Release v1.114.2

### Breaking Changes

### New Features and Improvements

* Support adopting pre-existing `databricks_postgres_branch` and `databricks_postgres_endpoint` resources via `replace_existing = true` argument.

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Mark `effective_file_event_queue` as computed with diff suppression in `databricks_external_location` to prevent Terraform drift when the Unity Catalog backend returns the server-populated field.
