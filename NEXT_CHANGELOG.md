# NEXT CHANGELOG

## Release v1.128.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_pipeline` with `ingestion_definition` failing when `serverless = false` is set alongside a `clusters` block ([#5783](https://github.com/databricks/terraform-provider-databricks/pull/5783)). The Go SDK marshals `serverless` with `omitempty`, so `false` was dropped and the platform defaulted ingestion pipelines to serverless, then rejected the cluster settings. The provider now force-sends `serverless` when it is explicitly set in the configuration.

### Documentation

* Document Genie budgets on `databricks_budget` with shared vs per-user examples, and blocking usage ([#5946](https://github.com/databricks/terraform-provider-databricks/pull/5946))

### Exporter

### Internal Changes
