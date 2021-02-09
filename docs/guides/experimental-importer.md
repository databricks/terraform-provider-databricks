---
page_title: "Experimental resource importer"
subcategory: "Guides"
---
# Experimental resource importer

Generates `*.tf` files for Databricks resources as well as `import.sh` to run import state. Available as part of provider binary. The only possible way to authenticate is through [environment variables](../index.md#Environment-variables).

## Example Usage

```bash
export DATABRICKS_HOST=...
export DATABRICKS_TOKEN=...
./terraform-provider-databricks importer \
    -services=groups,secrets,access,compute,users,jobs,storage \
    -listing=jobs,compute \
    -last-active-days=90 \
    -module=data_platform \
    -debug
sh import.sh
```

## Argument Reference

All arguments are optional and they tune what code is being generated.

* `-directory` - Path to directory, where `*.tf` and `import.sh` files would be written. By default it's set to current working directory.
* `-module` - Name of module in Terraform state, that would affect reference resolution and prefixes for generated commands in `import.sh`.
* `-last-active-days` - Items with older than `-last-active-days` won't be imported. By default the value is set to 3650 (10 years). Has effect on listing [databricks_cluster](../resources/cluster.md) and [databricks_job](../resources/job.md) resources.
* `-services` - Coma-separated list of services to import. By default all services are imported. 
* `listing` - Coma-separated list of services to be listed and further passed on for importing. `-services` parameter controls which transitive dependencies will be processed. We recommend limiting with `-listing` more often, than with `-services`.
* `-match` - Match resource names during listing operation. This filter applies to all resources that are getting listed, so if you want to import all dependencies of just one cluster, specify `-match=autoscaling -listing=compute`. By default is empty, which matches everything.
* `-mounts` - List DBFS mount points, which is a relatively slow operation and would not trigger unless explicitly specified.

## Services

~> `secrets` service works the same way as any other services in this tool and if no corresponding secret reference is found in state resources, `string_value` will contain a plain-text key. It will try to see if there's a matching secret anywhere in the terraform state, but if nothing is found - it'll dump it clear. So don't check the generated code into version control without verification.

Services are just logical groups of resources used for filtering and organization in files written in `-directory`. All resources are globally sorted by their resource name, which technically allows you to use generated files for compliance purposes. Nevertheless, managing entire Databricks workspace with Terraform is the prefered way. With the exception of notebooks and possibly libraries, which may have their own CI/CD processes.
* `groups` - [databricks_group](../data-sources/group.md) with [membership](../resources/group_member.md) and [data access](../resources/group_instance_profile.md).
* `users` - [databricks_user](../resources/user.md) are written to their own file, simply because of their amount. If you use SCIM provisioning, the only use-case for importing `users` service is to migrate workspaces.
* `compute` - **listing** [databricks_cluster](../resources/cluster.md). Includes [policies](../resources/cluster_policy.md), [permissions](../resources/permissions.md), [pools](../resources/instance_pool.md).
* `jobs` - **listing** [databricks_job](../resources/job.md). Usually there are more automated jobs, than interactive clusters, so they get their own file in this tool's output.
* `access` - [databricks_permissions](../resources/permissions.md) and [databricks_instance_profile](../resources/instance_profile.md).
* `secrets` - **listing** [databricks_secret_scope](../resources/secret_scope.md) along with [keys](../resources/secret.md) and [ACLs](../resources/secret_acl.md). 
* `storage` - any [databricks_dbfs_file](../resources/dbfs_file.md) will be downloaded locally and propertly arranged into terraform state.
* `mounts` - works only in combination with `-mounts` for [databricks_s3_mount].

## Secrets

For security reasons, [databricks_secret](../resources/secret.md) cannot contain actual plaintext secrets. Importer will create variable in `vars.tf`, that would have the same name as secret. You are supposed to [fill in the value of the secret](https://blog.gruntwork.io/a-comprehensive-guide-to-managing-secrets-in-your-terraform-code-1d586955ace1#0e7d) after that.