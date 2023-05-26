package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestUcAccVolumesResourceWithoutInitialOwnerFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	unityWorkspaceLevel(t, step{
		Template: fmt.Sprintf(`
		resource "databricks_schema" "this" {
			name 		 = "schema-%[1]s"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-%[1]s"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-%[1]s"
			url             = "s3://{env.TEST_BUCKET}/somepath-%[1]s"
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
		}`, randomName),
	}, step{
		Template: fmt.Sprintf(`
		resource "databricks_schema" "this" {
			name 		 = "schema-%[1]s"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-%[1]s"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-%[1]s"
			url             = "s3://{env.TEST_BUCKET}/somepath-%[1]s"
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
		}`, randomName),
	})
}

func TestUcAccVolumesResourceWithInitialOnwerFullLifecycle(t *testing.T) {
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	unityWorkspaceLevel(t, step{
		Template: fmt.Sprintf(`
		resource "databricks_schema" "this" {
			name 		 = "schema-%[1]s"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-%[1]s"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-%[1]s"
			url             = "s3://{env.TEST_BUCKET}/somepath-%[1]s"
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
		}`, randomName),
	}, step{
		Template: fmt.Sprintf(`
		resource "databricks_schema" "this" {
			name 		 = "schema-%[1]s"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-%[1]s"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-%[1]s"
			url             = "s3://{env.TEST_BUCKET}/somepath-%[1]s"
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
		}`, randomName),
	})
}
