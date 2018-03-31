package lrfucache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Position struct {
	Lat float64
	Lon float64
}

const BookingHQKey = "thebank"

var BookingHQ = Position{Lat: 52.3655546, Lon: 4.896351}

const AmsterdamKey = "ams"

var Amsterdam = Position{Lat: 52.3702, Lon: 4.8952}

const METUCENGKey = "metuceng"

var METUCENG = Position{Lat: 39.891839, Lon: 32.7811584}

const GoogleHQKey = "googlehq"

var GoogleHQ = Position{Lat: 37.4220, Lon: 122.0841}

const AnitkabirKey = "atam_izindeyiz"

var Anitkabir = Position{Lat: 39.925054, Lon: 32.8347552}

const BrouwerijtIJKey = "coolest_bar_near_my_aparment"

var BrouwerijtIJ = Position{Lat: 52.3677145, Lon: 4.9210217}

const WorkLocationKey = "work"

var UPOffice = Position{Lat: 52.3775935, Lon: 4.9140735}

const FacebookLondonKey = "fb-london"

var FacebookLondon = Position{Lat: 51.5167849, Lon: -0.136269}

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
	position, ok = cache.Get(AmsterdamKey)
	if ok == true || position != nil {
		t.Errorf("should not get Amsterdam's location from empty cache")
	}

	// Setting ams in the cache
	evicted = cache.Set(AmsterdamKey, Amsterdam)
	if evicted {
		t.Errorf("setting amsterdam should not cause eviction")
	}

	// Getting ams back from the cache
	position, ok = cache.Get(AmsterdamKey)
	if ok != true || position != Amsterdam {
		t.Errorf("should get Amsterdam's location from the cache")
	}

	evicted = cache.Set(BookingHQKey, BookingHQ)
	if evicted {
		t.Errorf("setting BookingHQ should not cause eviction")
	}

	evicted = cache.Set(AnitkabirKey, Anitkabir)
	if evicted {
		t.Errorf("setting Anitkabir should not cause eviction")
	}

	evicted = cache.Set(METUCENGKey, METUCENG)
	if evicted {
		t.Errorf("setting metu ceng should not cause eviction")
	}

	evicted = cache.Set(GoogleHQKey, GoogleHQ)
	if evicted {
		t.Errorf("setting googlehq should not cause eviction")
	}

	// google, metuceng, anitkabir, booking, amsterdam
	// cache.PrintCache()

	// brouwerijt, google, metuceng, anitkabir, booking
	evicted = cache.Set(BrouwerijtIJKey, BrouwerijtIJ)
	if evicted != true {
		t.Errorf("setting BrouwerijtIJ should cause eviction")
	}

	position, ok = cache.Get(AmsterdamKey)
	if ok == true || position != nil {
		t.Errorf("should not get Amsterdam's location from the cache, it should have been evicted")
	}

	// work, brouwerijt, google, metuceng, anitkabir
	evicted = cache.Set(WorkLocationKey, UPOffice)
	if evicted != true {
		t.Errorf("setting work should cause eviction")
	}

	position, ok = cache.Get(BookingHQKey)
	if ok == true || position != nil {
		t.Errorf("should not get thebank's location from the cache, it should have been evicted")
	}

	cache.PrintCache()

	// This key is quite special and should never ever be evicted :)
	// Call get to make sure to move it to the head.
	// anitkabir, work, brouwerijt, google, metuceng
	cache.Get(AnitkabirKey)

	// work, anitkabir, brouwerijt, google, metuceng
	// longer commute :(
	evicted = cache.Set(WorkLocationKey, BookingHQ)
	if evicted {
		t.Errorf("no need to evict anything, updating existing key work in the cache")
	}

	coolOffices := []Position{
		FacebookLondon,
		GoogleHQ,
		BookingHQ,
	}
	rand.Seed(time.Now().UnixNano())
	nextLocation := coolOffices[rand.Intn(len(coolOffices)-1)]
	fmt.Printf("NEXT: %v\n", nextLocation)
	cache.Set(WorkLocationKey, nextLocation)

	cache.Get(AnitkabirKey)
	desiredKeys := []string{
		AnitkabirKey,
		WorkLocationKey,
		BrouwerijtIJKey,
		GoogleHQKey,
		METUCENGKey,
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

	desiredLocations := []Position{
		Anitkabir,
		nextLocation,
		BrouwerijtIJ,
		GoogleHQ,
		METUCENG,
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

	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("cache clear failed")
	}

	cache.PrintCache()
}
