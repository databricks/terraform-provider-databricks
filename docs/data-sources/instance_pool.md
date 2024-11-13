---
subcategory: "Compute"
---

# databricks_instance_pool Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about [databricks_instance_pool](../resources/instance_pool.md).

## Example Usage

Referring to an instance pool by name:

```hcl
data "databricks_instance_pool" "pool" {
  name = "All spot"
}

resource "databricks_cluster" "my_cluster" {
  instance_pool_id = data.databricks_instance_pool.pool.id
  # ...
}
```

## Argument Reference

Data source allows you to pick instance pool by the following attribute

- `name` - Name of the instance pool. The instance pool must exist before this resource can be planned.

## Attribute Reference

Data source exposes the following attributes:

- `id` - The id of the instance pool.
- `pool_info` - block describing instance pool and its state. Check documentation for [databricks_instance_pool](../resources/instance_pool.md) for a list of exposed attributes.
