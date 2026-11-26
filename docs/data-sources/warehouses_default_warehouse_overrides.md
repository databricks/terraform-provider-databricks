---
subcategory: "Databricks SQL"
---
# databricks_warehouses_default_warehouse_overrides Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/warehouses)

This data source can be used to list all default warehouse overrides in the workspace.

-> **Note** This data source requires workspace admin permissions.








## Example Usage
### List All Overrides
Returns a list of all default warehouse overrides in the workspace:

```hcl
data "databricks_warehouses_default_warehouse_overrides" "all" {
}
```








## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of overrides to return. The service may return fewer than
  this value.
  If unspecified, at most 100 overrides will be returned.
  The maximum value is 1000; values above 1000 will be coerced to 1000
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `default_warehouse_overrides`. It is a list of resources, each with the following attributes:
* `default_warehouse_override_id` (string) - The numeric ID of the user that the override applies to. It is the last component of `name`
* `name` (string) - The resource name of the default warehouse override.
  Format: default-warehouse-overrides/{default_warehouse_override_id}
  Responses always contain the numeric user ID, even when the request used `me`
* `type` (string) - The type of override behavior. Possible values are: `CUSTOM`, `LAST_SELECTED`
* `warehouse_id` (string) - The ID of the SQL warehouse to preselect when `type` is `CUSTOM`. It is required for `CUSTOM`
  and must be a valid warehouse ID. It is ignored for `LAST_SELECTED` and is not returned for that type