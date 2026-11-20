# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_lakewatch_parser`.
* Add resource and data sources for `databricks_lakewatch_parser_version`.
* Add data sources for `databricks_lakewatch_builtin_parser`.
* Add data sources for `databricks_lakewatch_builtin_parser_version`.

### Bug Fixes
* Escape `#`, `?` and `%` in the full name of a `databricks_volume` when reading, updating and deleting it ([#6206](https://github.com/databricks/terraform-provider-databricks/pull/6206)).

  Without escaping, these calls for the volume `main.replay.raw#2` went to `main.replay.raw`: a refresh could drop the volume from state and plan to create it again, and a destroy could delete a different volume.
* Escape `#`, `?` and `%` in the `name` of a `databricks_volume` data source when reading the volume ([#6195](https://github.com/databricks/terraform-provider-databricks/pull/6195)).

  Without escaping, the data source for `main.replay.raw#2` read `main.replay.raw`, so it returned a different volume or failed with a 404.

### Documentation

### Exporter

### Internal Changes
