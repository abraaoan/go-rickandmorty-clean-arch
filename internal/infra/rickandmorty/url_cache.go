package rickandmorty

import (
	"sync"
	"time"
)

type URLCache struct {
	mu     sync.RWMutex
	data   map[string][]byte
	expiry map[string]time.Time
	ttl    time.Duration
}

func NewURLCache(ttl time.Duration) *URLCache {
	return &URLCache{
		data:   make(map[string][]byte),
		expiry: make(map[string]time.Time),
		ttl:    ttl,
	}
}

func (c *URLCache) Get(url string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	exp, ok := c.expiry[url]
	if !ok || time.Now().After(exp) {
		return nil, false
	}

	data, exists := c.data[url]
	return data, exists
}

func (c *URLCache) Set(url string, body []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[url] = body
	c.expiry[url] = time.Now().Add(c.ttl)
}

func (c *URLCache) Delete(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, url)
	delete(c.expiry, url)
}

func (c *URLCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string][]byte)
	c.expiry = make(map[string]time.Time)
}
