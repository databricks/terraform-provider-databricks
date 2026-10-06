package terraformtest

import (
	"context"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Step describes one Terraform configuration and all HTTP traffic expected
// while Terraform refreshes, plans, and applies it.
type Step struct {
	// Name identifies this step in the fixture recording.
	Name string
	// Template is the Terraform configuration applied by this step.
	Template string
	// Check validates the resulting Terraform state after this step is applied. It works with both
	// SDKv2 and Plugin Framework resources.
	Check resource.TestCheckFunc
	// ConfigPlanChecks validates Terraform's provider-independent plan before apply or after apply,
	// before or after refresh.
	ConfigPlanChecks resource.ConfigPlanChecks
	// ExpectError requires this step to fail with an error matching the regular expression.
	ExpectError *regexp.Regexp
	// PlanOnly stops this step after Terraform creates the plan.
	PlanOnly bool
}

// TestCase describes a multi-step Terraform lifecycle and its final cleanup.
type TestCase struct {
	// FixtureFile is the JSONL recording path, relative to the calling test's package.
	FixtureFile string
	// HostMetadata is served as a test-global reusable response to host metadata requests. When nil,
	// those exchanges must be present in the JSONL recording.
	HostMetadata *config.HostMetadata
	// Steps are applied in order against the same Terraform state.
	Steps []Step
}

type fixturePhases interface {
	SetFixtures([]qa.HTTPFixture)
}

// Run executes a credential-free Terraform lifecycle against the strict local HTTP fixtures in
// TestCase.FixtureFile.
func Run(t *testing.T, testCase TestCase) {
	t.Helper()
	if testCase.FixtureFile == "" {
		t.Fatal("Terraform test case has no fixture file")
	}
	fixtures, err := loadFixtureFile(testCase.FixtureFile, testCase.Steps)
	if err != nil {
		t.Fatalf("load Terraform HTTP fixtures: %s", err)
	}
	handler := qa.NewHTTPFixtureHandler(t)
	server := httptest.NewServer(handler)
	t.Cleanup(func() {
		server.Close()
		handler.AssertAllFixturesConsumed()
	})
	fixtures = withHostMetadataFixtures(fixtures, testCase.HostMetadata, server.URL+"/oidc")

	providerFactories := providerFactories(server.URL)
	steps := terraformSteps(testCase, fixtures, handler, providerFactories)
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Steps:      steps,
	})
}

func terraformSteps(
	testCase TestCase,
	fixtures [][]qa.HTTPFixture,
	phases fixturePhases,
	providerFactories map[string]func() (tfprotov6.ProviderServer, error),
) []resource.TestStep {
	steps := make([]resource.TestStep, 0, len(testCase.Steps)+1)
	for stepIndex, step := range testCase.Steps {
		step := step
		stepFixtures := fixtures[stepIndex]
		steps = append(steps, resource.TestStep{
			PreConfig: func() {
				phases.SetFixtures(stepFixtures)
			},
			Config:                   step.Template,
			Check:                    step.Check,
			ConfigPlanChecks:         step.ConfigPlanChecks,
			ExpectError:              step.ExpectError,
			PlanOnly:                 step.PlanOnly,
			ProtoV6ProviderFactories: providerFactories,
		})
	}
	if len(testCase.Steps) == 0 {
		return steps
	}
	lastTemplate := testCase.Steps[len(testCase.Steps)-1].Template
	destroyFixtures := fixtures[len(testCase.Steps)]
	steps = append(steps, resource.TestStep{
		PreConfig: func() {
			phases.SetFixtures(destroyFixtures)
		},
		Config:                   lastTemplate,
		Destroy:                  true,
		ProtoV6ProviderFactories: providerFactories,
	})
	return steps
}

func providerFactories(host string) map[string]func() (tfprotov6.ProviderServer, error) {
	customizeConfig := func(cfg *config.Config) error {
		cfg.Host = host
		cfg.Token = "..."
		cfg.AuthType = "pat"
		cfg.AzureEnvironment = "PUBLIC"
		return nil
	}
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			sdkV2Provider := sdkv2.DatabricksProvider(sdkv2.WithConfigCustomizer(customizeConfig))
			pluginFrameworkProvider := pluginfw.GetDatabricksProviderPluginFramework(
				pluginfw.WithConfigCustomizer(customizeConfig),
			)
			return providers.GetProviderServer(
				context.Background(),
				providers.WithSdkV2Provider(sdkV2Provider),
				providers.WithPluginFrameworkProvider(pluginFrameworkProvider),
			)
		},
	}
}
