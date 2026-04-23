package scim_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func groupRoleProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_group" "this" {
		display_name = "tf-test-group"
	}
	resource "databricks_group_role" "this" {
		group_id = databricks_group.this.id
		role     = "arn:aws:iam::999999999999:role/fake-role"
		%s
	}
	`, providerConfig)
}

func TestAccGroupRole_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groupRoleProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccGroupRole_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groupRoleProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
