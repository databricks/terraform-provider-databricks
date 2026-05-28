# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

* Add support for managing permissions of Agent Bricks resources ([#5708](https://github.com/databricks/terraform-provider-databricks/issues/5669)). Reverts [#5582](https://github.com/databricks/terraform-provider-databricks/pull/5708).

### Bug Fixes
* Fix permanent permissions drift when `user_name` casing in `databricks_permissions` `access_control` blocks differs from the API response ([#5183](https://github.com/databricks/terraform-provider-databricks/issues/5183))

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.
* Fixed destroying of UC objects when workspace binding removed before actual destroy ([#5581](https://github.com/databricks/terraform-provider-databricks/pull/5581)).
* Fixed handling of the case when library is removed outside of Terraform ([#5678](https://github.com/databricks/terraform-provider-databricks/pull/5678)).
* Fix `databricks_vector_search_index` hardcoded 15-minute creation timeout: increased default to 75 minutes (consistent with `databricks_vector_search_endpoint`) and made it user-overridable via the `timeouts` block.
* Fixed child groups collection in `databricks_group` data source ([#5679](https://github.com/databricks/terraform-provider-databricks/pull/5679)).

### Documentation

* Document that some `databricks_mws_*` resources on GCP require Google-issued OIDC tokens (not Databricks OAuth) ([#5654](https://github.com/databricks/terraform-provider-databricks/issues/5654)).
* Remove non-existent field from the `databricks_vector_search_index` doc ([#5605](https://github.com/databricks/terraform-provider-databricks/pull/5605)).

### Exporter

* Support `alert_task` when exporting `databricks_job` ([#5629](https://github.com/databricks/terraform-provider-databricks/pull/5629)).
* Add support for exporting Agent Bricks resources ([#5704](https://github.com/databricks/terraform-provider-databricks/issues/5704)).

### Internal Changes

* Add `internal/retrier` package for unified retry and backoff handling ([#5746](https://github.com/databricks/terraform-provider-databricks/pull/5746)).
* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).
