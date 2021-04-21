package acceptance

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
)

func TestPreviewAccDashboard(t *testing.T) {
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
				resource "databricks_sql_dashboard" "d1" {
					name = "tf-{var.RANDOM}"

					// The SQLA API doesn't save tags on create, only on update.
					// Uncomment the following when this is fixed.
					// tags = [
					// "tf-{var.RANDOM}",
					// ]
				}

				resource "databricks_sql_widget" "d1w1" {
					dashboard_id = databricks_sql_dashboard.d1.id
					text = "hello there!"

					position {
						size_x = 3
						size_y = 4
						pos_x = 0
						pos_y = 0
					}
				}

				resource "databricks_sql_widget" "d1w2" {
					dashboard_id = databricks_sql_dashboard.d1.id
					visualization_id = databricks_sql_visualization.q1v1.visualization_id

					position {
						size_x = 3
						size_y = 4
						pos_x = 3
						pos_y = 0
					}
				}

				resource "databricks_sql_query" "q1" {
					data_source_id = databricks_sql_endpoint.this.data_source_id
					name = "tf-{var.RANDOM}"
					query = "SELECT 1"
				}

				resource "databricks_sql_visualization" "q1v1" {
					query_id = databricks_sql_query.q1.id
					type = "table"
					name = "My Table"

					options = jsonencode({})
				}

				resource "databricks_sql_endpoint" "this" {
					name = "tf-{var.RANDOM}"
					cluster_size = "Small"
					max_num_clusters = 1
				}
			`,
		},
	})
}
