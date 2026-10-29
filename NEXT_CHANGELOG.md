# NEXT CHANGELOG

## Release v1.141.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Added `cascade_on_destroy` attribute to `databricks_pipeline` to control whether destroying a pipeline also deletes its datasets (materialized views, streaming tables, and views). Defaults to `true`; set to `false` to preserve the datasets on destroy ([#5860](https://github.com/databricks/terraform-provider-databricks/pull/5860)).

### Bug Fixes

### Documentation

* Fix the Azure Databricks US Gov account console host in the provider authentication docs to `https://accounts.usgov.databricks.azure.us` ([#6071](https://github.com/databricks/terraform-provider-databricks/pull/6071)).

### Exporter

### Internal Changes
