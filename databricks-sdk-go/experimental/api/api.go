// Package api provides utilities to make API calls against the Databricks API.
//
// This package draws inspiration from the AWS and GCP SDKs.
package api

import (
	"context"
	"time"
)

// Call represents a call to a Databricks API.
type Call func(context.Context) error

// Execute makes a call to a Databricks API using the given options.
func Execute(ctx context.Context, call Call, opts ...Option) error {
	options := Options{}
	for _, opt := range opts {
		if err := opt.Apply(&options); err != nil {
			return err
		}
	}
	return execute(ctx, call, options, sleep)
}

// ExecuteWithResult returns the result of calling op with the given options.
// It is a convenience wrapper around Execute for operations that return a
// value. In case of error, the zero value of T is returned.
func ExecuteWithResult[T any](ctx context.Context, op func(context.Context) (T, error), opts ...Option) (T, error) {
	var result T
	err := Execute(ctx, func(ctx context.Context) error {
		var err error
		result, err = op(ctx)
		return err
	}, opts...)
	if err != nil {
		var zero T // guarantee zero value on error
		return zero, err
	}
	return result, nil
}

// sleep sleeps for the given duration. It is mostly equivalent to time.Sleep,
// but can be interrupted by ctx.Done() if the context completes before the
// duration elapses.
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

// sleeper is a convenience type for readability.
type sleeper func(ctx context.Context, d time.Duration) error

// execute is the actual implementation of Execute. Its purpose is to ease
// testing by providing a convenient way to mock the sleeping logic.
func execute(ctx context.Context, apiCall Call, opts Options, sleep sleeper) error {
	// Optionally update the context with the timeout. If the context already
	// has a deadline, that deadline is updated to the minimum of the context's
	// deadline and the timeout.
	if opts.timeout != 0 {
		c, f := context.WithTimeout(ctx, opts.timeout)
		defer f()
		ctx = c
	}

	// Get a new retrier for this specific execution. This is instantiated
	// lazilly if and if the first call execution returns an error.
	var retrier Retrier

	for {
		if opts.rateLimiter != nil { // no limiter == no wait
			if err := opts.rateLimiter.Wait(ctx); err != nil {
				return err
			}
		}

		err := apiCall(ctx)
		if err == nil {
			return nil // nothing to retry
		}

		if retrier == nil {
			if opts.retrier != nil {
				retrier = opts.retrier() // lazily instantiate the retrier
			}
			if retrier == nil {
				return err // no retrier == no retry
			}
		}
		delay, ok := retrier.IsRetriable(err)
		if !ok {
			return err // not retriable
		}
		if err := sleep(ctx, delay); err != nil {
			return err
		}
	}
}
