package sdkgoutility

import (
	"cmp"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth"
	"github.com/databricks/sdk-go/core/ops"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/sdk-go/options/client"
	"github.com/databricks/sdk-go/options/internaloptions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// clientOptionsTestHost is a host in the form that a resolved config has.
	clientOptionsTestHost = "https://workspace.cloud.databricks.com"
	// clientOptionsModelBody is the answer of the test servers to every call.
	clientOptionsModelBody         = `{"registered_model_databricks":{"name":"m"}}`
	clientOptionsWorkspaceIDHeader = "X-Databricks-Workspace-Id"
	// clientOptionsTransportHeader marks what clientOptionsPlainTransport sends.
	clientOptionsTransportHeader = "X-Test-Sent-Over"
)

// clientOptionsFakeCredentials sign a request with fixed headers.
type clientOptionsFakeCredentials struct{ headers []auth.Header }

func (c clientOptionsFakeCredentials) Name() string { return "fake" }

func (c clientOptionsFakeCredentials) AuthHeaders(context.Context) ([]auth.Header, error) {
	return c.headers, nil
}

var clientOptionsTestCredentials = clientOptionsFakeCredentials{headers: []auth.Header{
	{Key: "Authorization", Value: "Bearer fake-token"},
	{Key: "X-Databricks-Azure-SP-Management-Token", Value: "fake-management-token"},
}}

// clientOptionsNilMap are credentials whose zero value is a nil map.
type clientOptionsNilMap map[string]string

func (c clientOptionsNilMap) Name() string { return "map" }

func (c clientOptionsNilMap) AuthHeaders(context.Context) ([]auth.Header, error) { return nil, nil }

// clientOptionsRoundTripFunc is a transport whose zero value is a nil func.
type clientOptionsRoundTripFunc func(*http.Request) (*http.Response, error)

func (f clientOptionsRoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// clientOptionsPlainTransport notes the deadline of every request, zero for
// none, marks the request with clientOptionsTransportHeader and sends it.
type clientOptionsPlainTransport struct {
	mu        sync.Mutex
	deadlines []time.Time
}

func (t *clientOptionsPlainTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	deadline, _ := r.Context().Deadline()
	t.mu.Lock()
	t.deadlines = append(t.deadlines, deadline)
	t.mu.Unlock()
	marked := r.Clone(r.Context())
	marked.Header.Set(clientOptionsTransportHeader, "plain")
	return http.DefaultTransport.RoundTrip(marked)
}

// clientOptionsServer records the requests that it gets and answers them all
// the same, after delay, unless the client gives up first.
type clientOptionsServer struct {
	*httptest.Server
	mu       sync.Mutex
	requests []*http.Request
}

func clientOptionsNewServer(t *testing.T, status int, body string, delay time.Duration) *clientOptionsServer {
	t.Helper()
	s := &clientOptionsServer{}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		s.requests = append(s.requests, r)
		s.mu.Unlock()
		select {
		case <-time.After(delay):
		case <-r.Context().Done():
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(s.Close)
	return s
}

// received returns the requests that came in since the last call.
func (s *clientOptionsServer) received() []*http.Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	requests := s.requests
	s.requests = nil
	return requests
}

// clientOptionsCleanEnvironment hides the Databricks configuration of the
// machine from a test: every variable that the classic or the modular SDK
// reads, and the config file in the home directory.
func clientOptionsCleanEnvironment(t *testing.T) {
	t.Helper()
	for _, attribute := range config.ConfigAttributes {
		for _, name := range attribute.EnvVars {
			t.Setenv(name, "")
		}
	}
	for _, entry := range os.Environ() {
		if name, _, _ := strings.Cut(entry, "="); strings.HasPrefix(name, "DATABRICKS_") {
			t.Setenv(name, "")
		}
	}
	t.Setenv("HOME", t.TempDir())
}

// clientOptionsClassicClient returns a real classic client for cfg. Building
// it resolves and authenticates its Config, which then keeps its host.
func clientOptionsClassicClient(t *testing.T, cfg *config.Config) *databricks.WorkspaceClient {
	t.Helper()
	cfg.Token = "classic-token"
	cfg.HostMetadataResolver = func(context.Context, string) (*config.HostMetadata, error) {
		return nil, nil
	}
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	require.Equal(t, "pat", cfg.AuthType, "building the client didn't authenticate its config")
	return w
}

// clientOptionsModularClient returns a real client built from ClientOptions.
func clientOptionsModularClient(t *testing.T, cfg *config.Config, transport http.RoundTripper) *modelregistry.Client {
	t.Helper()
	options, err := ClientOptions(cfg, clientOptionsTestCredentials, transport)
	require.NoError(t, err)
	c, err := modelregistry.NewClient(context.Background(), options...)
	require.NoError(t, err)
	return c
}

// clientOptionsGetModel gets the model m and returns its name.
func clientOptionsGetModel(ctx context.Context, c *modelregistry.Client, options ...call.Option) (string, error) {
	name := "m"
	model, err := c.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: &name}, options...)
	if err != nil {
		return "", err
	}
	return *model.RegisteredModelDatabricks.Name, nil
}

func TestClientOptions_Errors(t *testing.T) {
	const (
		nilConfig      = "modular SDK client options: config is nil"
		nilCredentials = "modular SDK client options: credentials are nil"
		nilTransport   = "modular SDK client options: transport is nil"
		noHost         = "modular SDK client options: config has no host"
	)
	transport := &clientOptionsPlainTransport{}
	resolved := &config.Config{Host: clientOptionsTestHost}
	tests := []struct {
		name        string
		cfg         *config.Config
		credentials auth.Credentials
		transport   http.RoundTripper
		// host, when set, is the host of a config that is not resolved, and
		// stands for cfg, credentials, transport and want.
		host string
		want string
	}{
		{name: "nil config", credentials: clientOptionsTestCredentials, transport: transport, want: nilConfig},
		{name: "nil config comes first", want: nilConfig},
		{name: "nil credentials", cfg: resolved, transport: transport, want: nilCredentials},
		{name: "nil credentials come before a nil transport", cfg: resolved, want: nilCredentials},
		{name: "nil pointer as credentials", cfg: resolved, credentials: (*clientOptionsFakeCredentials)(nil), transport: transport, want: nilCredentials},
		{name: "nil map as credentials", cfg: resolved, credentials: clientOptionsNilMap(nil), transport: transport, want: nilCredentials},
		{name: "nil transport", cfg: resolved, credentials: clientOptionsTestCredentials, want: nilTransport},
		{name: "nil pointer as transport", cfg: resolved, credentials: clientOptionsTestCredentials, transport: (*http.Transport)(nil), want: nilTransport},
		{name: "nil func as transport", cfg: resolved, credentials: clientOptionsTestCredentials, transport: clientOptionsRoundTripFunc(nil), want: nilTransport},
		{name: "nil transport comes before the host", cfg: &config.Config{}, credentials: clientOptionsTestCredentials, want: nilTransport},
		{name: "no host", cfg: &config.Config{WorkspaceID: "1", AccountID: "a"}, credentials: clientOptionsTestCredentials, transport: transport, want: noHost},
		{name: "host without scheme", host: "workspace.cloud.databricks.com"},
		{name: "host and port without scheme", host: "localhost:8080"},
		{name: "host that starts with slashes", host: "//workspace.cloud.databricks.com"},
		{name: "scheme without host", host: "https:"},
		{name: "scheme and slashes without host", host: "https://"},
		{name: "host that is no URL", host: "://workspace"},
		{name: "host with a control character", host: "https://workspace\n"},
		{name: "host with a trailing slash", host: clientOptionsTestHost + "/"},
		{name: "host with a path", host: clientOptionsTestHost + "/api/2.0"},
		{name: "host with the workspace ID in the query", host: clientOptionsTestHost + "?o=123"},
		{name: "host with a fragment", host: clientOptionsTestHost + "#home"},
		{name: "host with a user", host: "https://user@workspace.cloud.databricks.com"},
		{name: "host with a scheme in capitals", host: "HTTPS://workspace.cloud.databricks.com"},
		{name: "host with spaces around it", host: " " + clientOptionsTestHost + " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.host != "" {
				tt.cfg, tt.credentials, tt.transport = &config.Config{Host: tt.host}, clientOptionsTestCredentials, transport
				tt.want = "modular SDK client options: config host " + strconv.Quote(tt.host) +
					" is not of the form scheme://host, resolve the config first"
			}
			options, err := ClientOptions(tt.cfg, tt.credentials, tt.transport)
			assert.EqualError(t, err, tt.want)
			assert.Nil(t, options)
		})
	}
}

func TestClientOptions_SetEveryOption(t *testing.T) {
	clientOptionsCleanEnvironment(t)
	transport := &clientOptionsPlainTransport{}
	tests := []struct {
		name        string
		cfg         *config.Config
		credentials auth.Credentials // clientOptionsTestCredentials when nil
	}{
		{name: "host only", cfg: &config.Config{Host: clientOptionsTestHost}},
		{name: "workspace ID", cfg: &config.Config{Host: clientOptionsTestHost, WorkspaceID: "1234567890123456"}},
		{name: "account ID", cfg: &config.Config{Host: clientOptionsTestHost, AccountID: "8f4c2a10-account"}},
		{name: "workspace ID and account ID", cfg: &config.Config{Host: clientOptionsTestHost, WorkspaceID: "conn-1a2b", AccountID: "8f4c2a10-account"}},
		{name: "host with an IPv6 address and a port", cfg: &config.Config{Host: "http://[::1]:8080"}},
		{name: "credentials that are a pointer", cfg: &config.Config{Host: clientOptionsTestHost}, credentials: &clientOptionsTestCredentials},
		{name: "credentials that are a map", cfg: &config.Config{Host: clientOptionsTestHost}, credentials: clientOptionsNilMap{}},
		{name: "retry timeout of a second is not carried", cfg: &config.Config{Host: clientOptionsTestHost, RetryTimeoutSeconds: 1}},
		{name: "retry timeout of an hour is not carried", cfg: &config.Config{Host: clientOptionsTestHost, RetryTimeoutSeconds: 3600}},
		{name: "negative retry timeout is not carried", cfg: &config.Config{Host: clientOptionsTestHost, RetryTimeoutSeconds: -1}},
		{name: "HTTP timeout and rate limit are not carried", cfg: &config.Config{Host: clientOptionsTestHost, HTTPTimeoutSeconds: 65, RateLimitPerSecond: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.credentials == nil {
				tt.credentials = clientOptionsTestCredentials
			}
			clientOptionsClassicClient(t, tt.cfg)
			options, err := ClientOptions(tt.cfg, tt.credentials, transport)
			require.NoError(t, err)

			// What a modular client constructor does with its options.
			var got internaloptions.ClientOptions
			for _, option := range options {
				require.NoError(t, option(&got))
			}
			require.NoError(t, got.Resolve())

			assert.Equal(t, tt.cfg.Host, got.Host)
			assert.Equal(t, tt.cfg.WorkspaceID, got.WorkspaceID)
			assert.Equal(t, tt.cfg.AccountID, got.AccountID)
			assert.Equal(t, tt.credentials, got.Credentials)
			assert.Same(t, transport, got.HTTPClient.Transport)
			assert.Zero(t, got.HTTPClient.Timeout)
			assert.Nil(t, got.HTTPClient.CheckRedirect)
			assert.Nil(t, got.HTTPClient.Jar)
			assert.True(t, got.DisableConfigFile)
			assert.True(t, got.DisableEnv)
			assert.Empty(t, got.ConfigFile)
			assert.Empty(t, got.ProfileName)
			require.NotNil(t, got.Timeout, "the client keeps the timeout of each method")
			assert.Zero(t, *got.Timeout, "the client sets a timeout of its own")
		})
	}
}

func TestClientOptions_CallsLikeClassicClient(t *testing.T) {
	clientOptionsCleanEnvironment(t)
	ctx := context.Background()
	tests := []struct {
		name string
		// host follows the URL of the server in the host that the classic
		// config is given, before it is resolved.
		host            string
		workspaceID     string
		accountID       string
		wantWorkspaceID []string
	}{
		{name: "no workspace ID"},
		{name: "numeric workspace ID", workspaceID: "1234567890123456", wantWorkspaceID: []string{"1234567890123456"}},
		{name: "connection ID as workspace ID", workspaceID: "conn-1a2b", wantWorkspaceID: []string{"conn-1a2b"}},
		{name: "account ID without workspace ID", accountID: "8f4c2a10-account"},
		{name: "workspace ID and account ID", workspaceID: "42", accountID: "8f4c2a10-account", wantWorkspaceID: []string{"42"}},
		{name: "host resolved from a trailing slash", host: "/"},
		{name: "workspace ID resolved from the query of the host", host: "/?o=77", wantWorkspaceID: []string{"77"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody, 0)
			w := clientOptionsClassicClient(t, &config.Config{Host: server.URL + tt.host, WorkspaceID: tt.workspaceID, AccountID: tt.accountID})
			classicModel, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "m"})
			require.NoError(t, err)
			classicRequests := server.received()
			require.Len(t, classicRequests, 1)

			transport := &clientOptionsPlainTransport{}
			modularModel, err := clientOptionsGetModel(ctx, clientOptionsModularClient(t, w.Config, transport))
			require.NoError(t, err)
			modularRequests := server.received()
			require.Len(t, modularRequests, 1)
			got, want := modularRequests[0], classicRequests[0]

			// The same call arrives at the same server, routed the same way.
			assert.Equal(t, strings.TrimPrefix(server.URL, "http://"), got.Host)
			assert.Equal(t, want.Host, got.Host)
			assert.Equal(t, want.Method, got.Method)
			assert.Equal(t, "/api/2.0/mlflow/databricks/registered-models/get?name=m", got.URL.String())
			assert.Equal(t, want.URL.String(), got.URL.String())
			assert.Equal(t, tt.wantWorkspaceID, got.Header.Values(clientOptionsWorkspaceIDHeader))
			assert.Equal(t, want.Header.Values(clientOptionsWorkspaceIDHeader), got.Header.Values(clientOptionsWorkspaceIDHeader))
			assert.Equal(t, classicModel.RegisteredModelDatabricks.Name, modularModel)

			// It is signed by the credentials and not by the classic config.
			assert.Equal(t, []string{"Bearer classic-token"}, want.Header.Values("Authorization"))
			assert.Equal(t, []string{"Bearer fake-token"}, got.Header.Values("Authorization"))
			assert.Equal(t, []string{"fake-management-token"}, got.Header.Values("X-Databricks-Azure-SP-Management-Token"))

			// It went over the transport.
			assert.Equal(t, []string{"plain"}, got.Header.Values(clientOptionsTransportHeader))
			assert.Len(t, transport.deadlines, 1)
		})
	}
}

func TestClientOptions_IgnoresEnvironment(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name string
		// env are the variables to set and files the files to write in the
		// home directory. {elsewhere} is the URL of another server than the
		// workspace and {home} is the home directory.
		env   map[string]string
		files map[string]string
		// A modular client without the options fails in NewClient with
		// leftAloneError, or calls elsewhere with these headers.
		leftAloneError         string
		leftAloneAuthorization string
		leftAloneWorkspaceID   []string
	}{
		{name: "DATABRICKS_HOST, DATABRICKS_TOKEN, DATABRICKS_WORKSPACE_ID and DATABRICKS_ACCOUNT_ID",
			env: map[string]string{"DATABRICKS_HOST": "{elsewhere}", "DATABRICKS_TOKEN": "env-token",
				"DATABRICKS_WORKSPACE_ID": "999", "DATABRICKS_ACCOUNT_ID": "env-account"},
			leftAloneAuthorization: "Bearer env-token",
			leftAloneWorkspaceID:   []string{"999"}},
		{name: "config file in the home directory",
			files:                  map[string]string{".databrickscfg": "[DEFAULT]\nhost = {elsewhere}\ntoken = home-token\nworkspace_id = 888\n"},
			leftAloneAuthorization: "Bearer home-token",
			leftAloneWorkspaceID:   []string{"888"}},
		{name: "DATABRICKS_CONFIG_PROFILE",
			env:                    map[string]string{"DATABRICKS_CONFIG_PROFILE": "elsewhere"},
			files:                  map[string]string{".databrickscfg": "[elsewhere]\nhost = {elsewhere}\ntoken = profile-token\n"},
			leftAloneAuthorization: "Bearer profile-token"},
		{name: "DATABRICKS_CONFIG_FILE",
			env:                    map[string]string{"DATABRICKS_CONFIG_FILE": "{home}/other.cfg"},
			files:                  map[string]string{"other.cfg": "[DEFAULT]\nhost = {elsewhere}\ntoken = file-token\n"},
			leftAloneAuthorization: "Bearer file-token"},
		{name: "DATABRICKS_CONFIG_FILE that doesn't exist",
			env:            map[string]string{"DATABRICKS_CONFIG_FILE": "{home}/missing.cfg"},
			leftAloneError: "config file not found"},
		{name: "DATABRICKS_CONFIG_PROFILE that doesn't exist",
			env:            map[string]string{"DATABRICKS_CONFIG_PROFILE": "missing"},
			files:          map[string]string{".databrickscfg": "[DEFAULT]\nhost = {elsewhere}\ntoken = home-token\n"},
			leftAloneError: "profile not found"},
		{name: "all of them",
			env: map[string]string{"DATABRICKS_HOST": "{elsewhere}", "DATABRICKS_TOKEN": "env-token",
				"DATABRICKS_WORKSPACE_ID": "999", "DATABRICKS_ACCOUNT_ID": "env-account",
				"DATABRICKS_CONFIG_PROFILE": "elsewhere", "DATABRICKS_CONFIG_FILE": "{home}/other.cfg"},
			files: map[string]string{
				".databrickscfg": "[DEFAULT]\nhost = {elsewhere}\ntoken = home-token\n",
				"other.cfg":      "[elsewhere]\nhost = {elsewhere}\ntoken = profile-token\nworkspace_id = 888\n",
			},
			leftAloneAuthorization: "Bearer env-token",
			leftAloneWorkspaceID:   []string{"999"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientOptionsCleanEnvironment(t)
			workspace := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody, 0)
			elsewhere := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody, 0)
			// The provider resolves the classic config before it builds a
			// modular client, so here before the environment points elsewhere.
			configs := []*config.Config{
				clientOptionsClassicClient(t, &config.Config{Host: workspace.URL}).Config,
				clientOptionsClassicClient(t, &config.Config{Host: workspace.URL, WorkspaceID: "123", AccountID: "config-account"}).Config,
			}
			home := os.Getenv("HOME")
			fill := strings.NewReplacer("{elsewhere}", elsewhere.URL, "{home}", home).Replace
			for path, content := range tt.files {
				require.NoError(t, os.WriteFile(filepath.Join(home, path), []byte(fill(content)), 0o600))
			}
			for name, value := range tt.env {
				t.Setenv(name, fill(value))
			}

			// The environment does divert a modular client that is left alone.
			leftAlone, err := modelregistry.NewClient(ctx, client.WithHTTPClient(&http.Client{Transport: &clientOptionsPlainTransport{}}))
			if tt.leftAloneError != "" {
				require.ErrorContains(t, err, tt.leftAloneError)
			} else {
				require.NoError(t, err)
				_, err = clientOptionsGetModel(ctx, leftAlone)
				require.NoError(t, err)
				diverted := elsewhere.received()
				require.Len(t, diverted, 1)
				assert.Equal(t, []string{tt.leftAloneAuthorization}, diverted[0].Header.Values("Authorization"))
				assert.Equal(t, tt.leftAloneWorkspaceID, diverted[0].Header.Values(clientOptionsWorkspaceIDHeader))
				assert.Empty(t, workspace.received())
			}

			for _, cfg := range configs {
				name, err := clientOptionsGetModel(ctx, clientOptionsModularClient(t, cfg, &clientOptionsPlainTransport{}))
				require.NoError(t, err)
				assert.Equal(t, "m", name)
				requests := workspace.received()
				require.Len(t, requests, 1)
				assert.Equal(t, strings.TrimPrefix(workspace.URL, "http://"), requests[0].Host)
				assert.Equal(t, []string{"Bearer fake-token"}, requests[0].Header.Values("Authorization"))
				var wantWorkspaceID []string
				if cfg.WorkspaceID != "" {
					wantWorkspaceID = []string{cfg.WorkspaceID}
				}
				assert.Equal(t, wantWorkspaceID, requests[0].Header.Values(clientOptionsWorkspaceIDHeader))
				assert.Empty(t, elsewhere.received())
			}
		})
	}
}

// A call has no deadline but the one of the context of the caller, whatever
// cfg.RetryTimeoutSeconds is, as the attempts of a classic call have.
func TestClientOptions_CallTimeout(t *testing.T) {
	clientOptionsCleanEnvironment(t)
	const slow = 1300 * time.Millisecond
	tests := []struct {
		name                string
		retryTimeoutSeconds int
		callerTimeout       time.Duration // of the context of the caller, if any
		delay               time.Duration // of the answer of the server
		// want is how far ahead the deadline of the request is, zero for none.
		want    time.Duration
		wantErr error
	}{
		{name: "unset retry timeout is no deadline", retryTimeoutSeconds: 0},
		{name: "retry timeout of 1 second is no deadline", retryTimeoutSeconds: 1},
		{name: "minus one, as the provider sets it, is no deadline", retryTimeoutSeconds: -1},
		{name: "any negative retry timeout is no deadline", retryTimeoutSeconds: -300},
		{name: "deadline of the caller before the retry timeout", retryTimeoutSeconds: 60, callerTimeout: 10 * time.Second, want: 10 * time.Second},
		{name: "deadline of the caller after the retry timeout", retryTimeoutSeconds: 60, callerTimeout: time.Hour, want: time.Hour},
		{name: "deadline of the caller after the 5 minutes of a modular method", callerTimeout: 20 * time.Minute, want: 20 * time.Minute},
		{name: "retry timeout doesn't end a slower attempt", retryTimeoutSeconds: 1, delay: slow},
		{name: "deadline of the caller ends a slower attempt after the retry timeout", retryTimeoutSeconds: 1,
			callerTimeout: slow, delay: 2 * slow, want: slow, wantErr: context.DeadlineExceeded},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody, tt.delay)
			transport := &clientOptionsPlainTransport{}
			w := clientOptionsClassicClient(t, &config.Config{Host: server.URL, RetryTimeoutSeconds: tt.retryTimeoutSeconds})
			c := clientOptionsModularClient(t, w.Config, transport)
			before := time.Now()
			ctx := context.Background()
			if tt.callerTimeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tt.callerTimeout)
				defer cancel()
			}

			_, err := clientOptionsGetModel(ctx, c)
			after := time.Now()

			require.ErrorIs(t, err, tt.wantErr)
			// The attempt ran until the server answered or the caller gave up.
			assert.GreaterOrEqual(t, after.Sub(before), min(tt.delay, cmp.Or(tt.callerTimeout, tt.delay)))
			require.Len(t, transport.deadlines, 1)
			deadline := transport.deadlines[0]
			if tt.want == 0 {
				assert.Zero(t, deadline, "the request has a deadline")
				return
			}
			assert.False(t, deadline.Before(before.Add(tt.want)), "the deadline is less than %s ahead", tt.want)
			assert.False(t, deadline.After(after.Add(tt.want)), "the deadline is more than %s ahead", tt.want)
		})
	}
}

// Both clients get cfg.RetryTimeoutSeconds of a second, and neither ends an
// attempt that outlasts it. Ending the retries is left to the retrier of a
// call: with one that never stops, as here, the modular call goes on until ctx.
func TestClientOptions_LeavesRetryTimeoutToRetrier(t *testing.T) {
	const callerTimeout = 2 * time.Second // of the modular call only
	tests := []struct {
		name                             string
		status                           int
		body                             string
		delay                            time.Duration // of every answer of the server
		wantClassicErr, wantModularErr   error         // nil when the call succeeds
		wantClassicTook, wantModularTook time.Duration // at least, and less than 2 seconds more
	}{
		{name: "one attempt outlasts the retry timeout", status: http.StatusOK, body: clientOptionsModelBody,
			delay: 1300 * time.Millisecond, wantClassicTook: 1300 * time.Millisecond, wantModularTook: 1300 * time.Millisecond},
		{name: "every attempt is throttled", status: http.StatusTooManyRequests, body: `{"error_code":"REQUEST_LIMIT_EXCEEDED","message":"slow down"}`,
			wantClassicErr: apierr.ErrTooManyRequests, wantModularErr: context.DeadlineExceeded, wantClassicTook: time.Second, wantModularTook: callerTimeout},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientOptionsCleanEnvironment(t)
			server := clientOptionsNewServer(t, tt.status, tt.body, tt.delay)
			w := clientOptionsClassicClient(t, &config.Config{Host: server.URL, RetryTimeoutSeconds: 1})
			c := clientOptionsModularClient(t, w.Config, &clientOptionsPlainTransport{})
			// A modular client retries only with a retrier, a call option.
			retry := call.WithRetrier(func() ops.Retrier {
				backoff := ops.BackoffPolicy{Initial: 50 * time.Millisecond, Maximum: 50 * time.Millisecond}
				return ops.RetryIf(backoff, func(error) bool { return true })
			})
			ctx, cancel := context.WithTimeout(context.Background(), callerTimeout)
			defer cancel()

			var classicErr, modularErr error
			var classicTook, modularTook time.Duration
			var wg sync.WaitGroup
			wg.Go(func() {
				start := time.Now()
				_, classicErr = w.ModelRegistry.GetModel(context.Background(), ml.GetModelRequest{Name: "m"})
				classicTook = time.Since(start)
			})
			wg.Go(func() {
				start := time.Now()
				_, modularErr = clientOptionsGetModel(ctx, c, retry)
				modularTook = time.Since(start)
			})
			wg.Wait()

			assert.ErrorIs(t, classicErr, tt.wantClassicErr)
			assert.ErrorIs(t, modularErr, tt.wantModularErr)
			assert.GreaterOrEqual(t, classicTook, tt.wantClassicTook, "the classic call ended early")
			assert.Less(t, classicTook, tt.wantClassicTook+2*time.Second, "the classic call went on")
			assert.GreaterOrEqual(t, modularTook, tt.wantModularTook, "the modular call ended early")
			assert.Less(t, modularTook, tt.wantModularTook+2*time.Second, "the modular call went on")
		})
	}
}

// Resources build their options while others call with the classic client of
// the same config. Under -race this fails if ClientOptions writes the config.
func TestClientOptions_Concurrent(t *testing.T) {
	clientOptionsCleanEnvironment(t)
	ctx := context.Background()
	server := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody, 0)
	w := clientOptionsClassicClient(t, &config.Config{Host: server.URL, WorkspaceID: "42"})
	const builders, classicCalls = 32, 10
	errs := make(chan error, builders+classicCalls)
	var wg sync.WaitGroup
	for i := range builders {
		wg.Go(func() {
			options, err := ClientOptions(w.Config, clientOptionsTestCredentials, &clientOptionsPlainTransport{})
			if err == nil {
				var c *modelregistry.Client
				if c, err = modelregistry.NewClient(ctx, options...); err == nil {
					_, err = clientOptionsGetModel(ctx, c)
				}
			}
			errs <- err
		})
		if i < classicCalls {
			wg.Go(func() {
				_, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "m"})
				errs <- err
			})
		}
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		assert.NoError(t, err)
	}
	requests := server.received()
	require.Len(t, requests, builders+classicCalls)
	for _, r := range requests {
		assert.Equal(t, []string{"42"}, r.Header.Values(clientOptionsWorkspaceIDHeader))
	}
}
