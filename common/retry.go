package common

import (
	"context"
	"errors"
	"regexp"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
)

var timeoutRegex = regexp.MustCompile(`request timed out after .* of inactivity`)

func RetryOnTimeout[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	r := retries.New[T](retries.WithRetryFunc(func(err error) bool {
		msg := err.Error()
		isTimeout := timeoutRegex.MatchString(msg)
		if isTimeout {
			logger.Debugf(ctx, "Retrying due to timeout: %s", msg)
		}
		return isTimeout
	}))
	return r.Run(ctx, func(ctx context.Context) (*T, error) {
		return f(ctx)
	})
}

// RetryOn504 returns a [retries.Retrier] that calls the given method
// until it either succeeds or returns an error that is different from
// [apierr.ErrDeadlineExceeded].
func RetryOn504[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	r := retries.New[T](retries.WithTimeout(-1), retries.WithRetryFunc(func(err error) bool {
		if !errors.Is(err, apierr.ErrDeadlineExceeded) {
			return false
		}
		logger.Debugf(ctx, "Retrying on error 504")
		return true
	}))
	return r.Run(ctx, func(ctx context.Context) (*T, error) {
		return f(ctx)
	})
}
