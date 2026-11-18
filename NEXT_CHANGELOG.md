# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix the plan of an existing Plugin Framework resource, such as `databricks_app`, whose `provider_config` or `provider_config.workspace_id` becomes unknown at plan time, for example when it's taken from another resource. On an account-level provider without `workspace_id`, the plan failed with `managing a workspace-level resource requires a workspace_id, but the previously configured workspace_id was removed`. The plan now leaves the value unknown for apply to resolve and replaces the resource, because the new workspace may differ ([#6228](https://github.com/databricks/terraform-provider-databricks/pull/6228)).

### Documentation

### Exporter

### Internal Changes
