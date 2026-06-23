# NEXT CHANGELOG

## Release v1.119.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Expand the "workspace-level resource requires a workspace_id" error to list all sources checked (resource `provider_config`, provider `workspace_id`, the configured profile, and the `DATABRICKS_WORKSPACE_ID` environment variable) so users know where to set it ([#5763](https://github.com/databricks/terraform-provider-databricks/pull/5763)).

### Documentation

* Document that read-only workspace bindings aren't applicable for non-catalog objects ([#5611](https://github.com/databricks/terraform-provider-databricks/pull/5611))

### Exporter

### Internal Changes
