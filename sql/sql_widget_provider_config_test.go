package sql_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func sqlWidgetProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_sql_widget" "this" {
		dashboard_id = "fake-dashboard-id"
		%s
	}
	`, providerConfig)
}

func TestAccSqlWidget_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: sqlWidgetProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
