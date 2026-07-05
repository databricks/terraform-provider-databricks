package privateendpointrule_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const ruleAddr = "databricks_mws_ncc_private_endpoint_rule.this"

// notCreatingAfterApply guards the post-migration contract: Plugin Framework
// polls until the rule leaves CREATING before returning from Create. If a
// future change accidentally reverts to fire-and-forget, the resource would
// land in state with connection_state == "CREATING" and this assertion would
// catch it.
func notCreatingAfterApply(resourceAddr string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			return fmt.Errorf("resource %s not in state", resourceAddr)
		}
		if got := rs.Primary.Attributes["connection_state"]; got == "CREATING" {
			return fmt.Errorf("connection_state after apply should not be CREATING; Create must poll until the state settles")
		}
		return nil
	}
}

// nccRuleConfig returns the HCL for an NCC config plus a private endpoint rule
// on the current cloud. The cloud-specific inputs mirror the SDKv2 acceptance
// test (mws/resource_mws_ncc_private_endpoint_rule_acc_test.go) so the same
// config is known to apply in CI and the migration tests compare like for like.
// The AWS form (resource_names) and the Azure form (group_id/resource_id)
// exercise the two different unset-field shapes that the "" -> null mapping in
// fromAPI has to handle.
//
// No in-place Update step is exercised here: the Azure backend rejects updates
// to enabled ("Updating fields enabled ... is not supported"), and the other
// updatable fields (domain_names/resource_names) conflict with group_id, so the
// Azure form has no updatable field. The update_mask logic is covered by unit
// tests instead (TestUpdate_SendsMaskedFieldsToServer and friends).
func nccRuleConfig(t *testing.T) string {
	var ncc, ruleBase string
	switch {
	case acceptance.IsAzure(t):
		ncc = `
		resource "databricks_mws_network_connectivity_config" "this" {
			name   = "tf-{var.STICKY_RANDOM}"
			region = "eastus2"
		}`
		ruleBase = `
		resource "databricks_mws_ncc_private_endpoint_rule" "this" {
			network_connectivity_config_id = databricks_mws_network_connectivity_config.this.network_connectivity_config_id
			resource_id                    = "{env.TEST_AZURE_STORAGE_RESOURCE_ID}"
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
		// Documented as Azure-GA / AWS-Public-Preview only
		// (docs/resources/mws_ncc_private_endpoint_rule.md). No GCP form factor.
		acceptance.Skipf(t)("NCC private endpoint rule not supported on GCP")
	default:
		acceptance.Skipf(t)("unrecognized cloud env for NCC private endpoint rule")
	}
	return ncc + ruleBase + `
		}
		`
}

// pluginFrameworkProviderFactory returns a ProtoV6ProviderFactories map that
// forces the resource through the plugin framework implementation. Used by
// both the dedicated PF test and the migration tests, since the resource is
// opt-in (default is SDKv2) and we can't rely on the
// DATABRICKS_TF_ENABLED_PF_RESOURCES env var for acceptance runs.
func pluginFrameworkProviderFactory() map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			sdkv2P, pluginfwP := acceptance.ProvidersWithPluginFrameworkOverrides([]string{"databricks_mws_ncc_private_endpoint_rule"})
			return providers.GetProviderServer(context.Background(),
				providers.WithSdkV2Provider(sdkv2P),
				providers.WithPluginFrameworkProvider(pluginfwP))
		},
	}
}

// TestMwsAccPrivateEndpointRulePluginFrameworkOptIn exercises the PF CRUD
// lifecycle by explicitly opting in (the default mux serves SDKv2). It covers
// Create (polled to a settled state) and import with ImportStateVerify, which
// asserts the imported attributes match the post-create state attribute by
// attribute, catching any fromAPI divergence on the import path (e.g. an unset
// scalar surfacing as "" instead of null).
func TestMwsAccPrivateEndpointRulePluginFrameworkOptIn(t *testing.T) {
	base := nccRuleConfig(t)
	acceptance.AccountLevel(t,
		acceptance.Step{
			ProtoV6ProviderFactories: pluginFrameworkProviderFactory(),
			Template:                 base,
			Check:                    notCreatingAfterApply(ruleAddr),
		},
		acceptance.Step{
			ProtoV6ProviderFactories:             pluginFrameworkProviderFactory(),
			ImportState:                          true,
			ImportStateVerify:                    true,
			ResourceName:                         ruleAddr,
			ImportStateIdFunc:                    acceptance.BuildImportStateIdFunc(ruleAddr, "id"),
			ImportStateVerifyIdentifierAttribute: "id",
		},
	)
}

// TestMwsAccPrivateEndpointRuleMigrationFromSDKv2 verifies forward
// compatibility: state written by the legacy SDKv2 implementation must be
// adopted by the Plugin Framework implementation with zero diff. Step 1 creates
// the rule via SDKv2 (the current default); step 2 re-applies the identical HCL
// with PF forced on and asserts the plan is a no-op. The no-op assertion is the
// real compatibility check: any mapping difference (e.g. an unset scalar
// surfacing as "" instead of null, or enabled drifting) would show as an
// Update here and fail the step.
func TestMwsAccPrivateEndpointRuleMigrationFromSDKv2(t *testing.T) {
	base := nccRuleConfig(t)
	acceptance.AccountLevel(t,
		acceptance.Step{
			Template: base,
		},
		acceptance.Step{
			ProtoV6ProviderFactories: pluginFrameworkProviderFactory(),
			Template:                 base,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(ruleAddr, plancheck.ResourceActionNoop),
				},
			},
		},
	)
}

// TestMwsAccPrivateEndpointRuleMigrationToSDKv2 verifies backward compatibility
// (provider rollback): state written by the Plugin Framework implementation
// must be adopted by the legacy SDKv2 implementation with zero diff. Step 1
// creates the rule via PF; step 2 re-applies the identical HCL with the default
// SDKv2 provider and asserts a no-op plan. This protects users who opt in and
// later downgrade.
func TestMwsAccPrivateEndpointRuleMigrationToSDKv2(t *testing.T) {
	base := nccRuleConfig(t)
	acceptance.AccountLevel(t,
		acceptance.Step{
			ProtoV6ProviderFactories: pluginFrameworkProviderFactory(),
			Template:                 base,
			Check:                    notCreatingAfterApply(ruleAddr),
		},
		acceptance.Step{
			Template: base,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(ruleAddr, plancheck.ResourceActionNoop),
				},
			},
		},
	)
}
