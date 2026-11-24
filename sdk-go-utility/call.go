package sdkgoutility

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/databricks/sdk-go/core/ops"
	"github.com/databricks/sdk-go/options/call"
)

// Call runs method, which wraps one method of a modular SDK client, as the
// classic SDK client runs a request: it retries by retrier, waits on limiter
// before every attempt and returns the error that classic makes of a failure.
//
// The method runs once, with retrier as its call.WithRetrier option and
// limiter as its call.WithLimiter option. A nil retrier or limiter leaves that
// option out. When the method fails, Call returns the zero value of T and
// classic(err), or err itself when classic is nil.
//
// One failure is reported as the classic client reports it, not as the modular
// SDK does: when the context ends while the SDK waits before a retry, err is
// the error of the last attempt, not that of the context.
//
// Call leaves to the caller which errors retrier retries, how fast limiter
// lets requests through and what classic converts, as well as:
//   - a new retrier for every call, because a retrier keeps the backoff of the
//     call it serves;
//   - the timeout, which stays that of the modular client;
//   - a method that returns an iterator, whose requests run after the method
//     has returned, outside Call.
func Call[T any](
	ctx context.Context,
	method func(context.Context, ...call.Option) (T, error),
	retrier ops.Retrier,
	limiter ops.Limiter,
	classic func(error) error,
) (T, error) {
	watch := &retryWatch{retrier: retrier, limiter: limiter}
	var opts []call.Option
	if retrier != nil {
		opts = append(opts, call.WithRetrier(func() ops.Retrier { return watch }))
	}
	if limiter != nil {
		opts = append(opts, call.WithLimiter(watch))
	}
	result, err := method(ctx, opts...)
	if err == nil {
		return result, nil
	}
	err = watch.failure(err)
	if classic != nil {
		err = classic(err)
	}
	var zero T
	return zero, err
}

// retryWatch stands between the SDK and the retrier and limiter of one Call.
// The SDK asks the retrier after every failed attempt and the limiter before
// every attempt, which tells whether it was waiting to retry when it gave up.
type retryWatch struct {
	retrier ops.Retrier
	limiter ops.Limiter
	// retried is the error of the attempt that the SDK is waiting to retry, and
	// nil outside that wait. It is atomic because ops requires a Limiter to be
	// safe for concurrent use.
	retried atomic.Pointer[error]
}

func (w *retryWatch) IsRetriable(err error) (time.Duration, bool) {
	delay, retry := w.retrier.IsRetriable(err)
	if retry {
		w.retried.Store(&err)
	} else {
		w.retried.Store(nil)
	}
	return delay, retry
}

func (w *retryWatch) Wait(ctx context.Context) error {
	// The SDK asks for the next attempt, so its wait to retry is over.
	w.retried.Store(nil)
	return w.limiter.Wait(ctx)
}

// failure returns the error that the classic client returns where the SDK
// returned err. The SDK returns the error of the context when the context ends
// during a wait to retry; the classic client returns the error it would have
// retried.
func (w *retryWatch) failure(err error) error {
	retried := w.retried.Load()
	contextEnded := errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
	if retried != nil && contextEnded {
		return *retried
	}
	return err
}
