---
subcategory: "Workspace"
---
# databricks_notebook_paths Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

This data source allows to list notebooks in the workspace

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
