# NEXT CHANGELOG

## Release v1.101.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_dashboard` resource to detect content changes when using the `file_path` attribute. Previously, only changes to the path string itself triggered updates, not changes to the file content.
* [Fix] Allow Updating Share Objects With shared_as Defined ([#5287](https://github.com/databricks/terraform-provider-databricks/pull/5287))

### Documentation

### Exporter

### Internal Changes

* Switch to use Go SDK struct in `databricks_metastore` resource ([#5088](https://github.com/databricks/terraform-provider-databricks/pull/5088))
