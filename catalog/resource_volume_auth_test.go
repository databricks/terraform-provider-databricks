package catalog

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
	providerclient "github.com/databricks/terraform-provider-databricks/internal/providers/client"
)

// These tests pin the credential headers and token requests of databricks_volume
// refreshes for credential types other than the PAT that request_headers.yaml pins.
// YAML replay can't express them: its client customizer forces PAT authentication,
// and its transport rejects the Azure AD, Azure instance metadata and Google token
// hosts. Here one httptest server answers every host, through a transport that keeps
// each request's original host, and the provider's PrepareDatabricksClient builds the
// client with the SDK's default credential strategies.

const (
	authVolumeID   = "main.replay.auth_volume"
	authVolumePath = workspaceHostVolumesPath + "/" + authVolumeID
	authVolumeJSON = `{"catalog_name": "main", "schema_name": "replay",
		"name": "auth_volume", "full_name": "main.replay.auth_volume",
		"volume_type": "MANAGED", "owner": "replay-admin@example.com",
		"storage_location": "s3://replay-metastore/volume-199",
		"metastore_id": "11111111-1111-1111-1111-111111111111",
		"volume_id": "00000000-0000-4000-8000-000000000199"}`
	authAWSHost    = workspaceHostWorkspaceDomain
	authAzureHost  = "adb-12345.6.azuredatabricks.net"
	authGCPHost    = "12345.6.gcp.databricks.com"
	authAADToken   = "login.microsoftonline.com/replay-tenant/oauth2/token"
	authMSIToken   = "169.254.169.254/metadata/identity/oauth2/token"
	authGoogleURL  = "oauth2.googleapis.com/token"
	authAzureAppID = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"
	authAzureMgmt  = "https://management.core.windows.net/"
	authAzureWSID  = "/subscriptions/replay-subscription/resourceGroups/replay-group" +
		"/providers/Microsoft.Databricks/workspaces/replay-workspace"
	authGCPAccount        = "replay-sa@replay-project.iam.gserviceaccount.com"
	authOIDCDiscoveryPath = "/oidc/.well-known/oauth-authorization-server"
	authM2MRejected       = `{"error": "invalid_client",` +
		` "error_description": "Bad secret"}`
	authAADRejected = `{"error": "invalid_client", "error_description": ` +
		`"AADSTS7000215: Invalid client secret provided. Trace ID: 1"}`
	authGoogleRejected = `{"error": "invalid_grant", "error_description": "Bad JWT"}`
)

// authIDToken is a Google ID token for the GCP workspace. The Google client reads its
// expiry from the claims, so they must decode, but nothing checks its signature.
var authIDToken = base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"RS256"}`)) + "." +
	base64.RawURLEncoding.EncodeToString([]byte(
		`{"aud":"https://`+authGCPHost+`","exp":4102444800}`)) + ".signature"

// authCredentialHeaders are the headers each recorded request keeps.
var authCredentialHeaders = []string{"Authorization", "Metadata",
	"X-Databricks-Azure-Sp-Management-Token", "X-Databricks-Azure-Workspace-Resource-Id",
	"X-Databricks-Gcp-Sa-Access-Token"}

// authRequest is a request as the server received it. headers holds the credential
// headers it carried, joined with ", " if repeated, so a missing key means the request
// had no such header.
type authRequest struct {
	method  string
	url     string
	body    string
	headers map[string]string
}

// authExchange is one request the server expects, in order, and its response.
type authExchange struct {
	request  authRequest
	status   int
	response string
}

// authServer answers host metadata on every host, and every other request with the
// next scripted response, recording it. A request past the script gets 404.
type authServer struct {
	metadata string
	script   []authExchange
	mu       sync.Mutex
	requests []authRequest
}

func (s *authServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/.well-known/databricks-config" {
		_, _ = io.WriteString(w, s.metadata)
		return
	}
	body, _ := io.ReadAll(r.Body)
	request := authRequest{method: r.Method, url: r.Host + r.URL.Path,
		body: readableAssertion(string(body)), headers: map[string]string{}}
	if r.URL.RawQuery != "" {
		request.url += "?" + r.URL.RawQuery
	}
	for _, name := range authCredentialHeaders {
		if values := r.Header.Values(name); len(values) > 0 {
			request.headers[name] = strings.Join(values, ", ")
		}
	}
	s.mu.Lock()
	next := len(s.requests)
	s.requests = append(s.requests, request)
	s.mu.Unlock()
	status := http.StatusNotFound
	response := `{"error_code": "NOT_FOUND", "message": "no fixture"}`
	if next < len(s.script) {
		status, response = s.script[next].status, s.script[next].response
	}
	w.WriteHeader(status)
	_, _ = io.WriteString(w, response)
}

func (s *authServer) expected() []authRequest {
	var requests []authRequest
	for _, exchange := range s.script {
		requests = append(requests, exchange.request)
	}
	return requests
}

func (s *authServer) received() []authRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	return slices.Clone(s.requests)
}

// readableAssertion replaces the signed JWT of a Google token request with its claims,
// without the iat and exp times, so the body is stable. Other bodies are unchanged.
func readableAssertion(body string) string {
	form, err := url.ParseQuery(body)
	if err != nil || form.Get("assertion") == "" {
		return body
	}
	parts := strings.Split(form.Get("assertion"), ".")
	if len(parts) != 3 {
		return body
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	var claims map[string]any
	if err != nil || json.Unmarshal(payload, &claims) != nil {
		return body
	}
	delete(claims, "iat")
	delete(claims, "exp")
	readable, _ := json.Marshal(claims)
	return fmt.Sprintf("assertion=%s&grant_type=%s", readable, form.Get("grant_type"))
}

// authTransport sends every request to the test server and keeps its original host
// in the Host header, so no request leaves the test.
type authTransport struct {
	server *url.URL
}

func (t authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Host = req.URL.Host
	req.URL.Scheme, req.URL.Host = t.server.Scheme, t.server.Host
	return http.DefaultTransport.RoundTrip(req)
}

// newAuthClient starts s and configures a client for cfg as the provider does,
// including its deprecated auth type mapping. Ambient profiles and environment
// variables are kept out of cfg.
func newAuthClient(t *testing.T, s *authServer,
	cfg *config.Config) (*common.DatabricksClient, error) {
	t.Helper()
	server := httptest.NewServer(s)
	t.Cleanup(server.Close)
	cfg.Loaders = []config.Loader{workspaceHostLoader{}}
	cfg.HTTPTransport = authTransport{server: mustParseURL(t, server.URL)}
	return providerclient.PrepareDatabricksClient(context.Background(), cfg, nil)
}

// isolateDefaultChain keeps the default chain's strategies that read the process
// environment from finding credentials: no Databricks CLI on PATH, no Azure DevOps
// token and no OIDC token variable.
func isolateDefaultChain(t *testing.T) {
	t.Setenv("PATH", "")
	t.Setenv("SYSTEM_ACCESSTOKEN", "")
	t.Setenv("DATABRICKS_OIDC_TOKEN", "")
}

// authVolumeState is the state of the volume after it was created on workspace 12345.
func authVolumeState() *terraform.InstanceState {
	return &terraform.InstanceState{ID: authVolumeID, Attributes: map[string]string{
		"id":                             authVolumeID,
		"name":                           "auth_volume",
		"catalog_name":                   "main",
		"schema_name":                    "replay",
		"volume_type":                    "MANAGED",
		"owner":                          "replay-admin@example.com",
		"storage_location":               "s3://replay-metastore/volume-199",
		"volume_path":                    "/Volumes/main/replay/auth_volume",
		"provider_config.#":              "1",
		"provider_config.0.workspace_id": "12345",
	}}
}

// authRefresh is the outcome of a refresh of the volume, and the next plan of its
// configuration from the refreshed state.
type authRefresh struct {
	client *common.DatabricksClient
	state  *terraform.InstanceState
	diags  diag.Diagnostics
	plan   *terraform.InstanceDiff
}

// refreshAuthVolume refreshes the volume with a client for cfg, as terraform plan does,
// checks that the requests match the script and plans the volume's configuration.
func refreshAuthVolume(t *testing.T, s *authServer, cfg *config.Config) authRefresh {
	t.Helper()
	c, err := newAuthClient(t, s, cfg)
	require.NoError(t, err)
	state, diags := routingRefreshVolume(c, false, authVolumeState())
	assert.Equal(t, s.expected(), s.received())

	r := ResourceVolume().ToResource()
	v := routingVolume{attributes: routingVolumeAttributes(map[string]string{
		"name": "auth_volume"})}
	planned := state
	if planned == nil {
		planned = &terraform.InstanceState{}
	}
	planned.RawConfig = v.rawConfig(t, r)
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	plan, err := r.Diff(ctx, planned, v.resourceConfig(), c)
	require.NoError(t, err)
	return authRefresh{client: c, state: state, diags: diags, plan: plan}
}

// assertKept checks that the refresh left state unchanged and the next plan empty.
func (a authRefresh) assertKept(t *testing.T) {
	t.Helper()
	require.NotNil(t, a.state)
	assert.Equal(t, authVolumeState().Attributes, a.state.Attributes)
	assert.Nil(t, a.plan, "the next plan must be empty")
}

func (a authRefresh) assertSucceeded(t *testing.T) {
	t.Helper()
	assert.False(t, a.diags.HasError(), "%v", a.diags)
	a.assertKept(t)
}

func (a authRefresh) assertFailed(t *testing.T, summary string) {
	t.Helper()
	require.Len(t, a.diags, 1)
	assert.Equal(t, summary, a.diags[0].Summary)
	a.assertKept(t)
}

// authVolumeRead is the volume GET with the given credential headers, answered 200.
func authVolumeRead(host string, headers map[string]string) authExchange {
	return authExchange{authRequest{method: "GET", url: host + authVolumePath,
		headers: headers}, http.StatusOK, authVolumeJSON}
}

// authBearerRead is the volume GET on the AWS workspace with a bearer token.
func authBearerRead(token string) authExchange {
	return authVolumeRead(authAWSHost,
		map[string]string{"Authorization": "Bearer " + token})
}

func basicAuthorization(user, password string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
}

func authBearerToken(token string, expiresIn int) string {
	return fmt.Sprintf(`{"access_token": %q, "token_type": "Bearer", "expires_in": %d}`,
		token, expiresIn)
}

func authAWSMetadata() string {
	return `{"host_type": "workspace", "cloud": "AWS", "workspace_id": "12345",
		"oidc_endpoint": "https://` + authAWSHost + `/oidc"}`
}

func authCloudMetadata(cloud string) string {
	return fmt.Sprintf(`{"host_type": "workspace", "cloud": %q, "workspace_id": "12345"}`,
		cloud)
}

// authOIDCDiscovery is the unauthenticated OIDC discovery GET on the AWS workspace.
var authOIDCDiscovery = authExchange{
	authRequest{method: "GET", url: authAWSHost + authOIDCDiscoveryPath,
		headers: map[string]string{}},
	http.StatusOK, `{"authorization_endpoint": "https://` + authAWSHost +
		`/oidc/v1/authorize", "token_endpoint": "https://` + authAWSHost +
		`/oidc/v1/token"}`,
}

// authM2MToken is the client_credentials request of service principal replay-sp.
func authM2MToken(status int, response string) authExchange {
	return authExchange{authRequest{method: "POST", url: authAWSHost + "/oidc/v1/token",
		body: "grant_type=client_credentials&scope=all-apis",
		headers: map[string]string{
			"Authorization": basicAuthorization("replay-sp", "replay-secret")}},
		status, response}
}

func authM2MConfig(authType string) *config.Config {
	return &config.Config{Host: "https://" + authAWSHost, AuthType: authType,
		ClientID: "replay-sp", ClientSecret: "replay-secret"}
}

// authFailedDefaultChain is one run of the default chain in which OAuth M2M fails its
// dry run. The github-oidc, env-oidc and file-oidc strategies then look up the token
// endpoint before they find no ID token, and no other strategy is configured.
var authFailedDefaultChain = []authExchange{
	authOIDCDiscovery,
	authM2MToken(http.StatusUnauthorized, authM2MRejected),
	authOIDCDiscovery,
	authOIDCDiscovery,
	authOIDCDiscovery,
}

func TestVolumeAuthBasicSendsBasicAuthorization(t *testing.T) {
	s := &authServer{metadata: authAWSMetadata(), script: []authExchange{
		authVolumeRead(authAWSHost, map[string]string{
			"Authorization": basicAuthorization("replay-user", "replay-password")}),
	}}

	refreshAuthVolume(t, s, &config.Config{Host: "https://" + authAWSHost,
		Username: "replay-user", Password: "replay-password"}).assertSucceeded(t)
}

func TestVolumeAuthM2MSendsBearerToken(t *testing.T) {
	// The default chain fetches the token in its dry run, and the GET reuses it.
	s := &authServer{metadata: authAWSMetadata(), script: []authExchange{
		authOIDCDiscovery,
		authM2MToken(http.StatusOK, authBearerToken("m2m-token", 3600)),
		authBearerRead("m2m-token"),
	}}

	refreshAuthVolume(t, s, authM2MConfig("")).assertSucceeded(t)
}

func TestVolumeAuthM2MTokenEndpointFailure(t *testing.T) {
	tests := []struct {
		name     string
		authType string
		script   []authExchange
		summary  string
	}{
		{
			// Without a dry run, the token fails in the GET's request visitor, which
			// doesn't retry.
			name:     "explicit_auth_type",
			authType: "oauth-m2m",
			script: []authExchange{
				authOIDCDiscovery,
				authM2MToken(http.StatusUnauthorized, authM2MRejected),
			},
			summary: `cannot read volume: failed during request visitor: error ` +
				`getting token: Post "https://` + authAWSHost + `/oidc/v1/token": ` +
				authM2MRejected,
		},
		{
			// A failed chain isn't kept, so it runs twice: when the provider builds the
			// workspace client, whose serving data plane asks for a token source, and in
			// the GET's request visitor.
			name:   "default_chain",
			script: slices.Concat(authFailedDefaultChain, authFailedDefaultChain),
			summary: "cannot read volume: failed during request visitor: default auth: " +
				"cannot configure default credentials, please check https://docs." +
				"databricks.com/en/dev-tools/auth.html#databricks-client-unified-" +
				"authentication to configure credentials for your preferred " +
				"authentication method. Config: host=https://" + authAWSHost +
				", workspace_id=12345, client_id=replay-sp, client_secret=***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isolateDefaultChain(t)
			s := &authServer{metadata: authAWSMetadata(), script: tt.script}

			a := refreshAuthVolume(t, s, authM2MConfig(tt.authType))

			a.assertFailed(t, tt.summary)
		})
	}
}

func TestVolumeAuthM2MTokenEndpointNotFoundDropsVolume(t *testing.T) {
	// The SDK maps a token endpoint's 404 to apierr.ErrNotFound, so the provider treats
	// the volume as deleted, and the next plan creates it.
	s := &authServer{metadata: authAWSMetadata(), script: []authExchange{
		authOIDCDiscovery,
		authM2MToken(http.StatusNotFound, `{"error": "not_found"}`),
	}}

	a := refreshAuthVolume(t, s, authM2MConfig("oauth-m2m"))

	assert.False(t, a.diags.HasError(), "%v", a.diags)
	assert.Nil(t, a.state)
	require.NotNil(t, a.plan)
	assert.Equal(t, &terraform.ResourceAttrDiff{New: "auth_volume"},
		a.plan.Attributes["name"])
	assert.True(t, a.plan.Attributes["volume_path"].NewComputed)
}

func TestVolumeAuthM2MAuthenticatesEachRetriedAttempt(t *testing.T) {
	// The first token expires during the 1.05 to 1.75 s backoff after the 429, so the
	// retried GET fetches and sends a second token.
	s := &authServer{metadata: authAWSMetadata(), script: []authExchange{
		authOIDCDiscovery,
		authM2MToken(http.StatusOK, authBearerToken("m2m-token-1", 1)),
		{authRequest{method: "GET", url: authAWSHost + authVolumePath,
			headers: map[string]string{"Authorization": "Bearer m2m-token-1"}},
			http.StatusTooManyRequests, tooManyRequestsJSON},
		authM2MToken(http.StatusOK, authBearerToken("m2m-token-2", 3600)),
		authBearerRead("m2m-token-2"),
	}}

	refreshAuthVolume(t, s, authM2MConfig("oauth-m2m")).assertSucceeded(t)
}

// authSeveralSourcesConfig sets a PAT, a username and password, and an OAuth client.
func authSeveralSourcesConfig(authType string) *config.Config {
	return &config.Config{Host: "https://" + authAWSHost, AuthType: authType,
		Token: "replay-pat", Username: "replay-user", Password: "replay-password",
		ClientID: "replay-sp", ClientSecret: "replay-secret"}
}

func TestVolumeAuthExplicitAuthTypeSelectsOneSource(t *testing.T) {
	tests := []struct {
		authType string
		script   []authExchange
	}{
		{"pat", []authExchange{
			authBearerRead("replay-pat"),
		}},
		{"basic", []authExchange{
			authVolumeRead(authAWSHost, map[string]string{
				"Authorization": basicAuthorization("replay-user", "replay-password")}),
		}},
		{"oauth-m2m", []authExchange{
			authOIDCDiscovery,
			authM2MToken(http.StatusOK, authBearerToken("m2m-token", 3600)),
			authBearerRead("m2m-token"),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.authType, func(t *testing.T) {
			s := &authServer{metadata: authAWSMetadata(), script: tt.script}

			a := refreshAuthVolume(t, s, authSeveralSourcesConfig(tt.authType))

			a.assertSucceeded(t)
			assert.Equal(t, tt.authType, a.client.Config.AuthType)
		})
	}
}

func TestVolumeAuthSeveralSourcesWithoutAuthTypeFailConfiguration(t *testing.T) {
	s := &authServer{metadata: authAWSMetadata()}

	_, err := newAuthClient(t, s, authSeveralSourcesConfig(""))

	assert.EqualError(t, err, "validate: more than one authorization method configured: "+
		"basic and oauth and pat. Config: host=https://"+authAWSHost+", token=***, "+
		"username=replay-user, password=***, client_id=replay-sp, client_secret=***")
	assert.Empty(t, s.received())
}

func TestVolumeAuthDefaultChainFallsBackAfterFailedDryRun(t *testing.T) {
	// OAuth M2M fails its dry run, github-oidc finds no GitHub token, and env-oidc
	// exchanges the OIDC token variable for a Databricks token with the same client ID.
	isolateDefaultChain(t)
	t.Setenv("DATABRICKS_OIDC_TOKEN", "replay-oidc-token")
	s := &authServer{metadata: authAWSMetadata(), script: []authExchange{
		authOIDCDiscovery,
		authM2MToken(http.StatusUnauthorized, authM2MRejected),
		authOIDCDiscovery,
		authOIDCDiscovery,
		{authRequest{method: "POST", url: authAWSHost + "/oidc/v1/token",
			body: "client_id=replay-sp" +
				"&grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange" +
				"&scope=all-apis&subject_token=replay-oidc-token" +
				"&subject_token_type=urn%3Aietf%3Aparams%3Aoauth%3Atoken-type%3Ajwt",
			headers: map[string]string{}},
			http.StatusOK, authBearerToken("federated-token", 3600)},
		authBearerRead("federated-token"),
	}}

	a := refreshAuthVolume(t, s, authM2MConfig(""))

	a.assertSucceeded(t)
	assert.Equal(t, "env-oidc", a.client.Config.AuthType)
}
