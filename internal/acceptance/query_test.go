package acceptance

import (
	"testing"
)

func TestAccQuery(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		resource "databricks_query" "this" {
			warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
			display_name = "tf-{var.RANDOM}"
			query_text = "SELECT 1 AS p1, 2 as p2"
		}
`,
	}, Step{
		Template: `
		resource "databricks_query" "this" {
			warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
			display_name = "tf-{var.RANDOM}"
			query_text = "SELECT 1 AS p1, 2 as p2"
  			parameter {
    			name = "foo"
    			text_value {
      				value = "bar"
    			}
    			title = "foo"
  			}
		}
`,
	})
}
