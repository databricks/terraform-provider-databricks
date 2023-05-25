package acceptance

import (
	"testing"
)

func TestAccVolumesResourceFullLifecycle(t *testing.T) {
	workspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.RANDOM}"
			catalog_name = "main"
		}

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

		resource "databricks_volumes" "this" {
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
			name 		 = "schema-{var.RANDOM}"
			catalog_name = "main"
		}

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

		resource "databricks_volumes" "this" {
			name = "name-def"
			comment = "comment-def"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}`,
	})
}
