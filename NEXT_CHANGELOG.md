# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Route SCIM `api = "account"` calls to the workspace-proxied Account SCIM endpoint (`/api/2.0/account/scim/v2/...`) when the provider is configured against a workspace host. The previous behavior emitted `/api/2.0/accounts/{account_id}/scim/v2/...` against the workspace host, which returns 404 — no real production path exercised it. The new behavior honors the Group Manager role per [Databricks docs](https://docs.databricks.com/aws/en/admin/users-groups/manage-groups#manage-groups-using-the-api), letting non-account-admin service principals manage account-group membership from a workspace-scoped provider. The account-host branch is unchanged.

### Documentation

### Exporter

### Internal Changes

* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).

  The provider only needs identity fields (`userName`, `id`, `externalId`) from `/Me`, never entitlements. Skipping the entitlement computation avoids an expensive `getEffectivePermissions` traversal on the SCIM backend, which has caused incidents on workspaces with large grant counts.
