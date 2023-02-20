package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccQuery(t *testing.T) {
	t.Parallel()
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
				resource "databricks_sql_query" "q1" {
					data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
					name = "tf-{var.RANDOM}"
					query = "SELECT {{ p1 }} AS p1, 2 as p2"

					parameter {
						name = "p1"
						title = "Title for p1"
						text {
							value = "default"
						}
					}

					tags = [
						"t1",
						"t2",
					]
				}

				resource "databricks_sql_visualization" "q1v1" {
					query_id = databricks_sql_query.q1.id
					type = "chart"
					name = "My Chart"
					description = "Some Description"

					options = jsonencode({
						columnMapping = {
							"p1" = "x"
							"p2" = "y"
						}
					})
				}
			`,
		},
	})
}
