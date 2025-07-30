---
page_title: "Experimental resource exporter"
---
# Experimental resource exporter

-> This tooling is experimental and provided as is. It has an evolving interface, which may change or be removed in future provider versions.

-> Use the same user who did the exporting to import the exported templates.  Otherwise, it could cause changes in the ownership of the objects.

Generates `*.tf` files for Databricks resources as well as `import.sh` to run import state. It's best used when you need to quickly export Terraform configuration for an existing Databricks workspace. After generating configuration, we strongly recommend manually review all created files.

## Installion
The Resource Exporter is available in your Terraform plugin cache once you have initialised a Terraform workspace that use the Databricks Terraform Provider (`.terraform/providers/registry.terraform.io/databricks/databricks/<provider_version>/<arch>/terraform-provider-databricks_v<provider_version>`).

If not you can also download the [latest released binary](https://github.com/databricks/terraform-provider-databricks/releases), unpack it, and place it in the same folder.

## Example Usage

Resource Exporter can in both interactive and none-interactive mode. 

When running in interactive mode, the Resource Exporter will prompt the user for the [Databricks Workspace URL and a Databricks Workspace PAT](../index.md#authenticating-with-hostname-and-token). It is also possible to authenticate using environment variables.

```bash
./terraform-provider-databricks_v<provider_version> exporter
```

Here's the tool in action:

[![asciicast](https://asciinema.org/a/Rv8ZFJQpfrfp6ggWddjtyXaOy.svg)](https://asciinema.org/a/Rv8ZFJQpfrfp6ggWddjtyXaOy)

-> Please note that in the interactive mode, the selected services are passed as the `-listing` option, not as `-services` option (see below).

The non-interactive mode allows for a more granular selection of services and dependencies. For example, the following command will list all resources related to `jobs` and `compute` services and import them with their dependencies from `groups,secrets,access,compute,users,jobs,storage` services.

```bash
export DATABRICKS_HOST=...
export DATABRICKS_TOKEN=...
./terraform-provider-databricks exporter -skip-interactive \
  -services=groups,secrets,access,compute,users,jobs,storage \
  -listing=jobs,compute \
  -debug
```

The exporter is also supported on the account level for resources that could be defined on an account level. For example, we can export everything defined on the account level:

```sh
export DATABRICKS_HOST=https://accounts.azuredatabricks.net
export DATABRICKS_ACCOUNT_ID=...
./terraform-provider-databricks exporter -skip-interactive
```

Or export only specific resources - users and groups:

```sh
DATABRICKS_HOST=https://accounts.azuredatabricks.net \
  DATABRICKS_ACCOUNT_ID=<uuid>  \
  ./terraform-provider-databricks exporter -directory output \
   -listing groups,users -skip-interactive
```

## Argument Reference

!> This tooling was only extensively tested with administrator privileges.

All arguments are optional, and they tune what code is being generated.

* `-directory` - Path to a directory, where `*.tf` and `import.sh` files would be written. By default, it's set to the current working directory.
* `-module` - Name of module in Terraform state that would affect reference resolution and prefixes for generated commands in `import.sh`.
* `-last-active-days` - Items older than `-last-active-days` won't be imported. By default, the value is set to 3650 (10 years). Has an effect on listing [databricks_cluster](../resources/cluster.md) and [databricks_job](../resources/job.md) resources.
* `-listing` - Comma-separated list of services to be listed and further passed on for importing. For each service specified, the exporter performs a listing of available resources using the `List` function and emits them for importing together with their dependencies. The `-services` parameter could be used to control which transitive dependencies will be also imported.
* `-services` - Comma-separated list of services to import. By default, all services are imported.
* `-match` - Match resource names during listing operation. This filter applies to all resources that are getting listed, so if you want to import all dependencies of just one cluster, specify `-match=autoscaling -listing=compute`. By default, it is empty, which matches everything.
* `-matchRegex` - Match resource names against a given regex during listing operation. Applicable to all resources selected for listing.
* `-excludeRegex` - Exclude resource names matching a given regex. Applied during the listing operation and has higher priority than `-match` and `-matchRegex`.  Applicable to all resources selected for listing.  Could be used to exclude things like `databricks_automl` notebooks, etc.
* `-filterDirectoriesDuringWorkspaceWalking` - if we should apply match logic to directory names when we're performing workspace tree walking.  *Note: be careful with it as it will be applied to all entries, so if you want to filter only specific users, then you will need to specify condition for `/Users` as well, so regex will be `^(/Users|/Users/[a-c].*)$`*.
* `-mounts` - List DBFS mount points, an extremely slow operation that would not trigger unless explicitly specified.
* `-generateProviderDeclaration` - the flag that toggles the generation of `databricks.tf` file with the declaration of the Databricks Terraform provider that is necessary for Terraform versions since Terraform 0.13 (disabled by default).
* `-prefix` - optional prefix that will be added to the name of all exported resources - that's useful for exporting resources from multiple workspaces for merging into a single one.
* `-skip-interactive` - optionally run in a non-interactive mode.
* `-includeUserDomains` - optionally include domain name into generated resource name for `databricks_user` resource.
* `-importAllUsers` - optionally include all users and service principals even if they are only part of the `users` group.
* `-exportDeletedUsersAssets` - optionally include assets of deleted users and service principals.
* `-incremental` - experimental option for incremental export of modified resources and merging with existing resources. *Please note that only a limited set of resources (notebooks, SQL queries/dashboards/alerts, ...) provides information about the last modified date - all other resources will be re-exported again! Also, it's impossible to detect the deletion of many resource types (i.e. clusters, jobs, ...), so you must do periodic full export if resources are deleted! For Workspace objects (notebooks, workspace files, and directories) exporter tries to detect deleted objects and remove them from generated code (requires the presence of `ws_objects.json` file that is written on each export that pulls all workspace objects).  For workspace objects renames are handled as deletion of existing/creation of new resource!*  **Requires** `-updated-since` option if no `exporter-run-stats.json` file exists in the output directory.
* `-updated-since` - timestamp (in ISO8601 format supported by Go language) for exporting of resources modified since a given timestamp. I.e., `2023-07-24T00:00:00Z`. If not specified, the exporter will try to load the last run timestamp from the `exporter-run-stats.json` file generated during the export and use it.
* `-notebooksFormat` - optional format for exporting of notebooks. Supported values are `SOURCE` (default), `DBC`, `JUPYTER`.  This option could be used to export notebooks with embedded dashboards.
* `-noformat` - optionally turn off the execution of `terraform fmt` on the exported files (enabled by default).
* `-debug` - turn on debug output.
* `-trace` - turn on trace output (includes debug level as well).
* `-native-import` - turns on generation of [native import blocks](https://developer.hashicorp.com/terraform/language/import) (requires Terraform 1.5+).  This option is recommended for cases when you want to start managing an existing workspace.
* `-export-secrets` - enables exporting of the secret values - they will be written into the `terraform.tfvars` file.  **Be very careful with this file!**

### Use of `-listing` and `-services` for granular resources selection

The `-listing` option is used to discover resources to export; if it's not specified, then all services are listed (if they have the `List` operation implemented). The `-services` restricts the export of resources only to those resources whose service type is in the list specified by this option.

For example, if we have a job comprising two notebooks and one SQL dashboard, and tasks have Python libraries on DBFS attached. If we just specify the `-listing jobs`, then it will export the following resources:

* job itself
* two notebooks
* directory where notebooks reside
* libraries from DBFS
* SQL dashboard, SQL queries that are used in it, and SQL warehouse that is used to run dashboard/queries
* directory where SQL objects reside
* permissions for all objects above
* user/group information based on permissions, and directories

During code generation, references/dependencies between these objects will be created, and the code will be portable between workspaces.

but if we also specify `-services notebooks,storage` then it will export only:

* job itself
* two notebooks
* directory where notebooks reside
* libraries from DBFS

The rest of the values, like SQL object IDs, etc. will be hard-coded and not portable between workspaces.

You can also use predefined aliases (`all` and `uc`) to specify multiple services at once.  For example, if `-listing` has value `all,-uc`, then we will discover all services except of Unity Catalog + vector search.

We can also exclude specific services  For example, we can specify `-services` as `-all,-uc-tables` and then we won't generate code for `databricks_sql_table`.

### Migration between workspaces with identity federation enabled

When Unity Catalog metastore is attached to a workspace, the Identity Federation is enabled on it.  With Identity Federation users, service principals, and groups are coming from the account level via assignment to a workspace.  But there is still an ability to create workspace-level groups via API and `databricks_group` resource uses it and always creates workspace-level.  As a result, we shouldn't generate resources for account-level groups, because they will be turned into workspace-level groups.  Due to the limitations of APIs we can't use `databricks_permission_assignment` on workspace-level to emulate the assignment.

So migration of resources between two workspaces with Identity Federation enabled should be done in a few steps:

1. On the account level export `databricks_mws_permission_assignment` resources for your source workspace:

  ```sh
  DATABRICKS_CONFIG_PROFILE=<cli-profile> DATABRICKS_ACCOUNT_ID=<account-id> ./terraform-provider-databricks exporter \
    -matchRegex '^<source-workspace-id>$' -listing idfed -services idfed \
    -directory output -skip-interactive -noformat
  ```

2. Replace source workspace ID with destination workspace ID in the generated `idfed.tf` file, i.e. with `sed`:

  ```sh
  sed -ibak -e 's|workspace_id = <source-workspace-id>|workspace_id = <destination-workspace-id>|' idfed.tf
  ```

  and do `terraform apply` on the account level to assign users, service principals, and groups to a destination workspace.

3. Export resources from the source workspace using the exporter on the workspace level. It will automatically detect that Identity Federation is enabled and export account-level objects as data sources instead of resources.

4. Apply exported code against the destination workspace.

## Services

Services are just logical groups of resources used for filtering and organization in files written in `-directory`. All resources are globally sorted by their resource name, which allows you to use generated files for compliance purposes. Nevertheless, managing the entire Databricks workspace with Terraform is the preferred way. Except for notebooks and possibly libraries, which may have their own CI/CD processes.

Services could be specified in combination with predefined aliases (`all` - for all services and listings, `uc` - for all UC services, including the vector search).  The service could be specified as the service name, or it could have `-` prepended to the service, to exclude it from the list (including `-uc` to exclude all UC-related services).

-> Please note that for services not marked with **listing**, we'll export resources only if they are referenced from other resources.

* `access` -  **listing** [databricks_permissions](../resources/permissions.md), [databricks_instance_profile](../resources/instance_profile.md), [databricks_ip_access_list](../resources/ip_access_list.md), and [databricks_access_control_rule_set](../resources/access_control_rule_set.md).   *Please note that for `databricks_permissions` we list only `authorization = "tokens"`, the permissions for other objects (notebooks, ...) will be emitted when corresponding objects are processed!*
* `alerts` - **listing** [databricks_alert](../resources/alert.md).
* `compute` - **listing** [databricks_cluster](../resources/cluster.md).
* `dashboards` - **listing** [databricks_dashboard](../resources/dashboard.md).
* `directories` - **listing** [databricks_directory](../resources/directory.md).  *Please note that directories aren't listed when running in the incremental mode! Only directories with updated notebooks will be emitted.*
* `dlt` - **listing** [databricks_pipeline](../resources/pipeline.md).
* `groups` - **listing** [databricks_group](../data-sources/group.md) with [membership](../resources/group_member.md) and [data access](../resources/group_instance_profile.md).   If Identity Federation is enabled on the workspace (when UC Metastore is attached), then account-level groups are exposed as data sources because they are defined on account level, and only workspace-level groups are exposed as resources.  See the note above on how to perform migration between workspaces with Identity Federation enabled.
* `idfed` - **listing** [databricks_mws_permission_assignment](../resources/mws_permission_assignment.md).  When listing allows filtering assignment only to specific workspace IDs as specified by `-match`, `-matchRegex`, and `-excludeRegex` options.  I.e., to export assignments only for two workspaces, use `-matchRegex '^1688808130562317|5493220389262917$'`.
* `jobs` - **listing** [databricks_job](../resources/job.md). Usually, there are more automated workflows than interactive clusters, so they get their own file in this tool's output.  *Please note that workflows deployed and maintained via [Databricks Asset Bundles](https://docs.databricks.com/en/dev-tools/bundles/index.html) aren't exported!*
* `mlflow-webhooks` - **listing** [databricks_mlflow_webhook](../resources/mlflow_webhook.md).
* `model-serving` - **listing** [databricks_model_serving](../resources/model_serving.md).
* `mounts` - **listing** works only in combination with `-mounts` command-line option.
* `nccs` - **listing** [databricks_mws_network_connectivity_config](../resources/mws_network_connectivity_config.md) and [databricks_mws_ncc_private_endpoint_rule](../resources/mws_ncc_private_endpoint_rule.md).  **Note** we can't export [databricks_mws_ncc_binding](../resources/mws_ncc_binding.md) because of the missing API.
* `notebooks` - **listing** [databricks_notebook](../resources/notebook.md).
* `policies` - **listing** [databricks_cluster_policy](../resources/cluster_policy).
* `pools` - **listing** [instance pools](../resources/instance_pool.md).
* `queries` - **listing** [databricks_query](../resources/query.md).
* `repos` - **listing** [databricks_repo](../resources/repo.md) (both classical Repos in `/Repos` and Git Folders in arbitrary locations).
* `secrets` - **listing** [databricks_secret_scope](../resources/secret_scope.md) along with [keys](../resources/secret.md) and [ACLs](../resources/secret_acl.md).
* `settings` - **listing** [databricks_notification_destination](../resources/notification_destination.md).
* `sql-dashboards` - **listing** Legacy [databricks_sql_dashboard](../resources/sql_dashboard.md) along with associated [databricks_sql_widget](../resources/sql_widget.md) and [databricks_sql_visualization](../resources/sql_visualization.md).
* `sql-endpoints` - **listing** [databricks_sql_endpoint](../resources/sql_endpoint.md).
* `storage` - only [databricks_dbfs_file](../resources/dbfs_file.md) and [databricks_file](../resources/file.md) referenced in other resources (libraries, init scripts, ...) will be downloaded locally and properly arranged into terraform state.
* `uc-artifact-allowlist` - **listing** exports [databricks_artifact_allowlist](../resources/artifact_allowlist.md) resources for Unity Catalog Allow Lists attached to the current metastore.
* `uc-catalogs` - **listing** [databricks_catalog](../resources/catalog.md) and [databricks_workspace_binding](../resources/workspace_binding.md)
* `uc-connections` - **listing** [databricks_connection](../resources/connection.md).  *Please note that because API doesn't return sensitive fields, such as, passwords, tokens, ..., the generated `options` block could be incomplete!*
* `uc-credentials` - **listing** exports [databricks_credential](../resources/credential.md) resources on workspace or account level.  *Please note that it will skip storage credentials! Use the `uc-storage-credentials` service for them*
* `uc-external-locations` - **listing** exports [databricks_external_location](../resources/external_location.md) resource.
* `uc-grants` -  [databricks_grants](../resources/grants.md). *Please note that during export the list of grants is expanded to include the identity that does the export! This is done to allow to creation of objects in case when catalogs/schemas have different owners than the current identity.*.
* `uc-metastores` - **listing** [databricks_metastore](../resources/metastore.md) and [databricks_metastore_assignment](../resource/metastore_assignment.md) (only on account-level).  *Please note that when using workspace-level configuration, only the metastores from the workspace's region are listed!*
* `uc-models` - **listing** (*we can't list directly, only via dependencies to top-level object*) [databricks_registered_model](../resources/registered_model.md)
* `uc-online-tables` - **listing** (*we can't list directly, only via dependencies to top-level object*) [databricks_online_table](../resources/online_table.md)
* `uc-schemas` - **listing** (*we can't list directly, only via dependencies to top-level object*) [databricks_schema](../resources/schema.md)
* `uc-shares` - **listing** [databricks_share](../resources/share.md) and [databricks_recipient](../resources/recipient.md)
* `uc-storage-credentials` - **listing** exports [databricks_storage_credential](../resources/storage_credential.md) resources on workspace or account level.
* `uc-system-schemas` - **listing** exports [databricks_system_schema](../resources/system_schema.md) resources for the UC metastore of the current workspace.
* `uc-tables` - **listing** (*we can't list directly, only via dependencies to top-level object*) [databricks_sql_table](../resources/sql_table.md) resource.
* `uc-volumes` - **listing** (*we can't list directly, only via dependencies to top-level object*) [databricks_volume](../resources/volume.md)
* `users` - **listing** [databricks_user](../resources/user.md) and [databricks_service_principal](../resources/service_principal.md) are written to their own file, simply because of their amount. If Identity Federation is enabled on the workspace (when UC Metastore is attached), then users and service principals are exposed as data sources because they are defined on an account level.  See the note above on how to perform migration between workspaces with Identity Federation enabled.
* `vector-search` - **listing** exports [databricks_vector_search_endpoint](../resources/vector_search_endpoint.md) and [databricks_vector_search_index](../resources/vector_search_index.md)
* `wsconf` - **listing** exports Workspace-level configuration: [databricks_workspace_conf](../resources/workspace_conf.md), [databricks_sql_global_config](../resources/sql_global_config.md) and [databricks_global_init_script](../resources/global_init_script.md).
* `wsfiles` - **listing** [databricks_workspace_file](../resources/workspace_file.md).

## Secrets

For security reasons, [databricks_secret](../resources/secret.md) cannot contain actual plaintext secrets. By default, the exporter will create a variable in `vars.tf`, with the same name as the secret. You are supposed to [fill in the value of the secret](https://blog.gruntwork.io/a-comprehensive-guide-to-managing-secrets-in-your-terraform-code-1d586955ace1#0e7d) after that.  You can use `-export-secrets` command-line option to generate the `terraform.tfvars` file with secret values.

## Parallel execution

To speed up export, Terraform Exporter performs many operations, such as listing & actual data exporting, in parallel using Goroutines.  Built-in defaults control the parallelism, but it's also possible to tune some parameters using environment variables specific to the exporter:

* `EXPORTER_WS_LIST_PARALLELISM` (default: `5`) controls how many Goroutines are used to perform parallel listing of Databricks Workspace objects (notebooks, directories, workspace files, ...).
* `EXPORTER_DIRECTORIES_CHANNEL_SIZE` (default: `300000`) controls the channel's capacity when listing workspace objects. Please ensure that this value is big enough (greater than the number of directories in the workspace; default value should be ok for most cases); otherwise, there is a chance of deadlock.
* `EXPORTER_DEDICATED_RESOUSE_CHANNELS` - by default, only specific resources (`databricks_user`, `databricks_service_principal`, `databricks_group`) have dedicated channels - the rest are handled by the shared channel.  This is done to prevent throttling by specific APIs.  You can override this by providing a comma-separated list of resources as this environment variable.
* `EXPORTER_PARALLELISM_NNN` - number of Goroutines used to process resources of a specific type (replace `NNN` with the exact resource name, for example, `EXPORTER_PARALLELISM_databricks_notebook=10` sets the number of Goroutines for `databricks_notebook` resource to `10`).  There is a shared channel (with name `default`) for handling resources for which there are no dedicated channels - use `EXPORTER_PARALLELISM_default` to increase its size (default size is `15`).   Defaults for some resources are defined by the `goroutinesNumber` map in `exporter/context.go` or equal to `2` if there is no value.  *Don't increase default values too much to avoid REST API throttling!*
* `EXPORTER_DEFAULT_HANDLER_CHANNEL_SIZE` is the size of the shared channel (default: `200000`). You may need to increase it if you have a huge workspace.

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
| [databricks_credential](../resources/credential.md) | Yes | Yes | Yes | No |
| [databricks_dashboard](../resources/dashboard.md) | Yes | No | Yes | No |
| [databricks_dbfs_file](../resources/dbfs_file.md) | Yes | No | Yes | No |
| [databricks_external_location](../resources/external_location.md) | Yes | Yes | Yes | No |
| [databricks_file](../resources/file.md) | Yes | No | Yes | No |
| [databricks_global_init_script](../resources/global_init_script.md) | Yes | Yes | Yes\*\* | No |
| [databricks_grants](../resources/grants.md) | Yes | No | Yes | No |
| [databricks_group](../resources/group.md) | Yes | No | Yes | Yes |
| [databricks_group_instance_profile](../resources/group_instance_profile.md) | Yes | No | Yes | No |
| [databricks_group_member](../resources/group_member.md) | Yes | No | Yes | Yes |
| [databricks_group_role](../resources/group_role.md) | Yes | No | Yes | Yes |
| [databricks_instance_pool](../resources/instance_pool.md) | Yes | No | Yes | No |
| [databricks_instance_profile](../resources/instance_profile.md) | Yes | No | Yes | No |
| [databricks_ip_access_list](../resources/ip_access_list.md) | Yes | Yes | Yes\*\* | No |
| [databricks_job](../resources/job.md) | Yes | No | Yes | No |
| [databricks_library](../resources/library.md) | Yes\* | No | Yes | No |
| [databricks_metastore](../resources/metastore.md) | Yes | Yes | No | Yes |
| [databricks_metastore_assignment](../resources/metastore_assignment.md) | Yes | No | No | Yes |
| [databricks_mlflow_experiment](../resources/mlflow_experiment.md) | No | No | No | No |
| [databricks_mlflow_model](../resources/mlflow_model.md) | No | No | No | No |
| [databricks_mlflow_webhook](../resources/mlflow_webhook.md) | Yes | Yes | Yes | No |
| [databricks_model_serving](../resources/model_serving) | Yes | Yes | Yes | No |
| [databricks_mws_network_connectivity_config](../resources/mws_network_connectivity_config.md) | Yes | Yes | No | Yes |
| [databricks_mws_ncc_binding](../resources/mws_ncc_binding.md) | No | No | No | No |
| [databricks_mws_ncc_private_endpoint_rule](../resources/mws_ncc_private_endpoint_rule.md) | Yes | No | No | Yes |
| [databricks_mws_permission_assignment](../resources/mws_permission_assignment.md) | Yes | No | No | Yes |
| [databricks_notebook](../resources/notebook.md) | Yes | Yes | Yes | No |
| [databricks_notification_destination](../resources/notification_destination.md) | Yes | No | Yes\*\* | No |
| [databricks_obo_token](../resources/obo_token.md) | Not Applicable | No | No | No |
| [databricks_online_table](../resources/online_table.md) | Yes | Yes | Yes | No |
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
| [databricks_vector_search_endpoint](../resources/vector_search_endpoint.md) | Yes | No | Yes | No |
| [databricks_vector_search_index](../resources/vector_search_index.md) | Yes | No | Yes | No |
| [databricks_volume](../resources/volume.md) | Yes | Yes | Yes | No |
| [databricks_workspace_binding](../resources/workspace_binding.md) | Yes | No | Yes | No |
| [databricks_workspace_conf](../resources/workspace_conf.md) | Yes (partial) | No | Yes\*\* | No |
| [databricks_workspace_file](../resources/workspace_file.md) | Yes | Yes | Yes | No |

Notes:

* \* - libraries are exported as blocks inside the cluster definition instead of generating `databricks_library` resources.  This is done to decrease the number of generated resources.
* \*\* - requires workspace admin permission.
