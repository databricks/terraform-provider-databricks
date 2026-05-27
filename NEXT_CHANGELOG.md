# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Bump `databricks-sdk-go` to pick up the workspace addressing header migration: workspace-scoped API requests now send `X-Databricks-Workspace-Id` instead of `X-Databricks-Org-Id`. The value still comes from the `workspace_id` provider attribute / `DATABRICKS_WORKSPACE_ID` environment variable. The previous request header continues to be accepted by the platform for rollback safety. The response header read by `CurrentWorkspaceID` on `/api/2.0/preview/scim/v2/Me` is unchanged in this release.
