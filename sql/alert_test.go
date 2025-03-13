package sql_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccAlert(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_query" "this" {
			warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
			display_name = "tf-{var.RANDOM}"
			query_text   = "SELECT 1 AS p1, 2 as p2"
		}

		resource "databricks_alert" "alert" {
			query_id = databricks_query.this.id
			display_name = "tf-alert-{var.RANDOM}"
			condition {
			    op = "EQUAL"
			    operand {
			      column {
			        name = "p2"
			      }
			    }
			    threshold {
			      value {
			        double_value = 2
			      }
			    }
			}
		}
`,
	}, acceptance.Step{
		Template: `
		resource "databricks_query" "this" {
			warehouse_id = "{env.TEST_DEFAULT_WAREHOUSE_ID}"
			display_name = "tf-{var.RANDOM}"
			query_text   = "SELECT 1 AS p1, 2 as p2"
		}

		resource "databricks_alert" "alert" {
			query_id = databricks_query.this.id
			display_name = "tf-alert-{var.RANDOM}"
			condition {
			    op = "GREATER_THAN"
			    operand {
			      column {
			        name = "p2"
			      }
			    }
			    threshold {
			      value {
			        double_value = 3
			      }
			    }
			}
		}`,
	})
}
