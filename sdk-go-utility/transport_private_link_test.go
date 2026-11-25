package sdkgoutility

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/service/ml"
	modularerr "github.com/databricks/sdk-go/core/apierr"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	transportPrivateLinkLogin = "/login.html?error=private-link-validation-error"
	transportPrivateLinkCode  = "PRIVATE_LINK_VALIDATION_ERROR"
	transportPrivateLinkPage  = "<!doctype html><html><body>Log in to Databricks</body></html>"
)

// transportPrivateLinkModel names the model of every call.
var transportPrivateLinkModel = "m"

// transportPrivateLinkWorkspace answers as a workspace that refuses requests
// from outside its private link: it redirects every API request to location
// with redirectStatus, and answers every other request with loginStatus and a
// login page. It serves over HTTP and, as a base transport, under the host
// name of any cloud.
type transportPrivateLinkWorkspace struct {
	location                    string
	redirectStatus, loginStatus int
	mu                          sync.Mutex
	requests                    []string
}

func (w *transportPrivateLinkWorkspace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	w.mu.Lock()
	w.requests = append(w.requests, r.Method+" "+r.URL.String())
	w.mu.Unlock()
	if strings.HasPrefix(r.URL.Path, "/api/") {
		http.Redirect(rw, r, w.location, w.redirectStatus)
		return
	}
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(w.loginStatus)
	io.WriteString(rw, transportPrivateLinkPage)
}

func (w *transportPrivateLinkWorkspace) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		req.Body.Close()
	}
	recorder := httptest.NewRecorder()
	w.ServeHTTP(recorder, req)
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

// taken returns the requests that the workspace received since the last call.
func (w *transportPrivateLinkWorkspace) taken() []string {
	w.mu.Lock()
	defer w.mu.Unlock()
	requests := w.requests
	w.requests = nil
	return requests
}

// transportPrivateLinkCalls are the same call with a classic and with a
// modular client, for each method that the model registry API uses.
var transportPrivateLinkCalls = []struct {
	method  string
	classic func(context.Context, *databricks.WorkspaceClient) error
	modular func(context.Context, *modelregistry.Client) error
}{{
	http.MethodGet,
	func(ctx context.Context, w *databricks.WorkspaceClient) error {
		_, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: transportPrivateLinkModel})
		return err
	},
	func(ctx context.Context, c *modelregistry.Client) error {
		return transportTestGet(ctx, c, transportPrivateLinkModel)
	},
}, {
	http.MethodPost,
	func(ctx context.Context, w *databricks.WorkspaceClient) error {
		_, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: transportPrivateLinkModel})
		return err
	},
	func(ctx context.Context, c *modelregistry.Client) error {
		_, err := c.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: &transportPrivateLinkModel})
		return err
	},
}, {
	http.MethodPatch,
	func(ctx context.Context, w *databricks.WorkspaceClient) error {
		_, err := w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{Name: transportPrivateLinkModel})
		return err
	},
	func(ctx context.Context, c *modelregistry.Client) error {
		_, err := c.UpdateRegisteredModel(ctx, modelregistry.UpdateRegisteredModelRequest{Name: &transportPrivateLinkModel})
		return err
	},
}, {
	http.MethodDelete,
	func(ctx context.Context, w *databricks.WorkspaceClient) error {
		return w.ModelRegistry.DeleteModel(ctx, ml.DeleteModelRequest{Name: transportPrivateLinkModel})
	},
	func(ctx context.Context, c *modelregistry.Client) error {
		_, err := c.DeleteRegisteredModel(ctx, modelregistry.DeleteRegisteredModelRequest{Name: &transportPrivateLinkModel})
		return err
	},
}}

// transportPrivateLinkOutcome is what the caller of a call got back. Status,
// Code, Message and Denied are those of an API error.
type transportPrivateLinkOutcome struct {
	Failed        bool
	Status        int
	Code, Message string
	Denied        bool
}

func transportPrivateLinkClassicOutcome(err error) transportPrivateLinkOutcome {
	var apiErr *apierr.APIError
	if !errors.As(err, &apiErr) {
		return transportPrivateLinkOutcome{Failed: err != nil}
	}
	return transportPrivateLinkOutcome{true, apiErr.StatusCode, apiErr.ErrorCode, apiErr.Message, errors.Is(err, apierr.ErrPermissionDenied)}
}

// transportPrivateLinkModularOutcome reads the API error of a modular call as
// ClassicError does: it parses the response that the error carries with the
// parser of the classic client.
func transportPrivateLinkModularOutcome(ctx context.Context, err error) transportPrivateLinkOutcome {
	var apiErr *modularerr.APIError
	if !errors.As(err, &apiErr) {
		return transportPrivateLinkOutcome{Failed: err != nil}
	}
	status, body := apiErr.HTTPStatusCode(), apiErr.HTTPBody()
	return transportPrivateLinkClassicOutcome(apierr.GetAPIError(ctx, common.ResponseWrapper{
		ReadCloser: io.NopCloser(bytes.NewReader(body)),
		Response: &http.Response{
			Status:     fmt.Sprintf("%d %s", status, http.StatusText(status)),
			StatusCode: status,
			Header:     apiErr.HTTPHeader(),
			Request:    &http.Request{URL: &url.URL{}},
		},
	}))
}

// TestTransport_PrivateLinkMatchesClassicClient calls a workspace that refuses
// requests from outside its private link with a real classic client and with a
// real modular client over Transport, with one config, and compares what the
// callers got back and what the workspace received. wantMessage is a part of
// the message of the private link error; without it the classic client reports
// the answer of the login page as any other answer.
func TestTransport_PrivateLinkMatchesClassicClient(t *testing.T) {
	const aws, azure, gcp = "https://dbc-1.cloud.databricks.com", "https://adb-1.2.azuredatabricks.net", "https://1.2.gcp.databricks.com"
	testCases := []struct {
		name, host, location  string
		redirect, loginStatus int
		wantMessage           string
	}{
		{name: "AWS workspace on a server", wantMessage: "AWS PrivateLink"},
		{name: "AWS workspace", host: aws, wantMessage: "AWS PrivateLink"},
		{name: "Azure workspace", host: azure, wantMessage: "Azure Private Link"},
		{name: "GCP workspace", host: gcp, wantMessage: "Private Service Connect"},
		{name: "login page of another host", host: aws, location: azure + transportPrivateLinkLogin, wantMessage: "Azure Private Link"},
		{name: "host with a port", host: azure + ":443", wantMessage: "AWS PrivateLink"},
		{name: "redirect with 303", host: azure, redirect: 303, wantMessage: "Azure Private Link"},
		{name: "redirect with 307, which keeps the method", host: azure, redirect: 307, wantMessage: "Azure Private Link"},
		{name: "error among other parameters", host: gcp, location: "/login.html?next_url=%2F&error=private-link-validation-error&o=1", wantMessage: "Private Service Connect"},
		{name: "error inside another parameter", host: gcp, location: "/login.html?xerror=private-link-validation-errors", wantMessage: "Private Service Connect"},
		{name: "login page with 204", host: azure, loginStatus: 204, wantMessage: "Azure Private Link"},
		{name: "another error", host: azure, location: "/login.html?error=other"},
		{name: "escaped error", host: azure, location: "/login.html?error=private%2Dlink-validation-error"},
		{name: "another page", host: azure, location: "/aad/login.html?error=private-link-validation-error"},
		{name: "escaped path", host: azure, location: "/login%2Ehtml?error=private-link-validation-error"},
		{name: "login page with 403", host: azure, loginStatus: 403},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			workspace := &transportPrivateLinkWorkspace{
				location:       cmp.Or(tc.location, transportPrivateLinkLogin),
				redirectStatus: cmp.Or(tc.redirect, http.StatusFound),
				loginStatus:    cmp.Or(tc.loginStatus, http.StatusOK),
			}
			cfg := transportTestConfig(tc.host)
			cfg.RateLimitPerSecond = 1000 // the classic client would wait between its calls
			if tc.host == "" {
				server := httptest.NewServer(workspace)
				t.Cleanup(server.Close)
				cfg.Host = server.URL
			} else {
				cfg.HTTPTransport = workspace
			}
			modular := transportTestModularClient(t, cfg)
			classic, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
			require.NoError(t, err)
			for _, call := range transportPrivateLinkCalls {
				want := transportPrivateLinkClassicOutcome(call.classic(t.Context(), classic))
				wantRequests := workspace.taken()
				got := transportPrivateLinkModularOutcome(t.Context(), call.modular(t.Context(), modular))
				assert.Equal(t, want, got, call.method)
				assert.Equal(t, wantRequests, workspace.taken(), call.method)
				assert.Len(t, wantRequests, 2, call.method)
				if tc.wantMessage != "" {
					assert.Equal(t, transportPrivateLinkOutcome{true, 403, transportPrivateLinkCode, got.Message, true}, got, call.method)
					assert.Contains(t, got.Message, tc.wantMessage, call.method)
					continue
				}
				assert.NotEqual(t, transportPrivateLinkCode, got.Code, call.method)
				// A login page with a 200 is no error to either client, so only
				// a call that decodes the response fails, and a delete succeeds.
				assert.Equal(t, call.method != http.MethodDelete || tc.loginStatus >= 400, got.Failed, call.method)
			}
		})
	}
}

// transportPrivateLinkBase is a base transport that is a function.
type transportPrivateLinkBase func(*http.Request) (*http.Response, error)

func (f transportPrivateLinkBase) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

// transportPrivateLinkOverBase returns Transport for a workspace at host that
// sends over base.
func transportPrivateLinkOverBase(t *testing.T, host string, base transportPrivateLinkBase) http.RoundTripper {
	t.Helper()
	cfg := transportTestConfig(host)
	cfg.AuthType, cfg.HTTPTransport = "pat", base
	transport, err := Transport(cfg)
	require.NoError(t, err)
	return transport
}

// TestTransport_PrivateLinkResponse checks which responses of the base
// transport the transport replaces, and with what. request is the request of
// the response: "sent" for the one that the base transport was given, "" for
// none, as a base transport other than http.Transport may leave it, or the URL
// of another request.
func TestTransport_PrivateLinkResponse(t *testing.T) {
	const host = "https://adb-1.2.azuredatabricks.net"
	const api, login = host + "/api/2.0/mlflow/registered-models/delete", host + transportPrivateLinkLogin
	classicErr := apierr.GetAPIError(t.Context(), common.ResponseWrapper{
		Response: &http.Response{StatusCode: http.StatusOK, Request: httptest.NewRequest(http.MethodGet, login, nil)},
	}).(*apierr.APIError)
	wantBody, err := json.Marshal(map[string]string{"error_code": classicErr.ErrorCode, "message": classicErr.Message})
	require.NoError(t, err)
	require.Contains(t, string(wantBody), `{"error_code":"PRIVATE_LINK_VALIDATION_ERROR","message":"The requested workspace has Azure Private Link enabled`)
	testCases := []struct {
		name, url string
		status    int
		request   string
		noBody    bool
		replaced  bool
	}{
		{"login page", login, 200, "sent", false, true},
		{"login page with a status that is no error", login, 302, "sent", false, true},
		{"response without a request", login, 200, "", false, true},
		{"response with the request of the login page", api, 200, login, false, true},
		{"response with the request of another page", login, 200, api, false, false},
		{"response without a body", login, 200, "sent", true, true},
		{"login page with an error status", login, 400, "sent", false, false},
		{"another page", api, 200, "sent", false, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			page := &transportTestBody{Reader: strings.NewReader(transportPrivateLinkPage)}
			fromBase := &http.Response{
				StatusCode: tc.status, Proto: "HTTP/2.0", ProtoMajor: 2,
				Header: http.Header{"Content-Type": {"text/html"}, "Set-Cookie": {"a=b"}},
			}
			if !tc.noBody {
				fromBase.Body = page
			}
			transport := transportPrivateLinkOverBase(t, host, func(req *http.Request) (*http.Response, error) {
				switch tc.request {
				case "sent":
					fromBase.Request = req
				case "":
				default:
					fromBase.Request = httptest.NewRequest(http.MethodGet, tc.request, nil)
				}
				return fromBase, nil
			})
			got, err := transport.RoundTrip(httptest.NewRequest(http.MethodDelete, tc.url, nil))
			require.NoError(t, err)
			if !tc.replaced {
				assert.Same(t, fromBase, got)
				assert.False(t, page.closed)
				return
			}
			body, err := io.ReadAll(got.Body)
			require.NoError(t, err)
			require.NoError(t, got.Body.Close())
			assert.Equal(t, string(wantBody), string(body))
			got.Body = nil
			assert.Equal(t, &http.Response{
				Status: "403 Forbidden", StatusCode: 403, Proto: "HTTP/2.0", ProtoMajor: 2,
				Header:        http.Header{"Content-Type": {"application/json"}},
				ContentLength: int64(len(wantBody)),
				Request:       fromBase.Request,
			}, got)
			assert.Equal(t, !tc.noBody, page.closed, "the login page is closed")
		})
	}
	// Without a response there is nothing to replace.
	for _, baseErr := range []error{nil, errTransportTest} {
		transport := transportPrivateLinkOverBase(t, host, func(*http.Request) (*http.Response, error) { return nil, baseErr })
		got, err := transport.RoundTrip(httptest.NewRequest(http.MethodGet, login, nil))
		assert.Nil(t, got)
		assert.Equal(t, baseErr, err)
	}
}
