package privateendpointrule_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

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

const nccPrivateEndpointRuleHCL = `
resource "databricks_mws_network_connectivity_config" "this" {
	name   = "tf-{var.STICKY_RANDOM}"
	region = "us-east-1"
}

resource "databricks_mws_ncc_private_endpoint_rule" "this" {
	network_connectivity_config_id = databricks_mws_network_connectivity_config.this.network_connectivity_config_id
	endpoint_service               = "com.amazonaws.us-east-1.s3"
	resource_names                 = ["{var.STICKY_RANDOM}-bucket"]
}
`

// pluginFrameworkProviderFactory returns a ProtoV6ProviderFactories map that
// forces the resource through the plugin framework implementation. Used by
// both the dedicated PF test and the second step of the migration test, since
// the resource is opt-in (default is SDKv2) and we can't rely on the
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

// TestMwsAccPrivateEndpointRulePluginFrameworkOptIn exercises the plugin
// framework implementation by explicitly opting in. The resource is in
// pluginFwOptInResources, so the default mux serves SDKv2; this test forces
// PF to keep coverage of the new implementation deterministic.
func TestMwsAccPrivateEndpointRulePluginFrameworkOptIn(t *testing.T) {
	acceptance.AccountLevel(t,
		acceptance.Step{
			ProtoV6ProviderFactories: pluginFrameworkProviderFactory(),
			Template:                 nccPrivateEndpointRuleHCL,
			Check:                    notCreatingAfterApply("databricks_mws_ncc_private_endpoint_rule.this"),
		},
		acceptance.Step{
			ProtoV6ProviderFactories:             pluginFrameworkProviderFactory(),
			ImportState:                          true,
			ResourceName:                         "databricks_mws_ncc_private_endpoint_rule.this",
			ImportStateIdFunc:                    acceptance.BuildImportStateIdFunc("databricks_mws_ncc_private_endpoint_rule.this", "id"),
			ImportStateVerifyIdentifierAttribute: "id",
		},
	)
}

// TestMwsAccPrivateEndpointRuleMigrationFromSDKv2 verifies behavioural parity
// between the legacy SDKv2 implementation and the new Plugin Framework
// implementation. Step 1 creates the resource via SDKv2 (the default now that
// this resource is opt-in); step 2 re-applies the same HCL with PF forced on.
// Terraform must report zero diff.
//
// Pattern lifted from sharing/resource_share_acc_test.go:333 TestUcAccShareMigrationFromSDKv2.
func TestMwsAccPrivateEndpointRuleMigrationFromSDKv2(t *testing.T) {
	acceptance.AccountLevel(t,
		acceptance.Step{
			Template: nccPrivateEndpointRuleHCL,
		},
		acceptance.Step{
			ProtoV6ProviderFactories: pluginFrameworkProviderFactory(),
			Template:                 nccPrivateEndpointRuleHCL,
			Check:                    notCreatingAfterApply("databricks_mws_ncc_private_endpoint_rule.this"),
		},
	)
}
