---
subcategory: "Databricks SQL"
---
# databricks_sql_visualization Resource

To manage [SQLA resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

-> documentation for this resource is a work in progress.

A visualization is always tied to a [query](sql_query.md). Every query may have one or more visualizations.

## Example Usage

```hcl
resource "databricks_sql_visualization" "q1v1" {
  query_id    = databricks_sql_query.q1.id
  type        = "table"
  name        = "My Table"
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

## Separating `visualization definition` from IAC configuration

Since `options` field contains the full JSON encoded string definition of how to render a visualization for the backend API - `sql/api/visualizations`, they can get quite verbose.

If you have lots of visualizations to declare, it might be cleaner to separate the `options` field and store them as separate `.json` files to be referenced.

### Example

- directory tree

    ```bash
    .
    ├── q1vx.tf
    └── visualizations
        ├── q1v1.json
        └── q1v2.json
    ```

- resource definitions
  
    ```hcl
    ##q1vx.tf

    resource "databricks_sql_visualization" "q1v1" {
      query_id    = databricks_sql_query.q1.id
      type        = "table"
      name        = "My Table"
      description = "Some Description"
      options     = file("${path.module}/visualizations/q1v1.json")
    }

    resource "databricks_sql_visualization" "q1v2" {
      query_id    = databricks_sql_query.q1.id
      type        = "chart"
      name        = "My Chart"
      description = "Some Description"
      options     = file("${path.module}/visualizations/q1v2.json")
    }
    ```

## Known Issues

As of 2022-09, databricks sql visualization backend API does not validate the content of what is passed via `options`, couple that with `options` being outputted as string in the module, it can lead to configurations which succeed `terraform plan` but do fail at `terraform apply`.

In some instances, incorrect definitions within `options` can [lead to stuck terraform states](https://github.com/databricks/terraform-provider-databricks/issues/1615).
In preparation for this operational scenario; you should be familiar with, and have sufficient access for, manual inspection and modification of your deployed [terraform state](https://www.terraform.io/language/state).

## Import

You can import a `databricks_sql_visualization` resource with ID like the following:

```hcl
import {
  to = databricks_sql_visualization.this
  id = "<query-id>/<visualization-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_sql_visualization.this <query-id>/<visualization-id>
```

## Related Resources

The following resources are often used in the same context:

- [End to end workspace management](../guides/workspace-management.md) guide.
- [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
- [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
- [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
- [databricks_grants](grant.md) to manage data access in Unity Catalog.
