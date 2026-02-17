package tokens_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func tokenProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_token" "this" {
		comment = "test-token"
		%s
	}
	`, providerConfig)
}

func TestAccToken_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: tokenProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccToken_ProviderConfig_Required(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: tokenProviderConfigTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccToken_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: tokenProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
