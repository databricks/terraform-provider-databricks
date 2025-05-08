---
subcategory: "Databricks SQL"
---
# databricks_sql_query Resource

!> This resource is deprecated! Please switch to [databricks_query](query.md#migrating-from-databricks_sql_query-resource).

To manage [SQLA resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

-> documentation for this resource is a work in progress.

A query may have one or more [visualizations](sql_visualization.md).

## Example Usage

```hcl
resource "databricks_directory" "shared_dir" {
  path = "/Shared/Queries"
}

resource "databricks_sql_query" "q1" {
  data_source_id = databricks_sql_endpoint.example.data_source_id
  name           = "My Query Name"
  query          = <<EOT
                        SELECT {{ p1 }} AS p1
                        WHERE 1=1
                        AND p2 in ({{ p2 }})
                        AND event_date > date '{{ p3 }}'
                    EOT

  parent      = "folders/${databricks_directory.shared_dir.object_id}"
  run_as_role = "viewer"

  parameter {
    name  = "p1"
    title = "Title for p1"
    text {
      value = "default"
    }
  }

  parameter {
    name  = "p2"
    title = "Title for p2"
    enum {
      options = ["default", "foo", "bar"]
      value   = "default"
      // passes to sql query as string `"foo", "bar"` if foo and bar are both selected in the front end
      multiple {
        prefix    = "\""
        suffix    = "\""
        separator = ","
      }

    }
  }

  parameter {
    name  = "p3"
    title = "Title for p3"
    date {
      value = "2022-01-01"
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

## Argument Reference

The following arguments are supported:

* `data_source_id` - Data source ID of a [SQL warehouse](sql_endpoint.md)
* `query` - The text of the query to be run.
* `name` - The title of this query that appears in list views, widget headings, and on the query page.
* `parent` - The identifier of the workspace folder containing the object.
* `description` - General description that conveys additional information about this query such as usage notes.
* `run_as_role` - Run as role. Possible values are `viewer`, `owner`.

### `parameter` configuration block

For parameter definition

* `title` - The text displayed in a parameter picking widget.
* `name` - The literal parameter marker that appears between double curly braces in the query text.
Parameters can have several different types. Type is specified using one of the following configuration blocks: `text`, `number`, `enum`, `query`, `date`, `datetime`, `datetimesec`, `date_range`, `datetime_range`, `datetimesec_range`.

For `text`, `number`, `date`, `datetime`, `datetimesec` block

* `value` - The default value for this parameter.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the unique ID of the SQL Query.

## Import

You can import a `databricks_sql_query` resource with ID like the following:

```bash
terraform import databricks_sql_query.this <query-id>
```

## Troubleshooting

In case you see `Error: cannot create sql query: Internal Server Error` during `terraform apply`; double check that you are using the correct [`data_source_id`](sql_endpoint.md)

Operations on `databricks_sql_query` schedules are ⛔️ deprecated. You can create, update or delete a schedule for SQLA and other Databricks resources using the [databricks_job](job.md#sql_task-configuration-block) resource.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
* [databricks_job](job.md#sql_task-configuration-block) to schedule Databricks SQL queries (as well as dashboards and alerts) using Databricks Jobs.
