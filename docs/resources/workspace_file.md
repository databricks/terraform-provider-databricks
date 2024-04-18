---
subcategory: "Workspace"
---
# databricks_workspace_file Resource

This resource allows you to manage [Databricks Workspace Files](https://docs.databricks.com/files/workspace.html).

## Example Usage

You can declare Terraform-managed workspace file by specifying `source` attribute of corresponding local file.

```hcl
data "databricks_current_user" "me" {
}

resource "databricks_workspace_file" "module" {
  source = "${path.module}/module.py"
  path   = "${data.databricks_current_user.me.home}/AA/BB/CC"
}
```

You can also create a managed workspace file with inline sources through `content_base64`  attribute.

```hcl
resource "databricks_workspace_file" "init_script" {
  content_base64 = base64encode(<<-EOT
    #!/bin/bash
    echo "Hello World"
    EOT
  )
  path = "/Shared/init-script.sh"
}
```

## Argument Reference

-> **Note** Files in Databricks workspace would only be changed, if Terraform stage did change. This means that any manual changes to managed workspace files won't be overwritten by Terraform, if there's no local change to file sources. Workspace files are identified by their path, so changing file's name manually on the workspace and then applying Terraform state would result in creation of workspace file from Terraform state.

The size of a workspace file source code must not exceed a few megabytes. The following arguments are supported:

* `path` -  (Required) The absolute path of the workspace file, beginning with "/", e.g. "/Demo".
* `source` - Path to file on local filesystem. Conflicts with `content_base64`.
* `content_base64` - The base64-encoded file content. Conflicts with `source`. Use of `content_base64` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances, like creating a workspace file with configuration properties for a data pipeline.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` -  Path of workspace file
* `url` - Routable URL of the workspace file
* `object_id` -  Unique identifier for a workspace file
* `workspace_path` - path on Workspace File System (WSFS) in form of `/Workspace` + `path`

## Access Control

* [databricks_permissions](permissions.md#workspace-file-usage) can control which groups or individual users can access workspace file.

## Import

The workspace file resource can be imported using workspace file path

```bash
terraform import databricks_workspace_file.this /path/to/file
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_job](job.md) to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).
* [databricks_pipeline](pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
* [databricks_secret](secret.md) to manage [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_secret_acl](secret_acl.md) to manage access to [secrets](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_secret_scope](secret_scope.md) to create [secret scopes](https://docs.databricks.com/security/secrets/index.html#secrets-user-guide) in Databricks workspace.
* [databricks_user](user.md) to [manage users](https://docs.databricks.com/administration-guide/users-groups/users.html), that could be added to [databricks_group](group.md) within the workspace.
* [databricks_user](../data-sources/user.md) data to retrieve information about [databricks_user](user.md).
