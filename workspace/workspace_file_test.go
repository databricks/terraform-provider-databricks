package workspace_test

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccWorkspaceFile(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}"
		}`,
	}, acceptance.Step{
		Template: `resource "databricks_workspace_file" "this" {
			source = "{var.CWD}/../storage/testdata/tf-test-python.py"
			path = "/Shared/provider-test/xx_{var.RANDOM}_renamed"
		}`,
	})
}

func TestAccWorkspaceFileEmptyFile(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "empty" {
			source = "{var.CWD}/acceptance/testdata/empty_file"
			path = "/Shared/provider-test/empty_{var.RANDOM}"
		}`,
	})
}

func TestAccWorkspaceFileZipFile(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "zipfile" {
			source = "{var.CWD}/acceptance/testdata/zipfile.zip"
			path = "/Shared/provider-test/zipfile_{var.RANDOM}.zip"
		}`,
	})
}

func TestAccWorkspaceFileCreate_NonExistentParent(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: `resource "databricks_workspace_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Shared/provider-test/workspace-file-parent-{var.RANDOM}/test.py"
		}`,
		Check: acceptance.ResourceCheck("databricks_workspace_file.this",
			func(ctx context.Context, client *common.DatabricksClient, id string) error {
				w, err := client.WorkspaceClient()
				if err != nil {
					return err
				}
				// Verify the file exists and has the right content type.
				info, err := w.Workspace.GetStatusByPath(ctx, id)
				require.NoError(t, err)
				assert.Equal(t, id, info.Path)
				// Verify the parent folder was created automatically.
				parent := id[:len(id)-len("/test.py")]
				_, err = w.Workspace.GetStatusByPath(ctx, parent)
				assert.NoError(t, err, "parent folder should have been created automatically")
				return nil
			}),
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
