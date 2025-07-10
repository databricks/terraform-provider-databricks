---
subcategory: "Databricks SQL"
---
# databricks_sql_warehouses Data Source

Retrieves a list of [databricks_sql_endpoint](../resources/sql_endpoint.md) ids, that were created by Terraform or manually.

-> This data source can only be used with a workspace-level provider!

## Example Usage

Retrieve IDs for all SQL warehouses:

```hcl
data "databricks_sql_warehouses" "all" {
}
```

Retrieve IDs for all clusters having "Shared" in the warehouse name:

```hcl
data "databricks_sql_warehouses" "all_shared" {
  warehouse_name_contains = "shared"
}
```

## Argument Reference

* `warehouse_name_contains` - (Optional) Only return [databricks_sql_endpoint](../resources/sql_endpoint.md) ids that match the given name string.

## Attribute Reference

This data source exports the following attributes:

* `ids` - list of [databricks_sql_endpoint](../resources/sql_endpoint.md) ids

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_instance_profile](../resources/instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](../resources/mount.md).
* [databricks_sql_dashboard](../resources/sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_global_config](../resources/sql_global_config.md) to configure the security policy, [databricks_instance_profile](../resources/instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_warehouse](sql_warehouse.md) of workspace.
* [databricks_grants](../resources/grant.md) to manage data access in Unity Catalog.