package acceptance

import (
	"fmt"
	"testing"
)

var (
	customAppIntegrationTemplate = `resource "databricks_custom_app_integration" "this" {
			name = "custom_integration_name"
			redirect_urls = ["https://example.com"]
			scopes = ["all"]
			token_access_policy {
				access_token_ttl_in_minutes = %s
				refresh_token_ttl_in_minutes = 30
			}
		}`
)

func TestMwsAccCustomAppIntegrationCreate(t *testing.T) {
	loadAccountEnv(t)
	AccountLevel(t, Step{
		Template: fmt.Sprintf(customAppIntegrationTemplate, "30"),
	})
}

func TestMwsAccCustomAppIntegrationUpdate(t *testing.T) {
	loadAccountEnv(t)
	AccountLevel(t, Step{
		Template: fmt.Sprintf(customAppIntegrationTemplate, "30"),
	}, Step{
		Template: fmt.Sprintf(customAppIntegrationTemplate, "60"),
	})
}
