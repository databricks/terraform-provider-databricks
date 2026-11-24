package sdkgoutility

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// credentialsVisitorFailed is what the classic client puts before the
	// error of a failed authentication.
	credentialsVisitorFailed = "failed during request visitor: "
	// credentialsRejectedClientID is the OAuth client that the token endpoint
	// rejects.
	credentialsRejectedClientID = "rejected"
	credentialsGetModelPath     = "/api/2.0/mlflow/databricks/registered-models/get?name=model"
)

// credentialsNoLoader replaces the loaders of a config, so that a test config
// holds only the attributes that the test sets, and nothing from the
// environment or the ~/.databrickscfg of the machine.
type credentialsNoLoader struct{}

func (credentialsNoLoader) Name() string                   { return "none" }
func (credentialsNoLoader) Configure(*config.Config) error { return nil }

// credentialsStrategy is a custom credentials strategy. Configure fails with
// configureErr, reports the strategy as not configured when it has no visit,
// and otherwise authenticates requests with visit.
type credentialsStrategy struct {
	name         string
	configureErr error
	visit        func(*http.Request) error
}

func (s credentialsStrategy) Name() string { return s.name }

func (s credentialsStrategy) Configure(context.Context, *config.Config) (credentials.CredentialsProvider, error) {
	if s.configureErr != nil {
		return nil, s.configureErr
	}
	if s.visit == nil {
		return nil, nil
	}
	return credentials.CredentialsProviderFn(s.visit), nil
}

// credentialsWorkspace is a test server that plays a workspace for both SDKs.
// It serves the OIDC discovery document, issues OAuth tokens and answers
// every API call, of which it keeps the headers.
type credentialsWorkspace struct {
	*httptest.Server
	// tokenTTLs holds the lifetimes in seconds of the tokens issued in turn.
	// The last one repeats.
	tokenTTLs []int

	mu     sync.Mutex
	tokens int
	calls  []http.Header
}

func credentialsNewWorkspace(t *testing.T, tokenTTLs ...int) *credentialsWorkspace {
	if len(tokenTTLs) == 0 {
		tokenTTLs = []int{3600}
	}
	ws := &credentialsWorkspace{tokenTTLs: tokenTTLs}
	ws.Server = httptest.NewServer(http.HandlerFunc(ws.serve))
	t.Cleanup(ws.Close)
	return ws
}

func (ws *credentialsWorkspace) serve(w http.ResponseWriter, r *http.Request) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	clientID, _, _ := r.BasicAuth()
	switch {
	case r.URL.Path == "/oidc/.well-known/oauth-authorization-server":
		fmt.Fprintf(w, `{"authorization_endpoint": "%[1]s/oidc/v1/authorize", "token_endpoint": "%[1]s/oidc/v1/token"}`, ws.URL)
	case r.URL.Path == "/oidc/v1/token" && clientID == credentialsRejectedClientID:
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `{"error": "invalid_client", "error_description": "Client authentication failed"}`)
	case r.URL.Path == "/oidc/v1/token":
		ttl := ws.tokenTTLs[min(ws.tokens, len(ws.tokenTTLs)-1)]
		ws.tokens++
		fmt.Fprintf(w, `{"access_token": "token-%d", "token_type": "Bearer", "expires_in": %d}`, ws.tokens, ttl)
	case strings.HasPrefix(r.URL.Path, "/api/"):
		ws.calls = append(ws.calls, r.Header.Clone())
		fmt.Fprint(w, `{"registered_model_databricks": {"name": "model"}}`)
	default:
		http.NotFound(w, r)
	}
}

// apiCalls returns the headers of the API calls that reached the workspace.
func (ws *credentialsWorkspace) apiCalls() []http.Header {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	return append([]http.Header(nil), ws.calls...)
}

func (ws *credentialsWorkspace) tokensIssued() int {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	return ws.tokens
}

// config returns a classic config for the workspace with the attributes that
// set gives it.
func (ws *credentialsWorkspace) config(set func(*config.Config)) *config.Config {
	cfg := &config.Config{
		Host:    ws.URL,
		Loaders: []config.Loader{credentialsNoLoader{}},
		// The workspace has no /.well-known/databricks-config to read.
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
			return nil, nil
		},
	}
	set(cfg)
	return cfg
}

// credentialsClients holds a classic client of a config and a modular client
// that authenticates with the Credentials of the same config.
type credentialsClients struct {
	classic func() (*databricks.WorkspaceClient, error)
	modular *modelregistry.Client
}

// clients returns both clients. The classic one is built on its first use, as
// the provider builds its workspace client, because building it already
// authenticates the config: until then, the credentials are the first to do
// so.
func (ws *credentialsWorkspace) clients(t *testing.T, cfg *config.Config) credentialsClients {
	modular, err := modelregistry.NewClient(context.Background(),
		client.WithHost(ws.URL),
		client.WithCredentials(Credentials(cfg)),
		client.WithoutEnv(),
		client.WithoutConfigFile(),
	)
	require.NoError(t, err)
	return credentialsClients{
		classic: sync.OnceValues(func() (*databricks.WorkspaceClient, error) {
			return databricks.NewWorkspaceClient((*databricks.Config)(cfg))
		}),
		modular: modular,
	}
}

func (c credentialsClients) classicGet(ctx context.Context) error {
	classic, err := c.classic()
	if err != nil {
		return err
	}
	_, err = classic.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "model"})
	return err
}

func (c credentialsClients) modularGet(ctx context.Context) error {
	name := "model"
	_, err := c.modular.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: &name})
	return err
}

// credentialsSentAuth returns the headers of an API call without those that
// either SDK adds to every request on its own. What is left is the
// authentication.
func credentialsSentAuth(call http.Header) http.Header {
	sent := call.Clone()
	for _, name := range []string{"Accept", "Accept-Encoding", "Traceparent", "User-Agent"} {
		sent.Del(name)
	}
	return sent
}

// credentialsOnWire returns headers as a server receives them.
func credentialsOnWire(headers []auth.Header) http.Header {
	sent := http.Header{}
	for _, h := range headers {
		sent.Add(h.Key, h.Value)
	}
	return sent
}

func credentialsBearer(token string) []auth.Header {
	return []auth.Header{{Key: "Authorization", Value: "Bearer " + token}}
}

func credentialsPAT(cfg *config.Config) { cfg.Token = "dapi-secret" }

func credentialsM2M(cfg *config.Config) { cfg.ClientID, cfg.ClientSecret = "client", "secret" }

func credentialsCustom(s credentialsStrategy) func(*config.Config) {
	return func(cfg *config.Config) { cfg.Credentials = s }
}

var credentialsTwoHeaders = credentialsStrategy{name: "two-headers", visit: func(r *http.Request) error {
	r.Header.Set("X-Custom-Auth", "second")
	r.Header.Set("Authorization", "Custom first")
	return nil
}}

func TestCredentials_SendsWhatClassicClientSends(t *testing.T) {
	cases := []struct {
		name     string
		set      func(*config.Config)
		wantName string
		want     []auth.Header
	}{
		{"personal access token", credentialsPAT, "pat", credentialsBearer("dapi-secret")},
		{"basic authentication", func(cfg *config.Config) { cfg.Username, cfg.Password = "user", "pass" }, "basic",
			[]auth.Header{{Key: "Authorization", Value: "Basic dXNlcjpwYXNz"}}},
		{"oauth m2m", credentialsM2M, "oauth-m2m", credentialsBearer("token-1")},
		{"custom strategy with two headers", credentialsCustom(credentialsTwoHeaders), "two-headers",
			[]auth.Header{{Key: "Authorization", Value: "Custom first"}, {Key: "X-Custom-Auth", Value: "second"}}},
		{"custom strategy with two values of one header", credentialsCustom(credentialsStrategy{name: "two-values", visit: func(r *http.Request) error {
			r.Header.Add("X-Custom-Auth", "b")
			r.Header.Add("X-Custom-Auth", "a")
			return nil
		}}), "two-values", []auth.Header{{Key: "X-Custom-Auth", Value: "b"}, {Key: "X-Custom-Auth", Value: "a"}}},
		{"custom strategy with no header", credentialsCustom(credentialsStrategy{name: "no-headers", visit: func(*http.Request) error { return nil }}),
			"no-headers", []auth.Header{}},
	}
	for _, tc := range cases {
		for _, first := range []string{"classic", "modular"} {
			t.Run(tc.name+", "+first+" client first", func(t *testing.T) {
				t.Parallel()
				ws := credentialsNewWorkspace(t)
				cfg := ws.config(tc.set)
				c := ws.clients(t, cfg)
				ctx := context.Background()
				gets := []func(context.Context) error{c.classicGet, c.modularGet}
				if first == "modular" {
					gets = []func(context.Context) error{c.modularGet, c.classicGet}
				}
				for _, get := range gets {
					require.NoError(t, get(ctx))
				}

				calls := ws.apiCalls()
				require.Len(t, calls, 2)
				for _, call := range calls {
					assert.Equal(t, credentialsOnWire(tc.want), credentialsSentAuth(call))
					assert.Contains(t, call.Get("User-Agent"), "auth/"+tc.wantName)
				}
				headers, err := Credentials(cfg).AuthHeaders(ctx)
				require.NoError(t, err)
				assert.Equal(t, tc.want, headers)
				assert.Equal(t, tc.wantName, Credentials(cfg).Name())
			})
		}
	}
}

func TestCredentials_OAuthTokenStaysWithClassicConfig(t *testing.T) {
	t.Parallel()
	const lifetime = 2 * time.Second
	ws := credentialsNewWorkspace(t, int(lifetime.Seconds()), 3600)
	cfg := ws.config(func(cfg *config.Config) {
		credentialsM2M(cfg)
		// The config refreshes a token in the background from half of its
		// lifetime on. Without that, the number of tokens issued is exact.
		cfg.DisableAsyncTokenRefresh = true
	})
	c := ws.clients(t, cfg)
	creds := Credentials(cfg)
	ctx := context.Background()

	callAll := func(token string) {
		for range 3 {
			headers, err := creds.AuthHeaders(ctx)
			require.NoError(t, err)
			assert.Equal(t, credentialsBearer(token), headers)
		}
		before := len(ws.apiCalls())
		require.NoError(t, c.modularGet(ctx))
		require.NoError(t, c.classicGet(ctx))
		for _, call := range ws.apiCalls()[before:] {
			assert.Equal(t, credentialsOnWire(credentialsBearer(token)), credentialsSentAuth(call))
		}
	}
	_, err := creds.AuthHeaders(ctx)
	require.NoError(t, err)
	fetched := time.Now()
	callAll("token-1")
	assert.Equal(t, 1, ws.tokensIssued(), "one token for all calls of both clients")

	time.Sleep(time.Until(fetched.Add(lifetime + 100*time.Millisecond)))
	callAll("token-2")
	assert.Equal(t, 2, ws.tokensIssued(), "one more token after the first one expired")
}

func TestCredentials_AuthenticationFails(t *testing.T) {
	boom := errors.New("boom")
	cases := []struct {
		name     string
		set      func(*config.Config)
		wantErr  string // HOST stands for the URL of the workspace
		wantIs   error
		wantName string
	}{
		{"strategy fails to configure", credentialsCustom(credentialsStrategy{name: "failing", configureErr: boom}),
			"failing auth: boom. Config: host=HOST", boom, credentialsUnknownName},
		{"strategy is not configured", credentialsCustom(credentialsStrategy{name: "absent"}),
			"absent auth: not configured. Config: host=HOST", nil, credentialsUnknownName},
		{"strategy fails to set headers", credentialsCustom(credentialsStrategy{name: "failing", visit: func(*http.Request) error { return boom }}),
			"boom", boom, "failing"},
		{"token endpoint rejects the client", func(cfg *config.Config) {
			cfg.AuthType, cfg.ClientID, cfg.ClientSecret = "oauth-m2m", credentialsRejectedClientID, "secret"
		}, `error getting token: Post "HOST/oidc/v1/token": {"error": "invalid_client", "error_description": "Client authentication failed"}`,
			nil, "oauth-m2m"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ws := credentialsNewWorkspace(t)
			cfg := ws.config(tc.set)
			c := ws.clients(t, cfg)
			ctx := context.Background()
			wantErr := credentialsVisitorFailed + strings.ReplaceAll(tc.wantErr, "HOST", ws.URL)

			headers, err := Credentials(cfg).AuthHeaders(ctx)
			assert.Nil(t, headers)
			for _, err := range []error{err, c.classicGet(ctx), c.modularGet(ctx)} {
				require.EqualError(t, err, wantErr)
				if tc.wantIs != nil {
					assert.ErrorIs(t, err, tc.wantIs)
				}
			}
			assert.Empty(t, ws.apiCalls())
			assert.Equal(t, tc.wantName, Credentials(cfg).Name())
		})
	}
}

func TestCredentials_ConfigFailsToResolve(t *testing.T) {
	t.Parallel()
	ws := credentialsNewWorkspace(t)
	twoMethods := func(cfg *config.Config) { cfg.Token, cfg.Username, cfg.Password = "dapi-secret", "user", "pass" }

	// The classic client fails when it is built, with the error that the
	// credentials wrap on every call.
	_, classicErr := databricks.NewWorkspaceClient((*databricks.Config)(ws.config(twoMethods)))
	require.ErrorContains(t, classicErr, "validate: more than one authorization method configured: basic and pat")

	creds := Credentials(ws.config(twoMethods))
	for range 2 {
		headers, err := creds.AuthHeaders(context.Background())
		assert.Nil(t, headers)
		require.EqualError(t, err, credentialsVisitorFailed+classicErr.Error())
	}
	assert.Equal(t, credentialsUnknownName, creds.Name())
}

func TestCredentials_CancelledContext(t *testing.T) {
	readsContext := credentialsStrategy{name: "reads-context", visit: func(r *http.Request) error {
		if err := r.Context().Err(); err != nil {
			return err
		}
		r.Header.Set("Authorization", "Custom first")
		return nil
	}}
	cases := []struct {
		name string
		set  func(*config.Config)
		// want is what a strategy that ignores the context still returns.
		want []auth.Header
	}{
		{"strategy that reads the context", credentialsCustom(readsContext), nil},
		{"strategy that ignores the context", credentialsPAT, credentialsBearer("dapi-secret")},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ws := credentialsNewWorkspace(t)
			cfg := ws.config(tc.set)
			c := ws.clients(t, cfg)
			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			headers, err := Credentials(cfg).AuthHeaders(ctx)
			assert.Equal(t, tc.want, headers)
			if tc.want == nil {
				require.EqualError(t, err, credentialsVisitorFailed+"context canceled")
				assert.ErrorIs(t, err, context.Canceled)
			} else {
				require.NoError(t, err)
			}
			// Neither client sends the call.
			assert.ErrorIs(t, c.classicGet(ctx), context.Canceled)
			assert.ErrorIs(t, c.modularGet(ctx), context.Canceled)
			assert.Empty(t, ws.apiCalls())

			// The cancelled call leaves nothing behind.
			headers, err = Credentials(cfg).AuthHeaders(context.Background())
			require.NoError(t, err)
			assert.Len(t, headers, 1)
		})
	}
}

func TestCredentials_NilContext(t *testing.T) {
	t.Parallel()
	var ctx context.Context
	headers, err := Credentials(credentialsNewWorkspace(t).config(credentialsPAT)).AuthHeaders(ctx)
	assert.Nil(t, headers)
	require.EqualError(t, err, "failed creating new request: net/http: nil Context")
}

func TestCredentials_ConcurrentCalls(t *testing.T) {
	cases := []struct {
		name       string
		set        func(*config.Config)
		wantName   string
		want       []auth.Header
		wantTokens int
	}{
		{"personal access token", credentialsPAT, "pat", credentialsBearer("dapi-secret"), 0},
		{"oauth m2m", credentialsM2M, "oauth-m2m", credentialsBearer("token-1"), 1},
		{"custom strategy with two headers", credentialsCustom(credentialsTwoHeaders), "two-headers",
			[]auth.Header{{Key: "Authorization", Value: "Custom first"}, {Key: "X-Custom-Auth", Value: "second"}}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			const workers = 32
			ws := credentialsNewWorkspace(t)
			cfg := ws.config(tc.set)
			// Resolved and not authenticated: the state in which the first
			// resources of a plan get the config of the provider.
			require.NoError(t, cfg.EnsureResolved())
			c := ws.clients(t, cfg)
			creds := Credentials(cfg)
			ctx := context.Background()

			// In the first round, all workers at once make the credentials
			// authenticate the config for the first time. The second round
			// adds calls of the classic client.
			rounds := [][]func(context.Context) error{{c.modularGet}, {c.modularGet, c.classicGet}}
			for round, gets := range rounds {
				headers := make([][]auth.Header, workers)
				names := make([]string, workers)
				errs := make([]error, workers)
				start := make(chan struct{})
				var wg sync.WaitGroup
				for i := range workers {
					wg.Go(func() {
						<-start
						var err error
						headers[i], err = creds.AuthHeaders(ctx)
						errs[i] = errors.Join(err, gets[i%len(gets)](ctx))
						names[i] = creds.Name()
					})
				}
				close(start)
				wg.Wait()

				for i := range workers {
					require.NoError(t, errs[i])
					assert.Equal(t, tc.want, headers[i])
					assert.Equal(t, tc.wantName, names[i])
				}
				require.Len(t, ws.apiCalls(), (round+1)*workers)
			}
			for _, call := range ws.apiCalls() {
				assert.Equal(t, credentialsOnWire(tc.want), credentialsSentAuth(call))
			}
			assert.Equal(t, tc.wantTokens, ws.tokensIssued())
		})
	}
}

func TestCredentials_Name(t *testing.T) {
	emptyName := credentialsStrategy{visit: func(r *http.Request) error {
		r.Header.Set("Authorization", "Custom first")
		return nil
	}}
	cases := []struct {
		name         string
		set          func(*config.Config)
		authenticate bool
		want         string
	}{
		{"before the config authenticates", credentialsPAT, false, credentialsUnknownName},
		{"type set on the config", func(cfg *config.Config) { cfg.AuthType = "azure-cli" }, false, "azure-cli"},
		{"type that authentication resolved", credentialsPAT, true, "pat"},
		{"strategy with an empty name", credentialsCustom(emptyName), true, credentialsUnknownName},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ws := credentialsNewWorkspace(t)
			cfg := ws.config(tc.set)
			creds := Credentials(cfg)
			if tc.authenticate {
				// The modular client puts the name in its User-Agent.
				require.NoError(t, ws.clients(t, cfg).modularGet(context.Background()))
				assert.Contains(t, ws.apiCalls()[0].Get("User-Agent"), "auth/"+tc.want)
			}
			assert.Equal(t, tc.want, creds.Name())
		})
	}
}

func TestCredentials_HeadersInStableOrder(t *testing.T) {
	t.Parallel()
	sets := map[string]string{"X-B": "2", "Authorization": "Custom first", "X-A": "1", "X-C": "3", "X-D": "4"}
	var calls atomic.Int64
	unordered := credentialsStrategy{name: "unordered", visit: func(r *http.Request) error {
		for name, value := range sets { // in a different order on every call
			r.Header.Set(name, value)
		}
		r.Header.Add("X-B", "1")
		r.Header["x-raw"] = []string{"raw"} // not canonical: http.Header.Set would make it X-Raw
		r.Header.Set("X-Call", fmt.Sprint(calls.Add(1)))
		return nil
	}}
	creds := Credentials(credentialsNewWorkspace(t).config(credentialsCustom(unordered)))
	for call := 1; call <= 50; call++ {
		headers, err := creds.AuthHeaders(context.Background())
		require.NoError(t, err)
		// Nothing is cached: the strategy runs on every call.
		assert.Equal(t, []auth.Header{
			{Key: "Authorization", Value: "Custom first"},
			{Key: "X-A", Value: "1"},
			{Key: "X-B", Value: "2"},
			{Key: "X-B", Value: "1"},
			{Key: "X-C", Value: "3"},
			{Key: "X-Call", Value: fmt.Sprint(call)},
			{Key: "X-D", Value: "4"},
			{Key: "x-raw", Value: "raw"},
		}, headers)
	}
}

// The classic client hands a request without host to the strategy, so the
// credentials do the same and never read cfg.Host. Both carry the context of
// the caller.
func TestCredentials_RequestGivenToStrategy(t *testing.T) {
	t.Parallel()
	type key struct{}
	type request struct {
		method, url, host string
		headers           int
		value             any
	}
	var mu sync.Mutex
	var seen []request
	recording := credentialsStrategy{name: "recording", visit: func(r *http.Request) error {
		mu.Lock()
		defer mu.Unlock()
		seen = append(seen, request{r.Method, r.URL.String(), r.Host, len(r.Header), r.Context().Value(key{})})
		return nil
	}}
	ws := credentialsNewWorkspace(t)
	cfg := ws.config(credentialsCustom(recording))
	c := ws.clients(t, cfg)
	ctx := context.WithValue(context.Background(), key{}, "from the caller")

	require.NoError(t, c.classicGet(ctx))
	require.NoError(t, c.modularGet(ctx))
	_, err := Credentials(cfg).AuthHeaders(ctx)
	require.NoError(t, err)

	synthetic := request{method: http.MethodGet, url: "/", value: "from the caller"}
	assert.Equal(t, []request{
		{method: http.MethodGet, url: credentialsGetModelPath, value: "from the caller"},
		synthetic,
		synthetic,
	}, seen)
}
