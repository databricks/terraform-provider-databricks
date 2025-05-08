---
subcategory: "Storage"
---
# databricks_dbfs_file Resource

This is a resource that lets you manage relatively small files on [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html). The best use cases are libraries for [databricks_cluster](cluster.md) or [databricks_job](job.md). You can also use [databricks_dbfs_file](../data-sources/dbfs_file.md) and [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data sources.

-> This resource can only be used with a workspace-level provider!

## Example Usage

In order to manage a file on Databricks File System with Terraform, you must specify the `source` attribute containing the full path to the file on the local filesystem.

```hcl
resource "databricks_dbfs_file" "this" {
  source = "${path.module}/main.tf"
  path   = "/tmp/main.tf"
}
```

Alternatively, you can create DBFS files with custom content, using [filesystem functions](https://www.terraform.io/docs/language/functions/templatefile.html).

```hcl
resource "databricks_dbfs_file" "this" {
  content_base64 = base64encode(<<-EOT
    Hello, world!
    Module is ${abspath(path.module)}
    EOT
  )
  path = "/tmp/this.txt"
}
```

Install [databricks_library](library.md) on all [databricks_clusters](../data-sources/clusters.md):

```hcl
data "databricks_clusters" "all" {
}

resource "databricks_dbfs_file" "app" {
  source = "${path.module}/baz.whl"
  path   = "/FileStore/baz.whl"
}

resource "databricks_library" "app" {
  for_each   = data.databricks_clusters.all.ids
  cluster_id = each.key
  whl        = databricks_dbfs_file.app.dbfs_path
}
```

## Argument Reference

-> DBFS files would only be changed, if Terraform stage did change. This means that any manual changes to managed file won't be overwritten by Terraform, if there's no local change.

The following arguments are supported:

* `source` - The full absolute path to the file. Conflicts with `content_base64`.
* `content_base64` - Encoded file contents. Conflicts with `source`. Use of `content_base64` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances, like creating a data pipeline configuration file.
* `path` - (Required) The path of the file in which you wish to save.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Same as `path`.
* `file_size` - The file size of the file that is being tracked by this resource in bytes.
* `dbfs_path` - Path, but with `dbfs:` prefix.

## Import

The resource dbfs file can be imported using the path of the file:

```hcl
import {
  to = databricks_dbfs_file.this
  id = "<path>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_dbfs_file.this <path>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_dbfs_file](../data-sources/dbfs_file.md) data to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_global_init_script](global_init_script.md) to manage [global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
* [databricks_mount](mount.md) to [mount your cloud storage](https://docs.databricks.com/data/databricks-file-system.html#mount-object-storage-to-dbfs) on `dbfs:/mnt/name`.
* [databricks_workspace_conf](workspace_conf.md) to manage workspace configuration for expert usage.
