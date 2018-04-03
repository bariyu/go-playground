package lrfucache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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
	var position interface{}

	// Getting ams from empty cache
	position, ok = cache.Get(amsterdamKey)
	if ok == true || position != nil {
		t.Errorf("should not get amsterdam's location from empty cache")
	}

	// Setting ams in the cache
	// ams: 1
	evicted = cache.Set(amsterdamKey, amsterdam)
	if evicted {
		t.Errorf("setting amsterdam should not cause eviction")
	}

	if cache.Size() != 1 {
		t.Errorf("size of the LFUCache should be 1")
	}

	// Getting ams back from the cache
	// ams: 2
	position, ok = cache.Get(amsterdamKey)
	if ok != true || position != amsterdam {
		t.Errorf("should get amsterdam's location from the cache")
	}

	// ams: 2, bookinghq: 1
	evicted = cache.Set(bookingHQKey, bookingHQ)
	if evicted {
		t.Errorf("setting bookingHQ should not cause eviction")
	}

	// cache.PrintCache()

	// ams: 2, anitkabir: 1, bookinghq: 1
	evicted = cache.Set(anitkabirKey, anitkabir)
	if evicted {
		t.Errorf("setting anitkabir should not cause eviction")
	}

	// ams: 2, metuceng: 1, anitkabir: 1, bookinghq: 1
	evicted = cache.Set(mETUCENGKey, mETUCENG)
	if evicted {
		t.Errorf("setting metu ceng should not cause eviction")
	}

	// cache.PrintCache()

	// ams: 2, googlehq: 1, metuceng: 1, anitkabir: 1, bookinghq: 1
	evicted = cache.Set(googleHQKey, googleHQ)
	if evicted {
		t.Errorf("setting googlehq should not cause eviction")
	}

	// ams: 2, googlehq: 1, metuceng: 1, anitkabir: 1, bookinghq: 1
	// cache.PrintCache()

	// ams: 2, brouwerij: 1, googlehq: 1, metuceng: 1, anitkabir: 1
	evicted = cache.Set(brouwerijtIJKey, brouwerijtIJ)
	if evicted != true {
		t.Errorf("setting brouwerijtIJ should cause eviction")
	}

	desiredKeys := []string{
		amsterdamKey,
		brouwerijtIJKey,
		googleHQKey,
		mETUCENGKey,
		anitkabirKey,
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

	// cache.PrintCache()

	// ams: 2, brouwerij: 1, googlehq: 1, metuceng: 1, anitkabir: 1
	position, ok = cache.Get(bookingHQKey)
	if ok == true || position != nil {
		t.Errorf("should not get bookinghq's location from the cache, it should have been evicted")
	}

	// anitkabir: 4, ams: 2, brouwerij: 1, googlehq: 1, metuceng: 1
	cache.Get(anitkabirKey)
	cache.Get(anitkabirKey)
	cache.Get(anitkabirKey)
	// cache.PrintCache()

	// anitkabir: 4, ams: 2, work: 1, brouwerij: 1, googlehq: 1
	evicted = cache.Set(workLocationKey, uPOffice)
	if evicted != true {
		t.Errorf("setting work should cause eviction")
	}

	// anitkabir: 4, ams: 2, work: 1, brouwerij: 1, googlehq: 1
	position, ok = cache.Get(mETUCENGKey)
	if ok == true || position != nil {
		t.Errorf("should not get metuceng's location from the cache, it should have been evicted")
	}

	// cache.PrintCache()

	// anitkabir: 4, work: 2, ams: 2, brouwerij: 1, googlehq: 1
	evicted = cache.Set(workLocationKey, bookingHQ)
	if evicted {
		t.Errorf("no need to evict anything, updating existing key work in the cache")
	}

	// cache.PrintCache()

	coolOffices := []testPosition{
		facebookLondon,
		googleHQ,
		bookingHQ,
	}
	rand.Seed(time.Now().UnixNano())
	nextLocation := coolOffices[rand.Intn(len(coolOffices)-1)]
	fmt.Printf("NEXT: %v\n", nextLocation)

	// anitkabir: 4, work: 3, ams: 2, brouwerij: 1, googlehq: 1
	cache.Set(workLocationKey, nextLocation)

	// anitkabir: 5, work: 3, ams: 2, brouwerij: 1, googlehq: 1
	cache.Get(anitkabirKey)
	desiredKeys = []string{
		anitkabirKey,
		workLocationKey,
		amsterdamKey,
		brouwerijtIJKey,
		googleHQKey,
	}

	keys = cache.Keys(true)
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
		amsterdam,
		brouwerijtIJ,
		googleHQ,
	}

	// anitkabir: 6, work: 4, ams: 3, googlehq: 2, brouwerij: 2
	// googlehq and brouwerij switches positions because we access google hq later..
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

	// cache.PrintCache()

	// anitkabir: 6, work: 4, ams: 3, googlehq: 2, brouwerij: 2
	deleted := cache.Delete(facebookLondonKey)
	if deleted == true {
		t.Errorf("shouldn't have deleted facebookLondon key from the cache because it shouldn't exist")
	}

	// anitkabir: 6, work: 4, ams: 3, brouwerij: 2
	deleted = cache.Delete(googleHQKey)
	if deleted != true {
		t.Errorf("should have deleted googlehq key from the cache because it should be in the cache before deleting")
	}

	cache.PrintCache()

	// anitkabir: 6, work: 4, ams: 3, brouwerij: 2
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

	// cache.PrintCache()

}
