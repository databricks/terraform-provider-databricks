---
subcategory: "Unity Catalog"
---
# databricks_external_locations Data Source

-> **Note** This data source could be only used with workspace-level provider!

Retrieves a list of [databricks_external_location](./external_location.md) objects, that were created by Terraform or manually, so that special handling could be applied.

## Example Usage

List all external locations in the metastore

```hcl
data "databricks_external_locations" "all" {}

output "all_external_locations" {
  value = data.databricks_external_locations.all.names
}
```

## Attribute Reference

This data source exports the following attributes:

* `names` - List of names of [databricks_external_location](./external_location.md) in the metastore

## Related Resources

The following resources are used in the same context:

* [databricks_external_location](./external_location.md) to get information about a single external location
* [databricks_external_location](../resources/external_location.md) to manage external locations within Unity Catalog.
