# NEXT CHANGELOG

## Release v1.144.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Support `statement_timeout` for all SQL warehouse types; omitting or removing the argument preserves the existing warehouse timeout, while explicit `0` clears the override ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Bug Fixes
* Fixed `enable_file_events = false` having no effect on `databricks_external_location` when Databricks had enabled file events automatically ([#5967](https://github.com/databricks/terraform-provider-databricks/issues/5967)).

### Documentation

### Exporter

* Preserve `statement_timeout`, including `0`, in exports of all SQL warehouse types ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Internal Changes
