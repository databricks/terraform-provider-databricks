package aws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccUserRole(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_user" "this" {
			user_name = "{var.RANDOM}@example.com"
		}
		resource "databricks_user_role" "this" {
			user_id = databricks_user.this.id
			role = "arn:aws:iam::999999999999:role/foo"
		}`,
	})
}
