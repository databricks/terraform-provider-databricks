terraform {
    required_providers {
      databricks = {
        source = "databricks/databricks"
      }
    }
}


resource "databricks_catalog" "testcatalog" {
  name    = "testcatalog-9aug"
  comment = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_schema" "testschema" {
  catalog_name = databricks_catalog.testcatalog.name
  name         = "testschema-9aug"
  comment      = "Plugin Framework PoC"
  properties = {
    purpose = "testing"
  }
}

resource "databricks_volume" "testvolume1" {
  name             = "testvolume1-9aug"
  catalog_name     = databricks_catalog.testcatalog.name
  schema_name      = databricks_schema.testschema.name
  volume_type      = "MANAGED"
  comment          = "Plugin Framework PoC"
}

resource "databricks_volume" "testvolume2" {
  name             = "testvolume2-9aug"
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


