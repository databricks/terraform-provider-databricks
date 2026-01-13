# NEXT CHANGELOG

## Release v1.103.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fixed `databricks_dashboard` resource to detect content changes when using the `file_path` attribute. Previously, only changes to the path string itself triggered updates, not changes to the file content.

### Documentation

### Exporter

* Added support for `databricks_account_network_policy` and `databricks_workspace_network_option` resources ([#5238](https://github.com/databricks/terraform-provider-databricks/pull/5238)).

### Internal Changes
