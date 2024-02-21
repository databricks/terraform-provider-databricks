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

			resource "databricks_sql_table" "myInferenceTable" {
				catalog_name = databricks_catalog.sandbox.id
				schema_name = databricks_schema.things.name
				name = "bar${var.RANDOM}_inference"
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

			resource "databricks_lakehouse_monitor" "testMonitorInference" {
				table_name = databricks_sql_table.myInferenceTable.id
				assets_dir = "/Shared/provider-test/databricks_lakehouse_monitoring/${databricks_sql_table.myInferenceTable.name}"
				output_schema_name = databricks_schema.things.id
				inference_log  {
				  granularities = ["1 day"]
				  timestamp_col = "timestamp"
				  prediction_col = "prediction"
				  model_id_col = "model_id"
				  problem_type = "PROBLEM_TYPE_REGRESSION"
				} 
			}

			resource "databricks_sql_table" "myTimeseries" {
				catalog_name = databricks_catalog.sandbox.id
				schema_name = databricks_schema.things.name
				name = "bar${var.RANDOM}_timeseries"
				table_type = "MANAGED"
				data_source_format = "DELTA"

				column {
					name = "timestamp"
					type = "int"
				}
			}

			resource "databricks_lakehouse_monitor" "testMonitorTimeseries" {
				table_name = databricks_sql_table.myTimeseries.id
				assets_dir = "/Shared/provider-test/databricks_lakehouse_monitoring/${databricks_sql_table.myTimeseries.name}"
				output_schema_name = databricks_schema.things.id
				time_series  {
				  granularities = ["1 day"]
				  timestamp_col = "timestamp"
				} 
			}
		`,
	})
}

func TestUcAccUpdateLakehouseMonitor(t *testing.T) {
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

			resource "databricks_sql_table" "myInferenceTable" {
				catalog_name = databricks_catalog.sandbox.id
				schema_name = databricks_schema.things.name
				name = "bar${var.RANDOM}_inference"
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

			resource "databricks_lakehouse_monitor" "testMonitorInference" {
				table_name = databricks_sql_table.myInferenceTable.id
				assets_dir = "/Shared/provider-test/databricks_lakehouse_monitoring/${databricks_sql_table.myInferenceTable.name}"
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
	}, step{
		Template: `
		resource "databricks_lakehouse_monitor" "testMonitorInference" {
			table_name = databricks_sql_table.myInferenceTable.id
			assets_dir = "/Shared/provider-test/databricks_lakehouse_monitoring/${databricks_sql_table.myInferenceTable.name}"
			output_schema_name = databricks_schema.things.id
			inference_log  {
			  granularities = ["1 hour"]
			  timestamp_col = "timestamp"
			  prediction_col = "prediction"
			  model_id_col = "model_id"
			  problem_type = "PROBLEM_TYPE_REGRESSION"
			} 
		}
		`
	})
}