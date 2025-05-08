---
subcategory: "Unity Catalog"
---
# databricks_sql_table (Resource)

Within a metastore, Unity Catalog provides a 3-level namespace for organizing data: Catalogs, databases (also called schemas), and tables/views.

A `databricks_sql_table` is contained within [databricks_schema](schema.md), and can represent either a managed table, an external table, or a view.

This resource creates and updates the Unity Catalog table/view by executing the necessary SQL queries on a special auto-terminating cluster it would create for this operation. You could also specify a SQL warehouse or cluster for the queries to be executed on.

-> This resource can only be used with a workspace-level provider!

~> This resource doesn't handle complex cases of schema evolution due to the limitations of Terraform itself.  If you need to implement schema evolution it's recommended to use specialized tools, such as, [Liquibase](https://medium.com/dbsql-sme-engineering/advanced-schema-management-on-databricks-with-liquibase-1900e9f7b9c0) and [Flyway](https://medium.com/dbsql-sme-engineering/databricks-schema-management-with-flyway-527c4a9f5d67).

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

## Use an Identity Column

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
  name               = "identity_table"
  catalog_name       = databricks_catalog.sandbox.name
  schema_name        = databricks_schema.things.name
  table_type         = "MANAGED"
  column {
    name     = "id"
    type     = "bigint"
    identity = "default"
  }
  column {
    name    = "name"
    type    = "string"
    comment = "name of thing"
  }
  comment = "this table is managed by terraform"
}
```

## Enable automatic clustering

```hcl
resource "databricks_sql_table" "thing" {
  provider           = databricks.workspace
  name               = "auto_cluster_table"
  catalog_name       = databricks_catalog.sandbox.name
  schema_name        = databricks_schema.things.name
  table_type         = "MANAGED"
  cluster_keys       = ["AUTO"]
  
  column {
    name    = "name"
    type    = "string"
    comment = "name of thing"
  }
  comment = "this table is managed by terraform"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of table relative to parent catalog and schema. Change forces the creation of a new resource.
* `catalog_name` - Name of parent catalog. Change forces the creation of a new resource.
* `schema_name` - Name of parent Schema relative to parent Catalog. Change forces the creation of a new resource.
* `table_type` - Distinguishes a view vs. managed/external Table. `MANAGED`, `EXTERNAL`, or `VIEW`. Change forces the creation of a new resource.
* `storage_location` - (Optional) URL of storage location for Table data (required for EXTERNAL Tables). Not supported for `VIEW` or `MANAGED` table_type.
* `data_source_format` - (Optional) External tables are supported in multiple data source formats. The string constants identifying these formats are `DELTA`, `CSV`, `JSON`, `AVRO`, `PARQUET`, `ORC`, and `TEXT`. Change forces the creation of a new resource. Not supported for `MANAGED` tables or `VIEW`.
* `view_definition` - (Optional) SQL text defining the view (for `table_type == "VIEW"`). Not supported for `MANAGED` or `EXTERNAL` table_type.
* `cluster_id` - (Optional) All table CRUD operations must be executed on a running cluster or SQL warehouse. If a cluster_id is specified, it will be used to execute SQL commands to manage this table. If empty, a cluster will be created automatically with the name `terraform-sql-table`. Conflicts with `warehouse_id`.
* `warehouse_id` - (Optional) All table CRUD operations must be executed on a running cluster or SQL warehouse. If a `warehouse_id` is specified, that SQL warehouse will be used to execute SQL commands to manage this table. Conflicts with `cluster_id`.
* `cluster_keys` - (Optional) a subset of columns to liquid cluster the table by. For automatic clustering, set `cluster_keys` to `["AUTO"]`. To turn off clustering, set it to `["NONE"]`. Conflicts with `partitions`.
* `partitions` - (Optional) a subset of columns to partition the table by. Change forces the creation of a new resource. Conflicts with `cluster_keys`.
* `storage_credential_name` - (Optional) For EXTERNAL Tables only: the name of storage credential to use. Change forces the creation of a new resource.
* `owner` - (Optional) User name/group name/sp application_id of the table owner.
* `comment` - (Optional) User-supplied free-form text. Changing the comment is not currently supported on the `VIEW` table type.
* `options` - (Optional) Map of user defined table options. Change forces creation of a new resource.
* `properties` - (Optional) A map of table properties.

### `column` configuration block

For table columns
Currently, changing the column definitions for a table will require dropping and re-creating the table

* `name` - User-visible name of column
* `type` - Column type spec (with metadata) as SQL text. Not supported for `VIEW` table_type.
* `identity` - (Optional) Whether the field is an identity column. Can be `default`, `always`, or unset. It is unset by default.
* `comment` - (Optional) User-supplied free-form text.
* `nullable` - (Optional) Whether field is nullable (Default: `true`)

## Attribute Reference

In addition to all the arguments above, the following attributes are exported:

* `id` - ID of this table in the form of `<catalog_name>.<schema_name>.<name>`.

## Import

This resource can be imported by its full name.

```bash
terraform import databricks_sql_table.this <catalog_name>.<schema_name>.<name>
```

## Migration from `databricks_table`

The `databricks_table` resource has been deprecated in favor of `databricks_sql_table`. To migrate from `databricks_table` to `databricks_sql_table`:

1. Define a `databricks_sql_table` resource with arguments corresponding to `databricks_table`.
2. Add a `removed` block to remove the `databricks_table` resource without deleting the existing table by using the `lifecycle` block. If you're using Terraform version below v1.7.0, you will need to use the `terraform state rm` command instead.
3. Add an `import` block to add the `databricks_sql_table` resource, corresponding to the existing table. If you're using Terraform version below v1.5.0, you will need to use `terraform import` command instead.

For example, suppose we have the following `databricks_table` resource:

```hcl
resource "databricks_table" "this" {
  catalog_name       = "catalog"
  schema_name        = "schema"
  name               = "table"
  table_type         = "MANAGED"
  data_source_format = "DELTA"
  column {
    name      = "col1"
    type_name = "STRING"
    type_json = "{\"type\":\"STRING\"}"
    comment   = "comment"
    nullable  = true
  }
  comment = "comment"
  properties = {
    key = "value"
  }
}
```

The migration would look like this:

```hcl
# Leave this resource definition as-is.
resource "databricks_table" "this" { ... }

# Remove the old resource without destroying the existing table.
removed {
  from = databricks_table.this.id
  lifecycle {
    destroy = false
  }
}

# Import the existing table as a databricks_sql_table.
import {
  to = databricks_sql_table.this
  id = "<catalog_name>.<schema_name>.<name>"
}

# Define the new databricks_sql_table resource.
resource "databricks_sql_table" "this" {
  catalog_name = "catalog"
  schema_name = "schema"
  name = "table"
  table_type = "MANAGED"
  data_source_format = "DELTA"
  column {
    name = "col1"
    type = "STRING"   # <-- changed from type_name
    type_json = "{\"type\":\"STRING\"}"
    comment = "comment"
    nullable = true
  }
  comment = "comment"
  properties = {
    key = "value"
  }
}
```

Finally, run `terraform plan` to verify the changes, followed by `terraform apply` to apply the changes.
