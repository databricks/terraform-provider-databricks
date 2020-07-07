# Databricks Terraform Provider

[![Build Status](https://travis-ci.org/databrickslabs/terraform-provider-databricks.svg?branch=master)](https://travis-ci.org/databrickslabs/terraform-provider-databricks)

[Documentation](https://databrickslabs.github.io/terraform-provider-databricks/provider/) | [Contributing and Development Guidelines](CONTRIBUTING.md)

To quickly install the binary please execute the following curl command in your shell or [install provider from source](CONTRIBUTING.md#installing-from-source).

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

Then run `terraform init` then `terraform apply` to apply the hcl code to your Databricks workspace. Please refer to the [end-user documentation](https://databrickslabs.github.io/terraform-provider-databricks/provider/) for detailed use of the provider. Also refer to these [examples](examples/) for more scenarios. 

## Project Support
Please note that all projects in the /databrickslabs github account are provided for your exploration only, and are not formally supported by Databricks with Service Level Agreements (SLAs).  They are provided AS-IS and we do not make any guarantees of any kind.  Please do not submit a support ticket relating to any issues arising from the use of these projects.

Any issues discovered through the use of this project should be filed as GitHub Issues on the Repo.  They will be reviewed as time permits, but there are no formal SLAs for support.