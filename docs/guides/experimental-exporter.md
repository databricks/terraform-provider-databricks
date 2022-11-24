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

## Services

Services are just logical groups of resources used for filtering and organization in files written in `-directory`. All resources are globally sorted by their resource name, which technically allows you to use generated files for compliance purposes. Nevertheless, managing the entire Databricks workspace with Terraform is the prefered way. With the exception of notebooks and possibly libraries, which may have their own CI/CD processes.
* `groups` - [databricks_group](../data-sources/group.md) with [membership](../resources/group_member.md) and [data access](../resources/group_instance_profile.md).
* `users` - [databricks_user](../resources/user.md) are written to their own file, simply because of their amount. If you use SCIM provisioning, the only use-case for importing `users` service is to migrate workspaces.
* `compute` - **listing** [databricks_cluster](../resources/cluster.md). Includes [policies](../resources/cluster_policy.md), [permissions](../resources/permissions.md), [pools](../resources/instance_pool.md).
* `jobs` - **listing** [databricks_job](../resources/job.md). Usually there are more automated jobs than interactive clusters, so they get their own file in this tool's output.
* `access` - [databricks_permissions](../resources/permissions.md) and [databricks_instance_profile](../resources/instance_profile.md).
* `secrets` - **listing** [databricks_secret_scope](../resources/secret_scope.md) along with [keys](../resources/secret.md) and [ACLs](../resources/secret_acl.md). 
* `storage` - any [databricks_dbfs_file](../resources/dbfs_file.md) will be downloaded locally and properly arranged into terraform state.
* `mounts` - works only in combination with `-mounts`.
* `notebooks` - [databricks_notebook](../resources/notebook.md)
* `workspace` - [databricks_workspace_conf](../resources/workspace_conf.md) and [databricks_global_init_script](../resources/global_init_script.md)

## Secrets

For security reasons, [databricks_secret](../resources/secret.md) cannot contain actual plaintext secrets. Importer will create a variable in `vars.tf`, that would have the same name as secret. You are supposed to [fill in the value of the secret](https://blog.gruntwork.io/a-comprehensive-guide-to-managing-secrets-in-your-terraform-code-1d586955ace1#0e7d) after that.

## Support Matrix

Exporter aims to generate HCL code for the most of resources within the Databricks workspace:

| Resource | Generated code |
| --- | --- |
| [databricks_cluster](../resources/cluster.md) | Yes |
| [databricks_cluster_policy](../resources/cluster_policy.md) | Yes |
| [databricks_dbfs_file](../resources/dbfs_file.md) | Yes |
| [databricks_global_init_script](../resources/global_init_script.md) | Yes |
| [databricks_group](../resources/group.md) | Yes |
| [databricks_group_instance_profile](../resources/group_instance_profile.md) | Yes |
| [databricks_group_member](../resources/group_member.md) | Yes |
| [databricks_instance_pool](../resources/instance_pool.md) | Yes |
| [databricks_instance_profile](../resources/instance_profile.md) | Yes |
| [databricks_ip_access_list](../resources/ip_access_list.md) | Yes |
| [databricks_job](../resources/job.md) | Yes |
| [databricks_library](../resources/library.md) | No |
| [databricks_mlflow_model](../resources/mlflow_model.md) | No |
| [databricks_mlflow_experiment](../resources/mlflow_experiment.md) | No |
| [databricks_notebook](../resources/notebook.md) | Yes |
| [databricks_obo_token](../resources/obo_token.md) | Not Applicable |
| [databricks_permissions](../resources/permissions.md) | Yes |
| [databricks_pipeline](../resources/pipeline.md) | Yes |
| [databricks_repo](../resources/repo.md) | Yes |
| [databricks_secret](../resources/secret.md) | Yes |
| [databricks_secret_acl](../resources/secret_acl.md) | Yes |
| [databricks_secret_scope](../resources/secret_scope.md) | Yes |
| [databricks_sql_dashboard](../resources/sql_dashboard.md) | Yes |
| [databricks_sql_endpoint](../resources/sql_endpoint.md) | Yes |
| [databricks_sql_global_config](../resources/sql_global_config.md) | No |
| [databricks_sql_permissions](../resources/sql_permissions.md) | No |
| [databricks_sql_query](../resources/sql_query.md) | Yes |
| [databricks_sql_visualization](../resources/sql_visualization.md) | Yes |
| [databricks_sql_widget](../resources/sql_widget.md) | Yes |
| [databricks_token](../resources/token.md) | Not Applicable |
| [databricks_user](../resources/user.md) | Yes |
| [databricks_user_instance_profile](../resources/user_instance_profile.md) | No (Deprecated) |
| [databricks_workspace_conf](../resources/workspace_conf.md) | Yes (partial) |
