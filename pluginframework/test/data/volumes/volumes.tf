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
  name    = "testcatalog-25july"
  comment = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_schema" "testschema" {
  catalog_name = databricks_catalog.testcatalog.name
  name         = "testschema-25july"
  comment      = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_volume" "testvolume1" {
  name             = "testvolume1-25july"
  catalog_name     = databricks_catalog.testcatalog.name
  schema_name      = databricks_schema.testschema.name
  volume_type      = "MANAGED"
  comment          = "Plugin Framework PoC"
}

resource "databricks_volume" "testvolume2" {
  name             = "testvolume2-25july"
  catalog_name     = databricks_catalog.testcatalog.name
  schema_name      = databricks_schema.testschema.name
  volume_type      = "MANAGED"
  comment          = "Plugin Framework PoC"
}

data "databricks_volumes_pluginframework" "testvolumes" {
  catalog_name = databricks_catalog.testcatalog.name
  schema_name  = databricks_schema.testschema.name
}

output "all_volumes" {
  value = data.databricks_volumes_pluginframework.testvolumes
}


