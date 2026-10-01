# NEXT CHANGELOG

## Release v1.136.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes

* Fixed `databricks_pipeline` dropping `serverless = false` from create/update requests. Because `serverless` is `omitempty` in the SDK and was never force-sent, classic (non-serverless) ingestion pipelines failed with `cannot provide cluster settings when using serverless compute` (the API defaults an omitted `serverless` to `true` for ingestion pipelines, then rejects the `cluster` block).

### Documentation

### Exporter

* Fixed a nil pointer panic in `emitRfaAccessRequestDestinations` when exporting workspace-level UC securables with an account-level provider ([#6028](https://github.com/databricks/terraform-provider-databricks/issues/6028)).

### Internal Changes
