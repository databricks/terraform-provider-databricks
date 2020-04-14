# Resource: db_notebook

This resource allows you to manage the import, export, and delete notebooks. The maximum allowed size of a 
request to resource is 10MB. 


## Example Usage

.. code-block:: tf

    resource "db_notebook" "my_databricks_notebook" {
      content = filebase64("${path.module}/demo-terraform.dbc")
      path = "/workspace/terraform-test-folder/"
      overwrite = false
      mkdirs = true
      format = "DBC"
    }
    
## Argument Reference

The following arguments are supported:

.. _r_notebook_content:
* :ref:`content <r_notebook_content>` - **(Required)** The base64-encoded content. If the limit (10MB) is exceeded, 
exception with error code MAX_NOTEBOOK_SIZE_EXCEEDED will be thrown.

.. _r_notebook_path:
* :ref:`path <r_notebook_path>` - **(Required)** 	The absolute path of the notebook or directory. 
Exporting a directory is supported only for DBC. This field is **required**.

.. _r_notebook_language:
* :ref:`language <r_notebook_language>` - **(Required)** The language. If format is set to SOURCE, 
this field is required; otherwise, it will be ignored. Possible choices are SCALA, PYTHON, SQL, R.

.. _r_notebook_overwrite:
* :ref:`overwrite <r_notebook_overwrite>` - **(Required)** The flag that specifies whether to overwrite existing object. 
It is false by default. For DBC format, overwrite is not supported since it may contain a directory.

.. _r_notebook_mkdirs:
* :ref:`mkdirs <r_notebook_mkdirs>` - **(Required)** Create the given directory and necessary parent directories 
if they do not exists. If there exists an object (not a directory) at any prefix of the input path, this call 
returns an error RESOURCE_ALREADY_EXISTS. If this operation fails it may have succeeded in creating some of the necessary parent directories.

.. _r_notebook_format:
* :ref:`format <r_notebook_format>` - **(Required)** This specifies the format of the file to be imported. 
By default, this is SOURCE. However it may be one of: SOURCE, HTML, JUPYTER, DBC. The value is case sensitive.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

.. _r_notebook_object_id:
* :ref:`object_id <r_notebook_object_id>` - Unique identifier for a NOTEBOOK or DIRECTORY.

.. _r_notebook_object_type:
* :ref:`object_type <r_notebook_object_type>` - 	The type of the object. It could be NOTEBOOK, DIRECTORY or LIBRARY.

## Import

.. Note:: Importing this resource is not currently supported.