package sdkgoutility

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
)

// attemptTimeoutDefault is the timeout of the classic client when
// http_timeout_seconds is unset: defaultHTTPTimeout of its httpclient package,
// which is not exported.
const attemptTimeoutDefault = 60 * time.Second

// AttemptTimeout returns a transport that sends each request over transport
// and ends it once the HTTP timeout of cfg (http_timeout_seconds, 60 seconds
// when unset) passes without progress, as the classic client does. A modular
// client has no such limit: its only timeout bounds the whole call.
//
// The wait starts when the request is sent and restarts whenever the transport
// reads from the request body or the caller reads from the response body. The
// arrival of the response headers does not restart it. The response body stays
// readable after RoundTrip returns, so the limit covers it too.
//
// An attempt that the timeout ends fails with "request timed out after
// <duration> of inactivity", from RoundTrip or from Read on the response body.
// An attempt that ends because the context of the request ended first fails
// with the error of transport, unchanged.
//
// The caller has to close the response body, which releases the timer at once;
// the timer of a body that is never closed lasts until the timeout passes
// without a read. The caller also decides whether to retry a timed-out
// attempt, which the classic client does not, and limits the whole call with
// the context of the request. The timeout is read from cfg once, here, which
// resolves cfg if it is not resolved yet.
func AttemptTimeout(cfg *config.Config, transport http.RoundTripper) (http.RoundTripper, error) {
	if cfg == nil {
		return nil, errors.New("attempt timeout: config is nil")
	}
	if transport == nil {
		return nil, errors.New("attempt timeout: transport is nil")
	}
	clientConfig, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("attempt timeout: %w", err)
	}
	timeout := clientConfig.HTTPTimeout
	if timeout == 0 {
		timeout = attemptTimeoutDefault
	}
	return &attemptTimeoutTransport{
		transport: transport,
		timeout:   timeout,
		expired:   fmt.Errorf("request timed out after %s of inactivity", timeout),
		afterFunc: time.AfterFunc,
	}, nil
}

// attemptTimeoutTransport is the transport that AttemptTimeout returns.
type attemptTimeoutTransport struct {
	transport http.RoundTripper
	timeout   time.Duration
	// expired is the error of an attempt that the timeout ended. It is one
	// value per transport, so that it identifies the cause of a cancellation.
	expired error
	// afterFunc is time.AfterFunc. Tests replace it to see every timer.
	afterFunc func(time.Duration, func()) *time.Timer
}

// RoundTrip sends req over the wrapped transport under a watch. The request
// is copied, because a RoundTripper must not modify the one it is given.
func (t *attemptTimeoutTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	watch := t.watch(req.Context())
	attempt := req.WithContext(watch.ctx)
	attempt.Body = watch.requestBody(req.Body)
	if req.GetBody != nil {
		// net/http rewinds the body with GetBody when it sends the request
		// again on another connection.
		attempt.GetBody = func() (io.ReadCloser, error) {
			body, err := req.GetBody()
			return watch.requestBody(body), err
		}
	}
	resp, err := t.transport.RoundTrip(attempt)
	if err != nil || resp == nil || resp.Body == nil {
		// Nothing is left to read, so the attempt is over.
		watch.stop()
		return resp, watch.failure(err)
	}
	resp.Body = &attemptTimeoutResponseBody{body: resp.Body, watch: watch}
	return resp, nil
}

// watch starts the wait of one attempt that runs under parent.
func (t *attemptTimeoutTransport) watch(parent context.Context) *attemptTimeoutWatch {
	ctx, cancel := context.WithCancelCause(parent)
	w := &attemptTimeoutWatch{ctx: ctx, cancel: cancel, timeout: t.timeout, expired: t.expired}
	// check can run as soon as the timer exists, and it needs the timer.
	w.mu.Lock()
	defer w.mu.Unlock()
	w.timer = t.afterFunc(t.timeout, w.check)
	return w
}

// attemptTimeoutWatch cancels the context of one attempt when its deadline
// passes. Progress moves the deadline, and the timer follows it when it fires,
// so that a read costs no timer operation.
type attemptTimeoutWatch struct {
	ctx     context.Context
	cancel  context.CancelCauseFunc
	timeout time.Duration
	expired error

	mu sync.Mutex
	// deadline is the end of the wait that the last progress started. It is
	// zero until then, because the timer first fires when the first wait ends.
	deadline time.Time
	timer    *time.Timer
}

// progress restarts the wait.
func (w *attemptTimeoutWatch) progress() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.deadline = time.Now().Add(w.timeout)
}

// check runs when the timer fires. It ends the attempt if the deadline has
// passed, and waits for the rest of it otherwise.
func (w *attemptTimeoutWatch) check() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.ctx.Err() != nil {
		return // The attempt is over, so the timer must not start again.
	}
	if rest := time.Until(w.deadline); rest > 0 {
		w.timer.Reset(rest)
		return
	}
	w.cancel(w.expired)
}

// stop releases the timer and the context of a finished attempt.
func (w *attemptTimeoutWatch) stop() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.timer.Stop()
	w.cancel(nil)
}

// failure returns the error of an attempt that failed with err. If the watch
// ended the attempt before its parent context did, a transport reports that
// as context.Canceled, which becomes the timeout error here, or as the cause
// of the context, as net/http does, which is the timeout error already.
func (w *attemptTimeoutWatch) failure(err error) error {
	if context.Cause(w.ctx) == w.expired && errors.Is(err, context.Canceled) {
		return w.expired
	}
	return err
}

// requestBody returns body with every read from it counted as progress. An
// absent body is returned as it is, because net/http tells it by identity.
func (w *attemptTimeoutWatch) requestBody(body io.ReadCloser) io.ReadCloser {
	if body == nil || body == http.NoBody {
		return body
	}
	return &attemptTimeoutRequestBody{body: body, watch: w}
}

// attemptTimeoutRequestBody is a request body whose reads restart the wait.
type attemptTimeoutRequestBody struct {
	body  io.ReadCloser
	watch *attemptTimeoutWatch
}

func (b *attemptTimeoutRequestBody) Read(p []byte) (int, error) {
	defer b.watch.progress()
	return b.body.Read(p)
}

func (b *attemptTimeoutRequestBody) Close() error {
	return b.body.Close()
}

// attemptTimeoutResponseBody is a response body whose reads restart the wait,
// and whose Close ends it.
type attemptTimeoutResponseBody struct {
	body  io.ReadCloser
	watch *attemptTimeoutWatch
}

func (b *attemptTimeoutResponseBody) Read(p []byte) (int, error) {
	n, err := b.body.Read(p)
	b.watch.progress()
	if err != nil {
		err = b.watch.failure(err)
	}
	return n, err
}

func (b *attemptTimeoutResponseBody) Close() error {
	defer b.watch.stop()
	return b.body.Close()
}
