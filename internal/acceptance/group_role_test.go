package acceptance

import (
	"testing"
)

func TestAccGroupRole(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `
		resource "databricks_group" "this" {
			display_name = "tf-{var.RANDOM}"
		}
		resource "databricks_group_role" "this" {
			group_id = databricks_group.this.id
			role = "arn:aws:iam::999999999999:role/foo"
		}`,
	})
}
