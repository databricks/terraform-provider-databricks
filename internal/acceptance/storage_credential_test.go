package acceptance

import (
	"os"
	"testing"
)

func TestUcAccStorageCredential(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "ucws":
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
	case "gcp-ucws":
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
