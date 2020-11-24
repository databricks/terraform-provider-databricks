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
  language = "PYTHON"
}
```
    
## Argument Reference

The following arguments are supported:

* `content` - (Required) The base64-encoded content. If the limit (10MB) is exceeded, an exception with error code MAX_NOTEBOOK_SIZE_EXCEEDED will be thrown.
* `path` -  (Required) The absolute path of the notebook or directory, beginning with "/", e.g. "/mynotebook". This field is **required**.
* `language` -  (Required) The language. If format is set to SOURCE, this field is required; otherwise, it will be ignored. Possible choices are SCALA, PYTHON, SQL, R.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  The id for the notebook object (which is the same as `path`)
* `object_id` -  Unique identifier for a NOTEBOOK or DIRECTORY.
* `object_type` -  The type of the object. It could be NOTEBOOK, DIRECTORY or LIBRARY.

## Access Control

* [databricks_permissions](permissions.md#Notebook-usage) can control which groups or individual users can access notebooks or folders.

## Import

The resource notebook can be imported using notebook path

```bash
$ terraform import databricks_notebook.this /path/to/notebook
```


