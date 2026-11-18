# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix the plan of an existing Plugin Framework resource, such as `databricks_app`, whose `provider_config` or `provider_config.workspace_id` becomes unknown at plan time, for example when it's taken from another resource. On an account-level provider without `workspace_id`, the plan failed with `managing a workspace-level resource requires a workspace_id, but the previously configured workspace_id was removed`. The plan now leaves the value unknown for apply to resolve and replaces the resource, because the new workspace may differ ([#6228](https://github.com/databricks/terraform-provider-databricks/pull/6228)).

* Fix the plan of `databricks_library`, `databricks_quality_monitor` and `databricks_share` when `provider_config.workspace_id` is unknown at plan time, for example when it's taken from another resource. On an account-level provider without `workspace_id`, the plan failed with `managing a workspace-level resource requires a workspace_id, but none was found`. The plan now leaves the value unknown for apply to resolve, and replaces an existing `databricks_quality_monitor` or `databricks_share` instead of planning an update that failed apply once the value resolved to another workspace ([#6238](https://github.com/databricks/terraform-provider-databricks/pull/6238)).

### Documentation

### Exporter

### Internal Changes
