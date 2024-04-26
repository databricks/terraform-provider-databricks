# TESTING SOP
# -----------
# Please make sure development overrides are in effect before running this manually
# https://github.com/databricks/terraform-provider-databricks/blob/main/CONTRIBUTING.md#developing-provider
# 1. touch ~/.terraformrc
# 2. Add the following to the file (update to your user.name in the path):
#     provider_installation {
#       dev_overrides {
#         "databricks/databricks" = "/Users/<user.name>/terraform-provider-databricks"
#       }
#       direct {}
#     }
# 3. run $ make in terraform-provider-databricks root directory to build the binary
# 4. cd terraform-provider-databricks/pluginframework/test
# 5. terraform init -upgrade
# 6. terraform apply

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

resource "databricks_lakehouse_monitor_plugin_framework" "testMonitor" { 
  table_name = "${databricks_catalog.testCatalog.name}.${databricks_schema.testSchema.name}.${databricks_sql_table.testTable.name}"
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