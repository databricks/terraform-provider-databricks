# NEXT CHANGELOG

## Release v1.135.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `string_value_wo` and `string_value_wo_version` attributes to `databricks_secret` resource ([#5480](https://github.com/databricks/terraform-provider-databricks/pull/5480)).

### Bug Fixes

* Expand the "workspace-level resource requires a workspace_id" error to list all sources checked (resource `provider_config`, provider `workspace_id`, the configured profile, and the `DATABRICKS_WORKSPACE_ID` environment variable) so users know where to set it ([#5763](https://github.com/databricks/terraform-provider-databricks/pull/5763)).

### Documentation

### Exporter

### Internal Changes
