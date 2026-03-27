package permissions_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func accessControlRuleSetProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_access_control_rule_set" "this" {
		name = "accounts/0000000000000000/servicePrincipals/0000000000000000/ruleSets/default"
		grant_rules {
			principals = ["groups/admins"]
			role       = "roles/servicePrincipal.manager"
		}
		%s
	}
	`, providerConfig)
}

func TestAccAccessControlRuleSet_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: accessControlRuleSetProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccAccessControlRuleSet_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: accessControlRuleSetProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
