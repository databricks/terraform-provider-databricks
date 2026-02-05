---
subcategory: "Databricks SQL"
---
# databricks_warehouses_default_warehouse_overrides Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to list all default warehouse overrides in the workspace.

-> **Note** This data source requires workspace admin permissions.








## Example Usage
### List All Overrides
Returns a list of all default warehouse overrides in the workspace:

```hcl
data "databricks_default_warehouse_overrides" "all" {
}
```








## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of overrides to return. The service may return fewer than
  this value.
  If unspecified, at most 100 overrides will be returned.
  The maximum value is 1000; values above 1000 will be coerced to 1000


## Attributes
This data source exports a single attribute, `default_warehouse_overrides`. It is a list of resources, each with the following attributes:
* `default_warehouse_override_id` (string) - The ID component of the resource name (user ID)
* `name` (string) - The resource name of the default warehouse override.
  Format: default-warehouse-overrides/{default_warehouse_override_id}
* `type` (string) - The type of override behavior. Possible values are: `CUSTOM`, `LAST_SELECTED`
* `warehouse_id` (string) - The specific warehouse ID when type is CUSTOM.
  Not set for LAST_SELECTED type