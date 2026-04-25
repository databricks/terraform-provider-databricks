package catalog_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func metastoreDataAccessProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_metastore_data_access" "this" {
		metastore_id = "fake-metastore-id"
		name         = "tf-test-dac"
		aws_iam_role {
			role_arn = "arn:aws:iam::123456789012:role/tf-test"
		}
		%s
	}
	`, providerConfig)
}

func TestAccMetastoreDataAccess_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreDataAccessProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccMetastoreDataAccess_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreDataAccessProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccMetastoreDataAccess_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreDataAccessProviderConfigTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}
