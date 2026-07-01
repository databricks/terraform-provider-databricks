package pipelines_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
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
	acceptance.WorkspaceLevel(t, acceptance.Step{
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
		}
		` + dltNotebookResource,
	})
}

func pipelineRunAsTemplate(runAs string) string {
	if runAs != "" {
		runAs = `run_as { ` + runAs + ` }`
	}
	return `
	data "databricks_current_user" "me" {}

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
				cluster_type = "maintenance"
			}
		}

		continuous = false

		` + runAs + `
	}
	` + dltNotebookResource
}

func TestAccPipelineRunAsUser(t *testing.T) {
	if !acceptance.IsGcp(t) {
		acceptance.Skipf(t)("Only GCP service principals are treated as users")
	}

	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: pipelineRunAsTemplate(`user_name = data.databricks_current_user.me.user_name`),
	})
}

func TestUcAccPipelineRunAsServicePrincipal(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	spId := acceptance.GetEnvOrSkipTest(t, "ACCOUNT_LEVEL_SERVICE_PRINCIPAL_ID")
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: pipelineRunAsTemplate(`service_principal_name = "` + spId + `"`),
	})
}

func TestUcAccPipelineRunAsMutations(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	spId := acceptance.GetEnvOrSkipTest(t, "ACCOUNT_LEVEL_SERVICE_PRINCIPAL_ID")
	// Note: the attribute must match the type of principal that the test is run as.
	ctx := context.Background()
	attribute := acceptance.GetRunAsAttribute(t, ctx)
	acceptance.UnityWorkspaceLevel(
		t,
		// Create a pipeline without a run_as configured.
		acceptance.Step{
			Template: pipelineRunAsTemplate(""),
		},
		// Update pipeline to a service principal `run_as`
		acceptance.Step{
			Template: pipelineRunAsTemplate(`service_principal_name = "` + spId + `"`),
		},
		// Update pipeline to a user `run_as`
		acceptance.Step{
			Template: pipelineRunAsTemplate(attribute + ` = data.databricks_current_user.me.user_name`),
		},
		// Update pipeline back to a service principal `run_as`
		acceptance.Step{
			Template: pipelineRunAsTemplate(`service_principal_name = "` + spId + `"`),
		},
		// Remove run_as, and there should be an error.
		acceptance.Step{
			Template:    pipelineRunAsTemplate(""),
			ExpectError: regexp.MustCompile("Once configured, run_as of the pipeline can not be set to null. Please set run_as to " + spId + " to preserve the pipeline's run_as"),
		},
	)
}

func TestAccAwsPipelineResource_CreatePipeline(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
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
	}, acceptance.Step{
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
	acceptance.WorkspaceLevel(t, acceptance.Step{
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
				num_workers = 0
				spark_conf = {
					"spark.databricks.cluster.profile" = "singleNode"
				}
			}

			continuous = false
		}
		` + dltNotebookResource,
		Check: acceptance.ResourceCheck("databricks_pipeline.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
	acceptance.WorkspaceLevel(t, acceptance.Step{
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
		}
		` + dltNotebookResource,
		Check: acceptance.ResourceCheck("databricks_pipeline.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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
	}, acceptance.Step{
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
			expected_last_modified = ` + fmt.Sprintf("%d", lastModified) + `
		}
		` + dltNotebookResource,
		Check: acceptance.ResourceCheck("databricks_pipeline.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
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

// TestUcAccPipelineIngestionServerlessFalse verifies that an ingestion pipeline
// (one with an ingestion_definition) can be created on classic compute by
// setting serverless = false and providing a cluster block.
//
// The API defaults ingestion pipelines to serverless compute. If the provider
// drops an explicit serverless = false (which the SDK does by default, because
// the field is marshaled with omitempty), the API treats the pipeline as
// serverless and rejects the cluster block with "You cannot provide cluster
// settings when using serverless compute". This test guards against that
// regression: creation succeeds only if serverless = false actually reaches the
// API, and the Check asserts the created pipeline is not serverless.
//
// The connection points at a dummy host on purpose. The serverless/cluster
// validation runs at pipeline creation, before any source connectivity is
// checked, so the connection does not need to be reachable for this test.
func TestUcAccPipelineIngestionServerlessFalse(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		data "databricks_node_type" "smallest" {
			local_disk = true
		}

		resource "databricks_connection" "this" {
			name            = "conn-{var.STICKY_RANDOM}"
			connection_type = "MYSQL"
			options = {
				host     = "test.mysql.database.azure.com"
				port     = "3306"
				user     = "user"
				password = "password"
			}
		}

		resource "databricks_schema" "this" {
			name         = "sch_{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_pipeline" "this" {
			name       = "pipeline-acceptance-{var.STICKY_RANDOM}"
			serverless = false
			catalog    = "main"
			schema     = databricks_schema.this.name
			channel    = "PREVIEW"

			cluster {
				label        = "default"
				node_type_id = data.databricks_node_type.smallest.id
				num_workers  = 1
			}

			ingestion_definition {
				connection_name = databricks_connection.this.name
				objects {
					table {
						source_catalog      = "srccat"
						source_schema       = "srcsch"
						source_table        = "srctbl"
						destination_catalog = "main"
						destination_schema  = databricks_schema.this.name
					}
				}
			}
		}
		`,
		Check: acceptance.ResourceCheck("databricks_pipeline.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			ctx = context.WithValue(ctx, common.Api, common.API_2_1)
			w, err := client.WorkspaceClient()
			assert.NoError(t, err)
			pipeline, err := w.Pipelines.Get(ctx, pipelines.GetPipelineRequest{
				PipelineId: id,
			})
			assert.NoError(t, err)
			assert.False(t, pipeline.Spec.Serverless, "pipeline should be created on classic compute")
			return nil
		}),
		// Ingestion pipelines report a computed source_type back on the
		// ingestion_definition, which produces a non-empty plan after apply.
		// This is unrelated to the serverless behavior under test.
		ExpectNonEmptyPlan: true,
	})
}
