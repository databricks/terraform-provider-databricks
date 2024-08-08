---
subcategory: "Workspace"
---
# databricks_dashboard Resource

This resource allows you to manage Databricks [Dashboards](https://docs.databricks.com/en/dashboards/index.html). To manage [Dashboards](https://docs.databricks.com/en/dashboards/index.html) you must have a warehouse access on your databricks workspace.

## Example Usage

Dashboard using `serialized_dashboard` attribute:

```hcl
data "databricks_sql_warehouse" "starter" {
    name = "Starter Warehouse"
}

resource "databricks_dashboard" "dashboard" {
    display_name	 = "New Dashboard"
    warehouse_id	 = data.databricks_sql_warehouse.starter.id
    serialized_dashboard = "{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
    embed_credentials    = false // Optional
    parent_path	         = "/Shared/provider-test"
}
```

Dashboard using `file_path` attribute:

```hcl
data "databricks_sql_warehouse" "starter" {
    name = "Starter Warehouse"
}

resource "databricks_dashboard" "dashboard" {
    display_name         = "New Dashboard"
    warehouse_id         = data.databricks_sql_warehouse.starter.id
    file_path	         = "${path.module}/dashboard.json"
    embed_credentials    = false // Optional
    parent_path	         = "/Shared/provider-test"
}
```


## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The display name of the dashboard.
* `warehouse_id` - (Required) The warehouse ID used to run the dashboard.
* `serialized_dashboard` - (Optional) The contents of the dashboard in serialized string form. Conflicts with `file_path`.
* `file_path` - (Optional) The path to the dashboard JSON file. Conflicts with `serialized_dashboard`.
* `embed_credentials` - (Optional) Whether to embed credentials in the dashboard. Default is `true`.
* `parent_path` - (Required) The workspace path of the folder containing the dashboard. Includes leading slash and no trailing slash.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID of the dashboard.

## Access Control

[databricks_permissions](permissions.md#dashboard-usage) can control which groups or individual users can *Manage*, *Edit*, *Read* or *Run* individual dashboards.

## Import

You can import a `databricks_dashboard` resource with ID like the following:

```bash
terraform import databricks_dashboard.this <dashboard-id>
```

## Notes
* Only one of `serialized_dashboard` or `file_path` can be used throughout the lifecycle of the dashboard. If you want to switch from one to the other, you must first destroy the dashboard resource and then recreate it with the new attribute.
* Dashboards managed by Terraform will be published automatically.
