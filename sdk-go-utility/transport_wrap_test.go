package sdkgoutility

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
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
