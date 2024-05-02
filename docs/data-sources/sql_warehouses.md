---
subcategory: "Databricks SQL"
---
# databricks_sql_warehouses Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves a list of [databricks_sql_endpoint](../resources/sql_endpoint.md#id) ids, that were created by Terraform or manually.

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

* `warehouse_name_contains` - (Optional) Only return [databricks_sql_endpoint](../resources/sql_endpoint.md#id) ids that match the given name string.

## Attribute Reference

This data source exports the following attributes:

* `ids` - list of [databricks_sql_endpoint](../resources/sql_endpoint.md#id) ids

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_sql_dashboard](sql_dashboard.md) to manage Databricks SQL [Dashboards](https://docs.databricks.com/sql/user/dashboards/index.html).
* [databricks_sql_global_config](sql_global_config.md) to configure the security policy, [databricks_instance_profile](instance_profile.md), and [data access properties](https://docs.databricks.com/sql/admin/data-access-configuration.html) for all [databricks_sql_endpoint](sql_endpoint.md) of workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
