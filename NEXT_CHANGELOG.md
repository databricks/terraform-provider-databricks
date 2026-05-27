# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.
* Fixed destroying of UC objects when workspace binding removed before actual destroy ([#5581](https://github.com/databricks/terraform-provider-databricks/pull/5581)).
* Fix `databricks_vector_search_index` hardcoded 15-minute creation timeout: increased default to 75 minutes (consistent with `databricks_vector_search_endpoint`) and made it user-overridable via the `timeouts` block.

### Documentation

* Remove non-existent field from the `databricks_vector_search_index` doc ([#5605](https://github.com/databricks/terraform-provider-databricks/pull/5605)).

### Exporter

* Support `alert_task` when exporting `databricks_job` ([#5629](https://github.com/databricks/terraform-provider-databricks/pull/5629)).

### Internal Changes

* Add `internal/retrier` package for unified retry and backoff handling ([#5746](https://github.com/databricks/terraform-provider-databricks/pull/5746)).
* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).
* `workspace_id` (provider attribute and `provider_config.workspace_id` block) now accepts workspace connection IDs in addition to classic numeric workspace IDs. On unified Databricks hosts, the platform gateway disambiguates the value server-side via the `X-Databricks-Workspace-Id` request header. The previous positive-integer validator has been relaxed to require only a non-empty string.

  Numeric workspace IDs continue to behave exactly as before, including the account-API workspace-deployment lookup. Connection IDs skip that lookup and route directly via the configured host. When the provider is configured at the workspace level (host + token), connection IDs surface a clear error directing the user to reconfigure with account-level credentials, since a workspace-level provider can only operate on a single workspace.
