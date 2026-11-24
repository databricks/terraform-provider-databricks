package api

import (
	"math/rand"
	"time"
)

// Retrier defines a retry behavior.
type Retrier interface {
	// IsRetriable returns whether an error is retriable and how long the
	// caller should wait before retrying. Implementation should assume that
	// the given error is never nil.
	IsRetriable(err error) (time.Duration, bool)
}

// RetryOn returns a Retrier that retries based on the isRetriable predicate
// and relies on the given backoff policy to decide how long to wait between
// retries.
//
// Important: the retrier has its own copy of the backoff policy which cannot be
// trivially reset by design. Users who need to reset the backoff policy should
// rather create a new retrier.
func RetryOn(bp BackoffPolicy, isRetriable func(error) bool) Retrier {
	return &retrier{
		isRetriable:   isRetriable,
		backoffPolicy: bp,
	}
}

type retrier struct {
	isRetriable   func(error) bool
	backoffPolicy BackoffPolicy
}

func (r *retrier) IsRetriable(err error) (time.Duration, bool) {
	if !r.isRetriable(err) {
		return 0, false
	}
	return r.backoffPolicy.Delay(), true
}

// BackoffPolicy implements an exponential backoff policy. The delay betwen
// retries is randomly computed between 0 and the "exponential delay" as
// recommended in [Exponential Backoff And Jitter]. The retry delay start from
// Initial and grows exponentially by Factor at every retry. The maximum
// retry delay is capped by Maximum.
//
// There is no parameter to limit the number of retries. This is intended as
// such logic should be implemented upstream (e.g. in a Retrier).
//
// Important: BackoffPolicies cannot conveniently be reset. This is intentional
// and client should rather create a new BackoffPolicy instead of resetting an
// existing one.
//
// [Exponential Backoff And Jitter]: https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
type BackoffPolicy struct {
	// Initial delay to be used for the first retry; defaults to 1 second.
	Initial time.Duration

	// Maximum value for the retry delay; defaults to 60 seconds.
	Maximum time.Duration

	// Factor by which the delay is multiplied after each retry. The value
	// must be greater or equal to 1. If not, it defaults to 2.
	Factor float64

	// Current delay before the next retry.
	current time.Duration

	// Whether the backoff policy has been initialized.
	initialized bool

	// Function to generate a random number between 0 and n. This is exposed for
	// testing purposes.
	int63n func(n int64) int64
}

func (bp *BackoffPolicy) Delay() time.Duration {
	if !bp.initialized {
		bp.setDefaults()
	}

	// Random duration in the range [0, bp.current].
	d := time.Duration(bp.int63n(int64(bp.current) + 1))

	// Grow delay for the next call.
	bp.current = time.Duration(float64(bp.current) * bp.Factor)
	if bp.current > bp.Maximum {
		bp.current = bp.Maximum
	}

	return d
}

func (bp *BackoffPolicy) setDefaults() {
	if bp.Initial == 0 {
		bp.Initial = 1 * time.Second
	}
	if bp.Maximum == 0 {
		bp.Maximum = 60 * time.Second
	}
	if bp.Initial > bp.Maximum {
		bp.Initial = bp.Maximum // Initial cannot be greater than Maximum
	}
	if bp.Factor < 1 {
		bp.Factor = 2
	}
	if bp.int63n == nil {
		bp.int63n = rand.Int63n
	}
	bp.current = bp.Initial
	bp.initialized = true
}
