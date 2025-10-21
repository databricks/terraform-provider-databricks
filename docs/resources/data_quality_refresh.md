---
subcategory: "Data Quality Monitoring"
---
# databricks_data_quality_refresh Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

This resource allows you to refresh the data quality monitoring checks on Unity Catalog tables.

You must either:
1. be an owner of the table's parent catalog, have **USE_SCHEMA** on the table's parent schema, and have **SELECT** access on the table
2. have **USE_CATALOG** on the table's parent catalog, be an owner of the table's parent schema, and have **SELECT** access on the table.
3. have the following permissions:
   - **USE_CATALOG** on the table's parent catalog
   - **USE_SCHEMA** on the table's parent schema
   - be an owner of the table.

-> **Note** This resource can only be used with a workspace-level provider!

## Example Usage
```hcl
resource "databricks_catalog" "sandbox" {
  name    = "sandbox"
  comment = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
}
resource "databricks_schema" "myTestSchema" {
  catalog_name = databricks_catalog.sandbox.id
  name         = "myTestSchema"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}
resource "databricks_sql_table" "myTestTable" {
  catalog_name       = "main"
  schema_name        = databricks_schema.myTestSchema.name
  name               = "bar"
  table_type         = "MANAGED"
  data_source_format = "DELTA"
  column {
    name = "timestamp"
    type = "int"
  }
}
resource "databricks_data_quality_monitor" "this" {
  object_type = "table"
  object_id = databricks_sql_table.myTestTable.id
  data_profiling_config = {
    output_schema = databricks_schema.myTestSchema.schema_id
  }
}
resource "databricks_data_quality_refresh" "this" {
  object_type = "table"
  object_id = databricks_sql_table.myTestTable.id
}
```

## Arguments
The following arguments are supported:
* `object_id` (string, required) - The UUID of the request object. For example, table id
* `object_type` (string, required) - The type of the monitored object. Can be one of the following: `schema`or `table`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `end_time_ms` (integer) - Time when the refresh ended (milliseconds since 1/1/1970 UTC)
* `message` (string) - An optional message to give insight into the current state of the refresh (e.g. FAILURE messages)
* `refresh_id` (integer) - Unique id of the refresh operation
* `start_time_ms` (integer) - Time when the refresh started (milliseconds since 1/1/1970 UTC)
* `state` (string) - The current state of the refresh. Possible values are: `MONITOR_REFRESH_STATE_CANCELED`, `MONITOR_REFRESH_STATE_FAILED`, `MONITOR_REFRESH_STATE_PENDING`, `MONITOR_REFRESH_STATE_RUNNING`, `MONITOR_REFRESH_STATE_SUCCESS`, `MONITOR_REFRESH_STATE_UNKNOWN`
* `trigger` (string) - What triggered the refresh. Possible values are: `MONITOR_REFRESH_TRIGGER_DATA_CHANGE`, `MONITOR_REFRESH_TRIGGER_MANUAL`, `MONITOR_REFRESH_TRIGGER_SCHEDULE`, `MONITOR_REFRESH_TRIGGER_UNKNOWN`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "object_type,object_id,refresh_id"
  to = databricks_data_quality_refresh.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_data_quality_refresh.this "object_type,object_id,refresh_id"
```