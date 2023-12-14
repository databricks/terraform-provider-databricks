package acceptance

import (
	"testing"
)

func TestUcAccStorageCredential(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}`,
	})
}

func TestUcAccStorageCredentialUpdateOwnerOnly(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}`,
	}, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
			owner = "account users"
		}`,
	})
}

func TestUcAccStorageCredentialUpdateOwnerAndOtherFields(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}`,
	}, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF -- Updated Comment"
			owner = "account users"
		}`,
	})
}

func TestUcAccStorageCredentialUpdateFieldsOtherThanOwner(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}`,
	}, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF -- Updated Comment"
		}`,
	})
}
