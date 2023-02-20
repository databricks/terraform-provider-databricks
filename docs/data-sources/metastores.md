---
subcategory: "Unity Catalog"
---
# databricks_metastores Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

Retrieves a list of [databricks_metastore](../resources/metastore.md) objects, that were created by Terraform or manually, so that special handling could be applied.

## Example Usage

Listing all catalogs:

```hcl
data "databricks_metastores" "all" {}

output "all_metastores" {
  value = data.databricks_metastores.all
}
```

## Attribute Reference

This data source exports the following attributes:

* `metastores` - list of [databricks_metastore](../resources/metastore.md)

## Related Resources

The following resources are used in the same context:

* [databricks_metastore](../resources/metastore.md) to manage Metastores within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
