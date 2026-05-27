# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_metastore` so that updating `external_access_enabled` from `true` to `false` is sent in the PATCH request. Previously the field was silently dropped from the request body, so the change never reached the API.

### Documentation

### Exporter

### Internal Changes

* Bump `databricks-sdk-go` to pick up the workspace addressing header migration: workspace-scoped API requests now send `X-Databricks-Workspace-Id` instead of `X-Databricks-Org-Id`. The value still comes from the `workspace_id` provider attribute / `DATABRICKS_WORKSPACE_ID` environment variable. The previous request header continues to be accepted by the platform for rollback safety. The response header read by `CurrentWorkspaceID` on `/api/2.0/preview/scim/v2/Me` is unchanged in this release.
* Add `internal/retrier` package for unified retry and backoff handling ([#5746](https://github.com/databricks/terraform-provider-databricks/pull/5746)).
* Pass `excludedAttributes=entitlements` on SCIM `/Me` requests ([#5725](https://github.com/databricks/terraform-provider-databricks/pull/5725)).
