---
subcategory: "Workspace"
---
# databricks_dashboards Data Source

This data source allows you to retrieve information about Databricks [Dashboards](https://docs.databricks.com/en/dashboards/index.html).

## Example Usage

```hcl
data "databricks_dashboards" "all" {
}

resource "databricks_permissions" "dashboards_permissions" {
  depends = [ data.databricks_dashboards.all ]
  for_each = data.databricks_dashboards.all.dashboards[*].dashboard_id

  dashboard_id = each.value

  access_control {
    group_name       = "Example Group"
    permission_level = "CAN_MANAGE"
  }
}

```

## Argument Reference

The following arguments are supported:

* `dashboard_name_contains` - (Optional) A **case-insensitive** substring to filter Dashboards by their name.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `dashboards` - A list of dashboards matching the specified criteria. Each element contains the following attributes:
  * `dashboard_id` - The unique ID of the dashboard.
  * `display_name` - The display name of the dashboard.
  * `create_time` - The timestamp of when the dashboard was created.

If no matches are found, an empty list will be returned.
