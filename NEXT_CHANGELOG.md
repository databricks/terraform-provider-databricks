# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.

### Documentation

* Documented `principal_id` argument for `databricks_git_credential` resource, allowing management of Git credentials on behalf of service principals.

### Exporter

### Internal Changes

* Add `internal/retrier` package for unified retry and backoff handling ([#5746](https://github.com/databricks/terraform-provider-databricks/pull/5746)).
* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).
