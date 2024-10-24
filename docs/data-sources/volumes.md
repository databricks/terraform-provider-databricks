---
subcategory: "Unity Catalog"
---
# databricks_volumes Data Source

-> **Note** This data source can only be used with a workspace-level provider!

Retrieves a list of [databricks_volume](../resources/volume.md) ids (full names), that were created by Terraform or manually.

## Plugin Framework Migration
The volumes data source has been migrated from sdkv2 to plugin framework。 If you encounter any problem with this data source and suspect it is due to the migration, you can fallback to sdkv2 by setting the environment variable in the following way `export USE_SDK_V2_DATA_SOURCES="databricks_volumes"`.

## Example Usage

Listing all volumes in a _things_ [databricks_schema](../resources/schema.md) of a  _sandbox_ [databricks_catalog](../resources/catalog.md):

```hcl
data "databricks_volumes" "this" {
  catalog_name = "sandbox"
  schema_name  = "things"
}

output "all_volumes" {
  value = data.databricks_volumes.this
}
```

## Argument Reference

* `catalog_name` - (Required) Name of [databricks_catalog](../resources/catalog.md)
* `schema_name` - (Required) Name of [databricks_schema](../resources/schema.md)

## Attribute Reference

This data source exports the following attributes:

* `ids` - a list of [databricks_volume](../resources/volume.md) full names: *`catalog`.`schema`.`volume`*

## Related Resources

The following resources are used in the same context:

* [databricks_volume](../resources/volume.md) to manage volumes within Unity Catalog.
* [databricks_schema](../resources/schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
