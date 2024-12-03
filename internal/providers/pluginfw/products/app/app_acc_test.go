package app_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

/*
	resource "databricks_sql_endpoint" "this" {
		name = "tf-{var.STICKY_RANDOM}"
		cluster_size = "2X-Small"
		max_num_clusters = 1

		tags {
			custom_tags {
				key   = "Owner"
				value = "eng-dev-ecosystem-team_at_databricks.com"
			}
		}
	}

	resource "databricks_job" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}

	resource "databricks_model_serving" "this" {
		name = "tf-{var.STICKY_RANDOM}"
		config {
			served_models {
				name = "prod_model"
				model_name = "experiment-fixture-model"
				model_version = "1"
				workload_size = "Small"
				scale_to_zero_enabled = true
			}
		}
	}
*/
const baseResources = `

	resource "databricks_secret_scope" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}

	resource "databricks_secret" "this" {
	    scope = databricks_secret_scope.this.name
		key = "tf-{var.STICKY_RANDOM}"
		string_value = "secret"
	}

`

/*

		resources {
			name = "warehouse"
			description = "warehouse for app"
			job {
				id = databricks_job.this.id
				permission = "CAN_MANAGE"
			}
		}
		resources {
			name = "serving endpoint"
			description = "serving endpoint for app"
			serving_endpoint {
				name = databricks_model_serving.this.name
				permission = "CAN_MANAGE"
			}
		}	
		resources {
			name = "sql warehouse"
			description = "sql warehouse for app"
			sql_warehouse {
				id = databricks_sql_endpoint.this.id
				permission = "CAN_MANAGE"
			}
		}
*/
func makeTemplate(description string) string {
	appTemplate := baseResources + `
	resource "databricks_app" "this" {
		name = "{var.STICKY_RANDOM}"
		description = "%s"
		resources {
			name = "secret"
			description = "secret for app"
			secret {
				scope = databricks_secret_scope.this.name
				key = databricks_secret.this.key
				permission = "MANAGE"
			}
		}
	}`
	return fmt.Sprintf(appTemplate, description)
}

var templateWithInvalidResource = `
	resource "databricks_app" "this" {
		name = "{var.STICKY_RANDOM}"
		description = "My app"
		resources {
			name = "invalid resource"
			description = "invalid resource for app"
			secret {
				permission = "CAN_MANAGE"
				key = "test"
				scope = "test"
			}
			sql_warehouse {
				id = "123"
				permission = "CAN_MANAGE"
			}
		}
	}`

func TestAccApp_InvalidResource(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: templateWithInvalidResource,
		ExpectError: regexp.MustCompile(regexp.QuoteMeta(`2 attributes specified when one (and only one) of
[resources[0].job.<.secret,resources[0].job.<.serving_endpoint,resources[0].job.<.sql_warehouse]
is required`)),
	})
}

func TestAccApp(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: makeTemplate("My app"),
	}, acceptance.Step{
		Template: makeTemplate("My new app"),
	})
}
