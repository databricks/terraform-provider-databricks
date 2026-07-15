# NEXT CHANGELOG

## Release v1.122.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Added `cascade_on_destroy` attribute to `databricks_pipeline` to control whether destroying a pipeline also deletes its datasets (materialized views, streaming tables, and views). Defaults to `true`; set to `false` to preserve the datasets on destroy ([#5860](https://github.com/databricks/terraform-provider-databricks/pull/5860)).

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
