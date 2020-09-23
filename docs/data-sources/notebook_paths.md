# databricks_notebook_paths Data Source

This data source allows to list notebooks in the workspace

## Example Usage

```hcl
data "databricks_notebook_paths" "prod" {
    path = "/Production"
    recursive = true
}
```

## Argument Reference

* `path` - (Required) Path to workspace directory
* `recursive` - (Required) Either or resursively walk given path

## Attribute Reference

This data source exports the following attributes:

* `notebook_path_list` - list of objects with `path` and `language` attributes