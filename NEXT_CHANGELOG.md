# NEXT CHANGELOG

## Release v1.122.0

### Important Changes

* Removed plan-time `workspace_id` reachability validation for unified-provider
  resources ([#PRNUM](https://github.com/databricks/terraform-provider-databricks/pull/PRNUM)).

  `workspace_id` reachability and mismatch errors (e.g. `workspace_id mismatch`,
  `managing workspace-level resources requires a workspace_id`) are now reported
  when the resource is applied instead of during `terraform plan`. Plans no longer
  make workspace-resolution API calls for these checks. This fixes false-positive
  plan failures for principals that can manage a resource but cannot call the
  workspace `/Me` endpoint, and for `provider_config.workspace_id` values that are
  only known after apply (for example a `workspace_id` sourced from another
  resource that is created in the same run). For newly-created workspace-level
  resources this removes plan-time workspace API calls entirely; existing
  resources still refresh at plan (an unavoidable read), which validates the
  workspace they already live in.

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes

* A provider-level `workspace_id` is now consistently validated when a
  workspace-level resource acquires its client, regardless of whether the value
  came from the resource's `provider_config` block or the provider configuration
  ([#PRNUM](https://github.com/databricks/terraform-provider-databricks/pull/PRNUM)).

  Previously a workspace-level provider silently ignored a mismatched
  provider-level `workspace_id` when a resource omitted `provider_config`. It now
  surfaces a `workspace_id mismatch` error at apply, matching the behavior when
  `provider_config.workspace_id` is set explicitly. For the common case where the
  provider-level `workspace_id` matches the configured workspace (including when
  it is auto-resolved from host metadata) there is no change.
