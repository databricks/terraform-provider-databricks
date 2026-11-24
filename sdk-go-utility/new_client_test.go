package sdkgoutility

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	classicapierr "github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/sdk-go/auth"
	modularapierr "github.com/databricks/sdk-go/core/apierr"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/client"
	"github.com/databricks/sdk-go/options/internaloptions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// newClientModel is a model name that has to be escaped in a query.
	newClientModel = "my model/ü&x=1"
	// newClientAnswer is what the workspace answers an API call with.
	newClientAnswer = `{"registered_model":{"name":"m","description":"d","creation_timestamp":1700000000000,"tags":[{"key":"k","value":"v"}]},` +
		`"registered_model_databricks":{"name":"m","id":"id-1","description":"d","permission_level":"CAN_MANAGE","tags":[{"key":"k","value":"v"}]}}`
)

var newClientTraceparent = regexp.MustCompile(`^00-[0-9a-f]{32}-[0-9a-f]{16}-01$`)

func newClientPtr[T any](value T) *T { return &value }

// newClientLoader replaces the loaders of a config, so that it reads nothing
// from the environment or a config file. It fails with err.
type newClientLoader struct{ err error }

func (l newClientLoader) Name() string                   { return "new-client-test" }
func (l newClientLoader) Configure(*config.Config) error { return l.err }

// newClientCounter is a transport that counts the requests sent over it.
type newClientCounter struct{ sent atomic.Int32 }

func (c *newClientCounter) RoundTrip(r *http.Request) (*http.Response, error) {
	c.sent.Add(1)
	return http.DefaultTransport.RoundTrip(r)
}

// newClientRequest is what the workspace received for one API call.
type newClientRequest struct {
	Method, Path, Query, Host, Body string
	Header                          http.Header
	conn                            string // the address it came from
}

// newClientWorkspace is a test server that plays a workspace for both SDKs:
// it serves OIDC discovery, issues OAuth tokens and records the API calls.
type newClientWorkspace struct {
	*httptest.Server
	useTLS bool
	// answer replaces newClientAnswer. before counts the earlier API calls.
	answer func(w http.ResponseWriter, r *http.Request, before int)

	mu                  sync.Mutex
	served, api, tokens int
	calls               []newClientRequest
}

func newClientStart(t *testing.T, ws *newClientWorkspace) *newClientWorkspace {
	ws.Server = httptest.NewUnstartedServer(http.HandlerFunc(ws.serve))
	if ws.useTLS {
		ws.StartTLS()
	} else {
		ws.Start()
	}
	t.Cleanup(ws.Close)
	return ws
}

func (ws *newClientWorkspace) serve(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	isAPI, isToken := strings.HasPrefix(r.URL.Path, "/api/"), r.URL.Path == "/oidc/v1/token"
	ws.mu.Lock()
	ws.served++
	before := ws.api
	if isAPI {
		ws.api++
		ws.calls = append(ws.calls, newClientRequest{r.Method, r.URL.Path, r.URL.RawQuery, r.Host, string(body), r.Header.Clone(), r.RemoteAddr})
	}
	if isToken {
		ws.tokens++
	}
	token := ws.tokens
	ws.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.URL.Path == "/oidc/.well-known/oauth-authorization-server":
		fmt.Fprintf(w, `{"authorization_endpoint":"%[1]s/oidc/v1/authorize","token_endpoint":"%[1]s/oidc/v1/token"}`, ws.URL)
	case isToken:
		fmt.Fprintf(w, `{"access_token":"token-%d","token_type":"Bearer","expires_in":3600}`, token)
	case isAPI && ws.answer != nil:
		ws.answer(w, r, before)
	case isAPI:
		io.WriteString(w, newClientAnswer)
	default:
		http.NotFound(w, r)
	}
}

// received returns the API calls that came in since it was last called.
func (ws *newClientWorkspace) received() []newClientRequest {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	calls := ws.calls
	ws.calls = nil
	return calls
}

// counts returns the number of requests of any kind and of tokens issued.
func (ws *newClientWorkspace) counts() (served, tokens int) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	return ws.served, ws.tokens
}

// config returns a classic config for the workspace, with a personal access
// token unless set changes that. It is neither resolved nor authenticated.
func (ws *newClientWorkspace) config(set func(*config.Config)) *config.Config {
	cfg := &config.Config{
		Host:    ws.URL,
		Token:   "dapi-secret",
		Loaders: []config.Loader{newClientLoader{}},
		// The workspace has no /.well-known/databricks-config to read.
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) { return nil, nil },
	}
	if set != nil {
		set(cfg)
	}
	return cfg
}

func newClientM2M(cfg *config.Config) {
	cfg.Token, cfg.ClientID, cfg.ClientSecret, cfg.WorkspaceID = "", "client", "secret", "42"
}

// newClientClassic builds the classic client of cfg, as the provider does
// before a resource runs. That resolves and authenticates cfg.
func newClientClassic(t *testing.T, cfg *config.Config) *databricks.WorkspaceClient {
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	require.NotEmpty(t, cfg.AuthType, "building the classic client didn't authenticate its config")
	return w
}

// newClientCalls are the calls that the provider makes for a registered
// model, with the classic and with the modular client. Each returns its response.
var newClientCalls = []struct {
	name    string
	classic func(context.Context, *databricks.WorkspaceClient) (any, error)
	modular func(context.Context, *modelregistry.Client) (any, error)
}{{
	name: "create",
	classic: func(ctx context.Context, w *databricks.WorkspaceClient) (any, error) {
		return w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: newClientModel, Description: "d", Tags: []ml.ModelTag{{Key: "k", Value: "v"}}})
	},
	modular: func(ctx context.Context, c *modelregistry.Client) (any, error) {
		return c.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: newClientPtr(newClientModel), Description: newClientPtr("d"),
			Tags: []modelregistry.RegisteredModelTag{{Key: newClientPtr("k"), Value: newClientPtr("v")}}})
	},
}, {
	name: "get",
	classic: func(ctx context.Context, w *databricks.WorkspaceClient) (any, error) {
		return w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: newClientModel})
	},
	modular: func(ctx context.Context, c *modelregistry.Client) (any, error) {
		return c.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: newClientPtr(newClientModel)})
	},
}, {
	name: "update",
	classic: func(ctx context.Context, w *databricks.WorkspaceClient) (any, error) {
		return w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{Name: newClientModel, Description: "d"})
	},
	modular: func(ctx context.Context, c *modelregistry.Client) (any, error) {
		return c.UpdateRegisteredModel(ctx, modelregistry.UpdateRegisteredModelRequest{Name: newClientPtr(newClientModel), Description: newClientPtr("d")})
	},
}, {
	name: "delete",
	classic: func(ctx context.Context, w *databricks.WorkspaceClient) (any, error) {
		return nil, w.ModelRegistry.DeleteModel(ctx, ml.DeleteModelRequest{Name: newClientModel})
	},
	modular: func(ctx context.Context, c *modelregistry.Client) (any, error) {
		_, err := c.DeleteRegisteredModel(ctx, modelregistry.DeleteRegisteredModelRequest{Name: newClientPtr(newClientModel)})
		return nil, err
	},
}}

// newClientSameRequest checks that the workspace received the same request
// from both clients, apart from the sdk-go- segments at the end of the
// User-Agent, the value of traceparent and the order of the keys in the body.
func newClientSameRequest(t *testing.T, classic, modular newClientRequest) {
	t.Helper()
	userAgent, traceparent := classic.Header.Get("User-Agent"), classic.Header.Get("Traceparent")
	assert.Regexp(t, "^"+regexp.QuoteMeta(userAgent)+` sdk-go-core/\S+ sdk-go-modelregistry/\S+$`, modular.Header.Get("User-Agent"))
	assert.Regexp(t, newClientTraceparent, traceparent)
	assert.Regexp(t, newClientTraceparent, modular.Header.Get("Traceparent"))
	assert.NotEqual(t, traceparent, modular.Header.Get("Traceparent"))
	if classic.Body != "" {
		assert.JSONEq(t, classic.Body, modular.Body)
		modular.Body = classic.Body
	}
	modular.Header = modular.Header.Clone()
	modular.Header.Set("User-Agent", userAgent)
	modular.Header.Set("Traceparent", traceparent)
	modular.conn = classic.conn
	assert.Equal(t, classic, modular)
}

// newClientDivert makes the environment and ~/.databrickscfg point at
// elsewhere, and checks that a modular client that is left alone goes there.
func newClientDivert(t *testing.T, elsewhere *newClientWorkspace) {
	require.NoError(t, os.WriteFile(filepath.Join(os.Getenv("HOME"), ".databrickscfg"),
		[]byte("[DEFAULT]\nhost = "+elsewhere.URL+"\ntoken = file-token\nworkspace_id = 888\n"), 0o600))
	for name, value := range map[string]string{"DATABRICKS_HOST": elsewhere.URL, "DATABRICKS_TOKEN": "env-token", "DATABRICKS_WORKSPACE_ID": "999"} {
		t.Setenv(name, value)
	}
	leftAlone, err := modelregistry.NewClient(t.Context())
	require.NoError(t, err)
	_, err = newClientCalls[1].modular(t.Context(), leftAlone)
	require.NoError(t, err)
	diverted := elsewhere.received()
	require.Len(t, diverted, 1)
	require.Equal(t, "Bearer env-token", diverted[0].Header.Get("Authorization"))
}

func TestNewClient_SendsWhatClassicClientSends(t *testing.T) {
	tests := []struct {
		name          string
		configure     func(*config.Config)
		modularFirst  bool // the config has not authenticated when NewClient gets it
		elsewhere     bool // the environment and ~/.databrickscfg point at another workspace
		httpTransport bool // the config has an HTTPTransport
		authorization string
		workspaceID   []string
		authType      string
		tokens        int
	}{
		{name: "personal access token and no workspace ID", authorization: "Bearer dapi-secret", authType: "pat"},
		{name: "workspace ID", configure: func(cfg *config.Config) { cfg.WorkspaceID = "1234567890" },
			authorization: "Bearer dapi-secret", workspaceID: []string{"1234567890"}, authType: "pat"},
		{name: "OAuth M2M with the token endpoint on the workspace", configure: newClientM2M,
			authorization: "Bearer token-1", workspaceID: []string{"42"}, authType: "oauth-m2m", tokens: 1},
		{name: "first request of an unresolved config is a modular one", configure: func(cfg *config.Config) { cfg.Host += "/?o=77" },
			modularFirst: true, authorization: "Bearer dapi-secret", workspaceID: []string{"77"}, authType: "pat"},
		{name: "first request of a config is a modular one that needs an OAuth token", configure: newClientM2M,
			modularFirst: true, authorization: "Bearer token-1", workspaceID: []string{"42"}, authType: "oauth-m2m", tokens: 1},
		{name: "DATABRICKS_ variables and ~/.databrickscfg point elsewhere", configure: func(cfg *config.Config) { cfg.Loaders = nil },
			elsewhere: true, authorization: "Bearer dapi-secret", authType: "pat"},
		{name: "cfg.HTTPTransport", httpTransport: true, authorization: "Bearer dapi-secret", authType: "pat"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Hide the configuration of the machine from the default loaders.
			for _, attribute := range config.ConfigAttributes {
				for _, name := range attribute.EnvVars {
					t.Setenv(name, "")
				}
			}
			t.Setenv("HOME", t.TempDir())
			ws, elsewhere := newClientStart(t, &newClientWorkspace{}), newClientStart(t, &newClientWorkspace{})
			cfg, through := ws.config(tc.configure), &newClientCounter{}
			if tc.httpTransport {
				cfg.HTTPTransport = through
			}
			classic := sync.OnceValue(func() *databricks.WorkspaceClient { return newClientClassic(t, cfg) })
			if !tc.modularFirst {
				classic()
			}
			if tc.elsewhere {
				newClientDivert(t, elsewhere)
			}
			ctx := useragent.InContext(t.Context(), "resource", "mlflow_model")
			modular, err := NewClient(ctx, cfg, modelregistry.NewClient)
			require.NoError(t, err)
			require.Equal(t, !tc.modularFirst, cfg.AuthType != "", "whether the config has authenticated before the first call")
			for _, call := range newClientCalls {
				fromModular, err := call.modular(ctx, modular)
				require.NoError(t, err, call.name)
				fromClassic, err := call.classic(ctx, classic())
				require.NoError(t, err, call.name)
				received := ws.received()
				require.Len(t, received, 2, call.name)
				newClientSameRequest(t, received[1], received[0])
				wantJSON, _ := json.Marshal(fromClassic)
				gotJSON, _ := json.Marshal(fromModular)
				assert.JSONEq(t, string(wantJSON), string(gotJSON), "what %s returns", call.name)
				assert.Equal(t, []string{tc.authorization}, received[0].Header.Values("Authorization"), call.name)
				assert.Equal(t, tc.workspaceID, received[0].Header.Values("X-Databricks-Workspace-Id"), call.name)
				assert.Contains(t, received[0].Header.Get("User-Agent"), " resource/mlflow_model auth/"+tc.authType+" ", call.name)
				assert.Equal(t, strings.TrimPrefix(ws.URL, "http://"), received[0].Host, call.name)
			}
			served, tokens := ws.counts()
			assert.Equal(t, tc.tokens, tokens, "both clients use the token of the config")
			assert.Empty(t, elsewhere.received(), "a call went to the workspace of the environment")
			if tc.httpTransport {
				assert.EqualValues(t, served, through.sent.Load(), "every request of both clients goes over cfg.HTTPTransport")
				assert.Equal(t, 2*len(newClientCalls), served)
			}
		})
	}
}

func TestNewClient_Errors(t *testing.T) {
	failed := errors.New("step failed")
	ws := newClientStart(t, &newClientWorkspace{})
	tests := []struct {
		name         string
		cfg          *config.Config
		noConstruct  bool
		constructErr error
		spoil        func(*newClientSteps) // makes a helper fail
		want         string                // how the error starts
		wantIs       error
	}{
		{name: "nil config", want: "modular SDK client: config is nil"},
		{name: "nil constructor", cfg: ws.config(nil), noConstruct: true, want: "modular SDK client: constructor is nil"},
		{name: "nil config and nil constructor", noConstruct: true, want: "modular SDK client: config is nil"},
		{name: "Transport fails: the config doesn't resolve",
			cfg:  &config.Config{Loaders: []config.Loader{newClientLoader{err: failed}}},
			want: "modular SDK client: Transport: resolve: step failed", wantIs: failed},
		{name: "AttemptTimeout fails", cfg: ws.config(nil),
			spoil: func(steps *newClientSteps) {
				steps.attemptTimeout = func(*config.Config, http.RoundTripper) (http.RoundTripper, error) { return nil, failed }
			},
			want: "modular SDK client: AttemptTimeout: step failed", wantIs: failed},
		{name: "ClientOptions fails: the config has no host",
			cfg:  &config.Config{Token: "dapi-secret", Loaders: []config.Loader{newClientLoader{}}},
			want: "modular SDK client: ClientOptions: modular SDK client options: config has no host"},
		{name: "constructor fails", cfg: ws.config(nil), constructErr: failed,
			want: "modular SDK client: constructor: step failed", wantIs: failed},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			constructed := 0
			construct := func(context.Context, ...client.Option) (*modelregistry.Client, error) {
				constructed++
				return &modelregistry.Client{}, tc.constructErr
			}
			if tc.noConstruct {
				construct = nil
			}
			build := func() (*modelregistry.Client, error) { return NewClient(t.Context(), tc.cfg, construct) }
			if tc.spoil != nil { // NewClient has the real helpers, which don't fail here.
				steps := newClientHelpers
				tc.spoil(&steps)
				build = func() (*modelregistry.Client, error) { return newClientBuild(t.Context(), tc.cfg, construct, steps) }
			}
			built, err := build()
			require.Error(t, err)
			assert.True(t, strings.HasPrefix(err.Error(), tc.want), "%q doesn't start with %q", err, tc.want)
			if tc.wantIs != nil {
				assert.ErrorIs(t, err, tc.wantIs)
			}
			assert.Nil(t, built, "the zero value comes back with an error")
			assert.Equal(t, tc.constructErr != nil, constructed == 1, "the constructor runs only after every step succeeded")
		})
	}
	// The zero value of a type that is not a pointer.
	built, err := NewClient(t.Context(), ws.config(nil), func(context.Context, ...client.Option) (int, error) { return 7, failed })
	assert.ErrorIs(t, err, failed)
	assert.Zero(t, built)
}

func TestNewClient_HandsEachHelperWhatItNeeds(t *testing.T) {
	type key struct{}
	ctx, cfg := context.WithValue(t.Context(), key{}, "caller"), &config.Config{}
	credentials, bare, watched := Credentials(cfg), &newClientCounter{}, &newClientCounter{}
	var order []string
	step := func(name string, got *config.Config) {
		order = append(order, name)
		assert.Same(t, cfg, got, name)
	}
	built, err := newClientBuild(ctx, cfg, func(got context.Context, options ...client.Option) (string, error) {
		order = append(order, "constructor")
		assert.Equal(t, "caller", got.Value(key{}), "the context of the caller")
		applied := internaloptions.ClientOptions{}
		require.Len(t, options, 1)
		require.NoError(t, options[0](&applied))
		return applied.Host, nil
	}, newClientSteps{
		credentials: func(got *config.Config) auth.Credentials {
			step("Credentials", got)
			return credentials
		},
		transport: func(got *config.Config) (http.RoundTripper, error) {
			step("Transport", got)
			return bare, nil
		},
		attemptTimeout: func(got *config.Config, transport http.RoundTripper) (http.RoundTripper, error) {
			step("AttemptTimeout", got)
			assert.Same(t, bare, transport, "AttemptTimeout goes around the transport of Transport")
			return watched, nil
		},
		clientOptions: func(got *config.Config, gotCredentials auth.Credentials, transport http.RoundTripper) ([]client.Option, error) {
			step("ClientOptions", got)
			assert.Equal(t, credentials, gotCredentials)
			assert.Same(t, watched, transport, "the client sends over the transport of AttemptTimeout")
			return []client.Option{client.WithHost("https://from-client-options")}, nil
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "https://from-client-options", built, "the constructor gets the options of ClientOptions")
	assert.Equal(t, []string{"Transport", "AttemptTimeout", "Credentials", "ClientOptions", "constructor"}, order)
}

func TestNewClient_EndsStalledAttemptLikeClassicClient(t *testing.T) {
	ws := newClientStart(t, &newClientWorkspace{answer: func(_ http.ResponseWriter, r *http.Request, _ int) { <-r.Context().Done() }})
	cfg := ws.config(func(cfg *config.Config) { cfg.HTTPTimeoutSeconds, cfg.RetryTimeoutSeconds = 1, 5 })
	classic := newClientClassic(t, cfg)
	modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
	require.NoError(t, err)
	errs, took := make([]error, 2), make([]time.Duration, 2)
	var wg sync.WaitGroup
	for i, call := range []func() (any, error){
		func() (any, error) { return newClientCalls[1].classic(t.Context(), classic) },
		func() (any, error) { return newClientCalls[1].modular(t.Context(), modular) },
	} {
		wg.Go(func() {
			start := time.Now()
			_, errs[i] = call()
			took[i] = time.Since(start)
		})
	}
	wg.Wait()
	want := `Get "` + ws.URL + `/api/2.0/mlflow/databricks/registered-models/get?name=my+model%2F%C3%BC%26x%3D1": request timed out after 1s of inactivity`
	for i, who := range []string{"classic", "modular"} {
		assert.EqualError(t, errs[i], want, who)
		assert.GreaterOrEqual(t, took[i], time.Second, who)
	}
	assert.Len(t, ws.received(), 2, "neither client tries again")
}

func TestNewClient_KeepsNoState(t *testing.T) {
	tests := []struct {
		name            string
		useTLS          bool
		configure       func(*config.Config)
		wantConnections int
		wantTokens      int
	}{
		{name: "clients share the token and the connections", configure: newClientM2M, wantConnections: 1, wantTokens: 1},
		{name: "InsecureSkipVerify gives every client a pool of its own", useTLS: true,
			configure: func(cfg *config.Config) { cfg.InsecureSkipVerify = true }, wantConnections: 3},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := newClientStart(t, &newClientWorkspace{useTLS: tc.useTLS})
			cfg := ws.config(tc.configure)
			clients, connections := map[*modelregistry.Client]bool{}, map[string]bool{}
			for range 3 {
				modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
				require.NoError(t, err)
				clients[modular] = true
				for _, call := range newClientCalls {
					_, err := call.modular(t.Context(), modular)
					require.NoError(t, err, call.name)
				}
			}
			for _, request := range ws.received() {
				connections[request.conn] = true
			}
			served, tokens := ws.counts()
			assert.Len(t, clients, 3, "every call builds a new client")
			assert.Len(t, connections, tc.wantConnections)
			assert.Equal(t, tc.wantTokens, tokens)
			var err error
			goroutines := runtime.NumGoroutine()
			allocations := testing.AllocsPerRun(100, func() {
				_, err = NewClient(t.Context(), cfg, modelregistry.NewClient)
			})
			require.NoError(t, err)
			servedAfter, _ := ws.counts()
			assert.Equal(t, served, servedAfter, "a call sends no request")
			assert.LessOrEqual(t, runtime.NumGoroutine(), goroutines, "a call starts no goroutine")
			assert.Less(t, allocations, 100.0, "allocations of a call")
		})
	}
}

func TestNewClient_ConcurrentWithClassicCalls(t *testing.T) {
	ws := newClientStart(t, &newClientWorkspace{})
	cfg := ws.config(newClientM2M)
	classic := newClientClassic(t, cfg)
	const goroutines, rounds = 16, 4
	var wg sync.WaitGroup
	for i := range goroutines {
		wg.Go(func() {
			for round := range rounds {
				call := newClientCalls[(i+round)%len(newClientCalls)]
				if i%4 == 0 {
					_, err := call.classic(t.Context(), classic)
					assert.NoError(t, err, call.name)
					continue
				}
				modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
				if assert.NoError(t, err) {
					_, err = call.modular(t.Context(), modular)
					assert.NoError(t, err, call.name)
				}
			}
		})
	}
	wg.Wait()
	received := ws.received()
	require.Len(t, received, goroutines*rounds)
	fromModular := 0
	for _, request := range received {
		assert.Equal(t, []string{"Bearer token-1"}, request.Header.Values("Authorization"))
		assert.Equal(t, []string{"42"}, request.Header.Values("X-Databricks-Workspace-Id"))
		assert.Contains(t, request.Header.Get("User-Agent"), " auth/oauth-m2m")
		if strings.Contains(request.Header.Get("User-Agent"), " sdk-go-modelregistry/") {
			fromModular++
		}
	}
	_, tokens := ws.counts()
	assert.Equal(t, goroutines*rounds*3/4, fromModular)
	assert.Equal(t, 1, tokens, "every client uses the one token of the config")
}

// TestNewClient_DifferencesFromClassicClient pins what a client of NewClient
// does differently from the classic client, as the pull request lists it.
func TestNewClient_DifferencesFromClassicClient(t *testing.T) {
	clients := func(t *testing.T, answer func(http.ResponseWriter, *http.Request, int)) (*newClientWorkspace, *databricks.WorkspaceClient, *modelregistry.Client) {
		ws := newClientStart(t, &newClientWorkspace{answer: answer})
		cfg := ws.config(nil)
		classic := newClientClassic(t, cfg)
		modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
		require.NoError(t, err)
		return ws, classic, modular
	}
	t.Run("the keys of a JSON body come in another order", func(t *testing.T) {
		ws, classic, modular := clients(t, nil)
		_, err := newClientCalls[0].classic(t.Context(), classic)
		require.NoError(t, err)
		_, err = newClientCalls[0].modular(t.Context(), modular)
		require.NoError(t, err)
		received := ws.received()
		require.Len(t, received, 2)
		name, _ := json.Marshal(newClientModel) // Escapes the ampersand.
		assert.Equal(t, `{"description":"d","name":`+string(name)+`,"tags":[{"key":"k","value":"v"}]}`, received[0].Body)
		assert.Equal(t, `{"name":`+string(name)+`,"tags":[{"key":"k","value":"v"}],"description":"d"}`, received[1].Body)
	})
	t.Run("a failed call returns the error of the modular SDK", func(t *testing.T) {
		_, classic, modular := clients(t, func(w http.ResponseWriter, _ *http.Request, _ int) {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, `{"error_code":"RESOURCE_DOES_NOT_EXIST","message":"no such model"}`)
		})
		_, classicErr := newClientCalls[1].classic(t.Context(), classic)
		_, modularErr := newClientCalls[1].modular(t.Context(), modular)
		assert.EqualError(t, classicErr, "no such model")
		assert.EqualError(t, modularErr, "databricks-api error: no such model")
		assert.ErrorIs(t, classicErr, classicapierr.ErrNotFound)
		assert.NotErrorIs(t, modularErr, classicapierr.ErrNotFound)
		var asClassic *classicapierr.APIError
		var asModular *modularapierr.APIError
		assert.False(t, errors.As(modularErr, &asClassic), "the modular error is a classic APIError")
		require.ErrorAs(t, modularErr, &asModular)
		assert.Equal(t, http.StatusNotFound, asModular.HTTPStatusCode())
	})
	t.Run("an attempt that fails with 503 is not retried", func(t *testing.T) {
		ws, classic, modular := clients(t, func(w http.ResponseWriter, _ *http.Request, before int) {
			if before%2 == 0 { // The first attempt of each call.
				w.WriteHeader(http.StatusServiceUnavailable)
				io.WriteString(w, `{"error_code":"TEMPORARILY_UNAVAILABLE","message":"try again"}`)
				return
			}
			io.WriteString(w, newClientAnswer)
		})
		_, err := newClientCalls[1].classic(t.Context(), classic)
		require.NoError(t, err)
		assert.Len(t, ws.received(), 2, "the classic client tries again")
		_, err = newClientCalls[1].modular(t.Context(), modular)
		assert.EqualError(t, err, "databricks-api error: try again")
		assert.Len(t, ws.received(), 1, "the modular client tried again")
	})
}
