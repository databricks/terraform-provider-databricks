package api

import "context"

// Limiter is anything that can wait. It is typically used to implement
// client-side rate limiting. Implementation of this interface must be
// thread-safe.
type Limiter interface {
	Wait(context.Context) error
}
