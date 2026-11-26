package common

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
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

func TestSdkGoVolumesClientWorkspaceRequest(t *testing.T) {
	server, calls := serveSequence(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, sdkGoTestVolumePath, r.URL.Path)
		assert.Equal(t, sdkGoTestVolumePath+"?", r.RequestURI, "the classic empty query")
		assert.Equal(t, "12345", r.Header.Get("X-Databricks-Workspace-Id"))
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		assert.NotEmpty(t, r.Header.Get("Traceparent"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))
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
		{"HTML error body with pre", 404,
			"<html><body><pre>Volume not found</pre></body></html>", "Volume not found",
			true},
		{"HTML error body without pre", 404, "<html><body>Volume not found</body></html>",
			"<html><body>Volume not found</body></html>", false},
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

// sentRequest is a request as the server received it, without its random Traceparent.
type sentRequest struct {
	method, requestURI, userAgent, body string
	header                              http.Header
}

func recordRequest(t *testing.T, r *http.Request) sentRequest {
	body, err := io.ReadAll(r.Body)
	require.NoError(t, err)
	header := r.Header.Clone()
	header.Del("Traceparent")
	header.Del("User-Agent")
	return sentRequest{r.Method, r.RequestURI, r.Header.Get("User-Agent"), string(body),
		header}
}

// The classic client formats the name into the URL unescaped, so in a name "#" drops the
// rest, "?" starts a query and "%41" is decoded. sdk-go reads send the same request for
// every name, and only add the sdk-go modules to the User-Agent.
func TestSdkGoVolumesClientSendsTheClassicRequestForEveryName(t *testing.T) {
	for _, name := range []string{"main.replay.volume", "main.replay.a#b",
		"main.replay.a?b", "main.replay.a%41", "main.replay.a b", "main.replay.a`b",
		"main.replay.caf\u00e9", "main.replay.a-b", "main.replay.A",
		"main.replay.a(b);c,d"} {
		t.Run(name, func(t *testing.T) {
			var sent []sentRequest
			record := func(w http.ResponseWriter, r *http.Request) {
				sent = append(sent, recordRequest(t, r))
				respond(http.StatusOK, sdkGoTestVolume)(w, r)
			}
			server, _ := serveSequence(t, record, record)
			w := newVolumesTestWorkspaceClient(t, server, nil)
			ctx := useragent.InContext(context.Background(), "resource", "volume")
			_, err := getVolume(ctx, t, w, name)
			require.NoError(t, err)
			_, err = w.Volumes.ReadByName(ctx, name)
			require.NoError(t, err)
			require.Len(t, sent, 2)
			sdkGo, classic := sent[0], sent[1]
			modules, ok := strings.CutPrefix(sdkGo.userAgent, classic.userAgent+" ")
			require.True(t, ok, "%q extends %q", sdkGo.userAgent, classic.userAgent)
			for _, module := range strings.Fields(modules) {
				assert.True(t, strings.HasPrefix(module, "sdk-go-"), module)
			}
			sdkGo.userAgent, classic.userAgent = "", ""
			assert.Equal(t, classic, sdkGo)
		})
	}
}

// A "%" that starts no valid escape fails to build the classic URL, so neither client
// authenticates or sends a request.
func TestSdkGoVolumesClientFailsUnparsableNamesLikeClassicClient(t *testing.T) {
	server, calls := serveSequence(t)
	w := newVolumesTestWorkspaceClient(t, server, nil)
	ctx := context.Background()
	_, err := getVolume(ctx, t, w, "main.replay.50%_raw")
	assert.EqualError(t, err, `failed creating new request: parse `+
		`"/api/2.1/unity-catalog/volumes/main.replay.50%_raw?": invalid URL escape "%_r"`)
	_, classicErr := w.Volumes.ReadByName(ctx, "main.replay.50%_raw")
	assert.EqualError(t, classicErr, err.Error())
	assert.EqualValues(t, 0, calls.Load())
}

// A classic request's body hook replaces the body sdk-go built on every attempt.
func TestSdkGoVolumesClientSendsTheClassicBody(t *testing.T) {
	var bodies []string
	record := func(status int, body string) handler {
		return func(w http.ResponseWriter, r *http.Request) {
			sent := recordRequest(t, r)
			bodies = append(bodies, sent.body)
			assert.Equal(t, r.ContentLength, int64(len(sent.body)))
			respond(status, body)(w, r)
		}
	}
	server, _ := serveSequence(t, record(429, sdkGoTestTooManyRequests),
		record(http.StatusOK, sdkGoTestVolume))
	ctx := context.Background()
	vc, err := SdkGoVolumesClient(ctx, newVolumesTestWorkspaceClient(t, server, nil))
	require.NoError(t, err)
	classic := classicVolumeRequest(http.MethodPatch, sdkGoTestVolumePath,
		sdkGoTestVolumePath)
	classic.body = func(body []byte) ([]byte, error) {
		var fields map[string]any
		if err := json.Unmarshal(body, &fields); err != nil {
			return nil, err
		}
		delete(fields, "full_name_arg")
		return json.Marshal(fields)
	}
	name, comment := sdkGoTestVolumeName, "updated"
	update := func(ctx context.Context,
		opts ...call.Option) (*sdkvolumes.VolumeInfo, error) {
		req := sdkvolumes.UpdateVolumeRequest{FullNameArg: &name, Comment: &comment}
		return vc.client.UpdateVolume(ctx, req, opts...)
	}
	_, err = callSdkGo(ctx, vc.bridge, classic, update)
	require.NoError(t, err)
	assert.Equal(t, []string{`{"comment":"updated"}`, `{"comment":"updated"}`}, bodies)
}

func TestSdkGoAttemptRejectsAnotherRequestThanTheClassicOne(t *testing.T) {
	classic := classicVolumeRequest(http.MethodGet, sdkGoTestVolumePath,
		sdkGoTestVolumePath+"?")
	attempt := &sdkGoAttempt{classic: classic}
	other := sdkGoTestVolumePath + "_other"
	r, err := http.NewRequest(http.MethodGet, "https://example.com"+other, nil)
	require.NoError(t, err)
	assert.EqualError(t, attempt.useClassicRequest(r), "sdk-go requested GET "+other+
		" instead of the classic GET "+sdkGoTestVolumePath)
}

// requestLogHeaders matches the request log lines whose values differ between the two
// clients: the User-Agent, by the sdk-go modules, and the random Traceparent.
var requestLogHeaders = regexp.MustCompile(`(?m)^> \* (User-Agent|Traceparent): .*$`)

// A success body is decoded as the classic client decodes it: an empty body is an empty
// volume, a body starting with "<" fails with the HTML error, and another body that
// doesn't decode fails with the classic type names and request log.
func TestSdkGoVolumesClientDecodesSuccessBodiesLikeClassicClient(t *testing.T) {
	for _, tc := range []struct{ name, body, message string }{
		{"empty body", "", ""},
		{"HTML", "<html><body>Sign in</body></html>",
			"received HTML response instead of JSON"},
		{"malformed JSON", `{"full_name": "main.replay.volume",`,
			"failed to unmarshal response body: unexpected end of JSON input. This is " +
				"likely a bug in the Databricks SDK for Go"},
		{"unparsable timestamp", `{"full_name": "main.replay.volume", "created_at": "x"}`,
			`failed to unmarshal response body: strconv.ParseInt: parsing "x": invalid ` +
				"syntax. This is likely a bug"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var v *sdkvolumes.VolumeInfo
			var err, classicErr error
			ctx := context.Background()
			server, calls := serveSequence(t, respond(http.StatusOK, tc.body))
			v, err = getVolume(ctx, t, newVolumesTestWorkspaceClient(t, server, nil),
				sdkGoTestVolumeName)
			classicServer, _ := serveSequence(t, respond(http.StatusOK, tc.body))
			classic := newVolumesTestWorkspaceClient(t, classicServer, nil)
			_, classicErr = classic.Volumes.ReadByName(ctx, sdkGoTestVolumeName)
			assert.EqualValues(t, 1, calls.Load())
			if tc.message == "" {
				require.NoError(t, err)
				require.NoError(t, classicErr)
				assert.Equal(t, sdkvolumes.VolumeInfo{}, *v)
				return
			}
			require.Error(t, err)
			require.Error(t, classicErr)
			assert.True(t, strings.HasPrefix(err.Error(), tc.message), err.Error())
			normalize := func(err error) string {
				return requestLogHeaders.ReplaceAllString(err.Error(), "> * $1: ...")
			}
			assert.Equal(t, normalize(classicErr), normalize(err))
		})
	}
}

// The classic client reads a JSON value of another type as a string field's text, such
// as 7 as "7", and a quoted number as an integer field, so sdk-go reads the same values.
func TestSdkGoVolumesClientCoercesValuesLikeClassicClient(t *testing.T) {
	for name, tc := range map[string]struct{ value, comment string }{
		"number":  {`7`, "7"},
		"boolean": {`true`, "true"},
		"object":  {`{"text": "x"}`, `{"text": "x"}`},
	} {
		t.Run(name, func(t *testing.T) {
			body := `{"full_name": 7, "created_at": "1700000000000", "comment": ` +
				tc.value + `}`
			ctx := context.Background()
			server, calls := serveSequence(t, respond(http.StatusOK, body))
			v, err := getVolume(ctx, t, newVolumesTestWorkspaceClient(t, server, nil),
				sdkGoTestVolumeName)
			require.NoError(t, err)
			classicServer, _ := serveSequence(t, respond(http.StatusOK, body))
			classic, err := newVolumesTestWorkspaceClient(t, classicServer, nil).
				Volumes.ReadByName(ctx, sdkGoTestVolumeName)
			require.NoError(t, err)
			assert.EqualValues(t, 1, calls.Load())
			assert.Equal(t, "7", classic.FullName)
			assert.Equal(t, tc.comment, classic.Comment)
			require.NotNil(t, v.FullName)
			require.NotNil(t, v.Comment)
			require.NotNil(t, v.CreatedAt)
			assert.Equal(t, classic.FullName, *v.FullName)
			assert.Equal(t, classic.Comment, *v.Comment)
			assert.Equal(t, classic.CreatedAt, *v.CreatedAt)
		})
	}
}

// The classic client decodes a success body after its retries, so a body that doesn't
// decode isn't retried, even when its request log holds a transient error message.
func TestSdkGoVolumesClientDoesNotRetryUndecodableBodies(t *testing.T) {
	body := `{"message": "REQUEST_LIMIT_EXCEEDED: retry later",`
	// A retry would send a second request within this timeout.
	customize := func(cfg *config.Config) { cfg.RetryTimeoutSeconds = 5 }
	r := readBothWays(t, context.Background, customize, respond(http.StatusOK, body))
	require.Error(t, r.err)
	assert.Contains(t, r.err.Error(), "REQUEST_LIMIT_EXCEEDED")
	assert.EqualValues(t, 1, r.calls)
	assert.Equal(t, r.classicCalls, r.calls)
	normalize := func(err error) string {
		return requestLogHeaders.ReplaceAllString(err.Error(), "> * $1: ...")
	}
	assert.Equal(t, normalize(r.classicErr), normalize(r.err))
}

// A host that doesn't parse fails each call in the classic host visitor, after the rate
// limiter and authentication, so neither client sends a request.
func TestSdkGoVolumesClientFailsHostsThatDontParseLikeClassicClient(t *testing.T) {
	badHost := func(cfg *config.Config) { cfg.Host = "https://bad host.example.com" }
	r := readBothWays(t, context.Background, badHost)
	assert.EqualError(t, r.err, `failed during request visitor: parse `+
		`"https://bad host.example.com": invalid character " " in host name`)
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.EqualValues(t, 0, r.calls)
	assert.EqualValues(t, 0, r.classicCalls)
}

// switchingProtocols answers with a 101 that no request asked for, then closes the
// connection, which ends the response body.
func switchingProtocols(w http.ResponseWriter, _ *http.Request) {
	conn, buf, err := w.(http.Hijacker).Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	_, _ = buf.WriteString("HTTP/1.1 101 Switching Protocols\r\n" +
		"Connection: Upgrade\r\nUpgrade: websocket\r\n\r\n")
	_ = buf.Flush()
}

// The classic client reads every status below 400 as a success, so sdk-go reads the
// same volume for each, and a delete answered with one succeeds on both.
func TestSdkGoVolumesClientReadsSuccessStatusesLikeClassicClient(t *testing.T) {
	for name, answer := range map[string]handler{
		"101 Switching Protocols": switchingProtocols,
		"202 Accepted":            respond(http.StatusAccepted, sdkGoTestVolume),
		"304 Not Modified":        respond(http.StatusNotModified, ""),
	} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			server, calls := serveSequence(t, answer, answer)
			vc, err := SdkGoVolumesClient(ctx, newVolumesTestWorkspaceClient(t, server, nil))
			require.NoError(t, err)
			v, err := vc.ReadByName(ctx, sdkGoTestVolumeName)
			require.NoError(t, err)
			require.NoError(t, vc.DeleteByName(ctx, sdkGoTestVolumeName))
			classicServer, classicCalls := serveSequence(t, answer, answer)
			classic := newVolumesTestWorkspaceClient(t, classicServer, nil).Volumes
			classicVolume, err := classic.ReadByName(ctx, sdkGoTestVolumeName)
			require.NoError(t, err)
			require.NoError(t, classic.DeleteByName(ctx, sdkGoTestVolumeName))
			var fullName string
			if v.FullName != nil {
				fullName = *v.FullName
			}
			assert.Equal(t, classicVolume.FullName, fullName)
			assert.EqualValues(t, 2, calls.Load())
			assert.Equal(t, classicCalls.Load(), calls.Load())
		})
	}
}
