package scim_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func servicePrincipalProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_service_principal" "this" {
		display_name = "tf-test-sp"
		%s
	}
	`, providerConfig)
}

func TestAccServicePrincipal_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: servicePrincipalProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
