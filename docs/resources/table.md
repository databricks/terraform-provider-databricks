---
subcategory: "Unity Catalog"
---
# databricks_table (Resource)

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, databases (also called schemas), and tables / views.

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

A `databricks_table` is contained within [databricks_schema](schema.md).

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

resource "databricks_table" "thing" {
  provider           = databricks.workspace
  name               = "quickstart_table"
  catalog_name       = databricks_catalog.sandbox.id
  schema_name        = databricks_schema.things.name
  table_type         = "MANAGED"
  data_source_format = "DELTA"
  column {
    name      = "id"
    position  = 0
    type_name = "INT"
    type_text = "int"
    type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
  }
  column {
    name      = "name"
    position  = 1
    type_name = "STRING"
    type_text = "varchar(64)"
    type_json = "{\"name\":\"name\",\"type\":\"varchar(64)\",\"nullable\":true,\"metadata\":{}}"
  }
  comment = "this table is managed by terraform"
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of table relative to parent catalog and schema. Change forces creation of a new resource.
* `catalog_name` - Name of parent catalog
* `schema_name` - Name of parent Schema relative to parent Catalog
* `table_type` - Distinguishes a view vs. managed/external Table. `MANAGED`, `EXTERNAL` or `VIEW`
* `storage_location` - URL of storage location for Table data (required for EXTERNAL Tables. For Managed Tables, if the path is provided it needs to be a Staging Table path that has been generated through the Staging Table API, otherwise should be empty)
* `data_source_format` - External tables are supported in multiple data source formats. The string constants identifying these formats are `DELTA`, `CSV`, `JSON`, `AVRO`, `PARQUET`, `ORC`, `TEXT`
* `view_definition` - (Optional) SQL text defining the view (for `table_type == "VIEW"`)
* `storage_credential_name` - (Optional) For EXTERNAL Tables only: the name of storage credential to use. This cannot be updated
* `owner` - (Optional) Username/groupname/sp application_id of the table owner.
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Table properties.

### `column` configuration block

-> **Note** `column` block would be simplified in future versions of provider

For table columns
* `name` - User-visible name of column
* `type_name` - Name of (outer) type
* `type_text` - Column type spec (with metadata) as SQL text
* `type_json` - Column type spec (with metadata) as JSON string
* `position` - Ordinal position of column, starting at 0.
* `type_precision` - (Optional) Digits of precision; applies to `DECIMAL` columns
* `type_scale` - (Optional) Digits to right of decimal; applies to `DECIMAL` columns 
* `type_interval_type` - (Optional) Format of `INTERVAL` columns
* `comment` - (Optional) User-supplied free-form text.
* `nullable` - (Optional) Whether field is nullable (Default: `true`)
* `partition_index` - (Optional) Partition ID

## Import

This resource can be imported by full name: *`catalog`.`schema`.`table`*:

```bash
$ terraform import databricks_table.this <full-name>
```

## Related Resources

The following resources are used in the same context:

* [databricks_table](../data-sources/tables.md) data to list tables within Unity Catalog.
* [databricks_schema](../data-sources/schemas.md) data to list schemas within Unity Catalog.
* [databricks_catalog](../data-sources/catalogs.md) data to list catalogs within Unity Catalog.
