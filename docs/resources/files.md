---
subcategory: "Storage"
---
# databricks_files Resource

This is a resource that lets you upload and download files upto 2GB in octet-stream.

## Example Usage

In order to manage a file on Databricks File System with Terraform, you must specify the `source` attribute containing the full path to the file on the local filesystem.

```hcl
resource "databricks_files" "this" {
  source = "${path.module}/main.tf"
  path   = "/tmp/main.tf"
}
```

## Argument Reference

The following arguments are supported:

* `source` - The full absolute path to the file. Conflicts with `content_base64`.
* `path` - (Required) The path of the file in which you wish to save.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Same as `path`.
* `file_size` - The file size of the file that is being tracked by this resource in bytes.
* `modification_time` - The last time stamp when the file was modified


## Import

The resource dbfs file can be imported using the path of the file:

```bash
$ terraform import databricks_files.this <path>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_workspace_file](./workspace_file.md)
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_dbfs_file](../data-sources/dbfs_file.md) data to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_global_init_script](global_init_script.md) to manage [global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
* [databricks_mount](mount.md) to [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`.
* [databricks_workspace_conf](workspace_conf.md) to manage workspace configuration for expert usage.
