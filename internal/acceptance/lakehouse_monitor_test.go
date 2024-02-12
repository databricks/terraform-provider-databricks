package acceptance

import (
	"testing"
)

func TestUcAccLakehouseMonitor(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
			resource "databricks_catalog" "sandbox" {
				name         = "sandbox${var.RANDOM}"
				comment      = "this catalog is managed by terraform"
				properties = {
					purpose = "testing"
				}
				force_destroy = true
			}

			resource "databricks_schema" "things" {
				catalog_name = databricks_catalog.sandbox.id
				name         = "things${var.RANDOM}"
				comment      = "this database is managed by terraform"
				properties = {
					kind = "various"
				}
			}

			resource "databricks_sql_table" "mytable" {
				catalog_name = databricks_catalog.sandbox.id
				schema_name = databricks_schema.things.name
				name = "bar${var.RANDOM}"
				table_type = "MANAGED"
				data_source_format = "DELTA"
				
				column {
					name = "model_id"
					type = "int"
				}
				column {
					name = "timestamp"
					type = "int"
				}
				column {
					name = "prediction"
					type = "int"
				}
			}

			resource "databricks_lakehouse_monitor" "testMonitor" {
				table_name = databricks_sql_table.mytable.id
				assets_dir = "/Shared/provider-test/databricks_lakehouse_monitoring/${databricks_sql_table.mytable.name}"
				output_schema_name = databricks_schema.things.id
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
