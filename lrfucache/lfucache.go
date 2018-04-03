package lrfucache

import (
	"fmt"

	"github.com/bariyu/go-playground/linkedlist"
	"github.com/fatih/color"
)

type lFUCacheFreqNode struct {
	frequency int
	entryList *linkedlist.LinkedList // linked list of lFUCacheEntry's
}

type lFUCacheEntry struct {
	key              interface{}
	value            interface{}
	holderFreqLLNode *linkedlist.Node // wrapper node of lFUCacheFreqNode which holds this cache entry
	llNode           *linkedlist.Node // linkedList Node wraps this entry in the entry list of lFUCacheFreqNode
}

type LFUCache struct {
	lookup   map[interface{}]*lFUCacheEntry
	list     *linkedlist.LinkedList
	capacity int
	size     int
}

func NewLFUCache(capacity int) (cache *LFUCache) {
	list := linkedlist.New() // linked list of lFUCacheFreqNode's
	lookup := make(map[interface{}]*lFUCacheEntry)
	cache = &LFUCache{lookup: lookup, list: list, capacity: capacity, size: 0}
	return cache
}

func (cache *LFUCache) Set(key, value interface{}) (evicted bool) {
	evicted = false
	cacheEntry, ok := cache.lookup[key]
	if ok {
		cache.increaseFrequency(cacheEntry)
		cacheEntry.value = value
	} else {
		if cache.Size() == cache.Capacity() {
			evicted = true

			// last freq node should evict an item
			llFreqNodeToTrim := cache.list.Tail()
			freqNodeToTrim := llFreqNodeToTrim.Value.(*lFUCacheFreqNode)
			keyToDelete := freqNodeToTrim.entryList.Tail().Value.(*lFUCacheEntry).key

			// delete evicted key from lookup
			delete(cache.lookup, keyToDelete)

			// remove last element from freq node
			freqNodeToTrim.entryList.RemoveTail()

			// if last freq node left with 0 elements after removing one cacheEntry
			// remove that freq node from cache's list
			if freqNodeToTrim.entryList.Len() == 0 {
				cache.list.RemoveTail()
			}
		}
		newCacheEntry := &lFUCacheEntry{key: key, value: value}
		tailCacheLLFreqNode := cache.list.Tail()
		// no tail freq element or tail has freq > 1 create a new one..
		if tailCacheLLFreqNode == nil || tailCacheLLFreqNode.Value.(*lFUCacheFreqNode).frequency != 1 {
			// new entry list for new freq node
			newLFUCacheEntryList := linkedlist.New()
			newLFUCacheEntryList.PushFront(newCacheEntry)

			// new lfu cache node with new freq
			tailCacheFreqNode := &lFUCacheFreqNode{frequency: 1, entryList: newLFUCacheEntryList}

			// new ll freq cache node should be the last one
			tailCacheLLFreqNode = cache.list.PushBack(tailCacheFreqNode)

			// update pointers in the new cache entry
			newCacheEntry.holderFreqLLNode = tailCacheLLFreqNode
			newCacheEntry.llNode = newLFUCacheEntryList.Head()
		} else {
			// Get tail freq cache node
			tailCacheFreqNode := tailCacheLLFreqNode.Value.(*lFUCacheFreqNode)

			// push this cache entry to tail freq node's front and update it's ll pointer
			newCacheEntry.llNode = tailCacheFreqNode.entryList.PushFront(newCacheEntry)

			// update holdfreqll node for new cache entry
			newCacheEntry.holderFreqLLNode = tailCacheLLFreqNode
		}

		// increase size of the cache if no eviction occurs
		if evicted == false {
			cache.size += 1
		}

		cache.lookup[key] = newCacheEntry
	}
	return evicted
}

// Increase frequencey of the given cache entry
// Change cacheNode if required or create new one with new freqy
func (cache *LFUCache) increaseFrequency(cacheEntry *lFUCacheEntry) {
	cacheLLFreqNode := cacheEntry.holderFreqLLNode
	prevLLCacheFreqNode := cacheEntry.holderFreqLLNode.Prev()
	cacheFreqNode := cacheLLFreqNode.Value.(*lFUCacheFreqNode)
	prevCacheFreqNode := prevLLCacheFreqNode.Value.(*lFUCacheFreqNode)

	currentFreq := cacheFreqNode.frequency
	newFreq := currentFreq + 1

	oldCacheEntryllNode := cacheEntry.llNode

	if prevCacheFreqNode.frequency == newFreq {
		// move this entry to prev cacheNode and update node pointer
		cacheEntry.llNode = prevCacheFreqNode.entryList.PushFront(cacheEntry)

		// this cache entry is now in prev ll node
		cacheEntry.holderFreqLLNode = prevLLCacheFreqNode
	} else {
		// new entry list for new freq node
		newLFUCacheEntryList := linkedlist.New()
		newLFUCacheEntryList.PushFront(cacheEntry)

		// new lfu cache node with new freq
		newLFUCacheFreqNode := &lFUCacheFreqNode{frequency: newFreq, entryList: newLFUCacheEntryList}

		// new ll cache node should come before current holder node of cacheEntry
		newLLCacheFreqNode := cache.list.PushBefore(newLFUCacheFreqNode, cacheEntry.holderFreqLLNode)

		// update pointers in cache entry
		cacheEntry.holderFreqLLNode = newLLCacheFreqNode
		cacheEntry.llNode = newLFUCacheEntryList.Head()
	}

	// the cacheNode holding keys with this frequency list has size 1
	// should be deleted since no other keys with this frequency left
	if cacheFreqNode.entryList.Len() == 1 {
		cache.list.RemoveNode(cacheLLFreqNode)
	} else {
		cacheFreqNode.entryList.RemoveNode(oldCacheEntryllNode)
	}
}

func (cache *LFUCache) Get(key interface{}) (value interface{}, ok bool) {
	cacheEntry, ok := cache.lookup[key]
	if ok {
		cache.increaseFrequency(cacheEntry)

		return cacheEntry.value, ok
	}
	return nil, ok
}

func (cache *LFUCache) Delete(key interface{}) (deleted bool) {
	deleted = false
	cacheEntry, ok := cache.lookup[key]
	if ok == true {
		deleted = true
		delete(cache.lookup, key)
		cacheLLFreqNode := cacheEntry.holderFreqLLNode
		cacheFreqNode := cacheLLFreqNode.Value.(*lFUCacheFreqNode)

		// the cacheNode holding that keys frequency list has size 1
		// should be deleted since we are deleting this key..
		if cacheFreqNode.entryList.Len() == 1 {
			cache.list.RemoveNode(cacheEntry.holderFreqLLNode)
		} else {
			cacheFreqNode.entryList.RemoveNode(cacheEntry.llNode)
		}
		cache.size -= 1
	}
	return deleted
}

func (cache *LFUCache) Contains(key interface{}) (ok bool) {
	_, ok = cache.lookup[key]
	return ok
}

func (cache *LFUCache) Keys(newest bool) (keys []interface{}) {
	keys = make([]interface{}, 0)

	runner := cache.list.Head()
	if newest == false {
		runner = cache.list.Tail()
	}

	for i := 0; i < cache.list.Len(); i++ {
		cacheNodeValue := runner.Value
		cacheFreqNode := cacheNodeValue.(*lFUCacheFreqNode)
		entryListRunner := cacheFreqNode.entryList.Head()
		if entryListRunner == nil {
			continue
		}
		if newest == false {
			entryListRunner = cacheFreqNode.entryList.Tail()
		}
		for j := 0; j < cacheFreqNode.entryList.Len(); j++ {
			cacheEntryListValue := entryListRunner.Value
			cacheEntry := cacheEntryListValue.(*lFUCacheEntry)
			keys = append(keys, cacheEntry.key)
			if newest {
				entryListRunner = entryListRunner.Next()
			} else {
				entryListRunner = entryListRunner.Prev()
			}
		}
		if newest == false {
			runner = runner.Prev()
		} else {
			runner = runner.Next()
		}
	}
	return keys
}

func (cache *LFUCache) Values(newest bool) (values []interface{}) {
	values = make([]interface{}, 0)

	runner := cache.list.Head()
	if newest == false {
		runner = cache.list.Tail()
	}

	for i := 0; i < cache.list.Len(); i++ {
		cacheNodeValue := runner.Value
		cacheFreqNode := cacheNodeValue.(*lFUCacheFreqNode)
		entryListRunner := cacheFreqNode.entryList.Head()
		if entryListRunner == nil {
			continue
		}
		if newest == false {
			entryListRunner = cacheFreqNode.entryList.Tail()
		}
		for j := 0; j < cacheFreqNode.entryList.Len(); j++ {
			cacheEntryListValue := entryListRunner.Value
			cacheEntry := cacheEntryListValue.(*lFUCacheEntry)
			values = append(values, cacheEntry.value)
			if newest {
				entryListRunner = entryListRunner.Next()
			} else {
				entryListRunner = entryListRunner.Prev()
			}
		}
		if newest == false {
			runner = runner.Prev()
		} else {
			runner = runner.Next()
		}
	}
	return values
}

func (cache *LFUCache) Enumerate(newest bool) (keys, values []interface{}) {
	return cache.Keys(newest), cache.Values(newest)
}

func (cache *LFUCache) Size() (size int) {
	return cache.size
}

func (cache *LFUCache) Capacity() (capacity int) {
	return cache.capacity
}

func (cache *LFUCache) Clear() {
	cache.lookup = make(map[interface{}]*lFUCacheEntry)
	cache.size = 0
	cache.list.Clear()
	return
}

func (cache *LFUCache) PrintCache() {
	fmt.Println("---LFUCACHE START---")
	fmt.Printf("Capacity: %v, Size: %v\n", cache.Capacity(), cache.Size())
	fmt.Printf("Keys %v, Values: %v\n", cache.Keys(true), cache.Values(true))
	fmt.Printf("Lookup: %v\n", cache.lookup)

	fmt.Println("\n---ACTUAL DATA STRUCTURES---")

	cacheEntryPrinter := func(value interface{}) {
		c := color.New(color.FgBlue) //.Add(color.BgBlue)
		entryToPrint := value.(*lFUCacheEntry)
		c.Printf("[key: %v, value %v] --> ", entryToPrint.key, entryToPrint.value)
	}

	freqNodePrinter := func(value interface{}) {
		freqNodeToPrint := value.(*lFUCacheFreqNode)
		color.Green("-FNode- freq: %d, entryList with %d elems:", freqNodeToPrint.frequency, freqNodeToPrint.entryList.Len())
		freqNodeToPrint.entryList.PrintList(cacheEntryPrinter)
		color.Red("nil\n")
	}

	cache.list.PrintList(freqNodePrinter)

	fmt.Println("---END ACTUAL DATA STRUCTURES---")
	fmt.Println()

	fmt.Println("---LFUCACHE END---")
}
