# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add the `USE_SDK_GO_RESOURCES` environment variable, which makes the listed resources read and write through the modular `github.com/databricks/sdk-go` clients ([#6229](https://github.com/databricks/terraform-provider-databricks/pull/6229), [#6231](https://github.com/databricks/terraform-provider-databricks/pull/6231)).

  `USE_SDK_GO_RESOURCES=databricks_volume` makes `databricks_volume` read and write through the `github.com/databricks/sdk-go/uc/volumes` v1 client. It is unset by default, and an unknown resource name fails provider configuration.

### Bug Fixes

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
