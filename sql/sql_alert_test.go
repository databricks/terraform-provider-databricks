package sql_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccSqlAlert(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_sql_query" "this" {
			data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
			name = "tf-{var.RANDOM}"
			query = "SELECT 1 AS p1, 2 as p2"
		}

        resource "databricks_permissions" "alert_usage" {
			sql_alert_id = databricks_sql_alert.alert.id
			access_control {
              group_name       = "users"
              permission_level = "CAN_RUN"
			}
		}

		resource "databricks_sql_alert" "alert" {
			query_id = databricks_sql_query.this.id
			name = "tf-alert-{var.RANDOM}"
			options {
				column = "p1"
				op = "=="
				value = "2"
				muted = false
			}
		}`,
	}, acceptance.Step{
		Template: `
		resource "databricks_sql_query" "this" {
			data_source_id = "{env.TEST_DEFAULT_WAREHOUSE_DATASOURCE_ID}"
			name = "tf-{var.RANDOM}"
			query = "SELECT 1 AS p1, 2 as p2"
		}

        resource "databricks_permissions" "alert_usage" {
			sql_alert_id = databricks_sql_alert.alert.id
			access_control {
              group_name       = "users"
              permission_level = "CAN_RUN"
			}
		}

		resource "databricks_sql_alert" "alert" {
			query_id = databricks_sql_query.this.id
			name = "tf-alert-{var.RANDOM}"
			options {
				column = "p1"
				op = ">="
				value = "3"
				muted = false
			}
		}`,
	})
}
