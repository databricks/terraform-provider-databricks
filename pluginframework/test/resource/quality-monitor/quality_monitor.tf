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

resource "databricks_catalog" "testCatalog" {
  name    = "testCatalog"
  comment = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_schema" "testSchema" {
  catalog_name = databricks_catalog.testCatalog.name
  name         = "testSchema"
  comment      = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_sql_table" "testTable" {
  catalog_name       = databricks_catalog.testCatalog.name
  schema_name        = databricks_schema.testSchema.name
  name               = "testTable"
  table_type         = "MANAGED"
  data_source_format = "DELTA"

  column {
    name     = "timestamp"
    type     = "int"
  }
}

resource "databricks_lakehouse_monitor_pluginframework" "testMonitor" { 
  table_name = "${databricks_catalog.testCatalog.name}.${databricks_schema.testSchema.name}.${databricks_sql_table.testTable.name}"
}

