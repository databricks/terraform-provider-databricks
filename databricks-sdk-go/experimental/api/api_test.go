package api

import (
	"context"
	"errors"
	"testing"
	"time"
)

type mockRetrier struct {
	fn func(error) (time.Duration, bool)
}

func (m *mockRetrier) IsRetriable(err error) (time.Duration, bool) {
	return m.fn(err)
}

type mockLimiter struct {
	fn func(context.Context) error
}

func (m *mockLimiter) Wait(ctx context.Context) error {
	return m.fn(ctx)
}

func TestExecute_success(t *testing.T) {
	testCases := []struct {
		name string
		call Call
		opts []Option
	}{
		{
			name: "successful call no options",
			call: func(ctx context.Context) error { return nil },
			opts: []Option{},
		},
		{
			name: "successful call with timeout option",
			call: func(ctx context.Context) error { return nil },
			opts: []Option{WithTimeout(5 * time.Second)},
		},
		{
			name: "successful call with retrier option",
			call: func(ctx context.Context) error { return nil },
			opts: []Option{WithRetrier(func() Retrier {
				return &mockRetrier{fn: func(err error) (time.Duration, bool) { return 0, true }}
			})},
		},
		{
			name: "successful call with limiter option",
			call: func(ctx context.Context) error { return nil },
			opts: []Option{WithLimiter(&mockLimiter{fn: func(ctx context.Context) error { return nil }})},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Execute(context.Background(), tc.call, tc.opts...)
			if got != nil {
				t.Errorf("Execute(): got error %v, want nil", got)
			}
		})
	}
}

func TestExecute_retries(t *testing.T) {
	// Default retrier that only retries if the error is a retriableError.
	retriableError := errors.New("retriable error")
	nonRetriableError := errors.New("non-retriable error")
	retrier := &mockRetrier{fn: func(err error) (time.Duration, bool) {
		return 0, err == retriableError
	}}

	testCases := []struct {
		name          string
		callErrors    []error // sequence of errors the call should return
		retrier       Retrier
		wantErr       error
		wantCallCount int
	}{
		{
			name:          "no retrier - fail immediately",
			callErrors:    []error{retriableError},
			retrier:       nil, // no retrier
			wantErr:       retriableError,
			wantCallCount: 1,
		},
		{
			name:          "non-retriable error - fail immediately",
			callErrors:    []error{nonRetriableError},
			retrier:       retrier,
			wantErr:       nonRetriableError,
			wantCallCount: 1,
		},
		{
			name:          "retriable error - retry once then succeed",
			callErrors:    []error{retriableError, nil},
			retrier:       retrier,
			wantCallCount: 2,
		},
		{
			name:          "retriable error - retry multiple times then succeed",
			callErrors:    []error{retriableError, retriableError, retriableError, nil},
			retrier:       retrier,
			wantCallCount: 4,
		},
		{
			name:          "retriable error - retry then fail with non-retriable",
			callErrors:    []error{retriableError, nonRetriableError},
			retrier:       retrier,
			wantErr:       nonRetriableError,
			wantCallCount: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCallCount := 0
			call := func(ctx context.Context) error {
				err := tc.callErrors[gotCallCount]
				gotCallCount++
				return err
			}

			gotErr := Execute(context.Background(), call, WithRetrier(func() Retrier { return tc.retrier }))

			if gotCallCount != tc.wantCallCount {
				t.Errorf("Execute(): got call count %d, want %d", gotCallCount, tc.wantCallCount)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("Execute(): got error %v, want %v", gotErr, tc.wantErr)
			}
		})
	}
}

func TestExecute_timeout(t *testing.T) {
	testCases := []struct {
		name        string
		ctxTimeout  time.Duration
		optTimeout  time.Duration
		callDelay   time.Duration
		wantTimeout bool
	}{
		{
			name:        "no timeout - call succeeds",
			callDelay:   10 * time.Millisecond,
			wantTimeout: false,
		},
		{
			name:        "context timeout - call times out",
			ctxTimeout:  10 * time.Millisecond,
			callDelay:   50 * time.Millisecond,
			wantTimeout: true,
		},
		{
			name:        "option timeout - call times out",
			optTimeout:  10 * time.Millisecond,
			callDelay:   50 * time.Millisecond,
			wantTimeout: true,
		},
		{
			name:        "minimum timeout - context timeout",
			ctxTimeout:  10 * time.Millisecond,
			optTimeout:  100 * time.Millisecond, // would not timeout
			callDelay:   50 * time.Millisecond,
			wantTimeout: true,
		},
		{
			name:        "minimum timeout - option timeout",
			ctxTimeout:  100 * time.Millisecond, // would not timeout
			optTimeout:  10 * time.Millisecond,
			callDelay:   50 * time.Millisecond,
			wantTimeout: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Cancellable call that succeed after the call delay or returns
			// the context error if the context is cancelled.
			call := func(ctx context.Context) error {
				select {
				case <-time.After(tc.callDelay):
					return nil
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			ctx := context.Background()
			if tc.ctxTimeout > 0 {
				c, f := context.WithTimeout(ctx, tc.ctxTimeout)
				defer f()
				ctx = c
			}

			var opts []Option
			if tc.optTimeout > 0 {
				opts = append(opts, WithTimeout(tc.optTimeout))
			}

			got := Execute(ctx, call, opts...)

			if tc.wantTimeout && !errors.Is(got, context.DeadlineExceeded) {
				t.Errorf("Execute(): got error %v, want context.DeadlineExceeded", got)
			}
			if !tc.wantTimeout && got != nil {
				t.Errorf("Execute(): got error %v, want nil", got)
			}
		})
	}
}

func TestExecute_rateLimiting(t *testing.T) {
	testError := errors.New("rate limited")

	// This tests purely focuses on verifying that the limiter is called and
	// that errors are handled correctly.
	testCases := []struct {
		name      string
		limiter   Limiter
		wantErr   error
		wantCalls int
	}{
		{
			name:      "no limiter - call proceeds",
			wantCalls: 1,
		},
		{
			name:      "limiter allows - call proceeds",
			limiter:   &mockLimiter{fn: func(ctx context.Context) error { return nil }},
			wantCalls: 1,
		},
		{
			name:    "limiter blocks - call fails",
			limiter: &mockLimiter{fn: func(ctx context.Context) error { return testError }},
			wantErr: testError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCalls := 0
			call := func(ctx context.Context) error {
				gotCalls++
				return nil
			}

			var opts []Option
			if tc.limiter != nil {
				opts = append(opts, WithLimiter(tc.limiter))
			}

			got := Execute(context.Background(), call, opts...)

			if gotCalls != tc.wantCalls {
				t.Errorf("Execute(): got call count %d, want %d", gotCalls, tc.wantCalls)
			}
			if !errors.Is(got, tc.wantErr) {
				t.Errorf("Execute(): got error %v, want %v", got, tc.wantErr)
			}
		})
	}
}

func TestExecute_contextCancellation(t *testing.T) {
	// Call and retrier for infinite retries.
	testErr := errors.New("test error")
	call := func(ctx context.Context) error {
		return testErr // always fail
	}
	retrier := &mockRetrier{
		fn: func(err error) (time.Duration, bool) {
			return 5 * time.Millisecond, true // always retry
		},
	}

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(10*time.Millisecond, cancel) // cancel after 10ms
	got := Execute(ctx, call, WithRetrier(func() Retrier { return retrier }))

	if !errors.Is(got, context.Canceled) {
		t.Errorf("Execute(): got error %v, want %v", got, context.Canceled)
	}
}

// errorOption implements an Option that always returns an error.
type errorOption struct {
	err error
}

func (e errorOption) Apply(*Options) error {
	return e.err
}

func TestExecuteWithResult(t *testing.T) {
	testErr := errors.New("fail")

	testCases := []struct {
		name       string
		fn         func(context.Context) (int, error)
		wantResult int
		wantErr    error
	}{
		{
			name:       "returns value on success",
			fn:         func(ctx context.Context) (int, error) { return 42, nil },
			wantResult: 42,
		},
		{
			name:       "returns zero value on error",
			fn:         func(ctx context.Context) (int, error) { return 0, testErr },
			wantResult: 0,
			wantErr:    testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ExecuteWithResult(context.Background(), tc.fn)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("got error %v, want %v", err, tc.wantErr)
			}
			if result != tc.wantResult {
				t.Errorf("got result %v, want %v", result, tc.wantResult)
			}
		})
	}
}

func TestExecute_optionErrors(t *testing.T) {
	testErr := errors.New("option error")

	call := func(ctx context.Context) error { return nil }
	got := Execute(context.Background(), call, errorOption{err: testErr})

	if !errors.Is(got, testErr) {
		t.Errorf("Execute(): got error %v, want %v", got, testErr)
	}
}
