# NEXT CHANGELOG

## Release v1.129.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

* Skip system-managed jobs during export and add missing file references for job task parameters ([#5956](https://github.com/databricks/terraform-provider-databricks/issues/5956)).
* Added support for exporting `databricks_endpoint` resource ([#5951](https://github.com/databricks/terraform-provider-databricks/pull/5951)).
* Add an `exporter` dimension to the user agent ([#5954](https://github.com/databricks/terraform-provider-databricks/pull/5954)).
* Preserve zero `value` fields when exporting `databricks_workspace_setting_v2` and `databricks_account_setting_v2` ([#5955](https://github.com/databricks/terraform-provider-databricks/issues/5955)).

### Internal Changes
