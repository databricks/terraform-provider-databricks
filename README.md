# Databricks Terraform Provider

![Resources](docs/resources.png)

[AWS](docs/guides/aws-workspace.md) tutorial
| [Azure](docs/guides/azure-workspace.md) tutorial
| [End-to-end](docs/guides/workspace-management.md) tutorial
| Migration from [0.3.x to 0.4.x](docs/guides/migration-0.4.x.md)
| [Changelog](CHANGELOG.md)
| [Authentication](docs/index.md)
| [databricks_aws_assume_role_policy](docs/data-sources/aws_assume_role_policy.md) data
| [databricks_aws_bucket_policy](docs/data-sources/aws_bucket_policy.md) data
| [databricks_aws_crossaccount_policy](docs/data-sources/aws_crossaccount_policy.md) data
| [databricks_cluster](docs/resources/cluster.md)
| [databricks_clusters](docs/data-sources/clusters.md) data
| [databricks_cluster_policy](docs/resources/cluster_policy.md)
| [databricks_current_user](docs/data-sources/current_user.md)
| [databricks_dbfs_file](docs/resources/dbfs_file.md)
| [databricks_dbfs_file_paths](docs/data-sources/dbfs_file_paths.md) data
| [databricks_dbfs_file](docs/data-sources/dbfs_file.md) data
| [databricks_global_init_script](docs/resources/global_init_script.md)
| [databricks_group](docs/resources/group.md)
| [databricks_group](docs/data-sources/group.md) data
| [databricks_group_instance_profile](docs/resources/group_instance_profile.md)
| [databricks_group_member](docs/resources/group_member.md)
| [databricks_instance_pool](docs/resources/instance_pool.md)
| [databricks_instance_profile](docs/resources/instance_profile.md)
| [databricks_ip_access_list](docs/resources/ip_access_list.md)
| [databricks_job](docs/resources/job.md)
| [databricks_mlflow_model](docs/resources/mlflow_model.md)
| [databricks_mlflow_experiment](docs/resources/mlflow_experiment.md)
| [databricks_mws_credentials](docs/resources/mws_credentials.md)
| [databricks_mws_customer_managed_keys](docs/resources/mws_customer_managed_keys.md)
| [databricks_mws_log_delivery](docs/resources/mws_log_delivery.md)
| [databricks_mws_networks](docs/resources/mws_networks.md)
| [databricks_mws_storage_configurations](docs/resources/mws_storage_configurations.md)
| [databricks_mws_workspaces](docs/resources/mws_workspaces.md)
| [databricks_node_type](docs/data-sources/node_type.md) data
| [databricks_notebook](docs/resources/notebook.md)
| [databricks_notebook](docs/data-sources/notebook.md) data
| [databricks_notebook_paths](docs/data-sources/notebook_paths.md) data
| [databricks_obo_token](docs/resources/obo_token.md)
| [databricks_permissions](docs/resources/permissions.md)
| [databricks_pipeline](docs/resources/pipeline.md)
| [databricks_repo](docs/resources/repo.md)
| [databricks_secret](docs/resources/secret.md)
| [databricks_secret_acl](docs/resources/secret_acl.md)
| [databricks_secret_scope](docs/resources/secret_scope.md)
| [databricks_spark_version](docs/data-sources/spark_version.md) data
| [databricks_sql_dashboard](docs/resources/sql_dashboard.md)
| [databricks_sql_endpoint](docs/resources/sql_endpoint.md)
| [databricks_sql_global_config](docs/resources/sql_global_config.md)
| [databricks_sql_permissions](docs/resources/sql_permissions.md)
| [databricks_sql_query](docs/resources/sql_query.md)
| [databricks_sql_visualization](docs/resources/sql_visualization.md)
| [databricks_sql_widget](docs/resources/sql_widget.md)
| [databricks_token](docs/resources/token.md)
| [databricks_user](docs/resources/user.md)
| [databricks_user_instance_profile](docs/resources/user_instance_profile.md)
| [databricks_workspace_conf](docs/resources/workspace_conf.md)
| [Contributing and Development Guidelines](CONTRIBUTING.md)

[![build](https://github.com/databrickslabs/terraform-provider-databricks/workflows/build/badge.svg?branch=master)](https://github.com/databrickslabs/terraform-provider-databricks/actions?query=workflow%3Abuild+branch%3Amaster) [![codecov](https://codecov.io/gh/databrickslabs/terraform-provider-databricks/branch/master/graph/badge.svg)](https://codecov.io/gh/databrickslabs/terraform-provider-databricks) ![lines](https://img.shields.io/tokei/lines/github/databrickslabs/terraform-provider-databricks) [![downloads](https://img.shields.io/github/downloads/databrickslabs/terraform-provider-databricks/total.svg)](https://hanadigital.github.io/grev/?user=databrickslabs&repo=terraform-provider-databricks)

If you use Terraform 0.13 or newer, please refer to instructions specified at [registry page](https://registry.terraform.io/providers/databrickslabs/databricks/latest). If you use older versions of Terraform or want to build it from sources, please refer to [contributing guidelines](CONTRIBUTING.md) page.

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databrickslabs/databricks"
      version = "0.4.0"
    }
  }
}
```

Then create a small sample file, named `main.tf` with approximately following contents. Replace `<your PAT token>` with newly created [PAT Token](https://docs.databricks.com/dev-tools/api/latest/authentication.html). 

```terraform
provider "databricks" {
  host = "https://abc-defg-024.cloud.databricks.com/"
  token = "<your PAT token>"
}

data "databricks_current_user" "me" {}
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_notebook" "this" {
  path     = "${data.databricks_current_user.me.home}/Terraform"
  language = "PYTHON"
  content_base64 = base64encode(<<-EOT
    # created from ${abspath(path.module)}
    display(spark.range(10))
    EOT
  )
}

resource "databricks_job" "this" {
  name = "Terraform Demo (${data.databricks_current_user.me.alphanumeric})"

  new_cluster {
    num_workers   = 1
    spark_version = data.databricks_spark_version.latest.id
    node_type_id  = data.databricks_node_type.smallest.id
  }

  notebook_task {
    notebook_path = databricks_notebook.this.path
  }

  email_notifications {}
}

output "notebook_url" {
  value = databricks_notebook.this.url
}

output "job_url" {
  value = databricks_job.this.url
}
```

Then run `terraform init` then `terraform apply` to apply the hcl code to your Databricks workspace. 

## Project Support

**Important:** Projects in the `databrickslabs` GitHub account, including the Databricks Terraform Provider, are not formally supported by Databricks. They are maintained by Databricks Field teams and provided as-is. There is no service level agreement (SLA). Databricks makes no guarantees of any kind. If you discover an issue with the provider, please file a GitHub Issue on the repo, and it will be reviewed by project maintainers as time permits.
