package acceptance

import (
	"testing"
)

func TestUcAccLakehouseMonitor(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
			resource "databricks_catalog" "sandbox" {
				name         = "sandbox{var.RANDOM}"
				comment      = "this catalog is managed by terraform"
				properties = {
					purpose = "testing"
				}
			}

			resource "databricks_schema" "things" {
				catalog_name = databricks_catalog.sandbox.id
				name         = "things{var.RANDOM}"
				comment      = "this database is managed by terraform"
				properties = {
					kind = "various"
				}
			}

			resource "databricks_table" "mytable" {
				catalog_name = databricks_catalog.sandbox.id
				schema_name = databricks_schema.things.name
				name = "bar"
				table_type = "MANAGED"
				data_source_format = "DELTA"
				
				column {
					name      = "model_id"
					position  = 0
					type_name = "INT"
					type_text = "int"
					type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
				}
				column {
					name      = "timestamp"
					position  = 1
					type_name = "INT"
					type_text = "int"
					type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
				}
				column {
					name      = "prediction"
					position  = 2
					type_name = "INT"
					type_text = "int"
					type_json = "{\"name\":\"id\",\"type\":\"integer\",\"nullable\":true,\"metadata\":{}}"
				}
			}

			resource "databricks_lakehouse_monitor" "testMonitor" {
				table_name = "{databricks_catalog.sandbox.id}.{databricks_schema.things.name}.{databricks_table.myTestTable.name}"
				assets_dir = "/Shared/provider-test/databricks_lakehouse_monitoring/{databricks_table.myTestTable.name}"
				output_schema_name = "{databricks_catalog.sandbox.id}.{databricks_schema.things.name}"
				inference_log  {
				  granularities = ["1 day"]
				  timestamp_col = "timestamp"
				  prediction_col = "prediction"
				  model_id_col = "model_id"
				  problem_type = "PROBLEM_TYPE_REGRESSION"
				} 
			}
		`,
	})
}
