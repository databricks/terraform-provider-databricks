package sdkgoutility

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go"
	classicapierr "github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth/credentials"
	"github.com/databricks/sdk-go/core/apierr"
	"github.com/databricks/sdk-go/core/ops"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const createModelPath = "/api/2.0/mlflow/registered-models/create"

// toClassic is the converter that the tests pass to Retrier as classic. It
// gives the response of a modular SDK API error to the error parser of the
// classic SDK and returns any other error as it is.
func toClassic(err error) error {
	modular, ok := errors.AsType[*apierr.APIError](err)
	if !ok {
		return err
	}
	return classicapierr.GetAPIError(context.Background(), common.ResponseWrapper{
		ReadCloser: io.NopCloser(bytes.NewReader(modular.HTTPBody())),
		Response: &http.Response{
			StatusCode: modular.HTTPStatusCode(),
			Header:     modular.HTTPHeader(),
			Request:    httptest.NewRequest(http.MethodPost, createModelPath, nil),
		},
	})
}

// step is how one attempt of a call ends: with a response, with the network
// error err, or with a connection that the host refuses.
type step struct {
	status int
	header http.Header
	body   string
	err    error
	refuse bool
}

var created = step{status: http.StatusOK, body: `{"registered_model": {"name": "model"}}`}

func failure(status int, message string) step {
	return step{status: status, body: fmt.Sprintf(`{"error_code": "FIXTURE", "message": %q}`, message)}
}

func networkError(message string) step {
	return step{err: errors.New(message)}
}

// scriptedAPI is the HTTP transport of the classic and of the modular client
// in a test. It ends the attempts of a call as script says, repeating the last
// step, and records them. With a server it sends every attempt over the
// network to that server, which answers with the step. Without one it answers
// itself, which a client in a synctest bubble needs.
type scriptedAPI struct {
	script   []step
	server   *httptest.Server
	refusing string // the address of a port that nothing listens on

	mu       sync.Mutex
	current  step
	attempts []string // the method and path of every attempt
	times    []time.Time
}

func (a *scriptedAPI) RoundTrip(r *http.Request) (*http.Response, error) {
	a.mu.Lock()
	s := a.script[min(len(a.attempts), len(a.script)-1)]
	a.current = s
	a.attempts = append(a.attempts, r.Method+" "+r.URL.Path)
	a.times = append(a.times, time.Now())
	a.mu.Unlock()
	switch {
	case s.err != nil:
		return nil, s.err
	case s.refuse:
		r = r.Clone(r.Context())
		r.URL.Host = a.refusing
		return http.DefaultTransport.RoundTrip(r)
	case a.server != nil:
		return http.DefaultTransport.RoundTrip(r)
	}
	recorder := httptest.NewRecorder()
	a.ServeHTTP(recorder, r)
	response := recorder.Result()
	response.Request = r
	return response, nil
}

func (a *scriptedAPI) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	a.mu.Lock()
	s := a.current
	a.mu.Unlock()
	maps.Copy(w.Header(), s.header)
	w.WriteHeader(s.status)
	_, _ = io.WriteString(w, s.body)
}

// fixtureAPI is a scriptedAPI that the classic client takes for one of the
// fixture transports of the classic SDK.
type fixtureAPI struct{ *scriptedAPI }

func (fixtureAPI) SkipRetryOnIO() bool { return true }

// retryLog records what a call logs about its retries, without the jittered
// duration of a wait.
type retryLog struct {
	logger.SimpleLogger
	mu    sync.Mutex
	lines []string
}

func (l *retryLog) Tracef(_ context.Context, format string, v ...any) { l.record(format, v...) }
func (l *retryLog) Debugf(_ context.Context, format string, v ...any) { l.record(format, v...) }

func (l *retryLog) record(format string, v ...any) {
	line := fmt.Sprintf(format, v...)
	if message, _, isWait := strings.Cut(line, ". Sleeping "); isWait {
		line = message + ". Sleeping"
	} else if !strings.HasPrefix(line, "Attempting retry") && !strings.HasPrefix(line, "non-retriable error") {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, line)
}

// testConfig returns a classic configuration that resolves without the
// network.
func testConfig(host string, retryTimeoutSeconds int) *config.Config {
	return &config.Config{
		Host:                host,
		Token:               "token",
		RetryTimeoutSeconds: retryTimeoutSeconds,
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
			return nil, nil
		},
	}
}

// createModel is a call that creates the model "model" with one of the two
// clients. It returns the name of the created model.
type createModel func(ctx context.Context, t *testing.T, cfg *config.Config) (string, error)

func classicCall(ctx context.Context, t *testing.T, cfg *config.Config) (string, error) {
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	response, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: "model"})
	if err != nil {
		return "", err
	}
	return response.RegisteredModel.Name, nil
}

// modularCall makes the call with the retrier of Retrier. It turns off the
// timeout of the modular client, as Retrier leaves it to its caller to do.
func modularCall(ctx context.Context, t *testing.T, cfg *config.Config) (string, error) {
	token, err := credentials.NewPATCredentials(cfg.Token)
	require.NoError(t, err)
	c, err := modelregistry.NewClient(ctx, client.WithHost(cfg.Host), client.WithCredentials(token),
		client.WithHTTPClient(&http.Client{Transport: cfg.HTTPTransport}),
		client.WithoutEnv(), client.WithoutConfigFile(), client.WithTimeout(0))
	require.NoError(t, err)
	retrier, err := Retrier(ctx, cfg, toClassic)
	require.NoError(t, err)
	response, err := c.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: new("model")},
		call.WithRetrier(func() ops.Retrier { return retrier }))
	if err != nil {
		return "", toClassic(err)
	}
	return *response.RegisteredModel.Name, nil
}

// scenario is a call whose attempts end as script says.
type scenario struct {
	name   string
	script []step
	// retryTimeoutSeconds is retry_timeout_seconds of the classic config.
	retryTimeoutSeconds int
	// fixture gives the classic config a fixture transport.
	fixture bool
	// timeout is the timeout of the context of the call, if it isn't zero.
	timeout time.Duration
	// cancelAfter is when the context of the call is cancelled, if it isn't
	// zero.
	cancelAfter time.Duration

	// attempts is the number of attempts that both clients make. Zero stands
	// for a number that the jitter of the backoff decides.
	attempts int
	// outcome is how the call of both clients ends, as result.outcome says it.
	outcome string
}

// port matches the port of a local address, which differs from call to call.
var port = regexp.MustCompile(`(127\.0\.0\.1):\d+`)

// result is what a call did.
type result struct {
	attempts []string
	waits    []time.Duration // the time between two attempts
	logs     []string
	outcome  string
	took     time.Duration
}

// run makes the call of tt with one of the clients, against an httptest.Server
// if overNetwork is set and against an in-memory API otherwise.
func (tt scenario) run(t *testing.T, overNetwork bool, create createModel) result {
	api := &scriptedAPI{script: tt.script}
	cfg := testConfig("https://retrier.test", tt.retryTimeoutSeconds)
	if overNetwork {
		api.server = httptest.NewServer(api)
		t.Cleanup(api.server.Close)
		cfg.Host = api.server.URL
		closed, err := net.Listen("tcp", "127.0.0.1:0")
		require.NoError(t, err)
		api.refusing = closed.Addr().String()
		require.NoError(t, closed.Close())
	}
	cfg.HTTPTransport = api
	if tt.fixture {
		cfg.HTTPTransport = fixtureAPI{api}
	}
	log := &retryLog{SimpleLogger: logger.SimpleLogger{Level: logger.LevelError + 1}}
	ctx, cancel := context.WithCancel(logger.NewContext(t.Context(), log))
	defer cancel()
	if tt.timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, tt.timeout)
		defer cancel()
	}
	if tt.cancelAfter > 0 {
		defer time.AfterFunc(tt.cancelAfter, cancel).Stop()
	}
	start := time.Now()
	model, err := create(ctx, t, cfg)
	r := result{attempts: api.attempts, outcome: "created " + model, took: time.Since(start)}
	if apiError, ok := errors.AsType[*classicapierr.APIError](err); ok {
		r.outcome = fmt.Sprintf("%d %s: %s", apiError.StatusCode, apiError.ErrorCode, apiError.Message)
	} else if err != nil {
		r.outcome = port.ReplaceAllString(err.Error(), "$1")
	}
	for _, line := range log.lines {
		r.logs = append(r.logs, port.ReplaceAllString(line, "$1"))
	}
	for i := 1; i < len(api.times); i++ {
		r.waits = append(r.waits, api.times[i].Sub(api.times[i-1]))
	}
	return r
}

// compare makes the call of tt with the classic and with the modular client
// and checks that both make the same attempts and end the same way.
func (tt scenario) compare(t *testing.T, overNetwork bool) (classic, modular result) {
	classic, modular = tt.run(t, overNetwork, classicCall), tt.run(t, overNetwork, modularCall)
	if tt.attempts > 0 {
		assert.Len(t, classic.attempts, tt.attempts)
		assert.Equal(t, classic.attempts, modular.attempts)
		assert.Equal(t, classic.logs, modular.logs)
	}
	assert.Equal(t, tt.outcome, classic.outcome)
	assert.Equal(t, tt.outcome, modular.outcome)
	return classic, modular
}

// The classic backoff before retry n is n seconds, 10 at most, and a jitter
// between these bounds.
const (
	minJitter  = 50 * time.Millisecond
	maxJitter  = 750 * time.Millisecond
	maxBackoff = 10*time.Second + maxJitter
)

func TestRetrierMatchesClassicClientOverNetwork(t *testing.T) {
	const retryAfter = 30 * time.Second
	throttled := failure(http.StatusTooManyRequests, "slow down")
	throttled.header = http.Header{"Retry-After": {fmt.Sprint(retryAfter.Seconds())}}
	tests := []scenario{
		{name: "429 then 200", script: []step{failure(429, "slow down"), created}, attempts: 2, outcome: "created model"},
		{name: "503 then 200", script: []step{failure(503, "unavailable"), created}, attempts: 2, outcome: "created model"},
		{name: "500 is not retried", script: []step{failure(500, "broken"), created}, attempts: 1, outcome: "500 FIXTURE: broken"},
		{name: "400 is not retried", script: []step{failure(400, "invalid"), created}, attempts: 1, outcome: "400 FIXTURE: invalid"},
		{name: "403 is not retried", script: []step{failure(403, "denied"), created}, attempts: 1, outcome: "403 FIXTURE: denied"},
		{name: "404 is not retried", script: []step{failure(404, "missing"), created}, attempts: 1, outcome: "404 FIXTURE: missing"},
		{name: "transient message then 200", script: []step{failure(500, "Timed out after 30s"), created}, attempts: 2, outcome: "created model"},
		{name: "REQUEST_LIMIT_EXCEEDED then 200", script: []step{failure(500, "REQUEST_LIMIT_EXCEEDED: too fast"), created}, attempts: 2, outcome: "created model"},
		{name: "refused connection then 200", script: []step{{refuse: true}, created}, attempts: 2, outcome: "created model"},
		{name: "429 with Retry-After then 200", script: []step{throttled, created}, attempts: 2, outcome: "created model"},
		{name: "429 until a retry timeout of 2 seconds", script: []step{failure(429, "slow down")}, retryTimeoutSeconds: 2, attempts: 2, outcome: "429 FIXTURE: slow down"},
		{name: "503 until a retry timeout of 1 second", script: []step{failure(503, "unavailable")}, retryTimeoutSeconds: 1, attempts: 1, outcome: "503 FIXTURE: unavailable"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			classic, modular := tt.compare(t, true)
			// Neither client waits as long as the Retry-After header asks.
			assert.Less(t, max(classic.took, modular.took), retryAfter)
		})
	}
}

func TestRetrierMatchesClassicClientInFakeTime(t *testing.T) {
	throttled := failure(http.StatusTooManyRequests, "slow down")
	throttled.header = http.Header{"Retry-After": {"120"}}
	html := step{status: http.StatusServiceUnavailable, body: "<html><pre>No webapps are available</pre></html>"}
	tests := []scenario{
		{name: "429 with Retry-After", script: []step{throttled, created}, attempts: 2},
		{name: "429 without a body", script: []step{{status: 429}, created}, attempts: 2},
		{name: "503 with an HTML body", script: []step{html, created}, attempts: 2},
		{name: "error code of a 429 on a 500", script: []step{{status: 500, body: `{"error_code": "RESOURCE_EXHAUSTED"}`}, created}, attempts: 2},
		{name: "unknown worker environment exception", script: []step{failure(400, "com.databricks.backend.manager.util.UnknownWorkerEnvironmentException"), created}, attempts: 2},
		{name: "no associated worker environments", script: []step{failure(400, "Workspace 1 does not have any associated worker environments"), created}, attempts: 2},
		{name: "no worker environment with id", script: []step{failure(400, "There is no worker environment with id 1"), created}, attempts: 2},
		{name: "unknown worker environment", script: []step{failure(400, "Unknown worker environment 1"), created}, attempts: 2},
		{name: "cluster not ready", script: []step{failure(400, "ClusterNotReadyException"), created}, attempts: 2},
		{name: "worker env not found", script: []step{failure(404, "worker env 1 not found"), created}, attempts: 2},
		{name: "timed out", script: []step{failure(500, "Timed out after 30s"), created}, attempts: 2},
		{name: "deadline exceeded", script: []step{failure(500, "deadline exceeded"), created}, attempts: 2},
		{name: "REQUEST_LIMIT_EXCEEDED", script: []step{failure(500, "REQUEST_LIMIT_EXCEEDED"), created}, attempts: 2},
		{name: "connection reset by peer", script: []step{networkError("read: connection reset by peer"), created}, attempts: 2},
		{name: "TLS handshake timeout", script: []step{networkError("net/http: TLS handshake timeout"), created}, attempts: 2},
		{name: "connection refused", script: []step{networkError("connect: connection refused"), created}, attempts: 2},
		{name: "unexpected error", script: []step{networkError("Unexpected error"), created}, attempts: 2},
		{name: "i/o timeout", script: []step{networkError("dial tcp: i/o timeout"), created}, attempts: 2},
		{name: "REQUEST_LIMIT_EXCEEDED in a network error", script: []step{networkError("REQUEST_LIMIT_EXCEEDED"), created}, attempts: 2},
		{name: "every kind of retriable error in turn", attempts: 5, script: []step{
			failure(429, "slow down"), networkError("i/o timeout"), failure(503, "unavailable"), failure(500, "REQUEST_LIMIT_EXCEEDED"), created,
		}},
		{name: "retries end at the first other error", attempts: 3, outcome: "409 FIXTURE: conflict", script: []step{
			failure(429, "slow down"), failure(503, "unavailable"), failure(409, "conflict"), created,
		}},
		{name: "13 waits", attempts: 14, script: append(slicesRepeat(failure(429, "slow down"), 13), created)},
		{name: "500 is not retried", script: []step{failure(500, "broken"), created}, attempts: 1, outcome: "500 FIXTURE: broken"},
		{name: "500 without a body is not retried", script: []step{{status: 500}, created}, attempts: 1, outcome: "500 : Internal Server Error"},
		{name: "401 is not retried", script: []step{failure(401, "who"), created}, attempts: 1, outcome: "401 FIXTURE: who"},
		{name: "504 is not retried", script: []step{failure(504, "late"), created}, attempts: 1, outcome: "504 FIXTURE: late"},
		{name: "EOF is not retried", script: []step{networkError("EOF"), created}, attempts: 1, outcome: `Post "https://retrier.test` + createModelPath + `": EOF`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				if tt.outcome == "" {
					tt.outcome = "created model"
				}
				classic, modular := tt.compare(t, false)
				require.Len(t, modular.waits, len(classic.waits))
				for i, wait := range modular.waits {
					// The clock of the bubble stands still during an attempt, so
					// the time between two attempts is the wait alone.
					seconds := time.Duration(min(i+1, 10)) * time.Second
					assert.Equal(t, seconds, classic.waits[i].Truncate(time.Second), "classic wait %d", i+1)
					assert.Equal(t, seconds, wait.Truncate(time.Second), "wait %d", i+1)
					assert.GreaterOrEqual(t, wait-seconds, minJitter, "wait %d", i+1)
					assert.LessOrEqual(t, wait-seconds, maxJitter, "wait %d", i+1)
				}
			})
		})
	}
}

func slicesRepeat(s step, count int) []step {
	steps := make([]step, count)
	for i := range steps {
		steps[i] = s
	}
	return steps
}

func TestRetrierStopsWhenClassicClientDoes(t *testing.T) {
	const throttledOutcome = "429 FIXTURE: slow down"
	tests := []struct {
		scenario
		// took is how long the classic call takes.
		took time.Duration
		// modularOutcome is how the modular call ends, where it differs.
		modularOutcome string
	}{
		{scenario: scenario{name: "unset retry timeout is 5 minutes"}, took: 5 * time.Minute},
		{scenario: scenario{name: "retry timeout of 1 second", retryTimeoutSeconds: 1, attempts: 1}, took: time.Second},
		{scenario: scenario{name: "retry timeout of 20 seconds", retryTimeoutSeconds: 20, attempts: 6}, took: 20 * time.Second},
		{scenario: scenario{name: "fixture transport has 20 minutes", retryTimeoutSeconds: 20, fixture: true}, took: 20 * time.Minute},
		{scenario: scenario{name: "deadline of the context", retryTimeoutSeconds: -1, timeout: 20 * time.Second, attempts: 6}, took: 20 * time.Second},
		{scenario: scenario{name: "deadline before the retry timeout", retryTimeoutSeconds: 60, timeout: 20 * time.Second, attempts: 6}, took: 20 * time.Second},
		{scenario: scenario{name: "retry timeout before the deadline", retryTimeoutSeconds: 20, timeout: time.Minute, attempts: 6}, took: 20 * time.Second},
		{scenario: scenario{name: "negative retry timeout retries until the context ends", retryTimeoutSeconds: -1, cancelAfter: time.Hour},
			took: time.Hour, modularOutcome: "context canceled"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				tt.script = []step{failure(429, "slow down")}
				classic, modular := tt.run(t, false, classicCall), tt.run(t, false, modularCall)
				assert.Equal(t, throttledOutcome, classic.outcome)
				assert.Equal(t, tt.took, classic.took)
				if tt.attempts > 0 {
					assert.Len(t, classic.attempts, tt.attempts)
					assert.Equal(t, classic.attempts, modular.attempts)
				}
				if tt.modularOutcome != "" {
					// The modular SDK returns the error of the context when the
					// context ends during a wait, and the retrier can't change it.
					assert.Equal(t, tt.modularOutcome, modular.outcome)
					assert.Equal(t, classic.took, modular.took)
					return
				}
				// The modular call ends when the wait that the classic call
				// doesn't finish would start, with the same error.
				assert.Equal(t, throttledOutcome, modular.outcome)
				assert.LessOrEqual(t, modular.took, classic.took)
				assert.Less(t, classic.took-modular.took, maxBackoff)
			})
		})
	}
}

func TestRetrierWaitsByClassicBackoff(t *testing.T) {
	retrier, err := Retrier(t.Context(), testConfig("https://retrier.test", -1), toClassic)
	require.NoError(t, err)
	throttled := apierr.FromHTTPError(http.StatusTooManyRequests, http.Header{"Retry-After": {"120"}}, nil)
	for n := 1; n <= 13; n++ {
		wait, ok := retrier.IsRetriable(throttled)
		require.True(t, ok, "retry %d", n)
		classic := retries.RetryConfig{}.Backoff(n)
		assert.Equal(t, classic.Truncate(time.Second), wait.Truncate(time.Second), "retry %d", n)
		assert.GreaterOrEqual(t, wait-wait.Truncate(time.Second), minJitter, "retry %d", n)
		assert.LessOrEqual(t, wait-wait.Truncate(time.Second), maxJitter, "retry %d", n)
	}
}

func TestRetrierDoesNotRetry(t *testing.T) {
	throttled := apierr.FromHTTPError(http.StatusTooManyRequests, nil, nil)
	ended, cancel := context.WithCancel(t.Context())
	cancel()
	tests := []struct {
		name    string
		ctx     context.Context
		classic func(error) error
		err     error
	}{
		{name: "an error that classic turns into nil", ctx: t.Context(), classic: func(error) error { return nil }, err: throttled},
		{name: "a retriable error after the context has ended", ctx: ended, classic: toClassic, err: throttled},
		{name: "a network message outside a url.Error", ctx: t.Context(), classic: toClassic, err: errors.New("connection refused")},
		{name: "a url.Error with another message", ctx: t.Context(), classic: toClassic, err: &url.Error{Op: "Post", Err: io.EOF}},
		{name: "an API error that the classic client doesn't retry", ctx: t.Context(), classic: toClassic,
			err: apierr.FromHTTPError(http.StatusInternalServerError, nil, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retrier, err := Retrier(tt.ctx, testConfig("https://retrier.test", 0), tt.classic)
			require.NoError(t, err)
			wait, ok := retrier.IsRetriable(tt.err)
			assert.False(t, ok)
			assert.Zero(t, wait)
		})
	}
}

func TestRetrierConvertsEveryErrorWithClassic(t *testing.T) {
	var converted []error
	retrier, err := Retrier(t.Context(), testConfig("https://retrier.test", 0), func(err error) error {
		converted = append(converted, err)
		return &classicapierr.APIError{StatusCode: http.StatusServiceUnavailable, Message: "unavailable"}
	})
	require.NoError(t, err)
	plain := errors.New("an error that only classic understands")
	_, ok := retrier.IsRetriable(plain)
	assert.True(t, ok, "the retrier decides by the error that classic returns")
	assert.Equal(t, []error{plain}, converted)
}

func TestRetrierErrors(t *testing.T) {
	unresolved := func() *config.Config {
		cfg := testConfig("https://retrier.test", 0)
		cfg.Username, cfg.Password = "user", "password"
		return cfg
	}
	_, classicErr := databricks.NewWorkspaceClient((*databricks.Config)(unresolved()))
	require.ErrorContains(t, classicErr, "more than one authorization method configured")
	tests := []struct {
		name    string
		cfg     *config.Config
		classic func(error) error
		err     string
	}{
		{name: "no config", classic: toClassic, err: "sdkgoutility.Retrier: cfg is nil"},
		{name: "no converter", cfg: testConfig("https://retrier.test", 0), err: "sdkgoutility.Retrier: classic is nil"},
		{name: "config that doesn't resolve", cfg: unresolved(), classic: toClassic, err: classicErr.Error()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retrier, err := Retrier(t.Context(), tt.cfg, tt.classic)
			assert.EqualError(t, err, tt.err)
			assert.Nil(t, retrier)
		})
	}
}
