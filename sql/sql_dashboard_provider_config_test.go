package sql_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func sqlDashboardProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_sql_dashboard" "this" {
		name = "test-dashboard"
		%s
	}
	`, providerConfig)
}

func TestAccSqlDashboard_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlDashboardProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
