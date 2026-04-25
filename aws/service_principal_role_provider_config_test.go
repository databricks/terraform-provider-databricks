package aws_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func servicePrincipalRoleProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_service_principal" "this" {
		display_name = "tf-test-sp"
	}
	resource "databricks_service_principal_role" "this" {
		service_principal_id = databricks_service_principal.this.id
		role                 = "arn:aws:iam::999999999999:role/fake-role"
		%s
	}
	`, providerConfig)
}

func TestAccServicePrincipalRole_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: servicePrincipalRoleProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccServicePrincipalRole_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: servicePrincipalRoleProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
