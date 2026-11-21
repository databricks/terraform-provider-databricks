package workspace_file_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

const workspaceFilePathsTemplate = `
resource "databricks_directory" "this" {
	path = "/Shared/provider-test/wsfp_{var.RANDOM}"
}

resource "databricks_workspace_file" "top" {
	path           = "${databricks_directory.this.path}/top.txt"
	content_base64 = base64encode("top")
}

resource "databricks_workspace_file" "nested" {
	path           = "${databricks_directory.this.path}/sub/nested.txt"
	content_base64 = base64encode("nested")
}

resource "databricks_notebook" "this" {
	path           = "${databricks_directory.this.path}/notebook"
	language       = "PYTHON"
	content_base64 = base64encode("# Databricks notebook source\nprint(1)")
}

data "databricks_workspace_file_paths" "flat" {
	path       = databricks_directory.this.path
	depends_on = [databricks_workspace_file.top, databricks_workspace_file.nested, databricks_notebook.this]
}

data "databricks_workspace_file_paths" "recursive" {
	path       = databricks_directory.this.path
	recursive  = true
	depends_on = [databricks_workspace_file.top, databricks_workspace_file.nested, databricks_notebook.this]
}
`

func TestAccDataSourceWorkspaceFilePaths(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: workspaceFilePathsTemplate,
		Check: func(s *terraform.State) error {
			flat, ok := s.Modules[0].Resources["data.databricks_workspace_file_paths.flat"]
			require.True(t, ok, "data.databricks_workspace_file_paths.flat has to be there")
			assert.Equal(t, "1", flat.Primary.Attributes["workspace_file_path_list.#"])

			recursive, ok := s.Modules[0].Resources["data.databricks_workspace_file_paths.recursive"]
			require.True(t, ok, "data.databricks_workspace_file_paths.recursive has to be there")
			assert.Equal(t, "2", recursive.Primary.Attributes["workspace_file_path_list.#"])
			return nil
		},
	})
}
