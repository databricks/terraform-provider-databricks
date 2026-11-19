package mws_workspaces_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestNativeWorkspaces(t *testing.T) {
	for _, name := range []string{"TestMwsAccWorkspaces", "TestMwsAccWorkspacesTokenUpdate", "TestMwsAccAwsChangeToServicePrincipal"} {
		t.Run(name, func(t *testing.T) {
			terraformtest.RunNative(t, "native_testdata/"+name+".yaml", func(tc *resource.TestCase, customize func(*config.Config) error) {
				local := workspaceConfig(customize)
				tc.ProtoV6ProviderFactories = nil
				for i := range tc.Steps {
					selected := local
					if name == "TestMwsAccAwsChangeToServicePrincipal" && i >= 1 && i <= 3 {
						selected = servicePrincipalConfig(local)
					}
					tc.Steps[i].ProtoV6ProviderFactories = terraformtest.ProviderFactories(selected)
					if name == "TestMwsAccWorkspacesTokenUpdate" && i < 2 {
						tc.Steps[i].Check = checkTokenComment(t, local, []string{"test foo", "test bar"}[i])
					} else if name == "TestMwsAccWorkspaces" && i == 0 {
						tc.Steps[i].Check = checkTokenComment(t, local, "Test replay")
					}
				}
			})
		})
	}
}

func checkTokenComment(t *testing.T, customize func(*config.Config) error, expected string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		workspace := s.RootModule().Resources["databricks_mws_workspaces.this"]
		require.NotNil(t, workspace)
		cfg := &config.Config{}
		require.NoError(t, customize(cfg))
		cfg.Host = workspace.Primary.Attributes["workspace_url"]
		client, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
		if err != nil {
			return err
		}
		tokens, err := client.Tokens.ListAll(context.Background())
		if err != nil {
			return err
		}
		// Keep the original foo/bar checks without treating fixtures as backend proof.
		found := false
		for _, token := range tokens {
			if token.Comment == expected {
				found = true
			}
			if (expected == "test foo" && token.Comment == "test bar") || (expected == "test bar" && token.Comment == "test foo") {
				return fmt.Errorf("unexpected previous token comment %q", token.Comment)
			}
		}
		if !found {
			return fmt.Errorf("token comment %q absent", expected)
		}
		return nil
	}
}

// The test changes the acting identity with fixed local credentials. OAuth token
// exchange and account-to-workspace propagation belong to the SDK/backend;
// request headers verify which identity the provider uses at each step.
func servicePrincipalConfig(customize func(*config.Config) error) func(*config.Config) error {
	return func(c *config.Config) error {
		if err := customize(c); err != nil {
			return err
		}
		c.Token = "fixture-service-principal-token"
		c.ClientID = "11111111-1111-4111-8111-111111111111"
		return nil
	}
}

// Account responses deliberately point to a second host. Only this explicit
// test alias is routed back to the fixture server; the runner rejects every
// other external destination, including redirects. This is local to these Go
// tests and does not add multi-host behavior to YAML replay.
func workspaceConfig(customize func(*config.Config) error) func(*config.Config) error {
	return func(c *config.Config) error {
		if err := customize(c); err != nil {
			return err
		}
		target, err := url.Parse(c.Host)
		if err != nil {
			return err
		}
		c.HTTPTransport = workspaceTransport{target: target, transport: c.HTTPTransport}
		return nil
	}
}

type workspaceTransport struct {
	target    *url.URL
	transport http.RoundTripper
}

func (w workspaceTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.URL.Scheme == "http" && request.URL.Host == "workspace.replay.invalid" {
		request = request.Clone(request.Context())
		destination := *request.URL
		destination.Scheme, destination.Host = w.target.Scheme, w.target.Host
		request.URL, request.Host = &destination, w.target.Host
	}
	return w.transport.RoundTrip(request)
}
