---
subcategory: "Workspace"
---

# databricks_directory Resource

This resource allows you to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).

-> This resource can only be used with a workspace-level provider!

## Example Usage

You can declare a Terraform-managed directory by specifying the `path` attribute of the corresponding directory.

```hcl
resource "databricks_directory" "my_custom_directory" {
  path = "/my_custom_directory"
}
```

## Argument Reference

The following arguments are supported:

- `path` - (Required) The absolute path of the directory, beginning with "/", e.g. "/Demo".
- `delete_recursive` - Whether or not to trigger a recursive delete of this directory and its resources when deleting this on Terraform. Defaults to `false`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - Path of directory on workspace
- `object_id` - Unique identifier for a DIRECTORY
- `workspace_path` - path on Workspace File System (WSFS) in form of `/Workspace` + `path`

## Access Control

- [databricks_permissions](permissions.md#Folder-usage) can control which groups or individual users can access folders.

## Import

The resource directory can be imported using directory path:

```hcl
import {
  to = databricks_directory.this
  id = "/path/to/directory"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_directory.this /path/to/directory
```

## Related Resources

The following resources are often used in the same context:

- [End to end workspace management](../guides/workspace-management.md) guide.
- [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
- [databricks_notebook](../data-sources/notebook.md) data to export a notebook from Databricks Workspace.
- [databricks_notebook_paths](../data-sources/notebook_paths.md) data to list notebooks in Databricks Workspace.
- [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
- [databricks_spark_version](../data-sources/spark_version.md) data to get [Databricks Runtime (DBR)](https://docs.databricks.com/runtime/dbr.html) version that could be used for `spark_version` parameter in [databricks_cluster](cluster.md) and other resources.
- [databricks_workspace_conf](workspace_conf.md) to manage workspace configuration for expert usage.
