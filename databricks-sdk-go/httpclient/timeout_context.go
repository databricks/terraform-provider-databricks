package httpclient

import (
	"context"
	"io"
	"sync"
	"time"
)

type timeoutContext struct {
	ctx    context.Context
	cancel context.CancelFunc

	// Timeout is constant.
	// Deadline is updated when Tick function is called.
	timeout  time.Duration
	deadline time.Time

	// Protect against concurrent deadline reads/writes.
	lock sync.Mutex
}

type TimeoutTicker interface {
	Tick()
	Cancel()
}

func newTimeoutContext(ctx context.Context, timeout time.Duration) (context.Context, TimeoutTicker) {
	ctx, cancel := context.WithCancel(ctx)
	t := &timeoutContext{
		ctx:      ctx,
		cancel:   cancel,
		timeout:  timeout,
		deadline: time.Now().Add(timeout),
	}

	// Start goroutine to cancel the context when the deadline is reached.
	go t.run()
	return ctx, t
}

// Tick updates the deadline to the current time plus the timeout.
func (t *timeoutContext) Tick() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.deadline = time.Now().Add(t.timeout)
}

// Cancel cancels the context.
func (t *timeoutContext) Cancel() {
	t.cancel()
}

// Deadline returns the current deadline.
func (t *timeoutContext) Deadline() time.Time {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.deadline
}

func (t *timeoutContext) run() {
	for {
		ttl := time.Until(t.Deadline())
		if ttl <= 0 {
			t.cancel()
			return
		}

		timer := time.NewTimer(ttl)
		select {
		case <-timer.C:
			// Check if the deadline has been updated
			continue
		case <-t.ctx.Done():
			timer.Stop()
			return
		}
	}
}

// tickingReadCloser wraps an io.ReadCloser and calls the tick function on each read.
type tickingReadCloser struct {
	rc io.ReadCloser
	t  TimeoutTicker
}

func (t tickingReadCloser) Read(p []byte) (n int, err error) {
	defer t.t.Tick()
	return t.rc.Read(p)
}

func (t tickingReadCloser) Close() error {
	return t.rc.Close()
}

// cancellingReadCloser wraps an io.ReadCloser and calls the cancel function on close.
type cancellingReadCloser struct {
	rc io.ReadCloser
	t  TimeoutTicker
}

func (t cancellingReadCloser) Read(p []byte) (n int, err error) {
	return t.rc.Read(p)
}

func (t cancellingReadCloser) Close() error {
	defer t.t.Cancel()
	return t.rc.Close()
}

func newRequestBodyTicker(t TimeoutTicker, r io.ReadCloser) io.ReadCloser {
	return tickingReadCloser{r, t}
}

func newResponseBodyTicker(t TimeoutTicker, r io.ReadCloser) io.ReadCloser {
	return cancellingReadCloser{tickingReadCloser{r, t}, t}
}
