---
subcategory: "Unity Catalog"
---
# databricks_storage_credentials Data Source

-> **Note** This data source can only be used with a workspace-level provider!

Retrieves a list of [databricks_storage_credential](./storage_credential.md) objects, that were created by Terraform or manually, so that special handling could be applied.

## Example Usage

List all storage credentials in the metastore

```hcl
data "databricks_storage_credentials" "all" {}

output "all_storage_credentials" {
  value = data.databricks_storage_credentials.all.names
}
```

## Attribute Reference

This data source exports the following attributes:

* `names` - List of names of [databricks_storage_credential](./storage_credential.md) in the metastore

## Related Resources

The following resources are used in the same context:

* [databricks_storage_credential](./storage_credential.md) to get information about a single credential
* [databricks_storage_credential](../resources/storage_credential.md) to manage Storage Credentials within Unity Catalog.
