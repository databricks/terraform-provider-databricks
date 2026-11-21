---
subcategory: "Workspace"
---
# databricks_workspace_file_paths Data Source

This data source allows to list [workspace files](https://docs.databricks.com/aws/en/files/workspace) in the Databricks Workspace. Only objects of type `FILE` are returned; directories, notebooks, repositories and other workspace objects are skipped. Use [databricks_notebook_paths](notebook_paths.md) to list notebooks.

-> This data source can only be used with a workspace-level provider!

## Example Usage

List all workspace files under a folder, including its subfolders:

```hcl
data "databricks_workspace_file_paths" "configs" {
  path      = "/Shared/configs"
  recursive = true
}

output "config_files" {
  value = [for f in data.databricks_workspace_file_paths.configs.workspace_file_path_list : f.path]
}
```

## Argument Reference

* `path` - (Required) Path to the workspace directory to list.
* `recursive` - (Optional) Whether to walk the given path recursively. Defaults to `false`, in which case only files directly under `path` are returned.
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exports the following attributes:

* `workspace_file_path_list` - list of workspace files, sorted by `path`. Each element has the following attributes:
  * `path` - absolute path of the file in the workspace.
  * `object_id` - unique identifier of the workspace object.
  * `resource_id` - unique identifier of the object that is consistent across all Databricks APIs.
  * `size` - size of the file in bytes.
  * `created_at` - creation time of the file, as a UTC timestamp in milliseconds since the epoch.
  * `modified_at` - last modification time of the file, as a UTC timestamp in milliseconds since the epoch.

## Related Resources

The following resources are used in the same context:

* [databricks_workspace_file](../resources/workspace_file.md) to manage workspace files.
* [databricks_directory](../resources/directory.md) to manage directories in the Databricks Workspace.
* [databricks_notebook_paths](notebook_paths.md) data to list notebooks in the Databricks Workspace.
