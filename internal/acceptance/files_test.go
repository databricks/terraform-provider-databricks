package acceptance

import (
	"testing"
)

func TestUcAccFilesFullLifeCycle(t *testing.T) {
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
		}
		
		resource "databricks_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
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
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}
		
		resource "databricks_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python2.py"
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
	})
}

func TestUcAccFilesBase64FullLifeCycle(t *testing.T) {
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
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
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
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjDg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
	})
}
