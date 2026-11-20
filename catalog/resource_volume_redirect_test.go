package catalog

import (
	"bytes"
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests pin the cloud named by the private-link login redirect on a
// databricks_volume refresh. The classic SDK picks the cloud from the final URL's host,
// and replay serves every call from 127.0.0.1, which falls back to AWS, so replay can't
// pin the Azure and GCP messages. Here the client uses a workspace domain of each cloud,
// and a transport sends that domain to an httptest server.

const (
	redirectVolumeID   = "main.replay.redirect_cloud_volume"
	redirectVolumePath = "/api/2.1/unity-catalog/volumes/" + redirectVolumeID
	redirectLoginPath  = "/login.html?error=private-link-validation-error"
)

// redirectCallLog records the calls the client sends, in order, as the client sees
// them.
type redirectCallLog struct {
	mu    sync.Mutex
	calls []string
}

func (l *redirectCallLog) add(call string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.calls = append(l.calls, call)
}

func (l *redirectCallLog) snapshot() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.calls...)
}

// redirectLoginServer answers host metadata, redirects the volume GET to the
// private-link login page, and serves that page. Other calls get 404.
func redirectLoginServer(metadata string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/.well-known/databricks-config":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, metadata)
		case r.Method == http.MethodGet && r.URL.Path == redirectVolumePath:
			w.Header().Set("Location", redirectLoginPath)
			w.WriteHeader(http.StatusFound)
		case r.Method == http.MethodGet && r.URL.RequestURI() == redirectLoginPath:
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = io.WriteString(w, "<html><body>Sign in</body></html>")
		default:
			http.NotFound(w, r)
		}
	})
}

// redirectCloudTransport sends the workspace domain to the httptest server, logs every
// call but host metadata, and rejects any other host. It returns each response with
// the client's request, because the SDK reads the final URL's host from it.
type redirectCloudTransport struct {
	domain string
	target *url.URL
	log    *redirectCallLog
}

func (t redirectCloudTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Host != t.domain {
		return nil, fmt.Errorf("request outside the test server: %s", req.URL)
	}
	var body []byte
	if req.Body != nil {
		var err error
		if body, err = io.ReadAll(req.Body); err != nil {
			return nil, err
		}
	}
	if req.URL.Path != "/.well-known/databricks-config" {
		t.log.add(fmt.Sprintf("%s %s body=%q", req.Method, req.URL, body))
	}
	out := req.Clone(req.Context())
	if req.Body != nil {
		out.Body = io.NopCloser(bytes.NewReader(body))
	}
	out.URL.Scheme, out.URL.Host, out.Host = t.target.Scheme, t.target.Host, t.target.Host
	resp, err := http.DefaultTransport.RoundTrip(out)
	if resp != nil {
		resp.Request = req
	}
	return resp, err
}

// redirectLoader keeps ambient profiles and environment variables out of the client
// configuration.
type redirectLoader struct{}

func (redirectLoader) Name() string { return "redirect test" }

func (redirectLoader) Configure(*config.Config) error { return nil }

func TestVolumeRefreshPrivateLinkRedirectNamesTheCloudOfTheHost(t *testing.T) {
	tests := []struct {
		name   string
		domain string
		cloud  string
		// diagnostic is the refresh's exact error.
		diagnostic string
	}{
		{
			name:   "azure",
			domain: "adb-1234567890123456.7.azuredatabricks.net",
			cloud:  "AZURE",
			diagnostic: "cannot read volume: The requested workspace has Azure Private " +
				"Link enabled and is not accessible from the current network. Ensure " +
				"that Azure Private Link is properly configured and that your device " +
				"has access to the Azure Private Link endpoint. For more information, " +
				"see https://learn.microsoft.com/en-us/azure/databricks/security/" +
				"network/classic/private-link-standard#authentication-troubleshooting.",
		},
		{
			name:   "gcp",
			domain: "1234567890123456.7.gcp.databricks.com",
			cloud:  "GCP",
			diagnostic: "cannot read volume: The requested workspace has Private " +
				"Service Connect enabled and is not accessible from the current " +
				"network. Ensure that Private Service Connect is properly configured " +
				"and that your device has access to the GCP VPC endpoint. For more " +
				"information, see https://docs.gcp.databricks.com/en/security/network/" +
				"classic/private-service-connect.html.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(redirectLoginServer(fmt.Sprintf(
				`{"host_type": "workspace", "cloud": %q, "workspace_id": "12345"}`,
				tt.cloud)))
			defer server.Close()
			target, err := url.Parse(server.URL)
			require.NoError(t, err)
			log := &redirectCallLog{}
			c := newRedirectClient(t, tt.domain,
				redirectCloudTransport{domain: tt.domain, target: target, log: log})

			before, after, diags := refreshRedirectVolume(t, c)

			assert.Equal(t, diag.Diagnostics{{Severity: diag.Error,
				Summary: tt.diagnostic}}, diags)
			host := "https://" + tt.domain
			assert.Equal(t, []string{
				// The classic client ends the GET's URL with an empty query.
				"GET " + host + redirectVolumePath + `? body=""`,
				"GET " + host + redirectLoginPath + ` body=""`,
			}, log.snapshot())
			// The failed refresh keeps the volume's ID and every attribute.
			assert.Equal(t, redirectVolumeID, after["id"])
			assert.Equal(t, before, after)
		})
	}
}

func newRedirectClient(
	t *testing.T, domain string, transport http.RoundTripper,
) *common.DatabricksClient {
	t.Helper()
	inner, err := client.New(&config.Config{
		Host:          "https://" + domain,
		Token:         "fixture-token",
		AuthType:      "pat",
		Credentials:   config.PatCredentials{},
		Loaders:       []config.Loader{redirectLoader{}},
		HTTPTransport: transport,
	})
	require.NoError(t, err)
	c := &common.DatabricksClient{DatabricksClient: inner}
	// Provider configuration caches the host metadata's workspace ID, so a refresh
	// doesn't look it up through /Me.
	c.SetCachedWorkspaceID(12345)
	return c
}

// refreshRedirectVolume runs the provider's wrapped Read for a volume already in state
// and returns its state attributes before and after the Read, with the diagnostics.
func refreshRedirectVolume(
	t *testing.T, c *common.DatabricksClient,
) (before, after map[string]string, diags diag.Diagnostics) {
	t.Helper()
	ctx := context.WithValue(context.Background(), common.ResourceName, "volume")
	r := ResourceVolume().ToResource()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"name":         "redirect_cloud_volume",
		"catalog_name": "main",
		"schema_name":  "replay",
		"volume_type":  "MANAGED",
		"owner":        "replay-admin@example.com",
		// A create puts the workspace ID here, so a refresh doesn't look it up.
		"provider_config": []any{map[string]any{"workspace_id": "12345"}},
	})
	d.SetId(redirectVolumeID)
	require.NoError(t, d.Set("volume_path", "/Volumes/main/replay/redirect_cloud_volume"))
	before = d.State().Attributes
	diags = r.ReadContext(ctx, d, c)
	return before, d.State().Attributes, diags
}
