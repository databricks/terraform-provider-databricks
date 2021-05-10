---
subcategory: "Workspace"
---
# databricks_notebook Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

This data source allows to export a notebook from workspace

## Example Usage

```hcl
data "databricks_notebook" "features" {
    path = "/Production/Features"
    format = "SOURCE"
}
```

## Argument Reference

* `path` - (Required) Notebook path on the workspace
* `format` - (Required) Notebook format to export. Either `SOURCE`, `HTML`, `JUPYTER`, or `DBC`.

## Attribute Reference

This data source exports the following attributes:

* `content` - notebook content in selected format
* `language` - notebook language
* `object_id` - notebook object ID
* `object_type` - notebook object type