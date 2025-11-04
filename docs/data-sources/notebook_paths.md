---
subcategory: "Workspace"
---
# databricks_notebook_paths Data Source

This data source allows to list notebooks in the Databricks Workspace.

-> This data source can only be used with a workspace-level provider!

## Example Usage

```hcl
data "databricks_notebook_paths" "prod" {
  path      = "/Production"
  recursive = true
}
```

## Argument Reference

* `path` - (Required) Path to workspace directory
* `recursive` - (Required) Either or recursively walk given path
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exports the following attributes:

* `notebook_path_list` - list of objects with `path` and `language` attributes
