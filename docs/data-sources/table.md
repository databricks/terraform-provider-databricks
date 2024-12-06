---
subcategory: "Unity Catalog"
---
# databricks_table Data Source

-> **Note** This data source can only be used with a workspace-level provider!

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves details of a specific table in Unity Catalog, that were created by Terraform or manually. Use [databricks_tables](tables.md) to retrieve multiple tables in Unity Catalog

## Example Usage

Read  on a specific table `main.certified.fct_transactions`:

```hcl
data "databricks_table" "fct_transactions" {
  name = "main.certified.fct_transactions"
}

resource "databricks_grants" "things" {
  table = data.databricks_table.fct_transactions.name

  grant {
    principal  = "sensitive"
    privileges = ["SELECT", "MODIFY"]
  }
}
```

## Argument Reference

* `name` - (Required) Full name of the databricks_table: _`catalog`.`schema`.`table`_

## Attribute Reference

This data source exports the following attributes:

* `table_info` - TableInfo object for a Unity Catalog table. This contains the following attributes:
  * `name` - Name of table, relative to parent schema.
  * `catalog_name` - Name of parent catalog.
  * `schema_name` - Name of parent schema relative to its parent catalog.
  * `table_type` - Table type, e.g. MANAGED, EXTERNAL, VIEW
  * `data_source_format` - Table format, e.g. DELTA, CSV, JSON
  * `view_definition` - View definition SQL (when `table_type` is VIEW, MATERIALIZED_VIEW, or STREAMING_TABLE)
  * `view_dependencies` - View dependencies (when `table_type` is VIEW or MATERIALIZED_VIEW, STREAMING_TABLE)
  * `columns` - Array of ColumnInfo objects of the table's columns
  * `owner` - Current owner of the table
  * `comment` - Free-form text description

## Related Resources

The following resources are used in the same context:

* [databricks_grant](../resources/grant.md) to manage grants within Unity Catalog.
* [databricks_tables](tables.md) to list all tables within a schema in Unity Catalog.
