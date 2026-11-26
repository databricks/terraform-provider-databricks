package cache

import (
	"sync"

	"golang.org/x/oauth2"
)

// NewInMemoryTokenCache returns a TokenCache that stores tokens in process
// memory only. Tokens do not persist across process restarts. This is the
// default cache used by PersistentAuth when no WithTokenCache option is
// provided. Most production consumers should supply a persistent cache
// implementation (for example, the file-based cache in
// github.com/databricks/cli/libs/auth/storage).
func NewInMemoryTokenCache() TokenCache {
	return &inMemoryTokenCache{tokens: map[string]*oauth2.Token{}}
}

type inMemoryTokenCache struct {
	mu     sync.Mutex
	tokens map[string]*oauth2.Token
}

func cloneToken(t *oauth2.Token) *oauth2.Token {
	if t == nil {
		return nil
	}
	clone := *t
	return &clone
}

// Store implements TokenCache. A nil token deletes the key.
func (c *inMemoryTokenCache) Store(key string, t *oauth2.Token) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if t == nil {
		delete(c.tokens, key)
		return nil
	}
	c.tokens[key] = cloneToken(t)
	return nil
}

// Lookup implements TokenCache. Returns ErrNotFound when the key is absent.
func (c *inMemoryTokenCache) Lookup(key string) (*oauth2.Token, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	t, ok := c.tokens[key]
	if !ok {
		return nil, ErrNotFound
	}
	return cloneToken(t), nil
}
