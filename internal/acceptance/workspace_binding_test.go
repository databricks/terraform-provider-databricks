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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			isolation_mode = "ISOLATION_MODE_ISOLATED"
		}
		
		resource "databricks_external_location" "some" {
			name            = "external-{var.RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/some{var.RANDOM}"
			credential_name = databricks_storage_credential.external.id
			isolation_mode  = "ISOLATION_MODE_ISOLATED"
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

		resource "databricks_workspace_binding" "ext" {
			securable_name = databricks_external_location.some.id
			securable_type = "external_location"
			workspace_id   = %s
		}

		resource "databricks_workspace_binding" "cred" {
			securable_name = databricks_storage_credential.external.id
			securable_type = "storage_credential"
			workspace_id   = %s
		}		
	`, workspaceId, workspaceId, workspaceId, workspaceId)
}

func TestUcAccWorkspaceBindingToOtherWorkspace(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: workspaceBindingTemplateWithWorkspaceId("{env.DUMMY_WORKSPACE_ID}"),
	})
}
