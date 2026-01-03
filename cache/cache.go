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
