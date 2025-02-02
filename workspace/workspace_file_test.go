package workspace_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccWorkspaceFile(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}

func TestAccWorkspaceFileEmptyFile(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "empty" {
			source = "{var.CWD}/../../workspace/acceptance/testdata/empty_file"
			path = "/Shared/provider-test/empty_{var.RANDOM}"
		}`,
	})
}

func TestAccWorkspaceFileZipFile(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "zipfile" {
			source = "{var.CWD}/../../workspace/acceptance/testdata/zipfile.zip"
			path = "/Shared/provider-test/zipfile_{var.RANDOM}.zip"
		}`,
	})
}

func TestAccWorkspaceFileBase64(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "this2" {
			content_base64 = "YWJjCg=="
			path = "/Shared/provider-test/xx2_{var.RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_workspace_file" "this2" {
			content_base64 = "YWJjCg=="
			path = "/Shared/provider-test/xx2_{var.RANDOM}_renamed"
		}`,
	})
}
