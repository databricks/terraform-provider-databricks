package sdkgoutility

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth/credentials"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	attemptTimeoutGetPath    = "/api/2.0/mlflow/databricks/registered-models/get"
	attemptTimeoutCreatePath = "/api/2.0/mlflow/registered-models/create"
	// attemptTimeoutModelJSON answers both calls of the tests.
	attemptTimeoutModelJSON = `{"registered_model":{"name":"m"},"registered_model_databricks":{"name":"m"}}`
	// attemptTimeoutHost is the host of the tests that use no server.
	attemptTimeoutHost = "https://attempt-timeout.test"
	// attemptTimeoutPause is the pause between two steps of a script. It is
	// shorter than the timeout of 10 seconds that the scripts are played with.
	attemptTimeoutPause = 9 * time.Second
)

// attemptTimeoutLoader replaces the loaders of a classic config, so that a
// test reads neither the environment nor a config file.
type attemptTimeoutLoader struct{ err error }

func (attemptTimeoutLoader) Name() string                     { return "attempt-timeout-test" }
func (l attemptTimeoutLoader) Configure(*config.Config) error { return l.err }

func attemptTimeoutConfig(host string, seconds int) *config.Config {
	return &config.Config{
		Host:                host,
		Token:               "token",
		HTTPTimeoutSeconds:  seconds,
		RetryTimeoutSeconds: -1, // No limit on the whole call, as in the provider.
		Loaders:             []config.Loader{attemptTimeoutLoader{}},
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
			return nil, nil
		},
	}
}

// attemptTimeoutClients are a real classic client and a real modular client
// of one config. Both send over the same transport, the modular one through
// AttemptTimeout.
type attemptTimeoutClients struct {
	classic *databricks.WorkspaceClient
	modular *modelregistry.Client
}

func attemptTimeoutNewClients(t *testing.T, cfg *config.Config, base http.RoundTripper) attemptTimeoutClients {
	t.Helper()
	cfg.HTTPTransport = base
	classic, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	transport, err := AttemptTimeout(cfg, base)
	require.NoError(t, err)
	creds, err := credentials.NewPATCredentials(cfg.Token)
	require.NoError(t, err)
	modular, err := modelregistry.NewClient(t.Context(),
		client.WithHost(cfg.Host), client.WithCredentials(creds),
		client.WithHTTPClient(&http.Client{Transport: transport}),
		client.WithoutConfigFile(), client.WithoutEnv(), client.WithTimeout(0))
	require.NoError(t, err)
	return attemptTimeoutClients{classic: classic, modular: modular}
}

// classicCall gets the model called name, or creates it with a description of
// 30 kB if create is set, and returns what it got or the text of the error.
func (c attemptTimeoutClients) classicCall(ctx context.Context, name string, create bool) string {
	if create {
		res, err := c.classic.ModelRegistry.CreateModel(ctx,
			ml.CreateModelRequest{Name: name, Description: strings.Repeat("d", 30000)})
		if err != nil {
			return err.Error()
		}
		return "created " + res.RegisteredModel.Name
	}
	res, err := c.classic.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: name})
	if err != nil {
		return err.Error()
	}
	return "got " + res.RegisteredModelDatabricks.Name
}

// modularCall makes the call of classicCall with the modular client.
func (c attemptTimeoutClients) modularCall(ctx context.Context, name string, create bool) string {
	if create {
		description := strings.Repeat("d", 30000)
		res, err := c.modular.CreateRegisteredModel(ctx,
			modelregistry.CreateRegisteredModelRequest{Name: &name, Description: &description})
		if err != nil {
			return err.Error()
		}
		return "created " + *res.RegisteredModel.Name
	}
	res, err := c.modular.GetRegisteredModelDatabricks(ctx,
		modelregistry.GetRegisteredModelDatabricksRequest{Name: &name})
	if err != nil {
		return err.Error()
	}
	return "got " + *res.RegisteredModelDatabricks.Name
}

// attemptTimeoutEnd is how a call ended and how long it took.
type attemptTimeoutEnd struct {
	outcome string
	after   time.Duration
}

// attemptTimeoutCall is one call that both clients make.
type attemptTimeoutCall struct {
	name   string
	create bool
	// ctx returns the context of the call, context.Background if nil.
	ctx func() (context.Context, context.CancelFunc)
}

func (c attemptTimeoutCall) run(do func(context.Context, string, bool) string) attemptTimeoutEnd {
	newContext := c.ctx
	if newContext == nil {
		newContext = func() (context.Context, context.CancelFunc) {
			return context.WithCancel(context.Background())
		}
	}
	ctx, cancel := newContext()
	defer cancel()
	start := time.Now()
	outcome := do(ctx, c.name, c.create)
	return attemptTimeoutEnd{outcome, time.Since(start)}
}

func attemptTimeoutCancelAfter(d time.Duration) func() (context.Context, context.CancelFunc) {
	return func() (context.Context, context.CancelFunc) {
		ctx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(d, cancel)
		return ctx, cancel
	}
}

func attemptTimeoutDeadlineAfter(d time.Duration) func() (context.Context, context.CancelFunc) {
	return func() (context.Context, context.CancelFunc) {
		return context.WithTimeout(context.Background(), d)
	}
}

// attemptTimeoutGetError and attemptTimeoutCreateError are the errors of the
// two calls when net/http fails them with text. HOST stands for the host.
func attemptTimeoutGetError(name, text string) string {
	return fmt.Sprintf("Get %q: %s", "HOST"+attemptTimeoutGetPath+"?name="+name, text)
}

func attemptTimeoutCreateError(text string) string {
	return fmt.Sprintf("Post %q: %s", "HOST"+attemptTimeoutCreatePath, text)
}

// attemptTimeoutServer answers the calls of the tests over proto and counts
// them. The name of the model picks what a get does: "never" gets no answer,
// "stalls" gets the headers and no body, "slow" gets the body byte by byte
// for longer than a timeout of one second, and any other name gets the model
// at once. A create gets the model once its body is read.
func attemptTimeoutServer(t *testing.T, proto string) (*httptest.Server, *atomic.Int32) {
	requests := &atomic.Int32{}
	release := make(chan struct{})
	handler := func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		assert.Equal(t, proto, r.Proto)
		name := r.URL.Query().Get("name")
		_, _ = io.Copy(io.Discard, r.Body)
		w.Header().Set("Content-Type", "application/json")
		switch name {
		case "never", "stalls":
			if name == "stalls" {
				w.(http.Flusher).Flush()
			}
			select {
			case <-r.Context().Done():
			case <-release:
			}
		case "slow":
			for i := range len(attemptTimeoutModelJSON) {
				_, _ = io.WriteString(w, attemptTimeoutModelJSON[i:i+1])
				w.(http.Flusher).Flush()
				time.Sleep(20 * time.Millisecond)
			}
		default:
			_, _ = io.WriteString(w, attemptTimeoutModelJSON)
		}
	}
	server := httptest.NewUnstartedServer(http.HandlerFunc(handler))
	if server.EnableHTTP2 = proto == "HTTP/2.0"; server.EnableHTTP2 {
		server.StartTLS()
	} else {
		server.Start()
	}
	t.Cleanup(func() {
		close(release) // Before Close, which waits for the handlers.
		server.Close()
	})
	return server, requests
}

// attemptTimeoutUpload is a transport that reads the request body as a slow
// connection does: in parts pieces, 100 ms apart, which takes longer than a
// timeout of one second if parts is 15. If stall is set, it stops reading
// after the first piece. With no parts it reads as the wrapped transport does.
type attemptTimeoutUpload struct {
	http.RoundTripper
	parts int
	stall bool
}

func (u attemptTimeoutUpload) RoundTrip(req *http.Request) (*http.Response, error) {
	if u.parts == 0 || req.Body == nil {
		return u.RoundTripper.RoundTrip(req)
	}
	paced := req.Clone(req.Context())
	paced.Body = &attemptTimeoutUploadBody{ReadCloser: req.Body, upload: u, ctx: req.Context(),
		size: int(req.ContentLength) / u.parts}
	return u.RoundTripper.RoundTrip(paced)
}

type attemptTimeoutUploadBody struct {
	io.ReadCloser
	upload attemptTimeoutUpload
	ctx    context.Context
	size   int
	reads  int
}

func (b *attemptTimeoutUploadBody) Read(p []byte) (int, error) {
	if b.reads++; b.upload.stall && b.reads > 1 {
		<-b.ctx.Done()
		return 0, context.Cause(b.ctx)
	}
	time.Sleep(100 * time.Millisecond)
	return b.ReadCloser.Read(p[:min(len(p), b.size)])
}

// The calls of the ticket, with a real classic client and a real modular
// client against one httptest.Server each, over HTTP/1.1 and over HTTP/2.
func TestAttemptTimeout_EndsLikeClassicClient(t *testing.T) {
	const expired = "request timed out after 1s of inactivity"
	tests := []struct {
		name   string
		call   attemptTimeoutCall
		upload attemptTimeoutUpload
		// classic is how the classic call ends, and no sooner than after. The
		// modular call ends the same way unless modular is set.
		classic, modular attemptTimeoutEnd
	}{{
		name:    "server answers at once",
		call:    attemptTimeoutCall{name: "m"},
		classic: attemptTimeoutEnd{outcome: "got m"},
	}, {
		name:    "server never answers",
		call:    attemptTimeoutCall{name: "never"},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("never", expired), time.Second},
	}, {
		// The classic client reports a stall in the body as the cancellation
		// that its timeout caused. AttemptTimeout names the timeout.
		name:    "server sends headers and then stalls",
		call:    attemptTimeoutCall{name: "stalls"},
		classic: attemptTimeoutEnd{"response body: context canceled", time.Second},
		modular: attemptTimeoutEnd{expired, time.Second},
	}, {
		name:    "server sends the body slowly but steadily for longer than the timeout",
		call:    attemptTimeoutCall{name: "slow"},
		classic: attemptTimeoutEnd{"got m", 1500 * time.Millisecond},
	}, {
		name:    "request body is sent slowly but steadily for longer than the timeout",
		call:    attemptTimeoutCall{name: "m", create: true},
		upload:  attemptTimeoutUpload{parts: 15},
		classic: attemptTimeoutEnd{"created m", 1500 * time.Millisecond},
	}, {
		name:    "request body stalls",
		call:    attemptTimeoutCall{name: "m", create: true},
		upload:  attemptTimeoutUpload{parts: 15, stall: true},
		classic: attemptTimeoutEnd{attemptTimeoutCreateError(expired), time.Second},
	}, {
		name:    "context is cancelled",
		call:    attemptTimeoutCall{name: "never", ctx: attemptTimeoutCancelAfter(100 * time.Millisecond)},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("never", "context canceled"), 100 * time.Millisecond},
	}, {
		name:    "context deadline passes before the timeout",
		call:    attemptTimeoutCall{name: "never", ctx: attemptTimeoutDeadlineAfter(100 * time.Millisecond)},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("never", "context deadline exceeded"), 100 * time.Millisecond},
	}}
	for _, proto := range []string{"HTTP/1.1", "HTTP/2.0"} {
		for _, tc := range tests {
			t.Run(proto+"/"+tc.name, func(t *testing.T) {
				t.Parallel()
				server, requests := attemptTimeoutServer(t, proto)
				tc.upload.RoundTripper = server.Client().Transport
				clients := attemptTimeoutNewClients(t, attemptTimeoutConfig(server.URL, 1), tc.upload)
				if tc.modular == (attemptTimeoutEnd{}) {
					tc.modular = tc.classic
				}
				for _, side := range []struct {
					want attemptTimeoutEnd
					do   func(context.Context, string, bool) string
				}{{tc.classic, clients.classicCall}, {tc.modular, clients.modularCall}} {
					got := tc.call.run(side.do)
					assert.Equal(t, strings.ReplaceAll(side.want.outcome, "HOST", server.URL), got.outcome)
					assert.GreaterOrEqual(t, got.after, side.want.after)
				}
				if tc.call.ctx == nil { // A context can end before the request is sent.
					assert.Equal(t, int32(2), requests.Load(), "each client sends one attempt")
				}
			})
		}
	}
}

// attemptTimeoutScript is a transport that plays one exchange in memory. Its
// steps are attemptTimeoutPause apart and block durably, so that a synctest
// bubble can play it on a fake clock.
type attemptTimeoutScript struct {
	// rewind makes the transport replace the request body with a new one from
	// GetBody first, as net/http does when it sends a request again.
	rewind bool
	// uploads is the number of reads from the request body, a pause before each.
	uploads int
	// headers is the wait for the response headers, negative for no answer.
	headers time.Duration
	// pieces is the number of pieces of the response body, a pause before
	// each. With no pieces the body is there at once.
	pieces int
	// stall leaves the response body without an end, and without a start if
	// it has no pieces.
	stall bool
	// ended returns the error for a context that ended. If nil, it is the
	// cause of the context, as in net/http.
	ended func(context.Context) error
}

func (s attemptTimeoutScript) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx, body := req.Context(), req.Body
	if s.rewind {
		body, _ = req.GetBody()
	}
	for range s.uploads {
		if err := s.wait(ctx, attemptTimeoutPause); err != nil {
			return nil, err
		}
		_, _ = body.Read(make([]byte, 8))
	}
	if err := s.wait(ctx, s.headers); err != nil {
		return nil, err
	}
	answer := &attemptTimeoutScriptBody{script: s, ctx: ctx}
	n := s.pieces
	if n == 0 && !s.stall {
		n = 1
	}
	for i := range n {
		from, to := i*len(attemptTimeoutModelJSON)/n, (i+1)*len(attemptTimeoutModelJSON)/n
		answer.pieces = append(answer.pieces, attemptTimeoutModelJSON[from:to])
	}
	header := http.Header{"Content-Type": {"application/json"}}
	return &http.Response{StatusCode: http.StatusOK, Header: header, Body: answer, Request: req}, nil
}

// wait blocks for d, or for good if d is negative, unless ctx ends first.
func (s attemptTimeoutScript) wait(ctx context.Context, d time.Duration) error {
	if d == 0 {
		return nil
	}
	var passed <-chan time.Time
	if d > 0 {
		timer := time.NewTimer(d)
		defer timer.Stop()
		passed = timer.C
	}
	select {
	case <-passed:
		return nil
	case <-ctx.Done():
		if s.ended != nil {
			return s.ended(ctx)
		}
		return context.Cause(ctx)
	}
}

type attemptTimeoutScriptBody struct {
	script attemptTimeoutScript
	ctx    context.Context
	pieces []string
}

func (b *attemptTimeoutScriptBody) Read(p []byte) (int, error) {
	if len(b.pieces) == 0 {
		if !b.script.stall {
			return 0, io.EOF
		}
		return 0, b.script.wait(b.ctx, -1)
	}
	if b.script.pieces > 0 {
		if err := b.script.wait(b.ctx, attemptTimeoutPause); err != nil {
			return 0, err
		}
	}
	n := copy(p, b.pieces[0])
	b.pieces = b.pieces[1:]
	return n, nil
}

func (b *attemptTimeoutScriptBody) Close() error { return nil }

// attemptTimeoutPlay makes call with both clients against script, with a
// timeout of seconds, in a synctest bubble. It returns how each call ended
// and how long it took on the fake clock.
func attemptTimeoutPlay(
	t *testing.T, seconds int, script attemptTimeoutScript, call attemptTimeoutCall,
) (classic, modular attemptTimeoutEnd) {
	synctest.Test(t, func(t *testing.T) {
		clients := attemptTimeoutNewClients(t, attemptTimeoutConfig(attemptTimeoutHost, seconds), script)
		classic = call.run(clients.classicCall)
		modular = call.run(clients.modularCall)
	})
	classic.outcome = strings.ReplaceAll(classic.outcome, attemptTimeoutHost, "HOST")
	modular.outcome = strings.ReplaceAll(modular.outcome, attemptTimeoutHost, "HOST")
	return classic, modular
}

// The rules of the wait, on a fake clock: both clients end the same exchange
// at the same instant. The timeout is 10 seconds and a pause is 9.
func TestAttemptTimeout_WaitsLikeClassicClient(t *testing.T) {
	const expired = "request timed out after 10s of inactivity"
	get, create := attemptTimeoutCall{name: "m"}, attemptTimeoutCall{name: "m", create: true}
	tests := []struct {
		name   string
		script attemptTimeoutScript
		call   attemptTimeoutCall
		// The modular call ends as the classic call unless modular is set.
		classic, modular attemptTimeoutEnd
	}{{
		name:    "no response headers",
		script:  attemptTimeoutScript{headers: -1},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("m", expired), 10 * time.Second},
	}, {
		name:    "response headers do not restart the wait",
		script:  attemptTimeoutScript{headers: 6 * time.Second, stall: true},
		classic: attemptTimeoutEnd{"response body: context canceled", 10 * time.Second},
		modular: attemptTimeoutEnd{expired, 10 * time.Second},
	}, {
		name:    "reads from the response body restart the wait",
		script:  attemptTimeoutScript{pieces: 3, stall: true},
		classic: attemptTimeoutEnd{"response body: context canceled", 37 * time.Second},
		modular: attemptTimeoutEnd{expired, 37 * time.Second},
	}, {
		name:    "slow response body that ends",
		script:  attemptTimeoutScript{pieces: 5},
		classic: attemptTimeoutEnd{"got m", 45 * time.Second},
	}, {
		name:    "reads from the request body restart the wait",
		script:  attemptTimeoutScript{uploads: 3, headers: 9 * time.Second},
		call:    create,
		classic: attemptTimeoutEnd{"created m", 36 * time.Second},
	}, {
		name:    "request body stalls",
		script:  attemptTimeoutScript{uploads: 2, headers: -1},
		call:    create,
		classic: attemptTimeoutEnd{attemptTimeoutCreateError(expired), 28 * time.Second},
	}, {
		// The classic client does not watch a body that net/http rewound.
		name:    "reads from a rewound request body restart the wait",
		script:  attemptTimeoutScript{rewind: true, uploads: 3, headers: 9 * time.Second},
		call:    create,
		classic: attemptTimeoutEnd{attemptTimeoutCreateError(expired), 10 * time.Second},
		modular: attemptTimeoutEnd{"created m", 36 * time.Second},
	}, {
		name:    "transport reports context.Canceled",
		script:  attemptTimeoutScript{headers: -1, ended: context.Context.Err},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("m", expired), 10 * time.Second},
	}, {
		// The classic client compares with context.Canceled by identity.
		name: "transport wraps context.Canceled",
		script: attemptTimeoutScript{headers: -1, ended: func(ctx context.Context) error {
			return fmt.Errorf("dial: %w", ctx.Err())
		}},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("m", "dial: context canceled"), 10 * time.Second},
		modular: attemptTimeoutEnd{attemptTimeoutGetError("m", expired), 10 * time.Second},
	}, {
		name: "transport fails with another error once the timeout passed",
		script: attemptTimeoutScript{headers: -1, ended: func(context.Context) error {
			return errors.New("broken pipe")
		}},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("m", "broken pipe"), 10 * time.Second},
	}, {
		name:    "context is cancelled first",
		script:  attemptTimeoutScript{headers: -1},
		call:    attemptTimeoutCall{name: "m", ctx: attemptTimeoutCancelAfter(3 * time.Second)},
		classic: attemptTimeoutEnd{attemptTimeoutGetError("m", "context canceled"), 3 * time.Second},
	}, {
		// The classic client adds "response body: " to every error of a read
		// from the body. The modular client adds nothing.
		name:    "context is cancelled first while the body is read",
		script:  attemptTimeoutScript{stall: true},
		call:    attemptTimeoutCall{name: "m", ctx: attemptTimeoutCancelAfter(3 * time.Second)},
		classic: attemptTimeoutEnd{"response body: context canceled", 3 * time.Second},
		modular: attemptTimeoutEnd{"context canceled", 3 * time.Second},
	}, {
		name:    "context deadline passes first while the body is read",
		script:  attemptTimeoutScript{pieces: 2, stall: true},
		call:    attemptTimeoutCall{name: "m", ctx: attemptTimeoutDeadlineAfter(25 * time.Second)},
		classic: attemptTimeoutEnd{"response body: context deadline exceeded", 25 * time.Second},
		modular: attemptTimeoutEnd{"context deadline exceeded", 25 * time.Second},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.call.name == "" {
				tc.call = get
			}
			if tc.modular == (attemptTimeoutEnd{}) {
				tc.modular = tc.classic
			}
			classic, modular := attemptTimeoutPlay(t, 10, tc.script, tc.call)
			assert.Equal(t, tc.classic, classic, "classic client")
			assert.Equal(t, tc.modular, modular, "modular client")
		})
	}
}

func TestAttemptTimeout_TakesTimeoutFromConfig(t *testing.T) {
	tests := []struct {
		name    string
		seconds int
		want    attemptTimeoutEnd
	}{
		{"unset", 0, attemptTimeoutEnd{"request timed out after 1m0s of inactivity", time.Minute}},
		{"default of the provider", 65, attemptTimeoutEnd{"request timed out after 1m5s of inactivity", 65 * time.Second}},
		{"one second", 1, attemptTimeoutEnd{"request timed out after 1s of inactivity", time.Second}},
		{"negative", -1, attemptTimeoutEnd{"request timed out after -1s of inactivity", 0}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.want.outcome = attemptTimeoutGetError("m", tc.want.outcome)
			classic, modular := attemptTimeoutPlay(t, tc.seconds,
				attemptTimeoutScript{headers: -1}, attemptTimeoutCall{name: "m"})
			assert.Equal(t, tc.want, classic, "classic client")
			assert.Equal(t, tc.want, modular, "modular client")
		})
	}
}

func TestAttemptTimeout_RejectsBadArguments(t *testing.T) {
	loadErr := errors.New("cannot load")
	unresolved := attemptTimeoutConfig(attemptTimeoutHost, 1)
	unresolved.Loaders = []config.Loader{attemptTimeoutLoader{err: loadErr}}
	tests := []struct {
		name      string
		cfg       *config.Config
		transport http.RoundTripper
		want      string
	}{
		{"nil config", nil, http.DefaultTransport, "attempt timeout: config is nil"},
		{"nil transport", attemptTimeoutConfig(attemptTimeoutHost, 1), nil, "attempt timeout: transport is nil"},
		{"config that does not resolve", unresolved, http.DefaultTransport, "attempt timeout: resolve: cannot load"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			transport, err := AttemptTimeout(tc.cfg, tc.transport)
			assert.Nil(t, transport)
			require.Error(t, err)
			assert.True(t, strings.HasPrefix(err.Error(), tc.want), err.Error())
			assert.Equal(t, tc.cfg == unresolved, errors.Is(err, loadErr))
		})
	}
}

// attemptTimeoutRoundTrip is a transport made of a function.
type attemptTimeoutRoundTrip func(*http.Request) (*http.Response, error)

func (f attemptTimeoutRoundTrip) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

// attemptTimeoutTimers records the timers that a transport starts.
type attemptTimeoutTimers struct {
	mu     sync.Mutex
	timers []*time.Timer
}

// attemptTimeoutWatched returns AttemptTimeout around base, with a timeout of
// seconds, and the timers that it starts.
func attemptTimeoutWatched(
	t *testing.T, seconds int, base http.RoundTripper,
) (http.RoundTripper, *attemptTimeoutTimers) {
	t.Helper()
	transport, err := AttemptTimeout(attemptTimeoutConfig(attemptTimeoutHost, seconds), base)
	require.NoError(t, err)
	timers := &attemptTimeoutTimers{}
	transport.(*attemptTimeoutTransport).afterFunc = func(d time.Duration, f func()) *time.Timer {
		timers.mu.Lock()
		defer timers.mu.Unlock()
		timers.timers = append(timers.timers, time.AfterFunc(d, f))
		return timers.timers[len(timers.timers)-1]
	}
	return transport, timers
}

// running stops all timers and returns how many there are and how many of
// them were still running, which is how many a finished attempt left behind.
func (l *attemptTimeoutTimers) running() (started, running int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, timer := range l.timers {
		if timer.Stop() {
			running++
		}
	}
	return len(l.timers), running
}

func TestAttemptTimeout_SendsCopyOfRequest(t *testing.T) {
	type key struct{}
	body := func() io.ReadCloser { return io.NopCloser(strings.NewReader("body")) }
	rewindErr := errors.New("cannot rewind")
	tests := []struct {
		name      string
		body      io.ReadCloser
		rewound   io.ReadCloser
		rewindErr error
	}{
		{name: "no body"},
		{name: "http.NoBody", body: http.NoBody, rewound: http.NoBody},
		{name: "body", body: body(), rewound: body()},
		{name: "body that cannot be rewound", body: body(), rewindErr: rewindErr},
	}
	// watched tells whether sent is body with its reads counted as progress.
	watched := func(body, sent io.ReadCloser) bool {
		wrapper, ok := sent.(*attemptTimeoutRequestBody)
		return ok && wrapper.body == body
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sent *http.Request
			transport, _ := attemptTimeoutWatched(t, 10, attemptTimeoutRoundTrip(
				func(req *http.Request) (*http.Response, error) {
					sent = req
					return nil, errors.New("sent")
				}))
			ctx := context.WithValue(context.Background(), key{}, "value")
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, attemptTimeoutHost, nil)
			require.NoError(t, err)
			req.Body = tc.body
			if tc.body != nil {
				req.GetBody = func() (io.ReadCloser, error) { return tc.rewound, tc.rewindErr }
			}
			_, err = transport.RoundTrip(req)
			require.EqualError(t, err, "sent")

			assert.Equal(t, ctx, req.Context(), "the request of the caller keeps its context")
			assert.True(t, req.Body == tc.body, "the request of the caller keeps its body")
			assert.NotSame(t, req, sent)
			assert.Equal(t, "value", sent.Context().Value(key{}))
			plain := tc.body == nil || tc.body == http.NoBody
			assert.Equal(t, plain, sent.Body == tc.body)
			assert.Equal(t, !plain, watched(tc.body, sent.Body))
			if tc.body == nil {
				assert.Nil(t, sent.GetBody)
				return
			}
			rewound, err := sent.GetBody()
			assert.Equal(t, tc.rewindErr, err)
			plain = tc.rewound == nil || tc.rewound == http.NoBody
			assert.Equal(t, plain, rewound == tc.rewound)
			assert.Equal(t, !plain, watched(tc.rewound, rewound))
		})
	}
}

// Every way in which an attempt finishes releases its timer and its context
// at once, and hands on what the transport returned.
func TestAttemptTimeout_ReleasesFinishedAttempt(t *testing.T) {
	failed := errors.New("failed")
	closeErr := errors.New("cannot close")
	read := func(t *testing.T, body io.ReadCloser) {
		got, err := io.ReadAll(body)
		assert.Equal(t, "body", string(got))
		assert.NoError(t, err)
	}
	tests := []struct {
		name     string
		response func() *http.Response
		err      error
		// finish is what the caller does with the body, if there is one.
		finish func(*testing.T, io.ReadCloser)
	}{
		{name: "transport fails", err: failed},
		{name: "transport returns nothing"},
		{name: "response has no body", response: func() *http.Response { return &http.Response{} }},
		{name: "response comes with an error", err: failed, response: func() *http.Response {
			return &http.Response{Body: io.NopCloser(strings.NewReader("body"))}
		}},
		{name: "body is closed unread", finish: func(t *testing.T, body io.ReadCloser) {
			assert.Equal(t, closeErr, body.Close())
		}},
		{name: "body is closed after it is read", finish: func(t *testing.T, body io.ReadCloser) {
			read(t, body)
			assert.Equal(t, closeErr, body.Close())
		}},
		{name: "body is closed twice and read again", finish: func(t *testing.T, body io.ReadCloser) {
			assert.Equal(t, closeErr, body.Close())
			assert.Equal(t, closeErr, body.Close())
			read(t, body)
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sent *http.Request
			inner := &http.Response{Body: attemptTimeoutCloser{strings.NewReader("body"), closeErr}}
			if tc.finish == nil {
				inner = nil
			}
			if tc.response != nil {
				inner = tc.response()
			}
			transport, timers := attemptTimeoutWatched(t, 10, attemptTimeoutRoundTrip(
				func(req *http.Request) (*http.Response, error) {
					sent = req
					return inner, tc.err
				}))
			req, err := http.NewRequest(http.MethodGet, attemptTimeoutHost, nil)
			require.NoError(t, err)
			resp, err := transport.RoundTrip(req)
			assert.Same(t, inner, resp)
			assert.Equal(t, tc.err, err)
			if tc.finish != nil {
				assert.NoError(t, sent.Context().Err(), "the attempt lasts while its body is open")
				tc.finish(t, resp.Body)
			}
			assert.Equal(t, context.Canceled, context.Cause(sent.Context()))
			started, running := timers.running()
			assert.Equal(t, 1, started)
			assert.Zero(t, running)
		})
	}
}

// attemptTimeoutCloser is a body whose Close fails with err.
type attemptTimeoutCloser struct {
	io.Reader
	err error
}

func (c attemptTimeoutCloser) Close() error { return c.err }

// A body that is never closed holds its timer until the timeout passes
// without a read, or for as long if the context of the caller ends before.
func TestAttemptTimeout_ReleasesUnclosedBody(t *testing.T) {
	for _, cancelled := range []bool{false, true} {
		t.Run(fmt.Sprintf("cancelled=%t", cancelled), func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				transport, timers := attemptTimeoutWatched(t, 10, attemptTimeoutScript{})
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, attemptTimeoutHost, nil)
				require.NoError(t, err)
				resp, err := transport.RoundTrip(req)
				require.NoError(t, err)
				time.Sleep(attemptTimeoutPause)
				got, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Equal(t, attemptTimeoutModelJSON, string(got))
				attempt := resp.Body.(*attemptTimeoutResponseBody).watch.ctx
				want := "request timed out after 10s of inactivity"
				if cancelled {
					cancel()
					want = "context canceled"
				}
				time.Sleep(10*time.Second - time.Nanosecond)
				synctest.Wait()
				assert.Equal(t, cancelled, attempt.Err() != nil, "the wait restarted with the read")
				time.Sleep(time.Nanosecond)
				synctest.Wait()
				assert.EqualError(t, context.Cause(attempt), want)
				started, running := timers.running()
				assert.Equal(t, 1, started)
				assert.Zero(t, running)
			})
		})
	}
}

// Many attempts of every kind over a real server leave no timer and no
// goroutine behind. This test is not parallel, because it counts goroutines.
func TestAttemptTimeout_LeavesNoTimerOrGoroutine(t *testing.T) {
	const rounds = 50
	before := runtime.NumGoroutine()
	var timers *attemptTimeoutTimers
	t.Run("attempts", func(t *testing.T) {
		server, _ := attemptTimeoutServer(t, "HTTP/1.1")
		var transport http.RoundTripper
		transport, timers = attemptTimeoutWatched(t, 1, server.Client().Transport)
		// attempt gets the model called name, reads the body if read is set
		// and closes it. The caller gives up after patience.
		attempt := func(name string, read bool, patience time.Duration) {
			ctx, cancel := context.WithTimeout(context.Background(), patience)
			defer cancel()
			url := server.URL + attemptTimeoutGetPath + "?name=" + name
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			assert.NoError(t, err)
			resp, err := (&http.Client{Transport: transport}).Do(req)
			if err != nil {
				return
			}
			if read {
				_, _ = io.ReadAll(resp.Body)
			}
			assert.NoError(t, resp.Body.Close())
		}
		var group sync.WaitGroup
		for range rounds {
			attempt("m", true, time.Minute)  // Answered, read and closed.
			attempt("m", false, time.Minute) // Answered and closed unread.
			for _, name := range []string{"never", "stalls"} {
				group.Go(func() { attempt(name, true, 50*time.Millisecond) }) // Ended by the caller.
				group.Go(func() { attempt(name, true, time.Minute) })         // Ended by the timeout.
			}
		}
		group.Wait()
	})
	started, running := timers.running()
	assert.Equal(t, 6*rounds, started)
	assert.Zero(t, running)
	for end := time.Now().Add(10 * time.Second); runtime.NumGoroutine() > before && time.Now().Before(end); {
		time.Sleep(10 * time.Millisecond)
	}
	assert.LessOrEqual(t, runtime.NumGoroutine(), before)
}
