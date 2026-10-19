---
subcategory: "Databricks SQL"
---

# databricks_sql_endpoint Resource

[API Documentation](https://docs.databricks.com/api/workspace/warehouses)

This resource is used to manage [Databricks SQL warehouses](https://docs.databricks.com/sql/admin/sql-endpoints.html). To create [SQL warehouses](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

-> This resource can only be used with a workspace-level provider!

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
* `cluster_size` - (Required) The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large", "5X-Large".
* `min_num_clusters` - Minimum number of clusters available when a SQL warehouse is running. The default is `1`. Must be `1` for REALTIME warehouses.
* `max_num_clusters` - Maximum number of clusters available when a SQL warehouse is running. Defaults to `1`. Must be `1` for REALTIME warehouses.
* `max_dbu` - (Optional, REALTIME only) Maximum compute capacity in DBU/hour. Must be a positive integer; the server enforces capacity limits. Omit to preserve the server value; `0` cannot clear it.
* `auto_stop_mins` - Idle minutes before the warehouse stops. Defaults to `120` for PRO/CLASSIC and `10` for REALTIME. Use `0` to disable autostop. REALTIME requires `0` or at least `5`.
* `statement_timeout` - (Optional, REALTIME only) Statement timeout in seconds: a nonnegative integer, default `3600`. Use `0` to clear the warehouse override.
* `tags` - Databricks tags all endpoint resources with these tags.
* `spot_instance_policy` - The spot policy to use for allocating instances to clusters: `COST_OPTIMIZED` or `RELIABILITY_OPTIMIZED`. Defaults to `COST_OPTIMIZED`, the only accepted value for REALTIME.
* `enable_photon` - Whether to enable [Photon](https://databricks.com/product/delta-engine). Enabled by default. Must be `true` for REALTIME.
* `enable_serverless_compute` - Whether to use serverless compute. Set explicitly to avoid workspace-dependent defaults.

  * For REALTIME, defaults to `true` and cannot be set to `false`.
  * For PRO/CLASSIC, defaults to `false`. Workspaces that created a warehouse through the API between September 1, 2022 and April 30, 2023 (November 1, 2022 to May 19, 2023 on Azure) retain a default of `true` if enabled for serverless and meeting its [requirements](https://docs.databricks.com/sql/admin/serverless.html#requirements). Workspace admins may need to accept updated [terms of use](https://docs.databricks.com/sql/admin/serverless.html#accept-terms) in the SQL UI.

* `channel` block, consisting of following fields:
  * `name` - Name of the Databricks SQL release channel. Possible values are: `CHANNEL_NAME_PREVIEW` and `CHANNEL_NAME_CURRENT`. Default is `CHANNEL_NAME_CURRENT`.

* `warehouse_type` - SQL warehouse type. See for [AWS](https://docs.databricks.com/sql/admin/sql-endpoints.html#switch-the-sql-warehouse-type-pro-classic-or-serverless) or [Azure](https://learn.microsoft.com/en-us/azure/databricks/sql/admin/create-sql-warehouse#--upgrade-a-pro-or-classic-sql-warehouse-to-a-serverless-sql-warehouse). Set to `PRO`, `CLASSIC`, or `REALTIME`. REALTIME must be selected explicitly; changing an existing warehouse to or from REALTIME is not supported. When omitted, defaults to `PRO` if `enable_serverless_compute` is `true`, otherwise `CLASSIC`. Serverless PRO/CLASSIC warehouses must use `PRO`.
* `no_wait` - (Optional) Skip waiting for `RUNNING` after creation. Defaults to `false`. REALTIME warehouses are created stopped; explicitly setting either value produces a warning because it has no effect.
* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `id` - the unique ID of the SQL warehouse.
* `jdbc_url` - JDBC connection string.
* `odbc_params` - ODBC connection params: `odbc_params.hostname`, `odbc_params.path`, `odbc_params.protocol`, and `odbc_params.port`.
* `data_source_id` - (Deprecated, will be removed) ID of the data source for this endpoint. This is used to bind a Databricks SQL query to an endpoint. Empty for REALTIME resources.
* `customer_version_label` - (Read-only, REALTIME only) Compute version label; empty if the API omits it.
* `creator_name` - The username of the user who created the endpoint.
* `num_active_sessions` - The current number of clusters used by the endpoint.
* `num_clusters` - The current number of clusters used by the endpoint.
* `state` - The current state of the endpoint.
* `health` - Health status of the endpoint.

## Access control

* [databricks_permissions](permissions.md#job-usage) can control which groups or individual users can *Can Use* or *Can Manage* SQL warehouses.
* `databricks_sql_access` on [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

## Timeouts

`timeouts.create` limits creation time. PRO/CLASSIC provisioning usually takes 10-20 minutes. REALTIME creation returns stopped, without waiting for query readiness.

```hcl
timeouts {
  create = "30m"
}
```

## Import

You can import a `databricks_sql_endpoint` resource with ID like the following:

```hcl
import {
  to = databricks_sql_endpoint.this
  id = "<endpoint-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_sql_endpoint.this <endpoint-id>
```

## Related resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_grants](grant.md) to manage data access in Unity Catalog.
