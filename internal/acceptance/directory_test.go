package acceptance

import (
	"testing"
)

func TestAccDirectoryResource(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_directory" "this" {
			path = "/Shared/provider-test/dir_{var.RANDOM}"
		}
		data "databricks_directory" "users" {
			path = "/Users"
		}
		`,
	})
}
