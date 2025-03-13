package workspace_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDirectoryResource(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_directory" "this" {
			path = "/Shared/provider-test/dir_{var.RANDOM}"
		}
		data "databricks_directory" "users" {
			path = "/Users"
		}
		`,
	})
}
