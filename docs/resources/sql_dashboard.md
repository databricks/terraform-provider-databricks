---
subcategory: "Databricks SQL"
---
# databricks_sql_dashboard Resource

!> This resource is deprecated! Please switch to [databricks_dashboard](dashboard.md) to author new AI/BI dashboards using the latest tooling.

This resource is used to manage [Legacy dashboards](https://docs.databricks.com/sql/user/dashboards/index.html). To manage [SQL resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

-> documentation for this resource is a work in progress.

A dashboard may have one or more [widgets](sql_widget.md).

## Example Usage

```hcl
resource "databricks_directory" "shared_dir" {
  path = "/Shared/Dashboards"
}

resource "databricks_sql_dashboard" "d1" {
  name   = "My Dashboard Name"
  parent = "folders/${databricks_directory.shared_dir.object_id}"

  tags = [
    "some-tag",
    "another-tag",
  ]
}
```

Example [permission](permissions.md) to share dashboard with all users:

```hcl
resource "databricks_permissions" "d1" {
  sql_dashboard_id = databricks_sql_dashboard.d1.id

  access_control {
    group_name       = data.databricks_group.users.display_name
    permission_level = "CAN_RUN"
  }
}
```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the unique ID of the SQL Dashboard.

## Import

You can import a `databricks_sql_dashboard` resource with ID like the following:

```hcl
import {
  to = databricks_sql_dashboard.this
  id = "<dashboard-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_sql_dashboard.this "<dashboard-id>"
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_grants](grant.md) to manage data access in Unity Catalog.
