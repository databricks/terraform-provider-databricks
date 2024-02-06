package acceptance

import (
	"testing"
)

func TestUcAccCatalogWorkspaceBindingToOtherWorkspace(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		# The dummy workspace needs to be assigned to the metastore for this test to pass
		resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY_WORKSPACE_ID}
		}

		resource "databricks_catalog" "dev" {
			name           = "dev{var.RANDOM}"
			isolation_mode = "ISOLATED"
		}

		resource "databricks_catalog_workspace_binding" "test" {
			catalog_name = databricks_catalog.dev.name
			workspace_id = {env.DUMMY_WORKSPACE_ID} # dummy workspace, not the authenticated workspace in this test
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

func TestUcAccSecurableWorkspaceBindingToSameWorkspaceReadOnly(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "dev" {
			name           = "dev{var.RANDOM}"
			isolation_mode = "ISOLATED"
		}

		resource "databricks_catalog_workspace_binding" "test" {
			securable_name = databricks_catalog.dev.name
			securable_type = "catalog"
			workspace_id   = {env.THIS_WORKSPACE_ID}
			binding_type   = "BINDING_TYPE_READ_ONLY"
		}
		`,
	})
}
