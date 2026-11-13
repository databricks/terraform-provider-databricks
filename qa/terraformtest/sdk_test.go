package terraformtest

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

func TestSdkImplementationValidation(t *testing.T) {
	for _, tc := range []struct{ name, value, message string }{
		{"unknown implementation", "{databricks_service_principal: modular}",
			"expected classic or sdk_go"},
		{"empty implementation", "{databricks_service_principal: ''}",
			"expected classic or sdk_go"},
		{"empty resource type", "{'': sdk_go}", "requires a resource type name"},
		{"resource without sdk-go", "{databricks_group: sdk_go}",
			"resource has no sdk_go implementation"},
		{"classic without sdk-go", "{databricks_group: classic}",
			"resource has no sdk_go implementation"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			selection := "name: create\n    sdk_implementations: " + tc.value
			input := strings.Replace(minimalDefinition, "name: create", selection, 1)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, "step 1 (create): sdk_implementations")
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestSdkImplementationSelectsSdkGoResources(t *testing.T) {
	selection := sdkImplementations{
		"databricks_service_principal": sdkGoImplementation,
		"databricks_other":             classicSdkImplementation,
	}
	require.Equal(t, []string{"databricks_service_principal"}, selection.sdkGoResources())
	require.Empty(t, sdkImplementations(nil).sdkGoResources())
}

func TestProviderFactoriesWithSdkGoRejectsUnavailableResource(t *testing.T) {
	factories := ProviderFactoriesWithSdkGo(fixtureConfig("http://127.0.0.1"),
		[]string{"databricks_group"})
	_, err := factories["databricks"]()
	require.ErrorContains(t, err,
		`sdk_implementations["databricks_group"]: resource has no sdk_go implementation`)
}

func TestSdkSelectionInstallsStepFactoriesAndRetainsCleanupSelection(t *testing.T) {
	sdkGo := sdkImplementations{"databricks_service_principal": sdkGoImplementation}
	d := &testDefinition{Steps: []step{
		{Name: "classic", Config: `provider "databricks" {}`},
		{Name: "sdk_go", Config: `provider "databricks" {}`, SdkImplementations: sdkGo},
	}}
	customize := fixtureConfig("http://127.0.0.1")
	tc := &resource.TestCase{
		ProtoV6ProviderFactories: ProviderFactories(customize),
		Steps:                    testSteps(t, d, httpfixture.NewHandler()),
	}
	configureProviderImplementations(d, tc, customize)
	require.Nil(t, tc.ProtoV6ProviderFactories)
	for i, s := range tc.Steps {
		require.Contains(t, s.ProtoV6ProviderFactories, "databricks", "step %d", i)
	}
	require.Equal(t, 1, d.cleanupConfigSteps()[len(d.Steps)-1])
}
