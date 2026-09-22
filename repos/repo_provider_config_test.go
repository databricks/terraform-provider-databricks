package repos_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func repoProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_repo" "this" {
		url = "https://github.com/databricks/databricks-sdk-go"
		%s
	}
	`, providerConfig)
}

func TestAccRepo_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: repoProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
