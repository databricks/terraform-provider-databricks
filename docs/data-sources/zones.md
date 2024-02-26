---
subcategory: "Deployment"
---
# databricks_zones Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

This data source allows you to fetch all available AWS availability zones on your workspace on AWS.

## Example Usage

```hcl
data "databricks_zones" "zones" {}
```

## Argument Reference

There are no arguments to this data source and only attributes that are computed.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id for the zone object.
* `default_zone` - This is the default zone that gets assigned to your workspace. This is the zone used by default for clusters and instance pools.
* `zones` - This is a list of all the zones available for your subnets in your Databricks workspace.
