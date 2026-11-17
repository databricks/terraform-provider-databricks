# NEXT CHANGELOG

## Release v1.144.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `thumbnail_path` to `databricks_app` to set the app's thumbnail from a local PNG or JPEG file ([#6196](https://github.com/databricks/terraform-provider-databricks/pull/6196)).
* Support `statement_timeout` for all SQL warehouse types; omitting or removing the argument preserves the existing warehouse timeout, while explicit `0` clears the override ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Bug Fixes
* Escape `#`, `?` and `%` in the full name of a `databricks_volume` when reading, updating and deleting it ([#6206](https://github.com/databricks/terraform-provider-databricks/pull/6206)).

  Without escaping, these calls for the volume `main.replay.raw#2` went to `main.replay.raw`: a refresh could drop the volume from state and plan to create it again, and a destroy could delete a different volume.

### Documentation

### Exporter

* Preserve `statement_timeout`, including `0`, in exports of all SQL warehouse types ([#6144](https://github.com/databricks/terraform-provider-databricks/pull/6144)).

### Internal Changes
