---
page_title: "Experimental resource exporter"
---
# Experimental resource exporter

-> **Note** This tooling is experimental and provided as is. It has an evolving interface, which may change or be removed in future provider versions.

-> **Note** Use the same user who did the exporting to import the exported templates.  Otherwise, it could cause changes in the ownership of the objects.

Generates `*.tf` files for Databricks resources together with `import.sh` that is used to import objects into the Terraform state. Available as part of provider binary. The only way to authenticate is through [environment variables](../index.md#Environment-variables). It's best used when you need to export Terraform configuration for an existing Databricks workspace quickly. After generating the configuration, we strongly recommend manually reviewing all created files.

## Example Usage

After downloading the [latest released binary](https://github.com/databricks/terraform-provider-databricks/releases), unpack it and place it in the same folder. You may have already downloaded this binary - check the `.terraform` folder of any state directory where you've used the `databricks` provider. It could also be in your plugin cache `~/.terraform.d/plugins/registry.terraform.io/databricks/databricks/*/*/terraform-provider-databricks`. Here's the tool in action:

[![asciicast](https://asciinema.org/a/Rv8ZFJQpfrfp6ggWddjtyXaOy.svg)](https://asciinema.org/a/Rv8ZFJQpfrfp6ggWddjtyXaOy)

Exporter can also be used in a non-interactive mode:

```bash
export DATABRICKS_HOST=...
export DATABRICKS_TOKEN=...
./terraform-provider-databricks exporter -skip-interactive \
 -services=groups,secrets,access,compute,users,jobs,storage \
 -listing=jobs,compute \
 -last-active-days=90 \
 -debug
```

## Argument Reference

!> **Warning** This tooling was only extensively tested with administrator privileges.

All arguments are optional, and they tune what code is being generated.

* `-directory` - Path to a directory, where `*.tf` and `import.sh` files would be written. By default, it's set to the current working directory.
* `-module` - Name of module in Terraform state that would affect reference resolution and prefixes for generated commands in `import.sh`.
* `-last-active-days` - Items older than `-last-active-days` won't be imported. By default, the value is set to 3650 (10 years). Has an effect on listing [databricks_cluster](../resources/cluster.md) and [databricks_job](../resources/job.md) resources.
* `-services` - Comma-separated list of services to import. By default, all services are imported.
* `-listing` - Comma-separated list of services to be listed and further passed on for importing. `-services` parameter controls which transitive dependencies will be processed. We recommend limiting with `-listing` more often than with `-services`.
* `-match` - Match resource names during listing operation. This filter applies to all resources that are getting listed, so if you want to import all dependencies of just one cluster, specify `-match=autoscaling -listing=compute`. By default, it is empty, which matches everything.
* `-mounts` - List DBFS mount points, an extremely slow operation that would not trigger unless explicitly specified.
* `-generateProviderDeclaration` - the flag that toggles the generation of `databricks.tf` file with the declaration of the Databricks Terraform provider that is necessary for Terraform versions since Terraform 0.13 (disabled by default).
* `-prefix` - optional prefix that will be added to the name of all exported resources - that's useful for exporting resources from multiple workspaces for merging into a single one.
* `-skip-interactive` - optionally run in a non-interactive mode.
* `-includeUserDomains` - optionally include domain name into generated resource name for `databricks_user` resource.
* `-importAllUsers` - optionally include all users and service principals even if they are only part of the `users` group.
* `-exportDeletedUsersAssets` - optionally include assets of deleted users and service principals.
* `-incremental` - experimental option for incremental export of modified resources and merging with existing resources. *Please note that only a limited set of resources (notebooks, SQL queries/dashboards/alerts, ...) provides information about the last modified date - all other resources will be re-exported again! Also, it's impossible to detect the deletion of many resource types (i.e. clusters, jobs, ...), so you must do periodic full export if resources are deleted! For Workspace objects (notebooks, workspace files and directories) exporter tries to detect deleted objects and remove them from generated code (requires presence of `ws_objects.json` file that is written on each export that pulls all workspace objects).  For workspace objects renames are handled as deletion of existing/creation of new resource!*  **Requires** `-updated-since` option if no `exporter-run-stats.json` file exists in the output directory.
* `-updated-since` - timestamp (in ISO8601 format supported by Go language) for exporting of resources modified since a given timestamp. I.e., `2023-07-24T00:00:00Z`. If not specified, the exporter will try to load the last run timestamp from the `exporter-run-stats.json` file generated during the export and use it.
* `-notebooksFormat` - optional format for exporting of notebooks. Supported values are `SOURCE` (default), `DBC`, `JUPYTER`.  This option could be used to export notebooks with embedded dashboards.
* `-noformat` - optionally turn off the execution of `terraform fmt` on the exported files (enabled by default).
* `-debug` - turn on debug output.
* `-trace` - turn on trace output (includes debug level as well).
* `-native-import` - turns on generation of [native import blocks](https://developer.hashicorp.com/terraform/language/import) (requires Terraform 1.5+).  This option is recommended for cases when you want to start to manage existing workspace.
* `-export-secrets` - enables export of the secret values - they will be written into the `terraform.tfvars` file.  **Be very careful with this file!**

## Services

Services are just logical groups of resources used for filtering and organization in files written in `-directory`. All resources are globally sorted by their resource name, which allows you to use generated files for compliance purposes. Nevertheless, managing the entire Databricks workspace with Terraform is the preferred way. Except for notebooks and possibly libraries, which may have their own CI/CD processes.

-> **Note**
  Please note that for services not marked with **listing**, we'll export resources only if they are referenced from other resources.

* `access` - [databricks_permissions](../resources/permissions.md), [databricks_instance_profile](../resources/instance_profile.md) and [databricks_ip_access_list](../resources/ip_access_list.md).
* `compute` - **listing** [databricks_cluster](../resources/cluster.md).
* `directories` - **listing** [databricks_directory](../resources/directory.md).
* `dlt` - **listing** [databricks_pipeline](../resources/pipeline.md).
* `groups` - **listing** [databricks_group](../data-sources/group.md) with [membership](../resources/group_member.md) and [data access](../resources/group_instance_profile.md).
* `jobs` - **listing** [databricks_job](../resources/job.md). Usually, there are more automated jobs than interactive clusters, so they get their own file in this tool's output.
* `mlflow-webhooks` - **listing** [databricks_mlflow_webhook](../resources/mlflow_webhook.md).
* `model-serving` - **listing** [databricks_model_serving](../resources/model_serving.md).
* `mounts` - **listing** works only in combination with `-mounts` command-line option.
* `notebooks` - **listing** [databricks_notebook](../resources/notebook.md) and [databricks_workspace_file](../resources/workspace_file.md).
* `policies` - **listing** [databricks_cluster_policy](../resources/cluster_policy).
* `pools` - **listing** [instance pools](../resources/instance_pool.md).
* `repos` - **listing** [databricks_repo](../resources/repo.md)
* `secrets` - **listing** [databricks_secret_scope](../resources/secret_scope.md) along with [keys](../resources/secret.md) and [ACLs](../resources/secret_acl.md).
* `sql-alerts` - **listing** [databricks_sql_alert](../resources/sql_alert.md).
* `sql-dashboards` - **listing** [databricks_sql_dashboard](../resources/sql_dashboard.md) along with associated [databricks_sql_widget](../resources/sql_widget.md) and [databricks_sql_visualization](../resources/sql_visualization.md).
* `sql-endpoints` - **listing** [databricks_sql_endpoint](../resources/sql_endpoint.md) along with [databricks_sql_global_config](../resources/sql_global_config.md).
* `sql-queries` - **listing** [databricks_sql_query](../resources/sql_query.md).
* `storage` - only [databricks_dbfs_file](../resources/dbfs_file.md) and [databricks_file](../resources/file.md) referenced in other resources (libraries, init scripts, ...) will be downloaded locally and properly arranged into terraform state.
* `uc-artifact-allowlist` - **listing** exports [databricks_artifact_allowlist](../resources/artifact_allowlist.md) resources for Unity Catalog Allow Lists attached to the current metastore.
* `uc-catalogs` - **listing** [databricks_catalog](../resources/catalog.md) and [databricks_catalog_workspace_binding](../resources/catalog_workspace_binding.md)
* `uc-connections` - **listing** [databricks_connection](../resources/connection.md).  *Please note that because API doesn't return sensitive fields, such as, passwords, tokens, ..., the generated `options` block could be incomplete!*
* `uc-external-locations` - **listing** exports [databricks_external_location](../resources/external_location.md) resource.
* `uc-grants` -  [databricks_grants](../resources/grants.md). *Please note that during export the list of grants is expanded to include the identity that does the export! This is done to allow to create objects in case when catalogs/schemas have different owners than current identity.*.
* `uc-metastores` - **listing** [databricks_metastore](../resources/metastore.md) and [databricks_metastore_assignment](../resource/metastore_assignment.md) (only on account-level).  *Please note that when using workspace-level configuration, only metastores from the workspace's region are listed!*
* `uc-models` - [databricks_registered_model](../resources/registered_model.md)
* `uc-schemas` -  [databricks_schema](../resources/schema.md)
* `uc-shares` - **listing** [databricks_share](../resources/share.md) and [databricks_recipient](../resources/recipient.md)
* `uc-storage-credentials` - **listing** exports [databricks_storage_credential](../resources/storage_credential.md) resources on workspace or account level.
* `uc-system-schemas` - **listing** exports [databricks_system_schema](../resources/system_schema.md) resources for the UC metastore of the current workspace.
* `uc-tables` - [databricks_sql_table](../resources/sql_table.md) resource.
* `uc-volumes` -  [databricks_volume](../resources/volume.md)
* `users` - [databricks_user](../resources/user.md) and [databricks_service_principal](../resources/service_principal.md) are written to their own file, simply because of their amount. If you use SCIM provisioning, migrating workspaces is the only use case for importing `users` service.
* `workspace` - **listing** [databricks_workspace_conf](../resources/workspace_conf.md) and [databricks_global_init_script](../resources/global_init_script.md)

## Secrets

For security reasons, [databricks_secret](../resources/secret.md) cannot contain actual plaintext secrets. By default importer will create a variable in `vars.tf`, with the same name as the secret. You are supposed to [fill in the value of the secret](https://blog.gruntwork.io/a-comprehensive-guide-to-managing-secrets-in-your-terraform-code-1d586955ace1#0e7d) after that.  You can use `-export-secrets` command-line option to generate the `terraform.tfvars` file with secret values.

## Parallel execution

To speed up export, Terraform Exporter performs many operations, such as listing & actual data exporting, in parallel using Goroutines.  Built-in defaults are controlling the parallelism, but it's also possible to tune some parameters using environment variables specific to the exporter:

* `EXPORTER_WS_LIST_PARALLELISM` (default: `5`) controls how many Goroutines are used to perform parallel listing of Databricks Workspace objects (notebooks, directories, workspace files, ...).
* `EXPORTER_DIRECTORIES_CHANNEL_SIZE` (default: `100000`) controls the channel's capacity when listing workspace objects. Please ensure that this value is big enough (greater than the number of directories in the workspace; default value should be ok for most cases); otherwise, there is a chance of deadlock.
* `EXPORTER_DEDICATED_RESOUSE_CHANNELS` - by default, only specific resources (`databricks_user`, `databricks_service_principal`, `databricks_group`) have dedicated channels - the rest are handled by the shared channel.  This is done to prevent throttling by specific APIs.  You can override this by providing a comma-separated list of resources as this environment variable.
* `EXPORTER_PARALLELISM_NNN` - number of Goroutines used to process resources of a specific type (replace `NNN` with the exact resource name, for example, `EXPORTER_PARALLELISM_databricks_notebook=10` sets the number of Goroutines for `databricks_notebook` resource to `10`).  There is a shared channel (with name `default`) for handling of resources for which there are no dedicated channels - use `EXPORTER_PARALLELISM_default` to increase it's size (default size is `15`).   Defaults for some resources are defined by the `goroutinesNumber` map in `exporter/context.go` or equal to `2` if there is no value.  *Don't increase default values too much to avoid REST API throttling!*
* `EXPORTER_DEFAULT_HANDLER_CHANNEL_SIZE` - the size of the shared channel (default: `200000`) - you may need to increase it if you have a huge workspace.

## Support Matrix

Exporter aims to generate HCL code for most of the resources within the Databricks workspace:

| Resource | Supported | Incremental | Workspace | Account |
| --- | --- | --- | --- | --- |
| [databricks_access_control_rule_set](../resources/access_control_rule_set.md) | Yes | No | No | Yes |
| [databricks_artifact_allowlist](../resources/artifact_allowlist.md) | Yes | No | Yes | No |
| [databricks_catalog](../resources/catalog.md) | Yes | Yes | Yes | No |
| [databricks_cluster](../resources/cluster.md) | Yes | No | Yes | No |
| [databricks_cluster_policy](../resources/cluster_policy.md) | Yes | No | Yes | No |
| [databricks_connection](../resources/connection.md) | Yes | Yes | Yes | No |
| [databricks_dbfs_file](../resources/dbfs_file.md) | Yes | No | Yes | No |
| [databricks_external_location](../resources/external_location.md) | Yes | Yes | Yes | No |
| [databricks_file](../resources/file.md) | Yes | No | Yes | No |
| [databricks_global_init_script](../resources/global_init_script.md) | Yes | Yes | Yes | No |
| [databricks_grants](../resources/grants.md) | Yes | No | Yes | No |
| [databricks_group](../resources/group.md) | Yes | No | Yes | Yes |
| [databricks_group_instance_profile](../resources/group_instance_profile.md) | Yes | No | Yes | No |
| [databricks_group_member](../resources/group_member.md) | Yes | No | Yes | Yes |
| [databricks_group_role](../resources/group_role.md) | Yes | No | Yes | Yes |
| [databricks_instance_pool](../resources/instance_pool.md) | Yes | No | Yes | No |
| [databricks_instance_profile](../resources/instance_profile.md) | Yes | No | Yes | No |
| [databricks_ip_access_list](../resources/ip_access_list.md) | Yes | Yes | Yes | No |
| [databricks_job](../resources/job.md) | Yes | No | Yes | No |
| [databricks_library](../resources/library.md) | Yes\* | No | Yes | No |
| [databricks_metastore](../resources/metastore.md) | Yes | Yes | No | Yes |
| [databricks_metastore_assignment](../resources/metastore_assignment.md) | Yes | No | No | Yes |
| [databricks_mlflow_experiment](../resources/mlflow_experiment.md) | No | No | No | No |
| [databricks_mlflow_model](../resources/mlflow_model.md) | No | No | No | No |
| [databricks_mlflow_webhook](../resources/mlflow_webhook.md) | Yes | Yes | Yes | No |
| [databricks_model_serving](../resources/model_serving) | Yes | Yes | Yes | No |
| [databricks_notebook](../resources/notebook.md) | Yes | Yes | Yes | No |
| [databricks_obo_token](../resources/obo_token.md) | Not Applicable | No | No | No |
| [databricks_permissions](../resources/permissions.md) | Yes | No | Yes | No |
| [databricks_pipeline](../resources/pipeline.md) | Yes | Yes | Yes | No |
| [databricks_recipient](../resources/recipient.md) | Yes | Yes | Yes | No |
| [databricks_registered_model](../resources/registered.md) | Yes | Yes | Yes | No |
| [databricks_repo](../resources/repo.md) | Yes | No | Yes | No |
| [databricks_schema](../resources/schema.md) | Yes | Yes | Yes | No |
| [databricks_secret](../resources/secret.md) | Yes | No | Yes | No |
| [databricks_secret_acl](../resources/secret_acl.md) | Yes | No | Yes | No |
| [databricks_secret_scope](../resources/secret_scope.md) | Yes | No | Yes | No |
| [databricks_service_principal](../resources/service_principal.md) | Yes | No | Yes | Yes |
| [databricks_service_principal_role](../resources/service_principal_role.md) | Yes | No | Yes | Yes |
| [databricks_share](../resources/share.md) | Yes | Yes | Yes | No |
| [databricks_sql_alert](../resources/sql_alert.md) | Yes | Yes | Yes | No |
| [databricks_sql_dashboard](../resources/sql_dashboard.md) | Yes | Yes | Yes | No |
| [databricks_sql_endpoint](../resources/sql_endpoint.md) | Yes | No | Yes | No |
| [databricks_sql_global_config](../resources/sql_global_config.md) | Yes | No | Yes | No |
| [databricks_sql_permissions](../resources/sql_permissions.md) | No | No | Yes | No |
| [databricks_sql_query](../resources/sql_query.md) | Yes | Yes | Yes | No |
| [databricks_sql_table](../resources/sql_table.md) | Yes | Yes | Yes | No |
| [databricks_sql_visualization](../resources/sql_visualization.md) | Yes | Yes | Yes | No |
| [databricks_sql_widget](../resources/sql_widget.md) | Yes | Yes | Yes | No |
| [databricks_storage_credential](../resources/storage_credential.md) | Yes | Yes | Yes | No |
| [databricks_system_schema](../resources/system_schema.md) | Yes | No | Yes | No |
| [databricks_token](../resources/token.md) | Not Applicable | No | Yes | No |
| [databricks_user](../resources/user.md) | Yes | No | Yes | Yes |
| [databricks_user_instance_profile](../resources/user_instance_profile.md) | No | No | No | No |
| [databricks_user_role](../resources/user_role.md) | Yes | No | Yes | Yes |
| [databricks_volume](../resources/volume.md) | Yes | Yes | Yes | No |
| [databricks_workspace_conf](../resources/workspace_conf.md) | Yes (partial) | No | Yes | No |
| [databricks_workspace_file](../resources/workspace_file.md) | Yes | Yes | Yes | No |

Notes:

* \* - libraries are exported as blocks inside the cluster definition instead of generating `databricks_library` resources.  This is done to decrease the number of generated resources.
