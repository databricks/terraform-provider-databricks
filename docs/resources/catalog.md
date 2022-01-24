---
subcategory: "Unity Catalog"
---
# databricks_catalog Resource

-> **Private Preview** This feature is in [Private Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Unity Catalog helps simplify security and governance of your data. It is contained within [databricks_metastore](metastore.md) and can have [databricks_schema](schema.md). By default, Databricks creates `default` schema for every new catalog, but Terraform plugin is removing this auto-created schema, so that resource destroy could be done in a clean way.

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

* `name` - Name of Catalog relative to parent metastore. Change forces new resource.
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Catalog properties.

## Import

This resource can be imported via name:

```bash
$ terraform import databricks_catalog.this <name>
```
