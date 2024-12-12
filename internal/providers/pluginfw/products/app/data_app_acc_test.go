package app_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

const fastApp = `
	resource "databricks_secret_scope" "this" {
		name = "tf-{var.STICKY_RANDOM}"
	}

	resource "databricks_secret" "this" {
	    scope = databricks_secret_scope.this.name
		key = "tf-{var.STICKY_RANDOM}"
		string_value = "secret"
	}

	resource "databricks_app" "this" {
		name = "{var.STICKY_RANDOM}"
		description = "%s"
		resources = [{
			name = "secret"
			description = "secret for app"
			secret = {
				scope = databricks_secret_scope.this.name
				key = databricks_secret.this.key
				permission = "MANAGE"
			}
		}]
	}`

func TestAccAppDataSource(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsGcp(t) {
		acceptance.Skipf(t)("not available on GCP")
	}
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: fastApp + `
		data "databricks_app" "this" {
			name = databricks_app.this.name
		}
		`,
	})
}
