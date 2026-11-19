package apps_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var (
	customAppIntegrationTemplate = `resource "databricks_custom_app_integration" "this" {
			name = "tf-{var.STICKY_RANDOM}-app-integration"
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

func customAppIntegrationClearValuesTemplate(userAuthorizedScopes bool, policy string) string {
	scopes := ""
	if userAuthorizedScopes {
		scopes = `user_authorized_scopes = ["sql"]`
	}
	return fmt.Sprintf(`resource "databricks_custom_app_integration" "this" {
			name = "tf-{var.STICKY_RANDOM}-app-integration"
			redirect_urls = ["https://example.com"]
			scopes = ["all-apis", "sql"]
			%s
			%s
		}`, scopes, policy)
}

func TestMwsAccCustomAppIntegrationClearValues(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	acceptance.AccountLevel(t, acceptance.Step{
		Template: customAppIntegrationClearValuesTemplate(true, `token_access_policy {
				access_token_ttl_in_minutes = 30
				refresh_token_ttl_in_minutes = 720
				enable_single_use_refresh_tokens = true
				absolute_session_lifetime_in_minutes = 1440
			}`),
	}, acceptance.Step{
		Template: customAppIntegrationClearValuesTemplate(true, `token_access_policy {
				access_token_ttl_in_minutes = 30
				refresh_token_ttl_in_minutes = 720
				enable_single_use_refresh_tokens = false
			}`),
		Check: resource.TestCheckResourceAttr("databricks_custom_app_integration.this",
			"token_access_policy.0.enable_single_use_refresh_tokens", "false"),
	}, acceptance.Step{
		Template: customAppIntegrationClearValuesTemplate(false, `token_access_policy {
				access_token_ttl_in_minutes = 30
				refresh_token_ttl_in_minutes = 720
				enable_single_use_refresh_tokens = false
			}`),
		ExpectError: regexp.MustCompile(`user_authorized_scopes can't be removed`),
	}, acceptance.Step{
		Template: customAppIntegrationClearValuesTemplate(true, ""),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("databricks_custom_app_integration.this",
				"token_access_policy.0.access_token_ttl_in_minutes", "60"),
			resource.TestCheckResourceAttr("databricks_custom_app_integration.this",
				"token_access_policy.0.refresh_token_ttl_in_minutes", "10080"),
		),
	})
}

func TestMwsAccCustomAppIntegrationNoPolicy(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	acceptance.AccountLevel(t, acceptance.Step{
		Template: customAppIntegrationClearValuesTemplate(false, ""),
	})
}
