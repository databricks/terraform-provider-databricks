---
subcategory: "SQL Analytics"
---
# databricks_sql_endpoint Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access.

To create [SQL endpoints](https://docs.databricks.com/sql/get-started/concepts.html) you must have `allow_sql_analytics_access` on your [databricks_group](group.md#allow_sql_analytics_access) or [databricks_user](user.md#allow_sql_analytics_access).

## Example usage

```hcl
data "databricks_current_user" "me" {}

resource "databricks_sql_endpoint" "this" {
  name = "Endpoint of ${data.databricks_current_user.me.alphanumeric}"
  cluster_size = "Small"
  max_num_clusters = 1

  tags {
    custom_tags {
        key = "City"
        value = "Amsterdam"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the SQL endpoint. Must be unique.
* `cluster_size` - (Required) The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".
* `min_num_clusters` - Minimum number of clusters available when a SQL endpoint is running. The default is 1.
* `max_num_clusters` - Maximum number of clusters available when a SQL endpoint is running. This field is required. If multi-cluster load balancing is not enabled, this is default to 1.
* `auto_stop_mins` - Time in minutes until an idle SQL endpoint terminates all clusters and stops. This field is optional. The default is 0, which means auto stop is disabled.
* `instance_profile_arn` - [databricks_instance_profile](instance_profile.md) used to access storage from the SQL endpoint. This field is optional.
* `tags` - Databricks tags all endpoint resources with these tags.
* `spot_instance_policy` - The spot policy to use for allocating instances to clusters: `COST_OPTIMIZED` or `RELIABILITY_OPTIMIZED`. This field is optional.
* `enable_photon` - Whether to enable [Photon](https://databricks.com/product/delta-engine). This field is optional.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `jdbc_url` - JDBC connection string.
* `odbc_params` - ODBC connection params: `odbc_params.host`, `odbc_params.path`, `odbc_params.protocol`, and `odbc_params.port`.

## Access Control

* [databricks_permissions](permissions.md#Job-Endpoint-usage) can control which groups or individual users can *Can Use* or *Can Manage* SQL endpoints.
* `allow_sql_analytics_access` on [databricks_group](group.md#allow_sql_analytics_access) or [databricks_user](user.md#allow_sql_analytics_access).

## Timeouts

The `timeouts` block allows you to specify `create` timeouts. It usually takes 10-20 minutes to provision SQL Analytics endpoint.

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
