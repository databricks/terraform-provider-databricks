package acceptance

import (
	"testing"
)

func TestUcAccCatalogWorkspaceBinding(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_catalog" "dev" {
			name = "dev{var.RANDOM}"
		}

		resource "databricks_catalog" "test" {
			name = "test{var.RANDOM}"
		}

		resource "databricks_catalog" "prod" {
			name = "prod{var.RANDOM}"
		}

		resource "databricks_catalog_workspace_binding" "dev" {
			catalog_name = databricks_catalog.dev.name
			workspace_id = {env.THIS_WORKSPACE_ID}
		}

		resource "databricks_catalog_workspace_binding" "test" {
			catalog_name = databricks_catalog.test.name
			workspace_id = {env.THIS_WORKSPACE_ID}
		}
		`,
	})
}
