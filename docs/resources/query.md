---
subcategory: "Databricks SQL"
---
# databricks_query Resource

This resource allows you to manage [Databricks SQL Queries](https://docs.databricks.com/en/sql/user/queries/index.html).  It supersedes [databricks_sql_query](sql_query.md) resource - see migration guide below for more details.

## Example Usage

```hcl
resource "databricks_directory" "shared_dir" {
  path = "/Shared/Queries"
}

# This will be replaced with new databricks_query resource
resource "databricks_query" "this" {
  warehouse_id = databricks_sql_endpoint.example.id
  display_name = "My Query Name"
  query_text   = "SELECT 42 as value"
  parent_path  = databricks_directory.shared_dir.path
}
```

## Argument Reference

The following arguments are available:

* `query_text` - (Required, String) Text of SQL query.
* `display_name` - (Required, String) Name of the query.
* `warehouse_id` - (Required, String) ID of a SQL warehouse which will be used to execute this query.
* `parent_path` - (Optional, String) The path to a workspace folder containing the query. The default is the user's home folder.  If changed, the query will be recreated.
* `owner_user_name` - (Optional, String) Query owner's username.
* `apply_auto_limit` - (Optional, Boolean) Whether to apply a 1000 row limit to the query result.
* `catalog` - (Optional, String) Name of the catalog where this query will be executed.
* `schema` - (Optional, String) Name of the schema where this query will be executed.
* `description` - (Optional, String) General description that conveys additional information about this query such as usage notes.
* `run_as_mode` - (Optional, String) Sets the "Run as" role for the object.
* `tags` - (Optional, List of strings) Tags that will be added to the query.
* `parameter` - (Optional, Block) Query parameter definition.  Consists of following attributes (one of `*_value` is required):
  * `name` - (Required, String) Literal parameter marker that appears between double curly braces in the query text.
  * `title` - (Optional, String) Text displayed in the user-facing parameter widget in the UI.
  * `text_value` - (Block) Text parameter value. Consists of following attributes:
    * `value` - (Required, String) - actual text value.
  * `numeric_value` -  (Block) Numeric parameter value. Consists of following attributes:
    * `value` - (Required, Double) - actual numeric value.
  * `date_value` - (Block) Date query parameter value. Consists of following attributes (Can only specify one of `dynamic_date_value` or `date_value`):
    * `date_value` - (String) Manually specified date-time value
    * `dynamic_date_value` - (String) Dynamic date-time value based on current date-time.  Possible values are `NOW`, `YESTERDAY`.
    * `precision` - (Optional, String) Date-time precision to format the value into when the query is run.  Possible values are `DAY_PRECISION`, `MINUTE_PRECISION`, `SECOND_PRECISION`.  Defaults to `DAY_PRECISION` (`YYYY-MM-DD`).
  * `date_range_value` - (Block) Date-range query parameter value. Consists of following attributes (Can only specify one of `dynamic_date_range_value` or `date_range_value`):
    * `date_range_value` - (Block) Manually specified date-time range value.  Consists of the following attributes:
      * `start` (Required, String) - begin of the date range.
      * `end` (Required, String) - end of the date range.
    * `dynamic_date_range_value` - (String) Dynamic date-time range value based on current date-time.  Possible values are `TODAY`, `YESTERDAY`, `THIS_WEEK`, `THIS_MONTH`, `THIS_YEAR`, `LAST_WEEK`, `LAST_MONTH`, `LAST_YEAR`, `LAST_HOUR`, `LAST_8_HOURS`, `LAST_24_HOURS`, `LAST_7_DAYS`, `LAST_14_DAYS`, `LAST_30_DAYS`, `LAST_60_DAYS`, `LAST_90_DAYS`, `LAST_12_MONTHS`.
    * `start_day_of_week` - (Optional, Int) Specify what day that starts the week.
    * `precision` - (Optional, String) Date-time precision to format the value into when the query is run.  Possible values are `DAY_PRECISION`, `MINUTE_PRECISION`, `SECOND_PRECISION`.  Defaults to `DAY_PRECISION` (`YYYY-MM-DD`).
  * `enum_value` - (Block) Dropdown parameter value. Consists of following attributes:
    * `enum_options` - (String) List of valid query parameter values, newline delimited.
    * `values` - (Array of strings) List of selected query parameter values.
    * `multi_values_options` - (Optional, Block) If specified, allows multiple values to be selected for this parameter. Consists of following attributes:
      * `prefix` - (Optional, String) Character that prefixes each selected parameter value.
      * `separator` - (Optional, String) Character that separates each selected parameter value. Defaults to a comma.
      * `suffix` - (Optional, String) Character that suffixes each selected parameter value.
  * `query_backed_value` - (Block) Query-based dropdown parameter value. Consists of following attributes:
    * `query_id` - (Required, String) ID of the query that provides the parameter values.
    * `values` - (Array of strings) List of selected query parameter values.
    * `multi_values_options` - (Optional, Block) If specified, allows multiple values to be selected for this parameter. Consists of following attributes:
      * `prefix` - (Optional, String) Character that prefixes each selected parameter value.
      * `separator` - (Optional, String) Character that separates each selected parameter value. Defaults to a comma.
      * `suffix` - (Optional, String) Character that suffixes each selected parameter value.

## Attribute Reference

In addition to all the arguments above, the following attributes are exported:

* `id` - unique ID of the created Query.
* `lifecycle_state` - The workspace state of the query. Used for tracking trashed status. (Possible values are `ACTIVE` or `TRASHED`).
* `last_modifier_user_name` - Username of the user who last saved changes to this query.
* `create_time` - The timestamp string indicating when the query was created.
* `update_time` - The timestamp string indicating when the query was updated.

## Migrating from `databricks_sql_query` resource

Under the hood, the new resource uses the same data as the `databricks_sql_query`, but exposed via different API. This means that we can migrate existing queries without recreating them.  This operation is done in few steps:

* Record the ID of existing `databricks_sql_query`, for example, by executing the `terraform state show databricks_sql_query.query` command.
* Create the code for the new implementation performing following changes:
  * the `name` attribute is now named `display_name`
  * the `parent` (if exists) is renamed to `parent_path` attribute, and should be converted from `folders/object_id` to the actual path.
  * Blocks that specify values in the `parameter` block were renamed (see above).
  
For example, if we have the original `databricks_sql_query` defined as:

```hcl
resource "databricks_sql_query" "query" {
  data_source_id = databricks_sql_endpoint.example.data_source_id
  query          = "select 42 as value"
  name           = "My Query"
  parent         = "folders/${databricks_directory.shared_dir.object_id}"

  parameter {
    name  = "p1"
    title = "Title for p1"
    text {
      value = "default"
    }
  }
}
```

we'll have a new resource defined as:

```hcl
resource "databricks_query" "query" {
  warehouse_id = databricks_sql_endpoint.example.id
  query_text   = "select 42 as value"
  display_name = "My Query"
  parent_path  = databricks_directory.shared_dir.path

  parameter {
    name  = "p1"
    title = "Title for p1"
    text_value {
      value = "default"
    }
  }
}
```

* Remove the old resource from the state with the `terraform state rm databricks_sql_query.query` command.
* Import new resource with the `terraform import databricks_query.query <query-id>` command.
* Run the `terraform plan` command to check possible changes, like, value type change, etc.

## Access Control

[databricks_permissions](permissions.md#sql-query-usage) can control which groups or individual users can *Manage*, *Edit*, *Run* or *View* individual queries.

```hcl
resource "databricks_permissions" "query_usage" {
  sql_query_id = databricks_query.query.id
  access_control {
    group_name       = "users"
    permission_level = "CAN_RUN"
  }
}
```

## Import

This resource can be imported using query ID:

```bash
terraform import databricks_query.this <query-id>
```

## Related Resources

The following resources are often used in the same context:

* [databricks_sql_query](sql_query.md) to manage Databricks SQL [Queries](https://docs.databricks.com/sql/user/queries/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
