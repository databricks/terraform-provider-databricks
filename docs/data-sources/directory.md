---
subcategory: "Workspace"
---
# databricks_directory Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

This data source allows to get information about a directory in a Databricks Workspace.

## Example Usage

```hcl
data "databricks_directory" "prod" {
  path = "/Production"
}
```

## Argument Reference

* `path` - (Required) Path to a directory in the workspace

## Attribute Reference

This data source exports the following attributes:

* `object_id` - directory object ID
* `workspace_path` - path on Workspace File System (WSFS) in form of `/Workspace` + `path`
