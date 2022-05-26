---
subcategory: "Unity Catalog"
---
# databricks_catalog Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

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

* `name` - Name of Catalog relative to parent metastore. Change forces creation of a new resource.
* `owner` - (Optional) Username/groupname of catalog owner.
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Catalog properties.

## Import

This resource can be imported by name:

```bash
$ terraform import databricks_catalog.this <name>
```

## Related Resources

The following resources are used in the same context:

* [databricks_table](../data-sources/tables.md) data to list tables within Unity Catalog.
* [databricks_schema](../data-sources/schemas.md) data to list schemas within Unity Catalog.
* [databricks_catalog](../data-sources/catalogs.md) data to list catalogs within Unity Catalog.
