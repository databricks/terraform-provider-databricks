package tokens_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccTokenResource(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_token" "this" {
			lifetime_seconds = 6000
			comment = "Testing token"
		}`,
	})
}

func TestAccTokenResource_WithScopes(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_token" "this" {
			lifetime_seconds = 6000
			comment = "Testing scoped token"
			scopes = ["sql", "jobs"]
		}`,
	})
}
