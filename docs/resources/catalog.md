---
subcategory: "Unity Catalog"
---
# databricks_catalog Resource

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, Databases (also called Schemas), and Tables / Views.

A `databricks_catalog` is contained within [databricks_metastore](metastore.md) and can contain [databricks_schema](schema.md). By default, Databricks creates `default` schema for every new catalog, but Terraform plugin is removing this auto-created schema, so that resource destruction could be done in a clean way.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  metastore_id = databricks_metastore.this.id
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of Catalog relative to parent metastore.
* `storage_root` - (Optional) Managed location of the catalog. Location in cloud storage where data for managed tables will be stored. If not specified, the location will default to the metastore root location. Change forces creation of a new resource.
* `provider_name` - (Optional) For Delta Sharing Catalogs: the name of the delta sharing provider. Change forces creation of a new resource.
* `share_name` - (Optional) For Delta Sharing Catalogs: the name of the share under the share provider. Change forces creation of a new resource.
* `connection_name` - (Optional) For Foreign Catalogs: the name of the connection to an external data source. Changes forces creation of a new resource.
* `owner` - (Optional) Username/groupname/sp application_id of the catalog owner.
* `isolation_mode` - (Optional) Whether the catalog is accessible from all workspaces or a specific set of workspaces. Can be `ISOLATED` or `OPEN`. Setting the catalog to `ISOLATED` will automatically allow access from the current workspace.
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Catalog properties.
* `options` - (Optional) For Foreign Catalogs: the name of the entity from an external data source that maps to a catalog. For example, the database name in a PostgreSQL server.
* `force_destroy` - (Optional) Delete catalog regardless of its contents.

## Import

This resource can be imported by name:

```bash
terraform import databricks_catalog.this <name>
```

## Related Resources

The following resources are used in the same context:

* [databricks_table](../data-sources/tables.md) data to list tables within Unity Catalog.
* [databricks_schema](../data-sources/schemas.md) data to list schemas within Unity Catalog.
* [databricks_catalog](../data-sources/catalogs.md) data to list catalogs within Unity Catalog.
