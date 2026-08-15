---
subcategory: "Deployment"
---

# databricks_mws_private_access_settings Data Source

-> **Note** This data source can only be used with an account-level provider!

This data source allows you to read the details of an existing [databricks_mws_private_access_settings](../resources/mws_private_access_settings.md) resource by its ID.

## Example Usage

```hcl
provider "databricks" {
  alias      = "mws"
  host       = "https://accounts.cloud.databricks.com"
  account_id = var.databricks_account_id
}

data "databricks_mws_private_access_settings" "this" {
  provider                   = databricks.mws
  private_access_settings_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}

output "private_access_settings_name" {
  value = data.databricks_mws_private_access_settings.this.private_access_settings_name
}
```

## Argument Reference

* `private_access_settings_id` - (Required) The ID of the private access settings object.

## Attribute Reference

This data source exports the following attributes:

* `account_id` - The Databricks account ID that hosts the private access settings.
* `allowed_vpc_endpoint_ids` - An array of Databricks VPC endpoint IDs that are allowed to connect to the workspace.
* `private_access_level` - The private access level (`ACCOUNT` or `ENDPOINT`).
* `private_access_settings_name` - The human-readable name of the private access settings object.
* `public_access_enabled` - Whether the workspace can be accessed over public internet.
* `region` - The cloud region for workspaces attached to this private access settings object.

## Related Resources

The following resources are used in the same context:

* [databricks_mws_private_access_settings](../resources/mws_private_access_settings.md) to manage private access settings.
* [databricks_mws_workspaces](../resources/mws_workspaces.md) to manage Databricks workspaces.
