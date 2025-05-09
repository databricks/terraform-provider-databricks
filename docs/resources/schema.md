---
subcategory: "Unity Catalog"
---
# databricks_schema Resource

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, Databases (also called Schemas), and Tables / Views.

-> This resource can only be used with a workspace-level provider!

A `databricks_schema` is contained within [databricks_catalog](catalog.md) and can contain tables & views.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  name    = "sandbox"
  comment = "this catalog is managed by terraform"
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
* `catalog_name` - Name of parent catalog. Change forces creation of a new resource.
* `storage_root` - (Optional) Managed location of the schema. Location in cloud storage where data for managed tables will be stored. If not specified, the location will default to the catalog root location. Change forces creation of a new resource.
* `owner` - (Optional) Username/groupname/sp application_id of the schema owner.
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Schema properties.
* `enable_predictive_optimization` - (Optional) Whether predictive optimization should be enabled for this object and objects under it. Can be `ENABLE`, `DISABLE` or `INHERIT`
* `force_destroy` - (Optional) Delete schema regardless of its contents.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this schema in form of `<catalog_name>.<name>`.

## Import

This resource can be imported by its full name:

```hcl
import {
  to = databricks_schema.this
  id = "<catalog_name>.<name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_schema.this "<catalog_name>.<name>"
```

## Related Resources

The following resources are used in the same context:

* [databricks_tables](../data-sources/tables.md) data to list tables within Unity Catalog.
* [databricks_schemas](../data-sources/schemas.md) data to list schemas within Unity Catalog.
* [databricks_catalogs](../data-sources/catalogs.md) data to list catalogs within Unity Catalog.
