# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes

* `databricks_mlflow_model` can run on the modular Databricks SDK for Go (`github.com/databricks/sdk-go/modelregistry/v1`) instead of the classic one. It does so only when the environment variable `USE_SDK_GO_RESOURCES` names `databricks_mlflow_model`, which is not offered to users yet, so nothing changes by default ([#6309](https://github.com/databricks/terraform-provider-databricks/pull/6309)).
