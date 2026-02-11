---
subcategory: "Databricks SQL"
---
# databricks_warehouses_default_warehouse_override Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

The Default Warehouse Override data source allows you to retrieve information about a user's default warehouse selection configuration in Databricks SQL.

You can use this data source to:
- Retrieve the current default warehouse override configuration for a user
- Check whether a user has a custom warehouse configured or uses last-selected behavior
- Get the warehouse ID if a custom warehouse is configured

-> **Note** The resource name format is `default-warehouse-overrides/{default_warehouse_override_id}`, where `default_warehouse_override_id` represents a user ID.


## Example Usage
### Get a User's Override
This example retrieves the default warehouse override for a specific user.
The name format is `default-warehouse-overrides/{default_warehouse_override_id}`:

```hcl
data "databricks_default_warehouse_override" "user" {
  # default_warehouse_override_id represents a user ID
  name = "default-warehouse-overrides/${data.databricks_user.example.id}"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the default warehouse override.
  Format: default-warehouse-overrides/{default_warehouse_override_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `default_warehouse_override_id` (string) - The ID component of the resource name (user ID)
* `name` (string) - The resource name of the default warehouse override.
  Format: default-warehouse-overrides/{default_warehouse_override_id}
* `type` (string) - The type of override behavior. Possible values are: `CUSTOM`, `LAST_SELECTED`
* `warehouse_id` (string) - The specific warehouse ID when type is CUSTOM.
  Not set for LAST_SELECTED type