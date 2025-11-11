# NEXT CHANGELOG

## Release v1.98.0

### Breaking Changes

### New Features and Improvements

* Add `databricks_users` data source ([#4028](https://github.com/databricks/terraform-provider-databricks/pull/4028))
* Improve `databricks_service_principals` data source ([#5164](https://github.com/databricks/terraform-provider-databricks/pull/5164))

### Bug Fixes

* Fixed `databricks_dashboard` resource to detect content changes when using the `file_path` attribute. Previously, only changes to the path string itself triggered updates, not changes to the file content.

### Documentation

### Exporter

* Fix typo in the name of environment variable ([#5158](https://github.com/databricks/terraform-provider-databricks/pull/5158)).

### Internal Changes
