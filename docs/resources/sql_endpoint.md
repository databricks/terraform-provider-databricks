---
subcategory: "Databricks SQL"
---
# databricks_sql_endpoint Resource

This resource is used to manage [Databricks SQL Endpoints](https://docs.databricks.com/sql/admin/sql-endpoints.html). To create [SQL endpoints](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

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

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the SQL endpoint. Must be unique.
* `cluster_size` - (Required) The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".
* `min_num_clusters` - Minimum number of clusters available when a SQL endpoint is running. The default is `1`.
* `max_num_clusters` - Maximum number of clusters available when a SQL endpoint is running. This field is required. If multi-cluster load balancing is not enabled, this is default to `1`.
* `auto_stop_mins` - Time in minutes until an idle SQL endpoint terminates all clusters and stops. This field is optional. The default is 120, set to 0 to disable the auto stop.
* `tags` - Databricks tags all endpoint resources with these tags.
* `spot_instance_policy` - The spot policy to use for allocating instances to clusters: `COST_OPTIMIZED` or `RELIABILITY_OPTIMIZED`. This field is optional. Default is `COST_OPTIMIZED`.
* `enable_photon` - Whether to enable [Photon](https://databricks.com/product/delta-engine). This field is optional and is enabled by default.
* `enable_serverless_compute` - Whether this SQL endpoint is a Serverless endpoint. To use a Serverless SQL endpoint, you must enable Serverless SQL endpoints for the workspace.
* `channel` block, consisting of following fields:
  * `name` - Name of the Databricks SQL release channel. Possible values are: `CHANNEL_NAME_PREVIEW` and `CHANNEL_NAME_CURRENT`. Default is `CHANNEL_NAME_CURRENT`.
* `warehouse_type` - [SQL Warehouse Type](https://docs.databricks.com/sql/admin/sql-endpoints.html#switch-the-sql-warehouse-type-pro-classic-or-serverless): `PRO` or `CLASSIC` (default).  If Serverless SQL is enabled, you can only specify `PRO`.
 
## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `jdbc_url` - JDBC connection string.
* `odbc_params` - ODBC connection params: `odbc_params.hostname`, `odbc_params.path`, `odbc_params.protocol`, and `odbc_params.port`.
* `data_source_id` - ID of the data source for this endpoint. This is used to bind an Databricks SQL query to an endpoint.

## Access Control

* [databricks_permissions](permissions.md#Job-Endpoint-usage) can control which groups or individual users can *Can Use* or *Can Manage* SQL endpoints.
* `databricks_sql_access` on [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

## Timeouts

The `timeouts` block allows you to specify `create` timeouts. It usually takes 10-20 minutes to provision a Databricks SQL endpoint.

```hcl
timeouts {
  create = "30m"
}
```

## Import

You can import a `databricks_sql_endpoint` resource with ID like the following:

```bash
$ terraform import databricks_sql_endpoint.this <endpoint-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
