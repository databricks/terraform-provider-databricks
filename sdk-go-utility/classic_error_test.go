package sdkgoutility

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth/credentials"
	modularerr "github.com/databricks/sdk-go/core/apierr"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// classicErrorSentinels are the sentinel errors of the classic apierr package.
var classicErrorSentinels = []error{
	apierr.ErrBadRequest, apierr.ErrUnauthenticated, apierr.ErrPermissionDenied, apierr.ErrNotFound,
	apierr.ErrResourceConflict, apierr.ErrTooManyRequests, apierr.ErrCancelled, apierr.ErrInternalError,
	apierr.ErrNotImplemented, apierr.ErrTemporarilyUnavailable, apierr.ErrDeadlineExceeded,
	apierr.ErrInvalidState, apierr.ErrInvalidParameterValue, apierr.ErrResourceDoesNotExist,
	apierr.ErrAborted, apierr.ErrAlreadyExists, apierr.ErrResourceAlreadyExists, apierr.ErrResourceExhausted,
	apierr.ErrRequestLimitExceeded, apierr.ErrUnknown, apierr.ErrDataLoss,
}

const (
	classicErrorHTMLWithPre    = `<html><body><h2>HTTP ERROR</h2><pre> Worker environment not ready. </pre></body></html>`
	classicErrorHTMLWithoutPre = `<html><head><title>Bad Gateway</title></head><body><center>nginx</center></body></html>`
	classicErrorTruncatedJSON  = `{"error_code":"RESOURCE_DOES_NOT_EXIST","mess`
	classicErrorWithDetails    = `{"error_code":"INVALID_PARAMETER_VALUE","message":"Invalid name","details":[` +
		`{"@type":"type.googleapis.com/google.rpc.ErrorInfo","reason":"INVALID_NAME","domain":"mlflow","metadata":{"name":"a/b"}},` +
		`{"@type":"type.googleapis.com/google.rpc.RequestInfo","request_id":"r-1","serving_data":"s"},` +
		`{"@type":"type.googleapis.com/google.rpc.RetryInfo","retry_delay":"3s"},{"@type":"unknown","x":1},"not an object"]}`
)

// classicErrorResponses are failed responses with what the classic client
// makes of them: the error code, the message, whether apierr.IsMissing holds
// and whether the client retries the call before it returns the error.
var classicErrorResponses = []struct {
	name          string
	status        int
	body          string
	code, message string
	missing       bool
	retried       bool
}{
	// A JSON error_code and message on every status that the classic client has a sentinel error for.
	{name: "400", status: 400, body: `{"error_code":"BAD_REQUEST","message":"bad"}`, code: "BAD_REQUEST", message: "bad"},
	{name: "401", status: 401, body: `{"error_code":"UNAUTHENTICATED","message":"who"}`, code: "UNAUTHENTICATED", message: "who"},
	{name: "403", status: 403, body: `{"error_code":"PERMISSION_DENIED","message":"no"}`, code: "PERMISSION_DENIED", message: "no"},
	{name: "404", status: 404, body: `{"error_code":"NOT_FOUND","message":"gone"}`, code: "NOT_FOUND", message: "gone", missing: true},
	{name: "409", status: 409, body: `{"error_code":"RESOURCE_CONFLICT","message":"busy"}`, code: "RESOURCE_CONFLICT", message: "busy"},
	{name: "429", status: 429, body: `{"error_code":"REQUEST_LIMIT_EXCEEDED","message":"slow"}`, code: "REQUEST_LIMIT_EXCEEDED", message: "slow", retried: true},
	{name: "499", status: 499, body: `{"error_code":"CANCELLED","message":"stop"}`, code: "CANCELLED", message: "stop"},
	{name: "500", status: 500, body: `{"error_code":"INTERNAL_ERROR","message":"oops"}`, code: "INTERNAL_ERROR", message: "oops"},
	{name: "501", status: 501, body: `{"error_code":"NOT_IMPLEMENTED","message":"later"}`, code: "NOT_IMPLEMENTED", message: "later"},
	{name: "503", status: 503, body: `{"error_code":"TEMPORARILY_UNAVAILABLE","message":"wait"}`, code: "TEMPORARILY_UNAVAILABLE", message: "wait", retried: true},
	{name: "504", status: 504, body: `{"error_code":"DEADLINE_EXCEEDED","message":"late"}`, code: "DEADLINE_EXCEEDED", message: "late"},
	// A status without a sentinel error.
	{name: "405", status: 405, body: `{"error_code":"METHOD_NOT_ALLOWED","message":"no"}`, code: "METHOD_NOT_ALLOWED", message: "no"},
	{name: "502", status: 502, body: `{"error_code":"BAD_GATEWAY","message":"down"}`, code: "BAD_GATEWAY", message: "down"},
	// The error code decides before the status does.
	{name: "404 RESOURCE_DOES_NOT_EXIST", status: 404, body: `{"error_code":"RESOURCE_DOES_NOT_EXIST","message":"gone"}`, code: "RESOURCE_DOES_NOT_EXIST", message: "gone", missing: true},
	{name: "404 FEATURE_DISABLED", status: 404, body: `{"error_code":"FEATURE_DISABLED","message":"off"}`, code: "FEATURE_DISABLED", message: "off", missing: true},
	{name: "404 INVALID_PARAMETER_VALUE", status: 404, body: `{"error_code":"INVALID_PARAMETER_VALUE","message":"bad"}`, code: "INVALID_PARAMETER_VALUE", message: "bad"},
	{name: "400 RESOURCE_DOES_NOT_EXIST", status: 400, body: `{"error_code":"RESOURCE_DOES_NOT_EXIST","message":"gone"}`, code: "RESOURCE_DOES_NOT_EXIST", message: "gone", missing: true},
	{name: "400 RESOURCE_ALREADY_EXISTS", status: 400, body: `{"error_code":"RESOURCE_ALREADY_EXISTS","message":"dup"}`, code: "RESOURCE_ALREADY_EXISTS", message: "dup"},
	{name: "400 INVALID_STATE", status: 400, body: `{"error_code":"INVALID_STATE","message":"busy"}`, code: "INVALID_STATE", message: "busy"},
	{name: "409 ABORTED", status: 409, body: `{"error_code":"ABORTED","message":"race"}`, code: "ABORTED", message: "race"},
	{name: "409 ALREADY_EXISTS", status: 409, body: `{"error_code":"ALREADY_EXISTS","message":"dup"}`, code: "ALREADY_EXISTS", message: "dup"},
	{name: "429 RESOURCE_EXHAUSTED", status: 429, body: `{"error_code":"RESOURCE_EXHAUSTED","message":"slow"}`, code: "RESOURCE_EXHAUSTED", message: "slow", retried: true},
	{name: "500 RESOURCE_DOES_NOT_EXIST", status: 500, body: `{"error_code":"RESOURCE_DOES_NOT_EXIST","message":"gone"}`, code: "RESOURCE_DOES_NOT_EXIST", message: "gone", missing: true},
	{name: "500 DATA_LOSS", status: 500, body: `{"error_code":"DATA_LOSS","message":"lost"}`, code: "DATA_LOSS", message: "lost"},
	{name: "500 UNKNOWN", status: 500, body: `{"error_code":"UNKNOWN","message":"what"}`, code: "UNKNOWN", message: "what"},
	// A message that the classic client retries on any status.
	{name: "400 transient message", status: 400, body: `{"error_code":"INVALID_STATE","message":"Workspace 1 does not have any associated worker environments"}`, code: "INVALID_STATE", message: "Workspace 1 does not have any associated worker environments", retried: true},
	// SCIM bodies.
	{name: "404 SCIM", status: 404, body: `{"schemas":["urn:ietf:params:scim:api:messages:2.0:Error"],"detail":"User 1 not found","status":"404"}`, code: "SCIM_404", message: "User 1 not found", missing: true},
	{name: "409 SCIM with scimType", status: 409, body: `{"detail":"User exists","status":"409","scimType":"uniqueness"}`, code: "SCIM_409", message: "uniqueness User exists"},
	{name: "500 SCIM with null detail", status: 500, body: `{"detail":"null","status":"500"}`, code: "SCIM_500", message: "SCIM API Internal Error"},
	// API 1.2 bodies, which have only error.
	{name: "404 API 1.2", status: 404, body: `{"error":"Command 1 not found"}`, code: "<nil>", message: "Command 1 not found", missing: true},
	{name: "500 API 1.2", status: 500, body: `{"error":"Command 1 failed"}`, code: "<nil>", message: "Command 1 failed"},
	// An integer error_code.
	{name: "404 integer error_code", status: 404, body: `{"error_code":404,"message":"gone"}`, code: "404", message: "gone", missing: true},
	{name: "500 integer error_code", status: 500, body: `{"error_code":500,"message":"oops"}`, code: "500", message: "oops"},
	// Error details.
	{name: "400 details", status: 400, body: classicErrorWithDetails, code: "INVALID_PARAMETER_VALUE", message: "Invalid name"},
	{name: "500 details", status: 500, body: classicErrorWithDetails, code: "INVALID_PARAMETER_VALUE", message: "Invalid name"},
	// Other JSON.
	{name: "404 message only", status: 404, body: `{"message":"gone"}`, code: "<nil>", message: "gone", missing: true},
	{name: "404 empty object", status: 404, body: `{}`, code: "<nil>", missing: true},
	{name: "404 null", status: 404, body: `null`, code: "<nil>", missing: true},
	{name: "404 array", status: 404, body: `[]`, code: "UNKNOWN", message: `[]`},
	{name: "404 truncated JSON", status: 404, body: classicErrorTruncatedJSON, code: "UNKNOWN", message: classicErrorTruncatedJSON},
	{name: "500 truncated JSON", status: 500, body: classicErrorTruncatedJSON, code: "UNKNOWN", message: classicErrorTruncatedJSON},
	// Plain text of the form CODE: message, which holds the error code.
	{name: "400 text", status: 400, body: `INVALID_PARAMETER_VALUE: bad name`, code: "INVALID_PARAMETER_VALUE", message: "bad name"},
	{name: "400 text RESOURCE_DOES_NOT_EXIST", status: 400, body: `RESOURCE_DOES_NOT_EXIST: gone`, code: "RESOURCE_DOES_NOT_EXIST", message: "gone", missing: true},
	{name: "404 text INVALID_PARAMETER_VALUE", status: 404, body: `INVALID_PARAMETER_VALUE: bad name`, code: "INVALID_PARAMETER_VALUE", message: "bad name"},
	{name: "500 text", status: 500, body: `INTERNAL_ERROR: oops`, code: "INTERNAL_ERROR", message: "oops"},
	{name: "404 other text", status: 404, body: "model not found\n", code: "UNKNOWN", message: "model not found\n"},
	// HTML with a <pre> block, which holds the message. The error code is the status text.
	{name: "400 HTML with pre", status: 400, body: classicErrorHTMLWithPre, code: "BAD_REQUEST", message: "Worker environment not ready"},
	{name: "404 HTML with pre", status: 404, body: classicErrorHTMLWithPre, code: "NOT_FOUND", message: "Worker environment not ready", missing: true},
	{name: "500 HTML with pre", status: 500, body: classicErrorHTMLWithPre, code: "INTERNAL_SERVER_ERROR", message: "Worker environment not ready"},
	{name: "503 HTML with pre", status: 503, body: classicErrorHTMLWithPre, code: "SERVICE_UNAVAILABLE", message: "Worker environment not ready", retried: true},
	// HTML without a <pre> block.
	{name: "404 HTML without pre", status: 404, body: classicErrorHTMLWithoutPre, code: "UNKNOWN", message: classicErrorHTMLWithoutPre},
	{name: "502 HTML without pre", status: 502, body: classicErrorHTMLWithoutPre, code: "UNKNOWN", message: classicErrorHTMLWithoutPre},
	{name: "503 HTML without pre", status: 503, body: classicErrorHTMLWithoutPre, code: "UNKNOWN", message: classicErrorHTMLWithoutPre, retried: true},
	// An empty body. The message is the status text.
	{name: "400 empty", status: 400, message: "Bad Request"},
	{name: "404 empty", status: 404, message: "Not Found", missing: true},
	{name: "429 empty", status: 429, message: "Too Many Requests", retried: true},
	{name: "500 empty", status: 500, message: "Internal Server Error"},
	{name: "503 empty", status: 503, message: "Service Unavailable", retried: true},
}

// classicErrorModelCalls are the calls of databricks_mlflow_model, each
// through the classic client and through the modular client.
var classicErrorModelCalls = []struct {
	name    string
	classic func(context.Context, *databricks.WorkspaceClient) error
	modular func(context.Context, *modelregistry.Client) error
}{
	{"read", func(ctx context.Context, w *databricks.WorkspaceClient) error {
		_, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "m"})
		return err
	}, func(ctx context.Context, c *modelregistry.Client) error {
		_, err := c.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: new("m")})
		return err
	}},
	{"create", func(ctx context.Context, w *databricks.WorkspaceClient) error {
		_, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: "m"})
		return err
	}, func(ctx context.Context, c *modelregistry.Client) error {
		_, err := c.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: new("m")})
		return err
	}},
	{"update", func(ctx context.Context, w *databricks.WorkspaceClient) error {
		_, err := w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{Name: "m"})
		return err
	}, func(ctx context.Context, c *modelregistry.Client) error {
		_, err := c.UpdateRegisteredModel(ctx, modelregistry.UpdateRegisteredModelRequest{Name: new("m")})
		return err
	}},
	{"delete", func(ctx context.Context, w *databricks.WorkspaceClient) error {
		return w.ModelRegistry.DeleteModel(ctx, ml.DeleteModelRequest{Name: "m"})
	}, func(ctx context.Context, c *modelregistry.Client) error {
		_, err := c.DeleteRegisteredModel(ctx, modelregistry.DeleteRegisteredModelRequest{Name: new("m")})
		return err
	}},
}

// classicErrorNoAmbientConfig keeps profiles and environment variables out of
// the configuration of the classic client.
type classicErrorNoAmbientConfig struct{}

func (classicErrorNoAmbientConfig) Name() string                   { return "classic error test" }
func (classicErrorNoAmbientConfig) Configure(*config.Config) error { return nil }

// classicErrorClients returns a classic and a modular client for server.
// Neither reads the environment or a config file. The classic client stops
// retrying after a second, the shortest time it can be set to, and the modular
// client doesn't retry.
func classicErrorClients(t *testing.T, server *httptest.Server) (*databricks.WorkspaceClient, *modelregistry.Client) {
	t.Helper()
	classic, err := databricks.NewWorkspaceClient(&databricks.Config{
		Host: server.URL, Token: "token", AuthType: "pat", Loaders: []config.Loader{classicErrorNoAmbientConfig{}},
		HTTPTransport: server.Client().Transport, RetryTimeoutSeconds: 1,
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) { return nil, nil },
	})
	require.NoError(t, err)
	pat, err := credentials.NewPATCredentials("token")
	require.NoError(t, err)
	modular, err := modelregistry.NewClient(t.Context(), client.WithHost(server.URL), client.WithCredentials(pat),
		client.WithHTTPClient(server.Client()), client.WithoutConfigFile(), client.WithoutEnv())
	require.NoError(t, err)
	return classic, modular
}

// classicErrorRespond answers every request with status and body. JSON goes
// out as JSON, truncated or not, and net/http names the content type of every
// other body.
func classicErrorRespond(status int, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if json.Valid([]byte(body)) || body == classicErrorTruncatedJSON {
			w.Header().Set("Content-Type", "application/json")
		}
		w.Header().Set("X-Request-Id", "r-1")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}
}

// classicErrorRespondRaw answers every request with statusLine as the HTTP/1.1
// status line, which a http.ResponseWriter can't write, and body as HTML.
func classicErrorRespondRaw(statusLine, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		conn, buffer, err := w.(http.Hijacker).Hijack()
		if err != nil {
			return
		}
		defer conn.Close()
		fmt.Fprintf(buffer, "HTTP/1.1 %s\r\nContent-Type: text/html\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s",
			statusLine, len(body), body)
		_ = buffer.Flush()
	}
}

// classicErrorServe starts a server for handler, over HTTP/2 with TLS when
// http2 is set and over HTTP/1.1 otherwise.
func classicErrorServe(t *testing.T, handler http.Handler, http2 bool) *httptest.Server {
	t.Helper()
	server := httptest.NewUnstartedServer(handler)
	if server.EnableHTTP2 = http2; http2 {
		server.StartTLS()
	} else {
		server.Start()
	}
	t.Cleanup(server.Close)
	return server
}

// classicErrorAPIError returns the classic APIError in the chain of err and
// fails the test when there is none.
func classicErrorAPIError(t *testing.T, err error) *apierr.APIError {
	t.Helper()
	apiErr, ok := errors.AsType[*apierr.APIError](err)
	require.True(t, ok, "no classic APIError in %#v", err)
	return apiErr
}

// classicErrorConvert makes the read call of databricks_mlflow_model fail
// through the modular client and returns its error and what ClassicError makes
// of it.
func classicErrorConvert(t *testing.T, modular *modelregistry.Client) (modularErr, converted error) {
	t.Helper()
	modularErr = classicErrorModelCalls[0].modular(t.Context(), modular)
	require.Error(t, modularErr)
	return modularErr, ClassicError(t.Context(), modularErr)
}

// classicErrorAssertSame fails the test unless converted, the result of
// ClassicError, reads like classicErr, the error of the classic client for the
// same response.
func classicErrorAssertSame(t *testing.T, classicErr, converted error) {
	t.Helper()
	want := classicErrorAPIError(t, classicErr)
	got, ok := converted.(*apierr.APIError)
	require.True(t, ok, "ClassicError returned %#v", converted)
	assert.Equal(t, classicErr.Error(), converted.Error(), "Error()")
	assert.Equal(t, apierr.IsMissing(classicErr), apierr.IsMissing(converted), "IsMissing")
	for _, sentinel := range classicErrorSentinels {
		assert.Equal(t, errors.Is(classicErr, sentinel), errors.Is(converted, sentinel), "errors.Is %q", sentinel)
	}
	assert.Equal(t, want.StatusCode, got.StatusCode, "StatusCode")
	assert.Equal(t, want.ErrorCode, got.ErrorCode, "ErrorCode")
	assert.Equal(t, want.Message, got.Message, "Message")
	assert.Equal(t, want.Details, got.Details, "Details")
	assert.Equal(t, want.ErrorDetails(), got.ErrorDetails(), "ErrorDetails()")
	assert.Equal(t, want.IsRetriable(t.Context()), got.IsRetriable(t.Context()), "IsRetriable")
	assert.Equal(t, want.ResponseWrapper.DebugBytes, got.ResponseWrapper.DebugBytes, "body")
	assert.Equal(t, want.ResponseWrapper.Response.Header.Get("X-Request-Id"),
		got.ResponseWrapper.Response.Header.Get("X-Request-Id"), "header")
}

func TestClassicError_SameAsClassicClient(t *testing.T) {
	for _, call := range classicErrorModelCalls {
		for _, response := range classicErrorResponses {
			if response.retried && call.name != "read" {
				continue // Each costs the second that the classic client retries for.
			}
			t.Run(call.name+"/"+response.name, func(t *testing.T) {
				t.Parallel()
				classic, modular := classicErrorClients(t, classicErrorServe(t, classicErrorRespond(response.status, response.body), false))
				classicErr := call.classic(t.Context(), classic)
				modularErr := call.modular(t.Context(), modular)
				_, ok := errors.AsType[*modularerr.APIError](modularErr)
				require.True(t, ok, "no modular APIError in %#v", modularErr)

				classicErrorAssertSame(t, classicErr, ClassicError(t.Context(), modularErr))

				want := classicErrorAPIError(t, classicErr)
				assert.Equal(t, response.status, want.StatusCode)
				assert.Equal(t, response.code, want.ErrorCode)
				assert.Equal(t, response.message, want.Message)
				assert.Equal(t, response.missing, apierr.IsMissing(classicErr))
				assert.Equal(t, response.retried, want.IsRetriable(t.Context()))
			})
		}
	}
}

// Over HTTP/2 a response has no status text, and net/http reports
// http.StatusText as ClassicError does. So the error code of an HTML response
// is the same even for a status that http.StatusText doesn't know.
func TestClassicError_SameAsClassicClientOverHTTP2(t *testing.T) {
	for _, status := range []int{404, 500, 599} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			t.Parallel()
			classic, modular := classicErrorClients(t, classicErrorServe(t, classicErrorRespond(status, classicErrorHTMLWithPre), true))
			classicErr := classicErrorModelCalls[0].classic(t.Context(), classic)
			_, converted := classicErrorConvert(t, modular)
			classicErrorAssertSame(t, classicErr, converted)
			assert.Equal(t, "HTTP/2.0", classicErrorAPIError(t, classicErr).ResponseWrapper.Response.Proto)
		})
	}
}

// Over HTTP/1.x the classic client makes the error code of an HTML response
// with a <pre> block from the status text that the server sends. ClassicError
// doesn't have it and takes http.StatusText instead.
func TestClassicError_DiffersByStatusText(t *testing.T) {
	type result struct {
		code    string
		missing bool
	}
	tests := []struct {
		name               string
		handler            http.HandlerFunc
		classic, converted result
	}{
		{"another status text", classicErrorRespondRaw("404 No Such Model", classicErrorHTMLWithPre), result{"NO_SUCH_MODEL", true}, result{"NOT_FOUND", true}},
		{"no status text", classicErrorRespondRaw("404", classicErrorHTMLWithPre), result{"UNKNOWN", false}, result{"NOT_FOUND", true}},
		// The server of net/http writes "status code 599" for a status that http.StatusText doesn't know.
		{"status without a text in net/http", classicErrorRespond(599, classicErrorHTMLWithPre), result{"STATUS_CODE_599", false}, result{"", false}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			classic, modular := classicErrorClients(t, classicErrorServe(t, test.handler, false))
			classicErr := classicErrorModelCalls[0].classic(t.Context(), classic)
			_, converted := classicErrorConvert(t, modular)

			assert.Equal(t, test.classic, result{classicErrorAPIError(t, classicErr).ErrorCode, apierr.IsMissing(classicErr)})
			assert.Equal(t, test.converted, result{classicErrorAPIError(t, converted).ErrorCode, apierr.IsMissing(converted)})
			assert.Equal(t, classicErr.Error(), converted.Error())
		})
	}
}

// The classic parser reports a 400 with INVALID_PARAMETER_VALUE from three GET
// paths as a missing resource. ClassicError doesn't have the request, so the
// same response stays an invalid parameter. The modular client gets the
// response through a model registry call, because go.mod has no modular client
// for these APIs.
func TestClassicError_DiffersByRequestOverride(t *testing.T) {
	tests := []struct {
		name, message string
		call          func(context.Context, *databricks.WorkspaceClient) error
	}{
		{"clusters get", "Cluster abc does not exist", func(ctx context.Context, w *databricks.WorkspaceClient) error {
			_, err := w.Clusters.GetByClusterId(ctx, "abc")
			return err
		}},
		{"jobs get", "Job 1 does not exist", func(ctx context.Context, w *databricks.WorkspaceClient) error {
			_, err := w.Jobs.GetByJobId(ctx, 1)
			return err
		}},
		{"jobs runs get", "Run 1 does not exist", func(ctx context.Context, w *databricks.WorkspaceClient) error {
			_, err := w.Jobs.GetRun(ctx, jobs.GetRunRequest{RunId: 1})
			return err
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			body := fmt.Sprintf(`{"error_code":"INVALID_PARAMETER_VALUE","message":%q}`, test.message)
			classic, modular := classicErrorClients(t, classicErrorServe(t, classicErrorRespond(400, body), false))
			classicErr := test.call(t.Context(), classic)
			_, converted := classicErrorConvert(t, modular)

			assert.ErrorIs(t, classicErr, apierr.ErrResourceDoesNotExist)
			assert.True(t, apierr.IsMissing(classicErr))
			assert.ErrorIs(t, converted, apierr.ErrInvalidParameterValue)
			assert.False(t, apierr.IsMissing(converted))
			assert.Equal(t, classicErr.Error(), converted.Error())
		})
	}
}

// A workspace refuses a request from outside its private link with a redirect
// to its login page. The classic client reports it as an error of its own. A
// modular client follows the redirect and fails to decode the page, which is
// no APIError, so ClassicError returns that error as it is.
func TestClassicError_DiffersForPrivateLinkRedirect(t *testing.T) {
	server := classicErrorServe(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login.html" {
			http.Redirect(w, r, "/login.html?error=private-link-validation-error", http.StatusFound)
			return
		}
		_, _ = io.WriteString(w, "<html><body>Log in</body></html>")
	}), false)
	classic, modular := classicErrorClients(t, server)

	classicErr := classicErrorAPIError(t, classicErrorModelCalls[0].classic(t.Context(), classic))
	assert.Equal(t, "PRIVATE_LINK_VALIDATION_ERROR", classicErr.ErrorCode)
	assert.ErrorIs(t, classicErr, apierr.ErrPermissionDenied)

	modularErr, converted := classicErrorConvert(t, modular)
	_, ok := errors.AsType[*modularerr.APIError](modularErr)
	assert.False(t, ok, "modular APIError in %#v", modularErr)
	assert.True(t, converted == modularErr, "ClassicError returned %#v", converted)
}

func TestClassicError_ReturnsOtherErrorsUnchanged(t *testing.T) {
	assert.NoError(t, ClassicError(t.Context(), nil))

	classicErr := &apierr.APIError{StatusCode: 404, ErrorCode: "RESOURCE_DOES_NOT_EXIST", Message: "gone"}
	var nilModularErr *modularerr.APIError
	tests := map[string]error{
		"plain error":                    errors.New("boom"),
		"context error":                  context.DeadlineExceeded,
		"network error":                  &url.Error{Op: "Get", URL: "https://example.com", Err: io.ErrUnexpectedEOF},
		"classic error":                  classicErr,
		"wrapped classic error":          fmt.Errorf("cannot read model: %w", classicErr),
		"nil modular error":              nilModularErr,
		"modular error without response": &modularerr.APIError{},
		"modular error below 400":        modularerr.FromHTTPError(101, http.Header{}, []byte(`{"error_code":"NOT_FOUND"}`)),
		"wrapped modular error below 400": fmt.Errorf("cannot read model: %w",
			modularerr.FromHTTPError(199, nil, nil)),
	}
	for name, err := range tests {
		t.Run(name, func(t *testing.T) {
			converted := ClassicError(t.Context(), err)
			assert.True(t, converted == err, "ClassicError returned %#v for %#v", converted, err)
		})
	}
}

func TestClassicError_FindsModularErrorInChain(t *testing.T) {
	modularErr := modularerr.FromHTTPError(404, nil, []byte(`{"error_code":"RESOURCE_DOES_NOT_EXIST","message":"gone"}`))
	tests := map[string]error{
		"itself":                modularErr,
		"wrapped":               fmt.Errorf("cannot read model: %w", modularErr),
		"wrapped twice":         fmt.Errorf("apply: %w", fmt.Errorf("cannot read model: %w", modularErr)),
		"joined":                errors.Join(errors.New("boom"), modularErr),
		"wrapped with an error": fmt.Errorf("%w: %w", context.Canceled, modularErr),
	}
	for name, err := range tests {
		t.Run(name, func(t *testing.T) {
			converted := ClassicError(t.Context(), err)
			classicErr, ok := converted.(*apierr.APIError)
			require.True(t, ok, "ClassicError returned %#v", converted)
			assert.Equal(t, apierr.APIError{StatusCode: 404, ErrorCode: "RESOURCE_DOES_NOT_EXIST", Message: "gone"},
				apierr.APIError{StatusCode: classicErr.StatusCode, ErrorCode: classicErr.ErrorCode, Message: classicErr.Message})
			assert.True(t, apierr.IsMissing(converted))
			// The classic error stands alone: nothing else of the chain is left.
			assert.Equal(t, "gone", converted.Error())
			assert.NotErrorIs(t, converted, context.Canceled)
			// Converting again changes nothing.
			assert.True(t, ClassicError(t.Context(), converted) == converted)
		})
	}
}

func TestClassicError_RebuildsResponse(t *testing.T) {
	header := http.Header{"Content-Type": {"application/json"}, "Retry-After": {"7"}}
	body := []byte(`{"error_code":"REQUEST_LIMIT_EXCEEDED","message":"slow"}`)
	converted := ClassicError(t.Context(), modularerr.FromHTTPError(429, header, body))

	wrapper := classicErrorAPIError(t, converted).ResponseWrapper
	assert.Equal(t, 429, wrapper.Response.StatusCode)
	assert.Equal(t, "429 Too Many Requests", wrapper.Response.Status)
	assert.Equal(t, header, wrapper.Response.Header)
	assert.Equal(t, body, wrapper.DebugBytes)
	responseBody, err := io.ReadAll(wrapper.Response.Body)
	require.NoError(t, err)
	assert.Equal(t, body, responseBody)
	// The stand-in request, and no request body.
	assert.Empty(t, wrapper.Response.Request.Method)
	assert.Equal(t, url.URL{}, *wrapper.Response.Request.URL)
	assert.Zero(t, wrapper.RequestBody)
}

func TestClassicError_ConcurrentUse(t *testing.T) {
	err := fmt.Errorf("cannot read model: %w", modularerr.FromHTTPError(404, http.Header{"X-Request-Id": {"r-1"}}, []byte(classicErrorWithDetails)))
	var group sync.WaitGroup
	for range 32 {
		group.Go(func() {
			converted := classicErrorAPIError(t, ClassicError(t.Context(), err))
			assert.Equal(t, "Invalid name", converted.Error())
			assert.Equal(t, "INVALID_NAME", converted.ErrorDetails().ErrorInfo.Reason)
		})
	}
	group.Wait()
}
