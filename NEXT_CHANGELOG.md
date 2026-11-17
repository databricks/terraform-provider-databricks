# NEXT CHANGELOG

## Release v1.144.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `thumbnail_path` to `databricks_app` to set the app's thumbnail from a local PNG or JPEG file ([#6196](https://github.com/databricks/terraform-provider-databricks/pull/6196)).
* Support `statement_timeout` for all SQL warehouse types; omitting or removing the argument preserves the existing warehouse timeout, while explicit `0` clears the override ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Bug Fixes
* Keep a valid `databricks_volume` ID when the create or update response has no `full_name`, so the read after the write no longer fails ([#6208](https://github.com/databricks/terraform-provider-databricks/pull/6208)).

### Documentation

### Exporter

* Preserve `statement_timeout`, including `0`, in exports of all SQL warehouse types ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Internal Changes
