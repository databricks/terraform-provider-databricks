+++
title = "notebook"
date = 2020-04-20T23:34:03-04:00
weight = 15
chapter = false
pre = ""
+++

## Resource: `databricks_notebook`

This resource allows you to manage the import, export, and delete notebooks. The maximum allowed size of a 
request to resource is 10MB. 


## Example Usage

```hcl
resource "databricks_notebook" "my_databricks_notebook" {
  content = filebase64("${path.module}/demo-terraform.dbc")
  path = "/workspace/terraform-test-folder/"
  overwrite = false
  mkdirs = true
  format = "DBC"
}
```

For deployment of an empty Python notebook, the following example might be useful:

```hcl
resource "databricks_notebook" "notebook" {
  content = base64encode("# Welcome to your Python notebook")
  path = "/mynotebook"
  overwrite = false
  mkdirs = true
  language = "PYTHON"
  format = "SOURCE"
}
```
    
## Argument Reference

The following arguments are supported:

#### - `content`:
> **(Required)** The base64-encoded content. If the limit (10MB) is exceeded, 
exception with error code MAX_NOTEBOOK_SIZE_EXCEEDED will be thrown.

#### - `path`:
> **(Required)** The absolute path of the notebook or directory, beginning with "/", e.g. "/mynotebook"
Exporting a directory is supported only for DBC. This field is **required**.

#### - `language`:
> **(Required)** The language. If format is set to SOURCE, 
this field is required; otherwise, it will be ignored. Possible choices are SCALA, PYTHON, SQL, R.

#### - `overwrite`:
> **(Required)** The flag that specifies whether to overwrite existing object. 
It is false by default. For DBC format, overwrite is not supported since it may contain a directory.

#### - `mkdirs`:
> **(Required)** Create the given directory and necessary parent directories 
if they do not exists. If there exists an object (not a directory) at any prefix of the input path, this call 
returns an error RESOURCE_ALREADY_EXISTS. If this operation fails it may have succeeded in creating some of the necessary parent directories.

#### - `format`:
> **(Required)** This specifies the format of the file to be imported. 
By default, this is SOURCE. However it may be one of: SOURCE, HTML, JUPYTER, DBC. The value is case sensitive. SOURCE is suitable for .scala, .py, .r, .sql extension based files, HTML for .html files, JUPYTER for .ipynb files, and DBC for .dbc files.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

#### - `id`:
> The id for the notebook object.

#### - `object_id`:
> Unique identifier for a NOTEBOOK or DIRECTORY.

#### - `object_type`:
> The type of the object. It could be NOTEBOOK, DIRECTORY or LIBRARY.

## Import

{{% notice note %}}
Importing this resource is not currently supported.
{{% /notice %}}
