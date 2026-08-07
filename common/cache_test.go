package common

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCacheEntry_FetchOnce(t *testing.T) {
	var e CacheEntry[string]
	var calls int32

	v1, err := e.GetOrFetch(func() (string, error) {
		atomic.AddInt32(&calls, 1)
		return "hello", nil
	})
	require.NoError(t, err)
	assert.Equal(t, "hello", v1)

	// Second call must use the cached value; fetch must NOT be invoked again.
	v2, err := e.GetOrFetch(func() (string, error) {
		t.Fatal("fetch should not be called on warm cache")
		return "", nil
	})
	require.NoError(t, err)
	assert.Equal(t, "hello", v2)
	assert.EqualValues(t, 1, atomic.LoadInt32(&calls))
}

func TestCacheEntry_FetchError(t *testing.T) {
	var e CacheEntry[string]
	var calls int32

	_, err := e.GetOrFetch(func() (string, error) {
		atomic.AddInt32(&calls, 1)
		return "", fmt.Errorf("api error")
	})
	assert.Error(t, err)

	// Cache must not be populated on error; next call should retry the fetch.
	_, err2 := e.GetOrFetch(func() (string, error) {
		atomic.AddInt32(&calls, 1)
		return "", fmt.Errorf("api error again")
	})
	assert.Error(t, err2)
	assert.EqualValues(t, 2, atomic.LoadInt32(&calls), "fetch must be retried after an error")
}

func TestCacheEntry_InvalidateThenGet(t *testing.T) {
	var e CacheEntry[string]
	var calls int32

	v1, err := e.GetOrFetch(func() (string, error) {
		n := atomic.AddInt32(&calls, 1)
		return fmt.Sprintf("value-%d", n), nil
	})
	require.NoError(t, err)
	assert.Equal(t, "value-1", v1)

	e.Invalidate()

	// After invalidation the cache is cold; a new fetch must be triggered.
	v2, err := e.GetOrFetch(func() (string, error) {
		n := atomic.AddInt32(&calls, 1)
		return fmt.Sprintf("value-%d", n), nil
	})
	require.NoError(t, err)
	assert.Equal(t, "value-2", v2)
	assert.EqualValues(t, 2, atomic.LoadInt32(&calls))
}

// TestCacheEntry_InvalidateDuringInflight verifies the generation-counter
// invariant: a singleflight leader that completes *after* Invalidate() must
// NOT commit its result to the cache, so the very next GetOrFetch triggers a
// fresh fetch.
func TestCacheEntry_InvalidateDuringInflight(t *testing.T) {
	var e CacheEntry[int]

	fetchStarted := make(chan struct{})
	fetchCanProceed := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	var inflightResult int
	var inflightErr error

	go func() {
		defer wg.Done()
		inflightResult, inflightErr = e.GetOrFetch(func() (int, error) {
			close(fetchStarted) // tell the test we are in the fetch
			<-fetchCanProceed   // block until the test has called Invalidate
			return 42, nil      // return a "stale" value
		})
	}()

	<-fetchStarted

	// Invalidate the cache while the fetch is still blocked.
	e.Invalidate()

	// Let the in-flight fetch finish.  It must return 42 to the caller (that
	// is the singleflight contract), but must NOT commit 42 to the cache.
	close(fetchCanProceed)
	wg.Wait()

	require.NoError(t, inflightErr)
	assert.Equal(t, 42, inflightResult, "in-flight caller should still receive the fetched value")

	// The next GetOrFetch must see a cold cache and call its own fetch, NOT
	// return the stale 42.
	fresh, err := e.GetOrFetch(func() (int, error) { return 99, nil })
	require.NoError(t, err)
	assert.Equal(t, 99, fresh, "cache must not be poisoned by the stale in-flight result")
}

// TestCacheEntry_ConcurrentReads verifies that multiple goroutines on a warm
// cache all receive the same value without triggering multiple fetches.
func TestCacheEntry_ConcurrentReads(t *testing.T) {
	var e CacheEntry[int]
	var calls int32

	// Warm the cache.
	_, err := e.GetOrFetch(func() (int, error) {
		atomic.AddInt32(&calls, 1)
		return 7, nil
	})
	require.NoError(t, err)

	var wg sync.WaitGroup
	results := make([]int, 100)
	for i := range results {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			v, _ := e.GetOrFetch(func() (int, error) {
				t.Error("fetch called on warm cache during concurrent reads")
				return 0, nil
			})
			results[idx] = v
		}(i)
	}
	wg.Wait()
	for i, v := range results {
		assert.Equal(t, 7, v, "goroutine %d got wrong value", i)
	}
	assert.EqualValues(t, 1, atomic.LoadInt32(&calls), "fetch must only be called once")
}

func TestKeyedCache_IndependentKeys(t *testing.T) {
	c := NewKeyedCache[string, int]()
	var calls int32

	v1, err := c.Get("a", func() (int, error) { atomic.AddInt32(&calls, 1); return 1, nil })
	require.NoError(t, err)
	assert.Equal(t, 1, v1)

	v2, err := c.Get("b", func() (int, error) { atomic.AddInt32(&calls, 1); return 2, nil })
	require.NoError(t, err)
	assert.Equal(t, 2, v2)

	// Both keys must now be cached; no extra fetches.
	v1b, _ := c.Get("a", func() (int, error) { t.Fatal("a should be cached"); return 0, nil })
	assert.Equal(t, 1, v1b)
	v2b, _ := c.Get("b", func() (int, error) { t.Fatal("b should be cached"); return 0, nil })
	assert.Equal(t, 2, v2b)

	assert.EqualValues(t, 2, atomic.LoadInt32(&calls))
}

func TestKeyedCache_InvalidateSpecificKey(t *testing.T) {
	c := NewKeyedCache[string, int]()
	var aCalls, bCalls int32

	c.Get("a", func() (int, error) { atomic.AddInt32(&aCalls, 1); return 10, nil }) //nolint:errcheck
	c.Get("b", func() (int, error) { atomic.AddInt32(&bCalls, 1); return 20, nil }) //nolint:errcheck

	c.Invalidate("a")

	// "a" must be re-fetched; "b" must remain cached.
	va, _ := c.Get("a", func() (int, error) { atomic.AddInt32(&aCalls, 1); return 11, nil })
	c.Get("b", func() (int, error) { t.Fatal("b should not be re-fetched"); return 0, nil }) //nolint:errcheck

	assert.Equal(t, 11, va)
	assert.EqualValues(t, 2, atomic.LoadInt32(&aCalls), "a should have been fetched twice")
	assert.EqualValues(t, 1, atomic.LoadInt32(&bCalls), "b should have been fetched once")
}

func TestKeyedCache_InvalidateMissingKey(t *testing.T) {
	c := NewKeyedCache[string, int]()
	// Must not panic when invalidating a key that has never been used.
	c.Invalidate("nonexistent")
}
