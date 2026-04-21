package access_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func sqlPermissionsProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_sql_permissions" "this" {
		database = "default"
		%s
	}
	`, providerConfig)
}

func TestAccSqlPermissions_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccSqlPermissions_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

func TestAccSqlPermissions_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlPermissionsProviderConfigTemplate(`
			provider_config {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id mismatch.*please check the workspace_id provided in provider_config`),
		PlanOnly:    true,
	})
}
