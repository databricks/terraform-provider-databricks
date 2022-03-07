---
subcategory: "Databricks SQL"
---
# databricks_sql_query Resource

To manage [SQLA resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

**Note:** documentation for this resource is a work in progress.

A query may have one or more [visualizations](sql_visualization.md).

## Example Usage

```hcl
resource "databricks_sql_query" "q1" {
  data_source_id = databricks_sql_endpoint.example.data_source_id
  name           = "My Query Name"
  query          = "SELECT {{ p1 }} AS p1, 2 as p2"
  run_as_role    = "viewer"

  schedule {
    continuous {
      interval_seconds = 5 * 60
    }
  }

  parameter {
    name  = "p1"
    title = "Title for p1"
    text {
      value = "default"
    }
  }

  tags = [
    "t1",
    "t2",
  ]
}
```

Example [permission](permissions.md) to share query with all users:

```hcl
resource "databricks_permissions" "q1" {
  sql_query_id = databricks_sql_query.q1.id

  access_control {
    group_name       = data.databricks_group.users.display_name
    permission_level = "CAN_RUN"
  }

  // You can only specify "CAN_EDIT" permissions if the query `run_as_role` equals `viewer`.
  access_control {
    group_name       = data.databricks_group.team.display_name
    permission_level = "CAN_EDIT"
  }
}
```

## Import

-> **Note** Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
