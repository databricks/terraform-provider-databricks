package acceptance

import (
	"testing"
)

func TestAccWorkspaceFile(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, Step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}

func TestAccWorkspaceFileEmptyFile(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_workspace_file" "empty" {
			source = "{var.CWD}/../../workspace/acceptance/testdata/empty_file"
			path = "/Shared/provider-test/empty_{var.RANDOM}"
		}`,
	})
}

func TestAccWorkspaceFileBase64(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: `resource "databricks_workspace_file" "this2" {
			content_base64 = "YWJjCg=="
			path = "/Shared/provider-test/xx2_{var.RANDOM}"
		}`,
	}, Step{
		Template: `resource "databricks_workspace_file" "this2" {
			content_base64 = "YWJjCg=="
			path = "/Shared/provider-test/xx2_{var.RANDOM}_renamed"
		}`,
	})
}
