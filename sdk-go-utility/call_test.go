package sdkgoutility

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth/credentials"
	modularerr "github.com/databricks/sdk-go/core/apierr"
	"github.com/databricks/sdk-go/core/ops"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/sdk-go/options/client"
	"github.com/databricks/sdk-go/options/internaloptions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	// errCallRetriable is the only error that newCallTestRetrier retries.
	errCallRetriable = errors.New("retriable")
	errCallFatal     = errors.New("fatal")
	// errCallLimiter wraps a context error, as a limiter does when ctx ends.
	errCallLimiter = fmt.Errorf("failed in rate limiter: %w", context.Canceled)
)

// callTestRetrier retries, after delay, the errors that retriable accepts, and
// keeps the errors that it was asked about.
type callTestRetrier struct {
	delay     time.Duration
	retriable func(error) bool
	asked     []error
}

func (r *callTestRetrier) IsRetriable(err error) (time.Duration, bool) {
	r.asked = append(r.asked, err)
	return r.delay, r.retriable(err)
}

func newCallTestRetrier(delay time.Duration) *callTestRetrier {
	return &callTestRetrier{delay: delay, retriable: func(err error) bool { return err == errCallRetriable }}
}

// callTestLimiter counts its waits. Like the limiter of the classic client, it
// fails when ctx has ended. Its wait number failOn fails with errCallLimiter.
type callTestLimiter struct {
	waits  atomic.Int32
	failOn int32
}

func (l *callTestLimiter) Wait(ctx context.Context) error {
	if l.waits.Add(1) == l.failOn {
		return errCallLimiter
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("failed in rate limiter: %w", err)
	}
	return nil
}

// callTestClassic is a converter whose result shows the error that it got.
func callTestClassic(err error) error { return fmt.Errorf("classic: %w", err) }

// callGiven returns value when given, and otherwise the zero value of T.
func callGiven[T any](given bool, value T) T {
	if given {
		return value
	}
	var zero T
	return zero
}

// callFakeMethod is a method of a modular client without the client. It
// resolves its options and runs its attempts through ops.Execute, as a
// generated method does, with its own options such as the timeout of a client.
// attempt gets the number of the attempt, from 1.
func callFakeMethod(
	attempt func(ctx context.Context, n int) (string, error), own ...ops.Option,
) func(context.Context, ...call.Option) (string, error) {
	return func(ctx context.Context, opts ...call.Option) (string, error) {
		var resolved internaloptions.CallOptions
		for _, opt := range opts {
			if err := opt(&resolved); err != nil {
				return "", err
			}
		}
		n, result := 0, ""
		err := ops.Execute(ctx, func(ctx context.Context) error {
			n++
			var err error
			result, err = attempt(ctx, n)
			return err
		}, append(own, ops.WithRetrier(resolved.Retrier), ops.WithLimiter(resolved.RateLimiter))...)
		return result, err
	}
}

// callScripted is a fake method whose attempts fail with errs in turn, and
// succeed with "done" for nil. A failed attempt returns a result too.
func callScripted(errs ...error) func(context.Context, ...call.Option) (string, error) {
	return callFakeMethod(func(_ context.Context, n int) (string, error) {
		if err := errs[n-1]; err != nil {
			return "partial", err
		}
		return "done", nil
	})
}

func TestCallPassesItsOptionsToTheMethod(t *testing.T) {
	tests := []struct {
		name             string
		retrier, limiter bool
		options          int
	}{
		{"retrier and limiter", true, true, 2},
		{"retrier only", true, false, 1},
		{"limiter only", false, true, 1},
		{"neither", false, false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			givenRetrier := &callTestRetrier{delay: time.Hour, retriable: func(error) bool { return true }}
			givenLimiter := &callTestLimiter{}
			retrier := callGiven[ops.Retrier](tt.retrier, givenRetrier)
			limiter := callGiven[ops.Limiter](tt.limiter, givenLimiter)
			method := func(ctx context.Context, opts ...call.Option) (int, error) {
				var resolved internaloptions.CallOptions
				for _, opt := range opts {
					require.NoError(t, opt(&resolved))
				}
				assert.Equal(t, tt.retrier, resolved.Retrier != nil)
				assert.Equal(t, tt.limiter, resolved.RateLimiter != nil)
				assert.Zero(t, resolved.Timeout, "Call sets no timeout")
				if tt.retrier {
					delay, retry := resolved.Retrier().IsRetriable(errCallFatal)
					assert.Equal(t, time.Hour, delay)
					assert.True(t, retry)
				}
				if tt.limiter {
					assert.NoError(t, resolved.RateLimiter.Wait(ctx))
				}
				return len(opts), nil
			}

			options, err := Call(t.Context(), method, retrier, limiter, callTestClassic)

			require.NoError(t, err)
			assert.Equal(t, tt.options, options)
			assert.Equal(t, tt.retrier, len(givenRetrier.asked) == 1, "the method asks the retrier given to Call")
			assert.Equal(t, tt.limiter, givenLimiter.waits.Load() == 1, "the method waits on the limiter given to Call")
		})
	}
}

func TestCallReturnsTheResultOrTheClassicError(t *testing.T) {
	tests := []struct {
		name string
		// limiterFailsOn is the number of the limiter wait that fails.
		method                          func(context.Context, ...call.Option) (string, error)
		limiterFailsOn                  int32
		noRetrier, noLimiter, noClassic bool
		// wantErr is the message of the error, which must wrap wantIs. asked
		// and waits count the calls of the retrier and of the limiter.
		want, wantErr string
		wantIs        error
		asked         int
		waits         int32
	}{
		{name: "success", method: callScripted(nil), want: "done", waits: 1},
		{name: "success without retrier, limiter and classic", method: callScripted(nil),
			noRetrier: true, noLimiter: true, noClassic: true, want: "done"},
		{name: "failure", method: callScripted(errCallFatal),
			wantErr: "classic: fatal", wantIs: errCallFatal, asked: 1, waits: 1},
		{name: "failure without classic", method: callScripted(errCallFatal), noClassic: true,
			wantErr: "fatal", wantIs: errCallFatal, asked: 1, waits: 1},
		{name: "failure without retrier, limiter and classic", method: callScripted(errCallFatal),
			noRetrier: true, noLimiter: true, noClassic: true, wantErr: "fatal", wantIs: errCallFatal},
		{name: "retriable failure without retrier", method: callScripted(errCallRetriable, nil), noRetrier: true,
			wantErr: "classic: retriable", wantIs: errCallRetriable, waits: 1},
		{name: "two retries, then success", method: callScripted(errCallRetriable, errCallRetriable, nil),
			want: "done", asked: 2, waits: 3},
		{name: "retry, then failure", method: callScripted(errCallRetriable, errCallFatal),
			wantErr: "classic: fatal", wantIs: errCallFatal, asked: 2, waits: 2},
		{name: "limiter fails before the first attempt", method: callScripted(nil), limiterFailsOn: 1,
			wantErr: "classic: failed in rate limiter: context canceled", wantIs: errCallLimiter, waits: 1},
		// The wait to retry is over, so the context error of the limiter stays.
		{name: "limiter fails after the wait to retry", method: callScripted(errCallRetriable, nil), limiterFailsOn: 2,
			wantErr: "classic: failed in rate limiter: context canceled", wantIs: errCallLimiter, asked: 1, waits: 2},
		// Only a context error is replaced, and only while the SDK waits to retry.
		{name: "method fails by itself after a retried attempt", noLimiter: true,
			method: func(ctx context.Context, opts ...call.Option) (string, error) {
				_, _ = callScripted(errCallRetriable, nil)(ctx, opts...)
				return "partial", errCallFatal
			},
			wantErr: "classic: fatal", wantIs: errCallFatal, asked: 1},
		{name: "method wraps the context error of an attempt",
			method: func(ctx context.Context, opts ...call.Option) (string, error) {
				_, err := callScripted(context.Canceled)(ctx, opts...)
				return "partial", fmt.Errorf("wrapped: %w", err)
			},
			wantErr: "classic: wrapped: context canceled", wantIs: context.Canceled, asked: 1, waits: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			givenRetrier := newCallTestRetrier(time.Millisecond)
			givenLimiter := &callTestLimiter{failOn: tt.limiterFailsOn}
			retrier := callGiven[ops.Retrier](!tt.noRetrier, givenRetrier)
			limiter := callGiven[ops.Limiter](!tt.noLimiter, givenLimiter)
			classic := callGiven(!tt.noClassic, callTestClassic)

			got, err := Call(t.Context(), tt.method, retrier, limiter, classic)

			assert.Equal(t, tt.want, got)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
				assert.ErrorIs(t, err, tt.wantIs)
			}
			if tt.noClassic && tt.wantIs != nil {
				assert.Same(t, tt.wantIs, err, "without classic the error is returned as it is")
			}
			assert.Len(t, givenRetrier.asked, tt.asked)
			assert.Equal(t, tt.waits, givenLimiter.waits.Load())
		})
	}
}

// TestCallWhenTheContextEnds runs on the fake clock of a synctest bubble, so
// that the context ends at an exact moment of the call. The retrier waits a
// minute before a retry.
func TestCallWhenTheContextEnds(t *testing.T) {
	fails := func(context.Context, int) (string, error) { return "partial", errCallRetriable }
	hangs := func(ctx context.Context, _ int) (string, error) {
		<-ctx.Done()
		return "partial", &url.Error{Op: "Get", URL: "https://fake.test/model", Err: ctx.Err()}
	}
	failsThenHangs := func(ctx context.Context, n int) (string, error) {
		if n == 1 {
			return fails(ctx, n)
		}
		return hangs(ctx, n)
	}
	tests := []struct {
		name string
		// endsAfter is how long after the start of the call the context ends.
		endsAfter time.Duration
		attempt   func(ctx context.Context, n int) (string, error)
		noLimiter bool
		// wantErr has CONTEXT for the message of the context error, and
		// contextError tells whether the error wraps that error.
		wantErr      string
		contextError bool
		asked        int
	}{
		{name: "before the call", attempt: fails,
			wantErr: "classic: failed in rate limiter: CONTEXT", contextError: true},
		{name: "before the call, without limiter", attempt: hangs, noLimiter: true,
			wantErr: `classic: Get "https://fake.test/model": CONTEXT`, contextError: true, asked: 1},
		// The classic client returns the error that it was about to retry.
		{name: "during the wait to retry", endsAfter: 30 * time.Second, attempt: fails,
			wantErr: "classic: retriable", asked: 1},
		{name: "during the wait to retry, without limiter", endsAfter: 30 * time.Second, attempt: fails,
			noLimiter: true, wantErr: "classic: retriable", asked: 1},
		{name: "during the wait for the second retry", endsAfter: 90 * time.Second, attempt: fails,
			wantErr: "classic: retriable", asked: 2},
		{name: "during the request", endsAfter: 30 * time.Second, attempt: hangs,
			wantErr: `classic: Get "https://fake.test/model": CONTEXT`, contextError: true, asked: 1},
		// The error of the request of the retry stays, although the context ended.
		{name: "during the request of a retry", endsAfter: 90 * time.Second, attempt: failsThenHangs,
			wantErr: `classic: Get "https://fake.test/model": CONTEXT`, contextError: true, asked: 2},
		{name: "during the request of a retry, without limiter", endsAfter: 90 * time.Second,
			attempt: failsThenHangs, noLimiter: true,
			wantErr: `classic: Get "https://fake.test/model": CONTEXT`, contextError: true, asked: 2},
	}
	ends := []struct {
		name string
		err  error
		ctx  func(parent context.Context, after time.Duration) (context.Context, context.CancelFunc)
	}{
		{"cancelled", context.Canceled, func(parent context.Context, after time.Duration) (context.Context, context.CancelFunc) {
			ctx, cancel := context.WithCancel(parent)
			if after == 0 {
				cancel()
			} else {
				time.AfterFunc(after, cancel)
			}
			return ctx, cancel
		}},
		{"deadline passed", context.DeadlineExceeded, context.WithTimeout},
	}
	for _, tt := range tests {
		for _, end := range ends {
			t.Run(tt.name+", "+end.name, func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					givenRetrier := newCallTestRetrier(time.Minute)
					limiter := callGiven[ops.Limiter](!tt.noLimiter, &callTestLimiter{})
					ctx, cancel := end.ctx(t.Context(), tt.endsAfter)
					defer cancel()
					start := time.Now()

					got, err := Call(ctx, callFakeMethod(tt.attempt), givenRetrier, limiter, callTestClassic)

					assert.Empty(t, got)
					assert.EqualError(t, err, strings.ReplaceAll(tt.wantErr, "CONTEXT", end.err.Error()))
					assert.Equal(t, tt.contextError, errors.Is(err, end.err))
					assert.Equal(t, tt.endsAfter, time.Since(start), "Call returns when the context ends")
					assert.Len(t, givenRetrier.asked, tt.asked)
				})
			})
		}
	}
}

// The classic client returns the error that it was about to retry also when
// its retry timeout passes. The timeout of a modular client ends the wait of
// the SDK without ending the context of the caller.
func TestCallWhenTheTimeoutOfTheMethodPasses(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		fails := func(context.Context, int) (string, error) { return "partial", errCallRetriable }
		method := callFakeMethod(fails, ops.WithTimeout(90*time.Second))
		givenRetrier := newCallTestRetrier(time.Minute)
		start := time.Now()

		got, err := Call(t.Context(), method, givenRetrier, &callTestLimiter{}, callTestClassic)

		assert.Empty(t, got)
		assert.EqualError(t, err, "classic: retriable")
		assert.Equal(t, 90*time.Second, time.Since(start))
		assert.Len(t, givenRetrier.asked, 2)
	})
}

func TestCallKeepsNothingBetweenCalls(t *testing.T) {
	limiter := &callTestLimiter{}
	var wg sync.WaitGroup
	for range 32 {
		wg.Go(func() {
			got, err := Call(t.Context(), callScripted(errCallRetriable, nil),
				newCallTestRetrier(time.Millisecond), limiter, callTestClassic)
			assert.NoError(t, err)
			assert.Equal(t, "done", got)
		})
	}
	wg.Wait()
	assert.Equal(t, int32(64), limiter.waits.Load(), "two attempts of each call")
}

const (
	callModelPath    = "/api/2.0/mlflow/databricks/registered-models/get"
	callModelRequest = "GET " + callModelPath + "?name=m"
	callModelFound   = `200 {"registered_model_databricks": {"name": "m", "id": "model-id"}}`
	callModelMissing = `404 {"error_code": "RESOURCE_DOES_NOT_EXIST", "message": "no such model"}`
	callTooMany      = `429 {"error_code": "TOO_MANY_REQUESTS", "message": "slow down"}`
	// callHangs holds a request until its client gives up.
	callHangs = "hangs"
)

// callRegistry is the model registry API that the classic and the modular
// client call in turn. It answers the attempts of a call with responses, each
// a status and a body, and calls onRequest when it gets a request for the model.
type callRegistry struct {
	responses []string
	mu        sync.Mutex
	onRequest func()
	requests  []string
}

func (s *callRegistry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != callModelPath {
		// A classic client first asks for /.well-known/databricks-config.
		http.NotFound(w, r)
		return
	}
	s.mu.Lock()
	s.requests = append(s.requests, r.Method+" "+r.URL.RequestURI())
	response, onRequest := s.responses[len(s.requests)-1], s.onRequest
	s.mu.Unlock()
	onRequest()
	if response == callHangs {
		<-r.Context().Done()
		return
	}
	status, _ := strconv.Atoi(response[:3])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(response[4:]))
}

// nextCall sets onRequest for the next call, and returns the requests of the
// last call and forgets them.
func (s *callRegistry) nextCall(onRequest func()) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	requests := s.requests
	s.requests, s.onRequest = nil, onRequest
	return requests
}

// callEndableContext is a context that ends with err when done is closed. It
// lets a test pass a deadline at a moment of its choice, which a context of
// package context does by the clock only.
type callEndableContext struct {
	context.Context
	done chan struct{}
	err  error
}

func (c *callEndableContext) Done() <-chan struct{} { return c.done }

func (c *callEndableContext) Err() error {
	select {
	case <-c.done:
		return c.err
	default:
		return nil
	}
}

// callClassicError converts the error of a modular client with the parser of
// the classic client, from the response that the modular error carries.
func callClassicError(err error) error {
	modularErr, ok := errors.AsType[*modularerr.APIError](err)
	if !ok {
		return err
	}
	return apierr.GetAPIError(context.Background(), common.ResponseWrapper{
		ReadCloser: io.NopCloser(bytes.NewReader(modularErr.HTTPBody())),
		Response: &http.Response{
			StatusCode: modularErr.HTTPStatusCode(),
			Header:     modularErr.HTTPHeader(),
			Request:    &http.Request{URL: &url.URL{}},
		},
	})
}

// callGet returns the ID of the model m, read from host.
type callGet func(ctx context.Context, t *testing.T, host string) (string, error)

func callClassicGet(ctx context.Context, t *testing.T, host string) (string, error) {
	w, err := databricks.NewWorkspaceClient(&databricks.Config{Host: host, Token: "token"})
	require.NoError(t, err)
	resp, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "m"})
	if err != nil {
		return "", err
	}
	return resp.RegisteredModelDatabricks.Id, nil
}

// callModularGet reads the model with a modular client through Call. Its
// retrier retries a 429, as the classic client does, after retryAfter.
func callModularGet(retryAfter time.Duration) callGet {
	return func(ctx context.Context, t *testing.T, host string) (string, error) {
		creds, err := credentials.NewPATCredentials("token")
		require.NoError(t, err)
		c, err := modelregistry.NewClient(t.Context(), client.WithHost(host), client.WithCredentials(creds),
			client.WithoutConfigFile(), client.WithoutEnv())
		require.NoError(t, err)
		retrier := &callTestRetrier{delay: retryAfter, retriable: func(err error) bool {
			modularErr, ok := errors.AsType[*modularerr.APIError](err)
			return ok && modularErr.HTTPStatusCode() == http.StatusTooManyRequests
		}}
		name := "m"
		method := func(ctx context.Context, opts ...call.Option) (*modelregistry.GetRegisteredModelDatabricksResponse, error) {
			return c.GetRegisteredModelDatabricks(ctx, modelregistry.GetRegisteredModelDatabricksRequest{Name: &name}, opts...)
		}
		resp, err := Call(ctx, method, retrier, &callTestLimiter{}, callClassicError)
		if err != nil {
			assert.Nil(t, resp)
			return "", err
		}
		return *resp.RegisteredModelDatabricks.Id, nil
	}
}

// callOutcome is what the server and the caller see of one call: the requests
// that the server got, the ID of the model that the call returned, and the
// error. err is its message, with HOST for the URL of the server, and status
// and code are those of the classic *apierr.APIError in it.
type callOutcome struct {
	requests                    []string
	model, err                  string
	status                      int
	code                        string
	missing, canceled, timedOut bool
}

func newCallOutcome(requests []string, model string, err error, host string) callOutcome {
	got := callOutcome{requests: requests, model: model,
		canceled: errors.Is(err, context.Canceled), timedOut: errors.Is(err, context.DeadlineExceeded)}
	if err != nil {
		got.err = strings.ReplaceAll(err.Error(), host, "HOST")
		got.missing = apierr.IsMissing(err)
	}
	if apiErr, ok := errors.AsType[*apierr.APIError](err); ok {
		got.status, got.code = apiErr.StatusCode, apiErr.ErrorCode
	}
	return got
}

const (
	callEndsBeforeTheCall    = "before the call"
	callEndsAfterTheResponse = "after the first response"
	callEndsDuringTheRequest = "during the request"
)

// TestCallReturnsWhatTheClassicClientReturns makes the same call with a real
// classic client and, through Call, with a real modular client, against one
// server, and compares what the server got and what the caller got back.
func TestCallReturnsWhatTheClassicClientReturns(t *testing.T) {
	requestError := `Get "HOST` + callModelPath + `?name=m": `
	one, two := []string{callModelRequest}, []string{callModelRequest, callModelRequest}
	// A retry would succeed. Both clients return the 429 that they were about
	// to retry, and nothing of the context.
	retried := callOutcome{requests: one, err: "slow down", status: 429, code: "TOO_MANY_REQUESTS"}
	tests := []struct {
		name      string
		responses []string
		// ends is the moment at which the context ends, if it does, and
		// endsWith the error of the context.
		ends     string
		endsWith error
		want     callOutcome
	}{
		{name: "success after a retried 429", responses: []string{callTooMany, callModelFound},
			want: callOutcome{requests: two, model: "model-id"}},
		{name: "404", responses: []string{callModelMissing},
			want: callOutcome{requests: one, err: "no such model", status: 404, code: "RESOURCE_DOES_NOT_EXIST", missing: true}},
		{name: "cancelled before the call", ends: callEndsBeforeTheCall, endsWith: context.Canceled,
			want: callOutcome{err: "failed in rate limiter: context canceled", canceled: true}},
		{name: "deadline passed before the call", ends: callEndsBeforeTheCall, endsWith: context.DeadlineExceeded,
			want: callOutcome{err: "failed in rate limiter: context deadline exceeded", timedOut: true}},
		{name: "cancelled during the wait to retry", responses: []string{callTooMany, callModelFound},
			ends: callEndsAfterTheResponse, endsWith: context.Canceled, want: retried},
		{name: "deadline passed during the wait to retry", responses: []string{callTooMany, callModelFound},
			ends: callEndsAfterTheResponse, endsWith: context.DeadlineExceeded, want: retried},
		{name: "cancelled during the request", responses: []string{callHangs},
			ends: callEndsDuringTheRequest, endsWith: context.Canceled,
			want: callOutcome{requests: one, err: requestError + "context canceled", canceled: true}},
		{name: "deadline passed during the request", responses: []string{callHangs},
			ends: callEndsDuringTheRequest, endsWith: context.DeadlineExceeded,
			want: callOutcome{requests: one, err: requestError + "context deadline exceeded", timedOut: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &callRegistry{responses: tt.responses}
			server := httptest.NewServer(api)
			defer server.Close()
			run := func(get callGet) callOutcome {
				done := make(chan struct{})
				end := sync.OnceFunc(func() { close(done) })
				onRequest, afterResponse := func() {}, func() {}
				switch tt.ends {
				case callEndsBeforeTheCall:
					end()
				case callEndsAfterTheResponse:
					afterResponse = end
				case callEndsDuringTheRequest:
					onRequest = end
				}
				// A connection is idle once its client has read the whole response,
				// which is right before a client that retries starts its wait.
				traced := httptrace.WithClientTrace(context.Background(),
					&httptrace.ClientTrace{PutIdleConn: func(error) { afterResponse() }})
				ctx := &callEndableContext{Context: traced, done: done, err: tt.endsWith}
				api.nextCall(onRequest)
				model, err := get(ctx, t, server.URL)
				return newCallOutcome(api.nextCall(nil), model, err, server.URL)
			}

			classic := run(callClassicGet)
			// Only the end of its context may end the wait to retry of the modular
			// client. A retry that is to happen happens at once.
			modular := run(callModularGet(callGiven(tt.ends != "", time.Minute)))

			assert.Equal(t, tt.want, classic, "the classic client")
			assert.Equal(t, classic, modular, "Call returns what the classic client returns")
		})
	}
}
