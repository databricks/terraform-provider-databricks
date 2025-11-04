package permissions

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestObjectMutexManager_LockUnlock(t *testing.T) {
	manager := &objectMutexManager{}

	objectType := "clusters"
	objectID := "test-cluster"

	// First lock should succeed immediately
	manager.Lock(objectType, objectID)

	// Verify mutex was created by checking we can load it
	key := "clusters/test-cluster"
	_, exists := manager.mutexes.Load(key)
	assert.True(t, exists, "Mutex should be created for object")

	// Unlock
	manager.Unlock(objectType, objectID)
}

func TestObjectMutexManager_ConcurrentAccess(t *testing.T) {
	manager := &objectMutexManager{}

	objectType := "clusters"
	objectID := "test-cluster"
	var counter int32
	var wg sync.WaitGroup

	// Simulate 10 concurrent operations on the same object
	// The mutex should ensure they execute serially
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			manager.Lock(objectType, objectID)
			defer manager.Unlock(objectType, objectID)

			// Simulate some work with a counter
			// Without proper locking, this would cause race conditions
			current := atomic.LoadInt32(&counter)
			time.Sleep(time.Millisecond) // Simulate some processing
			atomic.StoreInt32(&counter, current+1)
		}()
	}

	wg.Wait()

	// All 10 operations should have completed
	assert.Equal(t, int32(10), atomic.LoadInt32(&counter), "All operations should complete")
}

func TestObjectMutexManager_DifferentObjects(t *testing.T) {
	manager := &objectMutexManager{}

	objectType1 := "clusters"
	objectID1 := "cluster-1"
	objectType2 := "clusters"
	objectID2 := "cluster-2"

	var wg sync.WaitGroup
	results := make([]string, 0, 2)
	var resultsMutex sync.Mutex

	// Operations on different objects should run concurrently
	wg.Add(2)

	go func() {
		defer wg.Done()
		manager.Lock(objectType1, objectID1)
		defer manager.Unlock(objectType1, objectID1)

		time.Sleep(50 * time.Millisecond)
		resultsMutex.Lock()
		results = append(results, "object1")
		resultsMutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		manager.Lock(objectType2, objectID2)
		defer manager.Unlock(objectType2, objectID2)

		time.Sleep(50 * time.Millisecond)
		resultsMutex.Lock()
		results = append(results, "object2")
		resultsMutex.Unlock()
	}()

	wg.Wait()

	// Both operations should complete (order doesn't matter)
	assert.Len(t, results, 2, "Both operations should complete")
}

func TestObjectMutexManager_DifferentObjectTypes(t *testing.T) {
	manager := &objectMutexManager{}

	objectType1 := "clusters"
	objectType2 := "jobs"
	objectID := "same-id" // Same ID but different types

	var wg sync.WaitGroup
	results := make([]string, 0, 2)
	var resultsMutex sync.Mutex

	// Operations on different object types should run concurrently even with same ID
	wg.Add(2)

	go func() {
		defer wg.Done()
		manager.Lock(objectType1, objectID)
		defer manager.Unlock(objectType1, objectID)

		time.Sleep(50 * time.Millisecond)
		resultsMutex.Lock()
		results = append(results, "clusters")
		resultsMutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		manager.Lock(objectType2, objectID)
		defer manager.Unlock(objectType2, objectID)

		time.Sleep(50 * time.Millisecond)
		resultsMutex.Lock()
		results = append(results, "jobs")
		resultsMutex.Unlock()
	}()

	wg.Wait()

	// Both operations should complete (order doesn't matter)
	assert.Len(t, results, 2, "Both operations on different object types should complete concurrently")
}
