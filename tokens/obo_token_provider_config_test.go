package tokens_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func oboTokenProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_obo_token" "this" {
		application_id = "fake-app-id"
		%s
	}
	`, providerConfig)
}

func TestAccOboToken_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: oboTokenProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
