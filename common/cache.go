package common

import (
	"sync"

	"golang.org/x/sync/singleflight"
)

// CacheEntry is a single cached value with generation-based invalidation safety.
//
// Concurrent reads on a warm cache proceed without blocking (RLock).  A
// singleflight group ensures at most one in-flight fetch when the cache is
// cold, so all waiting goroutines share a single result.  A generation counter
// prevents a late-arriving singleflight leader from overwriting data that was
// invalidated after the fetch started: Invalidate increments the generation
// before clearing the value, and the leader only commits if the generation is
// unchanged.
type CacheEntry[V any] struct {
	mu          sync.RWMutex
	generation  uint64
	initialized bool
	value       V
	sg          singleflight.Group
}

// GetOrFetch returns the cached value if available, otherwise calls fetch
// exactly once (via singleflight) and caches the result.
func (e *CacheEntry[V]) GetOrFetch(fetch func() (V, error)) (V, error) {
	// Fast path: warm cache — many goroutines can hold RLock simultaneously.
	e.mu.RLock()
	if e.initialized {
		v := e.value
		e.mu.RUnlock()
		return v, nil
	}
	e.mu.RUnlock()

	// Slow path: cold cache.  Use singleflight so exactly one goroutine calls
	// the API; all others block until the result is available.
	type wrapper struct{ v V }
	raw, err, _ := e.sg.Do("fetch", func() (any, error) {
		// Double-check now that we hold the singleflight slot.
		e.mu.RLock()
		if e.initialized {
			v := e.value
			e.mu.RUnlock()
			return wrapper{v}, nil
		}
		gen := e.generation
		e.mu.RUnlock()

		fetched, err := fetch()
		if err != nil {
			return wrapper{}, err
		}
		e.mu.Lock()
		// Only commit if Invalidate has not been called in the interim.
		if e.generation == gen {
			e.value = fetched
			e.initialized = true
		}
		e.mu.Unlock()
		return wrapper{fetched}, nil
	})
	if err != nil {
		var zero V
		return zero, err
	}
	return raw.(wrapper).v, nil
}

// Invalidate clears the cached value and bumps the generation counter so that
// any in-flight fetch that started before the invalidation cannot re-populate
// the cache with stale pre-mutation data.
func (e *CacheEntry[V]) Invalidate() {
	e.mu.Lock()
	e.generation++
	e.initialized = false
	var zero V
	e.value = zero
	e.mu.Unlock()
}

// KeyedCache maps arbitrary comparable keys to CacheEntry values. It is safe
// for concurrent use by multiple goroutines.
type KeyedCache[K comparable, V any] struct {
	mu    sync.Mutex
	cache map[K]*CacheEntry[V]
}

// NewKeyedCache creates an empty KeyedCache.
func NewKeyedCache[K comparable, V any]() *KeyedCache[K, V] {
	return &KeyedCache[K, V]{cache: make(map[K]*CacheEntry[V])}
}

func (c *KeyedCache[K, V]) getOrCreateEntry(key K) *CacheEntry[V] {
	c.mu.Lock()
	defer c.mu.Unlock()
	if e, ok := c.cache[key]; ok {
		return e
	}
	e := &CacheEntry[V]{}
	c.cache[key] = e
	return e
}

// Get returns the value for key, calling fetch to populate it if absent or
// previously invalidated.
func (c *KeyedCache[K, V]) Get(key K, fetch func() (V, error)) (V, error) {
	return c.getOrCreateEntry(key).GetOrFetch(fetch)
}

// Invalidate clears the cache for key, triggering a fresh fetch on the next
// Get.  It is safe to call concurrently with Get.
func (c *KeyedCache[K, V]) Invalidate(key K) {
	c.mu.Lock()
	e, ok := c.cache[key]
	c.mu.Unlock()
	if ok {
		e.Invalidate()
	}
}
