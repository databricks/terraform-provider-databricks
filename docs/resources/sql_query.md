---
subcategory: "Databricks SQL"
---
# databricks_sql_query Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access.

To manage [SQLA resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `allow_sql_analytics_access` on your [databricks_group](group.md#allow_sql_analytics_access) or [databricks_user](user.md#allow_sql_analytics_access).

**Note:** documentation for this resource is a work in progress.

A query may have one or more [visualizations](sql_visualization.md).

## Example Usage

```hcl
resource "databricks_sql_query" "q1" {
  data_source_id = databricks_sql_endpoint.example.data_source_id
  name = "My Query Name"
  query = "SELECT {{ p1 }} AS p1, 2 as p2"
  run_as_role = "viewer"

  schedule {
    continuous {
      interval_seconds = 5 * 60
    }
  }

  parameter {
    name = "p1"
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
