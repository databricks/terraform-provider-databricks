package acceptance

import (
	"fmt"
	"testing"
)

func workspaceBindingTemplateWithWorkspaceId(workspaceId string) string {
	return fmt.Sprintf(`
		# The dummy workspace needs to be assigned to the metastore for this test to pass
		resource "databricks_metastore_assignment" "this" {
			metastore_id = "{env.TEST_METASTORE_ID}"
			workspace_id = {env.DUMMY_WORKSPACE_ID}
		}

		resource "databricks_catalog" "dev" {
			name           = "dev{var.RANDOM}"
			isolation_mode = "ISOLATED"
		}

		resource "databricks_catalog" "prod" {
			name           = "prod{var.RANDOM}"
			isolation_mode = "ISOLATED"
		}			

		resource "databricks_workspace_binding" "dev" {
			catalog_name = databricks_catalog.dev.name
			workspace_id = %s
		}

		resource "databricks_workspace_binding" "prod" {
			securable_name = databricks_catalog.prod.name
			securable_type = "catalog"
			workspace_id   = %s
			binding_type   = "BINDING_TYPE_READ_ONLY"
		}			
	`, workspaceId, workspaceId)
}

func TestUcAccWorkspaceBindingToOtherWorkspace(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: workspaceBindingTemplateWithWorkspaceId("{env.DUMMY_WORKSPACE_ID}"),
	})
}

func TestUcAccWorkspaceBindingToSameWorkspace(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: workspaceBindingTemplateWithWorkspaceId("{env.THIS_WORKSPACE_ID}"),
	})
}
