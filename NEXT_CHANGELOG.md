# NEXT CHANGELOG

## Release v1.133.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* `databricks_connection` now supports schema-level (L3) connections via the new `parent` argument (format `schemas/{catalog}.{schema}`). When set, the connection is created inside that schema and addressed by its `full_name` (`{catalog}.{schema}.{name}`); when omitted, the connection stays metastore-level, so existing configurations and state are unchanged ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).

### Bug Fixes
* `databricks_connection`: schema-level connections no longer show a perpetual `environment_settings` diff. The backend returns an empty `environment_settings` object for schema-level connections, which the provider materialized as an empty block that always diffed against configurations omitting it; the empty object is now dropped on read ([#6003](https://github.com/databricks/terraform-provider-databricks/pull/6003)).
* Fix `databricks_grant` and `databricks_grants` for the AI Gateway securables (`model_service`, `mcp_service`, `model_provider_service`). Wiring `databricks_ai_gateway_*.<x>.name` into the corresponding grant field previously failed on apply with `No API found for 'GET /unity-catalog/permissions/model_provider_service/model-provider-services/<full_name>'`, because that `name` attribute carries a resource-name prefix (e.g. `model-provider-services/`) that the permissions API does not expect. The prefix is now stripped before it reaches the API path and the resource ID, and a `DiffSuppressFunc` treats the prefixed and bare forms as equal, so both the `.name` reference and a bare full name apply cleanly with no perpetual diff.

### Documentation

### Exporter

### Internal Changes
