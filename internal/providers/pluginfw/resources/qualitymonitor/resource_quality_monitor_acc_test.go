package qualitymonitor_test

import (
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

var commonPartQualityMonitoring = `resource "databricks_catalog" "sandbox" {
	name         = "sandbox{var.STICKY_RANDOM}"
	comment      = "this catalog is managed by terraform"
	properties = {
		purpose = "testing"
	}
	force_destroy = true
}

resource "databricks_schema" "things" {
	catalog_name = databricks_catalog.sandbox.id
	name         = "things{var.STICKY_RANDOM}"
	comment      = "this database is managed by terraform"
	properties = {
		kind = "various"
	}
}

resource "databricks_sql_table" "myInferenceTable" {
	catalog_name = databricks_catalog.sandbox.id
	schema_name = databricks_schema.things.name
	name = "bar{var.STICKY_RANDOM}_inference"
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

`

func TestUcAccQualityMonitor(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		t.Skipf("databricks_quality_monitor resource is not available on GCP")
	}
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: commonPartQualityMonitoring + `

			resource "databricks_quality_monitor_pluginframework" "testMonitorInference" {
				table_name = databricks_sql_table.myInferenceTable.id
				assets_dir = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myInferenceTable.name}"
				output_schema_name = databricks_schema.things.id
				inference_log = {
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
				name = "bar{var.STICKY_RANDOM}_timeseries"
				table_type = "MANAGED"
				data_source_format = "DELTA"

				column {
					name = "timestamp"
					type = "int"
				}
			}

			resource "databricks_quality_monitor_pluginframework" "testMonitorTimeseries" {
				table_name = databricks_sql_table.myTimeseries.id
				assets_dir = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myTimeseries.name}"
				output_schema_name = databricks_schema.things.id
				time_series = {
				  granularities = ["1 day"]
				  timestamp_col = "timestamp"
				} 
			}

			resource "databricks_sql_table" "mySnapshot" {
				catalog_name = databricks_catalog.sandbox.id
				schema_name = databricks_schema.things.name
				name = "bar{var.STICKY_RANDOM}_snapshot"
				table_type = "MANAGED"
				data_source_format = "DELTA"

				column {
					name = "timestamp"
					type = "int"
				}
			}

			resource "databricks_quality_monitor_pluginframework" "testMonitorSnapshot" {
				table_name = databricks_sql_table.mySnapshot.id
				assets_dir = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myTimeseries.name}"
				output_schema_name = databricks_schema.things.id
				snapshot = {
				} 
			}
		`,
	})
}

func TestUcAccUpdateQualityMonitor(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		t.Skipf("databricks_quality_monitor resource is not available on GCP")
	}
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: commonPartQualityMonitoring + `
			resource "databricks_quality_monitor_pluginframework" "testMonitorInference" {
				table_name = databricks_sql_table.myInferenceTable.id
				assets_dir = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myInferenceTable.name}"
				output_schema_name = databricks_schema.things.id
				inference_log = {
				  granularities = ["1 day"]
				  timestamp_col = "timestamp"
				  prediction_col = "prediction"
				  model_id_col = "model_id"
				  problem_type = "PROBLEM_TYPE_REGRESSION"
				}
			}
		`,
	}, acceptance.Step{
		Template: commonPartQualityMonitoring + `
		resource "databricks_quality_monitor_pluginframework" "testMonitorInference" {
			table_name = databricks_sql_table.myInferenceTable.id
			assets_dir = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myInferenceTable.name}"
			output_schema_name = databricks_schema.things.id
			inference_log = {
				granularities = ["1 hour"]
				timestamp_col = "timestamp"
				prediction_col = "prediction"
				model_id_col = "model_id"
				problem_type = "PROBLEM_TYPE_REGRESSION"
			}
		}
		`,
	})
}

func TestUcAccQualityMonitorImportPluginFramework(t *testing.T) {
	if os.Getenv("GOOGLE_CREDENTIALS") != "" {
		t.Skipf("databricks_quality_monitor resource is not available on GCP")
	}
	acceptance.UnityWorkspaceLevel(t,
		acceptance.Step{
			Template: commonPartQualityMonitoring + `

			resource "databricks_quality_monitor_pluginframework" "testMonitorInference" {
				table_name = databricks_sql_table.myInferenceTable.id
				assets_dir = "/Shared/provider-test/databricks_quality_monitoring/${databricks_sql_table.myInferenceTable.name}"
				output_schema_name = databricks_schema.things.id
				inference_log = {
				  granularities = ["1 day"]
				  timestamp_col = "timestamp"
				  prediction_col = "prediction"
				  model_id_col = "model_id"
				  problem_type = "PROBLEM_TYPE_REGRESSION"
				} 
			}
		`,
		},
		acceptance.Step{
			ImportState:                          true,
			ResourceName:                         "databricks_quality_monitor_pluginframework.testMonitorInference",
			ImportStateIdFunc:                    acceptance.BuildImportStateIdFunc("databricks_quality_monitor_pluginframework.testMonitorInference", "table_name"),
			ImportStateVerify:                    true,
			ImportStateVerifyIdentifierAttribute: "table_name",
		},
	)
}
