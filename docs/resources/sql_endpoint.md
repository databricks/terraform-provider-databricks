---
subcategory: "Databricks SQL"
---
# databricks_sql_endpoint Resource

To create [SQL endpoints](https://docs.databricks.com/sql/get-started/concepts.html) you must have `databricks_sql_access` on your [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

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
* `auto_stop_mins` - Time in minutes until an idle SQL endpoint terminates all clusters and stops. This field is optional. The default is 0, which means auto stop is disabled.
* `instance_profile_arn` - [databricks_instance_profile](instance_profile.md) used to access storage from the SQL endpoint. This field is optional.
* `tags` - Databricks tags all endpoint resources with these tags.
* `spot_instance_policy` - The spot policy to use for allocating instances to clusters: `COST_OPTIMIZED` or `RELIABILITY_OPTIMIZED`. This field is optional. Default is `COST_OPTIMIZED`.
* `enable_photon` - Whether to enable [Photon](https://databricks.com/product/delta-engine). This field is optional and is enabled by default.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `jdbc_url` - JDBC connection string.
* `odbc_params` - ODBC connection params: `odbc_params.hostname`, `odbc_params.path`, `odbc_params.protocol`, and `odbc_params.port`.
* `data_source_id` - ID of the data source for this endpoint. This is used to bind an SQLA query to an endpoint.

## Access Control

* [databricks_permissions](permissions.md#Job-Endpoint-usage) can control which groups or individual users can *Can Use* or *Can Manage* SQL endpoints.
* `databricks_sql_access` on [databricks_group](group.md#databricks_sql_access) or [databricks_user](user.md#databricks_sql_access).

## Timeouts

The `timeouts` block allows you to specify `create` timeouts. It usually takes 10-20 minutes to provision Databricks SQL endpoint.

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
