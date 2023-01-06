---
subcategory: "AWS"
---
# databricks_mws_credentials Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Lists all [databricks_mws_credentials](../resources/mws_credentials.md) in Databricks Account.

-> **Note** [`account_id`](../index.md#account_id) provider configuration property is required for this resource to work.

## Example Usage

Listing all workspaces in 

```hcl
provider "databricks" {
  // other configuration
  account_id = "<databricks account id>"
}

data "databricks_mws_credentials" "all" {}

output "all_mws_credentials" {
  value = data.databricks_mws_credentials.all.ids
}
```

## Attribute Reference

-> **Note** This resource has an evolving interface, which may change in future versions of the provider.

This data source exports the following attributes:

* `ids` - name-to-id map for all of the credentials in the account

## Related Resources

The following resources are used in the same context:

* [databricks_mws_workspaces](../resources/mws_workspaces.md) to manage Databricks E2 Workspaces.
* [databricks_metastore_assignment](../resources/metastore_assignment.md) to assign [databricks_metastore](docs/resources/metastore.md) to [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace)
