package acceptance

import (
	"testing"
)

func TestAccWorkspaceFile(t *testing.T) {
	workspaceLevel(t, step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}
