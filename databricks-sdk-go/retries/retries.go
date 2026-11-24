package retries

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
)

// Deprecated: use return types from non-*AndWait methods
type Info[T any] struct {
	Info    *T
	Timeout time.Duration
}

// Deprecated: use return types from non-*AndWait methods
type Option[T any] func(*Info[T])

// Deprecated: use return types from non-*AndWait methods
func Timeout[T any](dur time.Duration) Option[T] {
	return func(i *Info[T]) {
		i.Timeout = dur
	}
}

// Deprecated: use return types from non-*AndWait methods
func OnPoll[T any](callback func(*T)) Option[T] {
	return func(i *Info[T]) {
		if i.Info == nil {
			return
		}
		callback(i.Info)
	}
}

type Err struct {
	Err  error
	Halt bool
}

func (e *Err) Error() string {
	return e.Err.Error()
}

func (e *Err) Unwrap() error {
	return e.Err
}

func Halt(err error) *Err {
	return &Err{err, true}
}

func Continue(err error) *Err {
	return &Err{err, false}
}

func Continues(msg string) *Err {
	return Continue(fmt.Errorf("%s", msg))
}

func Continuef(format string, err error, args ...interface{}) *Err {
	wrapped := fmt.Errorf(format, append([]interface{}{err}, args...))
	return Continue(wrapped)
}

var maxWait = 10 * time.Second
var minJitter = 50 * time.Millisecond
var maxJitter = 750 * time.Millisecond

func backoff(attempt int) time.Duration {
	wait := time.Duration(attempt) * time.Second
	if wait > maxWait {
		wait = maxWait
	}
	// add some random jitter
	rand.Seed(time.Now().UnixNano())
	jitter := rand.Intn(int(maxJitter)-int(minJitter)+1) + int(minJitter)
	wait += time.Duration(jitter)
	return wait
}

type ErrTimedOut struct {
	err error
}

func (et *ErrTimedOut) Error() string {
	return fmt.Sprintf("timed out: %s", et.err)
}

func (et *ErrTimedOut) Unwrap() error {
	return et.err
}

// RetryOption is a function that sets part of the retry configuration for a retrier.
type RetryOption func(*RetryConfig)

// RetryConfig is the configuration for a retrier.
type RetryConfig struct {
	timeout     time.Duration
	shouldRetry func(error) bool
	backoff     func(int) time.Duration
}

func (r RetryConfig) Timeout() time.Duration {
	if r.timeout == 0 {
		return 20 * time.Minute
	}
	return r.timeout
}

func (r RetryConfig) Backoff(attempt int) time.Duration {
	if r.backoff == nil {
		return backoff(attempt)
	}
	return r.backoff(attempt)
}

func (r RetryConfig) ShouldRetry(err error) bool {
	if r.shouldRetry == nil {
		return err != nil
	}
	return r.shouldRetry(err)
}

// WithTimeout sets the timeout for the retrier.
func WithTimeout(timeout time.Duration) RetryOption {
	return func(rc *RetryConfig) {
		rc.timeout = timeout
	}
}

// OnErrors sets the errors that should be retried.
func OnErrors(on ...error) RetryOption {
	return func(rc *RetryConfig) {
		rc.shouldRetry = func(err error) bool {
			for _, e := range on {
				if errors.Is(err, e) {
					return true
				}
			}
			return false
		}
	}
}

// WithRetryFunc sets the function that determines whether an error should halt the retrier.
// If the function returns true, the retrier will continue. If it returns false, the retrier will halt.
func WithRetryFunc(halt func(error) bool) RetryOption {
	return func(rc *RetryConfig) {
		rc.shouldRetry = halt
	}
}

// Retrier is a struct that can retry an operation until it succeeds or the timeout is reached.
// The empty struct indicates that the retrier should run for 20 minutes and retry on any non-nil error.
// The type parameter is the return type of the Run() method. When using the Wait() method, this can be struct{}.
//
// Example:
//
//	r := retries.New[struct{}](retries.WithTimeout(5 * time.Minute), retries.OnErrors(apierr.ErrResourceConflict))
//	err := r.Wait(ctx, func(ctx context.Context) error {
//		return a.Workspaces.Delete(ctx, provisioning.DeleteWorkspaceRequest{
//			WorkspaceId: workspace.WorkspaceId,
//		})
//	})
type Retrier[T any] struct {
	config RetryConfig
}

// New creates a new retrier with the given configuration.
// If no timeout is specified, the default is 20 minutes. If the timeout is negative, the retrier will run indefinitely.
// If no retry function is specified, the default is to retry on all errors.
func New[T any](configOpts ...RetryOption) Retrier[T] {
	config := RetryConfig{}
	for _, opt := range configOpts {
		opt(&config)
	}
	return Retrier[T]{config}
}

// Wait runs the given function until it succeeds or the timeout is reached. On
// success, it returns nil. On timeout, it returns an error wrapping the last
// error returned by the function.
func (r Retrier[T]) Wait(ctx context.Context, fn func(ctx context.Context) error) error {
	_, err := r.Run(ctx, func(ctx context.Context) (*T, error) {
		return nil, fn(ctx)
	})
	return err
}

// Run runs the given function until it succeeds or the timeout is reached, returning the result.
// On timeout, it returns an error wrapping the last error returned by the function.
func (r Retrier[T]) Run(ctx context.Context, fn func(context.Context) (*T, error)) (*T, error) {
	timeout := r.config.Timeout()
	if timeout > 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, r.config.Timeout())
		defer cancel()
	}
	var attempt int
	var lastErr error
	for {
		attempt++
		entity, err := fn(ctx)
		if err == nil {
			return entity, nil
		}
		if !r.config.ShouldRetry(err) {
			logger.Debugf(ctx, "non-retriable error: %s", err)
			return nil, err
		}
		lastErr = err
		wait := r.config.Backoff(attempt)
		timer := time.NewTimer(wait)
		logger.Tracef(ctx, "%s. Sleeping %s",
			strings.TrimSuffix(err.Error(), "."),
			wait.Round(time.Millisecond))
		select {
		// stop when either this or parent context times out
		case <-ctx.Done():
			timer.Stop()
			return nil, &ErrTimedOut{lastErr}
		case <-timer.C:
		}
	}
}

func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	e := err.(*Err)
	if e == nil {
		return false
	}
	return !e.Halt
}

func Wait(ctx context.Context, timeout time.Duration, fn func() *Err) error {
	return New[struct{}](WithTimeout(timeout), WithRetryFunc(shouldRetry)).Wait(ctx, func(_ context.Context) error {
		err := fn()
		if err != nil {
			return err
		}
		return nil
	})
}

func Poll[T any](ctx context.Context, timeout time.Duration, fn func() (*T, *Err)) (*T, error) {
	return New[T](WithTimeout(timeout), WithRetryFunc(shouldRetry)).Run(ctx, func(_ context.Context) (*T, error) {
		res, err := fn()
		if err != nil {
			return res, err
		}
		return res, nil
	})
}
