# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.

### Documentation

### Exporter

### Internal Changes

* Add `internal/retrier` package for unified retry and backoff handling ([#XXXX](https://github.com/databricks/terraform-provider-databricks/pull/XXXX)).

  Introduces a value-aware retry loop (`Run`, `RunErr`) and deterministic exponential `BackoffPolicy` that consolidate the multiple retry implementations currently scattered across the provider (`terraform-plugin-sdk/helper/retry`, `databricks-sdk-go/retries`, hand-rolled loops). Each `Run` invocation gets its own retrier instance via a factory, making the loop safe to use from multiple goroutines without locking. Defaults: `Initial=10s`, `Maximum=5m`, `Factor=2`, no jitter. State-driven polling (the dominant pattern across waiters in this provider) is supported natively because the retrier predicate sees both the polled value and any error. This is a no-op for users; callers will be migrated in follow-up changes.

* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).

  The provider only needs identity fields (`userName`, `id`, `externalId`) from `/Me`, never entitlements. Skipping the entitlement computation avoids an expensive `getEffectivePermissions` traversal on the SCIM backend, which has caused incidents on workspaces with large grant counts.
