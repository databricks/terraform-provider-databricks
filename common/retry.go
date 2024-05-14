package common

import (
	"context"
	"log"
	"regexp"

	"github.com/databricks/databricks-sdk-go/retries"
)

var timeoutRegex = regexp.MustCompile(`request timed out after .* of inactivity`)

func RetryOnTimeout[T any](ctx context.Context, f func(context.Context) (*T, error)) (*T, error) {
	r := retries.New[T](retries.WithRetryFunc(func(err error) bool {
		msg := err.Error()
		isTimeout := timeoutRegex.MatchString(msg)
		if isTimeout {
			log.Printf("[DEBUG] Retrying due to timeout: %s", msg)
		}
		return isTimeout
	}))
	return r.Run(ctx, func(ctx context.Context) (*T, error) {
		return f(ctx)
	})
}
