# Resource: databricks_dbfs_file

This is a resource that lets you create, get and delete files in DBFS (Databricks File System).

## Example Usage

.. code-block:: tf

    resource "databricks_dbfs_file" "my_dbfs_file" {
      content = filebase64("README.md")
      path = "/sri/terraformdbfs/example/README.md"
      overwrite = true
      mkdirs = true
      validate_remote_file = true
    }
    
## Argument Reference

The following arguments are supported:

.. _r_dbfs_file_content:
* :ref:`content <r_dbfs_file_content>` - **(Required)** The content of the file as a base64 encoded string.

.. _r_dbfs_file_path:
* :ref:`path <r_dbfs_file_path>` - **(Optional)** The path of the file in which you wish to save.

.. _r_dbfs_file_overwrite:
* :ref:`overwrite <r_dbfs_file_overwrite>` - **(Optional)** This is used to determine whether it should delete the 
existing file when with the same name when it writes. The default is set to false.

.. _r_dbfs_file_mkdirs:
* :ref:`mkdirs <r_dbfs_file_mkdirs>` - **(Optional)** When the resource is created, this field is used to determine
if it needs to make the parent directories. The default value is set to true.

.. _r_dbfs_file_validate_remote_file:
* :ref:`validate_remote_file <r_dbfs_file_validate_remote_file>` - **(Optional)** This is used to compare the 
actual contents of the file to determine if the remote file is valid or not. If the base64 content is different 
it will attempt to do a delete, create.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_dbfs_file_id:
* :ref:`id <r_dbfs_file_id>` - The id for the dbfs file object.

.. _r_dbfs_file_file_size:
* :ref:`file_size <r_dbfs_file_file_size>` - The file size of the file that is being tracked by this resource in bytes.


## Import

.. Note:: Importing this resource is not currently supported.