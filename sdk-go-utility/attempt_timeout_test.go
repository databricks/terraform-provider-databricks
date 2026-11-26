package sdkgoutility

import (
	"cmp"
	"context"
	"encoding/json"
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
	// attemptTimeoutFailed starts an error of net/http, without the host.
	attemptTimeoutFailed = `Post "/api/2.0/mlflow/registered-models/create": `
	attemptTimeoutModel  = `{"registered_model":{"name":"m"}}`
	attemptTimeoutHost   = "https://attempt-timeout.test" // For tests without a server.
	attemptTimeoutNever  = time.Duration(1 << 62)         // A wait that does not end.
)

// attemptTimeoutLoader keeps a config from reading the environment or a file.
type attemptTimeoutLoader struct{ err error }

func (attemptTimeoutLoader) Name() string                     { return "attempt-timeout-test" }
func (l attemptTimeoutLoader) Configure(*config.Config) error { return l.err }

func attemptTimeoutConfig(host string, seconds int) *config.Config {
	return &config.Config{
		Host: host, Token: "token", HTTPTimeoutSeconds: seconds,
		RetryTimeoutSeconds:  -1, // No limit on the whole call, as in the provider.
		Loaders:              []config.Loader{attemptTimeoutLoader{}},
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) { return nil, nil },
	}
}

// attemptTimeoutClients are a real classic client and a real modular client
// of one config. Both send over base, the modular one through AttemptTimeout.
type attemptTimeoutClients struct {
	host    string
	classic *databricks.WorkspaceClient
	modular *modelregistry.Client
}

func attemptTimeoutNewClients(t *testing.T, cfg *config.Config, base http.RoundTripper) attemptTimeoutClients {
	cfg.HTTPTransport = base
	classic, classicErr := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	transport, transportErr := AttemptTimeout(cfg, base)
	creds, credsErr := credentials.NewPATCredentials(cfg.Token)
	modular, modularErr := modelregistry.NewClient(t.Context(),
		client.WithHost(cfg.Host), client.WithCredentials(creds),
		client.WithHTTPClient(&http.Client{Transport: transport}),
		client.WithoutConfigFile(), client.WithoutEnv(), client.WithTimeout(0))
	require.NoError(t, errors.Join(classicErr, transportErr, credsErr, modularErr))
	return attemptTimeoutClients{host: cfg.Host, classic: classic, modular: modular}
}

// attemptTimeoutCall is the call that both clients make: they create the
// model called name with a description of 30 kB, under a context that is
// cancelled after cancelAfter and has a deadline of deadline, if set.
type attemptTimeoutCall struct {
	name                  string
	cancelAfter, deadline time.Duration
}

// attemptTimeoutEnd is how a call ended, "ok" or the text of its error
// without the host, and how long it took.
type attemptTimeoutEnd struct {
	outcome string
	after   time.Duration
}

func (c attemptTimeoutClients) ends(call attemptTimeoutCall) (classic, modular attemptTimeoutEnd) {
	big := strings.Repeat("d", 30000)
	run := func(create func(context.Context) error) attemptTimeoutEnd {
		start, outcome := time.Now(), "ok"
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		defer time.AfterFunc(cmp.Or(call.cancelAfter, attemptTimeoutNever), cancel).Stop()
		if call.deadline > 0 {
			ctx, cancel = context.WithTimeout(ctx, call.deadline)
			defer cancel()
		}
		if err := create(ctx); err != nil {
			outcome = strings.ReplaceAll(err.Error(), c.host, "")
		}
		return attemptTimeoutEnd{outcome, time.Since(start)}
	}
	classic = run(func(ctx context.Context) error {
		_, err := c.classic.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{Name: call.name, Description: big})
		return err
	})
	modular = run(func(ctx context.Context) error {
		_, err := c.modular.CreateRegisteredModel(ctx,
			modelregistry.CreateRegisteredModelRequest{Name: &call.name, Description: &big})
		return err
	})
	return classic, modular
}

// attemptTimeoutServer answers the calls of the tests over TLS and proto, and
// counts them. The name of the model, in the body or in the query, picks the
// answer: none for "never", the headers alone for "stalls", the model byte by
// byte for 1.65 seconds for "slow", and the model at once for any other name.
// It aborts when the client hangs up: an answer could still reach the client.
func attemptTimeoutServer(t *testing.T, proto string) (*httptest.Server, *atomic.Int32) {
	requests, release := &atomic.Int32{}, make(chan struct{})
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		assert.Equal(t, proto, r.Proto)
		var model struct{ Name string }
		if err := json.NewDecoder(r.Body).Decode(&model); err != nil && r.ContentLength > 0 {
			panic(http.ErrAbortHandler)
		}
		w.Header().Set("Content-Type", "application/json")
		switch cmp.Or(r.URL.Query().Get("name"), model.Name) {
		case "stalls":
			w.(http.Flusher).Flush()
			fallthrough
		case "never":
			select {
			case <-r.Context().Done():
			case <-release:
			}
			panic(http.ErrAbortHandler)
		case "slow":
			for i := range len(attemptTimeoutModel) {
				_, _ = io.WriteString(w, attemptTimeoutModel[i:i+1])
				w.(http.Flusher).Flush()
				time.Sleep(50 * time.Millisecond)
			}
		default:
			_, _ = io.WriteString(w, attemptTimeoutModel)
		}
	}))
	server.EnableHTTP2 = proto == "HTTP/2.0"
	server.StartTLS()
	t.Cleanup(func() {
		close(release) // Before Close, which waits for the handlers.
		server.Close()
	})
	return server, requests
}

// attemptTimeoutScript is a transport that plays the slow parts of one
// exchange. Its steps are pause apart and block durably, so that a synctest
// bubble can play them on a fake clock.
type attemptTimeoutScript struct {
	// base gets the request with its body paced. If nil, the script answers.
	base  http.RoundTripper
	pause time.Duration
	// rewind takes a new request body from GetBody first, as net/http can.
	rewind bool
	// uploads is the number of pieces in which the request body is read, a
	// pause before each and before its end, or zero for no pacing.
	// uploadStall stops the reading after the first piece.
	uploads     int
	uploadStall bool
	headers     time.Duration // The wait for the response headers.
	// pieces is the number of pieces of the response body, a pause before
	// each and before its end. stall leaves out that end, and the whole body
	// if pieces is zero.
	pieces int
	stall  bool
	// ended returns the error for a context that ended, its cause if nil.
	ended func(context.Context) error
}

func (s attemptTimeoutScript) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx, body := req.Context(), req.Body
	if s.rewind {
		body, _ = req.GetBody()
	}
	if s.uploads > 0 {
		size := (int(req.ContentLength) + s.uploads - 1) / s.uploads
		body = &attemptTimeoutPaced{ReadCloser: body, script: s, ctx: ctx, size: size, stalled: s.uploadStall, left: 1}
	}
	if s.base != nil {
		paced := req.Clone(ctx)
		paced.Body = body
		return s.base.RoundTrip(paced)
	}
	// The buffer holds the whole body, so that a read is a piece.
	if _, err := io.CopyBuffer(struct{ io.Writer }{io.Discard}, body, make([]byte, 1<<16)); err != nil {
		return nil, err
	}
	if err := s.wait(ctx, s.headers); err != nil {
		return nil, err
	}
	n := max(s.pieces, 1)
	answer := &attemptTimeoutPaced{ReadCloser: io.NopCloser(strings.NewReader(attemptTimeoutModel)),
		script: s, ctx: ctx, size: (len(attemptTimeoutModel) + n - 1) / n, stalled: s.stall, left: s.pieces}
	header := http.Header{"Content-Type": {"application/json"}}
	return &http.Response{StatusCode: http.StatusOK, Header: header, Body: answer, Request: req}, nil
}

// wait blocks for d unless ctx ends first.
func (s attemptTimeoutScript) wait(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		if s.ended != nil {
			return s.ended(ctx)
		}
		return context.Cause(ctx)
	}
}

// attemptTimeoutPaced is a body on a slow connection: a read gets size bytes
// at most after a pause of the script, and nothing after left reads if stalled.
type attemptTimeoutPaced struct {
	io.ReadCloser
	script  attemptTimeoutScript
	ctx     context.Context
	size    int
	stalled bool
	left    int
}

func (b *attemptTimeoutPaced) Read(p []byte) (int, error) {
	pause := b.script.pause
	if b.left--; b.stalled && b.left < 0 {
		pause = attemptTimeoutNever
	}
	if err := b.script.wait(b.ctx, pause); err != nil {
		return 0, err
	}
	return b.ReadCloser.Read(p[:min(len(p), b.size)])
}

// The cases of the ticket: a real classic client and a real modular client
// against one httptest.Server each, over HTTP/1.1 and HTTP/2, timeout of 1 s.
func TestAttemptTimeout_EndsLikeClassicClient(t *testing.T) {
	type call = attemptTimeoutCall
	type end = attemptTimeoutEnd
	const expired, failed, ms = "request timed out after 1s of inactivity", attemptTimeoutFailed, time.Millisecond
	tests := []struct {
		name   string
		call   call
		script attemptTimeoutScript
		// classic is how the classic call ends, and no sooner than after. The
		// modular call ends the same way unless modular is set.
		classic, modular end
	}{
		{name: "server answers at once", call: call{name: "m"}, classic: end{outcome: "ok"}},
		{name: "server never answers", call: call{name: "never"}, classic: end{failed + expired, 1000 * ms}},
		// The classic client reports a stall in the body as the cancellation
		// that its timeout caused. AttemptTimeout names the timeout.
		{name: "server sends headers and then stalls", call: call{name: "stalls"},
			classic: end{"response body: context canceled", 1000 * ms}, modular: end{expired, 1000 * ms}},
		{name: "server sends the body slowly but steadily for longer than the timeout",
			call: call{name: "slow"}, classic: end{"ok", 1500 * ms}},
		{name: "request body is sent slowly but steadily for longer than the timeout", call: call{name: "m"},
			script: attemptTimeoutScript{uploads: 15}, classic: end{"ok", 1500 * ms}},
		{name: "request body stalls", call: call{name: "m"}, classic: end{failed + expired, 1000 * ms},
			script: attemptTimeoutScript{uploads: 15, uploadStall: true}},
		{name: "context is cancelled", call: call{name: "never", cancelAfter: 500 * ms},
			classic: end{failed + "context canceled", 500 * ms}},
		{name: "context deadline passes before the timeout", call: call{name: "never", deadline: 500 * ms},
			classic: end{failed + "context deadline exceeded", 500 * ms}},
	}
	for _, proto := range []string{"HTTP/1.1", "HTTP/2.0"} {
		for _, tc := range tests {
			t.Run(proto+"/"+tc.name, func(t *testing.T) {
				t.Parallel()
				server, requests := attemptTimeoutServer(t, proto)
				tc.script.base, tc.script.pause = server.Client().Transport, 100*ms
				tc.modular = cmp.Or(tc.modular, tc.classic)
				clients := attemptTimeoutNewClients(t, attemptTimeoutConfig(server.URL, 1), tc.script)
				classic, modular := clients.ends(tc.call)
				assert.Equal(t, tc.classic.outcome, classic.outcome, "classic client")
				assert.Equal(t, tc.modular.outcome, modular.outcome, "modular client")
				assert.GreaterOrEqual(t, classic.after, tc.classic.after, "classic client")
				assert.GreaterOrEqual(t, modular.after, tc.modular.after, "modular client")
				assert.LessOrEqual(t, requests.Load(), int32(2), "no client sends a second attempt")
			})
		}
	}
}

// The rules of the wait, on the fake clock of a synctest bubble: both clients
// end the same exchange at the same instant. The timeout is 10 seconds unless
// seconds or unset says otherwise, and the pause of the scripts is 9.
func TestAttemptTimeout_WaitsLikeClassicClient(t *testing.T) {
	type script = attemptTimeoutScript
	type call = attemptTimeoutCall
	type end = attemptTimeoutEnd
	const expired, failed, s = "request timed out after 10s of inactivity", attemptTimeoutFailed, time.Second
	const never = attemptTimeoutNever
	wrapped := func(ctx context.Context) error { return fmt.Errorf("dial: %w", ctx.Err()) }
	broken := func(context.Context) error { return errors.New("broken pipe") }
	eof := func(context.Context) error { return io.EOF }
	tests := []struct {
		name    string
		seconds int
		unset   bool
		script  script
		call    call
		// The modular call ends as the classic call unless modular is set.
		classic, modular end
	}{
		{name: "no response headers", script: script{headers: never}, classic: end{failed + expired, 10 * s}},
		{name: "response headers do not restart the wait", script: script{headers: 6 * s, stall: true},
			classic: end{"response body: context canceled", 10 * s}, modular: end{expired, 10 * s}},
		{name: "reads from the response body restart the wait", script: script{pieces: 3, stall: true},
			classic: end{"response body: context canceled", 37 * s}, modular: end{expired, 37 * s}},
		{name: "slow response body that ends", script: script{pieces: 5}, classic: end{"ok", 54 * s}},
		{name: "reads from the request body restart the wait", script: script{uploads: 3}, classic: end{"ok", 54 * s}},
		{name: "request body stalls", script: script{uploads: 3, uploadStall: true}, classic: end{failed + expired, 19 * s}},
		// The classic client does not watch a body that net/http rewound.
		{name: "reads from a rewound request body restart the wait", script: script{rewind: true, uploads: 3},
			classic: end{failed + expired, 10 * s}, modular: end{"ok", 54 * s}},
		{name: "transport reports context.Canceled", script: script{headers: never, ended: context.Context.Err},
			classic: end{failed + expired, 10 * s}},
		// The classic client compares with context.Canceled by identity.
		{name: "transport wraps context.Canceled", script: script{headers: never, ended: wrapped},
			classic: end{failed + "dial: context canceled", 10 * s}, modular: end{failed + expired, 10 * s}},
		{name: "transport fails with another error once the timeout passed",
			script: script{headers: never, ended: broken}, classic: end{failed + "broken pipe", 10 * s}},
		{name: "context is cancelled first", script: script{headers: never}, call: call{cancelAfter: 3 * s},
			classic: end{failed + "context canceled", 3 * s}},
		// The classic client adds "response body: " to every error of a read
		// from the body. The modular client adds nothing.
		{name: "context is cancelled first while the body is read", script: script{stall: true},
			call:    call{cancelAfter: 3 * s},
			classic: end{"response body: context canceled", 3 * s}, modular: end{"context canceled", 3 * s}},
		{name: "timeout is unset", unset: true, script: script{headers: never},
			classic: end{failed + "request timed out after 1m0s of inactivity", 60 * s}},
		// A server can end its response in good order when the client hangs up.
		{name: "transport ends the body once the timeout passed", script: script{stall: true, ended: eof},
			classic: end{"ok", 10 * s}, modular: end{expired, 10 * s}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.script.pause, tc.call.name = 9*s, "m"; tc.seconds == 0 && !tc.unset {
				tc.seconds = 10
			}
			synctest.Test(t, func(t *testing.T) {
				cfg := attemptTimeoutConfig(attemptTimeoutHost, tc.seconds)
				classic, modular := attemptTimeoutNewClients(t, cfg, tc.script).ends(tc.call)
				assert.Equal(t, tc.classic, classic, "classic client")
				assert.Equal(t, cmp.Or(tc.modular, tc.classic), modular, "modular client")
			})
		})
	}
}

func TestAttemptTimeout_RejectsBadArguments(t *testing.T) {
	loadErr := errors.New("cannot load")
	unresolved := attemptTimeoutConfig(attemptTimeoutHost, 1)
	unresolved.Loaders = []config.Loader{attemptTimeoutLoader{err: loadErr}}
	for want, tc := range map[string]struct {
		cfg       *config.Config
		transport http.RoundTripper
	}{
		"attempt timeout: config is nil":        {nil, http.DefaultTransport},
		"attempt timeout: transport is nil":     {attemptTimeoutConfig(attemptTimeoutHost, 1), nil},
		"attempt timeout: resolve: cannot load": {unresolved, http.DefaultTransport},
	} {
		t.Run(want, func(t *testing.T) {
			transport, err := AttemptTimeout(tc.cfg, tc.transport)
			assert.Nil(t, transport)
			require.Error(t, err)
			assert.True(t, strings.HasPrefix(err.Error(), want), err.Error())
			assert.Equal(t, tc.cfg == unresolved, errors.Is(err, loadErr))
		})
	}
}

// attemptTimeoutSeen is a transport that keeps what a test wants to see of
// AttemptTimeout around it: its timers and the last request that it sent.
type attemptTimeoutSeen struct {
	transport http.RoundTripper // AttemptTimeout around this transport.
	inner     func(*http.Request) (*http.Response, error)
	mu        sync.Mutex
	timers    []*time.Timer
	sent      *http.Request
}

func (s *attemptTimeoutSeen) RoundTrip(req *http.Request) (*http.Response, error) {
	s.mu.Lock()
	s.sent = req
	s.mu.Unlock()
	return s.inner(req)
}

// attemptTimeoutWatched puts AttemptTimeout, with a timeout of seconds, around
// inner and returns what a test sees of it.
func attemptTimeoutWatched(
	t *testing.T, seconds int, inner func(*http.Request) (*http.Response, error),
) *attemptTimeoutSeen {
	seen := &attemptTimeoutSeen{inner: inner}
	transport, err := AttemptTimeout(attemptTimeoutConfig(attemptTimeoutHost, seconds), seen)
	require.NoError(t, err)
	transport.(*attemptTimeoutTransport).afterFunc = func(d time.Duration, f func()) *time.Timer {
		seen.mu.Lock()
		defer seen.mu.Unlock()
		seen.timers = append(seen.timers, time.AfterFunc(d, f))
		return seen.timers[len(seen.timers)-1]
	}
	seen.transport = transport
	return seen
}

// released checks that every attempt started one timer and left none behind.
func (s *attemptTimeoutSeen) released(t *testing.T, attempts int) {
	assert.Len(t, s.timers, attempts)
	for _, timer := range s.timers {
		assert.False(t, timer.Stop(), "the timer of a finished attempt still ran")
	}
}

func TestAttemptTimeout_SendsCopyOfRequest(t *testing.T) {
	type key struct{}
	body := func() io.ReadCloser { return attemptTimeoutCloser{strings.NewReader("body"), io.ErrClosedPipe} }
	// watched tells whether the transport got body with its reads counted
	// as progress, or body itself if there is nothing to read.
	watched := func(body, sent io.ReadCloser) bool {
		if body == nil || body == http.NoBody {
			return sent == body
		}
		wrapper, ok := sent.(*attemptTimeoutRequestBody)
		return ok && wrapper.body == body && sent.Close() == body.Close()
	}
	for name, tc := range map[string]struct {
		body, rewound io.ReadCloser
		rewindErr     error
	}{
		"no body":                     {},
		"http.NoBody":                 {body: http.NoBody, rewound: http.NoBody},
		"body":                        {body: body(), rewound: body()},
		"body that cannot be rewound": {body: body(), rewindErr: errors.New("cannot rewind")},
	} {
		t.Run(name, func(t *testing.T) {
			seen := attemptTimeoutWatched(t, 10, func(*http.Request) (*http.Response, error) {
				return nil, errors.New("sent")
			})
			ctx := context.WithValue(context.Background(), key{}, "value")
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, attemptTimeoutHost, nil)
			require.NoError(t, err)
			if req.Body = tc.body; tc.body != nil {
				req.GetBody = func() (io.ReadCloser, error) { return tc.rewound, tc.rewindErr }
			}
			_, err = seen.transport.RoundTrip(req)
			require.EqualError(t, err, "sent")
			assert.True(t, req.Context() == ctx && req.Body == tc.body, "the request of the caller is unchanged")
			sent := seen.sent
			assert.NotSame(t, req, sent)
			assert.Equal(t, "value", sent.Context().Value(key{}))
			assert.True(t, watched(tc.body, sent.Body))
			if tc.body == nil {
				assert.Nil(t, sent.GetBody)
				return
			}
			rewound, err := sent.GetBody()
			assert.Equal(t, tc.rewindErr, err)
			assert.True(t, watched(tc.rewound, rewound))
		})
	}
}

// attemptTimeoutCloser is a body whose Close fails with err.
type attemptTimeoutCloser struct {
	io.Reader
	err error
}

func (c attemptTimeoutCloser) Close() error { return c.err }

// An attempt hands on what the transport returned, and releases its context
// and timer when it is over, or when the timeout passes for an unclosed body.
func TestAttemptTimeout_ReleasesAttempt(t *testing.T) {
	const expired, s = "request timed out after 10s of inactivity", time.Second
	failed, closeErr := errors.New("failed"), errors.New("cannot close")
	response := func() *http.Response {
		return &http.Response{Body: attemptTimeoutCloser{strings.NewReader("body"), closeErr}}
	}
	tests := []struct {
		name     string
		response *http.Response
		err      error
		// What the caller does: read after 9 seconds, Close, cancel its context.
		read, close, cancel bool
		// The context of the attempt ends after held, with cause or as cancelled.
		held  time.Duration
		cause string
	}{
		{name: "transport fails", err: failed},
		{name: "transport returns nothing"},
		{name: "response has no body", response: &http.Response{}},
		{name: "response comes with an error", response: response(), err: failed},
		{name: "body is closed unread", response: response(), close: true},
		{name: "body is read and closed", response: response(), read: true, close: true, held: 9 * s},
		{name: "body is read and never closed", response: response(), read: true, held: 19 * s, cause: expired},
		{name: "unclosed body is read and the caller cancels", response: response(), read: true, cancel: true, held: 9 * s},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				seen := attemptTimeoutWatched(t, 10, func(*http.Request) (*http.Response, error) {
					return tc.response, tc.err
				})
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, attemptTimeoutHost, nil)
				require.NoError(t, err)
				start := time.Now()
				resp, err := seen.transport.RoundTrip(req)
				assert.Same(t, tc.response, resp)
				assert.Equal(t, tc.err, err)
				if tc.read {
					time.Sleep(9 * s)
					got, err := io.ReadAll(resp.Body)
					assert.Equal(t, "body", string(got))
					assert.NoError(t, err)
				}
				if tc.close {
					assert.Equal(t, closeErr, resp.Body.Close())
					assert.Equal(t, closeErr, resp.Body.Close(), "a second Close reaches the body too")
					_, err := resp.Body.Read(make([]byte, 1))
					assert.Equal(t, tc.read, err == io.EOF, "a Read after Close reaches the body too")
				}
				if tc.cancel {
					cancel()
				}
				<-seen.sent.Context().Done()
				assert.Equal(t, tc.held, time.Since(start))
				assert.EqualError(t, context.Cause(seen.sent.Context()), cmp.Or(tc.cause, "context canceled"))
				if tc.cancel { // The timer lasts until it fires, and does not start again.
					time.Sleep(10*s - tc.held)
				}
				synctest.Wait()
				seen.released(t, 1)
			})
		})
	}
}

// Many attempts of every kind over a real server leave no timer and no
// goroutine behind. This test is not parallel, because it counts goroutines.
func TestAttemptTimeout_LeavesNoTimerOrGoroutine(t *testing.T) {
	const rounds = 50
	before := runtime.NumGoroutine()
	var seen *attemptTimeoutSeen
	t.Run("attempts", func(t *testing.T) {
		server, _ := attemptTimeoutServer(t, "HTTP/1.1")
		seen = attemptTimeoutWatched(t, 1, server.Client().Transport.RoundTrip)
		// attempt gets the model called name and closes the body, read or not.
		attempt := func(name string, read bool, patience time.Duration) {
			ctx, cancel := context.WithTimeout(context.Background(), patience)
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"?name="+name, nil)
			assert.NoError(t, err)
			resp, err := (&http.Client{Transport: seen.transport}).Do(req)
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
	seen.released(t, 6*rounds)
	for end := time.Now().Add(10 * time.Second); runtime.NumGoroutine() > before && time.Now().Before(end); {
		time.Sleep(10 * time.Millisecond)
	}
	assert.LessOrEqual(t, runtime.NumGoroutine(), before)
}
