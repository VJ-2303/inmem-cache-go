package cache

import "sync"

// Cache represents an basic cache data structure
// it contains an Map to contains the data and an
// mutex to allow concurrent access without any dataraces
type Cache[K comparable, V any] struct {
	items map[K]V    // Storing the key-value pairs
	mu    sync.Mutex // Controlling concurrent access
}

// New creates a new cache instance
func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items: make(map[K]V),
	}
}

// Set adds or update a key-value pair in the cache
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}

// Get returns the value associated with the given key
// from the cache. the bool value will return true if no matching key is found
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.items[key]
	return value, found
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

	value, found := c.items[key]

	if found {
		delete(c.items, key)
	}
	return value, found
}
