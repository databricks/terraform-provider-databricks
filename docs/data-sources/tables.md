---
subcategory: "Unity Catalog"
---
# databricks_tables Data Source

Retrieves a list of managed or external table full names in Unity Catalog, that were created by Terraform or manually. Use [databricks_views](views.md) for retrieving a list of views.

-> This data source can only be used with a workspace-level provider!

## Example Usage

Granting `SELECT` and `MODIFY` to `sensitive` group on all tables a _things_ [databricks_schema](../resources/schema.md) from _sandbox_ [databricks_catalog](../resources/catalog.md):

```hcl
data "databricks_tables" "things" {
  catalog_name = "sandbox"
  schema_name  = "things"
}

resource "databricks_grants" "things" {
  for_each = data.databricks_tables.things.ids

  table = each.value

  grant {
    principal  = "sensitive"
    privileges = ["SELECT", "MODIFY"]
  }
}
```

## Argument Reference

* `catalog_name` - (Required) Name of [databricks_catalog](../resources/catalog.md)
* `schema_name` - (Required) Name of [databricks_schema](../resources/schema.md)

## Attribute Reference

This data source exports the following attributes:

* `ids` - set of databricks_table full names: *`catalog`.`schema`.`table`*

## Related Resources

The following resources are used in the same context:

* [databricks_schema](../resources/schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
