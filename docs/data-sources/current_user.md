---
subcategory: "Security"
---
# databricks_current_user Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about [databricks_user](../resources/user.md) or [databricks_service_principal](../resources/service_principal.md), that is calling Databricks REST API. Might be useful in applying the same Terraform by different users in the shared workspace for testing purposes.

## Example Usage

Create personalized [databricks_job](../resources/job.md) and [databricks_notebook](../resources/notebook.md):

```hcl
data "databricks_current_user" "me" {}
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_notebook" "this" {
  path     = "${data.databricks_current_user.me.home}/Terraform"
  language = "PYTHON"
  content_base64 = base64encode(<<-EOT
    # created from ${abspath(path.module)}
    display(spark.range(10))
    EOT
  )
}

resource "databricks_job" "this" {
  name = "Terraform Demo (${data.databricks_current_user.me.alphanumeric})"

  task {
    task_key = "task1"

    new_cluster {
      num_workers   = 1
      spark_version = data.databricks_spark_version.latest.id
      node_type_id  = data.databricks_node_type.smallest.id
    }

    notebook_task {
      notebook_path = databricks_notebook.this.path
    }
  }
}

output "notebook_url" {
  value = databricks_notebook.this.url
}

output "job_url" {
  value = databricks_job.this.url
}
```

## Exported attributes

Data source exposes the following attributes:

* `id` -  The id of the calling user.
* `external_id` - ID of the user in an external identity provider.
* `user_name` - Name of the [user](../resources/user.md), e.g. `mr.foo@example.com`. If the currently logged-in identity is a [service principal](../resources/service_principal.md), returns the application ID, e.g. `11111111-2222-3333-4444-555666777888`
* `home` - Home folder of the [user](../resources/user.md), e.g. `/Users/mr.foo@example.com`.
* `repos` - Personal Repos location of the [user](../resources/user.md), e.g. `/Repos/mr.foo@example.com`.
* `alphanumeric` - Alphanumeric representation of user local name. e.g. `mr_foo`.
* `workspace_url` - URL of the current Databricks workspace.
* `acl_principal_id` - identifier for use in [databricks_access_control_rule_set](../resources/access_control_rule_set.md), e.g. `users/mr.foo@example.com` if current user is user, or `servicePrincipals/00000000-0000-0000-0000-000000000000` if current user is service principal.

## Related Resources

The following resources are used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide
* [databricks_directory](../resources/directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
* [databricks_notebook](../resources/notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_repo](../resources/repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
