package acceptance

import (
	"testing"
)

func TestUcAccStorageCredential(t *testing.T) {
	loadUcwsEnv(t)
	if isAws(t) {
		unityWorkspaceLevel(t, step{
			Template: `
				resource "databricks_storage_credential" "external" {
					name = "cred-{var.RANDOM}"
					aws_iam_role {
						role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
					}
					skip_validation = true
					comment = "Managed by TF"
				}`,
		})
	} else if isGcp(t) {
		unityWorkspaceLevel(t, step{
			Template: `
				resource "databricks_storage_credential" "external" {
					name = "cred-{var.RANDOM}"
					databricks_gcp_service_account {}
					skip_validation = true
					comment = "Managed by TF"
				}`,
		})
	}
}

func TestAccStorageCredentialOwner(t *testing.T) {
	unityAccountLevel(t, step{
		Template: `
			resource "databricks_service_principal" "test_acc_storage_credential_owner" {
				display_name = "test_acc_storage_credential_owner {var.RANDOM}"
			}

			resource "databricks_storage_credential" "test_acc_storage_credential_owner" {
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
