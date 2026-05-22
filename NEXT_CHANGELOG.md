# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_vector_search_index` hardcoded 15-minute creation timeout: increased default to 75 minutes (consistent with `databricks_vector_search_endpoint`) and made it user-overridable via the `timeouts` block.

### Documentation

### Exporter

### Internal Changes

* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).

  The provider only needs identity fields (`userName`, `id`, `externalId`) from `/Me`, never entitlements. Skipping the entitlement computation avoids an expensive `getEffectivePermissions` traversal on the SCIM backend, which has caused incidents on workspaces with large grant counts.
