package lock

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/qa/lock/core"
	"github.com/databricks/databricks-sdk-go/qa/lock/databricks"
	"github.com/google/uuid"
)

// Acquire acquires a lock according to the provided lock options. If successful,
// the lock is returned. The caller is responsible for calling the Unlock() method
// on the returned lock when the lock is no longer needed. If the lock is held
// at the moment Acquire() is called, the method will block for up to the lease
// duration, retrying periodically to acquire the lock. If the lock is not
// acquired in time, an error will be returned.
//
// When the lock is no longer needed, the caller should call the Unlock() method
// on the returned lock. This will release the lock and return any errors that
// occurred while releasing the lock. If a *testing.T is provided as a
// LockOption, the Unlock() method will be called automatically when the test
// completes.
//
// By default, the lock will be acquired for 1 minute. This can be changed by
// passing the WithLeaseDuration() LockOption.
//
// By default, the lock will be acquired using a Databricks backend. This can be
// changed by passing the WithBackend() LockOption.
func Acquire(ctx context.Context, lockable core.Lockable, os ...LockOption) (*Lock, error) {
	opts := LockOptions{}
	for _, o := range os {
		o(&opts)
	}
	if opts.Backend == nil {
		opts.Backend = &databricks.Backend{}
	}
	if opts.LeaseDuration == 0 {
		opts.LeaseDuration = time.Minute
	}

	lockId := lockable.GetLockId()
	err := opts.Backend.PrepareBackend(ctx, lockId)
	if err != nil {
		return nil, fmt.Errorf("error preparing backend: %w", err)
	}

	leaseId := uuid.New().String()
	timesToPoll := int(opts.LeaseDuration/(10*time.Second)) + 1
	for i := 0; i < timesToPoll; i++ {
		state := core.NewLockState(lockable, leaseId, opts.LeaseDuration, opts.T)
		err = opts.Backend.AcquireLock(ctx, state)
		if err == nil {
			break
		}
		if i == timesToPoll-1 {
			return nil, fmt.Errorf("max wait duration exceeded, last error: %w", err)
		}
		fmt.Printf("Failed to acquire lock %s, retrying in 10 seconds. error: %s\n", lockId, err)
		time.Sleep(10 * time.Second)
	}

	// Channel to signal when the lock should be released
	done := make(chan struct{})

	// Channel to communicate errors renewing the lease or deleting the blob
	errs := make(chan error, 1)

	// Renew the lock periodically until the lock is released.
	go func() {
		duration := opts.Backend.RefreshDuration()
		ticker := time.NewTicker(duration)
		testTimer := time.NewTimer(opts.LeaseDuration)
		defer close(errs)
		defer ticker.Stop()
		for {
			select {
			case <-done:
				err := opts.Backend.ReleaseLock(ctx, leaseId)
				if err != nil {
					errs <- err
				}
				return
			case <-ticker.C:
				err := opts.Backend.RenewLock(ctx, leaseId)
				if err != nil {
					errs <- err
					return
				}
				fmt.Printf("Successfully renewed lease %s\n", leaseId)
			case <-testTimer.C:
				errs <- errors.New("max wait duration exceeded")
				return
			}
		}
	}()

	lock := &Lock{
		done: done,
		errs: errs,
	}

	if opts.T != nil {
		opts.T.Cleanup(func() {
			err := lock.Unlock()
			if err != nil {
				opts.T.Errorf("error unlocking: %s", err)
			}
		})
	}

	return lock, nil
}

// Lock is a lock that is acquired by a single test. Locks can be acquired
// by calling the Acquire() method on a LockOptions struct. The lock is
// released when the Unlock() method is called. Failures to acquire the lock
// in time, to renew the leease, or to release the lock will cause the test
// to fail.
type Lock struct {
	done chan<- struct{}
	errs <-chan error
}

func (l *Lock) Unlock() error {
	close(l.done)
	return <-l.errs
}
