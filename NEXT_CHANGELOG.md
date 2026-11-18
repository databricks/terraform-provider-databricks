# NEXT CHANGELOG

## Release v1.144.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `thumbnail_path` to `databricks_app` to set the app's thumbnail from a local PNG or JPEG file ([#6196](https://github.com/databricks/terraform-provider-databricks/pull/6196)).
* Support `statement_timeout` for all SQL warehouse types; omitting or removing the argument preserves the existing warehouse timeout, while explicit `0` clears the override ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Bug Fixes
* Stop `databricks_volume` sending `new_name` on every update after an import by a full name with upper case ([#6209](https://github.com/databricks/terraform-provider-databricks/pull/6209)).

### Documentation

### Exporter

* Preserve `statement_timeout`, including `0`, in exports of all SQL warehouse types ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Internal Changes
