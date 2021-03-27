package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestPreviewAccPipelineResource_CreatePipeline(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			locals {
				name = "pipeline-acceptance-{var.RANDOM}"
			}
			resource "databricks_pipeline" "this" {
				name = locals.name
				storage = "/test/${locals.name}"
				configuration = {
					key1 = "value1"
					key2 = "value2"
				}
				clusters {
					label = "default"
					num_workers = 2
					custom_tags = {
						cluster_type = "default"
					}
				}

				cluster {
					label = "maintenance"
					num_workers = 1
					custom_tags = {
						cluster_type = "maintenance
					}
				}

				library {
					maven {
						coordinates = "com.microsoft.azure:azure-eventhubs-spark_2.11:2.3.7"
					}
				}
				filters {
					include = ["com.databricks.include"]
					exclude = ["com.databricks.exclude"]
				}
				continuous = false
			}
			`,
		},
	})
}
