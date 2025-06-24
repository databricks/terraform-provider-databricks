package secrets_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccSecretResource(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_secret_scope" "this" {
			name = "tf-scope-{var.RANDOM}"
		}
		resource "databricks_secret" "this" {
			scope = databricks_secret_scope.this.name
			string_value = "{var.RANDOM}"
			key = "password"
		}`,
	})
}
