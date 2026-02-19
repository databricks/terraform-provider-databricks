# NEXT CHANGELOG

## Release v1.108.0

### Breaking Changes

### New Features and Improvements

* Added `database_project_name` to `databricks_permissions` for managing Lakebase database project permissions.
* Added `node_type_flexibility` block to `databricks_instance_pool` resource ([#5381](https://github.com/databricks/terraform-provider-databricks/pull/5381)).

### Bug Fixes

* Fixed `databricks_cluster` to preserve externally-set `spark_env_vars` (e.g. from cluster policies) during updates when not configured in Terraform.
* Handle error during WorkspaceClient() creation in databricks_grant and databricks_grants resources ([#5403](https://github.com/databricks/terraform-provider-databricks/pull/5403))

### Documentation

* Added note to `databricks_mws_ncc_binding` that a workspace can only have one NCC binding at a time.

### Exporter

### Internal Changes
