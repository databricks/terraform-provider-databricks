# NEXT CHANGELOG

## Release v1.131.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `databricks_pipeline` dropping `serverless = false` from create/update requests. Because `serverless` is `omitempty` in the SDK and was never force-sent, classic (non-serverless) ingestion pipelines failed with `cannot provide cluster settings when using serverless compute` (the API defaults an omitted `serverless` to `true` for ingestion pipelines, then rejects the `cluster` block).

### Documentation

### Exporter

### Internal Changes
