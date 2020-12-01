# databricks_dbfs_file Resource

This is a resource that lets you manage relatively small files on Databricks File System (DBFS). The best use cases are libraries or some configuration files.

## Example Usage

```hcl
resource "databricks_dbfs_file" "readme" {
  source = "${path.module}/README.md"
  path = "/sri/terraformdbfs/example/README.md"
}
```

## Argument Reference

The following arguments are supported:

* `source` - The full absolute path to the file. 
* `path` - (Required) The path of the file in which you wish to save.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Same as `path`.
* `file_size` - The file size of the file that is being tracked by this resource in bytes.


## Import

The resource dbfs file can be imported using the path of the file

```bash
$ terraform import databricks_dbfs_file.this <dbfs-file-id>
```
