package lock

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/qa/lock/core"
	"github.com/stretchr/testify/assert"
)

type testBackend struct {
	prepareBackend func(lockId string) error
	acquireLock    func(contents *core.LockState) error
	renewLock      func(leaseId string) error
	releaseLock    func(leaseId string) error
}

func (b *testBackend) PrepareBackend(_ context.Context, lockId string) error {
	return b.prepareBackend(lockId)
}

func (b *testBackend) AcquireLock(_ context.Context, contents *core.LockState) error {
	return b.acquireLock(contents)
}

func (b *testBackend) RenewLock(_ context.Context, leaseId string) error {
	return b.renewLock(leaseId)
}

func (b *testBackend) ReleaseLock(_ context.Context, leaseId string) error {
	return b.releaseLock(leaseId)
}

func (b *testBackend) RefreshDuration() time.Duration {
	return 10 * time.Millisecond
}

var _ core.LockBackend = &testBackend{}

func getOptions(backend *testBackend) []LockOption {
	return []LockOption{
		WithBackend(backend),
		WithLeaseDuration(1 * time.Second),
	}
}

func TestAcquire_PrepareBackendFails(t *testing.T) {
	errTest := errors.New("test error")
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return errTest
		},
	}
	_, err := Acquire(context.Background(), NewLockable("lock"), getOptions(backend)...)
	assert.ErrorIs(t, err, errTest)
}

func TestAcquire_AcquireLockFails(t *testing.T) {
	errTest := errors.New("test error")
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return nil
		},
		acquireLock: func(contents *core.LockState) error {
			return errTest
		},
	}
	_, err := Acquire(context.Background(), NewLockable("lock"), getOptions(backend)...)
	assert.ErrorIs(t, err, errTest)
}

func TestAcquire_UnlockFails(t *testing.T) {
	errTest := errors.New("test error")
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return nil
		},
		acquireLock: func(contents *core.LockState) error {
			return nil
		},
		releaseLock: func(leaseId string) error {
			return errTest
		},
		renewLock: func(leaseId string) error {
			return nil
		},
	}
	lock, err := Acquire(context.Background(), NewLockable("lock"), getOptions(backend)...)
	assert.NoError(t, err)
	err = lock.Unlock()
	assert.ErrorIs(t, err, errTest)
}

func TestAcquire_UnlockSucceeds(t *testing.T) {
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return nil
		},
		acquireLock: func(contents *core.LockState) error {
			return nil
		},
		releaseLock: func(leaseId string) error {
			return nil
		},
		renewLock: func(leaseId string) error {
			return nil
		},
	}
	lock, err := Acquire(context.Background(), NewLockable("lock"), getOptions(backend)...)
	assert.NoError(t, err)
	err = lock.Unlock()
	assert.NoError(t, err)
}
