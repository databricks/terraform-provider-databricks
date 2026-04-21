package scim_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func groupMemberProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_group" "parent" {
		display_name = "tf-test-parent"
	}
	resource "databricks_group" "child" {
		display_name = "tf-test-child"
	}
	resource "databricks_group_member" "this" {
		group_id  = databricks_group.parent.id
		member_id = databricks_group.child.id
		%s
	}
	`, providerConfig)
}

func TestAccGroupMember_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groupMemberProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccGroupMember_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groupMemberProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
