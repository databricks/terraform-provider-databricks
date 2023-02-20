package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDashboard(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
				resource "databricks_sql_dashboard" "d1" {
					name = "tf-{var.RANDOM}-dashboard"

					// The SQLA API doesn't save tags on create, only on update.
					// Uncomment the following when this is fixed.
					// tags = [
					// "tf-{var.RANDOM}-dashboard",
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
					data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
					name = "tf-{var.RANDOM}-query"
					query = "SELECT 1"
				}

				resource "databricks_sql_visualization" "q1v1" {
					query_id = databricks_sql_query.q1.id
					type = "table"
					name = "My Table"

					options = jsonencode({})
				}

				resource "databricks_sql_visualization" "q1v2" {
					query_id = databricks_sql_query.q1.id
					type = "table"
					name = "My Table (1)"

					options = jsonencode({})

					# Note: this resource differs from the one above in that
					# the query plan is set. This tests that it can either
					# be unset or set and in both cases yield a consistent result.
					query_plan = jsonencode({
						# The value should be non-empty to check we don't have
						# plan changes due to whitespace changes in JSON serialization.
						groups = [
						]
					})
				}
			`,
		},
	})
}
