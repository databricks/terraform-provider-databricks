package catalog

import (
	"context"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

const (
	workspaceHostAccountDomain   = "accounts.cloud.databricks.com"
	workspaceHostWorkspaceDomain = "replay-workspace-12345.cloud.databricks.com"
	workspaceHostLookupPath      = "/api/2.0/accounts/replay-account/workspaces/12345"
	workspaceHostVolumesPath     = "/api/2.1/unity-catalog/volumes"
	workspaceHostVolumePath      = workspaceHostVolumesPath + "/main.replay.host_volume"
	workspaceHostVolumeJSON      = `{"catalog_name": "main", "schema_name": "replay",
		"name": "host_volume", "full_name": "main.replay.host_volume",
		"volume_type": "MANAGED", "owner": "replay-admin@example.com",
		"storage_location": "s3://replay-metastore/volume-79",
		"metastore_id": "11111111-1111-1111-1111-111111111111",
		"volume_id": "00000000-0000-4000-8000-000000000079"}`
	workspaceHostWorkspaceJSON = `{"account_id": "replay-account", "workspace_id": 12345,
		"workspace_name": "replay-workspace-12345",
		"deployment_name": "replay-workspace-12345", "workspace_status": "RUNNING"}`
)

// workspaceHostResponse is a literal HTTP response served for one method and path.
type workspaceHostResponse struct {
	status int
	body   string
}

// workspaceHostVolumeRoutes answers the volume calls on both servers, so the call log
// alone shows which host received them.
var workspaceHostVolumeRoutes = map[string]workspaceHostResponse{
	"POST " + workspaceHostVolumesPath:  {http.StatusOK, workspaceHostVolumeJSON},
	"GET " + workspaceHostVolumePath:    {http.StatusOK, workspaceHostVolumeJSON},
	"DELETE " + workspaceHostVolumePath: {http.StatusOK, `{}`},
}

// workspaceHostLog records the calls both servers receive, in order.
type workspaceHostLog struct {
	mu    sync.Mutex
	calls []string
}

func (l *workspaceHostLog) add(call string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.calls = append(l.calls, call)
}

func (l *workspaceHostLog) snapshot() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.calls...)
}

// workspaceHostServer answers host metadata and its routes, logging every other call
// with the server name and X-Databricks-Workspace-Id header. Unknown routes get 404.
type workspaceHostServer struct {
	name     string
	metadata string
	routes   map[string]workspaceHostResponse
	log      *workspaceHostLog
}

func (s workspaceHostServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := workspaceHostResponse{http.StatusOK, s.metadata}
	if r.URL.Path != "/.well-known/databricks-config" {
		s.log.add(fmt.Sprintf("%s %s %s workspace_id=%s", s.name, r.Method, r.URL.Path,
			r.Header.Get("X-Databricks-Workspace-Id")))
		var ok bool
		response, ok = s.routes[r.Method+" "+r.URL.Path]
		if !ok {
			response = workspaceHostResponse{http.StatusNotFound,
				`{"error_code": "NOT_FOUND", "message": "no fixture"}`}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.status)
	_, _ = io.WriteString(w, response.body)
}

// workspaceHostTransport sends each known domain to its httptest server and rejects
// any other host, so no call leaves the test.
type workspaceHostTransport map[string]*url.URL

func (t workspaceHostTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	target, ok := t[req.URL.Host]
	if !ok {
		return nil, fmt.Errorf("request outside the test servers: %s", req.URL)
	}
	req = req.Clone(req.Context())
	req.URL.Scheme, req.URL.Host, req.Host = target.Scheme, target.Host, target.Host
	return http.DefaultTransport.RoundTrip(req)
}

// workspaceHostLoader keeps ambient profiles and environment variables out of the
// client configuration.
type workspaceHostLoader struct{}

func (workspaceHostLoader) Name() string { return "workspace host test" }

func (workspaceHostLoader) Configure(*config.Config) error { return nil }

// TestVolumeAccountProviderWorkspaceHost pins which host serves volume calls when an
// account-level provider routes through provider_config.workspace_id. Replay serves
// every host from one local server, and on a host outside a known DNS zone the SDK
// keeps the account host for the workspace client, so replay can't tell the two hosts
// apart. Here the provider uses an AWS account domain, and a transport sends the
// account domain and the workspace's deployment domain to two httptest servers.
func TestVolumeAccountProviderWorkspaceHost(t *testing.T) {
	tests := []struct {
		name   string
		lookup workspaceHostResponse
		// volumeHost names the server that receives the volume calls.
		volumeHost string
	}{
		{
			name:       "lookup_found_routes_to_workspace_host",
			lookup:     workspaceHostResponse{http.StatusOK, workspaceHostWorkspaceJSON},
			volumeHost: "workspace",
		},
		{
			name: "lookup_not_found_keeps_account_host",
			lookup: workspaceHostResponse{http.StatusNotFound, `{"error_code":
				"RESOURCE_DOES_NOT_EXIST", "message": "Workspace 12345 does not exist"}`},
			volumeHost: "account",
		},
		{
			name: "lookup_forbidden_keeps_account_host",
			lookup: workspaceHostResponse{http.StatusForbidden, `{"error_code":
				"PERMISSION_DENIED", "message": "Caller cannot read workspace 12345"}`},
			volumeHost: "account",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &workspaceHostLog{}
			accountRoutes := maps.Clone(workspaceHostVolumeRoutes)
			accountRoutes["GET "+workspaceHostLookupPath] = tt.lookup
			account := httptest.NewServer(workspaceHostServer{name: "account",
				metadata: `{"host_type": "account", "cloud": "AWS",
					"account_id": "replay-account"}`,
				routes: accountRoutes, log: log})
			defer account.Close()
			workspace := httptest.NewServer(workspaceHostServer{name: "workspace",
				metadata: `{"host_type": "workspace", "cloud": "AWS",
					"workspace_id": "12345"}`,
				routes: workspaceHostVolumeRoutes, log: log})
			defer workspace.Close()

			c := newWorkspaceHostClient(t, workspaceHostTransport{
				workspaceHostAccountDomain:   mustParseURL(t, account.URL),
				workspaceHostWorkspaceDomain: mustParseURL(t, workspace.URL),
			})
			diags := createReadDeleteVolume(t, c)

			volumeCall := tt.volumeHost + " %s workspace_id=12345"
			assert.Equal(t, []string{
				"account GET " + workspaceHostLookupPath + " workspace_id=",
				fmt.Sprintf(volumeCall, "POST "+workspaceHostVolumesPath),
				fmt.Sprintf(volumeCall, "GET "+workspaceHostVolumePath),
				fmt.Sprintf(volumeCall, "GET "+workspaceHostVolumePath),
				fmt.Sprintf(volumeCall, "DELETE "+workspaceHostVolumePath),
			}, log.snapshot())
			assert.False(t, diags.HasError(), "%v", diags)
		})
	}
}

func newWorkspaceHostClient(
	t *testing.T, transport http.RoundTripper,
) *common.DatabricksClient {
	t.Helper()
	inner, err := client.New(&config.Config{
		Host:          "https://" + workspaceHostAccountDomain,
		AccountID:     "replay-account",
		Token:         "fixture-token",
		AuthType:      "pat",
		Credentials:   config.PatCredentials{},
		Loaders:       []config.Loader{workspaceHostLoader{}},
		HTTPTransport: transport,
	})
	require.NoError(t, err)
	return &common.DatabricksClient{DatabricksClient: inner}
}

func mustParseURL(t *testing.T, raw string) *url.URL {
	t.Helper()
	u, err := url.Parse(raw)
	require.NoError(t, err)
	return u
}

// createReadDeleteVolume runs the provider's wrapped Create, Read and Delete on one
// client for a volume whose provider_config sets workspace_id 12345. It stops at the
// first operation that fails and returns its diagnostics.
func createReadDeleteVolume(t *testing.T, c *common.DatabricksClient) diag.Diagnostics {
	t.Helper()
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	r := ResourceVolume().ToResource()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"name":            "host_volume",
		"catalog_name":    "main",
		"schema_name":     "replay",
		"volume_type":     "MANAGED",
		"provider_config": []any{map[string]any{"workspace_id": "12345"}},
	})
	if diags := r.CreateContext(ctx, d, c); diags.HasError() {
		return diags
	}
	assert.Equal(t, "main.replay.host_volume", d.Id())
	assert.Equal(t, "12345", d.Get("provider_config.0.workspace_id"))
	if diags := r.ReadContext(ctx, d, c); diags.HasError() {
		return diags
	}
	return r.DeleteContext(ctx, d, c)
}
