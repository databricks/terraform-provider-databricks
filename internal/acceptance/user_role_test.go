package acceptance

import (
	"testing"
)

func TestAccUserRole(t *testing.T) {
	workspaceLevel(t, LegacyStep{
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
