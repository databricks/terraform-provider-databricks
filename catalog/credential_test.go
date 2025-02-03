package catalog_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func awsCredentialWithComment(comment string) string {
	return fmt.Sprintf(`
				resource "databricks_credential" "external" {
					name = "service-cred-{var.STICKY_RANDOM}"
					aws_iam_role {
						role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
					}
					purpose = "SERVICE"
					skip_validation = true
					comment = "%s"
				}`, comment)
}

func gcpCredentialWithComment(comment string) string {
	return fmt.Sprintf(`
				resource "databricks_credential" "external" {
					name = "service-cred-{var.STICKY_RANDOM}"
					databricks_gcp_service_account {}
					purpose = "SERVICE"
					skip_validation = true
					comment = "%s"
				}`, comment)
}

func TestUcAccCredential(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	if acceptance.IsAws(t) {
		acceptance.UnityWorkspaceLevel(t, acceptance.Step{
			Template: awsCredentialWithComment("Managed by TF"),
		}, acceptance.Step{
			Template: awsCredentialWithComment("Managed by TF updated"),
		})
	} else if acceptance.IsGcp(t) {
		acceptance.UnityWorkspaceLevel(t, acceptance.Step{
			Template: gcpCredentialWithComment("Managed by TF"),
		}, acceptance.Step{
			Template: gcpCredentialWithComment("Managed by TF updated"),
		})
	}
}

func TestAccCredentialOwner(t *testing.T) {
	acceptance.UnityAccountLevel(t, acceptance.Step{
		Template: `
			resource "databricks_service_principal" "test_acc_storage_credential_owner" {
				display_name = "test_acc_storage_credential_owner {var.RANDOM}"
			}

			resource "databricks_credential" "test_acc_storage_credential_owner" {
				name = "test_acc_storage_credential_owner-{var.RANDOM}"
				owner = databricks_service_principal.test_acc_storage_credential_owner.application_id
				purpose = "SERVICE"
				aws_iam_role {
					role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
				}
			}
		`,
	})
}
