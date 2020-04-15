# Resource: db_dbfs_file_sync

This resource will let you create a dbfs file sync which will synchronize files between systems depending on file size 
changes. 

.. Warning This dbfs file sync resource only determines differences via file size so if data changes but file size does not
it will provide a false negative in terms of determining difference.

## Example Usage

.. code-block:: tf
    
    
    data "db_dbfs_file" "my_dbfs_file_data" {
      path = "/install_wheels_sri.sh"
      limit_file_size = true
    }
    
    resource "db_dbfs_file_sync" "my-file1-sync" {
      src_path = data.db_dbfs_file.my_dbfs_file_data.path
      tgt_path = "/terraformdbfs/example/wheels.sh"
      file_size = data.db_dbfs_file.my-dbfs-file-data.file_size
      mkdirs = true
    
      host = "https://<domain>.com/"
      token = "dapiherehereherehere"
    }


    
## Argument Reference

The following arguments are supported:

.. _r_dbfs_file_sync_src_path:
* :ref:`src_path <r_dbfs_file_sync_src_path>` - **(Required)** The source path in dbfs where to fetch the file from. 
It should look like "dbfs:/path/to/my/file". 

.. _r_dbfs_file_sync_tgt_path:
* :ref:`tgt_path <r_dbfs_file_sync_tgt_path>` - **(Required)** The target path in dbfs where to sync the file to. 
It should look like "dbfs:/path/to/my/file".

.. _r_dbfs_file_sync_file_size:
* :ref:`file_size <r_dbfs_file_sync_file_size>` - **(Required)** This is the value that is used to determine file change.
Unfortunately at this point in time there are no file checksums provided by the dbfs api. To download the file everytime 
terraform provides a Read operation during state refresh is not efficient.

.. _r_dbfs_file_sync_mkdirs:
* :ref:`mkdirs <r_dbfs_file_sync_mkdirs>` - **(Optional)** This is whether the resource should create the required,
parent directories to sync the file. The default value is true.

.. _r_dbfs_file_sync_host:
* :ref:`host <r_dbfs_file_sync_host>` - **(Optional)** This is an optional api host for the databricks api if the source 
file resides on another workspace/dbfs system.

.. _r_dbfs_file_sync_token:
* :ref:`token <r_dbfs_file_sync_token>` - **(Optional)** This is an optional api token for the databricks api if the source 
file resides on another workspace/dbfs system.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_dbfs_file_sync_id:
* :ref:`id <r_dbfs_file_sync_id>` - Id for the file sync object.


## Import

.. Note:: Importing this resource is not currently supported.