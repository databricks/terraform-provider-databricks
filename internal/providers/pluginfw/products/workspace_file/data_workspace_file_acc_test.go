package workspace_file_test

import (
	"strconv"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const workspaceFileTemplate = `
resource "databricks_directory" "this" {
	path = "/Shared/provider-test/workspace_file_{var.RANDOM}"
}

resource "databricks_workspace_file" "file_a" {
	path           = "${databricks_directory.this.path}/file_a.txt"
	content_base64 = base64encode("Hello World A")
}

resource "databricks_workspace_file" "file_b" {
	path           = "${databricks_directory.this.path}/file_b.txt"
	content_base64 = base64encode("Hello World B")
}

resource "databricks_directory" "subdir" {
	path = "${databricks_directory.this.path}/subdir"
}

resource "databricks_directory" "empty_subdir" {
	path = "${databricks_directory.this.path}/empty_subdir"
}

resource "databricks_workspace_file" "file_c" {
	path           = "${databricks_directory.subdir.path}/file_c.txt"
	content_base64 = base64encode("Hello World C")
}
`

func checkWorkspaceFilesPopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_workspace_file.this"]
		require.True(t, ok, "data.databricks_workspace_file.this has to be there")
		numFiles, _ := strconv.Atoi(r.Primary.Attributes["workspace_files.#"])
		assert.Equal(t, 2, numFiles)
		return nil
	}
}

func TestAccWorkspaceFileDataSourceNonRecursive(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: workspaceFileTemplate + `
		data "databricks_workspace_file" "this" {
			depends_on = [
				databricks_workspace_file.file_a,
				databricks_workspace_file.file_b,
				databricks_workspace_file.file_c,
			]

			path = databricks_directory.this.path
		}
		output "num_files" {
			value = length(data.databricks_workspace_file.this.workspace_files)
		}
		`,
		Check: checkWorkspaceFilesPopulated(t),
	})
}

func TestAccWorkspaceFileDataSourceEmptyDirectory(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: workspaceFileTemplate + `
		data "databricks_workspace_file" "this" {
			depends_on = [
				databricks_directory.empty_subdir,
			]

			path = databricks_directory.empty_subdir.path
		}
		`,
		Check: func(s *terraform.State) error {
			r, ok := s.Modules[0].Resources["data.databricks_workspace_file.this"]
			require.True(t, ok, "data.databricks_workspace_file.this has to be there")
			numFiles, _ := strconv.Atoi(r.Primary.Attributes["workspace_files.#"])
			assert.Equal(t, 0, numFiles)
			return nil
		},
	})
}

func TestAccWorkspaceFileDataSourceRecursive(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: workspaceFileTemplate + `
		data "databricks_workspace_file" "this" {
			depends_on = [
				databricks_workspace_file.file_a,
				databricks_workspace_file.file_b,
				databricks_workspace_file.file_c,
			]

			path      = databricks_directory.this.path
			recursive = true
		}
		output "num_files" {
			value = length(data.databricks_workspace_file.this.workspace_files)
		}
		`,
		Check: func(s *terraform.State) error {
			r, ok := s.Modules[0].Resources["data.databricks_workspace_file.this"]
			require.True(t, ok, "data.databricks_workspace_file.this has to be there")
			numFiles, _ := strconv.Atoi(r.Primary.Attributes["workspace_files.#"])
			assert.Equal(t, 3, numFiles)
			return nil
		},
	})
}
