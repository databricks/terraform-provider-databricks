---
subcategory: "Unity Catalog"
---
# databricks_catalog Data Source

Retrieves details of a specific catalog in Unity Catalog, that were created by Terraform or manually. Use [databricks_catalogs](catalogs.md) to retrieve IDs of multiple catalogs from Unity Catalog

-> This data source can only be used with a workspace-level provider!

## Example Usage

Read  on a specific catalog `test`:

```hcl
data "databricks_catalog" "test" {
  name = "test"
}
resource "databricks_grants" "things" {
  catalog = data.databricks_catalog.test.name
  grant {
    principal  = "sensitive"
    privileges = ["USE_CATALOG"]
  }
}
```

## Argument Reference

* `name` - (Required) name of the catalog

## Attribute Reference

This data source exports the following attributes:

* `id` - same as the `name`
* `catalog_info` - the [CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo) object for a Unity Catalog catalog. This contains the following attributes (see ):
  * `name` - Name of the catalog
  * `full_name` The full name of the catalog. Corresponds with the name field.
  * `catalog_type` - Type of the catalog, e.g. `MANAGED_CATALOG`, `DELTASHARING_CATALOG`, `SYSTEM_CATALOG`,
  * `owner` - Current owner of the catalog
  * `comment` - Free-form text description
  * `storage_location` -  Storage Location URL (full path) for managed tables within catalog.
  * `storage_root` - Storage root URL for managed tables within catalog.
  * `connection_name` - The name of the connection to an external data source.
  * `provider_name` - The name of delta sharing provider.
  * `share_name` -  The name of the share under the share provider.
  * `created_at` - Time at which this catalog was created, in epoch milliseconds.
  * `created_by` - Username of catalog creator.
  * `updated_at` - Time at which this catalog was last modified, in epoch milliseconds.
  * `updated_by` - Username of user who last modified catalog.
  * `effective_predictive_optimization_flag` - object describing applied predictive optimization flag.
  * `enable_predictive_optimization` - Whether predictive optimization should be enabled for this object and objects under it.
  * `isolation_mode` - Whether the current securable is accessible from all workspaces or a  specific set of workspaces.
  * `metastore_id` - Unique identifier of parent metastore.
  * `options` - A map of key-value properties attached to the securable.
  * `properties` - A map of key-value properties attached to the securable.
  * `securable_kind` - Kind of catalog securable.
  * `securable_type` - Securable type.

## Related Resources

The following resources are used in the same context:

* [databricks_grant](../resources/grant.md) to manage grants within Unity Catalog.
* [databricks_catalogs](catalogs.md) to list all catalogs within Unity Catalog metastore.
