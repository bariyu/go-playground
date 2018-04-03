package lrfucache

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	var cache LRFUCache
	cache = NewLFUCache(5)

	if cache == nil {
		t.Errorf("cannot initialize new LFUCache")
	}

	if cache.Capacity() != 5 {
		t.Errorf("capactiy of the LRUCaLFUCacheche should be 5")
	}

	if cache.Size() != 0 {
		t.Errorf("size of the LFUCache should be 0")
	}

	var evicted, ok bool
	// var place interface{}
	var position interface{}

	// Getting ams from empty cache
	position, ok = cache.Get(amsterdamKey)
	if ok == true || position != nil {
		t.Errorf("should not get amsterdam's location from empty cache")
	}

	// Setting ams in the cache
	evicted = cache.Set(amsterdamKey, amsterdam)
	if evicted {
		t.Errorf("setting amsterdam should not cause eviction")
	}

	if cache.Size() != 1 {
		t.Errorf("size of the LFUCache should be 1")
	}

	// Getting ams back from the cache
	position, ok = cache.Get(amsterdamKey)
	if ok != true || position != amsterdam {
		t.Errorf("should get amsterdam's location from the cache")
	}

	cache.PrintCache()

}
