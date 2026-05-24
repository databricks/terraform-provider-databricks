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

// transientBackoff returns a fresh backoff tuned for transient errors:
// start fast, cap quickly.
func transientBackoff() retrier.BackoffPolicy {
	return retrier.BackoffPolicy{Initial: time.Second, Maximum: 30 * time.Second}
}

// RetryOnTimeout retries f while it returns an SDK inactivity-timeout error.
func RetryOnTimeout[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	return retrier.Run(ctx,
		retrier.RetryIf(transientBackoff(), func(_ *T, err error) bool {
			if err == nil || !timeoutRegex.MatchString(err.Error()) {
				return false
			}
			logger.Debugf(ctx, "Retrying due to timeout: %s", err.Error())
			return true
		}),
		f,
	)
}

// RetryOn504 retries f while it returns an error wrapping [apierr.ErrDeadlineExceeded].
func RetryOn504[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	return retrier.Run(ctx,
		retrier.RetryIf(transientBackoff(), func(_ *T, err error) bool {
			if !errors.Is(err, apierr.ErrDeadlineExceeded) {
				return false
			}
			logger.Debugf(ctx, "Retrying on error 504")
			return true
		}),
		f,
	)
}
