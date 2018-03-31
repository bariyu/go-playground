package lrfucache

// LRFUCache is the interface for simple LRU/LFU cache
type LRFUCache interface {
	// Sets given value in the cache, returns true if an another element is evicted
	Set(key, value interface{}) (evicted bool)

	// Return's keys value from the cache and it's existense
	// Updates the frequency or recentness of that key.
	Get(key interface{}) (value interface{}, ok bool)

	// Tries to delete given key fron the cache
	// returns true if key is deleted false othwervise
	Delete(key interface{}) (deleted bool)

	// Checks if key is in the cache or not
	// doesn't update frequency/recentness of the key.
	Contains(key interface{}) (ok bool)

	// Returns keys contained in this cache in order specified by
	// parameter newest, as a slice
	Keys(newest bool) (keys []interface{})

	// Returns values contained in this cache in order specified by
	// paramater newest, as a slice
	Values(newest bool) (values []interface{})

	// Returns keys and values contained in this cache in order specified by
	// paramater newest, as a slices
	Enumerate(newest bool) (keys, values []interface{})

	// Returns number of elements present in the cache.
	Size() (size int)

	// Returns the capacity of the cache
	Capacity() (capacity int)

	// Removes all keys/values from the cache
	Clear()

	// Just prints the cache to the console
	// Used for debugging
	PrintCache()
}
