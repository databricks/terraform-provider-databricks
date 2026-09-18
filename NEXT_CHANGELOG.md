# NEXT CHANGELOG

## Release v1.133.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fix `databricks_grant` and `databricks_grants` for the AI Gateway securables (`model_service`, `mcp_service`, `model_provider_service`). Wiring `databricks_ai_gateway_*.<x>.name` into the corresponding grant field previously failed on apply with `No API found for 'GET /unity-catalog/permissions/model_provider_service/model-provider-services/<full_name>'`, because that `name` attribute carries a resource-name prefix (e.g. `model-provider-services/`) that the permissions API does not expect. The prefix is now stripped before it reaches the API path and the resource ID, and a `DiffSuppressFunc` treats the prefixed and bare forms as equal, so both the `.name` reference and a bare full name apply cleanly with no perpetual diff.

* Fixed `databricks_pipeline` dropping `serverless = false` from create/update requests. Because `serverless` is `omitempty` in the SDK and was never force-sent, classic (non-serverless) ingestion pipelines failed with `cannot provide cluster settings when using serverless compute` (the API defaults an omitted `serverless` to `true` for ingestion pipelines, then rejects the `cluster` block).

### Documentation

### Exporter

### Internal Changes
