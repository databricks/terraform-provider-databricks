package sdkgoutility

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newClientLog is a logger of the classic SDK that keeps its debug messages.
type newClientLog struct {
	mu       sync.Mutex
	messages []string
}

func (l *newClientLog) Enabled(_ context.Context, level logger.Level) bool {
	return level >= logger.LevelDebug
}

func (l *newClientLog) Debugf(_ context.Context, format string, v ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages = append(l.messages, fmt.Sprintf(format, v...))
}

func (l *newClientLog) Tracef(context.Context, string, ...any) {}
func (l *newClientLog) Infof(context.Context, string, ...any)  {}
func (l *newClientLog) Warnf(context.Context, string, ...any)  {}
func (l *newClientLog) Errorf(context.Context, string, ...any) {}

// newClientLogged runs call with a context that logs to a logger of its own,
// and returns the debug messages and the error of the call.
func newClientLogged(ctx context.Context, call func(context.Context) (any, error)) ([]string, error) {
	log := &newClientLog{}
	_, err := call(logger.NewContext(ctx, log))
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.messages, err
}

// newClientLogDiffers matches the request header lines that the debug log of
// a client of NewClient doesn't share with the one of the classic client.
var newClientLogDiffers = regexp.MustCompile(`(?m)^> \* (Host|Traceparent|User-Agent): .*\n`)

// newClientLogOnlyIn returns the names of the differing header lines of
// message that other doesn't have, and every other such line as it is.
func newClientLogOnlyIn(message, other string) (lines []string) {
	for _, line := range strings.SplitAfter(message, "\n") {
		if !strings.Contains(other, line) {
			lines = append(lines, newClientLogDiffers.ReplaceAllString(line, "$1"))
		}
	}
	return lines
}

func TestNewClient_WritesDebugLogOfClassicClient(t *testing.T) {
	for _, headers := range []bool{false, true} {
		ws := newClientStart(t, &newClientWorkspace{})
		cfg := ws.config(func(cfg *config.Config) {
			cfg.DebugHeaders, cfg.WorkspaceID = headers, "42"
			if headers { // Shows the User-Agent to its end.
				cfg.DebugTruncateBytes = 1000
			}
		})
		classic := newClientClassic(t, cfg)
		modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
		require.NoError(t, err)
		for _, call := range newClientCalls {
			t.Run(fmt.Sprintf("%s/debug_headers=%t", call.name, headers), func(t *testing.T) {
				fromClassic, err := newClientLogged(t.Context(), func(ctx context.Context) (any, error) { return call.classic(ctx, classic) })
				require.NoError(t, err)
				fromModular, err := newClientLogged(t.Context(), func(ctx context.Context) (any, error) { return call.modular(ctx, modular) })
				require.NoError(t, err)
				require.Len(t, fromClassic, 1)
				require.Len(t, fromModular, 1, "one message for the one attempt")
				want, got := fromClassic[0], fromModular[0]
				assert.Regexp(t, `^[A-Z]+ /api/2\.0/mlflow/.*\n`, got)
				assert.Contains(t, got, "\n< HTTP/1.1 200 OK\n")
				assert.Contains(t, got, "\n<   \"registered_model\": {\n")
				// Debug logging doesn't change what is sent.
				received := ws.received()
				require.Len(t, received, 2)
				newClientSameRequest(t, received[0], received[1])
				if !headers {
					assert.Equal(t, want, got)
					assert.NotContains(t, got, "> * ")
					return
				}
				assert.Equal(t, newClientLogDiffers.ReplaceAllString(want, ""), newClientLogDiffers.ReplaceAllString(got, ""))
				assert.Contains(t, got, "\n> * Accept: application/json\n> * Authorization: Bearer dapi-secret\n")
				assert.Contains(t, got, "\n> * X-Databricks-Workspace-Id: 42\n")
				assert.Contains(t, got, "\n< * Content-Type: application/json\n")
				assert.Equal(t, []string{"Host", "Traceparent", "User-Agent"}, newClientLogOnlyIn(want, got))
				assert.Equal(t, []string{"Host", "Traceparent", "User-Agent"}, newClientLogOnlyIn(got, want))
				// Both logs show the headers that the workspace received, but
				// for the Host, which the classic client logs as empty.
				for i, message := range []string{want, got} {
					for _, name := range []string{"Traceparent", "User-Agent"} {
						assert.Contains(t, message, "\n> * "+name+": "+received[i].Header.Get(name)+"\n")
					}
				}
				assert.Contains(t, want, "\n> * Host: \n")
				assert.Contains(t, got, "\n> * Host: "+received[1].Host+"\n")
				assert.Equal(t, received[0].Host, received[1].Host)
			})
		}
	}
}

// newClientNoAnswer is a transport that answers no request. When the context
// of a request ends it returns that error bare, as a transport other than
// the one of net/http may, where net/http returns the cause of the context.
type newClientNoAnswer struct{}

func (newClientNoAnswer) RoundTrip(r *http.Request) (*http.Response, error) {
	<-r.Context().Done()
	return nil, r.Context().Err()
}

// newClientBounded returns a context that ends a call of t after 10 seconds.
// The client of NewClient has no timeout of its own, and a test whose helpers
// don't end a stalled attempt fails with it instead of waiting for the server.
func newClientBounded(t *testing.T) context.Context {
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	t.Cleanup(cancel)
	return ctx
}

func TestNewClient_EndsStalledAttemptLikeClassicClient(t *testing.T) {
	for _, tc := range []struct {
		name          string
		httpTransport http.RoundTripper
		wantReceived  int
	}{
		{name: "server that never answers", wantReceived: 2},
		{name: "cfg.HTTPTransport that never answers and reports the bare cancellation", httpTransport: newClientNoAnswer{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ws := newClientStart(t, &newClientWorkspace{answer: func(_ http.ResponseWriter, r *http.Request, _ int) { <-r.Context().Done() }})
			cfg := ws.config(func(cfg *config.Config) {
				cfg.HTTPTimeoutSeconds, cfg.HTTPTransport = 1, tc.httpTransport
			})
			classic := newClientClassic(t, cfg)
			modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
			require.NoError(t, err)
			ctx := newClientBounded(t)
			errs, took, logs := make([]error, 2), make([]time.Duration, 2), make([][]string, 2)
			var wg sync.WaitGroup
			for i, call := range []func(context.Context) (any, error){
				func(ctx context.Context) (any, error) { return newClientCalls[1].classic(ctx, classic) },
				func(ctx context.Context) (any, error) { return newClientCalls[1].modular(ctx, modular) },
			} {
				wg.Go(func() {
					start := time.Now()
					logs[i], errs[i] = newClientLogged(ctx, call)
					took[i] = time.Since(start)
				})
			}
			wg.Wait()
			want := `Get "` + ws.URL + `/api/2.0/mlflow/databricks/registered-models/get?name=my+model%2F%C3%BC%26x%3D1": request timed out after 1s of inactivity`
			for i, who := range []string{"classic", "modular"} {
				assert.EqualError(t, errs[i], want, who)
				assert.GreaterOrEqual(t, took[i], time.Second, who)
			}
			assert.Len(t, ws.received(), tc.wantReceived, "neither client tries again")
			// Both write the attempt with the timeout as its error. The retry
			// loop of the classic client adds a message of its own.
			attempt := "GET /api/2.0/mlflow/databricks/registered-models/get?name=my model/ü&x=1\n< Error: " + want
			assert.Equal(t, []string{attempt, "non-retriable error: " + want}, logs[0], "classic")
			assert.Equal(t, []string{attempt}, logs[1], "modular")
		})
	}
}

// TestNewClient_DebugLogShowsWhatIsSent covers the requests that Transport
// changes, refuses or answers itself: DebugLog is inside Transport, so the
// log shows each request as it is sent and each response as it arrives.
func TestNewClient_DebugLogShowsWhatIsSent(t *testing.T) {
	refused := errors.New("hook refused")
	redirect := func(w http.ResponseWriter, r *http.Request, _ int) {
		http.Redirect(w, r, "/login.html?error=private-link-validation-error", http.StatusFound)
	}
	const request = "DELETE /api/2.0/mlflow/registered-models/delete?name=my model/ü&x=1\n"
	for _, tc := range []struct {
		name      string
		configure func(*config.Config)
		answer    func(http.ResponseWriter, *http.Request, int)
		wantErr   string   // a part of the error of both calls
		classic   []string // how each message of the classic client starts
		modular   []string
	}{{
		name: "header of the cfg.Headers hook",
		configure: func(cfg *config.Config) {
			cfg.DebugHeaders, cfg.Headers = true, config.StaticHeaders(map[string]string{"X-From-Hook": "hook"})
		},
		classic: []string{request + "> * Host: \n> * Accept: application/json\n> * Authorization: Bearer dapi-secret\n> * Traceparent: "},
		modular: []string{request + "> * Host: 127.0.0.1:"},
	}, {
		// Only the retry loop of the classic client writes a message.
		name:      "request that the cfg.Headers hook refuses",
		configure: func(cfg *config.Config) { cfg.Headers = func(*http.Request) error { return refused } },
		wantErr:   "failed during request visitor: hook refused",
		classic:   []string{"non-retriable error: failed during request visitor: hook refused"},
	}, {
		// The classic client logs the first request with the last response.
		// DebugLog logs each request of the redirect, and the login page
		// without its text, because Transport closes the page unread. Both
		// calls fail, although the workspace answers the delete with a 200.
		name:    "private link login page",
		answer:  redirect,
		wantErr: "The requested workspace has AWS PrivateLink enabled",
		classic: []string{request + "< HTTP/1.1 200 OK\n< " + newClientLoginPage, "non-retriable error: The requested workspace has "},
		modular: []string{request + "< HTTP/1.1 302 Found", "GET /login.html?error=private-link-validation-error\n< HTTP/1.1 200 OK"},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			ws := newClientStart(t, &newClientWorkspace{answer: tc.answer})
			cfg := ws.config(tc.configure)
			classic := newClientClassic(t, cfg)
			modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
			require.NoError(t, err)
			fromClassic, classicErr := newClientLogged(t.Context(), func(ctx context.Context) (any, error) { return newClientCalls[3].classic(ctx, classic) })
			fromModular, modularErr := newClientLogged(t.Context(), func(ctx context.Context) (any, error) { return newClientCalls[3].modular(ctx, modular) })
			for who, err := range map[string]error{"classic": classicErr, "modular": modularErr} {
				if tc.wantErr == "" {
					assert.NoError(t, err, who)
				} else {
					assert.ErrorContains(t, err, tc.wantErr, who)
				}
			}
			for who, logged := range map[string][2][]string{"classic": {tc.classic, fromClassic}, "modular": {tc.modular, fromModular}} {
				require.Len(t, logged[1], len(logged[0]), who)
				for i, want := range logged[0] {
					assert.True(t, strings.HasPrefix(logged[1][i], want), "%s message %d is %q", who, i+1, logged[1][i])
				}
			}
			for who, logged := range map[string][]string{"classic": fromClassic, "modular": fromModular} {
				all := strings.Join(logged, "\n")
				assert.Equal(t, cfg.DebugHeaders, strings.Contains(all, "\n> * X-From-Hook: hook\n"), who)
				assert.Equal(t, who == "classic" && tc.answer != nil, strings.Contains(all, newClientLoginPage), who)
			}
		})
	}
}

// TestNewClient_TimesAttemptFromTheConnection pins a difference from the
// classic client, whose inactivity timeout runs while its request visitors
// do: a cfg.Headers hook that outlasts the timeout fails the classic attempt.
// AttemptTimeout is inside Transport, which runs the visitors, so the wait of
// the modular client starts after them and the call succeeds.
func TestNewClient_TimesAttemptFromTheConnection(t *testing.T) {
	t.Parallel()
	ws := newClientStart(t, &newClientWorkspace{})
	cfg := ws.config(func(cfg *config.Config) {
		cfg.HTTPTimeoutSeconds = 1
		cfg.Headers = func(*http.Request) error { time.Sleep(1300 * time.Millisecond); return nil }
	})
	classic := newClientClassic(t, cfg)
	modular, err := NewClient(t.Context(), cfg, modelregistry.NewClient)
	require.NoError(t, err)
	ctx := newClientBounded(t)
	var classicErr, modularErr error
	var wg sync.WaitGroup
	wg.Go(func() { _, classicErr = newClientCalls[1].classic(ctx, classic) })
	wg.Go(func() { _, modularErr = newClientCalls[1].modular(ctx, modular) })
	wg.Wait()
	assert.ErrorContains(t, classicErr, ": request timed out after 1s of inactivity")
	assert.NoError(t, modularErr)
	received := ws.received()
	require.Len(t, received, 1, "the classic client sends nothing")
	assert.Contains(t, received[0].Header.Get("User-Agent"), " sdk-go-modelregistry/")
}
