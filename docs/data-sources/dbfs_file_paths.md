---
subcategory: "Storage"
---
# databricks_dbfs_file_paths Data Source

This data source allows to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_dbfs_file_paths" "partitions" {
  path      = "dbfs:/user/hive/default.db/table"
  recursive = false
}
```

## Argument Reference

* `path` - (Required) Path on DBFS for the file to perform listing
* `recursive` - (Required) Either or not recursively list all files

## Attribute Reference

This data source exports the following attributes:

* `path_list` - returns list of objects with `path` and `file_size` attributes in each

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_dbfs_file](dbfs_file.md) data to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file_paths](dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file](../resources/dbfs_file.md) to manage relatively small files on [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_library](../resources/library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](../resources/cluster.md).
* [databricks_mount](../resources/mount.md) to [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`.
