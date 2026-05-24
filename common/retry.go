package common

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/terraform-provider-databricks/internal/retrier"
)

var timeoutRegex = regexp.MustCompile(`request timed out after .* of inactivity`)

// transientErrorBackoff is the backoff used by both helpers below. Transient
// errors typically resolve in a few seconds, so we start fast and cap quickly
// rather than using the retrier package defaults of 10s/5m.
var transientErrorBackoff = retrier.BackoffPolicy{
	Initial: 1 * time.Second,
	Maximum: 30 * time.Second,
	Factor:  2,
}

// RetryOnTimeout retries f while it returns an error whose message matches
// the SDK's "request timed out after ... of inactivity" pattern. Any other
// error, or a successful call, halts.
func RetryOnTimeout[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	return retrier.Run(ctx,
		retrier.RetryIf(transientErrorBackoff, func(_ *T, err error) bool {
			if err == nil {
				return false
			}
			isTimeout := timeoutRegex.MatchString(err.Error())
			if isTimeout {
				logger.Debugf(ctx, "Retrying due to timeout: %s", err.Error())
			}
			return isTimeout
		}),
		f,
	)
}

// RetryOn504 retries f while it returns an error wrapping
// [apierr.ErrDeadlineExceeded] (HTTP 504). Any other error, or a successful
// call, halts.
func RetryOn504[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	return retrier.Run(ctx,
		retrier.RetryIf(transientErrorBackoff, func(_ *T, err error) bool {
			if err == nil {
				return false
			}
			if !errors.Is(err, apierr.ErrDeadlineExceeded) {
				return false
			}
			logger.Debugf(ctx, "Retrying on error 504")
			return true
		}),
		f,
	)
}
