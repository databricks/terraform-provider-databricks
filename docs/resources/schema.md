---
subcategory: "Unity Catalog"
---
# databricks_schema Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, Databases (also called Schemas), and Tables / Views.

A `databricks_schema` is contained within [databricks_catalog](catalog.md) and can contain tables & views.

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

resource "databricks_schema" "things" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of Schema relative to parent catalog. Change forces creation of a new resource.
* `catalog_name` - Name of parent catalog
* `owner` - (Optional) Username/groupname/sp application_id of the schema owner.
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Schema properties.

## Import

This resource can be imported by name:

```bash
$ terraform import databricks_schema.this <name>
```

## Related Resources

The following resources are used in the same context:

* [databricks_table](../data-sources/tables.md) data to list tables within Unity Catalog.
* [databricks_schema](../data-sources/schemas.md) data to list schemas within Unity Catalog.
* [databricks_catalog](../data-sources/catalogs.md) data to list catalogs within Unity Catalog.
