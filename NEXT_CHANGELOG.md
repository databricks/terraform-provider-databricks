# NEXT CHANGELOG

## Release v1.117.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Expand the "workspace-level resource requires a workspace_id" error to list all sources checked (resource `provider_config`, provider `workspace_id`, the configured profile, and the `DATABRICKS_WORKSPACE_ID` environment variable) so users know where to set it ([#5763](https://github.com/databricks/terraform-provider-databricks/pull/5763)).
* Reject `workspace_id` values with a leading zero (e.g. `0470576644108500`) in `provider_config` so the literal value is no longer written to state where it would not match the canonical numeric workspace ID returned by the API ([#5764](https://github.com/databricks/terraform-provider-databricks/pull/5764)).

### Documentation

### Exporter

### Internal Changes
