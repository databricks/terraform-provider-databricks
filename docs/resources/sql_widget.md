---
subcategory: "Databricks SQL"
---
# databricks_sql_widget Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access.

To manage [SQLA resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `allow_sql_analytics_access` on your [databricks_group](group.md#allow_sql_analytics_access) or [databricks_user](user.md#allow_sql_analytics_access).

**Note:** documentation for this resource is a work in progress.

A widget is always tied to a [dashboard](sql_dashboard.md). Every dashboard may have one or more widgets.

## Example Usage

```
resource "databricks_sql_widget" "d1w1" {
  dashboard_id = databricks_sql_dashboard.d1.id
  text = "Hello! I'm a **text widget**!"

  position {
    size_x = 3
    size_y = 4
    pos_x = 0
    pos_y = 0
  }
}

resource "databricks_sql_widget" "d1w2" {
  dashboard_id = databricks_sql_dashboard.d1.id
  visualization_id = databricks_sql_visualization.q1v1.id

  position {
    size_x = 3
    size_y = 4
    pos_x = 3
    pos_y = 0
  }
}
```
