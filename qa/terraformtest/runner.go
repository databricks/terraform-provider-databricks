package terraformtest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"

	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

// run keeps the fixture server alive through the framework's deferred destroy,
// then checks for requests or fixtures left over after cleanup.
func run(t *testing.T, definition *testDefinition) {
	t.Helper()
	configureTerraform(t)
	handler := httpfixture.NewHandler()
	handler.SetEnforceCallOrder(definition.EnforceCallOrder)
	if err := handler.SetReusableFixtures(definition.metadataCalls()); err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(handler)
	t.Cleanup(func() {
		server.Close()
		if err := handler.Verify(); err != nil {
			t.Error(err)
		}
	})
	steps := testSteps(t, definition, handler)
	testCase := resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: providerFactories(server.URL),
		Steps:                    steps,
		ErrorCheck: func(err error) error {
			// The testing framework attempts destroy after an interrupted lifecycle.
			cleanupErr := handler.StartStep(stepLabel(len(definition.Steps), "destroy"), definition.Destroy.Calls)
			return errors.Join(err, cleanupErr)
		},
	}
	if definition.InitialState != nil {
		seedInitialCheckpoint(t, definition, &testCase)
	}
	resource.Test(t, testCase)
}

// configureTerraform disables implicit CLI downloads and optional extra
// refreshes so machine-local settings cannot change a recorded HTTP transcript.
func configureTerraform(t *testing.T) {
	t.Helper()
	path := os.Getenv("TF_ACC_TERRAFORM_PATH")
	if path == "" {
		path = "terraform"
	}
	resolved, err := exec.LookPath(path)
	if err != nil {
		t.Fatal("Terraform CLI is required: install it on PATH or set TF_ACC_TERRAFORM_PATH")
	}
	t.Setenv("TF_ACC_TERRAFORM_PATH", resolved)
	t.Setenv("TF_ACC_TERRAFORM_VERSION", "")
	t.Setenv("CHECKPOINT_DISABLE", "1")
	t.Setenv("TF_IN_AUTOMATION", "1")
	// Checkpoints refer to the real provider address, not HashiCorp's test default.
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "databricks")
	t.Setenv("TF_ACC_PROVIDER_HOST", "registry.terraform.io")
	// Optional framework refresh modes would change the required HTTP transcript.
	t.Setenv("TF_ACC_REFRESH_AFTER_APPLY", "")
	if err := os.Unsetenv("TF_ACC_REFRESH_AFTER_APPLY"); err != nil {
		t.Fatal(err)
	}
}

// testSteps appends a numbered destroy using the last successful configuration.
// Cleanup always requires empty state, independently of optional step checks.
func testSteps(t *testing.T, definition *testDefinition, handler *httpfixture.Handler) []resource.TestStep {
	t.Helper()
	steps := make([]resource.TestStep, 0, len(definition.Steps)+1)
	cleanupConfig := `provider "databricks" {}`
	hasApplyConfig := false
	for i, s := range definition.Steps {
		label := stepLabel(i, s.Name)
		testStep := terraformStep(t, handler, label, s.Config, s.Calls, s.State)
		configurePlan(&testStep, s, label)
		if s.ExpectError != nil {
			testStep.ExpectError = regexp.MustCompile(*s.ExpectError)
		} else if !s.PlanOnly || !hasApplyConfig {
			cleanupConfig = s.Config
			hasApplyConfig = !s.PlanOnly
		}
		steps = append(steps, testStep)
	}
	label := stepLabel(len(definition.Steps), "destroy")
	// Invalid resource configuration must not prevent cleanup of earlier steps.
	cleanup := terraformStep(t, handler, label, cleanupConfig, definition.Destroy.Calls, nil)
	cleanup.Destroy = true
	cleanup.ConfigStateChecks = []statecheck.StateCheck{emptyStateAssertion{label: label}}
	steps = append(steps, cleanup)
	return steps
}

// terraformStep treats plan, apply, and consistency reads as one fixture phase.
func terraformStep(t *testing.T, handler *httpfixture.Handler, label, hcl string, calls []httpfixture.Call, state *expectedState) resource.TestStep {
	t.Helper()
	s := resource.TestStep{
		Config: hcl,
		PreConfig: func() {
			t.Log(label)
			if err := handler.StartStep(label, calls); err != nil {
				t.Error(err)
			}
		},
		PostApplyFunc: func() {
			// Terraform's post-apply consistency plan can consume additional reads.
			if err := handler.Verify(); err != nil {
				t.Error(err)
			}
		},
	}
	if state != nil && len(state.Resources) != 0 {
		s.ConfigStateChecks = []statecheck.StateCheck{stateAssertion{label: label, expected: *state}}
	}
	return s
}

// providerFactories uses the production mux so both provider implementations
// share the same local HTTP boundary.
func providerFactories(host string) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			customize := fixtureConfig(host)
			return providers.GetProviderServer(context.Background(),
				providers.WithSdkV2Provider(sdkv2.DatabricksProvider(sdkv2.WithConfigCustomizer(customize))),
				providers.WithPluginFrameworkProvider(pluginfw.GetDatabricksProviderPluginFramework(pluginfw.WithConfigCustomizer(customize))))
		},
	}
}

type fixtureLoader struct{}

func (fixtureLoader) Name() string                   { return "fixture configuration" }
func (fixtureLoader) Configure(*config.Config) error { return nil }

// fixtureConfig forces local PAT authentication while retaining routing fields,
// such as account_id, explicitly set in HCL.
func fixtureConfig(host string) func(*config.Config) error {
	return func(cfg *config.Config) error {
		cfg.Host, cfg.Token, cfg.AuthType = host, "fixture-token", "pat"
		cfg.Credentials = config.PatCredentials{}
		// An explicit loader prevents ambient credentials and profiles from filling
		// configuration fields, including account/workspace routing identifiers.
		cfg.Loaders = []config.Loader{fixtureLoader{}}
		cfg.HTTPTransport = fixtureTransport{host: host, transport: &http.Transport{}}
		return nil
	}
}

type fixtureTransport struct {
	host      string
	transport http.RoundTripper
}

// RoundTrip also rejects redirects outside the fixture server, preventing
// fixture responses from causing accidental live API calls.
func (f fixtureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	expected, err := url.Parse(f.host)
	if err != nil {
		return nil, err
	}
	if req.URL.Scheme != expected.Scheme || req.URL.Host != expected.Host {
		return nil, fmt.Errorf("hermetic replay rejected request outside fixture server: %s", req.URL)
	}
	return f.transport.RoundTrip(req)
}
