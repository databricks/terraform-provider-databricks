---
subcategory: "Databricks SQL"
---
# databricks_sql_widget Resource

-> Please switch to [databricks_dashboard](dashboard.md) to author new AI/BI dashboards using the latest tooling

To manage [SQL resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

-> documentation for this resource is a work in progress.

A widget is always tied to a [Legacy dashboard](sql_dashboard.md). Every dashboard may have one or more widgets.

## Example Usage

```hcl
resource "databricks_sql_widget" "d1w1" {
  dashboard_id = databricks_sql_dashboard.d1.id
  text         = "Hello! I'm a **text widget**!"

  position {
    size_x = 3
    size_y = 4
    pos_x  = 0
    pos_y  = 0
  }
}

resource "databricks_sql_widget" "d1w2" {
  dashboard_id     = databricks_sql_dashboard.d1.id
  visualization_id = databricks_sql_visualization.q1v1.id

  position {
    size_x = 3
    size_y = 4
    pos_x  = 3
    pos_y  = 0
  }
}
```

## Import

You can import a `databricks_sql_widget` resource with ID like the following:

```hcl
import {
  to = databricks_sql_widget.this
  id = "<dashboard-id>/<widget-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_sql_widget.this <dashboard-id>/<widget-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_grants](grant.md) to manage data access in Unity Catalog.
