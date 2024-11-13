package acceptance

import (
	"testing"
)

func TestUcAccCredential(t *testing.T) {
	loadUcwsEnv(t)
	if isAws(t) {
		UnityWorkspaceLevel(t, Step{
			Template: `
				resource "databricks_credential" "external" {
					name = "cred-{var.RANDOM}"
					aws_iam_role {
						role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
					}
					skip_validation = true
					comment = "Managed by TF"
				}`,
		})
	}
}

func TestAccCredentialOwner(t *testing.T) {
	UnityAccountLevel(t, Step{
		Template: `
			resource "databricks_service_principal" "test_acc_storage_credential_owner" {
				display_name = "test_acc_storage_credential_owner {var.RANDOM}"
			}

			resource "databricks_credential" "test_acc_storage_credential_owner" {
				name = "test_acc_storage_credential_owner-{var.RANDOM}"
				owner = databricks_service_principal.test_acc_storage_credential_owner.application_id
				metastore_id = "{env.TEST_METASTORE_ID}"
				aws_iam_role {
					role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
				}
			}
		`,
	})
}
