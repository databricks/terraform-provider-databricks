package terraformtest

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestProviderImplementationValidation(t *testing.T) {
	for _, tc := range []struct{ name, value, message string }{
		{"unknown implementation", "{databricks_share: framework}", "expected sdkv2 or plugin_framework"},
		{"empty implementation", "{databricks_share: ''}", "expected sdkv2 or plugin_framework"},
		{"empty resource type", "{'': sdkv2}", "requires a resource type name"},
		{"old boolean", "{databricks_share: true}", "expected sdkv2 or plugin_framework"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Replace(minimalDefinition, "name: create", "name: create\n    provider_implementations: "+tc.value, 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, "step 1 (create): provider_implementations")
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestProviderSelectionRejectsUnavailableImplementation(t *testing.T) {
	for _, tc := range []struct {
		name           string
		implementation providerImplementation
	}{
		{"databricks_unknown_resource", sdkv2Implementation},
		{"databricks_share.this", sdkv2Implementation},
		{"databricks_app", sdkv2Implementation},
		{"databricks_group", pluginFrameworkImplementation},
	} {
		t.Run(tc.name, func(t *testing.T) {
			factories := providerFactories(fixtureConfig("http://127.0.0.1"), providerImplementations{tc.name: tc.implementation})
			_, err := factories["databricks"]()
			require.ErrorContains(t, err, "provider_implementations[\""+tc.name+"\"]")
			require.ErrorContains(t, err, "implementation is not available")
		})
	}
}

func TestProviderSelectionSupportsMultipleResources(t *testing.T) {
	for name, selections := range map[string]providerImplementations{
		"two SDKv2 fallbacks":     {"databricks_share": sdkv2Implementation, "databricks_quality_monitor": sdkv2Implementation},
		"two framework resources": {"databricks_quality_monitor": pluginFrameworkImplementation, "databricks_mws_ncc_private_endpoint_rule": pluginFrameworkImplementation},
		"mixed implementations":   {"databricks_share": sdkv2Implementation, "databricks_quality_monitor": pluginFrameworkImplementation, "databricks_mws_ncc_private_endpoint_rule": pluginFrameworkImplementation},
	} {
		t.Run(name, func(t *testing.T) {
			factories := providerFactories(fixtureConfig("http://127.0.0.1"), selections)
			server, err := factories["databricks"]()
			require.NoError(t, err)
			response, err := server.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
			require.NoError(t, err)
			for _, diagnostic := range response.Diagnostics {
				require.NotEqual(t, tfprotov6.DiagnosticSeverityError, diagnostic.Severity, diagnostic.Detail)
			}
			for resource := range selections {
				require.Contains(t, response.ResourceSchemas, resource)
			}
			// Unselected prerequisites remain available through the production mux.
			require.Contains(t, response.ResourceSchemas, "databricks_schema")
			require.Contains(t, response.DataSourceSchemas, "databricks_current_user")
		})
	}
}

func TestProviderSelectionRetainsCleanupConfiguration(t *testing.T) {
	persistent, discarded, diagnostic := true, false, "expected error"
	first := step{Name: "apply", Config: "applied", ProviderImplementations: providerImplementations{"databricks_share": sdkv2Implementation}}
	for _, tc := range []struct {
		name       string
		steps      []step
		checkpoint bool
		retained   int
	}{
		{"successful migration", []step{first, {Name: "migrate", Config: "migrated"}}, false, 1},
		{"failed update", []step{first, {Name: "reject", Config: "invalid", ExpectError: &diagnostic}}, false, 0},
		{"preview after apply", []step{first, {Name: "preview", Config: "preview", PlanOnly: true}}, false, 0},
		{"discarded import", []step{first, {Name: "verify", Config: "import", Import: &importStep{Persist: &discarded}}}, false, 0},
		{"persistent import", []step{first, {Name: "adopt", Config: "import", Import: &importStep{Persist: &persistent}}}, false, 1},
		{"preview only", []step{{Name: "preview", Config: "preview", PlanOnly: true}}, false, 0},
		{"error only", []step{{Name: "reject", Config: "invalid", ExpectError: &diagnostic}}, false, -1},
		{"checkpoint import", []step{{Name: "verify", Config: "import", Import: &importStep{Persist: &discarded}}}, true, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := &testDefinition{Steps: tc.steps}
			if tc.checkpoint {
				d.InitialState = &initialState{Checkpoint: "checkpoint.tfstate"}
			}
			require.Equal(t, tc.retained, d.cleanupConfigSteps()[len(tc.steps)-1])
		})
	}
}

func TestProviderSelectionSurvivesFileStaging(t *testing.T) {
	d := &testDefinition{Steps: []step{{Name: "read", Config: `provider "databricks" {}`, ProviderImplementations: providerImplementations{"databricks_group": sdkv2Implementation}, fileContents: map[string][]byte{"input.txt": []byte("input")}}}}
	tc := &resource.TestCase{ProtoV6ProviderFactories: ProviderFactories(fixtureConfig("http://127.0.0.1")), Steps: testSteps(t, d, httpfixture.NewHandler())}
	configureProviderImplementations(d, tc, fixtureConfig("http://127.0.0.1"))
	tc.Steps[0].ProtoV6ProviderFactories["step-specific"] = nil
	configureStepFiles(t, d, tc)
	require.Nil(t, tc.ProtoV6ProviderFactories)
	require.Contains(t, tc.Steps[0].ProtoV6ProviderFactories, "step-specific")
	require.NotContains(t, tc.Steps[1].ProtoV6ProviderFactories, "step-specific")
}
