# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

 * Added `warehouse_id` support to `databricks_sql_permissions` resource, allowing `GRANT`/`REVOKE`/`SHOW GRANT` to execute via a SQL warehouse instead of a cluster.

### Bug Fixes

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.

### Documentation

### Exporter

### Internal Changes

* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).

  The provider only needs identity fields (`userName`, `id`, `externalId`) from `/Me`, never entitlements. Skipping the entitlement computation avoids an expensive `getEffectivePermissions` traversal on the SCIM backend, which has caused incidents on workspaces with large grant counts.
