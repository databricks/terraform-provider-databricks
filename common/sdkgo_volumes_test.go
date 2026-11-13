package common

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/sdk-go/options/call"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	sdkGoTestVolumeName = "main.replay.volume"
	sdkGoTestVolumePath = "/api/2.1/unity-catalog/volumes/" + sdkGoTestVolumeName
	sdkGoTestVolume     = `{"name": "volume", "catalog_name": "main",
		"schema_name": "replay", "full_name": "main.replay.volume",
		"volume_type": "MANAGED"}`
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

// readVolumeBothWays reads the volume through sdk-go and the existing client, each
// against a fresh server with the same handlers.
func readVolumeBothWays(t *testing.T, handlers ...handler) bothWays {
	t.Helper()
	var result bothWays
	ctx := context.Background()
	server, calls := serveSequence(t, handlers...)
	_, result.err = getVolume(ctx, t, newVolumesTestWorkspaceClient(t, server, nil),
		sdkGoTestVolumeName)
	classicServer, classicCalls := serveSequence(t, handlers...)
	w := newVolumesTestWorkspaceClient(t, classicServer, nil)
	_, result.classicErr = w.Volumes.ReadByName(ctx, sdkGoTestVolumeName)
	result.calls, result.classicCalls = calls.Load(), classicCalls.Load()
	return result
}

func TestSdkGoVolumesClientWorkspaceRequest(t *testing.T) {
	server, calls := serveSequence(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, sdkGoTestVolumePath, r.URL.Path)
		assert.Equal(t, sdkGoTestVolumePath, r.RequestURI, "no query, not even an empty one")
		assert.Equal(t, "12345", r.Header.Get("X-Databricks-Workspace-Id"))
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.NotEmpty(t, r.Header.Get("Traceparent"))
		assert.Empty(t, r.Header.Values("Accept"), "the existing client sends Accept")
		userAgent := r.Header.Get("User-Agent")
		assert.Contains(t, userAgent, "databricks-sdk-go/")
		assert.Contains(t, userAgent, "resource/volume")
		assert.Contains(t, userAgent, "auth/pat")
		assert.Contains(t, userAgent, " sdk-go-uc-volumes/0.7.0")
		respond(http.StatusOK, sdkGoTestVolume)(w, r)
	})
	ctx := useragent.InContext(context.Background(), "resource", "volume")
	v, err := getVolume(ctx, t, newVolumesTestWorkspaceClient(t, server, nil),
		sdkGoTestVolumeName)
	require.NoError(t, err)
	assert.Equal(t, sdkvolumes.VolumeType_Managed, v.VolumeType)
	assert.EqualValues(t, 1, calls.Load())
}

// The provider resolves a workspace client per workspace_id, which can have another
// host, workspace ID and credentials than the provider's own client.
func TestSdkGoVolumesClientUsesTheWorkspaceClientsConfig(t *testing.T) {
	providerServer, providerCalls := serveSequence(t)
	workspaceServer, workspaceCalls := serveSequence(t,
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "67890", r.Header.Get("X-Databricks-Workspace-Id"))
			assert.Equal(t, "Bearer workspace-token", r.Header.Get("Authorization"))
			respond(http.StatusOK, sdkGoTestVolume)(w, r)
		})
	newSdkGoTestClient(t, providerServer, nil)
	w := newVolumesTestWorkspaceClient(t, workspaceServer, func(cfg *config.Config) {
		cfg.WorkspaceID = "67890"
		cfg.Token = "workspace-token"
	})
	_, err := getVolume(context.Background(), t, w, sdkGoTestVolumeName)
	require.NoError(t, err)
	assert.EqualValues(t, 0, providerCalls.Load())
	assert.EqualValues(t, 1, workspaceCalls.Load())
}

func TestSdkGoVolumesClientSharesTheWorkspaceClientsRateLimit(t *testing.T) {
	server, calls := serveSequence(t, respond(http.StatusOK, sdkGoTestVolume),
		respond(http.StatusOK, sdkGoTestVolume))
	// One request per second with a burst of one: a classic read through w uses the only
	// token, so the sdk-go read that follows waits for the next one.
	w := newVolumesTestWorkspaceClient(t, server,
		func(cfg *config.Config) { cfg.RateLimitPerSecond = 1 })
	ctx := context.Background()
	_, err := w.Volumes.ReadByName(ctx, sdkGoTestVolumeName)
	require.NoError(t, err)
	start := time.Now()
	_, err = getVolume(ctx, t, w, sdkGoTestVolumeName)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, time.Since(start), 800*time.Millisecond)
	assert.EqualValues(t, 2, calls.Load())
}

func TestSdkGoVolumesClientRejectsMockWorkspaceClients(t *testing.T) {
	_, err := SdkGoVolumesClient(context.Background(), &databricks.WorkspaceClient{})
	assert.EqualError(t, err, "*databricks.WorkspaceClient has no API client to share")
}

func TestSdkGoVolumesClientMapsErrorsLikeClassicClient(t *testing.T) {
	for _, tc := range []struct {
		name    string
		status  int
		body    string
		message string
		missing bool
	}{
		{"not found", 404, `{"error_code": "RESOURCE_DOES_NOT_EXIST",
			"message": "Volume 'main.replay.volume' does not exist."}`,
			"Volume 'main.replay.volume' does not exist.", true},
		{"forbidden", 403, `{"error_code": "PERMISSION_DENIED",
			"message": "User does not have READ VOLUME on Volume."}`,
			"User does not have READ VOLUME on Volume.", false},
		{"non-JSON error body", 500, "upstream connect error",
			"upstream connect error", false},
		{"undecodable success body", 200, "not json",
			"failed to unmarshal response body: " +
				"invalid character 'o' in literal null (expecting 'u')", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := readVolumeBothWays(t, respond(tc.status, tc.body))
			require.Error(t, r.err)
			assert.Equal(t, tc.message, r.err.Error())
			assert.Equal(t, tc.missing, apierr.IsMissing(r.err))
			require.Error(t, r.classicErr)
			// The classic decode error also prints the request and response.
			classicMessage := r.classicErr.Error()
			assert.True(t, strings.HasPrefix(classicMessage, tc.message), classicMessage)
			assert.Equal(t, apierr.IsMissing(r.classicErr), apierr.IsMissing(r.err))
		})
	}
}

func TestSdkGoVolumesClientRetriesLikeClassicClient(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		body   string
		calls  int32
	}{
		{"429 is retried", 429, `{"error_code": "TOO_MANY_REQUESTS",
			"message": "Too many requests."}`, 2},
		{"503 is retried", 503, `{"error_code": "TEMPORARILY_UNAVAILABLE",
			"message": "Unavailable."}`, 2},
		{"rate limit error code is retried", 400, `{"error_code": "RESOURCE_EXHAUSTED",
			"message": "Quota exceeded."}`, 2},
		{"transient message is retried", 400, `{"error_code": "BAD_REQUEST",
			"message": "REQUEST_LIMIT_EXCEEDED: retry later."}`, 2},
		{"504 is not retried", 504, "upstream request timeout", 1},
		{"500 is not retried", 500, `{"error_code": "INTERNAL_ERROR",
			"message": "Internal error."}`, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := readVolumeBothWays(t, respond(tc.status, tc.body),
				respond(http.StatusOK, sdkGoTestVolume))
			assert.Equal(t, tc.calls, r.calls)
			assert.Equal(t, tc.calls == 1, r.err != nil, "error: %v", r.err)
			assert.Equal(t, r.classicCalls, r.calls)
			if r.err != nil {
				assert.EqualError(t, r.classicErr, r.err.Error())
			}
		})
	}
}

// Volume names may contain "#", "?" and "%". sdk-go escapes them in the path, while the
// existing client parses them as a fragment, a query or an escape.
func TestSdkGoVolumesClientEscapesReservedCharactersInNames(t *testing.T) {
	for _, tc := range []struct {
		name, sdkGoURI, classicURI string
	}{
		{"main.replay.a#b", "/api/2.1/unity-catalog/volumes/main.replay.a%23b",
			"/api/2.1/unity-catalog/volumes/main.replay.a"},
		{"main.replay.a?b", "/api/2.1/unity-catalog/volumes/main.replay.a%3Fb",
			"/api/2.1/unity-catalog/volumes/main.replay.a?b?"},
		{"main.replay.a%41", "/api/2.1/unity-catalog/volumes/main.replay.a%2541",
			"/api/2.1/unity-catalog/volumes/main.replay.a%41?"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var uris []string
			record := func(w http.ResponseWriter, r *http.Request) {
				uris = append(uris, r.RequestURI)
				respond(http.StatusOK, sdkGoTestVolume)(w, r)
			}
			server, _ := serveSequence(t, record, record)
			w := newVolumesTestWorkspaceClient(t, server, nil)
			ctx := context.Background()
			_, err := getVolume(ctx, t, w, tc.name)
			require.NoError(t, err)
			_, err = w.Volumes.ReadByName(ctx, tc.name)
			require.NoError(t, err)
			assert.Equal(t, []string{tc.sdkGoURI, tc.classicURI}, uris)
		})
	}
}
