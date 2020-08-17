# databricks_dbfs_file Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `path` - (Required) (String) 

* `overwrite` - (Optional) (Bool) 

* `mkdirs` - (Optional) (Bool) 

* `validate_remote_file` - (Optional) (Bool) 

* `content` - (Optional) (String) 

* `source` - (Optional) (String) 

* `content_b64_md5` - (Required) (String) 




## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the dbfs file.

* `file_size` - (Integer) 


## Import

The resource dbfs file can be imported using the `object`, e.g.

```bash
$ terraform import databricks_dbfs_file.object <dbfs file id>
```