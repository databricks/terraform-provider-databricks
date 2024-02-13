---
subcategory: "Unity Catalog"
---
# databricks_storage_credentials Data Source

Retrieves a list of [databricks_storage_credential](./storage_credential.md) objects, that were created by Terraform or manually, so that special handling could be applied.

## Example Usage

List all storage credentials in the metastore

```hcl
data "databricks_storage_credentials" "all" {}

output "all_metastores" {
  value = data.databricks_metastores.all.names
}
```

## Attribute Reference

This data source exports the following attributes:

* `namé` - List of names of [databricks_storage_credential](./storage_credential.md) in the metastore

## Related Resources

The following resources are used in the same context:

* [databricks_storage_credential](./storage_credential.md) to get information about a single credential
* [databricks_storage_credential](../resources/storage_credential.md) to manage Storage Credentials within Unity Catalog.
