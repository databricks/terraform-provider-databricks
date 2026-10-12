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
	"path/filepath"
	"strings"
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

// Run loads testdata/<test name>.yaml and runs its steps with the real muxed
// provider. Terraform must be installed; replay never downloads the CLI.
func Run(t *testing.T) {
	t.Helper()
	path := filepath.Join("testdata", strings.ReplaceAll(t.Name(), "/", "_")+".yaml")
	definition, err := loadDefinition(path)
	if err != nil {
		t.Fatalf("%s: %v", path, err)
	}
	run(t, definition)
}

func run(t *testing.T, definition *testDefinition) {
	t.Helper()
	configureTerraform(t)
	handler, err := httpfixture.NewHandler(definition.metadataCalls())
	if err != nil {
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
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: providerFactories(server.URL),
		Steps:                    steps,
		ErrorCheck: func(err error) error {
			// The testing framework attempts destroy after an interrupted lifecycle.
			cleanupErr := handler.StartStep(stepLabel(len(definition.Steps), "destroy"), definition.Destroy.Calls)
			return errors.Join(err, cleanupErr)
		},
	})
}

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
	// Optional framework refresh modes would change the required HTTP transcript.
	t.Setenv("TF_ACC_REFRESH_AFTER_APPLY", "")
	if err := os.Unsetenv("TF_ACC_REFRESH_AFTER_APPLY"); err != nil {
		t.Fatal(err)
	}
}

func testSteps(t *testing.T, definition *testDefinition, handler *httpfixture.Handler) []resource.TestStep {
	t.Helper()
	steps := make([]resource.TestStep, 0, len(definition.Steps)+1)
	for i, s := range definition.Steps {
		label := stepLabel(i, s.Name)
		steps = append(steps, terraformStep(t, handler, label, s.Terraform, s.Calls, s.State))
	}
	last := definition.Steps[len(definition.Steps)-1]
	empty := &expectedState{Resources: map[string]expectedResource{}}
	cleanup := terraformStep(t, handler, stepLabel(len(definition.Steps), "destroy"), last.Terraform, definition.Destroy.Calls, empty)
	cleanup.Destroy = true
	steps = append(steps, cleanup)
	return steps
}

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
			if err := handler.Verify(); err != nil {
				t.Error(err)
			}
		},
	}
	if state != nil {
		s.ConfigStateChecks = []statecheck.StateCheck{stateAssertion{label: label, expected: *state}}
	}
	return s
}

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
