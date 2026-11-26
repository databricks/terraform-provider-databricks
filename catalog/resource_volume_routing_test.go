package catalog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	providerclient "github.com/databricks/terraform-provider-databricks/internal/providers/client"
)

const (
	routingWorkspace67890Domain = "replay-workspace-67890.cloud.databricks.com"
	routingVolumePath           = workspaceHostVolumesPath + "/main.replay.routing_volume"
	routingMePath               = "/api/2.0/preview/scim/v2/Me"
	routingWorkspace12345JSON   = `{"host_type": "workspace", "cloud": "AWS",
		"workspace_id": "12345"}`
	routingAccountJSON = `{"host_type": "account", "cloud": "AWS",
		"account_id": "replay-account"}`
)

// routingVolumeJSON is the volume's API representation with the given owner and comment.
func routingVolumeJSON(owner, comment string) string {
	return fmt.Sprintf(`{"catalog_name": "main", "schema_name": "replay",
		"name": "routing_volume", "full_name": "main.replay.routing_volume",
		"volume_type": "MANAGED", "owner": %q, "comment": %q,
		"storage_location": "s3://replay-metastore/volume-1920",
		"metastore_id": "11111111-1111-1111-1111-111111111111",
		"volume_id": "00000000-0000-4000-8000-000000001920"}`, owner, comment)
}

// routingResponse is a literal HTTP response with optional headers.
type routingResponse struct {
	status int
	header http.Header
	body   string
}

// routingServer answers host metadata and its routes by method and path, and 404 to any
// other call. Tests change routes between operations to model later API responses.
type routingServer struct {
	metadata string
	mu       sync.Mutex
	routes   map[string]routingResponse
}

func newRoutingServer(t *testing.T, metadata string,
	routes map[string]routingResponse) (*routingServer, *url.URL) {
	t.Helper()
	s := &routingServer{metadata: metadata, routes: maps.Clone(routes)}
	server := httptest.NewServer(s)
	t.Cleanup(server.Close)
	return s, mustParseURL(t, server.URL)
}

func (s *routingServer) set(route string, response routingResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.routes[route] = response
}

func (s *routingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := routingResponse{status: http.StatusOK, body: s.metadata}
	if r.URL.Path != "/.well-known/databricks-config" {
		s.mu.Lock()
		var ok bool
		response, ok = s.routes[r.Method+" "+r.URL.Path]
		s.mu.Unlock()
		if !ok {
			response = routingResponse{status: http.StatusNotFound,
				body: `{"error_code": "NOT_FOUND", "message": "no fixture"}`}
		}
	}
	maps.Copy(w.Header(), response.header)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.status)
	_, _ = io.WriteString(w, response.body)
}

// routingRecorder logs every call except host metadata with its destination host and
// X-Databricks-Workspace-Id values before passing it on, so a missing header shows as []
// and a call to an unknown host is logged too.
type routingRecorder struct {
	mu    sync.Mutex
	calls []string
	next  http.RoundTripper
}

func (r *routingRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path != "/.well-known/databricks-config" {
		r.mu.Lock()
		r.calls = append(r.calls, fmt.Sprintf("%s %s %s workspace_id=%q", req.URL.Host,
			req.Method, req.URL.Path, req.Header.Values("X-Databricks-Workspace-Id")))
		r.mu.Unlock()
	}
	return r.next.RoundTrip(req)
}

func (r *routingRecorder) snapshot() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.calls...)
}

// newRoutingClient configures a client for host as the provider does, including the
// workspace_id reconciliation with host metadata. Tests take a fresh client per
// operation, as each Terraform command starts a fresh provider.
func newRoutingClient(t *testing.T, host, accountID string,
	transport http.RoundTripper) *common.DatabricksClient {
	t.Helper()
	c, err := providerclient.PrepareDatabricksClient(context.Background(), &config.Config{
		Host:          "https://" + host,
		AccountID:     accountID,
		Token:         "fixture-token",
		AuthType:      "pat",
		Credentials:   config.PatCredentials{},
		Loaders:       []config.Loader{workspaceHostLoader{}},
		HTTPTransport: transport,
		// The provider's default retries until the context is cancelled.
		RetryTimeoutSeconds: 30,
	}, nil)
	require.NoError(t, err)
	return c
}

// routingVolume is a volume configuration with an optional provider_config.workspace_id.
type routingVolume struct {
	attributes  map[string]string
	workspaceID string
}

func (v routingVolume) resourceConfig() *terraform.ResourceConfig {
	raw := map[string]any{}
	for k, value := range v.attributes {
		raw[k] = value
	}
	if v.workspaceID != "" {
		raw["provider_config"] = []any{map[string]any{"workspace_id": v.workspaceID}}
	}
	return terraform.NewResourceConfigRaw(raw)
}

// rawConfig is what Terraform's gRPC calls carry, so provider_config is read from the
// configuration as it is outside tests.
func (v routingVolume) rawConfig(t *testing.T, r *schema.Resource) cty.Value {
	t.Helper()
	values := map[string]cty.Value{}
	for k, value := range v.attributes {
		values[k] = cty.StringVal(value)
	}
	if v.workspaceID != "" {
		values["provider_config"] = cty.ListVal([]cty.Value{cty.ObjectVal(
			map[string]cty.Value{"workspace_id": cty.StringVal(v.workspaceID)})})
	}
	raw, err := r.CoreConfigSchema().CoerceValue(cty.ObjectVal(values))
	require.NoError(t, err)
	return raw
}

func routingVolumeAttributes(extra map[string]string) map[string]string {
	attributes := map[string]string{"name": "routing_volume", "catalog_name": "main",
		"schema_name": "replay", "volume_type": "MANAGED"}
	maps.Copy(attributes, extra)
	return attributes
}

// routingApplyVolume plans v against state and applies the plan with c, as one terraform
// apply does, through sdk-go if sdkGo is set, and returns the new state. A nil state
// creates the volume.
func routingApplyVolume(
	t *testing.T,
	c *common.DatabricksClient,
	state *terraform.InstanceState,
	v routingVolume,
	sdkGo bool,
) (*terraform.InstanceState, diag.Diagnostics) {
	t.Helper()
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	r := volumeResource(sdkGo)
	raw := v.rawConfig(t, r)
	if state == nil {
		state = &terraform.InstanceState{}
	}
	state.RawConfig = raw
	diff, err := r.Diff(ctx, state, v.resourceConfig(), c)
	require.NoError(t, err)
	require.NotNil(t, diff)
	diff.RawConfig = raw
	return r.Apply(ctx, state, diff, c)
}

func routingRefreshVolume(c *common.DatabricksClient, state *terraform.InstanceState,
	sdkGo bool) (*terraform.InstanceState, diag.Diagnostics) {
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	return volumeResource(sdkGo).RefreshWithoutUpgrade(ctx, state, c)
}

func routingDestroyVolume(c *common.DatabricksClient, state *terraform.InstanceState,
	sdkGo bool) diag.Diagnostics {
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	_, diags := volumeResource(sdkGo).Apply(ctx, state,
		&terraform.InstanceDiff{Destroy: true}, c)
	return diags
}

// TestVolumeAccountProviderPatchesGoToWorkspaceHost pins that an account-level provider
// sends the owner PATCH after create, and the owner and fields PATCHes of an update, to
// the workspace's deployment host found by the account lookup, as it does POST, GET and
// DELETE in TestVolumeAccountProviderWorkspaceHost. Replay serves every host from one
// local server, so it can't tell the hosts apart.
func TestVolumeAccountProviderPatchesGoToWorkspaceHost(t *testing.T) {
	forEachVolumeClient(t, testVolumeAccountProviderPatchesGoToWorkspaceHost)
}

func testVolumeAccountProviderPatchesGoToWorkspaceHost(t *testing.T, sdkGo bool) {
	routes := map[string]routingResponse{
		"POST " + workspaceHostVolumesPath: {http.StatusOK,
			nil, routingVolumeJSON("replay-admin@example.com", "")},
		"PATCH " + routingVolumePath: {http.StatusOK,
			nil, routingVolumeJSON("data-engineers", "")},
		"GET " + routingVolumePath: {http.StatusOK,
			nil, routingVolumeJSON("data-engineers", "")},
	}
	accountRoutes := maps.Clone(routes)
	accountRoutes["GET "+workspaceHostLookupPath] = routingResponse{http.StatusOK,
		nil, workspaceHostWorkspaceJSON}
	_, accountURL := newRoutingServer(t, routingAccountJSON, accountRoutes)
	workspace, workspaceURL := newRoutingServer(t, routingWorkspace12345JSON, routes)
	recorder := &routingRecorder{next: workspaceHostTransport{
		workspaceHostAccountDomain:   accountURL,
		workspaceHostWorkspaceDomain: workspaceURL,
	}}
	newClient := func() *common.DatabricksClient {
		return newRoutingClient(t, workspaceHostAccountDomain, "replay-account", recorder)
	}

	state, diags := routingApplyVolume(t, newClient(), nil, routingVolume{
		attributes:  routingVolumeAttributes(map[string]string{"owner": "data-engineers"}),
		workspaceID: "12345",
	}, sdkGo)
	require.False(t, diags.HasError(), "%v", diags)
	workspace.set("GET "+routingVolumePath, routingResponse{http.StatusOK,
		nil, routingVolumeJSON("data-stewards", "routed comment")})
	_, diags = routingApplyVolume(t, newClient(), state, routingVolume{
		attributes: routingVolumeAttributes(map[string]string{"owner": "data-stewards",
			"comment": "routed comment"}),
		workspaceID: "12345",
	}, sdkGo)
	require.False(t, diags.HasError(), "%v", diags)

	lookup := workspaceHostAccountDomain + " GET " + workspaceHostLookupPath +
		" workspace_id=[]"
	volumeCall := workspaceHostWorkspaceDomain + ` %s workspace_id=["12345"]`
	assert.Equal(t, []string{
		lookup,
		fmt.Sprintf(volumeCall, "POST "+workspaceHostVolumesPath),
		fmt.Sprintf(volumeCall, "PATCH "+routingVolumePath),
		fmt.Sprintf(volumeCall, "GET "+routingVolumePath),
		lookup,
		fmt.Sprintf(volumeCall, "PATCH "+routingVolumePath),
		fmt.Sprintf(volumeCall, "PATCH "+routingVolumePath),
		fmt.Sprintf(volumeCall, "GET "+routingVolumePath),
	}, recorder.snapshot())
}

// TestVolumeAccountProviderLookupWithoutDeploymentName pins an account lookup answered
// 200 without deployment_name on an AWS account domain. The SDK builds the workspace
// host from the empty deployment name, so the volume POST goes to
// https://.cloud.databricks.com. The test transport rejects that host, standing in for
// the DNS failure a real run gets, and create fails with the transport error.
func TestVolumeAccountProviderLookupWithoutDeploymentName(t *testing.T) {
	forEachVolumeClient(t, testVolumeAccountProviderLookupWithoutDeploymentName)
}

func testVolumeAccountProviderLookupWithoutDeploymentName(t *testing.T, sdkGo bool) {
	_, accountURL := newRoutingServer(t, routingAccountJSON, map[string]routingResponse{
		"GET " + workspaceHostLookupPath: {http.StatusOK, nil, `{
			"account_id": "replay-account", "workspace_id": 12345,
			"workspace_name": "replay-workspace-12345", "workspace_status": "RUNNING"}`},
	})
	recorder := &routingRecorder{next: workspaceHostTransport{
		workspaceHostAccountDomain: accountURL,
	}}
	c := newRoutingClient(t, workspaceHostAccountDomain, "replay-account", recorder)

	_, diags := routingApplyVolume(t, c, nil, routingVolume{
		attributes:  routingVolumeAttributes(nil),
		workspaceID: "12345",
	}, sdkGo)

	assert.Equal(t, []string{
		workspaceHostAccountDomain + " GET " + workspaceHostLookupPath + " workspace_id=[]",
		`.cloud.databricks.com POST ` + workspaceHostVolumesPath + ` workspace_id=["12345"]`,
	}, recorder.snapshot())
	require.Len(t, diags, 1)
	assert.Equal(t, `cannot create volume: Post "https://.cloud.databricks.com`+
		workspaceHostVolumesPath+`": request outside the test servers: `+
		`https://.cloud.databricks.com`+workspaceHostVolumesPath, diags[0].Summary)
}

// TestVolumeHostChangeFailsRefreshBeforeRead pins a refresh after the provider host
// changes from workspace 12345's host to workspace 67890's while state records 12345.
// Host metadata seeds the new provider's workspace ID, so the refresh fails with a
// workspace_id mismatch without any call to the new host.
func TestVolumeHostChangeFailsRefreshBeforeRead(t *testing.T) {
	forEachVolumeClient(t, testVolumeHostChangeFailsRefreshBeforeRead)
}

func testVolumeHostChangeFailsRefreshBeforeRead(t *testing.T, sdkGo bool) {
	routes := map[string]routingResponse{
		"POST " + workspaceHostVolumesPath: {http.StatusOK,
			nil, routingVolumeJSON("replay-admin@example.com", "")},
		"GET " + routingVolumePath: {http.StatusOK,
			nil, routingVolumeJSON("replay-admin@example.com", "")},
	}
	_, hostA := newRoutingServer(t, routingWorkspace12345JSON, routes)
	_, hostB := newRoutingServer(t, `{"host_type": "workspace", "cloud": "AWS",
		"workspace_id": "67890"}`, routes)
	recorder := &routingRecorder{next: workspaceHostTransport{
		workspaceHostWorkspaceDomain: hostA,
		routingWorkspace67890Domain:  hostB,
	}}

	state, diags := routingApplyVolume(t,
		newRoutingClient(t, workspaceHostWorkspaceDomain, "", recorder), nil,
		routingVolume{attributes: routingVolumeAttributes(nil)}, sdkGo)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "12345", state.Attributes["provider_config.0.workspace_id"])
	_, diags = routingRefreshVolume(
		newRoutingClient(t, routingWorkspace67890Domain, "", recorder), state, sdkGo)

	volumeCall := workspaceHostWorkspaceDomain + ` %s workspace_id=["12345"]`
	assert.Equal(t, []string{
		fmt.Sprintf(volumeCall, "POST "+workspaceHostVolumesPath),
		fmt.Sprintf(volumeCall, "GET "+routingVolumePath),
	}, recorder.snapshot())
	require.Len(t, diags, 1)
	assert.Equal(t, "cannot read volume: failed to validate workspace_id: workspace_id "+
		"mismatch: provider is configured for workspace 67890 but got 12345 in "+
		"provider_config. please check the workspace_id provided in provider_config",
		diags[0].Summary)
}

// TestVolumeWithoutWorkspaceIDSendsNoWorkspaceHeader pins that a workspace-level
// provider without any known workspace ID, from configuration or host metadata, sends
// no X-Databricks-Workspace-Id header on any call, including after SCIM /Me resolved
// the ID into state. Replay ignores unlisted headers, so it can't assert an absence.
// Each operation uses a fresh client, as each Terraform command does.
func TestVolumeWithoutWorkspaceIDSendsNoWorkspaceHeader(t *testing.T) {
	forEachVolumeClient(t, testVolumeWithoutWorkspaceIDSendsNoWorkspaceHeader)
}

func testVolumeWithoutWorkspaceIDSendsNoWorkspaceHeader(t *testing.T, sdkGo bool) {
	workspace, workspaceURL := newRoutingServer(t, `{"host_type": "workspace",
		"cloud": "AWS"}`, map[string]routingResponse{
		"POST " + workspaceHostVolumesPath: {http.StatusOK,
			nil, routingVolumeJSON("replay-admin@example.com", "")},
		"PATCH " + routingVolumePath: {http.StatusOK,
			nil, routingVolumeJSON("data-engineers", "")},
		"GET " + routingVolumePath: {http.StatusOK,
			nil, routingVolumeJSON("data-engineers", "")},
		"DELETE " + routingVolumePath: {http.StatusOK, nil, `{}`},
		"GET " + routingMePath: {http.StatusOK,
			http.Header{"X-Databricks-Org-Id": {"12345"}}, `{"id": "1001",
			"userName": "replay-admin@example.com"}`},
	})
	recorder := &routingRecorder{next: workspaceHostTransport{
		workspaceHostWorkspaceDomain: workspaceURL,
	}}
	newClient := func() *common.DatabricksClient {
		return newRoutingClient(t, workspaceHostWorkspaceDomain, "", recorder)
	}

	state, diags := routingApplyVolume(t, newClient(), nil, routingVolume{
		attributes: routingVolumeAttributes(map[string]string{"owner": "data-engineers"}),
	}, sdkGo)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "12345", state.Attributes["provider_config.0.workspace_id"])
	workspace.set("GET "+routingVolumePath, routingResponse{http.StatusOK,
		nil, routingVolumeJSON("data-stewards", "routed comment")})
	state, diags = routingApplyVolume(t, newClient(), state, routingVolume{
		attributes: routingVolumeAttributes(map[string]string{"owner": "data-stewards",
			"comment": "routed comment"}),
	}, sdkGo)
	require.False(t, diags.HasError(), "%v", diags)
	diags = routingDestroyVolume(newClient(), state, sdkGo)
	require.False(t, diags.HasError(), "%v", diags)

	call := workspaceHostWorkspaceDomain + " %s workspace_id=[]"
	assert.Equal(t, []string{
		fmt.Sprintf(call, "POST "+workspaceHostVolumesPath),
		fmt.Sprintf(call, "PATCH "+routingVolumePath),
		fmt.Sprintf(call, "GET "+routingVolumePath),
		fmt.Sprintf(call, "GET "+routingMePath),
		fmt.Sprintf(call, "GET "+routingMePath),
		fmt.Sprintf(call, "PATCH "+routingVolumePath),
		fmt.Sprintf(call, "PATCH "+routingVolumePath),
		fmt.Sprintf(call, "GET "+routingVolumePath),
		fmt.Sprintf(call, "GET "+routingMePath),
		fmt.Sprintf(call, "DELETE "+routingVolumePath),
	}, recorder.snapshot())
}

// TestVolumeWorkspaceIDChangeToAnotherWorkspaceFailsUpdate pins a change of
// provider_config.workspace_id to a workspace other than that of a workspace-level
// provider. The plan accepts it, and Update fails on the mismatch before any PATCH. The
// failed update still records the new workspace ID in state, so the next refresh fails
// on the same mismatch. A replay journey can't pin this, as its destroy would fail too.
func TestVolumeWorkspaceIDChangeToAnotherWorkspaceFailsUpdate(t *testing.T) {
	forEachVolumeClient(t, testVolumeWorkspaceIDChangeToAnotherWorkspaceFailsUpdate)
}

func testVolumeWorkspaceIDChangeToAnotherWorkspaceFailsUpdate(t *testing.T, sdkGo bool) {
	_, workspaceURL := newRoutingServer(t, routingWorkspace12345JSON,
		map[string]routingResponse{
			"POST " + workspaceHostVolumesPath: {http.StatusOK,
				nil, routingVolumeJSON("replay-admin@example.com", "")},
			"GET " + routingVolumePath: {http.StatusOK,
				nil, routingVolumeJSON("replay-admin@example.com", "")},
		})
	recorder := &routingRecorder{next: workspaceHostTransport{
		workspaceHostWorkspaceDomain: workspaceURL,
	}}
	newClient := func() *common.DatabricksClient {
		return newRoutingClient(t, workspaceHostWorkspaceDomain, "", recorder)
	}

	state, diags := routingApplyVolume(t, newClient(), nil, routingVolume{
		attributes: routingVolumeAttributes(nil), workspaceID: "12345"}, sdkGo)
	require.False(t, diags.HasError(), "%v", diags)
	state, diags = routingApplyVolume(t, newClient(), state, routingVolume{
		attributes: routingVolumeAttributes(nil), workspaceID: "67890"}, sdkGo)

	mismatch := "failed to validate workspace_id: workspace_id mismatch: provider is " +
		"configured for workspace 12345 but got 67890 in provider_config. please check " +
		"the workspace_id provided in provider_config"
	require.Len(t, diags, 1)
	assert.Equal(t, "cannot update volume: "+mismatch, diags[0].Summary)
	require.NotNil(t, state)
	assert.Equal(t, "67890", state.Attributes["provider_config.0.workspace_id"])
	_, diags = routingRefreshVolume(newClient(), state, sdkGo)
	require.Len(t, diags, 1)
	assert.Equal(t, "cannot read volume: "+mismatch, diags[0].Summary)
	volumeCall := workspaceHostWorkspaceDomain + ` %s workspace_id=["12345"]`
	assert.Equal(t, []string{
		fmt.Sprintf(volumeCall, "POST "+workspaceHostVolumesPath),
		fmt.Sprintf(volumeCall, "GET "+routingVolumePath),
	}, recorder.snapshot())
}

// brokenAfterFirstLoader loads nothing for the provider's configuration and fails every
// later load, as a profile does that becomes unreadable after the provider starts.
type brokenAfterFirstLoader struct {
	loads *atomic.Int32
}

func (brokenAfterFirstLoader) Name() string { return "profile broken after start" }

func (l brokenAfterFirstLoader) Configure(*config.Config) error {
	if l.loads.Add(1) > 1 {
		return errors.New("profile replay is no longer readable")
	}
	return nil
}

// An account-level provider whose account lookup of the workspace is forbidden falls back
// to a client for the workspace on the same host, which loads the provider's config
// again. When that load fails, the volume call fails before any request to the
// workspace.
func TestVolumeAccountProviderFailsWhenTheWorkspaceClientCantBeBuilt(t *testing.T) {
	forEachVolumeClient(t, testVolumeAccountProviderFailsWhenTheWorkspaceClientCantBeBuilt)
}

func testVolumeAccountProviderFailsWhenTheWorkspaceClientCantBeBuilt(t *testing.T,
	sdkGo bool) {
	_, accountURL := newRoutingServer(t, routingAccountJSON, map[string]routingResponse{
		"GET " + workspaceHostLookupPath: {http.StatusForbidden, nil,
			`{"error_code": "PERMISSION_DENIED", "message": "no account access"}`},
	})
	recorder := &routingRecorder{next: workspaceHostTransport{
		workspaceHostAccountDomain: accountURL,
	}}
	c, err := providerclient.PrepareDatabricksClient(context.Background(), &config.Config{
		Host:                "https://" + workspaceHostAccountDomain,
		AccountID:           "replay-account",
		Token:               "fixture-token",
		AuthType:            "pat",
		Credentials:         config.PatCredentials{},
		Loaders:             []config.Loader{brokenAfterFirstLoader{&atomic.Int32{}}},
		HTTPTransport:       recorder,
		RetryTimeoutSeconds: 30,
	}, nil)
	require.NoError(t, err)

	_, diags := routingApplyVolume(t, c, nil, routingVolume{
		attributes:  routingVolumeAttributes(nil),
		workspaceID: "12345",
	}, sdkGo)

	assert.Equal(t, []string{
		workspaceHostAccountDomain + " GET " + workspaceHostLookupPath + " workspace_id=[]",
	}, recorder.snapshot())
	require.Len(t, diags, 1)
	assert.Regexp(t, `^cannot create volume: failed to get workspace client with `+
		`workspace_id 12345: resolve: profile replay is no longer readable\.`,
		diags[0].Summary)
}
