---
subcategory: "Workspace"
---
# databricks_notebook Resource

This resource allows you to manage Databricks notebooks. You can also work with [databricks_notebook](../data-sources/notebook.md) and [databricks_notebook_paths](../data-sources/notebook_paths.md) data sources.

## Example Usage

You can declare Terraform-managed notebook by specifying `source` attribute of corresponding local file. Only `.scala`, `.py`, `.sql` and `.r` extensions are supported, if you would like to omit `language` attribute.

```hcl
data "databricks_current_user" "me" {
}

resource "databricks_notebook" "ddl" {
  source = "${path.module}/DDLgen.py"
  path   = "${data.databricks_current_user.me.home}/AA/BB/CC"
}
```

You can also create managed notebook with inline sources through `content_base64` and `language` attributes.

```hcl
resource "databricks_notebook" "notebook" {
  content_base64 = base64encode(<<-EOT
    # created from ${abspath(path.module)}
    display(spark.range(10))
    EOT
  )
  path     = "/Shared/Demo"
  language = "PYTHON"
}
```

You can also manage [Databricks Archives](https://docs.databricks.com/notebooks/notebooks-manage.html#databricks-archive) to import the whole folders of notebooks statically. Whenever you update `.dbc` file, terraform-managed notebook folder is removed and replaced with contents of the new `.dbc` file. You are strongly advised to use `.dbc` format only with `source` attribute of the resource:

```hcl
resource "databricks_notebook" "lesson" {
  source = "${path.module}/IntroNotebooks.dbc"
  path   = "/Shared/Intro"
}
```

## Argument Reference

-> **Note** Notebook on Databricks workspace would only be changed, if Terraform stage did change. This means that any manual changes to managed notebook won't be overwritten by Terraform, if there's no local change to notebook sources. Notebooks are identified by their path, so changing notebook's name manually on the workspace and then applying Terraform state would result in creation of notebook from Terraform state.

The size of a notebook source code must not exceed few megabytes. The following arguments are supported:

* `path` -  (Required) The absolute path of the notebook or directory, beginning with "/", e.g. "/Demo". 
* `source` - Path to notebook in source code format on local filesystem. Conflicts with `content_base64`.
* `content_base64` - The base64-encoded notebook source code. Conflicts with `source`. Use of `content_base64` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances, like creating a notebook with configuration properties for a data pipeline.
* `language` -  (required with `content_base64`) One of `SCALA`, `PYTHON`, `SQL`, `R`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Path of notebook on workspace
* `url` - Routable URL of the notebook
* `object_id` -  Unique identifier for a NOTEBOOK

## Access Control

* [databricks_permissions](permissions.md#Notebook-usage) can control which groups or individual users can access notebooks or folders.

## Import

The resource notebook can be imported using notebook path

```bash
$ terraform import databricks_notebook.this /path/to/notebook
```
