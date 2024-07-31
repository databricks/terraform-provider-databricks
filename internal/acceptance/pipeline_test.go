package acceptance

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/stretchr/testify/assert"
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
	workspaceLevel(t, step{
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
	})
}

func TestAccAwsPipelineResource_CreatePipeline(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		locals {
			name = "pipeline-acceptance-aws-{var.STICKY_RANDOM}"
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
	}, step{
		Template: `
		locals {
			name = "pipeline-acceptance-aws-{var.STICKY_RANDOM}"
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
	})
}

func TestAccPipelineResource_CreatePipelineWithoutWorkers(t *testing.T) {
	workspaceLevel(t, step{
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
				num_workers = 0
				spark_conf = {
					"spark.databricks.cluster.profile" = "singleNode"
				}
			}

			continuous = false
		}
		` + dltNotebookResource,
		Check: resourceCheck("databricks_pipeline.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			ctx = context.WithValue(ctx, common.Api, common.API_2_1)
			w, err := client.WorkspaceClient()
			assert.NoError(t, err)
			pipeline, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
				PipelineId: id,
			})
			assert.NoError(t, err)
			cluster := pipeline.Spec.Clusters[0]
			assert.Nil(t, cluster.Autoscale)
			assert.Equal(t, 0, cluster.NumWorkers)
			// Check that the zero was indeed send by the server, and it's not a default value
			assert.Contains(t, cluster.ForceSendFields, "NumWorkers")
			return nil
		},
		),
	})
}

func TestAccPipelineResourcLastModified(t *testing.T) {
	var lastModified int64
	workspaceLevel(t, step{
		Template: `
		locals {
			name = "pipeline-acceptance-{var.STICKY_RANDOM}"
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
			development = false
		}
		` + dltNotebookResource,
		Check: resourceCheck("databricks_pipeline.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			ctx = context.WithValue(ctx, common.Api, common.API_2_1)
			w, err := client.WorkspaceClient()
			assert.NoError(t, err)
			pipeline, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
				PipelineId: id,
			})
			assert.NoError(t, err)
			assert.Equal(t, pipeline.CreatorUserName, pipeline.RunAsUserName)
			lastModified = pipeline.LastModified
			return nil
		}),
	}, step{
		Template: `
		locals {
			name = "pipeline-acceptance-{var.STICKY_RANDOM}"
		}
		resource "databricks_pipeline" "this" {
			name = local.name
			storage = "/test/${local.name}"

			configuration = {
				key1 = "value1"
				key2 = "value2"
				key3 = "value3"
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
			development = false
			expected_last_modified = ` + fmt.Sprintf("%d", lastModified) + `
		}
		` + dltNotebookResource,
		Check: resourceCheck("databricks_pipeline.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			ctx = context.WithValue(ctx, common.Api, common.API_2_1)
			w, err := client.WorkspaceClient()
			assert.NoError(t, err)
			pipeline, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
				PipelineId: id,
			})
			assert.NoError(t, err)
			assert.NotEqual(t, pipeline.LastModified, lastModified)
			return nil
		}),
	})

}
