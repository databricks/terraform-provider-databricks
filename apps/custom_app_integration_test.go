package apps_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

var (
	customAppIntegrationTemplate = `resource "databricks_custom_app_integration" "this" {
			name = "custom_integration_name"
			redirect_urls = ["https://example.com"]
			scopes = ["all-apis"]
			token_access_policy {
				access_token_ttl_in_minutes = %s
				refresh_token_ttl_in_minutes = 30
			}
		}`
)

func TestMwsAccCustomAppIntegrationCreate(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	acceptance.AccountLevel(t, acceptance.Step{
		Template: fmt.Sprintf(customAppIntegrationTemplate, "30"),
	})
}

func TestMwsAccCustomAppIntegrationUpdate(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	acceptance.AccountLevel(t, acceptance.Step{
		Template: fmt.Sprintf(customAppIntegrationTemplate, "30"),
	}, acceptance.Step{
		Template: fmt.Sprintf(customAppIntegrationTemplate, "15"),
	})
}
