package acceptance

import (
	"testing"
)

func TestUcAccVolumesResourceWithoutInitialOwnerAWSFullLifecycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-def"
			comment = "comment-def"
			owner = "account users"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}`,
	})
}

func TestUcAccVolumesResourceWithInitialOnwerAWSFullLifecycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			owner = "account users"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-def"
			comment = "comment-def"
			owner = "{env.TEST_DATA_ENG_GROUP}"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}`,
	})
}
