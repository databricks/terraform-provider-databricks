package common

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/useragent"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSdkGoVolumesClientWorkspaceRequest(t *testing.T) {
	server, calls := serveSequence(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, sdkGoTestVolumePath, r.URL.Path)
		assert.Equal(t, sdkGoTestVolumePath, r.RequestURI,
			"no query, not even an empty one")
		assert.Equal(t, "12345", r.Header.Get("X-Databricks-Workspace-Id"))
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.NotEmpty(t, r.Header.Get("Traceparent"))
		assert.Empty(t, r.Header.Values("Accept"), "the existing client sends Accept")
		userAgent := r.Header.Get("User-Agent")
		assert.Contains(t, userAgent, "databricks-sdk-go/")
		assert.Contains(t, userAgent, "resource/volume")
		assert.Contains(t, userAgent, "auth/pat")
		assert.Contains(t, userAgent, " sdk-go-uc-volumes/0.8.0")
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
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := readBothWays(t, context.Background, nil, respond(tc.status, tc.body))
			require.Error(t, r.err)
			assert.Equal(t, tc.message, r.err.Error())
			assert.Equal(t, tc.missing, apierr.IsMissing(r.err))
			var apiErr *apierr.APIError
			require.ErrorAs(t, r.err, &apiErr)
			assert.Equal(t, tc.status, apiErr.StatusCode)
			assert.EqualError(t, r.classicErr, tc.message)
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
			r := readBothWays(t, context.Background, nil, respond(tc.status, tc.body),
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

// withoutHeaderLines drops the header lines of the request log in a decode error. The
// clients send different User-Agent, Accept and traceparent headers.
func withoutHeaderLines(err error) string {
	var lines []string
	for line := range strings.SplitSeq(err.Error(), "\n") {
		if !strings.HasPrefix(line, "> * ") && !strings.HasPrefix(line, "< * ") {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "\n")
}

// Both clients decode a successful response after their retries, so neither retries a
// body it can't decode, even one with a transient error message.
func TestSdkGoVolumesClientFailsToDecodeLikeClassicClient(t *testing.T) {
	for name, body := range map[string]string{
		"malformed JSON":          `{"full_name": "main.replay.volume",`,
		"not JSON":                "not json",
		"transient error message": `{"message": "REQUEST_LIMIT_EXCEEDED: retry later.`,
		"HTML":                    "<html><body>Sign in</body></html>",
		"timestamp that isn't a number": `{"full_name": "main.replay.volume",
			"created_at": "yesterday"}`,
		"timestamp out of range": `{"full_name": "main.replay.volume",
			"created_at": "99999999999999999999"}`,
		"timestamp with a transient error message": `{"full_name": "main.replay.volume",
			"created_at": "REQUEST_LIMIT_EXCEEDED"}`,
	} {
		t.Run(name, func(t *testing.T) {
			r := readBothWays(t, context.Background, nil, respond(http.StatusOK, body),
				respond(http.StatusOK, sdkGoTestVolume))
			assert.EqualValues(t, 1, r.calls)
			assert.Equal(t, r.classicCalls, r.calls)
			require.Error(t, r.err)
			require.Error(t, r.classicErr)
			assert.Equal(t, withoutHeaderLines(r.classicErr), withoutHeaderLines(r.err))
			assert.Equal(t, errors.Is(r.classicErr, httpclient.ErrHTMLContent),
				errors.Is(r.err, httpclient.ErrHTMLContent))
		})
	}
}

func TestSdkGoVolumesClientDecodeErrorPrintsTheRequestLog(t *testing.T) {
	r := readBothWays(t, context.Background, nil, respond(http.StatusOK, "not json"))
	assert.Equal(t, "failed to unmarshal response body: invalid character 'o' in "+
		"literal null (expecting 'u'). "+fmt.Sprintf(classicDecodeErrorRequestLog,
		"GET "+sdkGoTestVolumePath+"\n< HTTP/1.1 200 OK\n< not json"),
		withoutHeaderLines(r.err))
}

// The classic decoder reads a value it can't decode into a string field as that value's
// text, so a response with a number for a comment reads as a comment.
func TestSdkGoVolumesClientCoercesValuesLikeClassicClient(t *testing.T) {
	for name, tc := range map[string]struct {
		value, comment string
	}{
		"number":  {`123`, "123"},
		"boolean": {`true`, "true"},
		"object":  {`{"text":"x"}`, `{"text":"x"}`},
	} {
		t.Run(name, func(t *testing.T) {
			body := `{"full_name": "main.replay.volume", "created_at": "1700000000000",
				"comment": ` + tc.value + `}`
			server, _ := serveSequence(t, respond(http.StatusOK, body),
				respond(http.StatusOK, body))
			w := newVolumesTestWorkspaceClient(t, server, nil)
			ctx := context.Background()
			v, err := getVolume(ctx, t, w, sdkGoTestVolumeName)
			require.NoError(t, err)
			classic, err := w.Volumes.ReadByName(ctx, sdkGoTestVolumeName)
			require.NoError(t, err)
			assert.Equal(t, tc.comment, classic.Comment)
			require.NotNil(t, v.Comment)
			assert.Equal(t, classic.Comment, *v.Comment)
			require.NotNil(t, v.CreatedAt)
			assert.Equal(t, classic.CreatedAt, *v.CreatedAt)
		})
	}
}

// The existing client doesn't decode an empty 200 body, so it reads a volume without
// fields, and sdk-go reads the same.
func TestSdkGoVolumesClientReadsEmptySuccessBodyLikeClassicClient(t *testing.T) {
	server, _ := serveSequence(t, respond(http.StatusOK, ""), respond(http.StatusOK, ""))
	w := newVolumesTestWorkspaceClient(t, server, nil)
	ctx := context.Background()
	v, err := getVolume(ctx, t, w, sdkGoTestVolumeName)
	require.NoError(t, err)
	assert.Equal(t, &sdkvolumes.VolumeInfo{}, v)
	classic, err := w.Volumes.ReadByName(ctx, sdkGoTestVolumeName)
	require.NoError(t, err)
	assert.Equal(t, &catalog.VolumeInfo{}, classic)
}
