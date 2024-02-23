---
subcategory: "Unity Catalog"
---
# databricks_schemas Data Source

-> **Note** This data source could be only used with workspace-level provider!

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves a list of [databricks_schema](../resources/schema.md) ids, that were created by Terraform or manually, so that special handling could be applied.

## Example Usage

Listing all schemas in a _sandbox_ [databricks_catalog](../resources/catalog.md):

```hcl
data "databricks_schemas" "sandbox" {
  catalog_name = "sandbox"
}

output "all_sandbox_schemas" {
  value = data.databricks_schemas.sandbox
}
```

## Argument Reference

* `catalog_name` - (Required) Name of [databricks_catalog](../resources/catalog.md)

## Attribute Reference

This data source exports the following attributes:

* `ids` - set of [databricks_schema](../resources/schema.md) full names: *`catalog`.`schema`*

## Related Resources

The following resources are used in the same context:

* [databricks_schema](../resources/schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
