package retrier

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBackoffPolicy_Next_initialization(t *testing.T) {
	testCases := []struct {
		name   string
		bp     BackoffPolicy
		wantBP BackoffPolicy
	}{
		{
			name: "default",
			bp:   BackoffPolicy{},
			wantBP: BackoffPolicy{
				Initial: 10 * time.Second,
				Maximum: 5 * time.Minute,
				Factor:  2,
			},
		},
		{
			name: "custom initial smaller than maximum",
			bp: BackoffPolicy{
				Initial: 100 * time.Millisecond,
			},
			wantBP: BackoffPolicy{
				Initial: 100 * time.Millisecond,
				Maximum: 5 * time.Minute,
				Factor:  2,
			},
		},
		{
			name: "custom initial greater than maximum",
			bp: BackoffPolicy{
				Initial: 10 * time.Second,
				Maximum: 1 * time.Second,
			},
			wantBP: BackoffPolicy{
				Initial: 1 * time.Second,
				Maximum: 1 * time.Second,
				Factor:  2,
			},
		},
		{
			name: "custom factor less than 1",
			bp: BackoffPolicy{
				Factor: 0.5,
			},
			wantBP: BackoffPolicy{
				Initial: 10 * time.Second,
				Maximum: 5 * time.Minute,
				Factor:  2,
			},
		},
		{
			name: "custom factor greater than 1",
			bp: BackoffPolicy{
				Factor: 1.5,
			},
			wantBP: BackoffPolicy{
				Initial: 10 * time.Second,
				Maximum: 5 * time.Minute,
				Factor:  1.5,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.bp.Next() // initializes defaults

			if diff := cmp.Diff(tc.bp, tc.wantBP, cmpopts.IgnoreUnexported(BackoffPolicy{})); diff != "" {
				t.Errorf("unexpected BackoffPolicy (-got +want):\n%s", diff)
			}
		})
	}
}

func TestBackoffPolicy_Next_exponential(t *testing.T) {
	bp := BackoffPolicy{
		Initial: 100 * time.Millisecond,
		Maximum: 10 * time.Second,
		Factor:  2.0,
	}

	wantDelays := []time.Duration{
		100 * time.Millisecond,
		200 * time.Millisecond,
		400 * time.Millisecond,
		800 * time.Millisecond,
		1600 * time.Millisecond,
		3200 * time.Millisecond,
		6400 * time.Millisecond,
		10000 * time.Millisecond, // capped by Maximum
		10000 * time.Millisecond, // stays capped
	}

	for i, want := range wantDelays {
		if got := bp.Next(); got != want {
			t.Errorf("Next() call %d = %v, want %v", i+1, got, want)
		}
	}
}

// mockRetrier implements Retrier[V] for use in tests.
type mockRetrier[V any] struct {
	fn func(V, error) (time.Duration, bool)
}

func (m *mockRetrier[V]) IsRetriable(val V, err error) (time.Duration, bool) {
	return m.fn(val, err)
}

// noSleep is a sleeper that returns immediately so retry tests do not wait.
func noSleep(ctx context.Context, _ time.Duration) error {
	return nil
}

func TestRun_retries(t *testing.T) {
	retriableErr := errors.New("retriable")
	nonRetriableErr := errors.New("non-retriable")

	// Factory that retries only on retriableErr; backoff delay is irrelevant
	// because tests inject a noSleep sleeper.
	retryOnRetriable := func() Retrier[int] {
		return &mockRetrier[int]{
			fn: func(_ int, err error) (time.Duration, bool) {
				return 0, errors.Is(err, retriableErr)
			},
		}
	}

	testCases := []struct {
		name      string
		callErrs  []error
		newR      func() Retrier[int]
		wantErr   error
		wantCalls int
	}{
		{
			name:      "nil factory - fail immediately",
			callErrs:  []error{retriableErr},
			newR:      nil,
			wantErr:   retriableErr,
			wantCalls: 1,
		},
		{
			name:      "factory returns nil - fail immediately",
			callErrs:  []error{retriableErr},
			newR:      func() Retrier[int] { return nil },
			wantErr:   retriableErr,
			wantCalls: 1,
		},
		{
			name:      "non-retriable error - fail immediately",
			callErrs:  []error{nonRetriableErr},
			newR:      retryOnRetriable,
			wantErr:   nonRetriableErr,
			wantCalls: 1,
		},
		{
			name:      "retriable error - retry once then succeed",
			callErrs:  []error{retriableErr, nil},
			newR:      retryOnRetriable,
			wantCalls: 2,
		},
		{
			name:      "retriable error - retry multiple times then succeed",
			callErrs:  []error{retriableErr, retriableErr, retriableErr, nil},
			newR:      retryOnRetriable,
			wantCalls: 4,
		},
		{
			name:      "retriable then fail with non-retriable",
			callErrs:  []error{retriableErr, nonRetriableErr},
			newR:      retryOnRetriable,
			wantErr:   nonRetriableErr,
			wantCalls: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCalls := 0
			fn := func(ctx context.Context) (int, error) {
				err := tc.callErrs[gotCalls]
				gotCalls++
				return gotCalls, err
			}

			_, gotErr := run(context.Background(), tc.newR, fn, noSleep)

			if gotCalls != tc.wantCalls {
				t.Errorf("call count = %d, want %d", gotCalls, tc.wantCalls)
			}
			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("err = %v, want %v", gotErr, tc.wantErr)
			}
		})
	}
}

func TestRun_valueAware(t *testing.T) {
	// State-driven: retry while the value is < 3. Verifies that the retrier
	// sees the value (not just the error) and that Run returns the final
	// terminal value.
	newR := func() Retrier[int] {
		return &mockRetrier[int]{
			fn: func(val int, _ error) (time.Duration, bool) {
				return 0, val < 3
			},
		}
	}

	calls := 0
	fn := func(ctx context.Context) (int, error) {
		calls++
		return calls, nil
	}

	gotVal, gotErr := run(context.Background(), newR, fn, noSleep)

	if gotErr != nil {
		t.Errorf("unexpected error: %v", gotErr)
	}
	if gotVal != 3 {
		t.Errorf("got val %d, want 3", gotVal)
	}
	if calls != 3 {
		t.Errorf("call count = %d, want 3", calls)
	}
}

func TestRun_contextCancellation(t *testing.T) {
	// When the sleeper reports a cancelled context, Run surfaces that error
	// and returns the most recent attempt's value. The kind of cancellation
	// (cancel vs deadline) is passed through unchanged.
	callErr := errors.New("call error")

	newR := func() Retrier[int] {
		return &mockRetrier[int]{
			fn: func(_ int, _ error) (time.Duration, bool) {
				return 5 * time.Millisecond, true // always retry
			},
		}
	}

	testCases := []struct {
		name    string
		sleeper sleeper
		wantErr error
	}{
		{
			name:    "context canceled during sleep",
			sleeper: func(_ context.Context, _ time.Duration) error { return context.Canceled },
			wantErr: context.Canceled,
		},
		{
			name:    "deadline exceeded during sleep",
			sleeper: func(_ context.Context, _ time.Duration) error { return context.DeadlineExceeded },
			wantErr: context.DeadlineExceeded,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			fn := func(ctx context.Context) (int, error) {
				calls++
				return calls, callErr
			}

			gotVal, gotErr := run(context.Background(), newR, fn, tc.sleeper)

			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("err = %v, want %v", gotErr, tc.wantErr)
			}
			if gotVal != 1 {
				t.Errorf("val = %d, want 1 (last fn return before cancellation)", gotVal)
			}
			if calls != 1 {
				t.Errorf("call count = %d, want 1", calls)
			}
		})
	}
}

func TestRunErr_retries(t *testing.T) {
	// RunErr is a thin adapter over Run; these cases confirm the adapter
	// wires the error-only predicate through correctly and returns only
	// the error.
	retriableErr := errors.New("retriable")
	nonRetriableErr := errors.New("non-retriable")

	newR := RetryIfErr(
		BackoffPolicy{Initial: time.Millisecond, Maximum: time.Millisecond, Factor: 1},
		func(err error) bool { return errors.Is(err, retriableErr) },
	)

	testCases := []struct {
		name      string
		callErrs  []error
		wantErr   error
		wantCalls int
	}{
		{
			name:      "success on first call",
			callErrs:  []error{nil},
			wantCalls: 1,
		},
		{
			name:      "retry then succeed",
			callErrs:  []error{retriableErr, retriableErr, nil},
			wantCalls: 3,
		},
		{
			name:      "non-retriable halts",
			callErrs:  []error{nonRetriableErr},
			wantErr:   nonRetriableErr,
			wantCalls: 1,
		},
		{
			name:      "retriable then non-retriable halts",
			callErrs:  []error{retriableErr, nonRetriableErr},
			wantErr:   nonRetriableErr,
			wantCalls: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			fn := func(ctx context.Context) error {
				err := tc.callErrs[calls]
				calls++
				return err
			}

			gotErr := RunErr(context.Background(), newR, fn)

			if !errors.Is(gotErr, tc.wantErr) {
				t.Errorf("err = %v, want %v", gotErr, tc.wantErr)
			}
			if calls != tc.wantCalls {
				t.Errorf("call count = %d, want %d", calls, tc.wantCalls)
			}
		})
	}
}

func TestRetryIf_constructedRetrierUsesBackoff(t *testing.T) {
	// RetryIf must wire both the predicate and the backoff into the returned
	// retrier. With Factor 1, the delay stays at Initial on every call.
	newR := RetryIf(
		BackoffPolicy{Initial: 7 * time.Millisecond, Maximum: 7 * time.Millisecond, Factor: 1},
		func(_ int, err error) bool { return err != nil },
	)

	r := newR()
	for i := range 3 {
		attempt := i + 1
		delay, retry := r.IsRetriable(0, errors.New("boom"))
		if !retry {
			t.Fatalf("attempt %d: retry = false, want true", attempt)
		}
		if delay != 7*time.Millisecond {
			t.Errorf("attempt %d: delay = %v, want 7ms", attempt, delay)
		}
	}

	if _, retry := r.IsRetriable(0, nil); retry {
		t.Errorf("on nil err: retry = true, want false")
	}
}

func TestRetryIf_factoryProducesIndependentInstances(t *testing.T) {
	// Each invocation of the factory must produce a retrier with its own
	// backoff state; advancing one must not advance another.
	newR := RetryIf(
		BackoffPolicy{Initial: time.Millisecond, Maximum: time.Second, Factor: 2},
		func(_ int, _ error) bool { return true },
	)

	boom := errors.New("boom")

	r1 := newR()
	r1.IsRetriable(0, boom)
	r1.IsRetriable(0, boom)
	d1, _ := r1.IsRetriable(0, boom) // r1 is now at attempt 4

	r2 := newR()
	d2, _ := r2.IsRetriable(0, boom) // r2's first attempt

	if d1 <= d2 {
		t.Errorf("expected r1's delay (%v) to exceed r2's first delay (%v); shared state suspected", d1, d2)
	}
}

func TestRetryIfErr_wraps(t *testing.T) {
	// RetryIfErr produces a Retrier[struct{}] that forwards to the
	// error-only predicate, ignoring the value argument.
	newR := RetryIfErr(
		BackoffPolicy{Initial: time.Millisecond, Maximum: time.Millisecond, Factor: 1},
		func(err error) bool { return errors.Is(err, context.Canceled) },
	)

	testCases := []struct {
		name      string
		err       error
		wantRetry bool
	}{
		{name: "matching error retries", err: context.Canceled, wantRetry: true},
		{name: "non-matching error halts", err: errors.New("other"), wantRetry: false},
		{name: "nil error halts", err: nil, wantRetry: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := newR()
			_, retry := r.IsRetriable(struct{}{}, tc.err)
			if retry != tc.wantRetry {
				t.Errorf("retry = %v, want %v", retry, tc.wantRetry)
			}
		})
	}
}

func TestSleep(t *testing.T) {
	// sleep should return ctx.Err() promptly when the context is done and
	// nil when the duration elapses normally.
	testCases := []struct {
		name     string
		ctxSetup func() (context.Context, context.CancelFunc)
		duration time.Duration
		wantErr  error
	}{
		{
			name: "pre-cancelled context returns immediately",
			ctxSetup: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx, func() {}
			},
			duration: time.Hour,
			wantErr:  context.Canceled,
		},
		{
			name: "deadline exceeded returns immediately",
			ctxSetup: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 5*time.Millisecond)
			},
			duration: time.Hour,
			wantErr:  context.DeadlineExceeded,
		},
		{
			name: "duration elapses normally",
			ctxSetup: func() (context.Context, context.CancelFunc) {
				return context.Background(), func() {}
			},
			duration: 5 * time.Millisecond,
			wantErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := tc.ctxSetup()
			defer cancel()

			start := time.Now()
			err := sleep(ctx, tc.duration)
			elapsed := time.Since(start)

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("err = %v, want %v", err, tc.wantErr)
			}
			if elapsed > 100*time.Millisecond {
				t.Errorf("sleep took %v; expected < 100ms", elapsed)
			}
		})
	}
}
