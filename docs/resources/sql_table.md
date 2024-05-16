---
subcategory: "Unity Catalog"
---
# databricks_sql_table (Resource)

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, databases (also called schemas), and tables / views.

A `databricks_sql_table` is contained within [databricks_schema](schema.md), and can represent either a managed table, an external table or a view.

This resource creates and updates the Unity Catalog table/view by executing the necessary SQL queries on a special auto-terminating cluster it would create for this operation. You could also specify a SQL warehouse or cluster for the queries to be executed on.

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

resource "databricks_sql_table" "thing" {
  provider           = databricks.workspace
  name               = "quickstart_table"
  catalog_name       = databricks_catalog.sandbox.name
  schema_name        = databricks_schema.things.name
  table_type         = "MANAGED"
  data_source_format = "DELTA"
  storage_location   = ""

  column {
    name = "id"
    type = "int"
  }
  column {
    name    = "name"
    type    = "string"
    comment = "name of thing"
  }
  comment = "this table is managed by terraform"
}

resource "databricks_sql_table" "thing_view" {
  provider     = databricks.workspace
  name         = "quickstart_table_view"
  catalog_name = databricks_catalog.sandbox.name
  schema_name  = databricks_schema.things.name
  table_type   = "VIEW"
  cluster_id   = "0423-201305-xsrt82qn"

  view_definition = format("SELECT name FROM %s WHERE id == 1", databricks_sql_table.thing.id)

  comment = "this view is managed by terraform"
}
```

### Use an existing warehouse to create a table

```hcl
resource "databricks_sql_endpoint" "this" {
  name             = "endpoint"
  cluster_size     = "2X-Small"
  max_num_clusters = 1
}

resource "databricks_sql_table" "thing" {
  provider           = databricks.workspace
  name               = "quickstart_table"
  catalog_name       = databricks_catalog.sandbox.name
  schema_name        = databricks_schema.things.name
  table_type         = "MANAGED"
  data_source_format = "DELTA"
  storage_location   = ""
  warehouse_id       = databricks_sql_endpoint.this.id

  column {
    name = "id"
    type = "int"
  }
  column {
    name    = "name"
    type    = "string"
    comment = "name of thing"
  }
  comment = "this table is managed by terraform"
}

resource "databricks_sql_table" "thing_view" {
  provider     = databricks.workspace
  name         = "quickstart_table_view"
  catalog_name = databricks_catalog.sandbox.name
  schema_name  = databricks_schema.things.name
  table_type   = "VIEW"
  warehouse_id = databricks_sql_endpoint.this.id

  view_definition = format("SELECT name FROM %s WHERE id == 1", databricks_sql_table.thing.id)

  comment = "this view is managed by terraform"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of table relative to parent catalog and schema. Change forces creation of a new resource.
* `catalog_name` - Name of parent catalog. Change forces creation of a new resource.
* `schema_name` - Name of parent Schema relative to parent Catalog. Change forces creation of a new resource.
* `table_type` - Distinguishes a view vs. managed/external Table. `MANAGED`, `EXTERNAL` or `VIEW`. Change forces creation of a new resource.
* `storage_location` - (Optional) URL of storage location for Table data (required for EXTERNAL Tables). Not supported for `VIEW` or `MANAGED` table_type.
* `data_source_format` - (Optional) External tables are supported in multiple data source formats. The string constants identifying these formats are `DELTA`, `CSV`, `JSON`, `AVRO`, `PARQUET`, `ORC`, `TEXT`. Change forces creation of a new resource. Not supported for `MANAGED` tables or `VIEW`.
* `view_definition` - (Optional) SQL text defining the view (for `table_type == "VIEW"`). Not supported for `MANAGED` or `EXTERNAL` table_type.
* `cluster_id` - (Optional) All table CRUD operations must be executed on a running cluster or SQL warehouse. If a cluster_id is specified, it will be used to execute SQL commands to manage this table. If empty, a cluster will be created automatically with the name `terraform-sql-table`.
* `warehouse_id` - (Optional) All table CRUD operations must be executed on a running cluster or SQL warehouse. If a `warehouse_id` is specified, that SQL warehouse will be used to execute SQL commands to manage this table. Conflicts with `cluster_id`.
* `cluster_keys` - (Optional) a subset of columns to liquid cluster the table by. Conflicts with `partitions`.
* `storage_credential_name` - (Optional) For EXTERNAL Tables only: the name of storage credential to use. Change forces creation of a new resource.
* `owner` - (Optional) Username/groupname/sp application_id of the schema owner.
* `comment` - (Optional) User-supplied free-form text. Changing comment is not currently supported on `VIEW` table_type.
* `options` - (Optional) Map of user defined table options. Change forces creation of a new resource.
* `properties` - (Optional) Map of table properties.
* `partitions` - (Optional) a subset of columns to partition the table by. Change forces creation of a new resource. Conflicts with `cluster_keys`.

### `column` configuration block

For table columns
Currently, changing the column definitions for a table will require dropping and re-creating the table

* `name` - User-visible name of column
* `type` - (Optional) Column type spec (with metadata) as SQL text. Not supported for `VIEW` table_type.
* `comment` - (Optional) User-supplied free-form text.
* `nullable` - (Optional) Whether field is nullable (Default: `true`)

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this table in form of `<catalog_name>.<schema_name>.<name>`.

## Import

This resource can be imported by its full name:

```bash
terraform import databricks_sql_table.this <catalog_name>.<schema_name>.<name>
```
