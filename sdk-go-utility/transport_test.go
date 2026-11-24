package sdkgoutility

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/sdk-go/auth"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	transportTestVisitorError  = "failed during request visitor: "
	errTransportTest           = errors.New("transport test error")
	transportTestTraceparent   = regexp.MustCompile(`^00-[0-9a-f]{32}-[0-9a-f]{16}-01$`)
	transportTestModularMarker = ` sdk-go-core/\S+ sdk-go-modelregistry/\S+$`
)

// transportTestLoader keeps a config from reading the environment and a config
// file, which the default loaders do. It fails with err.
type transportTestLoader struct{ err error }

func (l transportTestLoader) Name() string                   { return "transport-test" }
func (l transportTestLoader) Configure(*config.Config) error { return l.err }

// transportTestConfig returns a config for host that has not authenticated.
func transportTestConfig(host string) *config.Config {
	return &config.Config{
		Host:                 host,
		Token:                "transport-test-token",
		Loaders:              []config.Loader{transportTestLoader{}},
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) { return nil, nil },
	}
}

// transportTestCredentials authenticates a modular client with cfg.Authenticate,
// as credentials built on the classic config do. The modular SDK asks for them
// before it sends a request, so cfg has authenticated when Transport sees it.
type transportTestCredentials struct{ cfg *config.Config }

func (c transportTestCredentials) Name() string { return "transport-test" }
func (c transportTestCredentials) AuthHeaders(ctx context.Context) ([]auth.Header, error) {
	req := httptest.NewRequest(http.MethodGet, c.cfg.Host, nil)
	if err := c.cfg.Authenticate(req); err != nil {
		return nil, err
	}
	return []auth.Header{{Key: "Authorization", Value: req.Header.Get("Authorization")}}, nil
}

// transportTestRequest is what the test server received for one request.
// Remote is the address of the client, which names the connection.
type transportTestRequest struct {
	Method, Path, Query, Host, Body, Remote string
	Header                                  http.Header
}

// transportTestServer records every request and answers it with status, or
// redirects it to the same path of redirectTo.
type transportTestServer struct {
	*httptest.Server
	status     int
	redirectTo string
	mu         sync.Mutex
	requests   []transportTestRequest
}

func transportTestNewServer(t *testing.T, useTLS bool) *transportTestServer {
	s := &transportTestServer{status: http.StatusOK}
	s.Server = httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		s.mu.Lock()
		s.requests = append(s.requests, transportTestRequest{r.Method, r.URL.Path, r.URL.RawQuery, r.Host, string(body), r.RemoteAddr, r.Header.Clone()})
		s.mu.Unlock()
		if s.redirectTo != "" {
			http.Redirect(w, r, s.redirectTo+r.URL.RequestURI(), http.StatusTemporaryRedirect)
			return
		}
		w.WriteHeader(s.status)
		io.WriteString(w, `{"registered_model_databricks":{"name":"m"},"registered_model":{"name":"m"}}`)
	}))
	// The server would log the TLS handshake that a test expects to fail.
	s.Config.ErrorLog = log.New(io.Discard, "", 0)
	if useTLS {
		s.StartTLS()
	} else {
		s.Start()
	}
	t.Cleanup(s.Close)
	return s
}

func (s *transportTestServer) received() []transportTestRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	return slices.Clone(s.requests)
}

// transportTestModularClient returns a modular client for cfg.Host that sends
// its requests over Transport(cfg) and authenticates with cfg.
func transportTestModularClient(t *testing.T, cfg *config.Config, opts ...client.Option) *modelregistry.Client {
	t.Helper()
	transport, err := Transport(cfg)
	require.NoError(t, err)
	opts = append([]client.Option{
		client.WithHost(cfg.Host),
		client.WithCredentials(transportTestCredentials{cfg}),
		client.WithHTTPClient(&http.Client{Transport: transport}),
		client.WithoutConfigFile(),
		client.WithoutEnv(),
	}, opts...)
	modular, err := modelregistry.NewClient(t.Context(), opts...)
	require.NoError(t, err)
	return modular
}

func transportTestGet(ctx context.Context, modular *modelregistry.Client, name string) error {
	_, err := modular.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: &name})
	return err
}

// transportTestBase records the request it is asked to send and answers 200.
type transportTestBase struct {
	got  *http.Request
	body string
}

func (b *transportTestBase) RoundTrip(req *http.Request) (*http.Response, error) {
	b.got = req
	if req.Body != nil {
		body, _ := io.ReadAll(req.Body)
		b.body = string(body)
		req.Body.Close()
	}
	return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
}

// transportTestOverBase returns Transport for a config that has an explicit
// authentication type and sends over the base transport that it also returns.
func transportTestOverBase(t *testing.T, configure func(*config.Config)) (http.RoundTripper, *transportTestBase) {
	t.Helper()
	base := &transportTestBase{}
	cfg := transportTestConfig("https://classic.test")
	cfg.AuthType, cfg.HTTPTransport = "pat", base
	if configure != nil {
		configure(cfg)
	}
	transport, err := Transport(cfg)
	require.NoError(t, err)
	return transport, base
}

// transportTestBody is a request body that records whether it was closed.
type transportTestBody struct {
	io.Reader
	closed bool
}

func (b *transportTestBody) Close() error { b.closed = true; return nil }

// TestTransport_MatchesClassicClient sends one call with a real classic client
// and the same call with a real modular client over Transport, with one config
// and one context, and compares what the server received.
func TestTransport_MatchesClassicClient(t *testing.T) {
	useragent.WithProduct("databricks-tf-provider", "1.2.3")
	t.Cleanup(func() { useragent.WithProduct("unknown", "0.0.0") })
	provider := func(ctx context.Context) context.Context {
		return useragent.InContext(useragent.InContext(ctx, "sdk", "sdkv2"), "resource", "mlflow_model")
	}
	testCases := []struct {
		name         string
		env          map[string]string
		configure    func(*config.Config)
		context      func(context.Context) context.Context
		modular      []client.Option
		create       bool
		modularFirst bool
		redirect     bool
		want         map[string]string // a part of the value of each header
	}{{
		name:    "provider user agent data",
		context: provider,
		want: map[string]string{"User-Agent": " sdk/sdkv2 resource/mlflow_model auth/pat",
			"Accept": "application/json", "Authorization": "Bearer transport-test-token"},
	}, {
		name: "no user agent data in the context",
		want: map[string]string{"User-Agent": " auth/pat"},
	}, {
		name:         "modular call first on a config that has not authenticated",
		context:      provider,
		modularFirst: true,
		want:         map[string]string{"User-Agent": " resource/mlflow_model auth/pat"},
	}, {
		name:    "request with a body",
		context: provider,
		create:  true,
		want:    map[string]string{"Content-Type": "application/json", "Content-Length": "16"},
	}, {
		name:      "workspace ID",
		configure: func(cfg *config.Config) { cfg.WorkspaceID = "1234567890" },
		modular:   []client.Option{client.WithWorkspaceID("1234567890")},
		want:      map[string]string{"X-Databricks-Workspace-Id": "1234567890"},
	}, {
		name: "Headers hook",
		configure: func(cfg *config.Config) {
			cfg.Headers = config.StaticHeaders(map[string]string{"X-Transport-Test": "hook", "Accept": "text/csv", "Authorization": "Bearer hook"})
		},
		want: map[string]string{"X-Transport-Test": "hook", "Accept": "application/json", "Authorization": "Bearer hook"},
	}, {
		name:    "detected CI/CD provider, runtime and agent",
		env:     map[string]string{"GITHUB_ACTIONS": "true", "DATABRICKS_RUNTIME_VERSION": "15.4", "CLAUDECODE": "1"},
		context: provider,
		want:    map[string]string{"User-Agent": " auth/pat cicd/github runtime/15.4 agent/"},
	}, {
		name: "upstream product",
		env:  map[string]string{"DATABRICKS_SDK_UPSTREAM": "upstream-tool", "DATABRICKS_SDK_UPSTREAM_VERSION": "4.5.6"},
		want: map[string]string{"User-Agent": " upstream/upstream-tool upstream-version/4.5.6 "},
	}, {
		name:     "redirect to another host",
		context:  provider,
		redirect: true,
		want:     map[string]string{"User-Agent": " resource/mlflow_model auth/pat"},
	}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for name, value := range tc.env {
				t.Setenv(name, value)
			}
			// The classic SDK caches what it detects in the environment.
			useragent.ClearCache()
			t.Cleanup(useragent.ClearCache)
			server, target := transportTestNewServer(t, false), transportTestNewServer(t, false)
			if tc.redirect {
				server.redirectTo = target.URL
			}
			cfg := transportTestConfig(server.URL)
			ctx := t.Context()
			if tc.configure != nil {
				tc.configure(cfg)
			}
			if tc.context != nil {
				ctx = tc.context(ctx)
			}
			modular := transportTestModularClient(t, cfg, tc.modular...)
			classic, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
			require.NoError(t, err)
			name := "model"
			calls := []func() error{func() error {
				if tc.create {
					_, err := classic.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: name})
					return err
				}
				_, err := classic.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: name})
				return err
			}, func() error {
				if tc.create {
					_, err := modular.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: &name})
					return err
				}
				return transportTestGet(ctx, modular, name)
			}}
			if tc.modularFirst {
				slices.Reverse(calls)
			}
			for _, call := range calls {
				require.NoError(t, call())
			}
			received := server.received()
			require.Len(t, received, 2)
			if tc.modularFirst {
				slices.Reverse(received)
			}
			transportTestSameAsClassic(t, received[0], received[1])
			if hops := target.received(); tc.redirect {
				require.Len(t, hops, 2, "both clients follow the redirect to the other host")
				transportTestSameAsClassic(t, hops[0], hops[1])
			}
			assert.Regexp(t, `^databricks-tf-provider/1\.2\.3 databricks-sdk-go/`, received[0].Header.Get("User-Agent"))
			for key, want := range tc.want {
				assert.Contains(t, received[0].Header.Get(key), want, key)
			}
		})
	}
}

// transportTestSameAsClassic checks that a server received the same request from
// both clients, apart from the sdk-go- marker, the traceparent and the connection.
func transportTestSameAsClassic(t *testing.T, classic, modular transportTestRequest) {
	t.Helper()
	classicUserAgent := classic.Header.Get("User-Agent")
	assert.Regexp(t, "^"+regexp.QuoteMeta(classicUserAgent)+transportTestModularMarker, modular.Header.Get("User-Agent"))
	assert.Regexp(t, transportTestTraceparent, classic.Header.Get("Traceparent"))
	assert.Regexp(t, transportTestTraceparent, modular.Header.Get("Traceparent"))
	assert.NotEqual(t, classic.Header.Get("Traceparent"), modular.Header.Get("Traceparent"))
	modular.Header.Set("User-Agent", classicUserAgent)
	modular.Header.Set("Traceparent", classic.Header.Get("Traceparent"))
	modular.Remote = classic.Remote
	assert.Equal(t, classic, modular)
}

// TestTransport_Headers checks the headers that the base transport is asked to send.
func TestTransport_Headers(t *testing.T) {
	hook := func(key, value string) func(*http.Request) error {
		return config.StaticHeaders(map[string]string{key: value})
	}
	const traceparent = "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
	testCases := []struct {
		name   string
		header http.Header
		hook   func(*http.Request) error
		want   map[string]string
	}{
		{"adds Accept", http.Header{}, nil, map[string]string{"Accept": "application/json"}},
		{"keeps the Accept of the request", http.Header{"Accept": {"text/plain"}}, nil, map[string]string{"Accept": "text/plain"}},
		{"keeps the traceparent of the request", http.Header{"Traceparent": {traceparent}}, nil, map[string]string{"Traceparent": traceparent}},
		{"request without headers", nil, nil, map[string]string{"Accept": "application/json"}},
		{"does not authenticate", http.Header{}, nil, map[string]string{"Authorization": ""}},
		{"keeps the credentials and the call headers", http.Header{"Authorization": {"Bearer modular"}, "Content-Type": {"text/plain"}}, nil,
			map[string]string{"Authorization": "Bearer modular", "Content-Type": "text/plain"}},
		{"Headers hook adds a header", http.Header{}, hook("X-Transport-Test", "hook"), map[string]string{"X-Transport-Test": "hook"}},
		{"Headers hook overrides the credentials", http.Header{"Authorization": {"Bearer modular"}}, hook("Authorization", "Bearer hook"),
			map[string]string{"Authorization": "Bearer hook"}},
		// The classic call sets Accept and User-Agent after the visitors as well.
		{"Headers hook sets Accept", http.Header{}, hook("Accept", "text/csv"), map[string]string{"Accept": "application/json"}},
		{"Headers hook sets traceparent", http.Header{}, hook("traceparent", "hook"), map[string]string{"Traceparent": "hook"}},
		{"Headers hook sets User-Agent", http.Header{}, hook("User-Agent", "hook/1.0.0"), nil},
		{"Headers hook removes the URL", http.Header{}, func(r *http.Request) error { r.URL = nil; return nil }, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transport, base := transportTestOverBase(t, func(cfg *config.Config) { cfg.Headers = tc.hook })
			req := httptest.NewRequest(http.MethodGet, "https://classic.test/api", nil)
			req.Header = tc.header
			_, err := transport.RoundTrip(req)
			require.NoError(t, err)
			for key, want := range tc.want {
				assert.Equal(t, want, base.got.Header.Get(key), key)
			}
			assert.Regexp(t, `^\S+ databricks-sdk-go/\S+ go/\S+ os/\S+ (.* )?auth/pat\b`, base.got.Header.Get("User-Agent"))
			assert.NotEmpty(t, base.got.Header.Get("Traceparent"))
			assert.Equal(t, tc.header, req.Header, "the request of the caller")
		})
	}
}

// TestTransport_UserAgent checks which modular segments follow the classic ones.
func TestTransport_UserAgent(t *testing.T) {
	ctx := useragent.InContext(t.Context(), "resource", "mlflow_model")
	classic := useragent.FromContext(ctx)
	testCases := []struct{ name, modular, want string }{
		{"no User-Agent", "", classic},
		{"modular SDK", "sdk-go-core/0.1.0 go/1.26.0 os/linux cicd/github sdk-go-modelregistry/0.7.0 auth/pat",
			classic + " sdk-go-core/0.1.0 sdk-go-modelregistry/0.7.0"},
		{"another client", "curl/8.5.0", classic},
		{"segments that contain the prefix", "databricks-sdk-go-x/1 my-sdk-go-y/2 SDK-GO-z/3", classic},
		{"extra spaces", "  sdk-go-core/0.1.0   ", classic + " sdk-go-core/0.1.0"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, transportUserAgent(ctx, tc.modular))
		})
	}
}

// TestTransport_VisitorErrors checks that the error or the panic of a visitor
// comes back wrapped, that nothing is sent and that the body is closed.
func TestTransport_VisitorErrors(t *testing.T) {
	testCases := []struct {
		name      string
		configure func(*config.Config)
		request   func(*http.Request)
		want      string
		wantIs    error
	}{{
		name:      "Headers hook fails",
		configure: func(cfg *config.Config) { cfg.Headers = func(*http.Request) error { return errTransportTest } },
		want:      transportTestVisitorError + "transport test error",
		wantIs:    errTransportTest,
	}, {
		name:    "request without a URL",
		request: func(req *http.Request) { req.URL = nil },
		want:    transportTestVisitorError + "no URL found in request",
	}, {
		name:      "config that has not authenticated",
		configure: func(cfg *config.Config) { cfg.AuthType = "" },
		want:      transportTestVisitorError + "panic: invalid alphanumeric or semver string: ",
	}, {
		name:      "Headers hook panics",
		configure: func(cfg *config.Config) { cfg.Headers = func(*http.Request) error { panic("hook panic") } },
		want:      transportTestVisitorError + "panic: hook panic",
	}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transport, base := transportTestOverBase(t, tc.configure)
			body := &transportTestBody{Reader: strings.NewReader("body")}
			req := httptest.NewRequest(http.MethodPost, "https://classic.test/api", body)
			if tc.request != nil {
				tc.request(req)
			}
			resp, err := transport.RoundTrip(req)
			assert.Nil(t, resp)
			assert.EqualError(t, err, tc.want)
			assert.True(t, tc.wantIs == nil || errors.Is(err, tc.wantIs))
			assert.Nil(t, base.got, "nothing is sent")
			assert.True(t, body.closed, "the body is closed")
		})
	}
	// A real modular call, which has no body, returns the error to its caller.
	server := transportTestNewServer(t, false)
	cfg := transportTestConfig(server.URL)
	cfg.Headers = func(*http.Request) error { return errTransportTest }
	err := transportTestGet(t.Context(), transportTestModularClient(t, cfg), "model")
	assert.ErrorIs(t, err, errTransportTest)
	assert.ErrorContains(t, err, transportTestVisitorError+"transport test error")
	assert.Empty(t, server.received())
}

// TestTransport_RequestOfTheCaller sends a request with a body for another host
// than cfg.Host, with a Headers hook that changes the request it is given.
func TestTransport_RequestOfTheCaller(t *testing.T) {
	transport, base := transportTestOverBase(t, func(cfg *config.Config) {
		cfg.Headers = func(r *http.Request) error {
			r.Header.Set("X-Transport-Test", "hook")
			r.Header.Add("X-Call", "hook")
			r.URL.Path += "/hook"
			return nil
		}
	})
	body := &transportTestBody{Reader: strings.NewReader(`{"name":"model"}`)}
	req := httptest.NewRequest(http.MethodPost, "http://modular.test:8080/api?name=model", body)
	req.Header = http.Header{"X-Call": {"call"}, "User-Agent": {"sdk-go-core/0.1.0"}}
	before := req.Clone(req.Context())
	_, err := transport.RoundTrip(req)
	require.NoError(t, err)
	assert.Equal(t, before.Header, req.Header)
	assert.Equal(t, before.URL, req.URL)
	assert.Equal(t, "modular.test:8080", req.Host)
	assert.Equal(t, before.Context(), req.Context())
	// The visitors send the request to cfg.Host, and the Host header follows.
	assert.Equal(t, "https://classic.test/api/hook?name=model", base.got.URL.String())
	assert.Equal(t, "classic.test", base.got.Host)
	assert.Equal(t, []string{"call", "hook"}, base.got.Header.Values("X-Call"))
	assert.Regexp(t, `^\S+ databricks-sdk-go/.* auth/pat\b.* sdk-go-core/0\.1\.0$`, base.got.Header.Get("User-Agent"))
	assert.Equal(t, `{"name":"model"}`, base.body, "the base transport reads the body of the caller")
	assert.True(t, body.closed)
}

// TestTransport_Config checks the errors of Transport and the base it picks.
func TestTransport_Config(t *testing.T) {
	_, err := Transport(nil)
	assert.EqualError(t, err, "cannot build a transport for a nil config")
	unresolvable := transportTestConfig("https://classic.test")
	unresolvable.Loaders = []config.Loader{transportTestLoader{errTransportTest}}
	_, err = Transport(unresolvable)
	assert.ErrorIs(t, err, errTransportTest)
	assert.ErrorContains(t, err, "resolve: transport test error")
	base := func(configure func(*config.Config)) http.RoundTripper {
		cfg := transportTestConfig("https://classic.test")
		configure(cfg)
		transport, err := Transport(cfg)
		require.NoError(t, err)
		return transport.(*transportRoundTripper).base
	}
	custom := &transportTestBase{}
	assert.Same(t, transportClassicDefault, base(func(*config.Config) {}), "modular clients share one default transport")
	assert.Same(t, custom, base(func(cfg *config.Config) { cfg.HTTPTransport = custom }))
	assert.Same(t, custom, base(func(cfg *config.Config) { cfg.HTTPTransport, cfg.InsecureSkipVerify = custom, true }))
	assert.Same(t, transportClassicSkipVerify, base(func(cfg *config.Config) { cfg.InsecureSkipVerify = true }), "and one for InsecureSkipVerify")
}

// TestTransport_MirrorsClassicDefault compares every exported field of the
// mirrored transport with the unexported one of a classic client, which it
// reads by reflection. That can't see the dialer, which both wrap in a method
// value. Both sides are new, because sending a request fills in fields.
func TestTransport_MirrorsClassicDefault(t *testing.T) {
	clientCfg := httpclient.ClientConfig{InsecureSkipVerify: true}
	classic := reflect.ValueOf(httpclient.NewApiClient(clientCfg)).Elem().
		FieldByName("httpClient").Elem().FieldByName("Transport").Elem().Elem()
	mirror := reflect.ValueOf(transportNewClassicSkipVerify()).Elem()
	require.Equal(t, classic.Type(), mirror.Type())
	transportTestSameFields(t, "Transport", classic, mirror)
}

func transportTestSameFields(t *testing.T, path string, classic, mirror reflect.Value) {
	for i := range classic.NumField() {
		field := classic.Type().Field(i)
		name := path + "." + field.Name
		c, m := classic.Field(i), mirror.Field(i)
		switch {
		case !field.IsExported():
		case field.Type.Kind() == reflect.Func:
			assert.Equal(t, c.Pointer(), m.Pointer(), name)
		case field.Type.Kind() == reflect.Pointer && !c.IsNil() && !m.IsNil() && c.Elem().Kind() == reflect.Struct:
			transportTestSameFields(t, name, c.Elem(), m.Elem())
		default:
			// %#v prints a duration as a number on both sides. %v prints it as
			// text for the mirror only, whose fields are readable as interfaces.
			assert.Equal(t, fmt.Sprintf("%#v", c), fmt.Sprintf("%#v", m), name)
		}
	}
}

// TestTransport_InsecureSkipVerify calls a TLS server with a self-signed
// certificate. Each call has its own config, Transport and modular client.
func TestTransport_InsecureSkipVerify(t *testing.T) {
	server := transportTestNewServer(t, true)
	call := func(skipVerify bool) error {
		cfg := transportTestConfig(server.URL)
		cfg.InsecureSkipVerify = skipVerify
		return transportTestGet(t.Context(), transportTestModularClient(t, cfg), "model")
	}
	require.NoError(t, call(true))
	require.NoError(t, call(true))
	received := server.received()
	require.Len(t, received, 2)
	assert.Equal(t, received[0].Remote, received[1].Remote, "the second client reuses the connection of the first")
	assert.Regexp(t, ` auth/pat\b.*`+transportTestModularMarker, received[1].Header.Get("User-Agent"))
	// A config that verifies doesn't use the connection that skipped it.
	err := call(false)
	assert.ErrorContains(t, err, "failed to verify certificate")
	assert.NotContains(t, err.Error(), transportTestVisitorError)
	assert.Len(t, server.received(), 2)
}

// TestTransport_ReturnsResponsesAsTheyAre sends once, without credentials.
func TestTransport_ReturnsResponsesAsTheyAre(t *testing.T) {
	for _, status := range []int{200, 401, 404, 429, 503} {
		server := transportTestNewServer(t, false)
		server.status = status
		cfg := transportTestConfig(server.URL)
		cfg.AuthType = "pat"
		transport, err := Transport(cfg)
		require.NoError(t, err)
		resp, err := transport.RoundTrip(httptest.NewRequest(http.MethodGet, server.URL+"/api", nil))
		require.NoError(t, err)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NoError(t, resp.Body.Close())
		assert.Equal(t, status, resp.StatusCode)
		assert.Contains(t, string(body), "registered_model_databricks", status)
		require.Len(t, server.received(), 1, "no retry of %d", status)
		assert.Empty(t, server.received()[0].Header.Get("Authorization"), "no authentication")
	}
}

// TestTransport_ConcurrentRequests sends modular and classic calls at once
// with one config, each with its own User-Agent data in the context.
func TestTransport_ConcurrentRequests(t *testing.T) {
	const workers, callsPerWorker = 8, 10
	server := transportTestNewServer(t, false)
	cfg := transportTestConfig(server.URL)
	cfg.RateLimitPerSecond = 1000 // the classic client would wait between its calls
	modular := transportTestModularClient(t, cfg)
	classic, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	var wg sync.WaitGroup
	for worker := range workers {
		wg.Go(func() {
			name := fmt.Sprintf("model%d", worker)
			ctx := useragent.InContext(t.Context(), "resource", name)
			for range callsPerWorker {
				assert.NoError(t, transportTestGet(ctx, modular, name))
				_, err := classic.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: name})
				assert.NoError(t, err)
			}
		})
	}
	wg.Wait()
	received := server.received()
	require.Len(t, received, 2*workers*callsPerWorker)
	traceparents, modularCalls := map[string]bool{}, 0
	for _, r := range received {
		segments := strings.Fields(r.Header.Get("User-Agent"))
		assert.Contains(t, segments, "resource/"+strings.TrimPrefix(r.Query, "name="))
		assert.Contains(t, segments, "auth/pat")
		assert.Equal(t, "Bearer transport-test-token", r.Header.Get("Authorization"))
		traceparents[r.Header.Get("Traceparent")] = true
		modularCalls += strings.Count(r.Header.Get("User-Agent"), " sdk-go-modelregistry/")
	}
	assert.Len(t, traceparents, len(received), "every request has its own traceparent")
	assert.Equal(t, workers*callsPerWorker, modularCalls)
}
