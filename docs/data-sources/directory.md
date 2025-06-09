---
subcategory: "Workspace"
---
# databricks_directory Data Source

This data source allows to get information about a directory in a Databricks Workspace.

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_directory" "prod" {
  path = "/Production"
}
```

## Argument Reference

* `path` - (Required) Path to a directory in the workspace

## Attribute Reference

This data source exports the following attributes:

* `object_id` - directory object ID
* `workspace_path` - path on Workspace File System (WSFS) in form of `/Workspace` + `path`
