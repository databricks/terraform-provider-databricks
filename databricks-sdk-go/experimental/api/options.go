package api

import (
	"time"
)

// Option is an option used by Execute to control the behavior of an API call.
// An Option essentially act as a convenient way to configure Options.
type Option interface {
	Apply(*Options) error
}

type optionFn func(*Options)

func (fn optionFn) Apply(opts *Options) error {
	fn(opts)
	return nil
}

// WithRetrier configures to use the given Retrier provider. If no retrier is
// provided, the API call is not retried.
//
// The provider function must be thread-safe.
func WithRetrier(provider func() Retrier) Option {
	return optionFn(func(o *Options) {
		o.retrier = provider
	})
}

// WithTimeout is a convenience option to set the timeout duration in the Call
// context. If the context already has a deadline, that deadline is updated to
// the minimum of the context's deadline and the timeout.
//
// The timeout covers the whole Call execution; it is not a timeout of each
// intermediary API call.
func WithTimeout(t time.Duration) Option {
	return optionFn(func(o *Options) {
		o.timeout = t
	})
}

// WithLimiter configures to use the given Limiter provider.
//
// The limiter is used to potential rate limit the API call. If no limiter is
// provided, the API call is not rate limited.
func WithLimiter(l Limiter) Option {
	return optionFn(func(o *Options) {
		o.rateLimiter = l
	})
}

// Options to control the behavior of an API call.
type Options struct {
	// Provides a new Retrier to be used to execute a Call. The function
	// is called for each Call and must be thread-safe. The retrier must be
	// fresh within the context of an Execute call (e.g. no need to reset a
	// BackoffStrategy).
	retrier func() Retrier

	rateLimiter Limiter

	timeout time.Duration
}
