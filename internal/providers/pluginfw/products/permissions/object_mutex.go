package permissions

import (
	"fmt"
	"sync"
)

// objectMutexManager manages mutexes per object to prevent concurrent
// operations on the same Databricks object that could lead to race conditions.
//
// This is particularly important for Delete operations where multiple
// databricks_permission resources for the same object might be deleted
// concurrently, each doing GET -> filter -> SET, which could result in
// lost permission updates.
type objectMutexManager struct {
	mutexes sync.Map // map[string]*sync.Mutex
}

// globalObjectMutexManager is the singleton instance used by all permission resources
var globalObjectMutexManager = &objectMutexManager{}

// Lock acquires a mutex for the given object type and ID.
// Each object gets its own mutex to allow concurrent operations on different objects
// while serializing operations on the same object.
func (m *objectMutexManager) Lock(objectType, objectID string) {
	key := fmt.Sprintf("%s/%s", objectType, objectID)

	// LoadOrStore returns the existing value if present, otherwise stores and returns the given value
	value, _ := m.mutexes.LoadOrStore(key, &sync.Mutex{})
	mu := value.(*sync.Mutex)

	// Lock the object-specific mutex
	mu.Lock()
}

// Unlock releases the mutex for the given object type and ID.
func (m *objectMutexManager) Unlock(objectType, objectID string) {
	key := fmt.Sprintf("%s/%s", objectType, objectID)

	value, ok := m.mutexes.Load(key)
	if ok {
		mu := value.(*sync.Mutex)
		mu.Unlock()
	}
}

// lockObject acquires a lock for the given object type and ID.
// This should be called at the start of any operation that modifies permissions.
func lockObject(objectType, objectID string) {
	globalObjectMutexManager.Lock(objectType, objectID)
}

// unlockObject releases the lock for the given object type and ID.
// This should be called at the end of any operation that modifies permissions.
// Use defer to ensure it's always called even if the operation panics.
func unlockObject(objectType, objectID string) {
	globalObjectMutexManager.Unlock(objectType, objectID)
}
