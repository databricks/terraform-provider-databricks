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

resource "databricks_volume" "testVolume1" {
  name             = "testVolume1"
  catalog_name     = databricks_catalog.testCatalog.name
  schema_name      = databricks_schema.testSchema.name
  volume_type      = "MANAGED"
  comment          = "Plugin Framework PoC"
}

resource "databricks_volume" "testVolume2" {
  name             = "testVolume2"
  catalog_name     = databricks_catalog.testCatalog.name
  schema_name      = databricks_schema.testSchema.name
  volume_type      = "MANAGED"
  comment          = "Plugin Framework PoC"
}

data "databricks_volumes_plugin_framework" "testVolumes" {
  catalog_name = databricks_catalog.testCatalog.name
  schema_name  = databricks_schema.testSchema.name
}

output "all_volumes" {
  value = data.databricks_volumes_plugin_framework.testVolumes
}


