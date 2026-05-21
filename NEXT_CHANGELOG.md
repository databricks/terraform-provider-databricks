# NEXT CHANGELOG

## Release v1.116.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* Bump `databricks-sdk-go` to pick up the workspace addressing header migration: workspace-scoped API calls now send `X-Databricks-Workspace-Id` instead of `X-Databricks-Org-Id`. The value still comes from the `workspace_id` provider attribute / `DATABRICKS_WORKSPACE_ID` environment variable and now accepts either a classic numeric workspace ID or another workspace identifier format that the platform understands. The previous header continues to be accepted by the platform for rollback safety.
