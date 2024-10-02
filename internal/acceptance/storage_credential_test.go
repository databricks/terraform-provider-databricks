package acceptance

import (
	"testing"
)

func TestUcAccStorageCredential(t *testing.T) {
	loadUcwsEnv(t)
	if isAws(t) {
		UnityWorkspaceLevel(t, Step{
			Template: `
				resource "databricks_storage_credential" "external" {
					name = "cred-{var.RANDOM}"
					aws_iam_role {
						role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
					}
					skip_validation = true
					comment = "Managed by TF"
				}
				resource "databricks_storage_credential" "r2" {
					name = "r2-{var.RANDOM}"
					cloudflare_api_token {
						account_id = "1234"
						access_key_id = "1234"
						secret_access_key = "1234"
					}
					skip_validation = true
					comment = "Managed by TF"
				}`,
		})
	} else if isGcp(t) {
		UnityWorkspaceLevel(t, Step{
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
	UnityAccountLevel(t, Step{
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
