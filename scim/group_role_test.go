package scim_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccGroupRole(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
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
