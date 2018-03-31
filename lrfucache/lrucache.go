package lrfucache

import (
	"fmt"

	"github.com/bariyu/go-playground/linkedlist"
)

type LRUCacheEntry struct {
	key   interface{}
	value interface{}
}

type LRUCache struct {
	lookup   map[interface{}]*linkedlist.Node
	list     *linkedlist.LinkedList
	capacity int
}

func NewLRUCache(capacity int) (cache *LRUCache) {
	list := linkedlist.New()
	lookup := make(map[interface{}]*linkedlist.Node)
	cache = &LRUCache{lookup: lookup, list: list, capacity: capacity}
	return cache
}

func (cache *LRUCache) Set(key, value interface{}) (evicted bool) {
	evicted = false
	listNode, ok := cache.lookup[key]
	if ok {
		cache.list.MoveFront(listNode)
		cacheEntry := listNode.Value
		entry := cacheEntry.(*LRUCacheEntry)
		entry.value = value
	} else {
		if cache.Size() == cache.Capacity() {
			evicted = true
			listNodeToEvict := cache.list.Tail()
			entry := listNodeToEvict.Value.(*LRUCacheEntry)
			delete(cache.lookup, entry.key)
			cache.list.RemoveNode(listNodeToEvict)
		}
		newLRUCacheEntry := &LRUCacheEntry{key: key, value: value}
		newListNode := cache.list.PushFront(newLRUCacheEntry)
		cache.lookup[key] = newListNode
	}
	return evicted
}

func (cache *LRUCache) Get(key interface{}) (value interface{}, ok bool) {
	node, ok := cache.lookup[key]
	if ok {
		cacheEntry := node.Value
		cache.list.MoveFront(node)
		entry := cacheEntry.(*LRUCacheEntry)
		return entry.value, ok
	}
	return nil, ok
}

func (cache *LRUCache) Delete(key interface{}) (deleted bool) {
	deleted = false
	node, ok := cache.lookup[key]
	if ok == true {
		deleted = true
		delete(cache.lookup, key)
		cache.list.RemoveNode(node)
	}
	return deleted
}

func (cache *LRUCache) Contains(key interface{}) (ok bool) {
	_, ok = cache.lookup[key]
	return ok
}

func (cache *LRUCache) Keys(newest bool) (keys []interface{}) {
	keys = make([]interface{}, 0)

	runner := cache.list.Head()
	if newest == false {
		runner = cache.list.Tail()
	}

	for i := 0; i < cache.Size(); i++ {
		cacheEntry := runner.Value
		entry := cacheEntry.(*LRUCacheEntry)
		keys = append(keys, entry.key)
		if newest {
			runner = runner.Next()
		} else {
			runner = runner.Prev()
		}
	}
	return keys
}

func (cache *LRUCache) Values(newest bool) (values []interface{}) {
	values = make([]interface{}, 0)

	runner := cache.list.Head()
	if newest == false {
		runner = cache.list.Tail()
	}

	for i := 0; i < cache.Size(); i++ {
		cacheEntry := runner.Value
		entry := cacheEntry.(*LRUCacheEntry)
		values = append(values, entry.value)
		if newest {
			runner = runner.Next()
		} else {
			runner = runner.Prev()
		}
	}
	return values
}

func (cache *LRUCache) Enumerate(newest bool) (keys, values []interface{}) {
	return cache.Keys(newest), cache.Values(newest)
}

func (cache *LRUCache) Size() (size int) {
	return cache.list.Len()
}

func (cache *LRUCache) Capacity() (capacity int) {
	return cache.capacity
}

func (cache *LRUCache) Clear() {
	cache.lookup = make(map[interface{}]*linkedlist.Node)
	cache.list.Clear()
	return
}

func (cache *LRUCache) PrintCache() {
	fmt.Println("---Start---")
	fmt.Printf("Capacity: %v, Size: %v\n", cache.Capacity(), cache.Size())
	fmt.Printf("Lookup: %v\n", cache.lookup)
	fmt.Println("---End---")
}
