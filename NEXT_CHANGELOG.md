# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fixed `lifecycle { ignore_changes }` not working for `Optional+Computed` fields whose API returns empty values. The shared `StructToData` gate in `common/reflect_resource.go` was silently dropping empty/nil values for all `Optional` fields, including `Computed` ones, causing perpetual `(known after apply)` diffs and preventing `ignore_changes` from preserving externally-set values.
* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.
||||||| parent of 7a333420 ([Fix] StructToData gate: allow empty values through for Optional+Computed fields)

### Documentation

### Exporter

### Internal Changes

* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).

  The provider only needs identity fields (`userName`, `id`, `externalId`) from `/Me`, never entitlements. Skipping the entitlement computation avoids an expensive `getEffectivePermissions` traversal on the SCIM backend, which has caused incidents on workspaces with large grant counts.
