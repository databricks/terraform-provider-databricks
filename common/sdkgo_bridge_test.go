package common

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/sdk-go/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

// countingCredentials sets a new token and an extra header on every request, like a
// refreshing Azure service principal.
type countingCredentials struct {
	calls int
	err   error
}

func (c *countingCredentials) Name() string { return "counting" }

func (c *countingCredentials) Configure(context.Context,
	*config.Config) (credentials.CredentialsProvider, error) {
	return credentials.CredentialsProviderFn(func(r *http.Request) error {
		if c.err != nil {
			return c.err
		}
		c.calls++
		r.Header.Set("Authorization", fmt.Sprintf("Bearer token-%d", c.calls))
		r.Header.Set("X-Databricks-Azure-SP-Management-Token", "management")
		return nil
	}), nil
}

func newCredentialsTestConfig(t *testing.T, creds *countingCredentials) *config.Config {
	t.Helper()
	cfg := &config.Config{
		Host:                 "https://example.cloud.databricks.com",
		Credentials:          creds,
		Loaders:              []config.Loader{sdkGoTestLoader{}},
		HostMetadataResolver: noHostMetadata,
	}
	require.NoError(t, cfg.EnsureResolved())
	return cfg
}

func TestClassicCredentialsAuthenticateEveryRequest(t *testing.T) {
	creds := &countingCredentials{}
	adapter := classicCredentials{newCredentialsTestConfig(t, creds), time.Minute}
	for i := 1; i <= 2; i++ {
		headers, err := adapter.AuthHeaders(context.Background())
		require.NoError(t, err)
		assert.ElementsMatch(t, []auth.Header{
			{Key: "Authorization", Value: fmt.Sprintf("Bearer token-%d", i)},
			{Key: "X-Databricks-Azure-Sp-Management-Token", Value: "management"},
		}, headers)
	}
}

func TestClassicCredentialsReportAuthenticationErrors(t *testing.T) {
	creds := &countingCredentials{err: errors.New("token expired")}
	adapter := classicCredentials{newCredentialsTestConfig(t, creds), time.Minute}
	_, err := adapter.AuthHeaders(context.Background())
	assert.EqualError(t, err, "failed during request visitor: token expired")
}

func TestSdkGoUserAgentKeepsClassicFieldsAndAddsModules(t *testing.T) {
	ctx := useragent.InContext(context.Background(), "resource", "service_principal")
	sdkGo := "sdk-go-core/0.1.0 go/1.26.0 os/linux sdk-go-scim/0.6.0 auth/classic-config"
	got := sdkGoUserAgent(ctx, sdkGo)
	want := useragent.FromContext(ctx) + " sdk-go-core/0.1.0 sdk-go-scim/0.6.0"
	assert.Equal(t, want, got)
	assert.Contains(t, got, "resource/service_principal")
}

func TestClassicTransportFollowsClientConfig(t *testing.T) {
	custom := http.RoundTripper(&http.Transport{})
	assert.Same(t, custom, classicTransport(httpclient.ClientConfig{Transport: custom}))

	insecureConfig := httpclient.ClientConfig{InsecureSkipVerify: true}
	insecure := classicTransport(insecureConfig).(*http.Transport)
	require.NotNil(t, insecure.TLSClientConfig)
	assert.True(t, insecure.TLSClientConfig.InsecureSkipVerify)

	shared := classicTransport(httpclient.ClientConfig{}).(*http.Transport)
	assert.Nil(t, shared.TLSClientConfig)
	assert.NotNil(t, shared.Proxy, "proxy settings come from the environment")
}

type skipRetryTransport struct{ http.RoundTripper }

func (skipRetryTransport) SkipRetryOnIO() bool { return true }

// newBridgeTestClient returns a classic client for cfg, as the provider builds it.
func newBridgeTestClient(t *testing.T, cfg *config.Config) *client.DatabricksClient {
	t.Helper()
	c, err := client.New(cfg)
	require.NoError(t, err)
	return c
}

// newTestBridge returns a new bridge for a classic client built from cfg.
func newTestBridge(t *testing.T, cfg *config.Config) *sdkGoBridge {
	t.Helper()
	classic := newBridgeTestClient(t, cfg)
	b, err := newSdkGoBridge(classic.Config, classic.ApiClient())
	require.NoError(t, err)
	return b
}

func TestSdkGoBridgeFollowsRateLimitAndTimeouts(t *testing.T) {
	cfg := newCredentialsTestConfig(t, &countingCredentials{})
	cfg.RateLimitPerSecond = 3
	cfg.RetryTimeoutSeconds = 7
	cfg.HTTPTimeoutSeconds = 11
	b := newTestBridge(t, cfg)
	assert.Equal(t, rate.Limit(3), b.limiter.Limit())
	assert.Equal(t, 3, b.limiter.Burst())
	assert.Equal(t, "7s", b.retry.Timeout().String())
	assert.Equal(t, "11s", b.httpTimeout.String())
	assert.Equal(t, "11s", b.httpClient.Transport.(*sdkGoTransport).httpTimeout.String())
}

func TestSdkGoBridgeDefaultsMatchClassicClient(t *testing.T) {
	cfg := newCredentialsTestConfig(t, &countingCredentials{})
	b := newTestBridge(t, cfg)
	assert.Equal(t, rate.Limit(15), b.limiter.Limit())
	assert.Equal(t, "5m0s", b.retry.Timeout().String())
	assert.Equal(t, "1m0s", b.httpTimeout.String())
}

func TestSdkGoBridgeSkipsRateLimitForTestFixtures(t *testing.T) {
	cfg := newCredentialsTestConfig(t, &countingCredentials{})
	cfg.HTTPTransport = skipRetryTransport{}
	b := newTestBridge(t, cfg)
	assert.Equal(t, rate.Inf, b.limiter.Limit())
}

// The classic limiter is read from an unexported field, so this test fails if a
// databricks-sdk-go upgrade moves it.
func TestClassicRateLimiterReadsTheClassicClientsLimiter(t *testing.T) {
	cfg := newCredentialsTestConfig(t, &countingCredentials{})
	cfg.RateLimitPerSecond = 4
	classic := newBridgeTestClient(t, cfg)
	limiter, err := classicRateLimiter(classic.ApiClient())
	require.NoError(t, err)
	assert.Equal(t, rate.Limit(4), limiter.Limit())
	again, err := classicRateLimiter(classic.ApiClient())
	require.NoError(t, err)
	assert.Same(t, limiter, again)
}

func TestSdkGoBridgeIsSharedPerApiClient(t *testing.T) {
	classic := newBridgeTestClient(t, newCredentialsTestConfig(t, &countingCredentials{}))
	first, err := sdkGoBridgeFor(classic.Config, classic.ApiClient())
	require.NoError(t, err)
	second, err := sdkGoBridgeFor(classic.Config, classic.ApiClient())
	require.NoError(t, err)
	assert.Same(t, first, second)
	// A workspace client built from the same config has its own API client.
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(classic.Config))
	require.NoError(t, err)
	api, err := workspaceApiClient(w)
	require.NoError(t, err)
	other, err := sdkGoBridgeFor(w.Config, api)
	require.NoError(t, err)
	assert.NotSame(t, first, other)
}

// The API client is read from an unexported field, so this test fails if a
// databricks-sdk-go upgrade moves it.
func TestWorkspaceApiClientReadsTheWorkspaceClientsApiClient(t *testing.T) {
	cfg := newCredentialsTestConfig(t, &countingCredentials{})
	cfg.RateLimitPerSecond = 4
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	api, err := workspaceApiClient(w)
	require.NoError(t, err)
	again, err := workspaceApiClient(w)
	require.NoError(t, err)
	assert.Same(t, api, again)
	limiter, err := classicRateLimiter(api)
	require.NoError(t, err)
	assert.Equal(t, rate.Limit(4), limiter.Limit())
}

func TestWorkspaceApiClientRejectsMockWorkspaceClients(t *testing.T) {
	_, err := workspaceApiClient(&databricks.WorkspaceClient{})
	assert.EqualError(t, err, "*databricks.WorkspaceClient has no API client to share")
}

func TestClassicErrorKeepsTransportErrors(t *testing.T) {
	cfg := newCredentialsTestConfig(t, &countingCredentials{})
	b := newTestBridge(t, cfg)
	transportErr := errors.New("dial tcp: connection refused")
	assert.Same(t, transportErr, b.classicError(context.Background(), transportErr, nil))
}
