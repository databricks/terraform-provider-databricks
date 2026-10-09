# NEXT CHANGELOG

## Release v1.138.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Retry SCIM requests (`databricks_group`, `databricks_user`, `databricks_service_principal`, `databricks_group_member`, entitlements and roles) that fail with HTTP 409 `Failed due to conflicting concurrent updates. Please retry.`, which happens when many identity changes run in parallel in the same account.

### Documentation

### Exporter

### Internal Changes
