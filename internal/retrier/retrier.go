// Package retrier provides a small, value-aware retry loop for waiting on
// asynchronous state transitions (e.g. cluster startup, workspace creation)
// and for retrying transient errors that the Databricks SDK does not absorb
// itself. Retriers are constructed fresh per Run so independent invocations
// have independent backoff state and can run concurrently.
package retrier

import (
	"context"
	"time"
)

// Retrier decides whether to retry an operation based on the outcome of
// the last attempt. IsRetriable is called after every attempt with the
// value and error produced by that attempt. It returns the delay to wait
// before the next attempt and whether to retry at all.
//
// Passing both value and error lets callers express either of the two
// dominant retry policies in this codebase:
//
//   - error-driven: retry transient errors (e.g. 504, timeout); halt on
//     terminal errors; halt on success.
//   - state-driven: retry while the value is in a non-terminal state
//     (e.g. a cluster in PENDING); halt when the value reaches a
//     terminal state; halt on any error.
//
// State-driven polling is the majority of retry sites in this provider,
// so the value is always passed in. Error-driven retriers ignore it.
type Retrier[V any] interface {
	IsRetriable(V, error) (time.Duration, bool)
}

// RetrierErr is the error-only specialisation of Retrier, for retry policies
// that do not inspect the operation's result value (e.g. retrying on transient
// errors, polling until a resource is deleted).
type RetrierErr = Retrier[struct{}]

// RetryIf returns a retrier factory suitable for passing to Run. The
// constructed retrier calls isRetriable after every attempt and, when it
// returns true, sleeps for the duration produced by a fresh copy of bp before
// the next attempt. bp is captured by value, so independent Run invocations
// get independent backoff state and can run concurrently.
func RetryIf[V any](bp BackoffPolicy, isRetriable func(V, error) bool) func() Retrier[V] {
	return func() Retrier[V] {
		return &retrier[V]{
			backoffPolicy: bp,
			isRetriable:   isRetriable,
		}
	}
}

// RetryIfErr is the error-only analog of RetryIf, suitable for passing to
// RunErr. The isRetriable predicate sees only the error.
func RetryIfErr(bp BackoffPolicy, isRetriable func(error) bool) func() RetrierErr {
	return RetryIf(bp, func(_ struct{}, err error) bool {
		return isRetriable(err)
	})
}

type retrier[V any] struct {
	backoffPolicy BackoffPolicy
	isRetriable   func(V, error) bool
}

func (r *retrier[V]) IsRetriable(val V, err error) (time.Duration, bool) {
	if !r.isRetriable(val, err) {
		return 0, false
	}
	return r.backoffPolicy.Next(), true
}

// Run executes fn, retrying according to the retrier constructed by
// newR. newR is invoked once per Run call so each invocation has its
// own independent retrier state; this makes Run safe to use from
// multiple goroutines with the same newR. If newR is nil, or if it
// returns nil, Run performs exactly one attempt with no retry logic.
// It returns the value and error of the final attempt.
//
// Run honours context cancellation between attempts and during sleeps;
// if ctx completes, Run returns the last attempt's value and ctx.Err().
func Run[V any](ctx context.Context, newR func() Retrier[V], fn func(context.Context) (V, error)) (V, error) {
	return run(ctx, newR, fn, sleep)
}

// RunErr is the error-only analog of Run. It applies the same lifecycle
// and concurrency contract as Run; see Run's documentation for the
// details on newR, single-shot semantics, and ctx cancellation.
func RunErr(ctx context.Context, newR func() RetrierErr, fn func(context.Context) error) error {
	_, err := Run(ctx, newR, func(ctx context.Context) (struct{}, error) {
		return struct{}{}, fn(ctx)
	})
	return err
}

// sleep sleeps for d, returning early with ctx.Err() if the context
// completes first.
func sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		t.Stop()
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// sleeper is the dependency injected into run so tests can substitute
// a deterministic sleep.
type sleeper func(ctx context.Context, d time.Duration) error

// run is the implementation of Run. It takes the sleeper as an argument
// so tests can substitute a deterministic sleep.
func run[V any](ctx context.Context, newR func() Retrier[V], fn func(context.Context) (V, error), sleep sleeper) (V, error) {
	// r is constructed lazily so single-shot Runs skip allocation and any
	// operation-start state in the retrier anchors at first need.
	var r Retrier[V]

	for {
		val, err := fn(ctx)

		if r == nil {
			if newR != nil {
				r = newR()
			}
			if r == nil {
				return val, err
			}
		}

		delay, retry := r.IsRetriable(val, err)
		if !retry {
			return val, err
		}

		if sleepErr := sleep(ctx, delay); sleepErr != nil {
			return val, sleepErr
		}
	}
}

// BackoffPolicy implements a deterministic exponential backoff. The retry
// delay starts at Initial and grows by Factor at every step, capped at
// Maximum. There is no jitter: the Databricks SDK already retries transient
// errors on its own, and state-polling waiters in this package do not need
// cross-client decorrelation.
//
// BackoffPolicies are stateful and cannot be reset. Construct a new one per
// Run invocation rather than reusing.
type BackoffPolicy struct {
	// Initial delay before the first retry; defaults to 10 seconds.
	Initial time.Duration

	// Maximum delay between retries; defaults to 5 minutes.
	Maximum time.Duration

	// Factor by which the delay grows after each retry. Must be >= 1;
	// defaults to 2 if zero or negative.
	Factor float64

	// current is the delay that will be returned by the next call to Next.
	current time.Duration

	// initialized tracks whether setDefaults has run.
	initialized bool
}

func (bp *BackoffPolicy) Next() time.Duration {
	if !bp.initialized {
		bp.setDefaults()
	}
	d := bp.current
	bp.current = min(time.Duration(float64(bp.current)*bp.Factor), bp.Maximum)
	return d
}

func (bp *BackoffPolicy) setDefaults() {
	if bp.Initial == 0 {
		bp.Initial = 10 * time.Second
	}
	if bp.Maximum == 0 {
		bp.Maximum = 5 * time.Minute
	}
	if bp.Initial > bp.Maximum {
		bp.Initial = bp.Maximum // Initial cannot be greater than Maximum
	}
	if bp.Factor < 1 {
		bp.Factor = 2
	}
	bp.current = bp.Initial
	bp.initialized = true
}
