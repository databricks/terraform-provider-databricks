# databricks_notebook Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource allows you to manage the import, export, and delete notebooks. The maximum allowed size of a
request to resource is 10MB.

-> **Note** Though the public workspace import api supports notebooks of type `DBC`, `JUPYTER` and `HTML` this resource does not support those types as determining changes is quite challenging as the format includes additional information such as timestamps execution counts thus a base64 diff causes issues.

## Example Usage

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

* `content` - (Required) The base64-encoded content. If the limit (10MB) is exceeded, an exception with error code MAX_NOTEBOOK_SIZE_EXCEEDED will be thrown. *This field is **deprecated** and will be renamed to `content_base64` in 0.3.*
* `path` -  (Required) The absolute path of the notebook or directory, beginning with "/", e.g. "/mynotebook".
* `language` -  (Required) The language. If format is set to SOURCE, this field is required; otherwise, it will be ignored. Possible choices are SCALA, PYTHON, SQL, R. *This field will become optional in the next release.*
* `overwrite` - (Required) The flag that specifies whether to overwrite an existing object. It is false by default. *This field is **deprecated** and will be removed in 0.3.*
* `mkdirs` - (Required) Create the given directory and necessary parent directories if they do not exist. If there exists an object (not a directory) at any prefix of the input path, this call returns an error RESOURCE_ALREADY_EXISTS. If this operation fails it may have succeeded in creating some of the necessary parent directories. *This field is **deprecated** and will be removed in 0.3.*
* `format` -  (Required) This specifies the format of the file to be imported. This resource currently only supports SOURCE. The value is case sensitive. SOURCE is suitable for .scala, .py, .r, .sql extension based files, HTML for .html files, JUPYTER for .ipynb files. Though the API supports DBC, HTML, and JUPYTER currently we do not support them as effectively identifying DIFF is currently not feasible. *This field is **deprecated** and will be removed in 0.3.*

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  The id for the notebook object, which is equal to notebook path
* `object_id` -  Unique identifier for a NOTEBOOK or DIRECTORY.
* `object_type` -  The type of the object. It could be NOTEBOOK, DIRECTORY or LIBRARY.

## Access Control

* [databricks_permissions](permissions.md#Notebook-usage) can control which groups or individual users can access notebooks or folders.

## Import

The resource notebook can be imported using notebook path

```bash
$ terraform import databricks_notebook.this /Path/to/notebook
```


