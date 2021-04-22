---
subcategory: "Storage"
---
# databricks_dbfs_file Resource

This is a resource that lets you manage relatively small files on Databricks File System (DBFS). The best use cases are libraries for [databricks_cluster](cluster.md) or [databricks_job](job.md). You can also use [databricks_dbfs_file](../data-sources/dbfs_file.md) and [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data sources.

## Example Usage

In order to manage file on Databricks File System with Terraform, you must specify `source` attribute containing full path to the file on local filesystem.

```hcl
resource "databricks_dbfs_file" "this" {
  source = "${path.module}/main.tf"
  path = "/tmp/main.tf"
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

## Argument Reference

-> **Note** DBFS files would only be changed, if Terraform stage did change. This means that any manual changes to managed file won't be overwritten by Terraform, if there's no local change. 

The following arguments are supported:

* `source` - The full absolute path to the file. Conflicts with `content_base64`.
* `content_base64` - Encoded file contents. Conflicts with `source`. Use of `content_base64` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances, like creating a data pipeline configuration file.
* `path` - (Required) The path of the file in which you wish to save.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Same as `path`.
* `file_size` - The file size of the file that is being tracked by this resource in bytes.
* `dbfs_path` - Path, but with `dbfs:` prefix


## Import

The resource dbfs file can be imported using the path of the file

```bash
$ terraform import databricks_dbfs_file.this <path>
```
