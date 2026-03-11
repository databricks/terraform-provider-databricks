---
subcategory: "Deployment"
---
# databricks_mws_workspaces Data Source

Lists all [databricks_mws_workspaces](../resources/mws_workspaces.md) in Databricks Account.

-> This data source can only be used with an account-level provider!

## Example Usage

Listing all workspaces in

```hcl
provider "databricks" {
  // other configuration
  account_id = "<databricks account id>"
}

data "databricks_mws_workspaces" "all" {}

output "all_mws_workspaces" {
  value = data.databricks_mws_workspaces.all.ids
}

output "all_mws_workspaces_deployment_names" {
  value = data.databricks_mws_workspaces.all.deployment_names
}
```

## Argument Reference

* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

-> This resource has an evolving interface, which may change in future versions of the provider.

This data source exports the following attributes:

* `ids` - name-to-id map for all of the workspaces in the account
* `deployment_names` - name-to-deployment-names map for all of the workspaces in the account

## Related Resources

The following resources are used in the same context:

* [databricks_mws_workspaces](../resources/mws_workspaces.md) to manage Databricks Workspaces on AWS and GCP.
* [databricks_metastore_assignment](../resources/metastore_assignment.md) to assign [databricks_metastore](../resources/metastore.md) to [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace)
