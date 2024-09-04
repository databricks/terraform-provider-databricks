package acceptance

import (
	"fmt"
	"testing"
)

const grantsTemplateForExternalLocation = `
	resource "databricks_grants" "some" {
		external_location = databricks_external_location.some.id
		grant {
			principal  = "{env.TEST_DATA_ENG_GROUP}"
			privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
		}
	}
`

func externalLocationTemplateWithOwner(comment string, owner string) string {
	return fmt.Sprintf(`
		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/some{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			isolation_mode  = "ISOLATION_MODE_ISOLATED"
			comment         = "%s"
			owner = "%s"
		}
	`, comment, owner)
}

func storageCredentialTemplateWithOwner(comment, owner string) string {
	return fmt.Sprintf(`
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment        = "%s"
			owner          = "%s"
			isolation_mode = "ISOLATION_MODE_ISOLATED"
			force_update   = true
		}
	`, comment, owner)
}

func TestUcAccExternalLocation(t *testing.T) {
	UnityWorkspaceLevel(t, Step{
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
		}` + grantsTemplateForExternalLocation,
	})
}

func TestUcAccExternalLocationForceDestroy(t *testing.T) {
	UnityWorkspaceLevel(t, Step{
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
		}` + grantsTemplateForExternalLocation,
	})
}

func TestUcAccExternalLocationUpdate(t *testing.T) {
	UnityWorkspaceLevel(t, Step{
		Template: storageCredentialTemplateWithOwner("Managed by TF", "account users") +
			externalLocationTemplateWithOwner("Managed by TF", "account users") +
			grantsTemplateForExternalLocation,
	}, Step{
		Template: storageCredentialTemplateWithOwner("Managed by TF -- Updated Comment", "account users") +
			externalLocationTemplateWithOwner("Managed by TF -- Updated Comment", "account users") +
			grantsTemplateForExternalLocation,
	}, Step{
		Template: storageCredentialTemplateWithOwner("Managed by TF -- Updated Comment", "{env.TEST_DATA_ENG_GROUP}") +
			externalLocationTemplateWithOwner("Managed by TF -- Updated Comment", "{env.TEST_DATA_ENG_GROUP}") +
			grantsTemplateForExternalLocation,
	}, Step{
		Template: storageCredentialTemplateWithOwner("Managed by TF -- Updated Comment 2", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}") +
			externalLocationTemplateWithOwner("Managed by TF -- Updated Comment 2", "{env.TEST_METASTORE_ADMIN_GROUP_NAME}") +
			grantsTemplateForExternalLocation,
	})
}
