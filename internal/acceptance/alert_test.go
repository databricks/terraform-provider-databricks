package acceptance

import (
	"testing"
)

func TestAccAlert(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		resource "databricks_sql_query" "this" {
			data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
			name = "tf-{var.RANDOM}"
			query = "SELECT 1 AS p1, 2 as p2"
		}

        resource "databricks_permissions" "alert_usage" {
			sql_alert_id = databricks_alert.alert.id
			access_control {
              group_name       = "users"
              permission_level = "CAN_RUN"
			}
		}

		resource "databricks_alert" "alert" {
			query_id = databricks_sql_query.this.id
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
		resource "databricks_sql_query" "this" {
			data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
			name = "tf-{var.RANDOM}"
			query = "SELECT 1 AS p1, 2 as p2"
		}

        resource "databricks_permissions" "alert_usage" {
			sql_alert_id = databricks_alert.alert.id
			access_control {
              group_name       = "users"
              permission_level = "CAN_RUN"
			}
		}

		resource "databricks_alert" "alert" {
			query_id = databricks_sql_query.this.id
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
