---
subcategory: "Databricks SQL"
---
# databricks_sql_alert Resource

!> This resource is deprecated! Please switch to [databricks_alert](alert.md#migrating-from-databricks_sql_alert-resource).

This resource allows you to manage [Databricks SQL Alerts](https://docs.databricks.com/sql/user/queries/index.html).

-> To manage [SQLA resources](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

## Example Usage

```hcl
resource "databricks_directory" "shared_dir" {
  path = "/Shared/Queries"
}

resource "databricks_sql_query" "this" {
  data_source_id = databricks_sql_endpoint.example.data_source_id
  name           = "My Query Name"
  query          = "SELECT 1 AS p1, 2 as p2"
  parent         = "folders/${databricks_directory.shared_dir.object_id}"
}

resource "databricks_sql_alert" "alert" {
  query_id = databricks_sql_query.this.id
  name     = "My Alert"
  parent   = "folders/${databricks_directory.shared_dir.object_id}"
  rearm    = 1
  options {
    column = "p1"
    op     = "=="
    value  = "2"
    muted  = false
  }
}
```

## Argument Reference

The following arguments are available:

* `query_id` - (Required, String) ID of the query evaluated by the alert.
* `name` - (Required, String) Name of the alert.
* `options` - (Required) Alert configuration options.
  * `column` - (Required, String) Name of column in the query result to compare in alert evaluation.
  * `op` - (Required, String Enum) Operator used to compare in alert evaluation. (Enum: `>`, `>=`, `<`, `<=`, `==`, `!=`)
  * `value` - (Required, String) Value used to compare in alert evaluation.
  * `muted` - (Optional, bool) Whether or not the alert is muted. If an alert is muted, it will not notify users and alert destinations when triggered.
  * `custom_subject` - (Optional, String) Custom subject of alert notification, if it exists. This includes email subject, Slack notification header, etc. See [Alerts API reference](https://docs.databricks.com/sql/user/alerts/index.html) for custom templating instructions.
  * `custom_body` - (Optional, String) Custom body of alert notification, if it exists. See [Alerts API reference](https://docs.databricks.com/sql/user/alerts/index.html) for custom templating instructions.
  * `empty_result_state` - (Optional, String) State that alert evaluates to when query result is empty.  Currently supported values are `unknown`, `triggered`, `ok` - check [API documentation](https://docs.databricks.com/api/workspace/alerts/create) for full list of supported values.
* `parent` - (Optional, String) The identifier of the workspace folder containing the alert. The default is ther user's home folder. The folder identifier is formatted as `folder/<folder_id>`.
* `rearm` - (Optional, Integer) Number of seconds after being triggered before the alert rearms itself and can be triggered again. If not defined, alert will never be triggered again.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - unique ID of the SQL Alert.

## Access Control

[databricks_permissions](permissions.md#sql-alert-usage) can control which groups or individual users can *Manage*, *Edit*, *Run* or *View* individual alerts.

## Import

This resource can be imported using alert ID:

```bash
terraform import databricks_sql_alert.this <alert-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_sql_query](sql_query.md) to manage Databricks SQL [Queries](https://docs.databricks.com/sql/user/queries/index.html).
* [databricks_sql_endpoint](sql_endpoint.md) to manage Databricks SQL [Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html).
* [databricks_directory](directory.md) to manage directories in [Databricks Workpace](https://docs.databricks.com/workspace/workspace-objects.html).
