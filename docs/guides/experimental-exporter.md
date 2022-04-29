---
page_title: "Experimental resource exporter"
---
# Experimental resource exporter

-> **Note** This tooling is experimental and provided as is. It has an evolving interfaces, which may change or be removed in future versions of the provider.

-> **Note** Use the same user who did the exporting to import the exported templates.  Otherwise it could cause the changes in the jobs ownership.

Generates `*.tf` files for Databricks resources as well as `import.sh` to run import state. Available as part of provider binary. The only possible way to authenticate is through [environment variables](../index.md#Environment-variables). It's best used, when you need to quickly export Terraform configuration for an existing Databricks workspace. After generating configuration, we strongly recommend to manually review all created files.

## Example Usage

After downloading the [latest released binary](https://github.com/databrickslabs/terraform-provider-databricks/releases), unpack it and place it in the same folder. In fact, you may have already downloaded this binary - check `.terraform` folder of any state directory, where you've used `databricks` provider. It could also be in your plugin cache `~/.terraform.d/plugins/registry.terraform.io/databrickslabs/databricks/*/*/terraform-provider-databricks`. Here's the tool in action:

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

!> **Warning** This tooling was only extensively tested with administrator priviliges. 

All arguments are optional and they tune what code is being generated.

* `-directory` - Path to directory, where `*.tf` and `import.sh` files would be written. By default it's set to current working directory.
* `-module` - Name of module in Terraform state, that would affect reference resolution and prefixes for generated commands in `import.sh`.
* `-last-active-days` - Items with older than `-last-active-days` won't be imported. By default the value is set to 3650 (10 years). Has effect on listing [databricks_cluster](../resources/cluster.md) and [databricks_job](../resources/job.md) resources.
* `-services` - Coma-separated list of services to import. By default all services are imported. 
* `-listing` - Coma-separated list of services to be listed and further passed on for importing. `-services` parameter controls which transitive dependencies will be processed. We recommend limiting with `-listing` more often, than with `-services`.
* `-match` - Match resource names during listing operation. This filter applies to all resources that are getting listed, so if you want to import all dependencies of just one cluster, specify `-match=autoscaling -listing=compute`. By default is empty, which matches everything.
* `-mounts` - List DBFS mount points, which is a extremely slow operation and would not trigger unless explicitly specified.
* `-generateProviderDeclaration` - flag that toggles generation of `databricks.tf` file with declaration of the Databricks Terraform provider that is necessary for Terraform versions since Terraform 0.13 (disabled by default).
* `-prefix` - optional prefix that will be added to the name of all exported resources - that's useful for exporting resources multiple workspaces for merging into single one.
* `-skip-interactive` - optionally run in a non-interactive mode.

## Services

Services are just logical groups of resources used for filtering and organization in files written in `-directory`. All resources are globally sorted by their resource name, which technically allows you to use generated files for compliance purposes. Nevertheless, managing entire Databricks workspace with Terraform is the prefered way. With the exception of notebooks and possibly libraries, which may have their own CI/CD processes.
* `groups` - [databricks_group](../data-sources/group.md) with [membership](../resources/group_member.md) and [data access](../resources/group_instance_profile.md).
* `users` - [databricks_user](../resources/user.md) are written to their own file, simply because of their amount. If you use SCIM provisioning, the only use-case for importing `users` service is to migrate workspaces.
* `compute` - **listing** [databricks_cluster](../resources/cluster.md). Includes [policies](../resources/cluster_policy.md), [permissions](../resources/permissions.md), [pools](../resources/instance_pool.md).
* `jobs` - **listing** [databricks_job](../resources/job.md). Usually there are more automated jobs, than interactive clusters, so they get their own file in this tool's output.
* `access` - [databricks_permissions](../resources/permissions.md) and [databricks_instance_profile](../resources/instance_profile.md).
* `secrets` - **listing** [databricks_secret_scope](../resources/secret_scope.md) along with [keys](../resources/secret.md) and [ACLs](../resources/secret_acl.md). 
* `storage` - any [databricks_dbfs_file](../resources/dbfs_file.md) will be downloaded locally and properly arranged into terraform state.
* `mounts` - works only in combination with `-mounts`.
* `notebooks` - [databricks_notebook](../resources/notebook.md)
* `workspace` - [databricks_workspace_conf](../resources/workspace_conf.md) and [databricks_global_init_script](../resources/global_init_script.md)

## Secrets

For security reasons, [databricks_secret](../resources/secret.md) cannot contain actual plaintext secrets. Importer will create variable in `vars.tf`, that would have the same name as secret. You are supposed to [fill in the value of the secret](https://blog.gruntwork.io/a-comprehensive-guide-to-managing-secrets-in-your-terraform-code-1d586955ace1#0e7d) after that.

## Support Matrix for the exporter

| Resource | Exporter? |
| --- | --- |
| [databricks_aws_assume_role_policy](docs/data-sources/aws_assume_role_policy.md) data | No |
| [databricks_aws_bucket_policy](docs/data-sources/aws_bucket_policy.md) data | No |
| [databricks_aws_crossaccount_policy](docs/data-sources/aws_crossaccount_policy.md) data | No |
| [databricks_catalog](docs/resources/catalog.md) | No |
| [databricks_catalogs](docs/data-sources/catalog.md) data | No |
| [databricks_cluster](docs/resources/cluster.md) | Yes |
| [databricks_clusters](docs/data-sources/clusters.md) data | Yes |
| [databricks_cluster_policy](docs/resources/cluster_policy.md) | Yes |
| [databricks_current_user](docs/data-sources/current_user.md) | Yes |
| [databricks_dbfs_file](docs/resources/dbfs_file.md) | Yes |
| [databricks_dbfs_file_paths](docs/data-sources/dbfs_file_paths.md) data | No |
| [databricks_external_location](docs/resources/external_location.md) | No |
| [databricks_global_init_script](docs/resources/global_init_script.md) | Yes |
| [databricks_grants](docs/resources/grants.md) | No |
| [databricks_group](docs/resources/group.md) | Yes |
| [databricks_group_instance_profile](docs/resources/group_instance_profile.md) | Yes |
| [databricks_group_member](docs/resources/group_member.md) | Yes |
| [databricks_instance_pool](docs/resources/instance_pool.md)       | Yes |
| [databricks_instance_profile](docs/resources/instance_profile.md) | Yes |
| [databricks_ip_access_list](docs/resources/ip_access_list.md) | Yes |
| [databricks_job](docs/resources/job.md) | Yes |
| [databricks_library](docs/resources/library.md) | No |
| [databricks_metastore](docs/resources/metastore.md) | No |
| [databricks_metastore_assignment](docs/resources/metastore_assignment.md) | No |
| [databricks_metastore_data_access](docs/resources/metastore_data_access.md) | No |
| [databricks_mlflow_model](docs/resources/mlflow_model.md) | No |
| [databricks_mlflow_experiment](docs/resources/mlflow_experiment.md) | No |
| [databricks_mws_credentials](docs/resources/mws_credentials.md) | No |
| [databricks_mws_customer_managed_keys](docs/resources/mws_customer_managed_keys.md) | No |
| [databricks_mws_log_delivery](docs/resources/mws_log_delivery.md) | No |
| [databricks_mws_networks](docs/resources/mws_networks.md) | No |
| [databricks_mws_storage_configurations](docs/resources/mws_storage_configurations.md) | No |
| [databricks_mws_workspaces](docs/resources/mws_workspaces.md) | No |
| [databricks_node_type](docs/data-sources/node_type.md) data | No |
| [databricks_notebook](docs/resources/notebook.md) | Yes |
| [databricks_notebook](docs/data-sources/notebook.md) data | Yes |
| [databricks_notebook_paths](docs/data-sources/notebook_paths.md) data | No |
| [databricks_obo_token](docs/resources/obo_token.md) | No |
| [databricks_permissions (job)](docs/resources/permissions.md) | Yes |
| [databricks_permissions (cluster)](docs/resources/permissions.md) | Yes |
| [databricks_permissions (instance_pool)](docs/resources/permissions.md) | Yes |
| [databricks_permissions (cluster_policy)](docs/resources/permissions.md) | Yes |
| [databricks_permissions (repo)](docs/resources/permissions.md) | Yes |
| [databricks_permissions (token)](docs/resources/permissions.md) | No |
| [databricks_permissions (password)](docs/resources/permissions.md) | No |
| [databricks_permissions (delta live tables pipeline)](docs/resources/permissions.md) | No |
| [databricks_permissions (notebook)](docs/resources/permissions.md) | No |
| [databricks_permissions (directory)](docs/resources/permissions.md) | No |
| [databricks_permissions (mlflow experiment)](docs/resources/permissions.md) | No |
| [databricks_permissions (mlflow registered model)](docs/resources/permissions.md) | No |
| [databricks_permissions (sql_endpoint) ](docs/resources/permissions.md) | No |
| [databricks_pipeline](docs/resources/pipeline.md) | No |
| [databricks_repo](docs/resources/repo.md) | Yes |
| [databricks_schema](docs/resources/schema.md) | No |
| [databricks_schemas](docs/data-sources/schema.md) data | No |
| [databricks_secret](docs/resources/secret.md) | Yes |
| [databricks_secret_acl](docs/resources/secret_acl.md) | Yes |
| [databricks_secret_scope](docs/resources/secret_scope.md) | Yes |
| [databricks_spark_version](docs/data-sources/spark_version.md) data | No |
| [databricks_sql_dashboard](docs/resources/sql_dashboard.md) | No |
| [databricks_sql_endpoint](docs/resources/sql_endpoint.md) | No |
| [databricks_sql_global_config](docs/resources/sql_global_config.md) | No |
| [databricks_sql_permissions](docs/resources/sql_permissions.md) | No |
| [databricks_sql_query](docs/resources/sql_query.md) | No |
| [databricks_sql_visualization](docs/resources/sql_visualization.md) | No |
| [databricks_sql_widget](docs/resources/sql_widget.md) | No |
| [databricks_storage_credential](docs/resources/storage_credential.md) | No |
| [databricks_table](docs/resources/table.md) | No |
| [databricks_tables](docs/data-sources/table.md) data | No |
| [databricks_token](docs/resources/token.md) | No |
| [databricks_user](docs/resources/user.md) | Yes |
| [databricks_user_instance_profile](docs/resources/user_instance_profile.md) | No (Depricated) |
| [databricks_workspace_conf](docs/resources/workspace_conf.md) | No |
