package common

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/sdk-go/options/call"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
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

const (
	sdkGoTestVolumeName = "main.replay.volume"
	sdkGoTestVolumePath = "/api/2.1/unity-catalog/volumes/" + sdkGoTestVolumeName
	sdkGoTestVolume     = `{"name": "volume", "catalog_name": "main",
		"schema_name": "replay", "full_name": "main.replay.volume",
		"volume_type": "MANAGED"}`
	sdkGoTestTooManyRequests = `{"error_code": "TOO_MANY_REQUESTS",
		"message": "Too many requests."}`
)

// newVolumesTestWorkspaceClient returns the workspace client the provider resolves for
// a workspace-level provider configured for server.
func newVolumesTestWorkspaceClient(t *testing.T, server *httptest.Server,
	customize func(*config.Config)) *databricks.WorkspaceClient {
	t.Helper()
	w, err := newSdkGoTestClient(t, server, customize).WorkspaceClient()
	require.NoError(t, err)
	return w
}

func getVolume(ctx context.Context, t *testing.T, w *databricks.WorkspaceClient,
	fullName string) (*sdkvolumes.VolumeInfo, error) {
	t.Helper()
	vc, err := SdkGoVolumesClient(ctx, w)
	require.NoError(t, err)
	get := func(ctx context.Context,
		opts ...call.Option) (*sdkvolumes.VolumeInfo, error) {
		req := sdkvolumes.GetVolumeRequest{FullNameArg: &fullName}
		return vc.GetVolume(ctx, req, opts...)
	}
	return CallSdkGoVolumes(ctx, vc, get)
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

type bothWays struct {
	err, classicErr     error
	calls, classicCalls int32
}

// readBothWays reads the volume through sdk-go and the existing client, each against a
// fresh server with the same handlers.
func readBothWays(t *testing.T, ctx func() context.Context,
	customize func(*config.Config), handlers ...handler) bothWays {
	t.Helper()
	var result bothWays
	server, calls := serveSequence(t, handlers...)
	w := newVolumesTestWorkspaceClient(t, server, customize)
	_, result.err = getVolume(ctx(), t, w, sdkGoTestVolumeName)
	classicServer, classicCalls := serveSequence(t, handlers...)
	classic := newVolumesTestWorkspaceClient(t, classicServer, customize)
	_, result.classicErr = classic.Volumes.ReadByName(ctx(), sdkGoTestVolumeName)
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

func TestSdkGoClientAppliesConfigHeaders(t *testing.T) {
	server, _ := serveSequence(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "extra", r.Header.Get("X-Extra"))
		respond(http.StatusOK, sdkGoTestVolume)(w, r)
	})
	w := newVolumesTestWorkspaceClient(t, server, func(cfg *config.Config) {
		cfg.Headers = func(r *http.Request) error {
			r.Header.Set("X-Extra", "extra")
			return nil
		}
	})
	_, err := getVolume(context.Background(), t, w, sdkGoTestVolumeName)
	require.NoError(t, err)
}

func TestSdkGoClientStopsRetryingAtRetryTimeout(t *testing.T) {
	tooMany := respond(429, sdkGoTestTooManyRequests)
	// The first classic backoff is at least one second, past the retry timeout.
	r := readBothWays(t, context.Background,
		func(cfg *config.Config) { cfg.RetryTimeoutSeconds = 1 }, tooMany, tooMany)
	assert.EqualError(t, r.err, "Too many requests.")
	assert.True(t, errors.Is(r.err, apierr.ErrTooManyRequests))
	assert.EqualValues(t, 1, r.calls)
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.Equal(t, r.classicCalls, r.calls)
}

// The classic client bounds only its backoff by the retry timeout, so a response that
// arrives after the timeout still succeeds there and must succeed through sdk-go too.
func TestSdkGoClientWaitsForAResponsePastRetryTimeoutLikeClassicClient(t *testing.T) {
	slow := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1500 * time.Millisecond)
		respond(http.StatusOK, sdkGoTestVolume)(w, r)
	}
	r := readBothWays(t, context.Background, func(cfg *config.Config) {
		cfg.RetryTimeoutSeconds = 1
		cfg.HTTPTimeoutSeconds = 10
	}, slow)
	assert.NoError(t, r.err)
	assert.EqualValues(t, 1, r.calls)
	assert.NoError(t, r.classicErr)
	assert.Equal(t, r.classicCalls, r.calls)
}

// htmlError answers like the legacy APIs that return an HTML page, which the classic
// client classifies by the response's status text.
func htmlError(status int, message string) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		respond(status, "<html><body><pre>"+message+"</pre></body></html>")(w, r)
	}
}

func TestSdkGoClientReadsHTMLNotFoundAsMissingLikeClassicClient(t *testing.T) {
	r := readBothWays(t, context.Background, nil,
		htmlError(http.StatusNotFound, "Volume not found."))
	var apiErr, classicAPIErr *apierr.APIError
	require.ErrorAs(t, r.err, &apiErr)
	require.ErrorAs(t, r.classicErr, &classicAPIErr)
	assert.Equal(t, "NOT_FOUND", classicAPIErr.ErrorCode)
	assert.Equal(t, classicAPIErr.ErrorCode, apiErr.ErrorCode)
	assert.True(t, apierr.IsMissing(r.err))
	assert.EqualError(t, r.err, r.classicErr.Error())
	assert.Equal(t, r.classicCalls, r.calls)
}

func TestSdkGoClientRetriesHTMLTooManyRequestsLikeClassicClient(t *testing.T) {
	r := readBothWays(t, context.Background, nil,
		htmlError(http.StatusTooManyRequests, "Too many requests."),
		respond(http.StatusOK, sdkGoTestVolume))
	assert.NoError(t, r.classicErr)
	assert.NoError(t, r.err)
	assert.EqualValues(t, 2, r.calls)
	assert.Equal(t, r.classicCalls, r.calls)
}

// The classic client decodes a successful response only after its retries, so a write
// whose response fails to decode isn't sent again, even when the decode error contains
// a transient error message.
func TestSdkGoClientDoesNotRetryUndecodableSuccessfulResponses(t *testing.T) {
	undecodable := respond(http.StatusOK, `{"full_name": "main.replay.volume",
		"created_at": "REQUEST_LIMIT_EXCEEDED"}`)
	ctx := context.Background()
	name, catalogName, schemaName := "volume", "main", "replay"
	server, calls := serveSequence(t, undecodable)
	vc, err := SdkGoVolumesClient(ctx, newVolumesTestWorkspaceClient(t, server, nil))
	require.NoError(t, err)
	create := func(ctx context.Context,
		opts ...call.Option) (*sdkvolumes.VolumeInfo, error) {
		return vc.CreateVolume(ctx, sdkvolumes.CreateVolumeRequest{
			CatalogName: &catalogName, SchemaName: &schemaName, Name: &name,
			VolumeType: sdkvolumes.VolumeType_Managed,
		}, opts...)
	}
	_, err = CallSdkGoVolumes(ctx, vc, create)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal response body")
	assert.EqualValues(t, 1, calls.Load())
	classicServer, classicCalls := serveSequence(t, undecodable)
	classic := newVolumesTestWorkspaceClient(t, classicServer, nil)
	_, err = classic.Volumes.Create(ctx, catalog.CreateVolumeRequestContent{
		CatalogName: catalogName, SchemaName: schemaName, Name: name,
		VolumeType: catalog.VolumeTypeManaged,
	})
	require.Error(t, err)
	assert.EqualValues(t, 1, classicCalls.Load())
}

func TestSdkGoClientReturnsLastErrorWhenCancelledDuringBackoff(t *testing.T) {
	var cancel context.CancelFunc
	newContext := func() context.Context {
		var ctx context.Context
		ctx, cancel = context.WithCancel(context.Background())
		return ctx
	}
	tooMany := func(w http.ResponseWriter, r *http.Request) {
		time.AfterFunc(200*time.Millisecond, cancel)
		respond(429, sdkGoTestTooManyRequests)(w, r)
	}
	r := readBothWays(t, newContext, nil, tooMany)
	assert.EqualError(t, r.err, "Too many requests.")
	assert.EqualValues(t, 1, r.calls)
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.Equal(t, r.classicCalls, r.calls)
}

func TestSdkGoClientTimesOutLikeClassicClient(t *testing.T) {
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

func TestSdkGoClientReportsRefusedConnectionsLikeClassicClient(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	server.Close()
	// A refused connection is retried, so a short retry timeout ends the call.
	customize := func(cfg *config.Config) { cfg.RetryTimeoutSeconds = 1 }
	ctx := context.Background()
	w := newVolumesTestWorkspaceClient(t, server, customize)
	_, err := getVolume(ctx, t, w, sdkGoTestVolumeName)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "connection refused")
	classic := newVolumesTestWorkspaceClient(t, server, customize)
	_, classicErr := classic.Volumes.ReadByName(ctx, sdkGoTestVolumeName)
	// The existing client sends an empty query, so its URL ends in "?".
	classicMessage := strings.Replace(err.Error(), sdkGoTestVolumePath+`"`,
		sdkGoTestVolumePath+`?"`, 1)
	assert.EqualError(t, classicErr, classicMessage)
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

func TestSdkGoClientTimesOutAuthenticationLikeClassicClient(t *testing.T) {
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

// refusedTokenCredentials fail like a token endpoint that refuses connections, which
// the classic client would retry if it came from an API call.
type refusedTokenCredentials struct {
	calls *atomic.Int32
}

func (refusedTokenCredentials) Name() string { return "refused-token" }

func (c refusedTokenCredentials) Configure(context.Context,
	*config.Config) (credentials.CredentialsProvider, error) {
	return credentials.CredentialsProviderFn(func(r *http.Request) error {
		c.calls.Add(1)
		return &url.Error{Op: "Post", URL: "https://login.example.com/token",
			Err: errors.New("dial tcp 127.0.0.1:1: connect: connection refused")}
	}), nil
}

func TestSdkGoClientDoesNotRetryAuthenticationFailures(t *testing.T) {
	var sdkGoAuth, classicAuth atomic.Int32
	withCredentials := func(calls *atomic.Int32) func(*config.Config) {
		return func(cfg *config.Config) {
			// A retried failure would make a second attempt within this timeout.
			cfg.RetryTimeoutSeconds = 5
			cfg.Credentials = refusedTokenCredentials{calls}
			cfg.AuthType = "refused-token"
		}
	}
	server, calls := serveSequence(t)
	ctx := context.Background()
	w := newVolumesTestWorkspaceClient(t, server, withCredentials(&sdkGoAuth))
	_, err := getVolume(ctx, t, w, sdkGoTestVolumeName)
	classic := newVolumesTestWorkspaceClient(t, server, withCredentials(&classicAuth))
	_, classicErr := classic.Volumes.ReadByName(ctx, sdkGoTestVolumeName)

	assert.EqualError(t, err, `failed during request visitor: Post `+
		`"https://login.example.com/token": `+
		`dial tcp 127.0.0.1:1: connect: connection refused`)
	assert.EqualError(t, classicErr, err.Error())
	assert.EqualValues(t, 1, sdkGoAuth.Load())
	assert.Equal(t, classicAuth.Load(), sdkGoAuth.Load())
	assert.EqualValues(t, 0, calls.Load())
}

func TestSdkGoClientTimesOutRequestVisitorsLikeClassicClient(t *testing.T) {
	customize := func(cfg *config.Config) {
		cfg.HTTPTimeoutSeconds = 1
		cfg.Headers = func(r *http.Request) error {
			<-r.Context().Done()
			return r.Context().Err()
		}
	}
	start := time.Now()
	r := readBothWays(t, context.Background, customize)
	assert.EqualError(t, r.err, "failed during request visitor: context canceled")
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.EqualValues(t, 0, r.calls)
	assert.Equal(t, r.classicCalls, r.calls)
	// Each client gives up on the visitor after the one-second HTTP timeout.
	assert.Less(t, time.Since(start), 10*time.Second)
}
