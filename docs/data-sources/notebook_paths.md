---
subcategory: "Workspace"
---
# databricks_notebook_paths Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

This data source allows to list notebooks in the Databricks Workspace.

## Example Usage

```hcl
data "databricks_notebook_paths" "prod" {
  path      = "/Production"
  recursive = true
}
```

## Argument Reference

* `path` - (Required) Path to workspace directory
* `recursive` - (Required) Either or recursively walk given path

## Attribute Reference

This data source exports the following attributes:

* `notebook_path_list` - list of objects with `path` and `language` attributes
