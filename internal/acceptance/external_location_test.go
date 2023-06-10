package acceptance

import (
	"testing"
)

func TestUcAccExternalLocation(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}
		
		resource "databricks_external_location" "some" {
			name            = "external-{var.RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/some{var.RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}
		
		resource "databricks_grants" "some" {
			external_location = databricks_external_location.some.id
			grant {
				principal  = "{env.TEST_DATA_ENG_GROUP}"
				privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
			}
		}`,
	})
}

func TestUcAccExternalLocationForceDestroy(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}
		
		resource "databricks_external_location" "some" {
			name            = "external-{var.RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/some{var.RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
			force_destroy   = true
		}
		
		resource "databricks_grants" "some" {
			external_location = databricks_external_location.some.id
			grant {
				principal  = "{env.TEST_DATA_ENG_GROUP}"
				privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
			}
		}`,
	})
}
