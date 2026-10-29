package terraformtest

import (
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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

// run keeps the fixture server alive through the framework's deferred destroy,
// then checks for requests or fixtures left over after cleanup.
func run(t *testing.T, definition *testDefinition) {
	runNative(t, definition, nil)
}

// RunNative runs one fixture file with native HashiCorp test customization.
// It is intended for Go scenarios requiring per-step provider factories or Go
// checks. The customizer supplies isolated authentication and the same strict
// local HTTP boundary as YAML replay; retain the generated cleanup hooks.
func RunNative(t *testing.T, path string, configure func(*resource.TestCase, func(*config.Config) error)) {
	t.Helper()
	definition, err := loadDefinition(path)
	if err != nil {
		t.Fatal(err)
	}
	runNative(t, definition, configure)
}

func runNative(t *testing.T, definition *testDefinition, configure func(*resource.TestCase, func(*config.Config) error)) {
	t.Helper()
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
	runTerraform(t, definition, handler, fixtureConfig(server.URL), configure)
}

// stepHTTP is shared by replay and recording. Terraform operations and checks
// stay identical; only the HTTP boundary changes.
type stepHTTP interface {
	StartStep(string, []httpfixture.Call) error
	Verify() error
}

func runTerraform(t *testing.T, definition *testDefinition, handler stepHTTP, customize func(*config.Config) error, configure func(*resource.TestCase, func(*config.Config) error)) {
	t.Helper()
	configureTerraform(t)
	steps := testSteps(t, definition, handler)
	testCase := resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: ProviderFactories(customize),
		Steps:                    steps,
		ErrorCheck: func(err error) error {
			// The testing framework attempts destroy after an interrupted lifecycle.
			cleanupErr := handler.StartStep(stepLabel(len(definition.Steps), "destroy"), definition.Destroy.Calls)
			return errors.Join(err, cleanupErr)
		},
	}
	configureProviderImplementations(definition, &testCase, customize)
	if definition.InitialState != nil {
		seedInitialCheckpoint(t, definition, &testCase)
	}
	configureStepFiles(t, definition, &testCase)
	if configure != nil {
		configure(&testCase, customize)
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
func testSteps(t *testing.T, definition *testDefinition, handler stepHTTP) []resource.TestStep {
	t.Helper()
	steps := make([]resource.TestStep, 0, len(definition.Steps)+1)
	cleanupConfig := `provider "databricks" {}`
	cleanupSteps := definition.cleanupConfigSteps()
	for i, s := range definition.Steps {
		label := stepLabel(i, s.Name)
		hcl := s.Config
		if s.Import != nil && hcl == "" {
			hcl = cleanupConfig
		}
		testStep := configuredStep(t, handler, label, hcl, s)
		if cleanupSteps[i] == i {
			cleanupConfig = hcl
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

// configuredStep maps a YAML operation to native checks without changing the
// configuration retained for later cleanup.
func configuredStep(t *testing.T, handler stepHTTP, label, hcl string, s step) resource.TestStep {
	t.Helper()
	native := terraformStep(t, handler, label, hcl, s.Calls, s.State)
	configurePlan(&native, s, label)
	if s.Import != nil {
		if err := configureImport(&native, *s.Import, handler); err != nil {
			t.Fatalf("%s: %v", label, err)
		}
	}
	if s.ExpectError != nil {
		native.ExpectError = regexp.MustCompile(*s.ExpectError)
	}
	return native
}

// terraformStep treats plan, apply, and consistency reads as one fixture phase.
func terraformStep(t *testing.T, handler stepHTTP, label, hcl string, calls []httpfixture.Call, state *expectedState) resource.TestStep {
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
