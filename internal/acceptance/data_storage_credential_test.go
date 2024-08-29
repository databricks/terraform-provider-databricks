package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

func checkStorageCredentialDataSourcePopulated(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_storage_credential.this"]
		require.True(t, ok, "data.databricks_storage_credential.this has to be there")
		storage_credential_info := r.Primary.Attributes["storage_credential_info.0.%"]
		if storage_credential_info == "" {
			return fmt.Errorf("StorageCredentialInfo is empty: %v", r.Primary.Attributes)
		}
		return nil
	}
}
func TestUcAccDataSourceStorageCredential(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_storage_credential" "external" {
			name = "cred-{var.RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			skip_validation = true
			comment = "Managed by TF"
		}

		data "databricks_storage_credential" "this" {
			name = databricks_storage_credential.external.name
		}`,
		Check: checkStorageCredentialDataSourcePopulated(t),
	})
}
