---
subcategory: "Databricks SQL"
---

# databricks_sql_endpoint Resource

This resource is used to manage [Databricks SQL warehouses](https://docs.databricks.com/sql/admin/sql-endpoints.html). To create [SQL warehouses](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

## Example usage

```hcl
data "databricks_current_user" "me" {}

resource "databricks_sql_endpoint" "this" {
  name             = "Endpoint of ${data.databricks_current_user.me.alphanumeric}"
  cluster_size     = "Small"
  max_num_clusters = 1

  tags {
    custom_tags {
      key   = "City"
      value = "Amsterdam"
    }
  }
}
```

## Argument reference

The following arguments are supported:

* `name` - (Required) Name of the SQL warehouse. Must be unique.
* `cluster_size` - (Required) The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".
* `min_num_clusters` - Minimum number of clusters available when a SQL warehouse is running. The default is `1`.
* `max_num_clusters` - Maximum number of clusters available when a SQL warehouse is running. This field is required. If multi-cluster load balancing is not enabled, this is default to `1`.
* `auto_stop_mins` - Time in minutes until an idle SQL warehouse terminates all clusters and stops. This field is optional. The default is 120, set to 0 to disable the auto stop.
* `tags` - Databricks tags all endpoint resources with these tags.
* `spot_instance_policy` - The spot policy to use for allocating instances to clusters: `COST_OPTIMIZED` or `RELIABILITY_OPTIMIZED`. This field is optional. Default is `COST_OPTIMIZED`.
* `enable_photon` - Whether to enable [Photon](https://databricks.com/product/delta-engine). This field is optional and is enabled by default.
* `enable_serverless_compute` - Whether this SQL warehouse is a serverless endpoint. See below for details about the default values. To avoid ambiguity, especially for organizations with many workspaces, Databricks recommends that you always set this field explicitly.

  * **For AWS**, If omitted, the default is `false` for most workspaces. However, if this workspace used the SQL Warehouses API to create a warehouse between September 1, 2022 and April 30, 2023, the default remains the previous behavior which is default to `true` if the workspace is enabled for serverless and fits the requirements for serverless SQL warehouses. If your account needs updated [terms of use](https://docs.databricks.com/sql/admin/serverless.html#accept-terms), workspace admins are prompted in the Databricks SQL UI. A workspace must meet the [requirements](https://docs.databricks.com/sql/admin/serverless.html#requirements) and might require an update to its instance profile role to [add a trust relationship](https://docs.databricks.com/sql/admin/serverless.html#aws-instance-profile-setup).

  * **For Azure**, If omitted, the default is `false` for most workspaces. However, if this workspace used the SQL Warehouses API to create a warehouse between November 1, 2022 and May 19, 2023, the default remains the previous behavior which is default to `true` if the workspace is enabled for serverless and fits the requirements for serverless SQL warehouses. A workspace must meet the [requirements](https://learn.microsoft.com/azure/databricks/sql/admin/serverless) and might require an update to its [Azure storage firewall](https://learn.microsoft.com/azure/databricks/sql/admin/serverless-firewall).

* `channel` block, consisting of following fields:
  * `name` - Name of the Databricks SQL release channel. Possible values are: `CHANNEL_NAME_PREVIEW` and `CHANNEL_NAME_CURRENT`. Default is `CHANNEL_NAME_CURRENT`.

* `warehouse_type` - SQL warehouse type. See for [AWS](https://docs.databricks.com/sql/admin/sql-endpoints.html#switch-the-sql-warehouse-type-pro-classic-or-serverless) or [Azure](https://learn.microsoft.com/en-us/azure/databricks/sql/admin/create-sql-warehouse#--upgrade-a-pro-or-classic-sql-warehouse-to-a-serverless-sql-warehouse). Set to `PRO` or `CLASSIC`. If the field `enable_serverless_compute` has the value `true` either explicitly or through the default logic (see that field above for details), the default is `PRO`, which is required for serverless SQL warehouses. Otherwise, the default is `CLASSIC`.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `id` - the unique ID of the SQL warehouse.
* `jdbc_url` - JDBC connection string.
* `odbc_params` - ODBC connection params: `odbc_params.hostname`, `odbc_params.path`, `odbc_params.protocol`, and `odbc_params.port`.
* `data_source_id` - ID of the data source for this endpoint. This is used to bind an Databricks SQL query to an endpoint.
* `creator_name` - The username of the user who created the endpoint.
* `num_active_sessions` - The current number of clusters used by the endpoint.
* `num_clusters` - The current number of clusters used by the endpoint.
* `state` - The current state of the endpoint.
* `health` - Health status of the endpoint.

## Access control

* [databricks_permissions](permissions.md#job-usage) can control which groups or individual users can *Can Use* or *Can Manage* SQL warehouses.
* `databricks_sql_access` on [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

## Timeouts

The `timeouts` block allows you to specify `create` timeouts. It usually takes 10-20 minutes to provision a Databricks SQL warehouse.

```hcl
timeouts {
  create = "30m"
}
```

## Import

You can import a `databricks_sql_endpoint` resource with ID like the following:

```bash
terraform import databricks_sql_endpoint.this <endpoint-id>
```

## Related resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
