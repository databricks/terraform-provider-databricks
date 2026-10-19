package exporter

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImportPipelineEmitsIngestionDependencies(t *testing.T) {
	ic := importContextForTest()
	services, _ := ic.allServicesAndListing()
	ic.enableServices(services)
	d := schema.TestResourceDataRaw(t, ic.Resources["databricks_pipeline"].Schema, map[string]any{
		"name":             "ingestion-pipeline",
		"budget_policy_id": "budget-policy-id",
		"gateway_definition": []any{map[string]any{
			"connection_name":         "gateway-connection",
			"gateway_storage_catalog": "main",
			"gateway_storage_schema":  "gateway_storage",
		}},
		"ingestion_definition": []any{map[string]any{
			"connection_name":      "source-connection",
			"ingestion_gateway_id": "gateway-pipeline-id",
			"netsuite_jar_path":    "dbfs:/jars/netsuite.jar",
			"objects": []any{map[string]any{
				"schema": []any{map[string]any{
					"destination_catalog": "main",
					"destination_schema":  "bronze",
					"fanout_options": []any{map[string]any{
						"transforms": []any{
							map[string]any{"format": "AVRO", "avro_options": []any{map[string]any{
								"schema_file_path": "/Volumes/main/default/schemas/events.avsc",
								"schema_registry":  []any{map[string]any{"connection_name": "schema-registry"}},
							}}},
							map[string]any{"format": "PROTOBUF", "protobuf_options": []any{map[string]any{
								"desc_file_path": "/Volumes/main/default/schemas/events.desc",
							}}},
						},
					}},
				}},
			}},
		}},
	})
	d.SetId("ingestion-pipeline-id")

	err := importPipeline(ic, &resource{Resource: "databricks_pipeline", ID: d.Id(), Name: "ingestion_pipeline", Data: d})
	require.NoError(t, err)
	for _, expected := range []string{
		"databricks_budget_policy[<unknown>] (id: budget-policy-id)",
		"databricks_connection[<unknown>] (id: gateway-connection)",
		"databricks_connection[<unknown>] (id: source-connection)",
		"databricks_connection[<unknown>] (id: schema-registry)",
		"databricks_pipeline[<unknown>] (id: gateway-pipeline-id)",
		"databricks_catalog[<unknown>] (id: main)",
		"databricks_schema[<unknown>] (id: main.gateway_storage)",
		"databricks_schema[<unknown>] (id: main.bronze)",
		"databricks_dbfs_file[<unknown>] (id: dbfs:/jars/netsuite.jar)",
		"databricks_file[<unknown>] (id: /Volumes/main/default/schemas/events.avsc)",
		"databricks_file[<unknown>] (id: /Volumes/main/default/schemas/events.desc)",
	} {
		assert.Contains(t, ic.testEmits, expected)
	}
}
