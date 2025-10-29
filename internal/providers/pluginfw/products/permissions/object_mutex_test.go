package permissions

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestObjectMutexManager_LockUnlock(t *testing.T) {
	manager := &objectMutexManager{
		mutexes: make(map[string]*sync.Mutex),
	}

	objectID := "/clusters/test-cluster"

	// First lock should succeed immediately
	manager.Lock(objectID)

	// Verify mutex was created
	manager.mapLock.Lock()
	_, exists := manager.mutexes[objectID]
	manager.mapLock.Unlock()
	assert.True(t, exists, "Mutex should be created for object ID")

	// Unlock
	manager.Unlock(objectID)
}

func TestObjectMutexManager_ConcurrentAccess(t *testing.T) {
	manager := &objectMutexManager{
		mutexes: make(map[string]*sync.Mutex),
	}

	objectID := "/clusters/test-cluster"
	var counter int32
	var wg sync.WaitGroup

	// Simulate 10 concurrent operations on the same object
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			manager.Lock(objectID)
			defer manager.Unlock(objectID)

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
	manager := &objectMutexManager{
		mutexes: make(map[string]*sync.Mutex),
	}

	objectID1 := "/clusters/cluster-1"
	objectID2 := "/clusters/cluster-2"

	var wg sync.WaitGroup
	results := make([]string, 0, 2)
	var resultsMutex sync.Mutex

	// Operations on different objects should run concurrently
	wg.Add(2)

	go func() {
		defer wg.Done()
		manager.Lock(objectID1)
		defer manager.Unlock(objectID1)

		time.Sleep(50 * time.Millisecond)
		resultsMutex.Lock()
		results = append(results, "object1")
		resultsMutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		manager.Lock(objectID2)
		defer manager.Unlock(objectID2)

		time.Sleep(50 * time.Millisecond)
		resultsMutex.Lock()
		results = append(results, "object2")
		resultsMutex.Unlock()
	}()

	wg.Wait()

	// Both operations should complete (order doesn't matter)
	assert.Len(t, results, 2, "Both operations should complete")
}

func TestObjectMutexManager_SerializeSameObject(t *testing.T) {
	manager := &objectMutexManager{
		mutexes: make(map[string]*sync.Mutex),
	}

	objectID := "/jobs/job-123"
	var executionOrder []int
	var orderMutex sync.Mutex
	var wg sync.WaitGroup

	// Launch 3 goroutines that should execute serially
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			manager.Lock(objectID)
			defer manager.Unlock(objectID)

			// Record execution order
			orderMutex.Lock()
			executionOrder = append(executionOrder, num)
			orderMutex.Unlock()

			time.Sleep(10 * time.Millisecond) // Simulate work
		}(i)
	}

	wg.Wait()

	// All 3 operations should complete
	assert.Len(t, executionOrder, 3, "All operations should complete")

	// Operations should be serialized (no concurrent execution)
	// We can't predict the exact order, but we can verify no race conditions occurred
}

func TestGlobalObjectMutexManager(t *testing.T) {
	// Test the global singleton
	objectID := "/notebooks/notebook-456"

	lockObject(objectID)

	// Verify we can unlock without error
	unlockObject(objectID)

	// Should be able to lock again after unlocking
	lockObject(objectID)
	unlockObject(objectID)
}

func TestObjectMutexManager_UnlockNonexistent(t *testing.T) {
	manager := &objectMutexManager{
		mutexes: make(map[string]*sync.Mutex),
	}

	// Unlocking a non-existent object should not panic
	assert.NotPanics(t, func() {
		manager.Unlock("/clusters/does-not-exist")
	}, "Unlocking non-existent object should not panic")
}
