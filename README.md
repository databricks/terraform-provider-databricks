# Databricks Terraform Provider

[![Build Status](https://travis-ci.org/databrickslabs/terraform-provider-databricks.svg?branch=master)](https://travis-ci.org/databrickslabs/terraform-provider-databricks) [![codecov](https://codecov.io/gh/databrickslabs/terraform-provider-databricks/branch/master/graph/badge.svg)](https://codecov.io/gh/databrickslabs/terraform-provider-databricks)

End-to-end workspace creation on [AWS](scripts/awsmt-integration) or [Azure](scripts/azvnet-integration/databricks.tf)
| [Authentication](docs/index.md)
| [databricks_aws_s3_mount](docs/resources/aws_s3_mount.md)
| [databricks_aws_assume_role_policy](docs/data-sources/aws_assume_role_policy.md) data
| [databricks_aws_bucket_policy](docs/data-sources/aws_bucket_policy.md) data
| [databricks_aws_crossaccount_policy](docs/data-sources/aws_crossaccount_policy.md) data
| [databricks_azure_adls_gen1_mount](docs/resources/azure_adls_gen1_mount.md)
| [databricks_azure_adls_gen2_mount](docs/resources/azure_adls_gen2_mount.md)
| [databricks_azure_blob_mount](docs/resources/azure_blob_mount.md)
| [databricks_cluster](docs/resources/cluster.md)
| [databricks_cluster_policy](docs/resources/cluster_policy.md)
| [databricks_dbfs_file](docs/resources/dbfs_file.md)
| [databricks_dbfs_file_paths](docs/data-sources/dbfs_file_paths.md) data
| [databricks_dbfs_file](docs/data-sources/dbfs_file.md) data
| [databricks_group](docs/resources/group.md)
| [databricks_group](docs/data-sources/group.md) data
| [databricks_group_instance_profile](docs/resources/group_instance_profile.md)
| [databricks_group_member](docs/resources/group_member.md)
| [databricks_instance_pool](docs/resources/instance_pool.md)
| [databricks_instance_profile](docs/resources/instance_profile.md)
| [databricks_job](docs/resources/job.md)
| [databricks_mws_credentials](docs/resources/mws_credentials.md)
| [databricks_mws_customer_managed_keys](docs/resources/mws_customer_managed_keys.md)
| [databricks_mws_networks](docs/resources/mws_networks.md)
| [databricks_mws_storage_configurations](docs/resources/mws_storage_configurations.md)
| [databricks_mws_workspaces](docs/resources/mws_workspaces.md)
| [databricks_notebook](docs/resources/notebook.md)
| [databricks_notebook](docs/data-sources/notebook.md) data
| [databricks_notebook_paths](docs/data-sources/notebook_paths.md) data
| [databricks_permissions](docs/resources/permissions.md)
| [databricks_secret](docs/resources/secret.md)
| [databricks_secret_acl](docs/resources/secret_acl.md)
| [databricks_secret_scope](docs/resources/secret_scope.md)
| [databricks_token](docs/resources/token.md)
| [databricks_user](docs/resources/user.md)
| [databricks_user_instance_profile](docs/resources/user_instance_profile.md)
| [Contributing and Development Guidelines](CONTRIBUTING.md)
| [Changelog](CHANGELOG.md)

If you use Terraform 0.13, please refer to instructions specified at [registry page](https://registry.terraform.io/providers/databrickslabs/databricks/latest):

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databrickslabs/databricks"
      version = ">= 0.2.7"
    }
  }
}
```

If you use Terraform 0.12, please execute the following curl command in your shell:

```bash
curl https://raw.githubusercontent.com/databrickslabs/databricks-terraform/master/godownloader-databricks-provider.sh | bash -s -- -b $HOME/.terraform.d/plugins
```

Then create a small sample file, named `main.tf` with approximately following contents. Replace `<your PAT token>` with newly created [PAT Token](https://docs.databricks.com/dev-tools/api/latest/authentication.html). It will create [a simple cluster](https://databrickslabs.github.io/terraform-provider-databricks/resources/cluster/).

```terraform
provider "databricks" {
  host = "https://abc-defg-024.cloud.databricks.com/"
  token = "<your PAT token>"
}

resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = "6.6.x-scala2.11"
  node_type_id            = "i3.xlarge"
  autotermination_minutes = 20

  autoscale {
    min_workers = 1
    max_workers = 50
  }
}
```

Then run `terraform init` then `terraform apply` to apply the hcl code to your Databricks workspace. 

## Project Support

**Important:** Projects in the `databrickslabs` GitHub account, including the Databricks Terraform Provider, are not formally supported by Databricks. They are maintained by Databricks Field teams and provided as-is. There is no service level agreement (SLA). Databricks makes no guarantees of any kind. If you discover an issue with the provider, please file a GitHub Issue on the repo, and it will be reviewed by project maintainers as time permits.
