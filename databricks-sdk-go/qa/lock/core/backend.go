package core

import (
	"context"
	"time"
)

// LockBackend is the interface for a lock backend.
type LockBackend interface {
	// PrepareBackend prepares the backend for use. This is called once per lock.
	// Locking is aborted if this fails.
	PrepareBackend(ctx context.Context, lockId string) error

	// AcquireLock acquires the lock. This is called once per lock. Locking is aborted
	// if this fails.
	AcquireLock(ctx context.Context, contents *LockState) error

	// RenewLock renews the lock. This is called periodically while the lock is held.
	RenewLock(ctx context.Context, leaseId string) error

	// ReleaseLock releases the lock. This is called once per lock.
	ReleaseLock(ctx context.Context, leaseId string) error

	// RefreshDuration returns the duration between calls to RenewLock.
	RefreshDuration() time.Duration
}
