---
subcategory: "Unity Catalog"
---
# databricks_metastores Data Source

Retrieves a mapping of name to id of [databricks_metastore](../resources/metastore.md) objects, that were created by Terraform or manually, so that special handling could be applied.

-> This data source can only be used with an account-level provider!

-> [`account_id`](../index.md#account_id) provider configuration property is required for this resource to work. Data resource will error in case of metastores with duplicate names. This data source is only available for users & service principals with account admin status

## Example Usage

Mapping of name to id of all metastores:

```hcl
data "databricks_metastores" "all" {}

output "all_metastores" {
  value = data.databricks_metastores.all.ids
}
```

## Attribute Reference

This data source exports the following attributes:

* `ids` - Mapping of name to id of [databricks_metastore](../resources/metastore.md)

## Related Resources

The following resources are used in the same context:

* [databricks_metastore](./metastore.md) to get information about a single metastore.
* [databricks_metastore](../resources/metastore.md) to manage Metastores within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
