package acceptance

import (
	"testing"
)

func TestAccAlert(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_sql_query" "q1" {
			data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
			name = "tf-{var.RANDOM}"
			query = "SELECT 1 AS p1, 2 as p2"
		}

		resource "databricks_sql_alert" "alert" {
			query_id = databricks_sql_query.q1.id
			name = "tf-alert-{var.RANDOM}"
			options {
				column = "p1"
				op = "=="
				value = "2"
				muted = false
			}
		}`,
	})
}
