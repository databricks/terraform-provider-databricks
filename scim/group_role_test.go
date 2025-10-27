package scim_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccGroupRole(t *testing.T) {
	if !acceptance.IsAws(t) {
		acceptance.Skipf(t)("TestAccGroupRole is failing on non-AWS environments, likely due to read-after-write inconsistency.")
	}
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
