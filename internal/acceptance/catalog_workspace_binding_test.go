package acceptance

import (
	"testing"
)

func TestUcAccCatalogWorkspaceBindingToOtherWorkspace(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "dev" {
			name           = "dev{var.RANDOM}"
			isolation_mode = "ISOLATED"
		}

		resource "databricks_catalog_workspace_binding" "test" {
			catalog_name = databricks_catalog.dev.name
			workspace_id = {env.TEST_WORKSPACE_ID} # dummy workspace, not the authenticated workspace in this test
		}
		`,
	})
}

func TestUcAccCatalogWorkspaceBindingToSameWorkspace(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "dev" {
			name           = "dev{var.RANDOM}"
			isolation_mode = "ISOLATED"
		}

		resource "databricks_catalog_workspace_binding" "test" {
			catalog_name = databricks_catalog.dev.name
			workspace_id = {env.THIS_WORKSPACE_ID}
		}
		`,
	})
}
