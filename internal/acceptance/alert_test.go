package acceptance

import (
	"testing"
)

func TestAccAlert(t *testing.T) {
	WorkspaceLevel(t, Step{
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
	}, Step{
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
