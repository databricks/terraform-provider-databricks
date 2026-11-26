package config

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
)

// tokenCountingTransport wraps an http.RoundTripper and counts the number of
// round trips made. It is used to verify that the token source does not cache
// tokens internally, ensuring that each Token() call results in an HTTP
// request.
type tokenCountingTransport struct {
	inner http.RoundTripper
	count atomic.Int32
}

func (t *tokenCountingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.count.Add(1)
	return t.inner.RoundTrip(req)
}

func (t *tokenCountingTransport) SkipRetryOnIO() bool {
	return true
}

// TestAzureClientSecretCredentials_tokenSourceFor_noCaching verifies that
// the auth.TokenSource returned by tokenSourceFor does not cache tokens
// internally. Before the fix, the code used clientcredentials.Config.TokenSource
// which returns an oauth2.ReuseTokenSource with a 10s expiryDelta. This inner
// cache swallowed proactive refresh attempts from the outer caching layers
// (azureReuseTokenSource and serviceToServiceVisitor), causing bursts
// of 401s at token expiry.
// See https://github.com/databricks/databricks-sdk-go/issues/1549.
func TestAzureClientSecretCredentials_tokenSourceFor_noCaching(t *testing.T) {
	counter := &tokenCountingTransport{
		inner: fixtures.MappingTransport{
			"POST /test-tenant-id/oauth2/token": {
				Response: map[string]any{
					"token_type":   "Bearer",
					"access_token": "test-token",
					"expires_in":   3600,
				},
			},
		},
	}

	cfg := &Config{
		AzureClientID:     "test-client-id",
		AzureClientSecret: "test-client-secret",
		AzureTenantID:     "test-tenant-id",
	}
	cfg.refreshClient = httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: counter,
	})

	c := AzureClientSecretCredentials{}
	ts := c.tokenSourceFor(context.Background(), cfg, "https://login.microsoftonline.com/", "https://management.azure.com/")

	const numCalls = 3
	for i := range numCalls {
		tok, err := ts.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() call %d: unexpected error: %v", i+1, err)
		}
		if tok.AccessToken != "test-token" {
			t.Errorf("Token() call %d: got access token %q, want %q", i+1, tok.AccessToken, "test-token")
		}
	}

	if got := int(counter.count.Load()); got != numCalls {
		t.Errorf("token endpoint was called %d times, want %d (inner caching is present)", got, numCalls)
	}
}

// TestAzureClientSecretCredentials_Configure verifies the happy path of the
// full Azure client secret credential flow, including both the inner
// (workspace) token and the management token.
func TestAzureClientSecretCredentials_Configure(t *testing.T) {
	tokenExpiry := time.Now().Add(time.Hour).Unix()
	assertHeaders(t, &Config{
		Host:              "https://adb-123.azuredatabricks.net",
		AzureClientID:     "test-client-id",
		AzureClientSecret: "test-client-secret",
		AzureTenantID:     "test-tenant-id",
		AuthType:          "azure-client-secret",
		HTTPTransport: fixtures.MappingTransport{
			"POST /test-tenant-id/oauth2/token": {
				Response: map[string]any{
					"token_type":   "Bearer",
					"access_token": "workspace-token",
					"expires_on":   fmt.Sprintf("%d", tokenExpiry),
				},
			},
		},
	}, map[string]string{
		"Authorization":                          "Bearer workspace-token",
		"X-Databricks-Azure-Sp-Management-Token": "workspace-token",
	})
}
