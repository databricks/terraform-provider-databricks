package mlflow

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	sdkgoutility "github.com/databricks/terraform-provider-databricks/sdk-go-utility"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	mlflowModelWireGetPath      = "/api/2.0/mlflow/databricks/registered-models/get"
	mlflowModelWireMetadataPath = "/.well-known/databricks-config"
	mlflowModelWireLoginPage    = "<html><body>Sign in</body></html>"
)

// mlflowModelWireCalls records the requests of a test in the order in which they
// were sent.
type mlflowModelWireCalls struct {
	mu    sync.Mutex
	calls []string
}

func (l *mlflowModelWireCalls) add(call string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.calls = append(l.calls, call)
}

func (l *mlflowModelWireCalls) snapshot() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.calls...)
}

// mlflowModelWireLoader keeps profiles and environment variables out of the config
// of a test client.
type mlflowModelWireLoader struct{}

func (mlflowModelWireLoader) Name() string { return "mlflow model wire test" }

func (mlflowModelWireLoader) Configure(*config.Config) error { return nil }

// mlflowModelWireClient returns the provider client of cfg with fixed local
// credentials. The provider caches the workspace ID of the host metadata when it is
// configured, so no operation looks it up.
func mlflowModelWireClient(t *testing.T, cfg *config.Config) *common.DatabricksClient {
	t.Helper()
	cfg.Token, cfg.AuthType, cfg.Credentials = "fixture-token", "pat", config.PatCredentials{}
	cfg.Loaders = []config.Loader{mlflowModelWireLoader{}}
	inner, err := client.New(cfg)
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	c.SetCachedWorkspaceID(12345)
	return c
}

// mlflowModelWireContext is the context of an operation of databricks_mlflow_model,
// which names the resource in its diagnostics.
func mlflowModelWireContext() context.Context {
	return context.WithValue(context.Background(), common.ResourceName, "mlflow_model")
}

// mlflowModelWireInState returns the data of a model that a create left in the
// state.
func mlflowModelWireInState(t *testing.T, r *schema.Resource, name string) *schema.ResourceData {
	t.Helper()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"name":        name,
		"description": "Before the refresh",
		// A create puts the workspace ID here, so a refresh doesn't look it up.
		"provider_config": []any{map[string]any{"workspace_id": "12345"}},
	})
	d.SetId(name)
	require.NoError(t, d.Set("registered_model_id", "0a1b2c3d4e5f60718293a4b5c6d7ea51"))
	return d
}

// mlflowModelDomainTransport sends the requests for domain to the test server at
// target and records them, apart from host metadata. It rejects every other host.
// It returns each response with the request of the client, because the classic
// client reads the host of the final URL from it.
type mlflowModelDomainTransport struct {
	domain string
	target *url.URL
	calls  *mlflowModelWireCalls
}

func (t mlflowModelDomainTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Host != t.domain {
		return nil, fmt.Errorf("request outside the test server: %s", req.URL)
	}
	if req.URL.Path != mlflowModelWireMetadataPath {
		t.calls.add(req.Method + " " + req.URL.String())
	}
	out := req.Clone(req.Context())
	out.URL.Scheme, out.URL.Host, out.Host = t.target.Scheme, t.target.Host, t.target.Host
	resp, err := http.DefaultTransport.RoundTrip(out)
	if resp != nil {
		resp.Request = req
	}
	return resp, err
}

// mlflowModelRedirectServer answers host metadata, redirects the read of a model to
// location, and serves the login page for every other GET.
func mlflowModelRedirectServer(metadata, location string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == mlflowModelWireMetadataPath:
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, metadata)
		case r.Method == http.MethodGet && r.URL.Path == mlflowModelWireGetPath:
			w.Header().Set("Location", location)
			w.WriteHeader(http.StatusFound)
		case r.Method == http.MethodGet:
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = io.WriteString(w, mlflowModelWireLoginPage)
		default:
			http.NotFound(w, r)
		}
	})
}

// The private link redirect is detected by the final URL alone: its path has to be
// /login.html and its query has to hold error=private-link-validation-error. Another
// value of error or another path is not detected, so the 200 of the login page fails
// as a response that can't be decoded, with the text of each SDK. The error of a
// detected redirect names the cloud of the host of the final URL. A replay journey
// can't pin that for Azure and GCP, because it serves every call from 127.0.0.1, which
// falls back to AWS. Here the client uses a workspace domain of each cloud, and a
// transport sends that domain to a local server. The responses are synthetic: inferred
// from the private link redirect, not recorded.
func TestModelPrivateLinkRedirectDetectionEdges(t *testing.T) {
	mlflowModelOnBothPaths(t, testModelPrivateLinkRedirectDetectionEdges)
}

func testModelPrivateLinkRedirectDetectionEdges(t *testing.T) {
	const (
		name        = "private_link_edges_model"
		privateLink = "/login.html?error=private-link-validation-error"
	)
	tests := []struct {
		name     string
		domain   string
		cloud    string
		location string
		// diagnostic is the exact error of the refresh, and sdkGoDiagnostic the
		// one on the modular SDK path where it differs.
		diagnostic      string
		sdkGoDiagnostic string
	}{
		{
			name:     "another_value_of_error_is_not_detected",
			domain:   "dbc-a1b2c3d4-e5f6.cloud.databricks.com",
			cloud:    "AWS",
			location: "/login.html?error=session-expired",
			diagnostic: "cannot read mlflow model: " +
				"received HTML response instead of JSON",
			sdkGoDiagnostic: "cannot read mlflow model: " +
				"invalid character '<' looking for beginning of value",
		},
		{
			name:     "another_path_is_not_detected",
			domain:   "dbc-a1b2c3d4-e5f6.cloud.databricks.com",
			cloud:    "AWS",
			location: "/login?error=private-link-validation-error",
			diagnostic: "cannot read mlflow model: " +
				"received HTML response instead of JSON",
			sdkGoDiagnostic: "cannot read mlflow model: " +
				"invalid character '<' looking for beginning of value",
		},
		{
			name:     "azure_host_names_azure_private_link",
			domain:   "adb-1234567890123456.7.azuredatabricks.net",
			cloud:    "AZURE",
			location: privateLink,
			diagnostic: "cannot read mlflow model: The requested workspace has Azure " +
				"Private Link enabled and is not accessible from the current network. " +
				"Ensure that Azure Private Link is properly configured and that your " +
				"device has access to the Azure Private Link endpoint. For more " +
				"information, see https://learn.microsoft.com/en-us/azure/databricks/" +
				"security/network/classic/private-link-standard" +
				"#authentication-troubleshooting.",
		},
		{
			name:     "gcp_host_names_private_service_connect",
			domain:   "1234567890123456.7.gcp.databricks.com",
			cloud:    "GCP",
			location: privateLink,
			diagnostic: "cannot read mlflow model: The requested workspace has Private " +
				"Service Connect enabled and is not accessible from the current " +
				"network. Ensure that Private Service Connect is properly configured " +
				"and that your device has access to the GCP VPC endpoint. For more " +
				"information, see https://docs.gcp.databricks.com/en/security/network/" +
				"classic/private-service-connect.html.",
		},
		{
			// The cloud of the host metadata is not asked: a host that is no
			// Databricks domain falls back to AWS.
			name:     "unknown_host_of_an_azure_workspace_names_aws_privatelink",
			domain:   "workspace.example.com",
			cloud:    "AZURE",
			location: privateLink,
			diagnostic: "cannot read mlflow model: The requested workspace has AWS " +
				"PrivateLink enabled and is not accessible from the current network. " +
				"Ensure that AWS PrivateLink is properly configured and that your " +
				"device has access to the AWS VPC endpoint. For more information, see " +
				"https://docs.databricks.com/en/security/network/classic/" +
				"privatelink.html.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(mlflowModelRedirectServer(fmt.Sprintf(
				`{"host_type": "workspace", "cloud": %q, "workspace_id": "12345"}`,
				tt.cloud), tt.location))
			defer server.Close()
			target, err := url.Parse(server.URL)
			require.NoError(t, err)
			calls := &mlflowModelWireCalls{}
			host := "https://" + tt.domain
			c := mlflowModelWireClient(t, &config.Config{
				Host: host,
				HTTPTransport: mlflowModelDomainTransport{
					domain: tt.domain, target: target, calls: calls},
			})
			r := ResourceMlflowModel().ToResource()
			d := mlflowModelWireInState(t, r, name)
			before := d.State().Attributes

			diags := r.ReadContext(mlflowModelWireContext(), d, c)

			want := tt.diagnostic
			if tt.sdkGoDiagnostic != "" && sdkgoutility.Enabled(mlflowModelResourceType) {
				want = tt.sdkGoDiagnostic
			}
			assert.Equal(t, diag.Diagnostics{{Severity: diag.Error, Summary: want}}, diags)
			// The client follows the redirect and sends nothing after it.
			assert.Equal(t, []string{
				"GET " + host + mlflowModelWireGetPath + "?name=" + name,
				"GET " + host + tt.location,
			}, calls.snapshot())
			// The failed refresh keeps the ID of the model and every attribute.
			assert.Equal(t, name, d.Id())
			assert.Equal(t, before, d.State().Attributes)
		})
	}
}
