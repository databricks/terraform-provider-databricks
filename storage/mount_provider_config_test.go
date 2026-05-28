package storage_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func mountProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_mount" "this" {
		name = "test-mount"
		%s
	}
	`, providerConfig)
}

func TestAccMount_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: mountProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
