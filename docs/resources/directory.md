---
subcategory: "Workspace"
---

# databricks_directory Resource

This resource allows you to manage Databricks directories.

## Example Usage

You can declare Terraform-managed directory by specifying `path` attribute of corresponding directory.

```hcl
resource "databricks_directory" "my_custom_directory" {
  path = "/my_custom_directory"
}
```

## Argument Reference

The following arguments are supported:

- `path` - (Required) The absolute path of the directory, beginning with "/", e.g. "/Demo".
- `delete_recursive` - Wether or not to trigger a recursive delete of this directory and its resources when deleting this on Terraform.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - Path of directory on workspace
- `object_id` - Unique identifier for a DIRECTORY

## Access Control

- [databricks_permissions](permissions.md#Folder-usage) can control which groups or individual users can access folders.

## Import

The resource notebook can be imported using notebook path

```bash
$ terraform import databricks_directory.this /path/to/notebook
```
