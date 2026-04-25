package catalog_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func metastoreProviderConfigTemplate(name string, providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_metastore" "this" {
		name          = "%s"
		force_destroy = true
		%s
	}
	`, name, providerConfig)
}

func TestAccMetastore_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreProviderConfigTemplate("tf-test-metastore-{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccMetastore_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreProviderConfigTemplate("tf-test-metastore-{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccMetastore_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: metastoreProviderConfigTemplate("tf-test-metastore-{var.STICKY_RANDOM}", `
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}
