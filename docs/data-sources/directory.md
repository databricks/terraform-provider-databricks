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
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exports the following attributes:

* `object_id` - directory object ID
* `workspace_path` - path on Workspace File System (WSFS) in form of `/Workspace` + `path`
