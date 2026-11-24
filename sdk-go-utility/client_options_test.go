package sdkgoutility

import (
	"context"
	"io"
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

// clientOptionsMapCredentials are credentials whose zero value is a nil map.
type clientOptionsMapCredentials map[string]string

func (c clientOptionsMapCredentials) Name() string { return "map" }

func (c clientOptionsMapCredentials) AuthHeaders(context.Context) ([]auth.Header, error) {
	return nil, nil
}

// clientOptionsRoundTripFunc is a transport whose zero value is a nil func.
type clientOptionsRoundTripFunc func(*http.Request) (*http.Response, error)

func (f clientOptionsRoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// clientOptionsPlainTransport notes the deadline of every request, marks it
// with clientOptionsTransportHeader and sends it.
type clientOptionsPlainTransport struct {
	mu   sync.Mutex
	sent []clientOptionsSentRequest
}

type clientOptionsSentRequest struct {
	deadline time.Time
	bounded  bool
}

func (t *clientOptionsPlainTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	deadline, bounded := r.Context().Deadline()
	t.mu.Lock()
	t.sent = append(t.sent, clientOptionsSentRequest{deadline, bounded})
	t.mu.Unlock()
	marked := r.Clone(r.Context())
	marked.Header.Set(clientOptionsTransportHeader, "plain")
	return http.DefaultTransport.RoundTrip(marked)
}

// clientOptionsServer records the requests that it gets and answers them all
// the same, after delay.
type clientOptionsServer struct {
	*httptest.Server
	mu       sync.Mutex
	requests []*http.Request
	delay    time.Duration
}

func clientOptionsNewServer(t *testing.T, status int, body string) *clientOptionsServer {
	t.Helper()
	s := &clientOptionsServer{}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		s.requests = append(s.requests, r)
		delay := s.delay
		s.mu.Unlock()
		time.Sleep(delay)
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

// clientOptionsClassicClient returns a real classic client for cfg. Its Config
// is resolved, as the one that the provider passes to ClientOptions.
func clientOptionsClassicClient(t *testing.T, cfg *config.Config) *databricks.WorkspaceClient {
	t.Helper()
	cfg.Token = "classic-token"
	cfg.HostMetadataResolver = func(context.Context, string) (*config.HostMetadata, error) {
		return nil, nil
	}
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
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
	model, err := c.GetRegisteredModelDatabricks(ctx,
		modelregistry.GetRegisteredModelDatabricksRequest{Name: &name}, options...)
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
		{name: "nil map as credentials", cfg: resolved, credentials: clientOptionsMapCredentials(nil), transport: transport, want: nilCredentials},
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
	transport := &clientOptionsPlainTransport{}
	tests := []struct {
		name        string
		cfg         *config.Config
		credentials auth.Credentials // clientOptionsTestCredentials when nil
		wantTimeout time.Duration
	}{
		{name: "host only", cfg: &config.Config{Host: clientOptionsTestHost}, wantTimeout: 5 * time.Minute},
		{name: "workspace ID", cfg: &config.Config{Host: clientOptionsTestHost, WorkspaceID: "1234567890123456"}, wantTimeout: 5 * time.Minute},
		{name: "account ID", cfg: &config.Config{Host: "https://accounts.cloud.databricks.com", AccountID: "8f4c2a10-account"}, wantTimeout: 5 * time.Minute},
		{name: "workspace ID and account ID", cfg: &config.Config{Host: clientOptionsTestHost, WorkspaceID: "conn-1a2b", AccountID: "8f4c2a10-account"}, wantTimeout: 5 * time.Minute},
		{name: "host with an IPv6 address and a port", cfg: &config.Config{Host: "http://[::1]:8080"}, wantTimeout: 5 * time.Minute},
		{name: "credentials that are a pointer", cfg: &config.Config{Host: clientOptionsTestHost}, credentials: &clientOptionsTestCredentials, wantTimeout: 5 * time.Minute},
		{name: "credentials that are a map", cfg: &config.Config{Host: clientOptionsTestHost}, credentials: clientOptionsMapCredentials{}, wantTimeout: 5 * time.Minute},
		{name: "retry timeout of a second", cfg: &config.Config{Host: clientOptionsTestHost, RetryTimeoutSeconds: 1}, wantTimeout: time.Second},
		{name: "retry timeout of an hour", cfg: &config.Config{Host: clientOptionsTestHost, RetryTimeoutSeconds: 3600}, wantTimeout: time.Hour},
		{name: "negative retry timeout turns the timeout off", cfg: &config.Config{Host: clientOptionsTestHost, RetryTimeoutSeconds: -1}, wantTimeout: 0},
		{name: "HTTP timeout and rate limit are not carried", cfg: &config.Config{Host: clientOptionsTestHost, HTTPTimeoutSeconds: 65, RateLimitPerSecond: 3}, wantTimeout: 5 * time.Minute},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.credentials == nil {
				tt.credentials = clientOptionsTestCredentials
			}
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
			require.NotNil(t, got.Timeout)
			assert.Equal(t, tt.wantTimeout, *got.Timeout)
			// EnsureResolved fills the loaders of a config that has none.
			assert.Empty(t, tt.cfg.Loaders, "ClientOptions resolved the config")
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
			server := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody)
			w := clientOptionsClassicClient(t, &config.Config{
				Host:        server.URL + tt.host,
				WorkspaceID: tt.workspaceID,
				AccountID:   tt.accountID,
			})
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
			assert.Len(t, transport.sent, 1)
		})
	}
}

func TestClientOptions_SendsOnlyOverTransport(t *testing.T) {
	clientOptionsCleanEnvironment(t)
	var urls []string
	// The host doesn't resolve, so a request over any other transport fails.
	transport := clientOptionsRoundTripFunc(func(r *http.Request) (*http.Response, error) {
		urls = append(urls, r.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(clientOptionsModelBody)),
			Request:    r,
		}, nil
	})
	c := clientOptionsModularClient(t, &config.Config{Host: "https://workspace.invalid"}, transport)

	name, err := clientOptionsGetModel(context.Background(), c)

	require.NoError(t, err)
	assert.Equal(t, "m", name)
	assert.Equal(t, []string{"https://workspace.invalid/api/2.0/mlflow/databricks/registered-models/get?name=m"}, urls)
}

func TestClientOptions_IgnoresEnvironment(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name string
		// env are the variables to set and files the files to write, by path
		// in the home directory. In both, {elsewhere} is the URL of a server
		// that is not the workspace and {home} is the home directory.
		env   map[string]string
		files map[string]string
		// What a modular client without the options does in this environment:
		// NewClient fails with leftAloneError, or the call goes elsewhere with
		// this Authorization and this workspace ID.
		leftAloneError         string
		leftAloneAuthorization string
		leftAloneWorkspaceID   []string
	}{
		{name: "DATABRICKS_HOST and DATABRICKS_TOKEN",
			env:                    map[string]string{"DATABRICKS_HOST": "{elsewhere}", "DATABRICKS_TOKEN": "env-token"},
			leftAloneAuthorization: "Bearer env-token"},
		{name: "DATABRICKS_WORKSPACE_ID and DATABRICKS_ACCOUNT_ID",
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
			workspace := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody)
			elsewhere := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody)
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

func TestClientOptions_CallTimeout(t *testing.T) {
	clientOptionsCleanEnvironment(t)
	server := clientOptionsNewServer(t, http.StatusOK, clientOptionsModelBody)
	tests := []struct {
		name                string
		retryTimeoutSeconds int
		// callerTimeout is the timeout of the context of the caller, if any.
		callerTimeout time.Duration
		// want is how far ahead the deadline of the request is, or zero when
		// the request has no deadline.
		want time.Duration
	}{
		{name: "unset is the 5 minutes of the classic client", retryTimeoutSeconds: 0, want: 5 * time.Minute},
		{name: "1 second", retryTimeoutSeconds: 1, want: time.Second},
		{name: "45 seconds", retryTimeoutSeconds: 45, want: 45 * time.Second},
		{name: "1 hour", retryTimeoutSeconds: 3600, want: time.Hour},
		{name: "minus one is no bound, as the provider sets it", retryTimeoutSeconds: -1},
		{name: "any negative value is no bound", retryTimeoutSeconds: -300},
		{name: "earlier deadline of the caller wins", retryTimeoutSeconds: 60, callerTimeout: 10 * time.Second, want: 10 * time.Second},
		{name: "later deadline of the caller loses", retryTimeoutSeconds: 60, callerTimeout: time.Hour, want: time.Minute},
		{name: "deadline of the caller is the only one without a bound", retryTimeoutSeconds: -1, callerTimeout: 20 * time.Minute, want: 20 * time.Minute},
		{name: "deadline of the caller after the 5 minutes loses", retryTimeoutSeconds: 0, callerTimeout: 20 * time.Minute, want: 5 * time.Minute},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport := &clientOptionsPlainTransport{}
			c := clientOptionsModularClient(t, &config.Config{Host: server.URL, RetryTimeoutSeconds: tt.retryTimeoutSeconds}, transport)
			before := time.Now()
			ctx := context.Background()
			if tt.callerTimeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tt.callerTimeout)
				defer cancel()
			}

			_, err := clientOptionsGetModel(ctx, c)
			after := time.Now()

			require.NoError(t, err)
			require.Len(t, transport.sent, 1)
			sent := transport.sent[0]
			if tt.want == 0 {
				assert.False(t, sent.bounded, "the request has the deadline %s", sent.deadline)
				return
			}
			require.True(t, sent.bounded, "the request has no deadline")
			assert.False(t, sent.deadline.Before(before.Add(tt.want)), "the deadline is less than %s ahead", tt.want)
			assert.False(t, sent.deadline.After(after.Add(tt.want)), "the deadline is more than %s ahead", tt.want)
		})
	}
}

// Both clients give a call the time of cfg.RetryTimeoutSeconds, here a second.
// The classic client looks at it between attempts and returns what the last
// attempt returned. The modular client puts it on the context of the call. No
// client option changes that, so these cases also pin where the two differ.
func TestClientOptions_BoundsCallLikeClassicClient(t *testing.T) {
	throttled := `{"error_code":"REQUEST_LIMIT_EXCEEDED","message":"slow down"}`
	tests := []struct {
		name           string
		status         int
		body           string
		delay          time.Duration // of every answer of the server
		wantClassicErr error         // nil when the classic call succeeds
		wantRetries    bool
	}{
		{name: "every attempt is throttled", status: http.StatusTooManyRequests, body: throttled,
			wantClassicErr: apierr.ErrTooManyRequests, wantRetries: true},
		{name: "one attempt outlasts the bound", status: http.StatusOK, body: clientOptionsModelBody, delay: 1300 * time.Millisecond},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientOptionsCleanEnvironment(t)
			ctx := context.Background()
			server := clientOptionsNewServer(t, tt.status, tt.body)
			server.mu.Lock()
			server.delay = tt.delay
			server.mu.Unlock()
			w := clientOptionsClassicClient(t, &config.Config{Host: server.URL, RetryTimeoutSeconds: 1})
			c := clientOptionsModularClient(t, w.Config, &clientOptionsPlainTransport{})
			// A modular client retries only with a retrier, a call option.
			retry := call.WithRetrier(func() ops.Retrier {
				backoff := ops.BackoffPolicy{Initial: 50 * time.Millisecond, Maximum: 50 * time.Millisecond}
				return ops.RetryIf(backoff, func(error) bool { return true })
			})

			var classicErr, modularErr error
			var classicTook, modularTook time.Duration
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				start := time.Now()
				_, classicErr = w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "m"})
				classicTook = time.Since(start)
			}()
			go func() {
				defer wg.Done()
				start := time.Now()
				_, modularErr = clientOptionsGetModel(ctx, c, retry)
				modularTook = time.Since(start)
			}()
			wg.Wait()

			assert.ErrorIs(t, classicErr, tt.wantClassicErr)
			assert.ErrorIs(t, modularErr, context.DeadlineExceeded)
			for name, took := range map[string]time.Duration{"classic": classicTook, "modular": modularTook} {
				assert.GreaterOrEqual(t, took, time.Second, "the %s client gave up early", name)
				assert.Less(t, took, 6*time.Second, "the %s client went on", name)
			}
			assert.Equal(t, tt.wantRetries, len(server.received()) > 2, "retries of the modular client")
		})
	}
}
