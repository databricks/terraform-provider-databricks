package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestUcAccExternalLocation(t *testing.T) {
	qa.RequireCloudEnv(t, "ucws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_storage_credential" "external" {
				name = "cred-{var.RANDOM}"
				aws_iam_role {
					role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
				}
				comment = "Managed by TF"
			}
			
			resource "databricks_external_location" "some" {
				name            = "external"
				url             = "s3://{env.TEST_BUCKET}/some{var.RANDOM}"
				credential_name = databricks_storage_credential.external.id
				comment         = "Managed by TF"
			}
			
			resource "databricks_grants" "some" {
				external_location = databricks_external_location.some.id
				grant {
					principal  = "{env.TEST_DATA_ENG_GROUP}"
					privileges = ["CREATE_TABLE", "READ_FILES"]
				}
			}`,
		},
	})
}
