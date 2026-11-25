# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix the plan of an existing SDKv2 workspace-level resource whose `provider_config.workspace_id` becomes unknown at plan time, for example when it's taken from another resource. On an account-level provider without `workspace_id`, the plan failed with `managing a workspace-level resource requires a workspace_id, but the previously configured workspace_id was removed`. The plan now leaves the value unknown for apply to resolve. On an account-level or unified provider, workspace-scoped resources such as `databricks_directory` and `databricks_job` are now replaced whenever the value is unknown, also when the provider sets `workspace_id`, because the resource may move to another workspace. On a workspace-level provider, which can't change the workspace, they are updated in place. Metastore-scoped resources such as `databricks_volume` are updated in place, except `databricks_online_table`, `databricks_workspace_binding` and `databricks_metastore_data_access`, which are replaced because they can't be updated ([#6227](https://github.com/databricks/terraform-provider-databricks/pull/6227)).

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
