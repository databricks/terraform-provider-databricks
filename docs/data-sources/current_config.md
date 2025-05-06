---
subcategory: "Deployment"
---
# databricks_current_config Data Source

Retrieves information about the currently configured provider to make a decision, for example, add a dynamic block based on the specific cloud.

-> This data source can be used with an account or workspace-level provider.

## Example Usage

Create cloud-specific [databricks_storage_credential](../resources/storage_credential.md):

```hcl
data "databricks_current_config" "this" {}

resource "databricks_storage_credential" "external" {
  name = "storage_cred"
  dynamic "aws_iam_role" {
    # map for a conditional block
    for_each = data.databricks_current_config.this.cloud_type == "aws" ? {} : { aws = true }
    content {
      role_arn = var.cloud_credential_id
    }
  }
  dynamic "azure_managed_identity" {
    # map for a conditional block
    for_each = data.databricks_current_config.this.cloud_type == "azure" ? {} : { azure = true }
    content {
      access_connector_id = var.cloud_credential_id
    }
  }
  dynamic "databricks_gcp_service_account" {
    # map for a conditional block
    for_each = data.databricks_current_config.this.cloud_type == "gcp" ? {} : { gcp = true }
    content {}
  }
  comment = "Managed by TF"
}
```

## Exported attributes

Data source exposes the following attributes:

* `is_account` - Whether the provider is configured at account-level
* `account_id` - Account Id if provider is configured at account-level
* `host` - Host of the Databricks workspace or account console
* `cloud_type` - Cloud type specified in the provider
* `auth_type` - Auth type used by the provider

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide
* [databricks_directory](../resources/directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_notebook](../resources/notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_git_folder](../resources/git_folder.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
