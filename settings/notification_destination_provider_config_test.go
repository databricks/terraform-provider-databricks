package settings_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func notificationDestinationProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_notification_destination" "this" {
		display_name = "test-notification"
		config {
			email {
				addresses = ["test@example.com"]
			}
		}
		%s
	}
	`, providerConfig)
}

func TestAccNotificationDestination_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: notificationDestinationProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
