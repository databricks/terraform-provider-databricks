package provider_migration_test

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/databricks/terraform-provider-databricks/qa/replay/providertest"
	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestProviderMigrations(t *testing.T) {
	for _, tc := range []struct {
		name      string
		resource  string
		framework []bool
	}{
		{"TestMwsAccPrivateEndpointRuleMigrationFromSDKv2_AWS", "databricks_mws_ncc_private_endpoint_rule", []bool{false, true}},
		{"TestMwsAccPrivateEndpointRuleMigrationFromSDKv2_Azure", "databricks_mws_ncc_private_endpoint_rule", []bool{false, true}},
		{"TestMwsAccPrivateEndpointRuleMigrationToSDKv2_AWS", "databricks_mws_ncc_private_endpoint_rule", []bool{true, false}},
		{"TestMwsAccPrivateEndpointRuleMigrationToSDKv2_Azure", "databricks_mws_ncc_private_endpoint_rule", []bool{true, false}},
		{"TestMwsAccPrivateEndpointRulePluginFrameworkOptIn_AWS", "databricks_mws_ncc_private_endpoint_rule", []bool{true, true}},
		{"TestMwsAccPrivateEndpointRulePluginFrameworkOptIn_Azure", "databricks_mws_ncc_private_endpoint_rule", []bool{true, true}},
		{"TestUcAccShareMigrationFromPluginFramework", "databricks_share", []bool{true, false}},
		{"TestUcAccShareMigrationFromSDKv2", "databricks_share", []bool{false, true}},
		{"TestUcAccShareVolume", "databricks_share", []bool{false, true}},
		{"TestUcAccUpdateQualityMonitorTransitionFromPluginFw", "databricks_quality_monitor", []bool{true, false}},
		{"TestUcAccUpdateQualityMonitorTransitionFromSdkV2", "databricks_quality_monitor", []bool{false, true}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			terraformtest.RunNative(t, "testdata/"+tc.name+".yaml", func(testCase *resource.TestCase, customize func(*config.Config) error) {
				testCase.ProtoV6ProviderFactories = nil
				for i := range testCase.Steps {
					index := min(i, len(tc.framework)-1)
					// Destroy uses the final implementation and the state left by the journey.
					testCase.Steps[i].ProtoV6ProviderFactories = providertest.Factories(customize, tc.resource, tc.framework[index])
				}
			})
		})
	}
}
