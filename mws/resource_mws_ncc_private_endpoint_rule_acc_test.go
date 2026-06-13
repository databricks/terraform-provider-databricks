package mws_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Reproduces ES-1957213 / #5347.
//
// Step 1 sets account_id explicitly in HCL — necessary in CI because
// common.StructToData (common/reflect_resource.go:715) deliberately skips
// writing API-returned values to state for non-Computed, non-bool fields the
// user didn't configure. Without the explicit set, account_id never lands in
// state in this environment and the bug can't manifest here. The customer
// hits the bug because their state has account_id from a path that bypasses
// that skip-logic (terraform import goes through MarkNewResource, an older
// provider version, etc.); the fix below addresses the underlying schema bug
// regardless of how account_id got into state.
//
// Step 2 removes account_id from HCL. Before the fix the Update CRUD is
// invoked but the spurious account_id drift isn't one of the updatable fields
// (enabled/domain_names/resource_names), so the request goes out with an
// empty update_mask and the backend rejects with "Update mask must be
// specified" — exactly the customer-reported error. After the fix the schema
// is Optional+Computed, the diff engine preserves state, no Update fires.
//
// Step 3 re-plans with the same HCL — asserts the plan stays empty on
// subsequent runs.
func TestMwsAccNccPrivateEndpointRule_NoSpuriousAccountIdDrift(t *testing.T) {
	var ncc, ruleBase string
	switch {
	case acceptance.IsAzure(t):
		ncc = `
		resource "databricks_mws_network_connectivity_config" "this" {
			name = "tf-{var.STICKY_RANDOM}"
			region = "eastus2"
		}`
		ruleBase = `
		resource "databricks_mws_ncc_private_endpoint_rule" "this" {
			network_connectivity_config_id = databricks_mws_network_connectivity_config.this.network_connectivity_config_id
			resource_id                    = "/subscriptions/2a5a4578-9ca9-47e2-ba46-f6ee6cc731f2/resourceGroups/deco-prod-azure-eastus2-rg/providers/Microsoft.Storage/storageAccounts/decotestprodunity"
			group_id                       = "blob"`
	case acceptance.IsAws(t):
		ncc = `
		resource "databricks_mws_network_connectivity_config" "this" {
			account_id = "{env.DATABRICKS_ACCOUNT_ID}"
			name       = "tf-{var.STICKY_RANDOM}"
			region     = "{env.AWS_REGION}"
		}`
		ruleBase = `
		resource "databricks_mws_ncc_private_endpoint_rule" "this" {
			network_connectivity_config_id = databricks_mws_network_connectivity_config.this.network_connectivity_config_id
			resource_names                 = ["{env.TEST_LOGDELIVERY_BUCKET}"]`
	case acceptance.IsGcp(t):
		// databricks_mws_ncc_private_endpoint_rule is documented as Azure-GA /
		// AWS-Public-Preview only (docs/resources/mws_ncc_private_endpoint_rule.md).
		// No GCP form factor today, so skip cleanly.
		acceptance.Skipf(t)("NCC private endpoint rule not supported on GCP")
	default:
		acceptance.Skipf(t)("unrecognized cloud env for NCC private endpoint rule")
	}
	withAccountID := ncc + ruleBase + `
			account_id                     = "{env.DATABRICKS_ACCOUNT_ID}"
		}
		`
	withoutAccountID := ncc + ruleBase + `
		}
		`
	acceptance.AccountLevel(t,
		acceptance.Step{
			Template: withAccountID,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["databricks_mws_ncc_private_endpoint_rule.this"]
				if !ok {
					return fmt.Errorf("resource not found in state")
				}
				if r.Primary.Attributes["account_id"] == "" {
					return fmt.Errorf("account_id not populated in state after Step 1; bug precondition not met. attrs: %v", r.Primary.Attributes)
				}
				return nil
			},
		},
		acceptance.Step{
			Template: withoutAccountID,
		},
		acceptance.Step{
			Template: withoutAccountID,
			PlanOnly: true,
		},
	)
}
