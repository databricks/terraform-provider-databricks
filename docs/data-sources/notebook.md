---
subcategory: "Workspace"
---
# databricks_notebook Data Source

This data source allows to export a notebook from workspace

## Example Usage

```hcl
data "databricks_notebook" "features" {
    path = "/Production/Features"
    format = "SOURCE"
}
```

## Argument Reference

* `path` - (Required) Notebook path on the workspace
* `format` - (Required) Notebook format to export. Either `SOURCE`, `HTML`, `JUPYTER`, or `DBC`.

## Attribute Reference

This data source exports the following attributes:

* `content` - notebook content in selected format
* `language` - notebook language
* `object_id` - notebook object ID
* `object_type` - notebook object type