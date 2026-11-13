package common

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/sdk-go/options/call"
	sdkscim "github.com/databricks/sdk-go/scim/v1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type sdkGoTestLoader struct{}

func (sdkGoTestLoader) Name() string                   { return "sdk-go test" }
func (sdkGoTestLoader) Configure(*config.Config) error { return nil }

func noHostMetadata(context.Context, string) (*config.HostMetadata, error) {
	return nil, nil
}

// newSdkGoTestClient returns a workspace client for server that authenticates with a
// fixed PAT and ignores ambient configuration.
func newSdkGoTestClient(t *testing.T, server *httptest.Server,
	customize func(*config.Config)) *DatabricksClient {
	t.Helper()
	cfg := &config.Config{
		Host:                 server.URL,
		Token:                "test-token",
		AuthType:             "pat",
		Credentials:          config.PatCredentials{},
		Loaders:              []config.Loader{sdkGoTestLoader{}},
		WorkspaceID:          "12345",
		AccountID:            "test-account",
		HostMetadataResolver: noHostMetadata,
	}
	if customize != nil {
		customize(cfg)
	}
	c, err := client.New(cfg)
	require.NoError(t, err)
	return &DatabricksClient{DatabricksClient: c}
}

func getServicePrincipal(ctx context.Context, t *testing.T, c *DatabricksClient,
	apiLevel string) (*sdkscim.ServicePrincipal, error) {
	t.Helper()
	sc, err := c.SdkGoScimClient(ctx, apiLevel)
	require.NoError(t, err)
	id := "201"
	get := func(ctx context.Context,
		opts ...call.Option) (*sdkscim.ServicePrincipal, error) {
		req := sdkscim.GetServicePrincipalRequest{Id: &id}
		return sc.GetServicePrincipal(ctx, req, opts...)
	}
	return CallSdkGoScim(ctx, sc, get)
}

// classicGetServicePrincipal makes the same read through the existing SCIM client.
func classicGetServicePrincipal(ctx context.Context, c *DatabricksClient) error {
	var sp map[string]any
	path := "/preview/scim/v2/ServicePrincipals/201"
	return c.Scim(ctx, http.MethodGet, path, nil, &sp, "")
}

type handler = func(http.ResponseWriter, *http.Request)

// serveSequence answers each request with the next handler and counts requests.
func serveSequence(t *testing.T, handlers ...handler) (*httptest.Server, *atomic.Int32) {
	t.Helper()
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		n := int(calls.Add(1))
		if n > len(handlers) {
			t.Errorf("unexpected request %d: %s %s", n, r.Method, r.URL)
			w.WriteHeader(http.StatusTeapot)
			return
		}
		handlers[n-1](w, r)
	}))
	t.Cleanup(server.Close)
	return server, &calls
}

func respond(status int, body string) handler {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(status)
		fmt.Fprint(w, body)
	}
}

const sdkGoTestServicePrincipal = `{"id": "201", "applicationId": "app-201",
	"displayName": "SPN"}`

type bothWays struct {
	err, classicErr     error
	calls, classicCalls int32
}

// readBothWays reads through sdk-go and the existing client, each against a fresh
// server with the same handlers.
func readBothWays(t *testing.T, ctx func() context.Context,
	customize func(*config.Config), handlers ...handler) bothWays {
	t.Helper()
	var result bothWays
	server, calls := serveSequence(t, handlers...)
	sdkGoClient := newSdkGoTestClient(t, server, customize)
	_, result.err = getServicePrincipal(ctx(), t, sdkGoClient, "")
	classicServer, classicCalls := serveSequence(t, handlers...)
	classicClient := newSdkGoTestClient(t, classicServer, customize)
	result.classicErr = classicGetServicePrincipal(ctx(), classicClient)
	result.calls, result.classicCalls = calls.Load(), classicCalls.Load()
	return result
}

func TestEnableSdkGoMarksOperationsAndImporter(t *testing.T) {
	seen := map[string]bool{}
	op := func(name string) resourceOperation {
		return func(ctx context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
			seen[name] = UsesSdkGo(ctx)
			return nil
		}
	}
	r := &schema.Resource{
		CreateContext: op("create"),
		ReadContext:   op("read"),
		UpdateContext: op("update"),
		DeleteContext: op("delete"),
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData,
				_ any) ([]*schema.ResourceData, error) {
				seen["import"] = UsesSdkGo(ctx)
				return []*schema.ResourceData{d}, nil
			},
		},
	}
	originalImporter := r.Importer
	EnableSdkGo(r)
	ctx := context.Background()
	r.CreateContext(ctx, nil, nil)
	r.ReadContext(ctx, nil, nil)
	r.UpdateContext(ctx, nil, nil)
	r.DeleteContext(ctx, nil, nil)
	_, err := r.Importer.StateContext(ctx, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, map[string]bool{
		"create": true, "read": true, "update": true, "delete": true, "import": true,
	}, seen)
	assert.False(t, UsesSdkGo(ctx))
	_, err = originalImporter.StateContext(ctx, nil, nil)
	require.NoError(t, err)
	assert.False(t, seen["import"], "the original importer must stay unwrapped")
}

func TestEnableSdkGoKeepsMissingOperations(t *testing.T) {
	r := &schema.Resource{}
	EnableSdkGo(r)
	assert.Nil(t, r.CreateContext)
	assert.Nil(t, r.ReadContext)
	assert.Nil(t, r.Importer)
}

func TestSdkGoScimClientWorkspaceRequest(t *testing.T) {
	server, calls := serveSequence(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/2.0/preview/scim/v2/ServicePrincipals/201", r.URL.Path)
		assert.Empty(t, r.URL.RawQuery)
		assert.Equal(t, "12345", r.Header.Get("X-Databricks-Workspace-Id"))
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.NotEmpty(t, r.Header.Get("Traceparent"))
		userAgent := r.Header.Get("User-Agent")
		assert.Contains(t, userAgent, "databricks-sdk-go/")
		assert.Contains(t, userAgent, "resource/service_principal")
		assert.Contains(t, userAgent, "auth/pat")
		assert.Contains(t, userAgent, " sdk-go-scim/")
		respond(http.StatusOK, sdkGoTestServicePrincipal)(w, r)
	})
	c := newSdkGoTestClient(t, server, nil)
	ctx := useragent.InContext(context.Background(), "resource", "service_principal")
	sp, err := getServicePrincipal(ctx, t, c, "")
	require.NoError(t, err)
	assert.Equal(t, "app-201", *sp.ApplicationId)
	assert.EqualValues(t, 1, calls.Load())
}

func TestSdkGoScimClientAccountRequest(t *testing.T) {
	server, _ := serveSequence(t, func(w http.ResponseWriter, r *http.Request) {
		path := "/api/2.0/accounts/test-account/scim/v2/ServicePrincipals/201"
		assert.Equal(t, path, r.URL.Path)
		assert.Empty(t, r.Header.Values("X-Databricks-Workspace-Id"))
		respond(http.StatusOK, sdkGoTestServicePrincipal)(w, r)
	})
	c := newSdkGoTestClient(t, server, nil)
	ctx := context.Background()
	sc, err := c.SdkGoScimClient(ctx, ApiLevelAccount)
	require.NoError(t, err)
	assert.True(t, sc.IsAccount)
	id := "201"
	get := func(ctx context.Context,
		opts ...call.Option) (*sdkscim.AccountServicePrincipal, error) {
		req := sdkscim.GetAccountServicePrincipalRequest{Id: &id}
		return sc.GetAccountServicePrincipal(ctx, req, opts...)
	}
	sp, err := CallSdkGoScim(ctx, sc, get)
	require.NoError(t, err)
	assert.Equal(t, "SPN", *sp.DisplayName)
}

func TestSdkGoScimClientAppliesConfigHeaders(t *testing.T) {
	server, _ := serveSequence(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "extra", r.Header.Get("X-Extra"))
		respond(http.StatusOK, sdkGoTestServicePrincipal)(w, r)
	})
	c := newSdkGoTestClient(t, server, func(cfg *config.Config) {
		cfg.Headers = func(r *http.Request) error {
			r.Header.Set("X-Extra", "extra")
			return nil
		}
	})
	_, err := getServicePrincipal(context.Background(), t, c, "")
	require.NoError(t, err)
}

func TestSdkGoScimClientMapsErrorsLikeClassicClient(t *testing.T) {
	for _, tc := range []struct {
		name    string
		status  int
		body    string
		message string
		missing bool
	}{
		{"not found", 404,
			`{"detail": "Service Principal with id 201 not found.", "status": "404"}`,
			"Service Principal with id 201 not found.", true},
		{"forbidden", 403, `{"detail": "User is not authorized.", "status": "403"}`,
			"User is not authorized.", false},
		{"scim type", 409,
			`{"scimType": "uniqueness", "detail": "exists", "status": "409"}`,
			"uniqueness exists", false},
		{"non-JSON error body", 500, "upstream connect error",
			"upstream connect error", false},
		{"undecodable success body", 200, "not json",
			"failed to unmarshal response body: " +
				"invalid character 'o' in literal null (expecting 'u')", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := readBothWays(t, context.Background, nil, respond(tc.status, tc.body))
			require.Error(t, r.err)
			assert.Equal(t, tc.message, r.err.Error())
			assert.Equal(t, tc.missing, apierr.IsMissing(r.err))
			if tc.status >= 400 {
				var apiErr *apierr.APIError
				require.ErrorAs(t, r.err, &apiErr)
				assert.Equal(t, tc.status, apiErr.StatusCode)
			}
			require.Error(t, r.classicErr)
			// The classic decode error also prints the request and response.
			classicMessage := r.classicErr.Error()
			assert.True(t, strings.HasPrefix(classicMessage, tc.message), classicMessage)
			assert.Equal(t, apierr.IsMissing(r.classicErr), apierr.IsMissing(r.err))
		})
	}
}

func TestSdkGoScimClientRetriesLikeClassicClient(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		body   string
		calls  int32
	}{
		{"429 is retried", 429, `{"detail": "Too many requests.", "status": "429"}`, 2},
		{"503 is retried", 503, `{"detail": "Unavailable.", "status": "503"}`, 2},
		{"transient message is retried", 500,
			`{"detail": "REQUEST_LIMIT_EXCEEDED", "status": "500"}`, 2},
		{"504 is not retried", 504, "upstream request timeout", 1},
		{"500 is not retried", 500, `{"detail": "Internal error.", "status": "500"}`, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := readBothWays(t, context.Background, nil, respond(tc.status, tc.body),
				respond(http.StatusOK, sdkGoTestServicePrincipal))
			assert.Equal(t, tc.calls, r.calls)
			assert.Equal(t, tc.calls == 1, r.err != nil, "error: %v", r.err)
			assert.Equal(t, r.classicCalls, r.calls)
			if r.err != nil {
				assert.EqualError(t, r.classicErr, r.err.Error())
			}
		})
	}
}

func TestSdkGoScimClientStopsRetryingAtRetryTimeout(t *testing.T) {
	tooMany := respond(429, `{"detail": "Too many requests.", "status": "429"}`)
	// The first classic backoff is at least one second, past the retry timeout.
	r := readBothWays(t, context.Background,
		func(cfg *config.Config) { cfg.RetryTimeoutSeconds = 1 }, tooMany, tooMany)
	assert.EqualError(t, r.err, "Too many requests.")
	assert.True(t, errors.Is(r.err, apierr.ErrTooManyRequests))
	assert.EqualValues(t, 1, r.calls)
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.Equal(t, r.classicCalls, r.calls)
}

func TestSdkGoScimClientReturnsLastErrorWhenCancelledDuringBackoff(t *testing.T) {
	var cancel context.CancelFunc
	newContext := func() context.Context {
		var ctx context.Context
		ctx, cancel = context.WithCancel(context.Background())
		return ctx
	}
	tooMany := func(w http.ResponseWriter, r *http.Request) {
		time.AfterFunc(200*time.Millisecond, cancel)
		respond(429, `{"detail": "Too many requests.", "status": "429"}`)(w, r)
	}
	r := readBothWays(t, newContext, nil, tooMany)
	assert.EqualError(t, r.err, "Too many requests.")
	assert.EqualValues(t, 1, r.calls)
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.Equal(t, r.classicCalls, r.calls)
}

func TestSdkGoScimClientTimesOutLikeClassicClient(t *testing.T) {
	release := make(chan struct{})
	defer close(release)
	hang := func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-release:
		case <-r.Context().Done():
		}
	}
	r := readBothWays(t, context.Background,
		func(cfg *config.Config) { cfg.HTTPTimeoutSeconds = 1 }, hang)
	require.Error(t, r.err)
	assert.Contains(t, r.err.Error(), "request timed out after 1s of inactivity")
	assert.EqualValues(t, 1, r.calls)
	require.Error(t, r.classicErr)
	assert.Contains(t, r.classicErr.Error(), "request timed out after 1s of inactivity")
	assert.Equal(t, r.classicCalls, r.calls)
}

func TestSdkGoScimClientReportsRefusedConnectionsLikeClassicClient(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	server.Close()
	// A refused connection is retried, so a short retry timeout ends the call.
	customize := func(cfg *config.Config) { cfg.RetryTimeoutSeconds = 1 }
	ctx := context.Background()
	_, err := getServicePrincipal(ctx, t, newSdkGoTestClient(t, server, customize), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "connection refused")
	classicClient := newSdkGoTestClient(t, server, customize)
	assert.EqualError(t, classicGetServicePrincipal(ctx, classicClient), err.Error())
}

func TestSdkGoScimClientSharesClassicRateLimit(t *testing.T) {
	server, calls := serveSequence(t, respond(http.StatusOK, sdkGoTestServicePrincipal),
		respond(http.StatusOK, sdkGoTestServicePrincipal))
	// One request per second with a burst of one: a classic read uses the only token,
	// so the sdk-go read that follows waits for the next one.
	c := newSdkGoTestClient(t, server,
		func(cfg *config.Config) { cfg.RateLimitPerSecond = 1 })
	ctx := context.Background()
	require.NoError(t, classicGetServicePrincipal(ctx, c))
	start := time.Now()
	_, err := getServicePrincipal(ctx, t, c, "")
	require.NoError(t, err)
	assert.GreaterOrEqual(t, time.Since(start), 800*time.Millisecond)
	assert.EqualValues(t, 2, calls.Load())
}

// blockingCredentials authenticate only when the request's context ends, like a token
// endpoint that doesn't answer.
type blockingCredentials struct{}

func (blockingCredentials) Name() string { return "blocking" }

func (blockingCredentials) Configure(context.Context,
	*config.Config) (credentials.CredentialsProvider, error) {
	return credentials.CredentialsProviderFn(func(r *http.Request) error {
		<-r.Context().Done()
		return r.Context().Err()
	}), nil
}

func TestSdkGoScimClientTimesOutAuthenticationLikeClassicClient(t *testing.T) {
	customize := func(cfg *config.Config) {
		cfg.HTTPTimeoutSeconds = 1
		cfg.Credentials = blockingCredentials{}
		cfg.AuthType = "blocking"
	}
	start := time.Now()
	r := readBothWays(t, context.Background, customize)
	assert.EqualError(t, r.err, "failed during request visitor: context canceled")
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.EqualValues(t, 0, r.calls)
	assert.Equal(t, r.classicCalls, r.calls)
	// Each client gives up on authentication after the one-second HTTP timeout.
	assert.Less(t, time.Since(start), 10*time.Second)
}
