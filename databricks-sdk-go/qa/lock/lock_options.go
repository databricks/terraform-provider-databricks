package lock

import (
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/qa/lock/core"
)

type LockOption func(*LockOptions)

func WithBackend(backend core.LockBackend) LockOption {
	return func(opts *LockOptions) {
		opts.Backend = backend
	}
}

func WithLeaseDuration(duration time.Duration) LockOption {
	return func(opts *LockOptions) {
		opts.LeaseDuration = duration
	}
}

func InTest(t *testing.T) LockOption {
	return func(opts *LockOptions) {
		opts.T = t
	}
}

type LockOptions struct {
	LeaseDuration time.Duration
	Backend       core.LockBackend
	T             *testing.T
}
