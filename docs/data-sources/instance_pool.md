---
subcategory: "Compute"
---

# databricks_instance_pool Data Source

Retrieves information about [databricks_instance_pool](../resources/instance_pool.md).

-> This data source can only be used with a workspace-level provider!

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
- `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  - `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

Data source exposes the following attributes:

- `id` - The id of the instance pool.
- `pool_info` - block describing instance pool and its state. Check documentation for [databricks_instance_pool](../resources/instance_pool.md) for a list of exposed attributes.
