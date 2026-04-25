package aws_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func userRoleProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_user" "this" {
		user_name = "test@example.com"
	}
	resource "databricks_user_role" "this" {
		user_id = databricks_user.this.id
		role    = "arn:aws:iam::999999999999:role/fake-role"
		%s
	}
	`, providerConfig)
}

func TestAccUserRole_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: userRoleProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccUserRole_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: userRoleProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
