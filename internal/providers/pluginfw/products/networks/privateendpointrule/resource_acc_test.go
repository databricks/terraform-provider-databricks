package privateendpointrule_test

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

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

// TestMwsAccPrivateEndpointRulePluginFramework exercises the new Plugin
// Framework implementation through the default mux (no fallbacks). It is the
// baseline coverage of the migrated resource.
func TestMwsAccPrivateEndpointRulePluginFramework(t *testing.T) {
	acceptance.AccountLevel(t,
		acceptance.Step{
			Template: nccPrivateEndpointRuleHCL,
		},
		acceptance.Step{
			ImportState:                          true,
			ResourceName:                         "databricks_mws_ncc_private_endpoint_rule.this",
			ImportStateIdFunc:                    acceptance.BuildImportStateIdFunc("databricks_mws_ncc_private_endpoint_rule.this", "id"),
			ImportStateVerifyIdentifierAttribute: "id",
		},
	)
}

// TestMwsAccPrivateEndpointRuleMigrationFromSDKv2 verifies behavioural parity
// between the legacy SDKv2 implementation and the new Plugin Framework
// implementation. Step 1 creates the resource via the SDKv2 fallback; Step 2
// re-applies the same HCL via the default mux (Plugin Framework). Terraform
// must report zero diff.
//
// Pattern lifted from sharing/resource_share_acc_test.go:333 TestUcAccShareMigrationFromSDKv2.
func TestMwsAccPrivateEndpointRuleMigrationFromSDKv2(t *testing.T) {
	acceptance.AccountLevel(t,
		acceptance.Step{
			ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
				"databricks": func() (tfprotov6.ProviderServer, error) {
					sdkv2P, pluginfwP := acceptance.ProvidersWithResourceFallbacks([]string{"databricks_mws_ncc_private_endpoint_rule"})
					return providers.GetProviderServer(context.Background(),
						providers.WithSdkV2Provider(sdkv2P),
						providers.WithPluginFrameworkProvider(pluginfwP))
				},
			},
			Template: nccPrivateEndpointRuleHCL,
		},
		acceptance.Step{
			Template: nccPrivateEndpointRuleHCL,
		},
	)
}
