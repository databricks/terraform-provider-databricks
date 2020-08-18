+++
title = "dbfs_file_sync"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_dbfs_file_sync`

This resource will let you create a dbfs file sync which will synchronize files between systems depending on file size 
changes. 

{{% notice warning %}}
This dbfs file sync resource only determines differences via file size so if data changes but file size does not
it will provide a false negative in terms of determining difference.
{{% /notice %}}

## Example Usage

```hcl
data "databricks_dbfs_file" "my_dbfs_file_data" {
  path = "/install_wheels_sri.sh"
  limit_file_size = true
}

resource "databricks_dbfs_file_sync" "my-file1-sync" {
  src_path = data.databricks_dbfs_file.my_dbfs_file_data.path
  tgt_path = "/terraformdbfs/example/wheels.sh"
  file_size = data.databricks_dbfs_file.my-dbfs-file-data.file_size
  mkdirs = true

  host = "https://<domain>.com/"
  token = "dapiherehereherehere"
}
```

    
## Argument Reference

The following arguments are supported:

#### - `src_path`:
> **(Required)** The source path in dbfs where to fetch the file from. 
It should look like "dbfs:/path/to/my/file". 

#### - `tgt_path`:
> **(Required)** The target path in dbfs where to sync the file to. 
It should look like "dbfs:/path/to/my/file".

#### - `file_size`:
> **(Required)** This is the value that is used to determine file change.
Unfortunately at this point in time there are no file checksums provided by the dbfs api. To download the file everytime 
terraform provides a Read operation during state refresh is not efficient.

#### - `mkdirs`:
> **(Optional)** This is whether the resource should create the required,
parent directories to sync the file. The default value is true.

#### - `host`:
> **(Optional)** This is an optional api host for the databricks api if the source 
file resides on another workspace/dbfs system.

#### - `token`:
> **(Optional)** This is an optional api token for the databricks api if the source 
file resides on another workspace/dbfs system.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> Id for the file sync object.


## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
