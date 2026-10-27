// Package clienttest replays helper-level scenarios that do not run Terraform.
package clienttest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

// WorkspaceClient loads a consumable HTTP transcript and supplies local PAT
// authentication. It checks for unused fixtures and unexpected calls on cleanup.
func WorkspaceClient(t *testing.T, path string) *databricks.WorkspaceClient {
	t.Helper()
	transcript := loadTranscript(t, path)
	handler := httpfixture.NewHandler()
	require.NoError(t, handler.SetReusableFixtures([]httpfixture.Call{{
		Request:  httpfixture.Request{Method: "GET", Path: "/.well-known/databricks-config"},
		Response: httpfixture.Response{Status: 200, Body: &httpfixture.Body{JSON: transcript.HostMetadata}},
	}}))
	require.NoError(t, handler.StartStep(t.Name(), transcript.Calls))
	server := httptest.NewServer(handler)
	t.Cleanup(func() {
		server.Close()
		require.NoError(t, handler.Verify())
	})
	client, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host: server.URL, Token: "fixture-token", AuthType: "pat",
		Credentials: config.PatCredentials{}, Loaders: []config.Loader{fixtureLoader{}},
		HTTPTransport: localTransport{host: server.URL, transport: server.Client().Transport},
	})
	require.NoError(t, err)
	return client
}

type transcript struct {
	HostMetadata map[string]any     `yaml:"host_metadata"`
	Calls        []httpfixture.Call `yaml:"calls"`
}

func loadTranscript(t *testing.T, path string) transcript {
	t.Helper()
	file, err := os.Open(path)
	require.NoError(t, err)
	defer file.Close()
	var result transcript
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	require.NoError(t, decoder.Decode(&result))
	require.ErrorIs(t, decoder.Decode(new(any)), io.EOF, "expected one YAML document")
	return result
}

type fixtureLoader struct{}

func (fixtureLoader) Name() string                   { return "replay fixture configuration" }
func (fixtureLoader) Configure(*config.Config) error { return nil }

type localTransport struct {
	host      string
	transport http.RoundTripper
}

func (l localTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	host, err := url.Parse(l.host)
	if err != nil {
		return nil, err
	}
	if request.URL.Scheme != host.Scheme || request.URL.Host != host.Host {
		return nil, fmt.Errorf("replay rejected request outside fixture server: %s", request.URL)
	}
	return l.transport.RoundTrip(request)
}
