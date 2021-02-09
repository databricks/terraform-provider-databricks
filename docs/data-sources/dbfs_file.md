---
subcategory: "Storage"
---
# databricks_dbfs_file Data Source

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