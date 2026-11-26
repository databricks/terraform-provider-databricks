package cache

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"golang.org/x/oauth2"
)

func TestInMemoryTokenCache_StoreAndLookup(t *testing.T) {
	c := NewInMemoryTokenCache()
	tok := &oauth2.Token{AccessToken: "abc", RefreshToken: "def"}
	if err := c.Store("key1", tok); err != nil {
		t.Fatalf("Store: %v", err)
	}
	got, err := c.Lookup("key1")
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	if got.AccessToken != "abc" {
		t.Errorf("AccessToken: want %q, got %q", "abc", got.AccessToken)
	}
	if got.RefreshToken != "def" {
		t.Errorf("RefreshToken: want %q, got %q", "def", got.RefreshToken)
	}
}

func TestInMemoryTokenCache_StoreAndLookupUseCopies(t *testing.T) {
	c := NewInMemoryTokenCache()
	tok := &oauth2.Token{AccessToken: "abc", RefreshToken: "def"}
	if err := c.Store("key1", tok); err != nil {
		t.Fatalf("Store: %v", err)
	}

	tok.RefreshToken = "mutated-after-store"

	got, err := c.Lookup("key1")
	if err != nil {
		t.Fatalf("Lookup: %v", err)
	}
	if got.RefreshToken != "def" {
		t.Fatalf("RefreshToken after store mutation: want %q, got %q", "def", got.RefreshToken)
	}

	got.RefreshToken = "mutated-after-lookup"

	gotAgain, err := c.Lookup("key1")
	if err != nil {
		t.Fatalf("Lookup after lookup mutation: %v", err)
	}
	if gotAgain.RefreshToken != "def" {
		t.Fatalf("RefreshToken after lookup mutation: want %q, got %q", "def", gotAgain.RefreshToken)
	}
}

func TestInMemoryTokenCache_StoreNilDeletesKey(t *testing.T) {
	c := NewInMemoryTokenCache()
	if err := c.Store("key1", &oauth2.Token{AccessToken: "abc"}); err != nil {
		t.Fatalf("Store: %v", err)
	}
	if err := c.Store("key1", nil); err != nil {
		t.Fatalf("Store(nil): %v", err)
	}
	_, err := c.Lookup("key1")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Lookup after nil Store: want ErrNotFound, got %v", err)
	}
}

func TestInMemoryTokenCache_LookupUnsetKey(t *testing.T) {
	c := NewInMemoryTokenCache()
	_, err := c.Lookup("missing")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Lookup(missing): want ErrNotFound, got %v", err)
	}
}

func TestInMemoryTokenCache_ConcurrentStoreAndLookup(t *testing.T) {
	c := NewInMemoryTokenCache()
	const goroutines = 20
	const iterations = 100
	var wg sync.WaitGroup
	wg.Add(goroutines * 2)
	for i := 0; i < goroutines; i++ {
		prefix := fmt.Sprintf("writer-%d", i)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := fmt.Sprintf("%s-%d", prefix, j)
				_ = c.Store(key, &oauth2.Token{AccessToken: key})
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_, _ = c.Lookup(fmt.Sprintf("%s-%d", prefix, j))
			}
		}()
	}
	wg.Wait()
}
