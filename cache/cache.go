package cache

import (
	"sync"
	"time"
)

// Cache represents an basic cache data structure
// it contains an Map to contains the data and an
// mutex to allow concurrent access without any dataraces
type Cache[K comparable, V any] struct {
	items map[K]item[V]
	mu    sync.Mutex // Controlling concurrent access
}

type item[V any] struct {
	value  V
	expiry time.Time
}

func (i *item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}

// New creates a new cache instance
func New[K comparable, V any]() *Cache[K, V] {
	c := &Cache[K, V]{
		items: make(map[K]item[V]),
	}

	go func() {
		c.mu.Lock()

		for range time.Tick(5 * time.Second) {
			for key, item := range c.items {
				if item.isExpired() {
					delete(c.items, key)
				}
			}
		}
		c.mu.Unlock()
	}()
	return c
}

// Set adds or update a key-value pair in the cache
func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item[V]{
		value:  value,
		expiry: time.Now().Add(ttl),
	}
}

// Get returns the value associated with the given key
// from the cache. the bool value will return true if no matching key is found
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		return item.value, false
	}
	if item.isExpired() {
		delete(c.items, key)
		return item.value, false
	}
	return item.value, true
}

// Remove deletes the key-value with the specified key from the cache
func (c *Cache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// Pop returns an key-value from the cache, and pop the value if its exists
func (c *Cache[K, V]) Pop(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]

	if !found {
		return item.value, false
	}
	delete(c.items, key)

	if item.isExpired() {
		return item.value, false
	}
	return item.value, true
}
