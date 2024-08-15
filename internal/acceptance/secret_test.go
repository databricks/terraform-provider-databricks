package acceptance

import (
	"testing"
)

func TestAccSecretResource(t *testing.T) {
	workspaceLevel(t, LegacyStep{
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
