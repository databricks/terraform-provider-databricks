---
subcategory: "Storage"
---
# databricks_dbfs_file Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

This data source allows to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).

## Example Usage

```hcl
data "databricks_dbfs_file" "report" {
  path            = "dbfs:/reports/some.csv"
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

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/passthrough-cluster-per-user.md) guide
* [databricks_dbfs_file_paths](dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file](../resources/dbfs_file.md) to manage relatively small files on [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_mount](../resources/mount.md) to [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`.
