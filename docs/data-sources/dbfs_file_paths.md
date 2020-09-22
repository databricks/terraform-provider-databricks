# databricks_dbfs_file_paths Data Source

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