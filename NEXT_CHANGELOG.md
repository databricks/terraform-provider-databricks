# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.
* Fixed destroying of UC objects when workspace binding removed before actual destroy ([#5581](https://github.com/databricks/terraform-provider-databricks/pull/5581)).
* Fixed handling of the case when library is removed outside of Terraform ([#5678](https://github.com/databricks/terraform-provider-databricks/pull/5678)).
* Fix `databricks_vector_search_index` hardcoded 15-minute creation timeout: increased default to 75 minutes (consistent with `databricks_vector_search_endpoint`) and made it user-overridable via the `timeouts` block.

### Documentation

* Remove non-existent field from the `databricks_vector_search_index` doc ([#5605](https://github.com/databricks/terraform-provider-databricks/pull/5605)).

### Exporter

* Support `alert_task` when exporting `databricks_job` ([#5629](https://github.com/databricks/terraform-provider-databricks/pull/5629)).
* Add support for exporting Agent Bricks resources ([#5704](https://github.com/databricks/terraform-provider-databricks/issues/5704)).

### Internal Changes

* Add `internal/retrier` package for unified retry and backoff handling ([#5746](https://github.com/databricks/terraform-provider-databricks/pull/5746)).
* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).
