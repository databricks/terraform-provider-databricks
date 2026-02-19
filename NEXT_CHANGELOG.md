# NEXT CHANGELOG

## Release v1.108.0

### Breaking Changes

### New Features and Improvements
* Added support for updating `git_repository` on `databricks_app` resource and switched updates for `description`, `resources`, `budget_policy_id`, `compute_size`, `usage_policy_id`, and `user_api_scopes` to use the async `CreateUpdate` API, as the previous synchronous `Update` API does not support `git_repository` changes and does not handle fields that require async processing ([#5378](https://github.com/databricks/terraform-provider-databricks/pull/5378))

* Added `database_project_name` to `databricks_permissions` for managing Lakebase database project permissions.
* Added `node_type_flexibility` block to `databricks_instance_pool` resource ([#5381](https://github.com/databricks/terraform-provider-databricks/pull/5381)).

### Bug Fixes

* Handle error during WorkspaceClient() creation in databricks_grant and databricks_grants resources ([#5403](https://github.com/databricks/terraform-provider-databricks/pull/5403))

### Documentation

* Added note to `databricks_mws_ncc_binding` that a workspace can only have one NCC binding at a time.

### Exporter

### Internal Changes

* Update Go SDK to v0.110.
