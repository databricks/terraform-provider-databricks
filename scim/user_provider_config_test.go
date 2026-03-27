package scim_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func userProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_user" "this" {
		user_name = "test@example.com"
		%s
	}
	`, providerConfig)
}

func TestAccUser_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: userProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccUser_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: userProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
