package pipelines_test

import (
	"fmt"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"testing"
)

var (
	dltNotebook = `
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

func TestAccDataSourcePipelines(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		locals {
			name = "pipeline-ds-acceptance-{var.RANDOM}"
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
		` + dltNotebook + `
		data "databricks_pipelines" "this" {
			pipeline_name = local.name
			depends_on = [databricks_pipeline.this]
		}`,
		Check: func(s *terraform.State) error {
			r, ok := s.RootModule().Resources["data.databricks_pipelines.this"]
			if !ok {
				return fmt.Errorf("data not found in state")
			}
			ids := r.Primary.Attributes["ids.#"]
			if ids == "0" {
				return fmt.Errorf("ids is empty: %v", r.Primary.Attributes)
			}
			return nil
		},
	})
}
