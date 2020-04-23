+++
title = "dbfs_file"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_dbfs_file`

This is a resource that lets you create, get and delete files in DBFS (Databricks File System).

## Example Usage

```hcl
resource "databricks_dbfs_file" "my_dbfs_file" {
  content = filebase64("README.md")
  path = "/sri/terraformdbfs/example/README.md"
  overwrite = true
  mkdirs = true
  validate_remote_file = true
}
```
    
## Argument Reference

The following arguments are supported:

#### - `content`:
> **(Required)** The content of the file as a base64 encoded string.

#### - `path`:
> **(Optional)** The path of the file in which you wish to save.

#### - `overwrite`:
> **(Optional)** This is used to determine whether it should delete the 
existing file when with the same name when it writes. The default is set to false.

#### - `mkdirs`:
> **(Optional)** When the resource is created, this field is used to determine
if it needs to make the parent directories. The default value is set to true.

#### - `validate_remote_file`:
> **(Optional)** This is used to compare the 
actual contents of the file to determine if the remote file is valid or not. If the base64 content is different 
it will attempt to do a delete, create.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the dbfs file object.

#### - `file_size`:
> The file size of the file that is being tracked by this resource in bytes.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
