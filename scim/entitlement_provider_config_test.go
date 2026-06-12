package scim_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func entitlementProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_entitlements" "this" {
		group_id = "123"
		%s
	}
	`, providerConfig)
}

func TestAccEntitlements_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: entitlementProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}

// TestMwsAccEntitlements_AccountProvider_ProviderConfig verifies that
// databricks_entitlements honors provider_config { workspace_id } when
// configured with an account-level provider. Regression test for the bug
// where Create's AccountHost guard fired before DatabricksClientForUnifiedProvider
// was called, making provider_config non-functional for entitlements.
func TestMwsAccEntitlements_AccountProvider_ProviderConfig(t *testing.T) {
	acceptance.AccountLevel(t, acceptance.Step{
		Template: `
		resource "databricks_service_principal" "this" {
			display_name = "tf-entitlements-{var.STICKY_RANDOM}"
		}

		resource "databricks_mws_permission_assignment" "this" {
			workspace_id = "{env.TEST_WORKSPACE_ID}"
			principal_id = databricks_service_principal.this.id
			permissions  = ["USER"]
		}

		resource "databricks_entitlements" "this" {
			service_principal_id  = databricks_service_principal.this.id
			databricks_sql_access = true

			provider_config {
				workspace_id = "{env.TEST_WORKSPACE_ID}"
			}

			depends_on = [databricks_mws_permission_assignment.this]
		}
		`,
	})
}
