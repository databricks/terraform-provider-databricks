package permissions

import (
	"sync"
)

// objectMutexManager manages mutexes per object ID to prevent concurrent
// operations on the same Databricks object that could lead to race conditions.
//
// This is particularly important for Delete operations where multiple
// databricks_permission resources for the same object might be deleted
// concurrently, each doing GET -> filter -> SET, which could result in
// lost permission updates.
type objectMutexManager struct {
	mutexes map[string]*sync.Mutex
	mapLock sync.Mutex
}

// globalObjectMutexManager is the singleton instance used by all permission resources
var globalObjectMutexManager = &objectMutexManager{
	mutexes: make(map[string]*sync.Mutex),
}

// Lock acquires a mutex for the given object ID.
// Each object ID gets its own mutex to allow concurrent operations on different objects
// while serializing operations on the same object.
func (m *objectMutexManager) Lock(objectID string) {
	m.mapLock.Lock()
	mu, exists := m.mutexes[objectID]
	if !exists {
		mu = &sync.Mutex{}
		m.mutexes[objectID] = mu
	}
	m.mapLock.Unlock()

	// Lock the object-specific mutex (outside the map lock to avoid deadlock)
	mu.Lock()
}

// Unlock releases the mutex for the given object ID.
func (m *objectMutexManager) Unlock(objectID string) {
	m.mapLock.Lock()
	mu, exists := m.mutexes[objectID]
	m.mapLock.Unlock()

	if exists {
		mu.Unlock()
	}
}

// lockObject acquires a lock for the given object ID.
// This should be called at the start of any operation that modifies permissions.
func lockObject(objectID string) {
	globalObjectMutexManager.Lock(objectID)
}

// unlockObject releases the lock for the given object ID.
// This should be called at the end of any operation that modifies permissions.
// Use defer to ensure it's always called even if the operation panics.
func unlockObject(objectID string) {
	globalObjectMutexManager.Unlock(objectID)
}
