---
subcategory: "Workspace"
---
# databricks_lakeview_dashboard Resource

This resource allows you to manage Databricks [Lakeview Dashboards](https://docs.databricks.com/api/workspace/lakeview). To manage [Lakeview Dashboards](https://docs.databricks.com/api/workspace/lakeview) you must have a warehouse access on your databricks workspace.

## Example Usage

Dashboard using `serialized_dashboard` attribute:

```hcl
resource "databricks_lakeview_dashboard" "dashboard" {
			display_name			= 	"New Dashboard"
			warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
			serialized_dashboard	        =	"{\"pages\":[{\"name\":\"new_name\",\"displayName\":\"New Page\"}]}"
			embed_credentials		=	false // Optional
			parent_path			= 	"/Shared/provider-test"
		}
```

Dashboard using `file_path` attribute:

```hcl
resource "databricks_lakeview_dashboard" "dashboard" {
            display_name			= 	"New Dashboard"
            warehouse_id			=	"{env.TEST_DEFAULT_WAREHOUSE_ID}"
            file_path				=	"${path.module}/dashboard.json"
            embed_credentials   		=	false // Optional
            parent_path				= 	"/Shared/provider-test"
        }
```


## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The display name of the Lakeview dashboard.
* `warehouse_id` - (Required) The warehouse ID used to run the dashboard.
* `serialized_dashboard` - (Optional) The contents of the dashboard in serialized string form. Conflicts with `file_path`.
* `file_path` - (Optional) The path to the dashboard JSON file. Conflicts with `serialized_dashboard`.
* `embed_credentials` - (Optional) Whether to embed credentials in the dashboard. Default is `true`.
* `parent_path` - (Required) The workspace path of the folder containing the dashboard. Includes leading slash and no trailing slash.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID of the Lakeview dashboard.

## Notes
* Currently, only one of `serialized_dashboard` or `file_path` can be used throughout the lifecycle of the dashboard. If you want to switch from one to the other, you must first destroy the dashboard resource and then recreate it with the new attribute.
* Currently, all the dashboards managed by Terraform will be published by default. There is no option to keep the dashboards in draft state. If required in future, functionality to manage the publish status of the dashboard will be added.