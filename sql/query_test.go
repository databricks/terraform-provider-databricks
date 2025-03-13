package sql_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccQuery(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_query" "this" {
			warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
			display_name = "tf-{var.RANDOM}"
			query_text = "SELECT 1 AS p1, 2 as p2"
		}
`,
	}, acceptance.Step{
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
