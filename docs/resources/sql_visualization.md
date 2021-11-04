---
subcategory: "Databricks SQL"
---
# databricks_sql_visualization Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html).

To manage [SQLA resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `allow_sql_analytics_access` on your [databricks_group](group.md#allow_sql_analytics_access) or [databricks_user](user.md#allow_sql_analytics_access).

**Note:** documentation for this resource is a work in progress.

A visualization is always tied to a [query](sql_query.md). Every query may have one or more visualizations.

## Example Usage

```hcl
resource "databricks_sql_visualization" "q1v1" {
  query_id = databricks_sql_query.q1.id
  type = "table"
  name = "My Table"
  description = "Some Description"

  // The options encoded in this field are passed verbatim to the SQLA API.
  options = jsonencode(
    {
      "itemsPerPage" : 25,
      "columns" : [
        {
          "name" : "p1",
          "type" : "string"
          "title" : "Parameter 1",
          "displayAs" : "string",
        },
        {
          "name" : "p2",
          "type" : "string"
          "title" : "Parameter 2",
          "displayAs" : "link",
          "highlightLinks" : true,
        }
      ]
    }
  )
}
```
