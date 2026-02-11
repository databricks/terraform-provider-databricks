---
subcategory: "Databricks SQL"
---
# databricks_warehouses_default_warehouse_override Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

The Default Warehouse Override resource allows you to configure a user's default warehouse selection behavior in Databricks SQL. This resource enables customization of how a user's default warehouse is selected for SQL operations.

Users can configure their default warehouse to either:
- Remember their last-selected warehouse (`LAST_SELECTED` type)
- Use a specific warehouse (`CUSTOM` type with a warehouse ID)

-> **Note** The `default_warehouse_override_id` field represents the **user ID** of the user whose default warehouse behavior is being configured.

### Permissions
- Users can manage their own default warehouse override
- Workspace administrators can manage overrides for any user

### Behavior
If no override exists for a user, the workspace default warehouse will be used.



## Example Usage
### Basic Example with Last Selected Type
This example creates a default warehouse override that remembers the user's last-selected warehouse.
The `default_warehouse_override_id` represents the user ID of the target user:

```hcl
resource "databricks_default_warehouse_override" "last_selected" {
  # The user ID to configure the default warehouse override for
  default_warehouse_override_id = data.databricks_user.example.id
  type                          = "LAST_SELECTED"
}
```

### Custom Warehouse Example
This example creates a default warehouse override that always uses a specific warehouse:

```hcl
resource "databricks_default_warehouse_override" "custom" {
  # The user ID to configure the default warehouse override for
  default_warehouse_override_id = data.databricks_user.example.id
  type                          = "CUSTOM"
  warehouse_id                  = databricks_sql_endpoint.example.id
}
```


## Arguments
The following arguments are supported:
* `default_warehouse_override_id` (string, required) - The ID component of the resource name (user ID)
* `type` (string, required) - The type of override behavior. Possible values are: `CUSTOM`, `LAST_SELECTED`
* `warehouse_id` (string, optional) - The specific warehouse ID when type is CUSTOM.
  Not set for LAST_SELECTED type
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `name` (string) - The resource name of the default warehouse override.
  Format: default-warehouse-overrides/{default_warehouse_override_id}

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_warehouses_default_warehouse_override.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_warehouses_default_warehouse_override.this "name"
```