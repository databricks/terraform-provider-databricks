terraform {
    required_providers {
      databricks = {
        source = "databricks/databricks"
      }
    }
}

provider "databricks" {
    profile = "aws-prod-ucws"
}

resource "databricks_catalog" "testcatalog" {
  name    = "testcatalog"
  comment = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_schema" "testschema" {
  catalog_name = databricks_catalog.testcatalog.name
  name         = "testSchema"
  comment      = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_sql_table" "testtable" {
  catalog_name       = databricks_catalog.testcatalog.name
  schema_name        = databricks_schema.testschema.name
  name               = "testTable"
  table_type         = "MANAGED"
  data_source_format = "DELTA"

  column {
    name     = "timestamp"
    type     = "int"
  }
}

resource "databricks_lakehouse_monitor_pluginframework" "testmonitor" { 
  table_name = "${databricks_catalog.testcatalog.name}.${databricks_schema.testschema.name}.${databricks_sql_table.testtable.name}"
}

