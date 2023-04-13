---
subcategory: "Unity Catalog"
---
# databricks_sql_table (Resource)

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, databases (also called schemas), and tables / views.

A `databricks_sql_table` is contained within [databricks_schema](schema.md).

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

resource "databricks_sql_table" "thing" {
  provider           = databricks.workspace
  name               = "quickstart_table"
  catalog_name       = databricks_catalog.sandbox.id
  schema_name        = databricks_schema.things.id
  table_type         = "MANAGED"
  data_source_format = "DELTA"
  storage_location   = ""

  column {
    name      = "id"
    type_text = "int"
  }
  column {
    name      = "name"
    type_text = "varchar(64)"
    comment   = "name of thing"
  }
  comment = "this table is managed by terraform"
}

resource "databricks_sql_table" "thing_view" {
  provider           = databricks.workspace
  name               = "quickstart_table_view"
  catalog_name       = databricks_catalog.sandbox.id
  schema_name        = databricks_schema.things.id
  table_type         = "VIEW"

  view_definition    = format("SELECT name FROM %s WHERE id == 1", databricks_sql_table.thing.id)

  comment = "this view is managed by terraform"
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of table relative to parent catalog and schema. Change forces creation of a new resource.
* `catalog_name` - Name of parent catalog
* `schema_name` - Name of parent Schema relative to parent Catalog
* `table_type` - Distinguishes a view vs. managed/external Table. `MANAGED`, `EXTERNAL` or `VIEW`. Change forces creation of a new resource.
* `storage_location` - URL of storage location for Table data (required for EXTERNAL Tables. For Managed Tables, if the path is provided it needs to be a Staging Table path that has been generated through the Staging Table API, otherwise should be empty)
* `data_source_format` - External tables are supported in multiple data source formats. The string constants identifying these formats are `DELTA`, `CSV`, `JSON`, `AVRO`, `PARQUET`, `ORC`, `TEXT`. Change forces creation of a new resource.
* `view_definition` - (Optional) SQL text defining the view (for `table_type == "VIEW"`)
* `storage_credential_name` - (Optional) For EXTERNAL Tables only: the name of storage credential to use. This cannot be updated
* `comment` - (Optional) User-supplied free-form text.
* `properties` - (Optional) Extensible Table properties.

### `column` configuration block
For table columns
Currently, changing the column definitions for a table will require dropping and re-creating the table

* `name` - User-visible name of column
* `type_text` - Column type spec (with metadata) as SQL text
* `comment` - (Optional) User-supplied free-form text.
* `nullable` - (Optional) Whether field is nullable (Default: `true`)

## Import

This resource can be imported by name:

```bash
$ terraform import databricks_sql_table.this <name>
```