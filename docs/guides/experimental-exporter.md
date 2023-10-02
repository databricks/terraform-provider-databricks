---
page_title: "Experimental resource exporter"
---
# Experimental resource exporter

-> **Note** This tooling is experimental and provided as is. It has an evolving interface, which may change or be removed in future versions of the provider.

-> **Note** Use the same user who did the exporting to import the exported templates.  Otherwise it could cause changes in the jobs ownership.

Generates `*.tf` files for Databricks resources as well as `import.sh` to run import state. Available as part of provider binary. The only possible way to authenticate is through [environment variables](../index.md#Environment-variables). It's best used when you need to quickly export Terraform configuration for an existing Databricks workspace. After generating configuration, we strongly recommend manually review all created files.

## Example Usage

After downloading the [latest released binary](https://github.com/databricks/terraform-provider-databricks/releases), unpack it and place it in the same folder. In fact, you may have already downloaded this binary - check the `.terraform` folder of any state directory, where you've used the `databricks` provider. It could also be in your plugin cache `~/.terraform.d/plugins/registry.terraform.io/databricks/databricks/*/*/terraform-provider-databricks`. Here's the tool in action:

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

All arguments are optional and they tune what code is being generated.

* `-directory` - Path to directory, where `*.tf` and `import.sh` files would be written. By default it's set to the current working directory.
* `-module` - Name of module in Terraform state, that would affect reference resolution and prefixes for generated commands in `import.sh`.
* `-last-active-days` - Items older than `-last-active-days` won't be imported. By default the value is set to 3650 (10 years). Has an effect on listing [databricks_cluster](../resources/cluster.md) and [databricks_job](../resources/job.md) resources.
* `-services` - Comma-separated list of services to import. By default all services are imported. 
* `-listing` - Comma-separated list of services to be listed and further passed on for importing. `-services` parameter controls which transitive dependencies will be processed. We recommend limiting with `-listing` more often, than with `-services`.
* `-match` - Match resource names during listing operation. This filter applies to all resources that are getting listed, so if you want to import all dependencies of just one cluster, specify `-match=autoscaling -listing=compute`. By default it is empty, which matches everything.
* `-mounts` - List DBFS mount points, which is an extremely slow operation and would not trigger unless explicitly specified.
* `-generateProviderDeclaration` - flag that toggles generation of `databricks.tf` file with declaration of the Databricks Terraform provider that is necessary for Terraform versions since Terraform 0.13 (disabled by default).
* `-prefix` - optional prefix that will be added to the name of all exported resources - that's useful for exporting resources multiple workspaces for merging into a single one.
* `-skip-interactive` - optionally run in a non-interactive mode.
* `-includeUserDomains` - optionally include domain name into generated resource name for `databricks_user` resource.
* `-importAllUsers` - optionally include all users and service principals even if they only part of the `users` group.
* `-incremental` - experimental option for incremental export of modified resources and merging with existing resources. *Please note that only limited set of resources (notebooks, SQL queries/dashboards/alerts, ...) provides information about last modified date - all other resources will be re-exported again! Also, it's not possible to detect deletion of the resources, so you will need to do periodic full export if resources are deleted!*   **Requires** `-updated-since` option if no `exporter-run-stats.json` file exists in the output directory.
* `-updated-since` - timestamp (in ISO8601 format supported by Go language) for exporting of resources modified since a giving timestamp. I.e. `2023-07-24T00:00:00Z`. If not specified, exporter will try to load last run timestamp from the `exporter-run-stats.json` file generated during the export, and use it.
* `-notebooksFormat` - optional format for exporting of notebooks. Supported values are `SOURCE` (default), `DBC`, `JUPYTER`.  This could be used to export of notebooks with embedded dashboards.

## Services

Services are just logical groups of resources used for filtering and organization in files written in `-directory`. All resources are globally sorted by their resource name, which technically allows you to use generated files for compliance purposes. Nevertheless, managing the entire Databricks workspace with Terraform is the prefered way. With the exception of notebooks and possibly libraries, which may have their own CI/CD processes.

* `access` - [databricks_permissions](../resources/permissions.md), [databricks_instance_profile](../resources/instance_profile.md) and [databricks_ip_access_list](../resources/ip_access_list.md).
* `compute` - **listing** [databricks_cluster](../resources/cluster.md). Includes [cluster policies](../resources/cluster_policy.md).
* `directories` - **listing** [databricks_directory](../resources/directory.md)
* `dlt` - **listing** [databricks_pipeline](../resources/pipeline.md)
* `groups` - [databricks_group](../data-sources/group.md) with [membership](../resources/group_member.md) and [data access](../resources/group_instance_profile.md).
* `jobs` - **listing** [databricks_job](../resources/job.md). Usually there are more automated jobs than interactive clusters, so they get their own file in this tool's output.
* `mlflow-webhooks` - **listing** [databricks_mlflow_webhook](../resources/mlflow_webhook.md).
* `model-serving` - **listing** [databricks_model_serving](../resources/model_serving.md).
* `mounts` - **listing** works only in combination with `-mounts` command-line option.
* `notebooks` - **listing** [databricks_notebook](../resources/notebook.md) and [databricks_workspace_file](../resources/workspace_file.md)
* `pools` - **listing** [instance pools](../resources/instance_pool.md).
* `repos` - **listing** [databricks_repo](../resources/repo.md)
* `secrets` - **listing** [databricks_secret_scope](../resources/secret_scope.md) along with [keys](../resources/secret.md) and [ACLs](../resources/secret_acl.md). 
* `sql-alerts` - **listing** [databricks_sql_alert](../resources/sql_alert.md).
* `sql-dashboards` - **listing** [databricks_sql_dashboard](../resources/sql_dashboard.md) along with associated [databricks_sql_widget](../resources/sql_widget.md) and [databricks_sql_visualization](../resources/sql_visualization.md)
* `sql-dashboards` - **listing** [databricks_sql_dashboard](../resources/sql_dashboard.md) along with associated [databricks_sql_widget](../resources/sql_widget.md) and [databricks_sql_visualization](../resources/sql_visualization.md).
* `sql-endpoints` - **listing** [databricks_sql_endpoint](../resources/sql_endpoint.md) along with [databricks_sql_global_config](../resources/sql_global_config.md)
* `sql-queries` - **listing** [databricks_sql_query](../resources/sql_query.md)
* `storage` - any referenced [databricks_dbfs_file](../resources/dbfs_file.md) will be downloaded locally and properly arranged into terraform state.
* `users` - [databricks_user](../resources/user.md) and [databricks_service_principal](../resources/service_principal.md) are written to their own file, simply because of their amount. If you use SCIM provisioning, the only use-case for importing `users` service is to migrate workspaces.
* `workspace` - [databricks_workspace_conf](../resources/workspace_conf.md) and [databricks_global_init_script](../resources/global_init_script.md)

## Secrets

For security reasons, [databricks_secret](../resources/secret.md) cannot contain actual plaintext secrets. Importer will create a variable in `vars.tf`, that would have the same name as secret. You are supposed to [fill in the value of the secret](https://blog.gruntwork.io/a-comprehensive-guide-to-managing-secrets-in-your-terraform-code-1d586955ace1#0e7d) after that.

## Parallel execution

To speedup export, Terraform Exporter performs many operations, such as, listing & actual data exporting, in parallel using Goroutines.  There are built-in defaults controlling the parallelism, but it's also possible to tune some parameters using environment variables specific to exporter:

* `EXPORTER_WS_LIST_PARALLELISM` (default: `5`) controls how many Goroutines are used to perform parallel listing of Databricks Workspace objects (notebooks, directories, workspace files, ...).
* `EXPORTER_DIRECTORIES_CHANNEL_SIZE` (default: `100000`) controls capacity of the channel that is used when listing workspace objects. Please make sure that this value is big enough (default value should be ok), otherwise there is a chance of deadlock. 
* `EXPORTER_PARALLELISM_NNN` - number of Goroutines used to process resources of specific type (replace `NNN` with resource name, for example, `EXPORTER_PARALLELISM_databricks_notebook=10` sets number of Goroutines for `databricks_notebook` resource to `10`).  Defaults for some resources are defined by the `goroutinesNumber` map in `exporter/context.go`, or equal to `2` if there is no value there.  *Don't increase default values too much to avoid REST API throttling!*


## Support Matrix

Exporter aims to generate HCL code for the most of resources within the Databricks workspace:

| Resource | Generated code | Incremental |
| --- | --- | --- |
| [databricks_cluster](../resources/cluster.md) | Yes | No |
| [databricks_cluster_policy](../resources/cluster_policy.md) | Yes | No |
| [databricks_dbfs_file](../resources/dbfs_file.md) | Yes | No |
| [databricks_global_init_script](../resources/global_init_script.md) | Yes | Yes |
| [databricks_group](../resources/group.md) | Yes | No |
| [databricks_group_instance_profile](../resources/group_instance_profile.md) | Yes | No |
| [databricks_group_member](../resources/group_member.md) | Yes | No |
| [databricks_group_role](../resources/group_role.md) | Yes | No |
| [databricks_instance_pool](../resources/instance_pool.md) | Yes | No |
| [databricks_instance_profile](../resources/instance_profile.md) | Yes | No |
| [databricks_ip_access_list](../resources/ip_access_list.md) | Yes | Yes |
| [databricks_job](../resources/job.md) | Yes | No |
| [databricks_library](../resources/library.md) | Yes | No |
| [databricks_mlflow_model](../resources/mlflow_model.md) | No | No |
| [databricks_mlflow_experiment](../resources/mlflow_experiment.md) | No | No |
| [databricks_mlflow_webhook](../resources/mlflow_webhook.md) | Yes | Yes |
| [databricks_model_serving](../resources/model_serving) | Yes | Yes |
| [databricks_notebook](../resources/notebook.md) | Yes | Yes |
| [databricks_obo_token](../resources/obo_token.md) | Not Applicable | No |
| [databricks_permissions](../resources/permissions.md) | Yes | No |
| [databricks_pipeline](../resources/pipeline.md) | Yes | Yes |
| [databricks_repo](../resources/repo.md) | Yes | No |
| [databricks_secret](../resources/secret.md) | Yes | No |
| [databricks_secret_acl](../resources/secret_acl.md) | Yes | No |
| [databricks_secret_scope](../resources/secret_scope.md) | Yes | No |
| [databricks_service_principal](../resources/service_principal.md) | Yes | No |
| [databricks_service_principal_role](../resources/service_principal_role.md) | Yes | No |
| [databricks_sql_alert](../resources/sql_alert.md) | Yes | Yes |
| [databricks_sql_dashboard](../resources/sql_dashboard.md) | Yes | Yes |
| [databricks_sql_endpoint](../resources/sql_endpoint.md) | Yes | No |
| [databricks_sql_global_config](../resources/sql_global_config.md) | Yes | No |
| [databricks_sql_permissions](../resources/sql_permissions.md) | No | No |
| [databricks_sql_query](../resources/sql_query.md) | Yes | Yes |
| [databricks_sql_visualization](../resources/sql_visualization.md) | Yes | Yes |
| [databricks_sql_widget](../resources/sql_widget.md) | Yes | Yes |
| [databricks_token](../resources/token.md) | Not Applicable | No |
| [databricks_user](../resources/user.md) | Yes | No |
| [databricks_user_instance_profile](../resources/user_instance_profile.md) | No (Deprecated) | No |
| [databricks_user_role](../resources/user_role.md) | Yes | No |
| [databricks_workspace_conf](../resources/workspace_conf.md) | Yes (partial) | No |
| [databricks_workspace_file](../resources/workspace_file.md) | Yes | Yes |

