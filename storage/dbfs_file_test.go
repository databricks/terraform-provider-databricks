package storage_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestAccDatabricksDBFSFile_CreateViaContent(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_dbfs_file" "this" {
			content_base64 = base64encode("{var.RANDOM}")
			path = "/tmp/tf-test/{var.RANDOM}.bin"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_dbfs_file" "this" {
			content_base64 = base64encode("{var.RANDOM}-changed")
			path = "/tmp/tf-test/{var.RANDOM}.bin"
		}`,
	})
}

func TestAccDatabricksDBFSFile_CreateViaSource(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_dbfs_file" "file_1" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/tmp/tf-test/file-source-{var.RANDOM}"
		}`,
	})
}
