package lrfucache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLRUCacheSetGetEviction(t *testing.T) {
	var cache LRFUCache
	cache = NewLRUCache(5)

	if cache == nil {
		t.Errorf("cannot initialize new LRUCache")
	}

	if cache.Capacity() != 5 {
		t.Errorf("capactiy of the LRUCache should be 5")
	}

	if cache.Size() != 0 {
		t.Errorf("size of the LRUCache should be 0")
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

	// Getting ams back from the cache
	position, ok = cache.Get(amsterdamKey)
	if ok != true || position != amsterdam {
		t.Errorf("should get amsterdam's location from the cache")
	}

	evicted = cache.Set(bookingHQKey, bookingHQ)
	if evicted {
		t.Errorf("setting bookingHQ should not cause eviction")
	}

	evicted = cache.Set(anitkabirKey, anitkabir)
	if evicted {
		t.Errorf("setting anitkabir should not cause eviction")
	}

	evicted = cache.Set(mETUCENGKey, mETUCENG)
	if evicted {
		t.Errorf("setting metu ceng should not cause eviction")
	}

	evicted = cache.Set(googleHQKey, googleHQ)
	if evicted {
		t.Errorf("setting googlehq should not cause eviction")
	}

	// google, metuceng, anitkabir, booking, amsterdam
	// cache.PrintCache()

	// brouwerijt, google, metuceng, anitkabir, booking
	evicted = cache.Set(brouwerijtIJKey, brouwerijtIJ)
	if evicted != true {
		t.Errorf("setting brouwerijtIJ should cause eviction")
	}

	position, ok = cache.Get(amsterdamKey)
	if ok == true || position != nil {
		t.Errorf("should not get amsterdam's location from the cache, it should have been evicted")
	}

	// work, brouwerijt, google, metuceng, anitkabir
	evicted = cache.Set(workLocationKey, uPOffice)
	if evicted != true {
		t.Errorf("setting work should cause eviction")
	}

	position, ok = cache.Get(bookingHQKey)
	if ok == true || position != nil {
		t.Errorf("should not get thebank's location from the cache, it should have been evicted")
	}

	cache.PrintCache()

	// This key is quite special and should never ever be evicted :)
	// Call get to make sure to move it to the head.
	// anitkabir, work, brouwerijt, google, metuceng
	cache.Get(anitkabirKey)

	// work, anitkabir, brouwerijt, google, metuceng
	// longer commute :(
	evicted = cache.Set(workLocationKey, bookingHQ)
	if evicted {
		t.Errorf("no need to evict anything, updating existing key work in the cache")
	}

	coolOffices := []testPosition{
		facebookLondon,
		googleHQ,
		bookingHQ,
	}
	rand.Seed(time.Now().UnixNano())
	nextLocation := coolOffices[rand.Intn(len(coolOffices)-1)]
	fmt.Printf("NEXT: %v\n", nextLocation)
	cache.Set(workLocationKey, nextLocation)

	cache.Get(anitkabirKey)
	desiredKeys := []string{
		anitkabirKey,
		workLocationKey,
		brouwerijtIJKey,
		googleHQKey,
		mETUCENGKey,
	}

	keys := cache.Keys(true)
	for i := 0; i < 5; i++ {
		ok = cache.Contains(desiredKeys[i])
		if ok != true {
			t.Errorf("key: [%v] should be in the cache.", desiredKeys[i])
		}
		if keys[i] != desiredKeys[i] {
			t.Errorf("key order doesn't match want: [%v], got: [%v]", desiredKeys[i], keys[i])
		}
	}

	desiredLocations := []testPosition{
		anitkabir,
		nextLocation,
		brouwerijtIJ,
		googleHQ,
		mETUCENG,
	}
	keys, values := cache.Enumerate(true)
	for i := 0; i < 5; i++ {
		position, ok = cache.Get(desiredKeys[i])
		if ok != true {
			t.Errorf("key: [%v] should be in the cache.", desiredKeys[i])
		}
		if keys[i] != desiredKeys[i] {
			t.Errorf("key order doesn't match want: [%v], got: [%v]", desiredKeys[i], keys[i])
		}
		if values[i] != desiredLocations[i] {
			t.Errorf("values order doesn't match want: [%v], got [%v]", desiredLocations[i], values[i])
		}
	}

	cache.PrintCache()

	deleted := cache.Delete(amsterdamKey)
	if deleted == true {
		t.Errorf("shouldn't have deleted amsterdam key from the cache because it shouldn't exist")
	}

	deleted = cache.Delete(googleHQKey)
	if deleted != true {
		t.Errorf("should have deleted googlehq key from the cache because it should be in the cache before deleting")
	}

	position, ok = cache.Get(googleHQKey)
	if ok == true {
		t.Errorf("should not get googlehq key from the cache because it should been deleted previously")
	}

	keysDesc := cache.Keys(true)
	keysAsc := cache.Keys(false)
	fmt.Printf("desc keys after deleting googlehq: %v\n", keysDesc)

	for i := 0; i < cache.Size(); i++ {
		descKey := keysDesc[i]
		ascKey := keysAsc[cache.Size()-1-i]

		if descKey != ascKey {
			t.Errorf("asc, desc key order doesn't match")
		}
	}

	cache.PrintCache()

	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("cache clear failed")
	}

	cache.PrintCache()
}
