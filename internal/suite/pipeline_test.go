package acceptance

import (
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

var (
	dltNotebookResource = `
	resource "databricks_notebook" "this" {
		content_base64 = base64encode(<<-EOT
			CREATE LIVE TABLE clickstream_raw AS 
			SELECT * FROM json.` + "`/databricks-datasets/wikipedia-datasets/data-001/clickstream/raw-uncompressed-json/2015_2_clickstream.json`" + `
			
			-- COMMAND ----------
			
			CREATE LIVE TABLE clickstream_clean(
			  CONSTRAINT valid_current_page EXPECT (current_page_id IS NOT NULL and current_page_title IS NOT NULL),
			  CONSTRAINT valid_count EXPECT (click_count > 0) ON VIOLATION FAIL UPDATE
			) TBLPROPERTIES ("quality" = "silver")
			AS SELECT
			  CAST (curr_id AS INT) AS current_page_id,
			  curr_title AS current_page_title,
			  CAST(n AS INT) AS click_count,
			  CAST (prev_id AS INT) AS previous_page_id,
			  prev_title AS previous_page_title
			FROM live.clickstream_raw
			
			-- COMMAND ----------
			
			CREATE LIVE TABLE top_spark_referers TBLPROPERTIES ("quality" = "gold")
			AS SELECT
			  previous_page_title as referrer,
			  click_count
			FROM live.clickstream_clean
			WHERE current_page_title = 'Apache_Spark'
			ORDER BY click_count DESC
			LIMIT 10					
		  EOT
		)
		path = "/Shared/${local.name}"
		language = "SQL"
	}
`
)

func TestAccPipelineResource_CreatePipeline(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			locals {
				name = "pipeline-acceptance-{var.RANDOM}"
			}
			resource "databricks_pipeline" "this" {
				name = local.name
				storage = "/test/${local.name}"
				
				configuration = {
					key1 = "value1"
					key2 = "value2"
				}

				library {
					notebook {
						path = databricks_notebook.this.path
					}
				}

				cluster {
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					label = "default"
					num_workers = 2
					custom_tags = {
						cluster_type = "default"
					}
				}

				cluster {
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					label = "maintenance"
					num_workers = 1
					custom_tags = {
						cluster_type = "maintenance"
					}
				}

				continuous = false
			}
			` + dltNotebookResource,
		},
	})
}

func TestAccAwsPipelineResource_CreatePipeline(t *testing.T) {
	if cloud, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloud != "aws" {
		t.Skip("Test is only for CLOUD_ENV=aws")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			locals {
				name = "pipeline-acceptance-aws-{var.RANDOM}"
			}
			resource "databricks_pipeline" "this" {
				name = local.name
				storage = "/test/${local.name}"		
				configuration = {
					key1 = "value1"
					key2 = "value2"
				}
				library {
					notebook {
						path = databricks_notebook.this.path
					}
				}

				cluster {
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					label = "default"
					num_workers = 2
					custom_tags = {
						cluster_type = "default"
					}
					aws_attributes {
						first_on_demand = 1
					}
				}
				cluster {
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					label = "maintenance"
					num_workers = 1
					custom_tags = {
						cluster_type = "maintenance"
					}
				}
				continuous = false
			}
			` + dltNotebookResource,
		},
		{
			Template: `
			locals {
				name = "pipeline-acceptance-aws-{var.RANDOM}"
			}
			resource "databricks_pipeline" "this" {
				name = local.name
				storage = "/test/${local.name}"		
				configuration = {
					key1 = "value1"
					key2 = "value2"
				}
				library {
					notebook {
						path = databricks_notebook.this.path
					}
				}

				cluster {
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					label = "default"
					num_workers = 3
					custom_tags = {
						cluster_type = "default"
					}
					aws_attributes {
						first_on_demand = 2
					}
				}
				cluster {
					instance_pool_id = "{env.TEST_INSTANCE_POOL_ID}"
					label = "maintenance"
					num_workers = 1
					custom_tags = {
						cluster_type = "maintenance"
					}
				}
				continuous = false
			}
			` + dltNotebookResource,
		},
	})
}
