---
subcategory: "Databricks SQL"
---
# databricks_sql_warehouse Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about a [databricks_sql_warehouse](../resources/sql_endpoint.md) using its id. This could be retrieved programmatically using [databricks_sql_warehouses](../data-sources/sql_warehouses.md) data source.

## Example usage

* Retrieve attributes of each SQL warehouses in a workspace:

```hcl
data "databricks_sql_warehouses" "all" {
}

data "databricks_sql_warehouse" "this" {
  for_each = data.databricks_sql_warehouses.all.ids
  id       = each.value
}
```

* Search for a specific SQL Warehouse by name:

```hcl
data "databricks_sql_warehouse" "all" {
  name = "Starter Warehouse"
}
```

## Argument reference

* `id` - (Required, if `name` isn't specified) The ID of the SQL warehouse.
* `name` - (Required, if `id` isn't specified) Name of the SQL warehouse to search (case-sensitive).

## Attribute reference

This data source exports the following attributes:

* `id` - The ID of the SQL warehouse.
* `name` - Name of the SQL warehouse.
* `cluster_size` - The size of the clusters allocated to the warehouse: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".
* `min_num_clusters` - Minimum number of clusters available when a SQL warehouse is running.
* `max_num_clusters` - Maximum number of clusters available when a SQL warehouse is running.
* `auto_stop_mins` - Time in minutes until an idle SQL warehouse terminates all clusters and stops.
* `tags` - tags used for SQL warehouse resources.
* `spot_instance_policy` - The spot policy to use for allocating instances to clusters: `COST_OPTIMIZED` or `RELIABILITY_OPTIMIZED`.
* `enable_photon` - Whether [Photon](https://databricks.com/product/delta-engine) is enabled.
* `enable_serverless_compute` - Whether this SQL warehouse is a serverless SQL warehouse.

  * **For AWS**:  If your account needs updated [terms of use](https://docs.databricks.com/sql/admin/serverless.html#accept-terms), workspace admins are prompted in the Databricks SQL UI. A workspace must meet the [requirements](https://docs.databricks.com/sql/admin/serverless.html#requirements) and might require an update its instance profile role to [add a trust relationship](https://docs.databricks.com/sql/admin/serverless.html#aws-instance-profile-setup).

  * **For Azure**, you must [enable your workspace for serverless SQL warehouse](https://learn.microsoft.com/azure/databricks/sql/admin/serverless).

* `warehouse_type` - SQL warehouse type. See for [AWS](https://docs.databricks.com/sql/index.html#warehouse-types) or [Azure](https://learn.microsoft.com/azure/databricks/sql/#warehouse-types).
* `channel` block, consisting of following fields:
  * `name` - Name of the Databricks SQL release channel. Possible values are: `CHANNEL_NAME_PREVIEW` and `CHANNEL_NAME_CURRENT`. Default is `CHANNEL_NAME_CURRENT`.
* `jdbc_url` - JDBC connection string.
* `odbc_params` - ODBC connection params: `odbc_params.hostname`, `odbc_params.path`, `odbc_params.protocol`, and `odbc_params.port`.
* `data_source_id` - ID of the data source for this warehouse. This is used to bind an Databricks SQL query to an warehouse.
* `creator_name` - The username of the user who created the endpoint.
* `num_active_sessions` - The current number of clusters used by the endpoint.
* `num_clusters` - The current number of clusters used by the endpoint.
* `state` - The current state of the endpoint.
* `health` - Health status of the endpoint.

## Related resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_instance_profile](../resources/instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](../resources/mount.md).
* [databricks_sql_dashboard](../resources/sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_global_config](../resources/sql_global_config.md) to configure the security policy, [databricks_instance_profile](../resources/instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_warehouse](sql_warehouse.md) of workspace.
* [databricks_sql_permissions](../resources/sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
