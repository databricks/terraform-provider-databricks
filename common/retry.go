package common

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/internal/retrier"
)

// RetryOnTimeout retries f while it returns an SDK inactivity-timeout error.
//
// TODO: Deprecate this function in favor of retrier.Run.
func RetryOnTimeout[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	return retrier.Run(ctx, retryOnErr[*T](isTimeoutError), f)
}

// RetryOn504 retries f while it returns an error wrapping [apierr.ErrDeadlineExceeded].
//
// TODO: Deprecate this function in favor of retrier.Run.
func RetryOn504[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	return retrier.Run(ctx, retryOnErr[*T](is504Error), f)
}

// retryOnErr adapts an error-only predicate to a value-aware retrier factory
// using the transient backoff policy.
func retryOnErr[V any](isRetriable func(error) bool) func() retrier.Retrier[V] {
	return retrier.RetryIf(transientBackoff(), func(_ V, err error) bool {
		return isRetriable(err)
	})
}

// transientBackoff returns a fresh backoff tuned for transient errors:
// start fast, cap quickly.
func transientBackoff() retrier.BackoffPolicy {
	return retrier.BackoffPolicy{Initial: time.Second, Maximum: 30 * time.Second}
}

// TODO: Replace the regex check with type-aware error inspection.
var timeoutRegex = regexp.MustCompile(`request timed out after .* of inactivity`)

func isTimeoutError(err error) bool {
	return err != nil && timeoutRegex.MatchString(err.Error())
}

func is504Error(err error) bool {
	return errors.Is(err, apierr.ErrDeadlineExceeded)
}
