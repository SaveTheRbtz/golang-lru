package simplelru

type option[K comparable, V any] func(c *LRU[K, V])

// WithOnEvictCallback is used to set the callback function for when an
// entry is evicted from the cache.
func WithOnEvictCallback[K comparable, V any](onEvict EvictCallback[K, V]) option[K, V] {
	return func(c *LRU[K, V]) {
		c.onEvict = onEvict
	}
}

// WithPreallocate is used to set the initial size of the map used to store
// the cache entries. This can be used to avoid resizing the map when the
// cache is first populated.
func WithPreallocate[K comparable, V any](preallocate int) option[K, V] {
	return func(c *LRU[K, V]) {
		c.preallocate = preallocate
	}
}
