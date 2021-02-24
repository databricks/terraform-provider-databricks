---
subcategory: "Storage"
---
# databricks_dbfs_file Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _Authentication is not configured for provider_ errors.

This data source allows to get file content from DBFS

## Example Usage

```hcl
data "databricks_dbfs_file" "report" {
    path = "dbfs:/reports/some.csv"
    limit_file_size = 10240
}
```
## Argument Reference

* `path` - (Required) Path on DBFS for the file to get content of
* `limit_file_size` - (Required) Do lot load content for files smaller than this in bytes

## Attribute Reference

This data source exports the following attributes:

* `content` - base64-encoded file contents
* `file_size` - size of the file in bytes