package terraformtest

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

const hostsDefinition = `version: 1
provider_host: account
hosts:
  account:
    url: https://accounts.cloud.databricks.com
  workspace:
    url: https://replay-workspace-1002.cloud.databricks.com
    metadata: {host_type: workspace, workspace_id: "1002"}
steps:
  - name: create
    config: 'provider "databricks" {}'
    calls:
      - request: {method: GET, host: workspace, path: /items, body: null}
        response: {status: 200}
destroy:
  calls: []
`

func TestHostsReplay(t *testing.T) {
	Run(t, "testdata/hosts")
}

func TestHostsShowcaseFailsOnWrongHost(t *testing.T) {
	contents, err := os.ReadFile("testdata/hosts/workspace_host.yaml")
	require.NoError(t, err)
	fixture := strings.Replace(string(contents), "          method: POST\n          host: workspace\n", "          method: POST\n          host: account\n", 1)
	require.NotEqual(t, string(contents), fixture)
	directory := t.TempDir()
	writeDiscoveryFile(t, directory, "workspace_host.yaml", fixture)
	output, err := discoverySubprocess(t, directory, "^TestDirectoryRunnerHelper$")
	require.Error(t, err, output)
	require.Contains(t, output, `step 1 (create): unexpected HTTP call: POST /api/2.0/workspace/mkdirs on host "workspace"`)
	require.Contains(t, output, `step 1 (create): unconsumed call 2: POST /api/2.0/workspace/mkdirs on host "account"`)
}

func TestDecodeHosts(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(hostsDefinition))
	require.NoError(t, err)
	require.Equal(t, "account", d.ProviderHost)
	require.Equal(t, "https://replay-workspace-1002.cloud.databricks.com", d.Hosts["workspace"].URL)
	require.Equal(t, "workspace", d.Steps[0].Calls[0].Request.Host)
	metadata := d.metadataCalls()
	require.Len(t, metadata, 1)
	require.Equal(t, "workspace", metadata[0].Request.Host)
	require.Equal(t, map[string]any{"host_type": "workspace", "workspace_id": "1002"}, metadata[0].Response.Body.JSON)
	for _, tc := range []struct {
		name, before, after, message string
	}{
		{"unknown provider host", "provider_host: account", "provider_host: accounts", `provider_host "accounts" is not one of hosts`},
		{"provider host without hosts", "hosts:\n  account:\n    url: https://accounts.cloud.databricks.com\n  workspace:\n    url: https://replay-workspace-1002.cloud.databricks.com\n    metadata: {host_type: workspace, workspace_id: \"1002\"}\n", "", `provider_host "account" is not one of hosts`},
		{"metadata twice", "provider_host: account", "provider_host: workspace\nhost_metadata: {host_type: workspace}", "host_metadata and hosts.workspace.metadata both describe the provider host"},
		{"missing scheme", "url: https://accounts.cloud.databricks.com", "url: accounts.cloud.databricks.com", `hosts.account.url: expected an http or https URL without a path, query or fragment, got "accounts.cloud.databricks.com"`},
		{"other scheme", "url: https://accounts", "url: ftp://accounts", "hosts.account.url: expected an http or https URL"},
		{"path", "url: https://accounts.cloud.databricks.com", "url: https://accounts.cloud.databricks.com/api", "hosts.account.url: expected an http or https URL"},
		{"query", "url: https://accounts.cloud.databricks.com", "url: https://accounts.cloud.databricks.com?a=b", "hosts.account.url: expected an http or https URL"},
		{"fragment", "url: https://accounts.cloud.databricks.com", "url: https://accounts.cloud.databricks.com#a", "hosts.account.url: expected an http or https URL"},
		{"user", "url: https://accounts.cloud.databricks.com", "url: https://user@accounts.cloud.databricks.com", "hosts.account.url: expected an http or https URL"},
		{"same authority", "url: https://replay-workspace-1002.cloud.databricks.com", "url: http://accounts.cloud.databricks.com", "hosts.workspace.url: accounts.cloud.databricks.com is also the host of hosts.account"},
		{"empty name", "  workspace:\n", "  \"\":\n", "hosts: a host name must not be empty"},
		{"unknown field", "    url: https://accounts.cloud.databricks.com", "    url: https://accounts.cloud.databricks.com\n    provider: true", "field provider not found"},
		{"unknown call host", "host: workspace, path: /items", "host: workspac, path: /items", `step 1 (create): call 1: host "workspac" is not a named host`},
		{"unknown destroy host", "destroy:\n  calls: []", "destroy:\n  calls:\n    - request: {method: GET, host: workspac, path: /items, body: null}\n      response: {status: 200}", `destroy: call 1: host "workspac" is not a named host`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Replace(hostsDefinition, tc.before, tc.after, 1)
			require.NotEqual(t, hostsDefinition, input)
			_, err := decodeDefinition(strings.NewReader(input))
			require.ErrorContains(t, err, tc.message)
		})
	}
}

func TestRecordingRejectsHosts(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(hostsDefinition))
	require.NoError(t, err)
	require.ErrorContains(t, validateRecordingSetup(d), "recording does not support hosts")
}

func TestFixtureTransportRoutesNamedHosts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(r.Host + " " + r.URL.RequestURI()))
	}))
	defer server.Close()
	d, err := decodeDefinition(strings.NewReader(hostsDefinition))
	require.NoError(t, err)
	hosts := d.replayHosts(server.URL)
	transport := fixtureTransport{host: server.URL, named: hosts.named, transport: http.DefaultTransport}
	for target, served := range map[string]string{
		server.URL + "/local":                                        strings.TrimPrefix(server.URL, "http://") + " /local",
		"https://accounts.cloud.databricks.com/api?x=1":              "accounts.cloud.databricks.com /api?x=1",
		"https://replay-workspace-1002.cloud.databricks.com/api/2.0": "replay-workspace-1002.cloud.databricks.com /api/2.0",
	} {
		req, err := http.NewRequest("GET", target, nil)
		require.NoError(t, err)
		resp, err := transport.RoundTrip(req)
		require.NoError(t, err)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		resp.Body.Close()
		require.Equal(t, served, string(body))
		// The caller's request is unchanged and is the response's request.
		require.Same(t, req, resp.Request)
		require.Equal(t, target, req.URL.String())
	}
	for _, target := range []string{"http://accounts.cloud.databricks.com/api", "https://.cloud.databricks.com/api", "https://example.invalid"} {
		req, err := http.NewRequest("GET", target, nil)
		require.NoError(t, err)
		_, err = transport.RoundTrip(req)
		require.EqualError(t, err, "hermetic replay rejected request outside fixture server: "+target)
	}
}

func TestFixtureConfigSelectsProviderHost(t *testing.T) {
	d, err := decodeDefinition(strings.NewReader(hostsDefinition))
	require.NoError(t, err)
	server := "http://127.0.0.1:1234"
	withoutProvider := d.replayHosts(server)
	withoutProvider.provider = ""
	for _, tc := range []struct {
		name       string
		hosts      replayHosts
		configured string
		expected   string
	}{
		{"no hosts ignore HCL", replayHosts{server: server}, "https://accounts.cloud.databricks.com", server},
		{"provider host", d.replayHosts(server), "", "https://accounts.cloud.databricks.com"},
		{"no provider host", withoutProvider, "", server},
		{"HCL names a host", d.replayHosts(server), "https://replay-workspace-1002.cloud.databricks.com/", "https://replay-workspace-1002.cloud.databricks.com/"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &config.Config{Host: tc.configured}
			require.NoError(t, tc.hosts.fixtureConfig()(cfg))
			require.Equal(t, tc.expected, cfg.Host)
			require.Equal(t, "fixture-token", cfg.Token)
		})
	}
	cfg := &config.Config{Host: "https://other.cloud.databricks.com"}
	require.EqualError(t, d.replayHosts(server).fixtureConfig()(cfg), "hermetic replay: provider host https://other.cloud.databricks.com is not one of the journey's hosts")
}

// Named hosts keep the cloud of their DNS zone in the SDK: host metadata without
// a cloud resolves it from the host, and private-link errors name that cloud.
func TestNamedHostsKeepTheirCloud(t *testing.T) {
	for _, tc := range []struct {
		host    string
		cloud   environment.Cloud
		service string
	}{
		{"https://replay-workspace-1002.cloud.databricks.com", environment.CloudAWS, "AWS PrivateLink"},
		{"https://adb-1234567890123456.7.azuredatabricks.net", environment.CloudAzure, "Azure Private Link"},
		{"https://1234567890123456.7.gcp.databricks.com", environment.CloudGCP, "Private Service Connect"},
	} {
		t.Run(string(tc.cloud), func(t *testing.T) {
			d := &testDefinition{Hosts: map[string]namedHost{"workspace": {URL: tc.host}}, ProviderHost: "workspace"}
			handler := httpfixture.NewHandler()
			require.NoError(t, d.setHosts(handler))
			require.NoError(t, handler.SetReusableFixtures([]httpfixture.Call{metadataCall("", map[string]any{"host_type": "workspace"})}))
			require.NoError(t, handler.StartStep("read", []httpfixture.Call{
				{Request: httpfixture.Request{Method: "GET", Path: "/api/2.0/items"}, Response: httpfixture.Response{Status: 302, Headers: map[string]string{"Location": "/login.html?error=private-link-validation-error"}}},
				{Request: httpfixture.Request{Method: "GET", Path: "/login.html", Query: map[string][]string{"error": {"private-link-validation-error"}}}, Response: httpfixture.Response{Status: 200}},
			}))
			server := httptest.NewServer(handler)
			defer server.Close()
			cfg := &config.Config{}
			require.NoError(t, d.replayHosts(server.URL).fixtureConfig()(cfg))
			c, err := client.New(cfg)
			require.NoError(t, err)
			require.Equal(t, tc.cloud, cfg.Cloud)
			err = c.Do(context.Background(), "GET", "/api/2.0/items", nil, nil, nil, nil)
			require.ErrorContains(t, err, "The requested workspace has "+tc.service+" enabled")
			require.NoError(t, handler.Verify())
		})
	}
}
