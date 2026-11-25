package sdkgoutility

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/useragent"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// transportWrapExchange is one request that a wrapped transport was asked to
// send, with the answer of the transport below it.
type transportWrapExchange struct {
	Method, URL, Host string
	Header            http.Header
	Status            int
	Body              string
}

// transportWrapRecorder is a transport that a caller puts inside Transport, as
// it would a debug log: it records every exchange and passes the answer of the
// transport below on, with its whole body.
type transportWrapRecorder struct {
	below     http.RoundTripper
	mu        sync.Mutex
	exchanges []transportWrapExchange
}

// wrap is the argument of Transport that puts r around the transport it is given.
func (r *transportWrapRecorder) wrap(base http.RoundTripper) (http.RoundTripper, error) {
	r.below = base
	return r, nil
}

func (r *transportWrapRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	exchange := transportWrapExchange{Method: req.Method, URL: req.URL.String(), Host: req.Host, Header: req.Header.Clone()}
	resp, err := r.below.RoundTrip(req)
	if resp != nil && resp.Body != nil {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewReader(body))
		exchange.Status, exchange.Body = resp.StatusCode, string(body)
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.exchanges = append(r.exchanges, exchange)
	return resp, err
}

func (r *transportWrapRecorder) seen() []transportWrapExchange {
	r.mu.Lock()
	defer r.mu.Unlock()
	return slices.Clone(r.exchanges)
}

// transportWrapModularClient returns a modular client for cfg that sends its
// requests over Transport(cfg, wrap).
func transportWrapModularClient(t *testing.T, cfg *config.Config, wrap func(http.RoundTripper) (http.RoundTripper, error)) *modelregistry.Client {
	t.Helper()
	transport, err := Transport(cfg, wrap)
	require.NoError(t, err)
	return transportTestModularClient(t, cfg, client.WithHTTPClient(&http.Client{Transport: transport}))
}

// TestTransport_WrapSeesTheRequestAsSent makes a call with a real modular
// client and compares the request that the wrapped transport was asked to send
// with the one that the server received.
func TestTransport_WrapSeesTheRequestAsSent(t *testing.T) {
	for _, create := range []bool{false, true} {
		server := transportTestNewServer(t, false)
		cfg := transportTestConfig(server.URL)
		cfg.Headers = config.StaticHeaders(map[string]string{"X-Transport-Test": "hook"})
		recorder := &transportWrapRecorder{}
		modular := transportWrapModularClient(t, cfg, recorder.wrap)
		assert.Same(t, transportClassicDefault, recorder.below, "wrap gets the transport that Transport sends over")
		ctx, name := useragent.InContext(t.Context(), "resource", "mlflow_model"), "model"
		if create {
			_, err := modular.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: &name})
			require.NoError(t, err)
		} else {
			require.NoError(t, transportTestGet(ctx, modular, name))
		}
		received, seen := server.received(), recorder.seen()
		require.Len(t, received, 1)
		require.Len(t, seen, 1)
		// http.Transport adds these two itself, the second from the length of
		// the request.
		sent := received[0].Header.Clone()
		sent.Del("Accept-Encoding")
		sent.Del("Content-Length")
		assert.Equal(t, sent, seen[0].Header)
		assert.Equal(t, received[0].Host, seen[0].Host)
		assert.Equal(t, received[0].Method, seen[0].Method)
		assert.Contains(t, seen[0].URL, server.URL+received[0].Path)
		assert.Regexp(t, ` resource/mlflow_model auth/pat\b.*`+transportTestModularMarker, seen[0].Header.Get("User-Agent"))
		assert.Equal(t, "application/json", seen[0].Header.Get("Accept"))
		assert.Regexp(t, transportTestTraceparent, seen[0].Header.Get("Traceparent"))
		assert.Equal(t, "hook", seen[0].Header.Get("X-Transport-Test"))
		assert.Equal(t, "Bearer transport-test-token", seen[0].Header.Get("Authorization"))
		assert.Equal(t, http.StatusOK, seen[0].Status)
		assert.Contains(t, seen[0].Body, "registered_model")
	}
}

// TestTransport_WrapBase checks which transport wrap is given for a config,
// and that requests go over the transport that wrap returns.
func TestTransport_WrapBase(t *testing.T) {
	custom := &transportTestBase{}
	testCases := []struct {
		name      string
		configure func(*config.Config)
		want      http.RoundTripper
	}{
		{"default", func(*config.Config) {}, transportClassicDefault},
		{"InsecureSkipVerify", func(cfg *config.Config) { cfg.InsecureSkipVerify = true }, transportClassicSkipVerify},
		{"HTTPTransport", func(cfg *config.Config) { cfg.HTTPTransport = custom }, custom},
		{"HTTPTransport and InsecureSkipVerify", func(cfg *config.Config) { cfg.HTTPTransport, cfg.InsecureSkipVerify = custom, true }, custom},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := transportTestConfig("https://classic.test")
			cfg.AuthType = "pat"
			tc.configure(cfg)
			replacement := &transportTestBase{}
			var given []http.RoundTripper
			transport, err := Transport(cfg, func(base http.RoundTripper) (http.RoundTripper, error) {
				given = append(given, base)
				return replacement, nil
			})
			require.NoError(t, err)
			for range 2 {
				_, err = transport.RoundTrip(httptest.NewRequest(http.MethodGet, "https://modular.test/api", nil))
				require.NoError(t, err)
			}
			require.Len(t, given, 1, "wrap is called once, not for every request")
			assert.Same(t, tc.want, given[0])
			require.NotNil(t, replacement.got, "requests go over the transport that wrap returned")
			assert.Nil(t, custom.got, "and not over the one that wrap was given")
			assert.Equal(t, "classic.test", replacement.got.Host)
			assert.Equal(t, "https://classic.test/api", replacement.got.URL.String())
		})
	}
}

// TestTransport_WrapErrors checks that an error of wrap and a nil transport
// from it are errors of Transport, and that a config that Transport can't use
// never reaches wrap.
func TestTransport_WrapErrors(t *testing.T) {
	cfg := transportTestConfig("https://classic.test")
	transport, err := Transport(cfg, func(http.RoundTripper) (http.RoundTripper, error) {
		return &transportTestBase{}, errTransportTest
	})
	assert.Nil(t, transport)
	assert.ErrorIs(t, err, errTransportTest)
	assert.EqualError(t, err, "cannot wrap the transport: transport test error")
	transport, err = Transport(cfg, func(http.RoundTripper) (http.RoundTripper, error) { return nil, nil })
	assert.Nil(t, transport)
	assert.EqualError(t, err, "cannot wrap the transport: wrap returned no transport")

	called := false
	wrap := func(base http.RoundTripper) (http.RoundTripper, error) {
		called = true
		return base, nil
	}
	_, err = Transport(nil, wrap)
	assert.EqualError(t, err, "cannot build a transport for a nil config")
	unresolvable := transportTestConfig("https://classic.test")
	unresolvable.Loaders = []config.Loader{transportTestLoader{errTransportTest}}
	_, err = Transport(unresolvable, wrap)
	assert.ErrorIs(t, err, errTransportTest)
	assert.False(t, called)
}

// TestTransport_WrapSeesNoRefusedRequest checks that a request that a visitor
// refuses reaches neither the wrapped transport nor the one below it.
func TestTransport_WrapSeesNoRefusedRequest(t *testing.T) {
	testCases := []struct {
		name      string
		configure func(*config.Config)
		want      string
	}{
		{"Headers hook fails", func(cfg *config.Config) { cfg.Headers = func(*http.Request) error { return errTransportTest } },
			transportTestVisitorError + "transport test error"},
		{"config that has not authenticated", func(cfg *config.Config) { cfg.AuthType = "" }, transportTestVisitorError + "panic: "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			base, recorder := &transportTestBase{}, &transportWrapRecorder{}
			cfg := transportTestConfig("https://classic.test")
			cfg.AuthType, cfg.HTTPTransport = "pat", base
			tc.configure(cfg)
			transport, err := Transport(cfg, recorder.wrap)
			require.NoError(t, err)
			resp, err := transport.RoundTrip(httptest.NewRequest(http.MethodGet, "https://classic.test/api", nil))
			assert.Nil(t, resp)
			assert.ErrorContains(t, err, tc.want)
			assert.Empty(t, recorder.seen())
			assert.Nil(t, base.got)
		})
	}
}

// TestTransport_WrapSeesThePrivateLinkLoginPage deletes a model in a workspace
// that refuses the request from outside its private link: the wrapped
// transport sees the redirect and the login page as the workspace sent them,
// and the caller gets the private link error.
func TestTransport_WrapSeesThePrivateLinkLoginPage(t *testing.T) {
	workspace := &transportPrivateLinkWorkspace{
		location: transportPrivateLinkLogin, redirectStatus: http.StatusFound, loginStatus: http.StatusOK,
	}
	server := httptest.NewServer(workspace)
	t.Cleanup(server.Close)
	recorder := &transportWrapRecorder{}
	modular := transportWrapModularClient(t, transportTestConfig(server.URL), recorder.wrap)
	_, err := modular.DeleteRegisteredModel(t.Context(), modelregistry.DeleteRegisteredModelRequest{Name: &transportPrivateLinkModel})
	got := transportPrivateLinkModularOutcome(t.Context(), err)
	assert.Equal(t, transportPrivateLinkOutcome{true, 403, transportPrivateLinkCode, got.Message, true}, got, "the caller gets the 403")

	seen := recorder.seen()
	require.Len(t, seen, 2)
	for _, exchange := range seen {
		assert.Regexp(t, ` auth/pat\b.*`+transportTestModularMarker, exchange.Header.Get("User-Agent"))
		assert.Regexp(t, transportTestTraceparent, exchange.Header.Get("Traceparent"))
	}
	assert.Equal(t, transportWrapExchange{
		Method: http.MethodDelete, URL: server.URL + "/api/2.0/mlflow/registered-models/delete?name=m",
		Host: server.Listener.Addr().String(), Status: http.StatusFound,
	}, transportWrapWithoutHeader(seen[0]))
	// http.Client gives the request of a redirect no Host of its own, so it is
	// sent to the host of its URL.
	assert.Equal(t, transportWrapExchange{
		Method: http.MethodGet, URL: server.URL + transportPrivateLinkLogin,
		Status: http.StatusOK, Body: transportPrivateLinkPage,
	}, transportWrapWithoutHeader(seen[1]))
}

func transportWrapWithoutHeader(exchange transportWrapExchange) transportWrapExchange {
	exchange.Header = nil
	return exchange
}

// TestTransport_WrapKeepsConnectionsShared makes two calls that each build
// their own config, Transport with a wrap and modular client: the second
// reaches the server over the connection of the first, because wrap is given
// the transport that all results of Transport share.
func TestTransport_WrapKeepsConnectionsShared(t *testing.T) {
	for _, useTLS := range []bool{false, true} {
		server := transportTestNewServer(t, useTLS)
		for range 2 {
			cfg := transportTestConfig(server.URL)
			cfg.InsecureSkipVerify = useTLS
			modular := transportWrapModularClient(t, cfg, (&transportWrapRecorder{}).wrap)
			require.NoError(t, transportTestGet(t.Context(), modular, "model"))
		}
		received := server.received()
		require.Len(t, received, 2)
		assert.Equal(t, received[0].Remote, received[1].Remote, "TLS: %v", useTLS)
	}
}

// transportTestBothClients makes the same call, a GET or a POST with a body,
// with a real classic client and then with a real modular client over
// Transport, both for cfg.
func transportTestBothClients(t *testing.T, cfg *config.Config, create bool, opts ...client.Option) {
	t.Helper()
	cfg.RateLimitPerSecond = 1000 // the classic client would wait between its calls
	modular := transportTestModularClient(t, cfg, opts...)
	classic, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	name := "model"
	if create {
		_, err = classic.ModelRegistry.CreateModel(t.Context(), ml.CreateModelRequest{Name: name})
		require.NoError(t, err)
		_, err = modular.CreateRegisteredModel(t.Context(), modelregistry.CreateRegisteredModelRequest{Name: &name})
		require.NoError(t, err)
		return
	}
	_, err = classic.ModelRegistry.GetModel(t.Context(), ml.GetModelRequest{Name: name})
	require.NoError(t, err)
	require.NoError(t, transportTestGet(t.Context(), modular, name))
}

// TestTransport_HeadersHookSetsAHeaderOfTheCall pins a difference from the
// classic client. The classic client sets the headers of the generated call
// after the cfg.Headers hook, so the call wins where both set a header.
// Transport runs the hook on a request that already has the headers of the
// modular call, so the hook wins. A header that the call doesn't set is the
// one of the hook with both clients.
func TestTransport_HeadersHookSetsAHeaderOfTheCall(t *testing.T) {
	const hookContentType, hookWorkspaceID = "text/hook", "hook"
	testCases := []struct {
		name        string
		create      bool
		workspaceID string
		// What the server receives from the classic client. From the modular
		// client it receives the values of the hook in every case.
		classicContentType, classicWorkspaceID string
	}{
		{"POST to a workspace with an ID", true, "1234567890", "application/json", "1234567890"},
		{"GET, which has no Content-Type of its own", false, "1234567890", hookContentType, "1234567890"},
		{"POST without a workspace ID, which the call then doesn't send", true, "", "application/json", hookWorkspaceID},
		{"GET without a workspace ID", false, "", hookContentType, hookWorkspaceID},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := transportTestNewServer(t, false)
			cfg := transportTestConfig(server.URL)
			cfg.WorkspaceID = tc.workspaceID
			cfg.Headers = config.StaticHeaders(map[string]string{
				"Content-Type": hookContentType, "X-Databricks-Workspace-Id": hookWorkspaceID,
			})
			var opts []client.Option
			if tc.workspaceID != "" {
				opts = append(opts, client.WithWorkspaceID(tc.workspaceID))
			}
			transportTestBothClients(t, cfg, tc.create, opts...)
			received := server.received()
			require.Len(t, received, 2)
			classic, modular := received[0], received[1]
			assert.Equal(t, tc.classicContentType, classic.Header.Get("Content-Type"), "classic client")
			assert.Equal(t, tc.classicWorkspaceID, classic.Header.Get("X-Databricks-Workspace-Id"), "classic client")
			assert.Equal(t, hookContentType, modular.Header.Get("Content-Type"), "modular client")
			assert.Equal(t, hookWorkspaceID, modular.Header.Get("X-Databricks-Workspace-Id"), "modular client")
			// Apart from these two headers the server received the same request.
			modular.Header.Set("Content-Type", classic.Header.Get("Content-Type"))
			modular.Header.Set("X-Databricks-Workspace-Id", classic.Header.Get("X-Databricks-Workspace-Id"))
			transportTestSameAsClassic(t, classic, modular)
		})
	}
}

// TestTransport_RedirectRunsTheHeadersHookAgain pins a difference from the
// classic client for a redirect that http.Client follows. The classic client
// runs its visitors before http.Client, which sends the redirected request
// with a copy of the headers of the first one: the cfg.Headers hook runs once
// and both requests have one traceparent. Transport sits inside http.Client,
// so it runs the hook again on the redirected request and gives it a
// traceparent of its own.
func TestTransport_RedirectRunsTheHeadersHookAgain(t *testing.T) {
	const hookHeader = "X-Transport-Test-Hook-Call"
	for _, create := range []bool{false, true} {
		server, target := transportTestNewServer(t, false), transportTestNewServer(t, false)
		server.redirectTo = target.URL
		var hookCalls atomic.Int32
		cfg := transportTestConfig(server.URL)
		cfg.Headers = func(r *http.Request) error {
			r.Header.Set(hookHeader, strconv.Itoa(int(hookCalls.Add(1))))
			return nil
		}
		transportTestBothClients(t, cfg, create)
		first, redirected := server.received(), target.received()
		require.Len(t, first, 2, "create: %v", create)
		require.Len(t, redirected, 2, "create: %v", create)
		assert.EqualValues(t, 3, hookCalls.Load(), "once for the classic call and twice for the modular one")

		// The classic client: the first call of the hook on both requests.
		assert.Equal(t, "1", first[0].Header.Get(hookHeader))
		assert.Equal(t, "1", redirected[0].Header.Get(hookHeader))
		assert.Equal(t, first[0].Header.Get("Traceparent"), redirected[0].Header.Get("Traceparent"))

		// The modular client: a call of the hook and a traceparent for each.
		assert.Equal(t, "2", first[1].Header.Get(hookHeader))
		assert.Equal(t, "3", redirected[1].Header.Get(hookHeader))
		assert.Regexp(t, transportTestTraceparent, first[1].Header.Get("Traceparent"))
		assert.NotEqual(t, first[1].Header.Get("Traceparent"), redirected[1].Header.Get("Traceparent"))

		// Apart from these, each server received the same request from both.
		first[1].Header.Set(hookHeader, "1")
		redirected[1].Header.Set(hookHeader, "1")
		transportTestSameAsClassic(t, first[0], first[1])
		transportTestSameAsClassic(t, redirected[0], redirected[1])
	}
}

// transportTestRedirectHosts is a base transport that answers for two hosts,
// which a test server can't have: it redirects every request for the host
// from to the same path of the host to, and records the headers of every
// request by its host.
type transportTestRedirectHosts struct {
	from, to string
	mu       sync.Mutex
	received map[string][]http.Header
}

func (h *transportTestRedirectHosts) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		req.Body.Close()
	}
	h.mu.Lock()
	h.received[req.URL.Host] = append(h.received[req.URL.Host], req.Header.Clone())
	h.mu.Unlock()
	recorder := httptest.NewRecorder()
	if req.URL.Host == h.from {
		http.Redirect(recorder, req, "https://"+h.to+req.URL.RequestURI(), http.StatusTemporaryRedirect)
	} else {
		io.WriteString(recorder, `{"registered_model_databricks":{"name":"m"}}`)
	}
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

// TestTransport_RedirectAndAuthorization checks the Authorization header that
// the host of a redirect receives from each client. http.Client leaves that
// header out of a request that it redirects to another domain. With the
// classic client that holds for the header of a cfg.Headers hook as well.
// Transport runs the hook again on the redirected request, so a hook that sets
// Authorization sets it there too: the one case where the hosts differ.
func TestTransport_RedirectAndAuthorization(t *testing.T) {
	const workspace, credentials, hook = "workspace.test", "Bearer transport-test-token", "Bearer hook"
	testCases := []struct {
		name, to                 string
		hookSetsIt               bool
		wantClassic, wantModular string
	}{
		{"credentials, to a subdomain", "sub." + workspace, false, credentials, credentials},
		{"credentials, to another domain", "other.test", false, "", ""},
		{"hook, to a subdomain", "sub." + workspace, true, hook, hook},
		{"hook, to another domain", "other.test", true, "", hook},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hosts := &transportTestRedirectHosts{from: workspace, to: tc.to, received: map[string][]http.Header{}}
			cfg := transportTestConfig("https://" + workspace)
			cfg.HTTPTransport = hosts
			sent := credentials
			if tc.hookSetsIt {
				cfg.Headers = config.StaticHeaders(map[string]string{"Authorization": hook, "X-Transport-Test": "hook"})
				sent = hook
			}
			transportTestBothClients(t, cfg, false)
			first, redirected := hosts.received[workspace], hosts.received[tc.to]
			require.Len(t, first, 2)
			require.Len(t, redirected, 2)
			assert.Equal(t, sent, first[0].Get("Authorization"), "classic client, first request")
			assert.Equal(t, sent, first[1].Get("Authorization"), "modular client, first request")
			assert.Equal(t, tc.wantClassic, redirected[0].Get("Authorization"), "classic client, redirected request")
			assert.Equal(t, tc.wantModular, redirected[1].Get("Authorization"), "modular client, redirected request")
			// A header of the hook that is no credential reaches both hosts.
			assert.Equal(t, redirected[0].Get("X-Transport-Test"), redirected[1].Get("X-Transport-Test"))
			assert.Equal(t, tc.hookSetsIt, redirected[1].Get("X-Transport-Test") == "hook")
		})
	}
}
