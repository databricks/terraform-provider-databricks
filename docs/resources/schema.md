---
subcategory: "Unity Catalog"
---
# databricks_schema Resource

-> **Private Preview** This feature is in [Private Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, databases (also called schemas), and tables / views.

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

* `name` - Name of Schema relative to parent catalog. Change forces new resource.
* `catalog_name` - Name of parent catalog
* `owner` - (Optional) Username/groupname of schema owner. Currently this field can only be changed after the resource is created.
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Schema properties.

## Import

This resource can be imported via name:

```bash
$ terraform import databricks_schema.this <name>
```
