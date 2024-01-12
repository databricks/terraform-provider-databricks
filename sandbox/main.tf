# An example terraform template that you can use to test the Terraform provider.
#
# Usage:
#
#  export TF_CLI_CONFIG_FILE=local-dev.tfrc
#  export TF_LOG=debug
#  teraform plan
#  terraform apply

terraform {
    required_providers {
        databricks = {
            source = "databricks/databricks"
            # The version is not necessary to provide
        }
    }
}

provider "databricks" {
    # Configure your provider here. For example:
    # host = "https://accounts.cloud.databricks.com"
    # client_id = "00000000-0000-0000-0000-000000000000"
    # client_secret = "dose1234567890abcdef1234567890abcdef"
}

# Add your resources here. For example:
#
#  resource "databricks_notebook" "example" {
#      content_base64 = base64encode(<<-EOT
#          # created from ${abspath(path.module)}
#          display(spark.range(3))
#          EOT
#      )
#      language = "PYTHON"
#      path = "/Shared/HelloWorld"
#  }