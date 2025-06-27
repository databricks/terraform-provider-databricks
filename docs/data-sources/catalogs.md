---
subcategory: "Unity Catalog"
---
# databricks_catalogs Data Source

Retrieves a list of [databricks_catalog](../resources/catalog.md) ids, that were created by Terraform or manually, so that special handling could be applied.

-> This data source can only be used with a workspace-level provider!

## Example Usage

Listing all catalogs:

```hcl
data "databricks_catalogs" "all" {}

output "all_catalogs" {
  value = data.databricks_catalogs.all
}
```

## Attribute Reference

This data source exports the following attributes:

* `ids` - set of [databricks_catalog](../resources/catalog.md) names

## Related Resources

The following resources are used in the same context:

* [databricks_schema](../resources/schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
