package sdkgoutility

import (
	"context"
	"errors"
	"fmt"
	"io"
	"maps"
	"net"
	"net/http"
	"net/http/httptest"
	"regexp"
	"slices"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth/credentials"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The tests make their calls over real connections to an httptest.Server, on
// the fake clock of a synctest bubble, so that no test waits out a backoff.
// The server runs outside the bubble. The clock of a bubble stands still while
// a connection is in use, so whatever takes time happens in the transport,
// before the request is sent.

// classicCallStep is how one attempt of a call ends: with a response, with a
// connection that is refused, that the server resets or that the server holds
// until the client gives up, or with no answer at all.
type classicCallStep struct {
	status                        int
	header                        http.Header
	body                          string
	refused, reset, holds, stalls bool
	// takes is how long the attempt takes to reach the server.
	takes time.Duration
}

var classicCallFound = classicCallStep{status: http.StatusOK,
	body: `{"registered_model_databricks": {"id": "model"}, "registered_model": {"name": "model"}}`}

func classicCallFails(status int, code, message string) classicCallStep {
	return classicCallStep{status: status, body: fmt.Sprintf(`{"error_code": %q, "message": %q}`, code, message)}
}

const (
	classicCallGetPath    = "/api/2.0/mlflow/databricks/registered-models/get"
	classicCallCreatePath = "/api/2.0/mlflow/registered-models/create"
	// classicCallMaxAttempts is more attempts than any call of the tests
	// makes. The transport fails the next one, so that a call that never stops
	// retrying fails its test instead of hanging it.
	classicCallMaxAttempts = 100
	// classicCallRefusingAddress refuses connections: nothing listens on a
	// port that only root can open.
	classicCallRefusingAddress = "127.0.0.1:1"
	// The classic backoff before retry n is n seconds and at most this jitter.
	classicCallMaxJitter = 750 * time.Millisecond
	// The moments at which the context of a call is cancelled.
	classicCallBeforeTheCall    = "before the call"
	classicCallDuringTheWait    = "during the wait to retry"
	classicCallDuringTheRequest = "during the request"
)

// classicCallAPI is the model registry of a test: an httptest.Server, and the
// HTTP transport through which the classic and the modular client reach it.
// It ends the attempts of a call as script says, repeating the last step.
type classicCallAPI struct {
	server *httptest.Server
	script []classicCallStep
	// held gets a value when the server holds a request.
	held chan struct{}

	mu sync.Mutex
	// attempts are the times at which the client started an attempt, and
	// requests are the requests that the server received.
	attempts []time.Time
	requests []string
}

func classicCallNewAPI(t *testing.T, script ...classicCallStep) *classicCallAPI {
	api := &classicCallAPI{script: script, held: make(chan struct{}, 1)}
	api.server = httptest.NewServer(api)
	t.Cleanup(api.server.Close)
	return api
}

// RoundTrip sends an attempt to the server over a connection of its own, so
// that no connection outlives its attempt.
func (a *classicCallAPI) RoundTrip(r *http.Request) (*http.Response, error) {
	a.mu.Lock()
	step := a.script[min(len(a.attempts), len(a.script)-1)]
	a.attempts = append(a.attempts, time.Now())
	attempts := len(a.attempts)
	a.mu.Unlock()
	if attempts > classicCallMaxAttempts {
		return nil, errors.New("the call doesn't stop retrying")
	}
	time.Sleep(step.takes)
	if step.stalls {
		<-r.Context().Done()
		return nil, r.Context().Err()
	}
	r = r.Clone(r.Context())
	if step.refused {
		r.URL.Host = classicCallRefusingAddress
	}
	if r.Body != nil && r.Body != http.NoBody {
		// The classic client rewinds the body of a request to retry it, which
		// it may only do once net/http has closed that body, so wait for that.
		closed := make(chan struct{})
		r.Body = classicCallBody{r.Body, sync.OnceFunc(func() { close(closed) })}
		defer func() { <-closed }()
	}
	return (&http.Transport{DisableKeepAlives: true}).RoundTrip(r)
}

// classicCallBody is the body of a request that tells when it is closed.
type classicCallBody struct {
	io.ReadCloser
	closed func()
}

func (b classicCallBody) Close() error {
	defer b.closed()
	return b.ReadCloser.Close()
}

func (a *classicCallAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// The server notices that a client is gone once it has read the request.
	_, _ = io.Copy(io.Discard, r.Body)
	a.mu.Lock()
	step := a.script[min(len(a.attempts), len(a.script))-1]
	a.requests = append(a.requests, r.Method+" "+r.URL.RequestURI())
	a.mu.Unlock()
	switch {
	case step.reset:
		conn, _, _ := w.(http.Hijacker).Hijack()
		// Closing without lingering sends a reset instead of an orderly end.
		_ = conn.(*net.TCPConn).SetLinger(0)
		_ = conn.Close()
	case step.holds:
		a.held <- struct{}{}
		<-r.Context().Done()
		// Returning would answer a client that may still be reading.
		panic(http.ErrAbortHandler)
	default:
		maps.Copy(w.Header(), step.header)
		w.WriteHeader(step.status)
		_, _ = io.WriteString(w, step.body)
	}
}

// restart returns the attempts and requests of the last call and forgets them.
func (a *classicCallAPI) restart() (attempts []time.Time, requests []string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	attempts, requests = a.attempts, a.requests
	a.attempts, a.requests = nil, nil
	return attempts, requests
}

// config returns a classic config for the workspace that a serves, with the
// fields that set sets. It is resolved, as the config of the provider is when
// a resource makes a call: calls that run at once can't resolve their config.
func (a *classicCallAPI) config(t *testing.T, set func(*config.Config)) *config.Config {
	cfg := &config.Config{Host: a.server.URL, Token: "token", HTTPTransport: a,
		// The classic client would ask the server for its metadata first.
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) { return nil, nil }}
	set(cfg)
	require.NoError(t, cfg.EnsureResolved())
	return cfg
}

// client returns a modular client for the workspace that a serves, with the
// options that the ticket of ClassicCall names, because the helper that
// builds clients is another one.
func (a *classicCallAPI) client(t *testing.T) *modelregistry.Client {
	token, err := credentials.NewPATCredentials("token")
	require.NoError(t, err)
	c, err := modelregistry.NewClient(t.Context(), client.WithHost(a.server.URL), client.WithCredentials(token),
		client.WithHTTPClient(&http.Client{Transport: a}),
		client.WithoutConfigFile(), client.WithoutEnv(), client.WithTimeout(0))
	require.NoError(t, err)
	return c
}

// classicCallGet reads the model "model" with c, as a resource does.
func classicCallGet(ctx context.Context, cfg *config.Config, c *modelregistry.Client) (*modelregistry.GetRegisteredModelDatabricksResponse, error) {
	return ClassicCall(ctx, cfg, func(ctx context.Context, opts ...call.Option) (*modelregistry.GetRegisteredModelDatabricksResponse, error) {
		return c.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: new("model")}, opts...)
	})
}

// classicCallClient makes one call for the model "model" with the classic or
// the modular client of api, and returns what the call returned of the model.
type classicCallClient func(ctx context.Context, t *testing.T, cfg *config.Config, api *classicCallAPI) (string, error)

func classicCallClassicGet(ctx context.Context, t *testing.T, cfg *config.Config, _ *classicCallAPI) (string, error) {
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	response, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "model"})
	if err != nil {
		assert.Equal(t, &ml.GetModelResponse{}, response, "the classic client returns an empty response with an error")
		return "", err
	}
	return response.RegisteredModelDatabricks.Id, nil
}

func classicCallClassicCreate(ctx context.Context, t *testing.T, cfg *config.Config, _ *classicCallAPI) (string, error) {
	w, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	response, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: "model"})
	if err != nil {
		assert.Equal(t, &ml.CreateModelResponse{}, response, "the classic client returns an empty response with an error")
		return "", err
	}
	return response.RegisteredModel.Name, nil
}

func classicCallModularGet(ctx context.Context, t *testing.T, cfg *config.Config, api *classicCallAPI) (string, error) {
	response, err := classicCallGet(ctx, cfg, api.client(t))
	if err != nil {
		assert.Nil(t, response, "ClassicCall returns the zero value with an error")
		return "", err
	}
	return *response.RegisteredModelDatabricks.Id, nil
}

func classicCallModularCreate(ctx context.Context, t *testing.T, cfg *config.Config, api *classicCallAPI) (string, error) {
	c := api.client(t)
	response, err := ClassicCall(ctx, cfg, func(ctx context.Context, opts ...call.Option) (*modelregistry.CreateRegisteredModelResponse, error) {
		return c.CreateRegisteredModel(ctx, modelregistry.CreateRegisteredModelRequest{Name: new("model")}, opts...)
	})
	if err != nil {
		assert.Nil(t, response, "ClassicCall returns the zero value with an error")
		return "", err
	}
	return *response.RegisteredModel.Name, nil
}

// classicCallOps are the calls that every scenario is made with. request is
// what the server receives, and failed is how the message of a network error
// starts, with HOST for the URL of the server.
var classicCallOps = []struct {
	name, request, failed string
	classic, modular      classicCallClient
}{
	{"GET", "GET " + classicCallGetPath + "?name=model", `Get "HOST` + classicCallGetPath + `?name=model": `,
		classicCallClassicGet, classicCallModularGet},
	{"POST", "POST " + classicCallCreatePath, `Post "HOST` + classicCallCreatePath + `": `,
		classicCallClassicCreate, classicCallModularCreate},
}

// classicCallOutcome is what the server and the caller see of one call.
type classicCallOutcome struct {
	// attempts is the number of attempts of the client, and requests the
	// number of those that the server received.
	attempts, requests int
	// model is what the call returned of the model.
	model string
	// err is the message of the error, with HOST for the URL of the server and
	// without the ports of local addresses. errType is its type, for the
	// classic client the one inside the *retries.Err that it wraps every error
	// in. status and code are those of the classic *apierr.APIError in it.
	err, errType      string
	status            int
	code              string
	missing, canceled bool
}

// classicCallResult is the outcome of a call, how long the call took and
// whether its error came in a *retries.Err.
type classicCallResult struct {
	classicCallOutcome
	took    time.Duration
	wrapped bool
}

// classicCallPort matches the port of a local address, which differs from
// connection to connection.
var classicCallPort = regexp.MustCompile(`(127\.0\.0\.1):\d+`)

func classicCallNewResult(host, model string, err error) classicCallResult {
	result := classicCallResult{classicCallOutcome: classicCallOutcome{model: model}}
	if wrapper, ok := err.(*retries.Err); ok {
		result.wrapped, err = true, wrapper.Err
	}
	if err != nil {
		result.err = classicCallPort.ReplaceAllString(strings.ReplaceAll(err.Error(), host, "HOST"), "$1")
		result.errType = fmt.Sprintf("%T", err)
		result.missing, result.canceled = apierr.IsMissing(err), errors.Is(err, context.Canceled)
	}
	if apiErr, ok := errors.AsType[*apierr.APIError](err); ok {
		result.status, result.code = apiErr.StatusCode, apiErr.ErrorCode
	}
	return result
}

// classicCallScenario is a call whose attempts end as script says.
type classicCallScenario struct {
	name   string
	script []classicCallStep
	// retryTimeoutSeconds and httpTimeoutSeconds are retry_timeout_seconds
	// and http_timeout_seconds of the classic config.
	retryTimeoutSeconds, httpTimeoutSeconds int
	// deadline is the timeout of the context of the call, if it isn't zero.
	deadline time.Duration
	// cancelled is when the context of the call is cancelled, if it is.
	cancelled string
	// want is the outcome of both clients. Its err has REQUEST for the start
	// of the message of a network error.
	want classicCallOutcome
	// took is how long the classic call takes, without the jitter of its
	// waits. sooner marks a call whose retries run out: ClassicCall returns
	// when the last wait would start and the classic client when it ends.
	took   time.Duration
	sooner bool
}

// run makes the call of tt with one of the two clients.
func (tt classicCallScenario) run(t *testing.T, api *classicCallAPI, request string, call classicCallClient) classicCallResult {
	cfg := api.config(t, func(cfg *config.Config) {
		cfg.RetryTimeoutSeconds, cfg.HTTPTimeoutSeconds = tt.retryTimeoutSeconds, tt.httpTimeoutSeconds
	})
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	if tt.deadline > 0 {
		ctx, cancel = context.WithTimeout(ctx, tt.deadline)
		defer cancel()
	}
	switch tt.cancelled {
	case classicCallBeforeTheCall:
		cancel()
	case classicCallDuringTheWait:
		// The first wait to retry lasts at least a second.
		defer time.AfterFunc(500*time.Millisecond, cancel).Stop()
	case classicCallDuringTheRequest:
		go func() { <-api.held; cancel() }()
	}
	start := time.Now()
	model, err := call(ctx, t, cfg, api)
	result := classicCallNewResult(api.server.URL, model, err)
	result.took = time.Since(start)
	attempts, requests := api.restart()
	result.attempts, result.requests = len(attempts), len(requests)
	for _, received := range requests {
		assert.Equal(t, request, received)
	}
	return result
}

// classicCallBetween checks that a call took at least low and at most high.
func classicCallBetween(t *testing.T, took, low, high time.Duration, client string) {
	assert.GreaterOrEqual(t, took, low, client)
	assert.LessOrEqual(t, took, high, client)
}

// TestClassicCallReturnsWhatTheClassicClientReturns makes every call with a
// real classic client and, through ClassicCall, with a real modular client,
// against one server, and compares what the server and what the caller got.
func TestClassicCallReturnsWhatTheClassicClientReturns(t *testing.T) {
	type steps = []classicCallStep
	type outcome = classicCallOutcome
	fails, found := classicCallFails, classicCallFound
	model := func(attempts int) outcome { return outcome{attempts: attempts, requests: attempts, model: "model"} }
	apiError := func(attempts, status int, code, message string) outcome {
		return outcome{attempts: attempts, requests: attempts, err: message, errType: "*apierr.APIError", status: status, code: code}
	}
	missing := func(want outcome) outcome { want.missing = true; return want }
	network := func(attempts, requests int, message string) outcome {
		return outcome{attempts: attempts, requests: requests, err: "REQUEST" + message, errType: "*url.Error"}
	}
	const refused = "dial tcp 127.0.0.1: connect: connection refused"
	const reset = "read tcp 127.0.0.1->127.0.0.1: read: connection reset by peer"
	throttled, unavailable := fails(429, "TOO_MANY_REQUESTS", "slow down"), fails(503, "TEMPORARILY_UNAVAILABLE", "come back")
	throttledOnce, throttledTwice := apiError(1, 429, "TOO_MANY_REQUESTS", "slow down"), apiError(2, 429, "TOO_MANY_REQUESTS", "slow down")
	retryAfter := http.Header{"Retry-After": {"120"}}
	after := func(takes time.Duration) classicCallStep { step := throttled; step.takes = takes; return step }
	tests := []classicCallScenario{
		{name: "success on the first attempt", script: steps{found}, want: model(1)},
		{name: "429, then success", script: steps{throttled, found}, want: model(2), took: time.Second},
		{name: "503, then success", script: steps{unavailable, found}, want: model(2), took: time.Second},
		{name: "429 and 503, then success", script: steps{throttled, unavailable, found}, want: model(3), took: 3 * time.Second},
		// Neither client waits the two minutes that the header asks for.
		{name: "429 and 503 with Retry-After, then success", want: model(3), took: 3 * time.Second,
			script: steps{{status: 429, header: retryAfter, body: throttled.body}, {status: 503, header: retryAfter, body: unavailable.body}, found}},
		{name: "retried error code in a 400", script: steps{fails(400, "RESOURCE_EXHAUSTED", "too much"), found}, want: model(2), took: time.Second},
		{name: "retried error code in a 500", script: steps{fails(500, "REQUEST_LIMIT_EXCEEDED", "too much"), found}, want: model(2), took: time.Second},
		{name: "retried message in a 400", script: steps{fails(400, "BAD_REQUEST", "ClusterNotReadyException"), found}, want: model(2), took: time.Second},
		{name: "retried message in a 500", script: steps{fails(500, "INTERNAL_ERROR", "Timed out after 30s"), found}, want: model(2), took: time.Second},
		{name: "REQUEST_LIMIT_EXCEEDED in the message of a 400", script: steps{fails(400, "BAD_REQUEST", "REQUEST_LIMIT_EXCEEDED: too fast"), found},
			want: model(2), took: time.Second},
		{name: "400", script: steps{fails(400, "INVALID_PARAMETER_VALUE", "bad name")}, want: apiError(1, 400, "INVALID_PARAMETER_VALUE", "bad name")},
		{name: "401", script: steps{fails(401, "UNAUTHENTICATED", "who are you")}, want: apiError(1, 401, "UNAUTHENTICATED", "who are you")},
		{name: "403", script: steps{fails(403, "PERMISSION_DENIED", "not yours")}, want: apiError(1, 403, "PERMISSION_DENIED", "not yours")},
		{name: "409", script: steps{fails(409, "RESOURCE_ALREADY_EXISTS", "taken")}, want: apiError(1, 409, "RESOURCE_ALREADY_EXISTS", "taken")},
		{name: "500", script: steps{fails(500, "INTERNAL_ERROR", "broken")}, want: apiError(1, 500, "INTERNAL_ERROR", "broken")},
		{name: "404 with RESOURCE_DOES_NOT_EXIST", script: steps{fails(404, "RESOURCE_DOES_NOT_EXIST", "gone")},
			want: missing(apiError(1, 404, "RESOURCE_DOES_NOT_EXIST", "gone"))},
		{name: "404 with INVALID_PARAMETER_VALUE", script: steps{fails(404, "INVALID_PARAMETER_VALUE", "bad name")},
			want: apiError(1, 404, "INVALID_PARAMETER_VALUE", "bad name")},
		{name: "404 with an empty body", script: steps{{status: 404}}, want: missing(apiError(1, 404, "", "Not Found"))},
		{name: "404 with an HTML body", script: steps{{status: 404, body: "<html><h1>Nothing here</h1></html>"}},
			want: apiError(1, 404, "UNKNOWN", "<html><h1>Nothing here</h1></html>")},
		{name: "404 with an HTML body that has a pre block", script: steps{{status: 404, body: "<html><pre>Nothing here</pre></html>"}},
			want: missing(apiError(1, 404, "NOT_FOUND", "Nothing here"))},
		{name: "refused connection, then success", script: steps{{refused: true}, found},
			want: outcome{attempts: 2, requests: 1, model: "model"}, took: time.Second},
		{name: "reset connection, then success", script: steps{{reset: true}, found}, want: model(2), took: time.Second},
		{name: "refused connection until the retries run out", script: steps{{refused: true}}, retryTimeoutSeconds: 1,
			want: network(1, 0, refused), took: time.Second, sooner: true},
		{name: "reset connection until the retries run out", script: steps{{reset: true}}, retryTimeoutSeconds: 1,
			want: network(1, 1, reset), took: time.Second, sooner: true},
		{name: "retries run out after a retry timeout of 1 second", script: steps{throttled}, retryTimeoutSeconds: 1,
			want: throttledOnce, took: time.Second, sooner: true},
		{name: "retries run out after a retry timeout of 3 seconds", script: steps{throttled}, retryTimeoutSeconds: 3,
			want: throttledTwice, took: 3 * time.Second, sooner: true},
		// The first attempt ends 3 seconds before the 5 minutes of an unset
		// retry timeout do, which leaves time for one retry.
		{name: "retries run out after 5 minutes without a retry timeout", script: steps{after(297 * time.Second), throttled},
			httpTimeoutSeconds: 600, want: throttledTwice, took: 5 * time.Minute, sooner: true},
		// A negative retry timeout, which the provider sets when the user sets
		// none, is no timeout: the call still retries after half an hour.
		{name: "retries run out at the deadline with a negative retry timeout", script: steps{after(30 * time.Minute), throttled},
			retryTimeoutSeconds: -1, httpTimeoutSeconds: 3600, deadline: 30*time.Minute + 3*time.Second,
			want: throttledTwice, took: 30*time.Minute + 3*time.Second, sooner: true},
		{name: "retries run out at the deadline without a retry timeout", script: steps{throttled}, deadline: 3 * time.Second,
			want: throttledTwice, took: 3 * time.Second, sooner: true},
		{name: "retries run out at the deadline before the retry timeout", script: steps{throttled}, retryTimeoutSeconds: 60,
			deadline: 3 * time.Second, want: throttledTwice, took: 3 * time.Second, sooner: true},
		{name: "retries run out at the retry timeout before the deadline", script: steps{throttled}, retryTimeoutSeconds: 3,
			deadline: time.Minute, want: throttledTwice, took: 3 * time.Second, sooner: true},
		{name: "cancelled before the call", script: steps{found}, cancelled: classicCallBeforeTheCall,
			want: outcome{err: "failed in rate limiter: context canceled", errType: "*fmt.wrapError", canceled: true}},
		// Both clients return the 429 that they were about to retry.
		{name: "cancelled during the wait to retry", script: steps{throttled, found}, retryTimeoutSeconds: -1,
			cancelled: classicCallDuringTheWait, want: throttledOnce, took: 500 * time.Millisecond},
		{name: "cancelled during the request", script: steps{{holds: true}}, cancelled: classicCallDuringTheRequest,
			want: func() outcome { want := network(1, 1, "context canceled"); want.canceled = true; return want }()},
	}
	for _, tt := range tests {
		for _, op := range classicCallOps {
			t.Run(tt.name+", "+op.name, func(t *testing.T) {
				api := classicCallNewAPI(t, tt.script...)
				synctest.Test(t, func(t *testing.T) {
					classic := tt.run(t, api, op.request, op.classic)
					modular := tt.run(t, api, op.request, op.modular)
					want := tt.want
					want.err = strings.Replace(want.err, "REQUEST", op.failed, 1)
					assert.Equal(t, want, classic.classicCallOutcome, "the classic client")
					assert.Equal(t, classic.classicCallOutcome, modular.classicCallOutcome, "ClassicCall returns what the classic client returns")
					// The differences from the classic client, all on purpose.
					assert.Equal(t, want.err != "", classic.wrapped, "the classic client wraps every error in a *retries.Err")
					assert.False(t, modular.wrapped, "ClassicCall returns the error itself")
					low, high := tt.took, tt.took+time.Duration(max(want.attempts-1, 0))*classicCallMaxJitter
					classicCallBetween(t, classic.took, low, high, "the classic client")
					if tt.sooner {
						// ClassicCall returns when the wait that the classic client
						// doesn't finish would start: the one before retry number
						// want.attempts, of that many seconds and a jitter.
						low, high = classic.took-time.Duration(want.attempts)*time.Second-classicCallMaxJitter, classic.took-1
					}
					classicCallBetween(t, modular.took, low, high, "ClassicCall")
				})
			})
		}
	}
}

// The classic client ends a request that gets no byte for http_timeout_seconds,
// a minute by default, and doesn't retry it. ClassicCall sets no timeout: the
// request lasts until the timeout of the modular client, which the tests turn
// off, or until the context ends.
func TestClassicCallLeavesARequestWithoutAnAnswerToTheClient(t *testing.T) {
	api := classicCallNewAPI(t, classicCallStep{stalls: true})
	synctest.Test(t, func(t *testing.T) {
		tt, op := classicCallScenario{deadline: 10 * time.Minute}, classicCallOps[0]
		classic := tt.run(t, api, op.request, op.classic)
		modular := tt.run(t, api, op.request, op.modular)
		assert.Equal(t, op.failed+"request timed out after 1m0s of inactivity", classic.err)
		assert.Equal(t, time.Minute, classic.took)
		assert.Equal(t, op.failed+"context deadline exceeded", modular.err)
		assert.Equal(t, 10*time.Minute, modular.took)
		assert.Equal(t, 1, classic.attempts)
		assert.Equal(t, 1, modular.attempts)
	})
}

// A retrier counts the retries of its call and holds the retry timeout from
// the start of that call. Calls that shared one would wait longer from retry
// to retry, and would not retry once the retry timeout of the first has passed.
func TestClassicCallGivesEveryCallARetrierOfItsOwn(t *testing.T) {
	throttled := classicCallFails(429, "TOO_MANY_REQUESTS", "slow down")
	apis := make([]*classicCallAPI, 8)
	for i := range apis {
		apis[i] = classicCallNewAPI(t, throttled, throttled, classicCallFound)
	}
	synctest.Test(t, func(t *testing.T) {
		cfg := apis[0].config(t, func(cfg *config.Config) { cfg.RetryTimeoutSeconds = 6 })
		// The calls of a round run at once, and the second round starts after
		// the retry timeout of the first has passed.
		for round := range 2 {
			var calls sync.WaitGroup
			for _, api := range apis {
				c := api.client(t)
				calls.Go(func() {
					response, err := classicCallGet(t.Context(), cfg, c)
					if assert.NoError(t, err) {
						assert.Equal(t, "model", *response.RegisteredModelDatabricks.Id)
					}
				})
			}
			calls.Wait()
			for i, api := range apis {
				attempts, _ := api.restart()
				require.Len(t, attempts, 3, "round %d, call %d", round, i)
				for retry := 1; retry <= 2; retry++ {
					wait, backoff := attempts[retry].Sub(attempts[retry-1]), time.Duration(retry)*time.Second
					classicCallBetween(t, wait, backoff, backoff+classicCallMaxJitter, fmt.Sprintf("round %d, call %d, retry %d", round, i, retry))
				}
			}
			time.Sleep(6 * time.Second)
		}
	})
}

func TestClassicCallSharesTheRateLimitOfAConfig(t *testing.T) {
	api := classicCallNewAPI(t, classicCallFound)
	synctest.Test(t, func(t *testing.T) {
		// Two configs with equal fields have a rate limit each.
		rate := func(cfg *config.Config) { cfg.RateLimitPerSecond = 2 }
		limited, other := api.config(t, rate), api.config(t, rate)
		c := api.client(t)
		var mu sync.Mutex
		var calls sync.WaitGroup
		admitted := map[*config.Config][]time.Duration{}
		start := time.Now()
		for _, cfg := range []*config.Config{limited, limited, limited, limited, limited, limited, other, other} {
			calls.Go(func() {
				_, err := classicCallGet(t.Context(), cfg, c)
				assert.NoError(t, err)
				mu.Lock()
				defer mu.Unlock()
				admitted[cfg] = append(admitted[cfg], time.Since(start))
			})
		}
		calls.Wait()
		slices.Sort(admitted[limited])
		// A rate of 2 admits two calls at once and then one every half second.
		half := time.Second / 2
		assert.Equal(t, []time.Duration{0, 0, half, 2 * half, 3 * half, 4 * half}, admitted[limited])
		assert.Equal(t, []time.Duration{0, 0}, admitted[other])
		_, requests := api.restart()
		assert.Len(t, requests, 8)
	})
}

func TestClassicCallErrors(t *testing.T) {
	unresolved := func() *config.Config {
		return &config.Config{Host: "https://classic-call.test", Token: "token", Username: "user", Password: "password"}
	}
	_, classicErr := databricks.NewWorkspaceClient((*databricks.Config)(unresolved()))
	require.ErrorContains(t, classicErr, "more than one authorization method configured")
	runs := 0
	method := func(context.Context, ...call.Option) (string, error) { runs++; return "result", nil }
	tests := []struct {
		name   string
		cfg    *config.Config
		method func(context.Context, ...call.Option) (string, error)
		err    string
	}{
		{name: "no config", method: method, err: "sdkgoutility.ClassicCall: cfg is nil"},
		{name: "no method", cfg: &config.Config{}, err: "sdkgoutility.ClassicCall: method is nil"},
		{name: "no config and no method", err: "sdkgoutility.ClassicCall: cfg is nil"},
		// Retrier fails as the classic client does when it is built.
		{name: "config that doesn't resolve", cfg: unresolved(), method: method, err: classicErr.Error()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ClassicCall(t.Context(), tt.cfg, tt.method)
			assert.EqualError(t, err, tt.err)
			assert.Empty(t, result)
			assert.Zero(t, runs, "the method is not run")
		})
	}
}
