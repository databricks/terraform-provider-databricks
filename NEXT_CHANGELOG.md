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

* Fix `databricks_pipeline` with `ingestion_definition` failing when `serverless = false` is set alongside a `clusters` block ([#5783](https://github.com/databricks/terraform-provider-databricks/pull/5783)). The Go SDK marshals `serverless` with `omitempty`, so `false` was dropped and the platform defaulted ingestion pipelines to serverless, then rejected the cluster settings. The provider now force-sends `serverless` when it is explicitly set in the configuration.

### Documentation

### Exporter

### Internal Changes
