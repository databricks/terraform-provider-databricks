---
subcategory: "Storage"
---
# databricks_dbfs_file_paths Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

This data source allows to get list of file names from DBFS

## Example Usage

```hcl
data "databricks_dbfs_file_paths" "partitions" {
    path = "dbfs:/user/hive/default.db/table"
    recursive = false
}
```
## Argument Reference

* `path` - (Required) Path on DBFS for the file to perform listing
* `recursive` - (Required) Either or not recursively list all files

## Attribute Reference

This data source exports the following attributes:

* `path_list` - returns list of objects with `path` and `file_size` attributes in each