package core

// Lockable is the interface for a lockable resource.
type Lockable interface {
	// GetLockId returns the lock id for the resource.
	GetLockId() string
}
