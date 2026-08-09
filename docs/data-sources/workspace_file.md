---
subcategory: "Workspace"
---
# databricks_workspace_file Data Source

This data source allows to list workspace files in the Databricks Workspace.

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_workspace_file" "all" {
  path      = "/Production"
  recursive = true
}
```

## Argument Reference

* `path` - (Required) Path to workspace directory.
* `recursive` - (Optional) Whether to recursively list files under the given path. Defaults to `false`.

## Attribute Reference

This data source exports the following attributes:

* `workspace_files` - list of objects with the following attributes:
  * `id` - the file path (same as `path`).
  * `created_at` - the creation UTC timestamp of the file.
  * `modified_at` - the last modified UTC timestamp of the file.
  * `object_id` - unique identifier of the workspace object.
  * `path` - the absolute path of the file.
  * `resource_id` - a unique identifier for the object that is consistent across all Databricks APIs.
  * `url` - the URL of the file in the Databricks workspace UI.
  * `workspace_path` - the absolute path prefixed with `/Workspace`.
